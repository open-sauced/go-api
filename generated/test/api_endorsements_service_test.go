/*
@open-sauced/api.opensauced.pizza

Testing EndorsementsServiceAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package client

import (
	"context"
	"testing"

	openapiclient "github.com/open-sauced/go-api"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_openapi_EndorsementsServiceAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test EndorsementsServiceAPIService CreateEndorsement", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.EndorsementsServiceAPI.CreateEndorsement(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EndorsementsServiceAPIService DeleteEndoresementById", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		httpRes, err := apiClient.EndorsementsServiceAPI.DeleteEndoresementById(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EndorsementsServiceAPIService FindAllByRepoOwnerOrUsername", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var repoOwnerOrUser string

		resp, httpRes, err := apiClient.EndorsementsServiceAPI.FindAllByRepoOwnerOrUsername(context.Background(), repoOwnerOrUser).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EndorsementsServiceAPIService FindAllEndorsements", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.EndorsementsServiceAPI.FindAllEndorsements(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EndorsementsServiceAPIService FindAllEndorsementsByRepo", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var owner string
		var repo string

		resp, httpRes, err := apiClient.EndorsementsServiceAPI.FindAllEndorsementsByRepo(context.Background(), owner, repo).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EndorsementsServiceAPIService FindAllUserCreatedEndorsements", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.EndorsementsServiceAPI.FindAllUserCreatedEndorsements(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EndorsementsServiceAPIService FindAllUserReceivedEndorsements", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.EndorsementsServiceAPI.FindAllUserReceivedEndorsements(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EndorsementsServiceAPIService FindAllUserReceivedEndorsements_1", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var username string

		resp, httpRes, err := apiClient.EndorsementsServiceAPI.FindAllUserReceivedEndorsements_1(context.Background(), username).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EndorsementsServiceAPIService FindAllUserReceivedEndorsements_2", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var username string

		resp, httpRes, err := apiClient.EndorsementsServiceAPI.FindAllUserReceivedEndorsements_2(context.Background(), username).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EndorsementsServiceAPIService FindEndorsementById", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.EndorsementsServiceAPI.FindEndorsementById(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
