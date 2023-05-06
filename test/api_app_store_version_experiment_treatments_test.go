/*
App Store Connect API

Testing AppStoreVersionExperimentTreatmentsApiService

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

func Test_openapi_AppStoreVersionExperimentTreatmentsApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test AppStoreVersionExperimentTreatmentsApiService AppStoreVersionExperimentTreatmentsAppStoreVersionExperimentTreatmentLocalizationsGetToManyRelated", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.AppStoreVersionExperimentTreatmentsApi.AppStoreVersionExperimentTreatmentsAppStoreVersionExperimentTreatmentLocalizationsGetToManyRelated(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AppStoreVersionExperimentTreatmentsApiService AppStoreVersionExperimentTreatmentsCreateInstance", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.AppStoreVersionExperimentTreatmentsApi.AppStoreVersionExperimentTreatmentsCreateInstance(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AppStoreVersionExperimentTreatmentsApiService AppStoreVersionExperimentTreatmentsDeleteInstance", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		httpRes, err := apiClient.AppStoreVersionExperimentTreatmentsApi.AppStoreVersionExperimentTreatmentsDeleteInstance(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AppStoreVersionExperimentTreatmentsApiService AppStoreVersionExperimentTreatmentsGetInstance", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.AppStoreVersionExperimentTreatmentsApi.AppStoreVersionExperimentTreatmentsGetInstance(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AppStoreVersionExperimentTreatmentsApiService AppStoreVersionExperimentTreatmentsUpdateInstance", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.AppStoreVersionExperimentTreatmentsApi.AppStoreVersionExperimentTreatmentsUpdateInstance(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
