/*
@open-sauced/api.opensauced.pizza

Testing ContributionServiceAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package client

import (
	"context"
	"testing"

	openapiclient "github.com/open-sauced/go-api/client"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_openapi_ContributionServiceAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ContributionServiceAPIService FindAllByOwnerAndRepo", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var owner string
		var repo string

		resp, httpRes, err := apiClient.ContributionServiceAPI.FindAllByOwnerAndRepo(context.Background(), owner, repo).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ContributionServiceAPIService FindAllByRepoId", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id float32

		resp, httpRes, err := apiClient.ContributionServiceAPI.FindAllByRepoId(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}