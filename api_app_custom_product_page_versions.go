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


// AppCustomProductPageVersionsApiService AppCustomProductPageVersionsApi service
type AppCustomProductPageVersionsApiService service

type ApiAppCustomProductPageVersionsAppCustomProductPageLocalizationsGetToManyRelatedRequest struct {
	ctx context.Context
	ApiService *AppCustomProductPageVersionsApiService
	id string
	filterLocale *[]string
	fieldsAppScreenshotSets *[]string
	fieldsAppCustomProductPageLocalizations *[]string
	fieldsAppCustomProductPageVersions *[]string
	fieldsAppPreviewSets *[]string
	limit *int32
	limitAppScreenshotSets *int32
	limitAppPreviewSets *int32
	include *[]string
}

// filter by attribute &#39;locale&#39;
func (r ApiAppCustomProductPageVersionsAppCustomProductPageLocalizationsGetToManyRelatedRequest) FilterLocale(filterLocale []string) ApiAppCustomProductPageVersionsAppCustomProductPageLocalizationsGetToManyRelatedRequest {
	r.filterLocale = &filterLocale
	return r
}

// the fields to include for returned resources of type appScreenshotSets
func (r ApiAppCustomProductPageVersionsAppCustomProductPageLocalizationsGetToManyRelatedRequest) FieldsAppScreenshotSets(fieldsAppScreenshotSets []string) ApiAppCustomProductPageVersionsAppCustomProductPageLocalizationsGetToManyRelatedRequest {
	r.fieldsAppScreenshotSets = &fieldsAppScreenshotSets
	return r
}

// the fields to include for returned resources of type appCustomProductPageLocalizations
func (r ApiAppCustomProductPageVersionsAppCustomProductPageLocalizationsGetToManyRelatedRequest) FieldsAppCustomProductPageLocalizations(fieldsAppCustomProductPageLocalizations []string) ApiAppCustomProductPageVersionsAppCustomProductPageLocalizationsGetToManyRelatedRequest {
	r.fieldsAppCustomProductPageLocalizations = &fieldsAppCustomProductPageLocalizations
	return r
}

// the fields to include for returned resources of type appCustomProductPageVersions
func (r ApiAppCustomProductPageVersionsAppCustomProductPageLocalizationsGetToManyRelatedRequest) FieldsAppCustomProductPageVersions(fieldsAppCustomProductPageVersions []string) ApiAppCustomProductPageVersionsAppCustomProductPageLocalizationsGetToManyRelatedRequest {
	r.fieldsAppCustomProductPageVersions = &fieldsAppCustomProductPageVersions
	return r
}

// the fields to include for returned resources of type appPreviewSets
func (r ApiAppCustomProductPageVersionsAppCustomProductPageLocalizationsGetToManyRelatedRequest) FieldsAppPreviewSets(fieldsAppPreviewSets []string) ApiAppCustomProductPageVersionsAppCustomProductPageLocalizationsGetToManyRelatedRequest {
	r.fieldsAppPreviewSets = &fieldsAppPreviewSets
	return r
}

// maximum resources per page
func (r ApiAppCustomProductPageVersionsAppCustomProductPageLocalizationsGetToManyRelatedRequest) Limit(limit int32) ApiAppCustomProductPageVersionsAppCustomProductPageLocalizationsGetToManyRelatedRequest {
	r.limit = &limit
	return r
}

// maximum number of related appScreenshotSets returned (when they are included)
func (r ApiAppCustomProductPageVersionsAppCustomProductPageLocalizationsGetToManyRelatedRequest) LimitAppScreenshotSets(limitAppScreenshotSets int32) ApiAppCustomProductPageVersionsAppCustomProductPageLocalizationsGetToManyRelatedRequest {
	r.limitAppScreenshotSets = &limitAppScreenshotSets
	return r
}

// maximum number of related appPreviewSets returned (when they are included)
func (r ApiAppCustomProductPageVersionsAppCustomProductPageLocalizationsGetToManyRelatedRequest) LimitAppPreviewSets(limitAppPreviewSets int32) ApiAppCustomProductPageVersionsAppCustomProductPageLocalizationsGetToManyRelatedRequest {
	r.limitAppPreviewSets = &limitAppPreviewSets
	return r
}

// comma-separated list of relationships to include
func (r ApiAppCustomProductPageVersionsAppCustomProductPageLocalizationsGetToManyRelatedRequest) Include(include []string) ApiAppCustomProductPageVersionsAppCustomProductPageLocalizationsGetToManyRelatedRequest {
	r.include = &include
	return r
}

func (r ApiAppCustomProductPageVersionsAppCustomProductPageLocalizationsGetToManyRelatedRequest) Execute() (*AppCustomProductPageLocalizationsResponse, *http.Response, error) {
	return r.ApiService.AppCustomProductPageVersionsAppCustomProductPageLocalizationsGetToManyRelatedExecute(r)
}

/*
AppCustomProductPageVersionsAppCustomProductPageLocalizationsGetToManyRelated Method for AppCustomProductPageVersionsAppCustomProductPageLocalizationsGetToManyRelated

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id the id of the requested resource
 @return ApiAppCustomProductPageVersionsAppCustomProductPageLocalizationsGetToManyRelatedRequest
*/
func (a *AppCustomProductPageVersionsApiService) AppCustomProductPageVersionsAppCustomProductPageLocalizationsGetToManyRelated(ctx context.Context, id string) ApiAppCustomProductPageVersionsAppCustomProductPageLocalizationsGetToManyRelatedRequest {
	return ApiAppCustomProductPageVersionsAppCustomProductPageLocalizationsGetToManyRelatedRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return AppCustomProductPageLocalizationsResponse
func (a *AppCustomProductPageVersionsApiService) AppCustomProductPageVersionsAppCustomProductPageLocalizationsGetToManyRelatedExecute(r ApiAppCustomProductPageVersionsAppCustomProductPageLocalizationsGetToManyRelatedRequest) (*AppCustomProductPageLocalizationsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AppCustomProductPageLocalizationsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AppCustomProductPageVersionsApiService.AppCustomProductPageVersionsAppCustomProductPageLocalizationsGetToManyRelated")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/appCustomProductPageVersions/{id}/appCustomProductPageLocalizations"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.filterLocale != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter[locale]", r.filterLocale, "csv")
	}
	if r.fieldsAppScreenshotSets != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields[appScreenshotSets]", r.fieldsAppScreenshotSets, "csv")
	}
	if r.fieldsAppCustomProductPageLocalizations != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields[appCustomProductPageLocalizations]", r.fieldsAppCustomProductPageLocalizations, "csv")
	}
	if r.fieldsAppCustomProductPageVersions != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields[appCustomProductPageVersions]", r.fieldsAppCustomProductPageVersions, "csv")
	}
	if r.fieldsAppPreviewSets != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields[appPreviewSets]", r.fieldsAppPreviewSets, "csv")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
	}
	if r.limitAppScreenshotSets != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit[appScreenshotSets]", r.limitAppScreenshotSets, "")
	}
	if r.limitAppPreviewSets != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit[appPreviewSets]", r.limitAppPreviewSets, "")
	}
	if r.include != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "include", r.include, "csv")
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

type ApiAppCustomProductPageVersionsCreateInstanceRequest struct {
	ctx context.Context
	ApiService *AppCustomProductPageVersionsApiService
	appCustomProductPageVersionCreateRequest *AppCustomProductPageVersionCreateRequest
}

// AppCustomProductPageVersion representation
func (r ApiAppCustomProductPageVersionsCreateInstanceRequest) AppCustomProductPageVersionCreateRequest(appCustomProductPageVersionCreateRequest AppCustomProductPageVersionCreateRequest) ApiAppCustomProductPageVersionsCreateInstanceRequest {
	r.appCustomProductPageVersionCreateRequest = &appCustomProductPageVersionCreateRequest
	return r
}

func (r ApiAppCustomProductPageVersionsCreateInstanceRequest) Execute() (*AppCustomProductPageVersionResponse, *http.Response, error) {
	return r.ApiService.AppCustomProductPageVersionsCreateInstanceExecute(r)
}

/*
AppCustomProductPageVersionsCreateInstance Method for AppCustomProductPageVersionsCreateInstance

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiAppCustomProductPageVersionsCreateInstanceRequest
*/
func (a *AppCustomProductPageVersionsApiService) AppCustomProductPageVersionsCreateInstance(ctx context.Context) ApiAppCustomProductPageVersionsCreateInstanceRequest {
	return ApiAppCustomProductPageVersionsCreateInstanceRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return AppCustomProductPageVersionResponse
func (a *AppCustomProductPageVersionsApiService) AppCustomProductPageVersionsCreateInstanceExecute(r ApiAppCustomProductPageVersionsCreateInstanceRequest) (*AppCustomProductPageVersionResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AppCustomProductPageVersionResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AppCustomProductPageVersionsApiService.AppCustomProductPageVersionsCreateInstance")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/appCustomProductPageVersions"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.appCustomProductPageVersionCreateRequest == nil {
		return localVarReturnValue, nil, reportError("appCustomProductPageVersionCreateRequest is required and must be specified")
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
	localVarPostBody = r.appCustomProductPageVersionCreateRequest
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

type ApiAppCustomProductPageVersionsGetInstanceRequest struct {
	ctx context.Context
	ApiService *AppCustomProductPageVersionsApiService
	id string
	fieldsAppCustomProductPageVersions *[]string
	include *[]string
	fieldsAppCustomProductPageLocalizations *[]string
	limitAppCustomProductPageLocalizations *int32
}

// the fields to include for returned resources of type appCustomProductPageVersions
func (r ApiAppCustomProductPageVersionsGetInstanceRequest) FieldsAppCustomProductPageVersions(fieldsAppCustomProductPageVersions []string) ApiAppCustomProductPageVersionsGetInstanceRequest {
	r.fieldsAppCustomProductPageVersions = &fieldsAppCustomProductPageVersions
	return r
}

// comma-separated list of relationships to include
func (r ApiAppCustomProductPageVersionsGetInstanceRequest) Include(include []string) ApiAppCustomProductPageVersionsGetInstanceRequest {
	r.include = &include
	return r
}

// the fields to include for returned resources of type appCustomProductPageLocalizations
func (r ApiAppCustomProductPageVersionsGetInstanceRequest) FieldsAppCustomProductPageLocalizations(fieldsAppCustomProductPageLocalizations []string) ApiAppCustomProductPageVersionsGetInstanceRequest {
	r.fieldsAppCustomProductPageLocalizations = &fieldsAppCustomProductPageLocalizations
	return r
}

// maximum number of related appCustomProductPageLocalizations returned (when they are included)
func (r ApiAppCustomProductPageVersionsGetInstanceRequest) LimitAppCustomProductPageLocalizations(limitAppCustomProductPageLocalizations int32) ApiAppCustomProductPageVersionsGetInstanceRequest {
	r.limitAppCustomProductPageLocalizations = &limitAppCustomProductPageLocalizations
	return r
}

func (r ApiAppCustomProductPageVersionsGetInstanceRequest) Execute() (*AppCustomProductPageVersionResponse, *http.Response, error) {
	return r.ApiService.AppCustomProductPageVersionsGetInstanceExecute(r)
}

/*
AppCustomProductPageVersionsGetInstance Method for AppCustomProductPageVersionsGetInstance

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id the id of the requested resource
 @return ApiAppCustomProductPageVersionsGetInstanceRequest
*/
func (a *AppCustomProductPageVersionsApiService) AppCustomProductPageVersionsGetInstance(ctx context.Context, id string) ApiAppCustomProductPageVersionsGetInstanceRequest {
	return ApiAppCustomProductPageVersionsGetInstanceRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return AppCustomProductPageVersionResponse
func (a *AppCustomProductPageVersionsApiService) AppCustomProductPageVersionsGetInstanceExecute(r ApiAppCustomProductPageVersionsGetInstanceRequest) (*AppCustomProductPageVersionResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AppCustomProductPageVersionResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AppCustomProductPageVersionsApiService.AppCustomProductPageVersionsGetInstance")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/appCustomProductPageVersions/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.fieldsAppCustomProductPageVersions != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields[appCustomProductPageVersions]", r.fieldsAppCustomProductPageVersions, "csv")
	}
	if r.include != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "include", r.include, "csv")
	}
	if r.fieldsAppCustomProductPageLocalizations != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields[appCustomProductPageLocalizations]", r.fieldsAppCustomProductPageLocalizations, "csv")
	}
	if r.limitAppCustomProductPageLocalizations != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit[appCustomProductPageLocalizations]", r.limitAppCustomProductPageLocalizations, "")
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
