/*
@open-sauced/api.opensauced.pizza

Testing InsightsServiceAPIService

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

func Test_openapi_InsightsServiceAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test InsightsServiceAPIService AddInsightForUser", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.InsightsServiceAPI.AddInsightForUser(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InsightsServiceAPIService AddMemberForInsight", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id int32

		resp, httpRes, err := apiClient.InsightsServiceAPI.AddMemberForInsight(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InsightsServiceAPIService FindAllInsightMembers", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id int32

		resp, httpRes, err := apiClient.InsightsServiceAPI.FindAllInsightMembers(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InsightsServiceAPIService FindAllInsightsByUserId", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.InsightsServiceAPI.FindAllInsightsByUserId(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InsightsServiceAPIService FindInsightPageById", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id int32

		resp, httpRes, err := apiClient.InsightsServiceAPI.FindInsightPageById(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InsightsServiceAPIService RemoveInsightForUser", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id int32

		resp, httpRes, err := apiClient.InsightsServiceAPI.RemoveInsightForUser(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InsightsServiceAPIService RemoveInsightMemberById", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id int32
		var memberId string

		resp, httpRes, err := apiClient.InsightsServiceAPI.RemoveInsightMemberById(context.Background(), id, memberId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InsightsServiceAPIService UpdateInsightForUser", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id int32

		resp, httpRes, err := apiClient.InsightsServiceAPI.UpdateInsightForUser(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InsightsServiceAPIService UpdateInsightMember", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id int32
		var memberId string

		resp, httpRes, err := apiClient.InsightsServiceAPI.UpdateInsightMember(context.Background(), id, memberId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
