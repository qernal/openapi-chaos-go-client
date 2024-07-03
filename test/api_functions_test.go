/*
Chaos

Testing FunctionsAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package openapi_chaos_client

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/qernal/openapi-chaos-go-client"
)

func Test_openapi_chaos_client_FunctionsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test FunctionsAPIService FunctionsCreate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.FunctionsAPI.FunctionsCreate(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FunctionsAPIService FunctionsDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var functionId string

		resp, httpRes, err := apiClient.FunctionsAPI.FunctionsDelete(context.Background(), functionId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FunctionsAPIService FunctionsGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var functionId string

		resp, httpRes, err := apiClient.FunctionsAPI.FunctionsGet(context.Background(), functionId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FunctionsAPIService FunctionsRevisionsGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var functionId string
		var functionRevisionId string

		resp, httpRes, err := apiClient.FunctionsAPI.FunctionsRevisionsGet(context.Background(), functionId, functionRevisionId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FunctionsAPIService FunctionsRevisionsList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var functionId string

		resp, httpRes, err := apiClient.FunctionsAPI.FunctionsRevisionsList(context.Background(), functionId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FunctionsAPIService FunctionsUpdate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var functionId string

		resp, httpRes, err := apiClient.FunctionsAPI.FunctionsUpdate(context.Background(), functionId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FunctionsAPIService ProjectsFunctionsList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var projectId string

		resp, httpRes, err := apiClient.FunctionsAPI.ProjectsFunctionsList(context.Background(), projectId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
