/*
EasyAPI REST Services API

Testing MailAPIService

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

func Test_openapi_MailAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test MailAPIService CreateMailmap", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var domain string

		resp, httpRes, err := apiClient.MailAPI.CreateMailmap(context.Background(), domain).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MailAPIService DeleteMailmap", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var domain string
		var mailmapId int32

		resp, httpRes, err := apiClient.MailAPI.DeleteMailmap(context.Background(), domain, mailmapId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MailAPIService ListMailmaps", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var domain string

		resp, httpRes, err := apiClient.MailAPI.ListMailmaps(context.Background(), domain).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MailAPIService UpdateMailmap", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var domain string
		var email string

		resp, httpRes, err := apiClient.MailAPI.UpdateMailmap(context.Background(), domain, email).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}