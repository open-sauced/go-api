/*
@open-sauced/api.opensauced.pizza

Testing UserServiceAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package client

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/open-sauced/go-api"
)

func Test_openapi_UserServiceAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test UserServiceAPIService FindAllHighlightsByUsername", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var username string

		resp, httpRes, err := apiClient.UserServiceAPI.FindAllHighlightsByUsername(context.Background(), username).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UserServiceAPIService FindAllTopReposByUsername", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var username string

		resp, httpRes, err := apiClient.UserServiceAPI.FindAllTopReposByUsername(context.Background(), username).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UserServiceAPIService FindContributorPullRequests", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var username string

		resp, httpRes, err := apiClient.UserServiceAPI.FindContributorPullRequests(context.Background(), username).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UserServiceAPIService FindOneUserByUserame", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var username string

		resp, httpRes, err := apiClient.UserServiceAPI.FindOneUserByUserame(context.Background(), username).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UserServiceAPIService FollowUserById", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var username string

		resp, httpRes, err := apiClient.UserServiceAPI.FollowUserById(context.Background(), username).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UserServiceAPIService GetFollowStatusByUsername", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var username string

		resp, httpRes, err := apiClient.UserServiceAPI.GetFollowStatusByUsername(context.Background(), username).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UserServiceAPIService GetTop10Highlights", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.UserServiceAPI.GetTop10Highlights(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UserServiceAPIService GetUserNotifications", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.UserServiceAPI.GetUserNotifications(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UserServiceAPIService UnfollowUserByUsername", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var username string

		resp, httpRes, err := apiClient.UserServiceAPI.UnfollowUserByUsername(context.Background(), username).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
