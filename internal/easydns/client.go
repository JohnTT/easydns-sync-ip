package easydns

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"strconv"
	"time"

	"github.com/JohnTT/easydns-sync-ip/internal/config"
	openapi "github.com/JohnTT/easydns-sync-ip/openapi/client/go"
)

type Client struct {
	basicAuthHttpClient *http.Client
	*openapi.APIClient
}

func NewClient() *Client {
	cfg := openapi.NewConfiguration()
	cfg.Debug = false

	// Create an HTTP client with the custom BasicAuthTransport
	cfg.HTTPClient = &http.Client{
		Transport: &BasicAuthTransport{
			Username:  config.Get().Token,
			Password:  config.Get().Key,
			Transport: http.DefaultTransport, // Use the default transport for actual HTTP requests
		},
	}

	// Override the server configurations from env
	cfg.Servers = openapi.ServerConfigurations{
		{
			URL: config.Get().Host,
		},
	}

	apiClient := openapi.NewAPIClient(cfg)
	return &Client{
		basicAuthHttpClient: cfg.HTTPClient,
		APIClient:           apiClient,
	}
}

// Update checks if the DNS record's host matches the public IP of this computer.
// If there is a mismatch, update the DNS record in EasyDNS.
func (c *Client) Update() error {
	ip, err := c.getPublicIp()
	if err != nil {
		return err
	}
	log.Printf("public IP address is %+v", ip)

	zoneRecord, err := c.findZoneRecord(config.Get().RecordHost)
	if err != nil {
		return fmt.Errorf("could not find zone record: %w", err)
	}
	log.Printf("found zone record %+v", zoneRecord)

	// Done, since our public IP address matches the DNS record
	if ip == zoneRecord.Rdata {
		return nil
	}

	// Otherwise, update the DNS record with the new public IP address
	// Sleep for a couple seconds to avoid being rate limited.
	time.Sleep(2 * time.Second)
	zoneRecord.Rdata = ip
	if err := c.updateZoneRecord(zoneRecord); err != nil {
		return err
	}

	return nil
}

func (c *Client) findZoneRecord(host string) (ZoneRecord, error) {
	// OpenAPI client cannot parse body correctly so it returns an error.
	// Ignore the error and parse the body manually
	domain := config.Get().Domain
	_, resp, err := c.ZoneAPI.ListZone(context.TODO(), config.Get().Domain).Execute()
	if resp == nil && err != nil {
		return ZoneRecord{}, err
	}
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return ZoneRecord{}, fmt.Errorf("response did not contain 2xx code: %+v", resp)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return ZoneRecord{}, err
	}

	listZoneResponse := ListZoneResponse{}
	err = json.Unmarshal([]byte(data), &listZoneResponse)
	if err != nil {
		return ZoneRecord{}, err
	}

	for _, record := range listZoneResponse.Data {
		if record.Host == host {
			return record, nil
		}
	}

	return ZoneRecord{}, fmt.Errorf("domain %s, could not find host %s", domain, host)
}

func (c *Client) getPublicIp() (string, error) {
	resp, err := http.Get("https://ifconfig.me/ip")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	ip := net.ParseIP(string(data))
	if ip == nil || ip.To4() == nil {
		return "", fmt.Errorf("%s is not a valid IPv4 address", string(data))
	}

	return ip.String(), nil
}

func (c *Client) updateZoneRecord(record ZoneRecord) error {
	log.Printf("mismatch detected, updating zone record to %+v", record)

	id, err := strconv.ParseInt(record.ID, 10, 64)
	if err != nil {
		return err
	}

	_, resp, err := c.ZoneAPI.ModZoneRec(context.TODO(), id).RequestModelZoneBodyData(
		openapi.RequestModelZoneBodyData{
			Domain: record.Domain,
			Host:   record.Host,
			Type:   record.Type,
			Rdata:  record.Rdata,
		},
	).Execute()
	if resp == nil && err != nil {
		return err
	}
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("response did not contain 2xx code: %+v", resp)
	}

	return nil
}
