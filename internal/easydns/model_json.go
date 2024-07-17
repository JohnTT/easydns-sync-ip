package easydns

type ZoneRecord struct {
	ID        string      `json:"id"`
	Domain    string      `json:"domain"`
	Host      string      `json:"host"`
	TTL       string      `json:"ttl"`
	Prio      interface{} `json:"prio"`
	Type      string      `json:"type"`
	Rdata     string      `json:"rdata"`
	GeozoneID string      `json:"geozone_id"`
	LastMod   string      `json:"last_mod"`
}

type ListZoneResponse struct {
	TM     int          `json:"tm"`
	Data   []ZoneRecord `json:"data"`
	Count  int          `json:"count"`
	Total  int          `json:"total"`
	Start  int          `json:"start"`
	Max    int          `json:"max"`
	Status int          `json:"status"`
}
