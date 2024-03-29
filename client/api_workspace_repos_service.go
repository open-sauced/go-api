/*
@open-sauced/api.opensauced.pizza

 ## Swagger-UI API Documentation  This REST API can be used to create, read, update or delete data from the Open Sauced community platform. The Swagger-UI provides useful information to get started and an overview of all available resources. Each API route is clickable and has their own detailed description on how to use it. The base URL for the API is [api.opensauced.pizza](https://api.opensauced.pizza).  [comment]: # (TODO: add bearer auth information)  ## Rate limiting  Every IP address is allowed to perform 5000 requests per hour. This is measured by saving the date of the initial request and counting all requests in the next hour. When an IP address goes over the limit, HTTP status code 429 is returned. The returned HTTP headers of any API request show the current rate limit status:  header | description --- | --- `X-RateLimit-Limit` | The maximum number of requests allowed per hour `X-RateLimit-Remaining` | The number of requests remaining in the current rate limit window `X-RateLimit-Reset` | The date and time at which the current rate limit window resets in [UTC epoch seconds](https://en.wikipedia.org/wiki/Unix_time)  [comment]: # (TODO: add pagination information)  ## Common response codes  Each route shows for each method which data they expect and which they will respond when the call succeeds. The table below shows most common response codes you can receive from our endpoints.  code | condition --- | --- [`200`](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/200) | The [`GET`](https://developer.mozilla.org/en-US/docs/Web/HTTP/Methods/GET) request was handled successfully. The response provides the requested data. [`201`](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/201) | The [`POST`](https://developer.mozilla.org/en-US/docs/Web/HTTP/Methods/POST) request was handled successfully. The response provides the created data. [`204`](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/204) | The [`PATCH`](https://developer.mozilla.org/en-US/docs/Web/HTTP/Methods/PATCH) or [`DELETE`](https://developer.mozilla.org/en-US/docs/Web/HTTP/Methods/DELETE) request was handled successfully. The response provides no data, generally. [`400`](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/400) | The server will not process the request due to something that is perceived to be a client error. Check the provided error for mote information. [`401`](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/401) | The request requires user authentication. Check the provided error for more information. [`403`](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/403) | The request was valid, but the server is refusing user access. Check the provided error for more information. [`404`](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/404) | The requested resource could not be found. Check the provided error for more information. [`429`](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/429) | The current API Key made too many requests in the last hour. Check [Rate limiting](#ratelimiting) for more information.  ## Additional links

API version: 2
Contact: hello@opensauced.pizza
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// WorkspaceReposServiceAPIService WorkspaceReposServiceAPI service
type WorkspaceReposServiceAPIService service

type ApiAddOneWorkspaceRepoForUserRequest struct {
	ctx        context.Context
	ApiService *WorkspaceReposServiceAPIService
	id         string
	owner      string
	repo       string
}

func (r ApiAddOneWorkspaceRepoForUserRequest) Execute() (*DbWorkspaceRepo, *http.Response, error) {
	return r.ApiService.AddOneWorkspaceRepoForUserExecute(r)
}

/*
AddOneWorkspaceRepoForUser Adds workspace repos for the authenticated user

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id
	@param owner
	@param repo
	@return ApiAddOneWorkspaceRepoForUserRequest
*/
func (a *WorkspaceReposServiceAPIService) AddOneWorkspaceRepoForUser(ctx context.Context, id string, owner string, repo string) ApiAddOneWorkspaceRepoForUserRequest {
	return ApiAddOneWorkspaceRepoForUserRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
		owner:      owner,
		repo:       repo,
	}
}

// Execute executes the request
//
//	@return DbWorkspaceRepo
func (a *WorkspaceReposServiceAPIService) AddOneWorkspaceRepoForUserExecute(r ApiAddOneWorkspaceRepoForUserRequest) (*DbWorkspaceRepo, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *DbWorkspaceRepo
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WorkspaceReposServiceAPIService.AddOneWorkspaceRepoForUser")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/workspaces/{id}/repos/{owner}/{repo}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"owner"+"}", url.PathEscape(parameterValueToString(r.owner, "owner")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"repo"+"}", url.PathEscape(parameterValueToString(r.repo, "repo")), -1)

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

type ApiAddWorkspaceReposForUserRequest struct {
	ctx                     context.Context
	ApiService              *WorkspaceReposServiceAPIService
	id                      string
	updateWorkspaceReposDto *UpdateWorkspaceReposDto
}

func (r ApiAddWorkspaceReposForUserRequest) UpdateWorkspaceReposDto(updateWorkspaceReposDto UpdateWorkspaceReposDto) ApiAddWorkspaceReposForUserRequest {
	r.updateWorkspaceReposDto = &updateWorkspaceReposDto
	return r
}

func (r ApiAddWorkspaceReposForUserRequest) Execute() (*DbWorkspaceRepo, *http.Response, error) {
	return r.ApiService.AddWorkspaceReposForUserExecute(r)
}

/*
AddWorkspaceReposForUser Adds workspace repos for the authenticated user

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id
	@return ApiAddWorkspaceReposForUserRequest
*/
func (a *WorkspaceReposServiceAPIService) AddWorkspaceReposForUser(ctx context.Context, id string) ApiAddWorkspaceReposForUserRequest {
	return ApiAddWorkspaceReposForUserRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
//
//	@return DbWorkspaceRepo
func (a *WorkspaceReposServiceAPIService) AddWorkspaceReposForUserExecute(r ApiAddWorkspaceReposForUserRequest) (*DbWorkspaceRepo, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *DbWorkspaceRepo
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WorkspaceReposServiceAPIService.AddWorkspaceReposForUser")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/workspaces/{id}/repos"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.updateWorkspaceReposDto == nil {
		return localVarReturnValue, nil, reportError("updateWorkspaceReposDto is required and must be specified")
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
	localVarPostBody = r.updateWorkspaceReposDto
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

type ApiDeleteOneWorkspaceRepoForUserRequest struct {
	ctx        context.Context
	ApiService *WorkspaceReposServiceAPIService
	id         string
	owner      string
	repo       string
}

func (r ApiDeleteOneWorkspaceRepoForUserRequest) Execute() (*DbWorkspace, *http.Response, error) {
	return r.ApiService.DeleteOneWorkspaceRepoForUserExecute(r)
}

/*
DeleteOneWorkspaceRepoForUser Delete a workspace repos for the authenticated user

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id
	@param owner
	@param repo
	@return ApiDeleteOneWorkspaceRepoForUserRequest
*/
func (a *WorkspaceReposServiceAPIService) DeleteOneWorkspaceRepoForUser(ctx context.Context, id string, owner string, repo string) ApiDeleteOneWorkspaceRepoForUserRequest {
	return ApiDeleteOneWorkspaceRepoForUserRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
		owner:      owner,
		repo:       repo,
	}
}

// Execute executes the request
//
//	@return DbWorkspace
func (a *WorkspaceReposServiceAPIService) DeleteOneWorkspaceRepoForUserExecute(r ApiDeleteOneWorkspaceRepoForUserRequest) (*DbWorkspace, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodDelete
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *DbWorkspace
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WorkspaceReposServiceAPIService.DeleteOneWorkspaceRepoForUser")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/workspaces/{id}/repos/{owner}/{repo}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"owner"+"}", url.PathEscape(parameterValueToString(r.owner, "owner")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"repo"+"}", url.PathEscape(parameterValueToString(r.repo, "repo")), -1)

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

type ApiDeleteWorkspaceReposForUserRequest struct {
	ctx                     context.Context
	ApiService              *WorkspaceReposServiceAPIService
	id                      string
	deleteWorkspaceReposDto *DeleteWorkspaceReposDto
}

func (r ApiDeleteWorkspaceReposForUserRequest) DeleteWorkspaceReposDto(deleteWorkspaceReposDto DeleteWorkspaceReposDto) ApiDeleteWorkspaceReposForUserRequest {
	r.deleteWorkspaceReposDto = &deleteWorkspaceReposDto
	return r
}

func (r ApiDeleteWorkspaceReposForUserRequest) Execute() (*DbWorkspace, *http.Response, error) {
	return r.ApiService.DeleteWorkspaceReposForUserExecute(r)
}

/*
DeleteWorkspaceReposForUser Deletes workspace repos for the authenticated user

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id
	@return ApiDeleteWorkspaceReposForUserRequest
*/
func (a *WorkspaceReposServiceAPIService) DeleteWorkspaceReposForUser(ctx context.Context, id string) ApiDeleteWorkspaceReposForUserRequest {
	return ApiDeleteWorkspaceReposForUserRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
//
//	@return DbWorkspace
func (a *WorkspaceReposServiceAPIService) DeleteWorkspaceReposForUserExecute(r ApiDeleteWorkspaceReposForUserRequest) (*DbWorkspace, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodDelete
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *DbWorkspace
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WorkspaceReposServiceAPIService.DeleteWorkspaceReposForUser")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/workspaces/{id}/repos"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.deleteWorkspaceReposDto == nil {
		return localVarReturnValue, nil, reportError("deleteWorkspaceReposDto is required and must be specified")
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
	localVarPostBody = r.deleteWorkspaceReposDto
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

type ApiGetWorkspaceReposForUserRequest struct {
	ctx               context.Context
	ApiService        *WorkspaceReposServiceAPIService
	id                string
	page              *int32
	limit             *int32
	orderDirection    *OrderDirectionEnum
	range_            *int32
	prevDaysStartDate *int32
}

func (r ApiGetWorkspaceReposForUserRequest) Page(page int32) ApiGetWorkspaceReposForUserRequest {
	r.page = &page
	return r
}

func (r ApiGetWorkspaceReposForUserRequest) Limit(limit int32) ApiGetWorkspaceReposForUserRequest {
	r.limit = &limit
	return r
}

func (r ApiGetWorkspaceReposForUserRequest) OrderDirection(orderDirection OrderDirectionEnum) ApiGetWorkspaceReposForUserRequest {
	r.orderDirection = &orderDirection
	return r
}

// Range in days
func (r ApiGetWorkspaceReposForUserRequest) Range_(range_ int32) ApiGetWorkspaceReposForUserRequest {
	r.range_ = &range_
	return r
}

// Number of days in the past to start range block
func (r ApiGetWorkspaceReposForUserRequest) PrevDaysStartDate(prevDaysStartDate int32) ApiGetWorkspaceReposForUserRequest {
	r.prevDaysStartDate = &prevDaysStartDate
	return r
}

func (r ApiGetWorkspaceReposForUserRequest) Execute() (*DbWorkspaceRepo, *http.Response, error) {
	return r.ApiService.GetWorkspaceReposForUserExecute(r)
}

/*
GetWorkspaceReposForUser Gets workspace repos for the authenticated user

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id
	@return ApiGetWorkspaceReposForUserRequest
*/
func (a *WorkspaceReposServiceAPIService) GetWorkspaceReposForUser(ctx context.Context, id string) ApiGetWorkspaceReposForUserRequest {
	return ApiGetWorkspaceReposForUserRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
//
//	@return DbWorkspaceRepo
func (a *WorkspaceReposServiceAPIService) GetWorkspaceReposForUserExecute(r ApiGetWorkspaceReposForUserRequest) (*DbWorkspaceRepo, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *DbWorkspaceRepo
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WorkspaceReposServiceAPIService.GetWorkspaceReposForUser")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/workspaces/{id}/repos"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page, "")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
	}
	if r.orderDirection != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "orderDirection", r.orderDirection, "")
	}
	if r.range_ != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "range", r.range_, "")
	}
	if r.prevDaysStartDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "prev_days_start_date", r.prevDaysStartDate, "")
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
