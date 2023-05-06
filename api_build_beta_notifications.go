/*
App Store Connect API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)


// BuildBetaNotificationsApiService BuildBetaNotificationsApi service
type BuildBetaNotificationsApiService service

type ApiBuildBetaNotificationsCreateInstanceRequest struct {
	ctx context.Context
	ApiService *BuildBetaNotificationsApiService
	buildBetaNotificationCreateRequest *BuildBetaNotificationCreateRequest
}

// BuildBetaNotification representation
func (r ApiBuildBetaNotificationsCreateInstanceRequest) BuildBetaNotificationCreateRequest(buildBetaNotificationCreateRequest BuildBetaNotificationCreateRequest) ApiBuildBetaNotificationsCreateInstanceRequest {
	r.buildBetaNotificationCreateRequest = &buildBetaNotificationCreateRequest
	return r
}

func (r ApiBuildBetaNotificationsCreateInstanceRequest) Execute() (*BuildBetaNotificationResponse, *http.Response, error) {
	return r.ApiService.BuildBetaNotificationsCreateInstanceExecute(r)
}

/*
BuildBetaNotificationsCreateInstance Method for BuildBetaNotificationsCreateInstance

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiBuildBetaNotificationsCreateInstanceRequest
*/
func (a *BuildBetaNotificationsApiService) BuildBetaNotificationsCreateInstance(ctx context.Context) ApiBuildBetaNotificationsCreateInstanceRequest {
	return ApiBuildBetaNotificationsCreateInstanceRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return BuildBetaNotificationResponse
func (a *BuildBetaNotificationsApiService) BuildBetaNotificationsCreateInstanceExecute(r ApiBuildBetaNotificationsCreateInstanceRequest) (*BuildBetaNotificationResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *BuildBetaNotificationResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BuildBetaNotificationsApiService.BuildBetaNotificationsCreateInstance")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/buildBetaNotifications"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.buildBetaNotificationCreateRequest == nil {
		return localVarReturnValue, nil, reportError("buildBetaNotificationCreateRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.buildBetaNotificationCreateRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 409 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
