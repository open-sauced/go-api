/*
@open-sauced/api.opensauced.pizza

Testing PullRequestsServiceAPIService

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

func Test_openapi_PullRequestsServiceAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test PullRequestsServiceAPIService GenerateCodeExplanation", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.PullRequestsServiceAPI.GenerateCodeExplanation(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PullRequestsServiceAPIService GenerateCodeRefactor", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.PullRequestsServiceAPI.GenerateCodeRefactor(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PullRequestsServiceAPIService GenerateCodeTest", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.PullRequestsServiceAPI.GenerateCodeTest(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PullRequestsServiceAPIService GeneratePRDescription", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.PullRequestsServiceAPI.GeneratePRDescription(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PullRequestsServiceAPIService GetPullRequestInsights", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.PullRequestsServiceAPI.GetPullRequestInsights(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PullRequestsServiceAPIService ListAllPullRequests", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.PullRequestsServiceAPI.ListAllPullRequests(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PullRequestsServiceAPIService SearchAllPullRequests", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.PullRequestsServiceAPI.SearchAllPullRequests(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
