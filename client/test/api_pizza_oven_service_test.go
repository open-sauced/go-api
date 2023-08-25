/*
@open-sauced/api.opensauced.pizza

Testing PizzaOvenServiceAPIService

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

func Test_openapi_PizzaOvenServiceAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test PizzaOvenServiceAPIService BakeARepositoryWithThePizzaOvenMicroservice", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.PizzaOvenServiceAPI.BakeARepositoryWithThePizzaOvenMicroservice(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PizzaOvenServiceAPIService FindBakedRepoById", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id int32

		resp, httpRes, err := apiClient.PizzaOvenServiceAPI.FindBakedRepoById(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PizzaOvenServiceAPIService FindCommitAuthorById", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id int32

		resp, httpRes, err := apiClient.PizzaOvenServiceAPI.FindCommitAuthorById(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PizzaOvenServiceAPIService FindCommitById", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id int32

		resp, httpRes, err := apiClient.PizzaOvenServiceAPI.FindCommitById(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PizzaOvenServiceAPIService ListAllBakedRepos", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.PizzaOvenServiceAPI.ListAllBakedRepos(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PizzaOvenServiceAPIService ListAllCommitAuthors", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.PizzaOvenServiceAPI.ListAllCommitAuthors(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PizzaOvenServiceAPIService ListAllCommitsByBakedRepoId", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id int32

		resp, httpRes, err := apiClient.PizzaOvenServiceAPI.ListAllCommitsByBakedRepoId(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PizzaOvenServiceAPIService ListAllCommitsByCommitAuthorId", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id int32

		resp, httpRes, err := apiClient.PizzaOvenServiceAPI.ListAllCommitsByCommitAuthorId(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
