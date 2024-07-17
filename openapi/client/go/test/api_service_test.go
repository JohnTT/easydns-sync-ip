/*
EasyAPI REST Services API

Testing ServiceAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package openapi

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func Test_openapi_ServiceAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ServiceAPIService GetServiceDescription", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var serviceId int32

		resp, httpRes, err := apiClient.ServiceAPI.GetServiceDescription(context.Background(), serviceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ServiceAPIService GetSubscriptionServiceDescription", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var subscriptionId int32

		resp, httpRes, err := apiClient.ServiceAPI.GetSubscriptionServiceDescription(context.Background(), subscriptionId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
