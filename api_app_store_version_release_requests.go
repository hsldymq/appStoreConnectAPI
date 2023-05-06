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


// AppStoreVersionReleaseRequestsApiService AppStoreVersionReleaseRequestsApi service
type AppStoreVersionReleaseRequestsApiService service

type ApiAppStoreVersionReleaseRequestsCreateInstanceRequest struct {
	ctx context.Context
	ApiService *AppStoreVersionReleaseRequestsApiService
	appStoreVersionReleaseRequestCreateRequest *AppStoreVersionReleaseRequestCreateRequest
}

// AppStoreVersionReleaseRequest representation
func (r ApiAppStoreVersionReleaseRequestsCreateInstanceRequest) AppStoreVersionReleaseRequestCreateRequest(appStoreVersionReleaseRequestCreateRequest AppStoreVersionReleaseRequestCreateRequest) ApiAppStoreVersionReleaseRequestsCreateInstanceRequest {
	r.appStoreVersionReleaseRequestCreateRequest = &appStoreVersionReleaseRequestCreateRequest
	return r
}

func (r ApiAppStoreVersionReleaseRequestsCreateInstanceRequest) Execute() (*AppStoreVersionReleaseRequestResponse, *http.Response, error) {
	return r.ApiService.AppStoreVersionReleaseRequestsCreateInstanceExecute(r)
}

/*
AppStoreVersionReleaseRequestsCreateInstance Method for AppStoreVersionReleaseRequestsCreateInstance

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiAppStoreVersionReleaseRequestsCreateInstanceRequest
*/
func (a *AppStoreVersionReleaseRequestsApiService) AppStoreVersionReleaseRequestsCreateInstance(ctx context.Context) ApiAppStoreVersionReleaseRequestsCreateInstanceRequest {
	return ApiAppStoreVersionReleaseRequestsCreateInstanceRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return AppStoreVersionReleaseRequestResponse
func (a *AppStoreVersionReleaseRequestsApiService) AppStoreVersionReleaseRequestsCreateInstanceExecute(r ApiAppStoreVersionReleaseRequestsCreateInstanceRequest) (*AppStoreVersionReleaseRequestResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AppStoreVersionReleaseRequestResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AppStoreVersionReleaseRequestsApiService.AppStoreVersionReleaseRequestsCreateInstance")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/appStoreVersionReleaseRequests"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.appStoreVersionReleaseRequestCreateRequest == nil {
		return localVarReturnValue, nil, reportError("appStoreVersionReleaseRequestCreateRequest is required and must be specified")
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
	localVarPostBody = r.appStoreVersionReleaseRequestCreateRequest
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
