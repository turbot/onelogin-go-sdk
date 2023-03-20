/*
OneLogin API

Testing SmartHooksApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package onelogin

import (
	"context"
	"testing"

	openapiclient "github.com/onelogin/onelogin-go-sdk"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_onelogin_SmartHooksApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test SmartHooksApiService CreateEnvironmentVariable", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.SmartHooksApi.CreateEnvironmentVariable(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SmartHooksApiService CreateHook", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.SmartHooksApi.CreateHook(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SmartHooksApiService DeleteEnvironmentVariable", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var envvarId string

		httpRes, err := apiClient.SmartHooksApi.DeleteEnvironmentVariable(context.Background(), envvarId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SmartHooksApiService DeleteHook", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var hookId string

		httpRes, err := apiClient.SmartHooksApi.DeleteHook(context.Background(), hookId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SmartHooksApiService GetEnvironmentVariable", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var envvarId string

		resp, httpRes, err := apiClient.SmartHooksApi.GetEnvironmentVariable(context.Background(), envvarId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SmartHooksApiService GetHook", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var hookId string

		resp, httpRes, err := apiClient.SmartHooksApi.GetHook(context.Background(), hookId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SmartHooksApiService GetLogs", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var hookId string

		resp, httpRes, err := apiClient.SmartHooksApi.GetLogs(context.Background(), hookId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SmartHooksApiService ListEnvironmentVariables", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.SmartHooksApi.ListEnvironmentVariables(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SmartHooksApiService ListHooks", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.SmartHooksApi.ListHooks(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SmartHooksApiService UpdateEnvironmentVariable", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var envvarId string

		resp, httpRes, err := apiClient.SmartHooksApi.UpdateEnvironmentVariable(context.Background(), envvarId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SmartHooksApiService UpdateHook", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var hookId string

		resp, httpRes, err := apiClient.SmartHooksApi.UpdateHook(context.Background(), hookId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}