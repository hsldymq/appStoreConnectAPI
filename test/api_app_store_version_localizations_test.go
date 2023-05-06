/*
App Store Connect API

Testing AppStoreVersionLocalizationsApiService

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

func Test_openapi_AppStoreVersionLocalizationsApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test AppStoreVersionLocalizationsApiService AppStoreVersionLocalizationsAppPreviewSetsGetToManyRelated", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.AppStoreVersionLocalizationsApi.AppStoreVersionLocalizationsAppPreviewSetsGetToManyRelated(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AppStoreVersionLocalizationsApiService AppStoreVersionLocalizationsAppScreenshotSetsGetToManyRelated", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.AppStoreVersionLocalizationsApi.AppStoreVersionLocalizationsAppScreenshotSetsGetToManyRelated(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AppStoreVersionLocalizationsApiService AppStoreVersionLocalizationsCreateInstance", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.AppStoreVersionLocalizationsApi.AppStoreVersionLocalizationsCreateInstance(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AppStoreVersionLocalizationsApiService AppStoreVersionLocalizationsDeleteInstance", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		httpRes, err := apiClient.AppStoreVersionLocalizationsApi.AppStoreVersionLocalizationsDeleteInstance(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AppStoreVersionLocalizationsApiService AppStoreVersionLocalizationsGetInstance", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.AppStoreVersionLocalizationsApi.AppStoreVersionLocalizationsGetInstance(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AppStoreVersionLocalizationsApiService AppStoreVersionLocalizationsUpdateInstance", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.AppStoreVersionLocalizationsApi.AppStoreVersionLocalizationsUpdateInstance(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}