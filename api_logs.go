/*
Chaos

Central Management API - publicly exposed set of APIs for managing deployments

API version: 1.0.0
Contact: support@qernal.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_chaos_client

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)


// LogsAPIService LogsAPI service
type LogsAPIService service

type LogsAPILogsListRequest struct {
	ctx context.Context
	ApiService *LogsAPIService
	page *OrganisationsListPageParameter
	fProject *string
	fFunction *string
	fTimestamps *LogsListFTimestampsParameter
	fQuery *string
	fLogType *string
}

// Query parameters for pagination
func (r LogsAPILogsListRequest) Page(page OrganisationsListPageParameter) LogsAPILogsListRequest {
	r.page = &page
	return r
}

// Project uuid reference
func (r LogsAPILogsListRequest) FProject(fProject string) LogsAPILogsListRequest {
	r.fProject = &fProject
	return r
}

// Function uuid reference
func (r LogsAPILogsListRequest) FFunction(fFunction string) LogsAPILogsListRequest {
	r.fFunction = &fFunction
	return r
}

// Timestamp restriction for query
func (r LogsAPILogsListRequest) FTimestamps(fTimestamps LogsListFTimestampsParameter) LogsAPILogsListRequest {
	r.fTimestamps = &fTimestamps
	return r
}

// Text query string
func (r LogsAPILogsListRequest) FQuery(fQuery string) LogsAPILogsListRequest {
	r.fQuery = &fQuery
	return r
}

// Type of log
func (r LogsAPILogsListRequest) FLogType(fLogType string) LogsAPILogsListRequest {
	r.fLogType = &fLogType
	return r
}

func (r LogsAPILogsListRequest) Execute() (*ListLogResponse, *http.Response, error) {
	return r.ApiService.LogsListExecute(r)
}

/*
LogsList Get logs

Retrieve logs for a specific project or function. Use the query parameter to search logs.

> Note: Logs are always returned in a descending order based on the timestamp.
> Note: A max size of 500 logs is returned per request (when using page[size]).


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return LogsAPILogsListRequest
*/
func (a *LogsAPIService) LogsList(ctx context.Context) LogsAPILogsListRequest {
	return LogsAPILogsListRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ListLogResponse
func (a *LogsAPIService) LogsListExecute(r LogsAPILogsListRequest) (*ListLogResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ListLogResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LogsAPIService.LogsList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/logs"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page, "deepObject", "")
	}
	if r.fProject != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "f_project", r.fProject, "form", "")
	}
	if r.fFunction != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "f_function", r.fFunction, "form", "")
	}
	if r.fTimestamps != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "f_timestamps", r.fTimestamps, "deepObject", "")
	}
	if r.fQuery != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "f_query", r.fQuery, "form", "")
	}
	if r.fLogType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "f_log_type", r.fLogType, "form", "")
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
			var v BadRequestResponse
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
