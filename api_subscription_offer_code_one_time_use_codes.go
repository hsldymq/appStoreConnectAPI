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
	"strings"
)


// SubscriptionOfferCodeOneTimeUseCodesApiService SubscriptionOfferCodeOneTimeUseCodesApi service
type SubscriptionOfferCodeOneTimeUseCodesApiService service

type ApiSubscriptionOfferCodeOneTimeUseCodesCreateInstanceRequest struct {
	ctx context.Context
	ApiService *SubscriptionOfferCodeOneTimeUseCodesApiService
	subscriptionOfferCodeOneTimeUseCodeCreateRequest *SubscriptionOfferCodeOneTimeUseCodeCreateRequest
}

// SubscriptionOfferCodeOneTimeUseCode representation
func (r ApiSubscriptionOfferCodeOneTimeUseCodesCreateInstanceRequest) SubscriptionOfferCodeOneTimeUseCodeCreateRequest(subscriptionOfferCodeOneTimeUseCodeCreateRequest SubscriptionOfferCodeOneTimeUseCodeCreateRequest) ApiSubscriptionOfferCodeOneTimeUseCodesCreateInstanceRequest {
	r.subscriptionOfferCodeOneTimeUseCodeCreateRequest = &subscriptionOfferCodeOneTimeUseCodeCreateRequest
	return r
}

func (r ApiSubscriptionOfferCodeOneTimeUseCodesCreateInstanceRequest) Execute() (*SubscriptionOfferCodeOneTimeUseCodeResponse, *http.Response, error) {
	return r.ApiService.SubscriptionOfferCodeOneTimeUseCodesCreateInstanceExecute(r)
}

/*
SubscriptionOfferCodeOneTimeUseCodesCreateInstance Method for SubscriptionOfferCodeOneTimeUseCodesCreateInstance

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiSubscriptionOfferCodeOneTimeUseCodesCreateInstanceRequest
*/
func (a *SubscriptionOfferCodeOneTimeUseCodesApiService) SubscriptionOfferCodeOneTimeUseCodesCreateInstance(ctx context.Context) ApiSubscriptionOfferCodeOneTimeUseCodesCreateInstanceRequest {
	return ApiSubscriptionOfferCodeOneTimeUseCodesCreateInstanceRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return SubscriptionOfferCodeOneTimeUseCodeResponse
func (a *SubscriptionOfferCodeOneTimeUseCodesApiService) SubscriptionOfferCodeOneTimeUseCodesCreateInstanceExecute(r ApiSubscriptionOfferCodeOneTimeUseCodesCreateInstanceRequest) (*SubscriptionOfferCodeOneTimeUseCodeResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SubscriptionOfferCodeOneTimeUseCodeResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SubscriptionOfferCodeOneTimeUseCodesApiService.SubscriptionOfferCodeOneTimeUseCodesCreateInstance")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/subscriptionOfferCodeOneTimeUseCodes"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.subscriptionOfferCodeOneTimeUseCodeCreateRequest == nil {
		return localVarReturnValue, nil, reportError("subscriptionOfferCodeOneTimeUseCodeCreateRequest is required and must be specified")
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
	localVarPostBody = r.subscriptionOfferCodeOneTimeUseCodeCreateRequest
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

type ApiSubscriptionOfferCodeOneTimeUseCodesGetInstanceRequest struct {
	ctx context.Context
	ApiService *SubscriptionOfferCodeOneTimeUseCodesApiService
	id string
	fieldsSubscriptionOfferCodeOneTimeUseCodes *[]string
	include *[]string
	fieldsSubscriptionOfferCodeOneTimeUseCodeValues *[]string
}

// the fields to include for returned resources of type subscriptionOfferCodeOneTimeUseCodes
func (r ApiSubscriptionOfferCodeOneTimeUseCodesGetInstanceRequest) FieldsSubscriptionOfferCodeOneTimeUseCodes(fieldsSubscriptionOfferCodeOneTimeUseCodes []string) ApiSubscriptionOfferCodeOneTimeUseCodesGetInstanceRequest {
	r.fieldsSubscriptionOfferCodeOneTimeUseCodes = &fieldsSubscriptionOfferCodeOneTimeUseCodes
	return r
}

// comma-separated list of relationships to include
func (r ApiSubscriptionOfferCodeOneTimeUseCodesGetInstanceRequest) Include(include []string) ApiSubscriptionOfferCodeOneTimeUseCodesGetInstanceRequest {
	r.include = &include
	return r
}

// the fields to include for returned resources of type subscriptionOfferCodeOneTimeUseCodeValues
func (r ApiSubscriptionOfferCodeOneTimeUseCodesGetInstanceRequest) FieldsSubscriptionOfferCodeOneTimeUseCodeValues(fieldsSubscriptionOfferCodeOneTimeUseCodeValues []string) ApiSubscriptionOfferCodeOneTimeUseCodesGetInstanceRequest {
	r.fieldsSubscriptionOfferCodeOneTimeUseCodeValues = &fieldsSubscriptionOfferCodeOneTimeUseCodeValues
	return r
}

func (r ApiSubscriptionOfferCodeOneTimeUseCodesGetInstanceRequest) Execute() (*SubscriptionOfferCodeOneTimeUseCodeResponse, *http.Response, error) {
	return r.ApiService.SubscriptionOfferCodeOneTimeUseCodesGetInstanceExecute(r)
}

/*
SubscriptionOfferCodeOneTimeUseCodesGetInstance Method for SubscriptionOfferCodeOneTimeUseCodesGetInstance

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id the id of the requested resource
 @return ApiSubscriptionOfferCodeOneTimeUseCodesGetInstanceRequest
*/
func (a *SubscriptionOfferCodeOneTimeUseCodesApiService) SubscriptionOfferCodeOneTimeUseCodesGetInstance(ctx context.Context, id string) ApiSubscriptionOfferCodeOneTimeUseCodesGetInstanceRequest {
	return ApiSubscriptionOfferCodeOneTimeUseCodesGetInstanceRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return SubscriptionOfferCodeOneTimeUseCodeResponse
func (a *SubscriptionOfferCodeOneTimeUseCodesApiService) SubscriptionOfferCodeOneTimeUseCodesGetInstanceExecute(r ApiSubscriptionOfferCodeOneTimeUseCodesGetInstanceRequest) (*SubscriptionOfferCodeOneTimeUseCodeResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SubscriptionOfferCodeOneTimeUseCodeResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SubscriptionOfferCodeOneTimeUseCodesApiService.SubscriptionOfferCodeOneTimeUseCodesGetInstance")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/subscriptionOfferCodeOneTimeUseCodes/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.fieldsSubscriptionOfferCodeOneTimeUseCodes != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields[subscriptionOfferCodeOneTimeUseCodes]", r.fieldsSubscriptionOfferCodeOneTimeUseCodes, "csv")
	}
	if r.include != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "include", r.include, "csv")
	}
	if r.fieldsSubscriptionOfferCodeOneTimeUseCodeValues != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields[subscriptionOfferCodeOneTimeUseCodeValues]", r.fieldsSubscriptionOfferCodeOneTimeUseCodeValues, "csv")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

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
		if localVarHTTPResponse.StatusCode == 404 {
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

type ApiSubscriptionOfferCodeOneTimeUseCodesUpdateInstanceRequest struct {
	ctx context.Context
	ApiService *SubscriptionOfferCodeOneTimeUseCodesApiService
	id string
	subscriptionOfferCodeOneTimeUseCodeUpdateRequest *SubscriptionOfferCodeOneTimeUseCodeUpdateRequest
}

// SubscriptionOfferCodeOneTimeUseCode representation
func (r ApiSubscriptionOfferCodeOneTimeUseCodesUpdateInstanceRequest) SubscriptionOfferCodeOneTimeUseCodeUpdateRequest(subscriptionOfferCodeOneTimeUseCodeUpdateRequest SubscriptionOfferCodeOneTimeUseCodeUpdateRequest) ApiSubscriptionOfferCodeOneTimeUseCodesUpdateInstanceRequest {
	r.subscriptionOfferCodeOneTimeUseCodeUpdateRequest = &subscriptionOfferCodeOneTimeUseCodeUpdateRequest
	return r
}

func (r ApiSubscriptionOfferCodeOneTimeUseCodesUpdateInstanceRequest) Execute() (*SubscriptionOfferCodeOneTimeUseCodeResponse, *http.Response, error) {
	return r.ApiService.SubscriptionOfferCodeOneTimeUseCodesUpdateInstanceExecute(r)
}

/*
SubscriptionOfferCodeOneTimeUseCodesUpdateInstance Method for SubscriptionOfferCodeOneTimeUseCodesUpdateInstance

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id the id of the requested resource
 @return ApiSubscriptionOfferCodeOneTimeUseCodesUpdateInstanceRequest
*/
func (a *SubscriptionOfferCodeOneTimeUseCodesApiService) SubscriptionOfferCodeOneTimeUseCodesUpdateInstance(ctx context.Context, id string) ApiSubscriptionOfferCodeOneTimeUseCodesUpdateInstanceRequest {
	return ApiSubscriptionOfferCodeOneTimeUseCodesUpdateInstanceRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return SubscriptionOfferCodeOneTimeUseCodeResponse
func (a *SubscriptionOfferCodeOneTimeUseCodesApiService) SubscriptionOfferCodeOneTimeUseCodesUpdateInstanceExecute(r ApiSubscriptionOfferCodeOneTimeUseCodesUpdateInstanceRequest) (*SubscriptionOfferCodeOneTimeUseCodeResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SubscriptionOfferCodeOneTimeUseCodeResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SubscriptionOfferCodeOneTimeUseCodesApiService.SubscriptionOfferCodeOneTimeUseCodesUpdateInstance")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/subscriptionOfferCodeOneTimeUseCodes/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.subscriptionOfferCodeOneTimeUseCodeUpdateRequest == nil {
		return localVarReturnValue, nil, reportError("subscriptionOfferCodeOneTimeUseCodeUpdateRequest is required and must be specified")
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
	localVarPostBody = r.subscriptionOfferCodeOneTimeUseCodeUpdateRequest
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
		if localVarHTTPResponse.StatusCode == 404 {
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

type ApiSubscriptionOfferCodeOneTimeUseCodesValuesGetToOneRelatedRequest struct {
	ctx context.Context
	ApiService *SubscriptionOfferCodeOneTimeUseCodesApiService
	id string
}

func (r ApiSubscriptionOfferCodeOneTimeUseCodesValuesGetToOneRelatedRequest) Execute() (string, *http.Response, error) {
	return r.ApiService.SubscriptionOfferCodeOneTimeUseCodesValuesGetToOneRelatedExecute(r)
}

/*
SubscriptionOfferCodeOneTimeUseCodesValuesGetToOneRelated Method for SubscriptionOfferCodeOneTimeUseCodesValuesGetToOneRelated

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id the id of the requested resource
 @return ApiSubscriptionOfferCodeOneTimeUseCodesValuesGetToOneRelatedRequest
*/
func (a *SubscriptionOfferCodeOneTimeUseCodesApiService) SubscriptionOfferCodeOneTimeUseCodesValuesGetToOneRelated(ctx context.Context, id string) ApiSubscriptionOfferCodeOneTimeUseCodesValuesGetToOneRelatedRequest {
	return ApiSubscriptionOfferCodeOneTimeUseCodesValuesGetToOneRelatedRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return string
func (a *SubscriptionOfferCodeOneTimeUseCodesApiService) SubscriptionOfferCodeOneTimeUseCodesValuesGetToOneRelatedExecute(r ApiSubscriptionOfferCodeOneTimeUseCodesValuesGetToOneRelatedRequest) (string, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  string
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SubscriptionOfferCodeOneTimeUseCodesApiService.SubscriptionOfferCodeOneTimeUseCodesValuesGetToOneRelated")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/subscriptionOfferCodeOneTimeUseCodes/{id}/values"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"text/csv", "application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
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
		if localVarHTTPResponse.StatusCode == 404 {
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
