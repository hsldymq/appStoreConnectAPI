/*
App Store Connect API

Testing AppStoreVersionExperimentTreatmentLocalizationsApiService

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

func Test_openapi_AppStoreVersionExperimentTreatmentLocalizationsApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test AppStoreVersionExperimentTreatmentLocalizationsApiService AppStoreVersionExperimentTreatmentLocalizationsAppPreviewSetsGetToManyRelated", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.AppStoreVersionExperimentTreatmentLocalizationsApi.AppStoreVersionExperimentTreatmentLocalizationsAppPreviewSetsGetToManyRelated(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AppStoreVersionExperimentTreatmentLocalizationsApiService AppStoreVersionExperimentTreatmentLocalizationsAppScreenshotSetsGetToManyRelated", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.AppStoreVersionExperimentTreatmentLocalizationsApi.AppStoreVersionExperimentTreatmentLocalizationsAppScreenshotSetsGetToManyRelated(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AppStoreVersionExperimentTreatmentLocalizationsApiService AppStoreVersionExperimentTreatmentLocalizationsCreateInstance", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.AppStoreVersionExperimentTreatmentLocalizationsApi.AppStoreVersionExperimentTreatmentLocalizationsCreateInstance(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AppStoreVersionExperimentTreatmentLocalizationsApiService AppStoreVersionExperimentTreatmentLocalizationsDeleteInstance", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		httpRes, err := apiClient.AppStoreVersionExperimentTreatmentLocalizationsApi.AppStoreVersionExperimentTreatmentLocalizationsDeleteInstance(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AppStoreVersionExperimentTreatmentLocalizationsApiService AppStoreVersionExperimentTreatmentLocalizationsGetInstance", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.AppStoreVersionExperimentTreatmentLocalizationsApi.AppStoreVersionExperimentTreatmentLocalizationsGetInstance(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
