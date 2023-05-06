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
	"os"
)


// FinanceReportsApiService FinanceReportsApi service
type FinanceReportsApiService service

type ApiFinanceReportsGetCollectionRequest struct {
	ctx context.Context
	ApiService *FinanceReportsApiService
	filterRegionCode *[]string
	filterReportDate *[]string
	filterReportType *[]string
	filterVendorNumber *[]string
}

// filter by attribute &#39;regionCode&#39;
func (r ApiFinanceReportsGetCollectionRequest) FilterRegionCode(filterRegionCode []string) ApiFinanceReportsGetCollectionRequest {
	r.filterRegionCode = &filterRegionCode
	return r
}

// filter by attribute &#39;reportDate&#39;
func (r ApiFinanceReportsGetCollectionRequest) FilterReportDate(filterReportDate []string) ApiFinanceReportsGetCollectionRequest {
	r.filterReportDate = &filterReportDate
	return r
}

// filter by attribute &#39;reportType&#39;
func (r ApiFinanceReportsGetCollectionRequest) FilterReportType(filterReportType []string) ApiFinanceReportsGetCollectionRequest {
	r.filterReportType = &filterReportType
	return r
}

// filter by attribute &#39;vendorNumber&#39;
func (r ApiFinanceReportsGetCollectionRequest) FilterVendorNumber(filterVendorNumber []string) ApiFinanceReportsGetCollectionRequest {
	r.filterVendorNumber = &filterVendorNumber
	return r
}

func (r ApiFinanceReportsGetCollectionRequest) Execute() (*os.File, *http.Response, error) {
	return r.ApiService.FinanceReportsGetCollectionExecute(r)
}

/*
FinanceReportsGetCollection Method for FinanceReportsGetCollection

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiFinanceReportsGetCollectionRequest
*/
func (a *FinanceReportsApiService) FinanceReportsGetCollection(ctx context.Context) ApiFinanceReportsGetCollectionRequest {
	return ApiFinanceReportsGetCollectionRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return *os.File
func (a *FinanceReportsApiService) FinanceReportsGetCollectionExecute(r ApiFinanceReportsGetCollectionRequest) (*os.File, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *os.File
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FinanceReportsApiService.FinanceReportsGetCollection")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/financeReports"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.filterRegionCode == nil {
		return localVarReturnValue, nil, reportError("filterRegionCode is required and must be specified")
	}
	if r.filterReportDate == nil {
		return localVarReturnValue, nil, reportError("filterReportDate is required and must be specified")
	}
	if r.filterReportType == nil {
		return localVarReturnValue, nil, reportError("filterReportType is required and must be specified")
	}
	if r.filterVendorNumber == nil {
		return localVarReturnValue, nil, reportError("filterVendorNumber is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "filter[regionCode]", r.filterRegionCode, "csv")
	parameterAddToHeaderOrQuery(localVarQueryParams, "filter[reportDate]", r.filterReportDate, "csv")
	parameterAddToHeaderOrQuery(localVarQueryParams, "filter[reportType]", r.filterReportType, "csv")
	parameterAddToHeaderOrQuery(localVarQueryParams, "filter[vendorNumber]", r.filterVendorNumber, "csv")
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/a-gzip", "application/json"}

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
