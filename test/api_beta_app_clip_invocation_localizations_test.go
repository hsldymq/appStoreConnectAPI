/*
App Store Connect API

Testing BetaAppClipInvocationLocalizationsApiService

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

func Test_openapi_BetaAppClipInvocationLocalizationsApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test BetaAppClipInvocationLocalizationsApiService BetaAppClipInvocationLocalizationsCreateInstance", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.BetaAppClipInvocationLocalizationsApi.BetaAppClipInvocationLocalizationsCreateInstance(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BetaAppClipInvocationLocalizationsApiService BetaAppClipInvocationLocalizationsDeleteInstance", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		httpRes, err := apiClient.BetaAppClipInvocationLocalizationsApi.BetaAppClipInvocationLocalizationsDeleteInstance(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BetaAppClipInvocationLocalizationsApiService BetaAppClipInvocationLocalizationsUpdateInstance", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.BetaAppClipInvocationLocalizationsApi.BetaAppClipInvocationLocalizationsUpdateInstance(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
