/*
App Store Connect API

Testing InAppPurchaseAppStoreReviewScreenshotsApiService

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

func Test_openapi_InAppPurchaseAppStoreReviewScreenshotsApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test InAppPurchaseAppStoreReviewScreenshotsApiService InAppPurchaseAppStoreReviewScreenshotsCreateInstance", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.InAppPurchaseAppStoreReviewScreenshotsApi.InAppPurchaseAppStoreReviewScreenshotsCreateInstance(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InAppPurchaseAppStoreReviewScreenshotsApiService InAppPurchaseAppStoreReviewScreenshotsDeleteInstance", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		httpRes, err := apiClient.InAppPurchaseAppStoreReviewScreenshotsApi.InAppPurchaseAppStoreReviewScreenshotsDeleteInstance(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InAppPurchaseAppStoreReviewScreenshotsApiService InAppPurchaseAppStoreReviewScreenshotsGetInstance", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.InAppPurchaseAppStoreReviewScreenshotsApi.InAppPurchaseAppStoreReviewScreenshotsGetInstance(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InAppPurchaseAppStoreReviewScreenshotsApiService InAppPurchaseAppStoreReviewScreenshotsUpdateInstance", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.InAppPurchaseAppStoreReviewScreenshotsApi.InAppPurchaseAppStoreReviewScreenshotsUpdateInstance(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
