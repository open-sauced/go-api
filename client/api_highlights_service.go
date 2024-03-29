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

// HighlightsServiceAPIService HighlightsServiceAPI service
type HighlightsServiceAPIService service

type ApiAddAFeaturedHighlightRequest struct {
	ctx        context.Context
	ApiService *HighlightsServiceAPIService
	id         int32
}

func (r ApiAddAFeaturedHighlightRequest) Execute() (*DbUserHighlight, *http.Response, error) {
	return r.ApiService.AddAFeaturedHighlightExecute(r)
}

/*
AddAFeaturedHighlight Add a highlight to the featured list

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id
	@return ApiAddAFeaturedHighlightRequest
*/
func (a *HighlightsServiceAPIService) AddAFeaturedHighlight(ctx context.Context, id int32) ApiAddAFeaturedHighlightRequest {
	return ApiAddAFeaturedHighlightRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
//
//	@return DbUserHighlight
func (a *HighlightsServiceAPIService) AddAFeaturedHighlightExecute(r ApiAddAFeaturedHighlightRequest) (*DbUserHighlight, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *DbUserHighlight
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "HighlightsServiceAPIService.AddAFeaturedHighlight")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/highlights/{id}/featured"
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

type ApiFindAllFeaturedHighlightsRequest struct {
	ctx               context.Context
	ApiService        *HighlightsServiceAPIService
	page              *int32
	limit             *int32
	orderDirection    *OrderDirectionEnum
	range_            *int32
	prevDaysStartDate *int32
	repo              *string
	contributor       *string
}

func (r ApiFindAllFeaturedHighlightsRequest) Page(page int32) ApiFindAllFeaturedHighlightsRequest {
	r.page = &page
	return r
}

func (r ApiFindAllFeaturedHighlightsRequest) Limit(limit int32) ApiFindAllFeaturedHighlightsRequest {
	r.limit = &limit
	return r
}

func (r ApiFindAllFeaturedHighlightsRequest) OrderDirection(orderDirection OrderDirectionEnum) ApiFindAllFeaturedHighlightsRequest {
	r.orderDirection = &orderDirection
	return r
}

// Range in days
func (r ApiFindAllFeaturedHighlightsRequest) Range_(range_ int32) ApiFindAllFeaturedHighlightsRequest {
	r.range_ = &range_
	return r
}

// Number of days in the past to start range block
func (r ApiFindAllFeaturedHighlightsRequest) PrevDaysStartDate(prevDaysStartDate int32) ApiFindAllFeaturedHighlightsRequest {
	r.prevDaysStartDate = &prevDaysStartDate
	return r
}

// Highlight Repo Filter
func (r ApiFindAllFeaturedHighlightsRequest) Repo(repo string) ApiFindAllFeaturedHighlightsRequest {
	r.repo = &repo
	return r
}

// Highlight User Filter
func (r ApiFindAllFeaturedHighlightsRequest) Contributor(contributor string) ApiFindAllFeaturedHighlightsRequest {
	r.contributor = &contributor
	return r
}

func (r ApiFindAllFeaturedHighlightsRequest) Execute() (*FindAllHighlightsByUsername200Response, *http.Response, error) {
	return r.ApiService.FindAllFeaturedHighlightsExecute(r)
}

/*
FindAllFeaturedHighlights Finds all featured highlights and paginates them

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiFindAllFeaturedHighlightsRequest
*/
func (a *HighlightsServiceAPIService) FindAllFeaturedHighlights(ctx context.Context) ApiFindAllFeaturedHighlightsRequest {
	return ApiFindAllFeaturedHighlightsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return FindAllHighlightsByUsername200Response
func (a *HighlightsServiceAPIService) FindAllFeaturedHighlightsExecute(r ApiFindAllFeaturedHighlightsRequest) (*FindAllHighlightsByUsername200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *FindAllHighlightsByUsername200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "HighlightsServiceAPIService.FindAllFeaturedHighlights")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/highlights/featured"

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
	if r.repo != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "repo", r.repo, "")
	}
	if r.contributor != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "contributor", r.contributor, "")
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

type ApiFindAllHighlightReposRequest struct {
	ctx               context.Context
	ApiService        *HighlightsServiceAPIService
	page              *int32
	limit             *int32
	orderDirection    *OrderDirectionEnum
	range_            *int32
	prevDaysStartDate *int32
}

func (r ApiFindAllHighlightReposRequest) Page(page int32) ApiFindAllHighlightReposRequest {
	r.page = &page
	return r
}

func (r ApiFindAllHighlightReposRequest) Limit(limit int32) ApiFindAllHighlightReposRequest {
	r.limit = &limit
	return r
}

func (r ApiFindAllHighlightReposRequest) OrderDirection(orderDirection OrderDirectionEnum) ApiFindAllHighlightReposRequest {
	r.orderDirection = &orderDirection
	return r
}

// Range in days
func (r ApiFindAllHighlightReposRequest) Range_(range_ int32) ApiFindAllHighlightReposRequest {
	r.range_ = &range_
	return r
}

// Number of days in the past to start range block
func (r ApiFindAllHighlightReposRequest) PrevDaysStartDate(prevDaysStartDate int32) ApiFindAllHighlightReposRequest {
	r.prevDaysStartDate = &prevDaysStartDate
	return r
}

func (r ApiFindAllHighlightReposRequest) Execute() (*GetUserListContributorHighlightedRepos200Response, *http.Response, error) {
	return r.ApiService.FindAllHighlightReposExecute(r)
}

/*
FindAllHighlightRepos Finds all highlight repos and paginates them

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiFindAllHighlightReposRequest
*/
func (a *HighlightsServiceAPIService) FindAllHighlightRepos(ctx context.Context) ApiFindAllHighlightReposRequest {
	return ApiFindAllHighlightReposRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetUserListContributorHighlightedRepos200Response
func (a *HighlightsServiceAPIService) FindAllHighlightReposExecute(r ApiFindAllHighlightReposRequest) (*GetUserListContributorHighlightedRepos200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetUserListContributorHighlightedRepos200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "HighlightsServiceAPIService.FindAllHighlightRepos")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/highlights/repos/list"

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

type ApiFindAllHighlightsRequest struct {
	ctx               context.Context
	ApiService        *HighlightsServiceAPIService
	page              *int32
	limit             *int32
	orderDirection    *OrderDirectionEnum
	range_            *int32
	prevDaysStartDate *int32
	repo              *string
	contributor       *string
}

func (r ApiFindAllHighlightsRequest) Page(page int32) ApiFindAllHighlightsRequest {
	r.page = &page
	return r
}

func (r ApiFindAllHighlightsRequest) Limit(limit int32) ApiFindAllHighlightsRequest {
	r.limit = &limit
	return r
}

func (r ApiFindAllHighlightsRequest) OrderDirection(orderDirection OrderDirectionEnum) ApiFindAllHighlightsRequest {
	r.orderDirection = &orderDirection
	return r
}

// Range in days
func (r ApiFindAllHighlightsRequest) Range_(range_ int32) ApiFindAllHighlightsRequest {
	r.range_ = &range_
	return r
}

// Number of days in the past to start range block
func (r ApiFindAllHighlightsRequest) PrevDaysStartDate(prevDaysStartDate int32) ApiFindAllHighlightsRequest {
	r.prevDaysStartDate = &prevDaysStartDate
	return r
}

// Highlight Repo Filter
func (r ApiFindAllHighlightsRequest) Repo(repo string) ApiFindAllHighlightsRequest {
	r.repo = &repo
	return r
}

// Highlight User Filter
func (r ApiFindAllHighlightsRequest) Contributor(contributor string) ApiFindAllHighlightsRequest {
	r.contributor = &contributor
	return r
}

func (r ApiFindAllHighlightsRequest) Execute() (*FindAllHighlightsByUsername200Response, *http.Response, error) {
	return r.ApiService.FindAllHighlightsExecute(r)
}

/*
FindAllHighlights Finds all highlights and paginates them

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiFindAllHighlightsRequest
*/
func (a *HighlightsServiceAPIService) FindAllHighlights(ctx context.Context) ApiFindAllHighlightsRequest {
	return ApiFindAllHighlightsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return FindAllHighlightsByUsername200Response
func (a *HighlightsServiceAPIService) FindAllHighlightsExecute(r ApiFindAllHighlightsRequest) (*FindAllHighlightsByUsername200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *FindAllHighlightsByUsername200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "HighlightsServiceAPIService.FindAllHighlights")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/highlights/list"

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
	if r.repo != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "repo", r.repo, "")
	}
	if r.contributor != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "contributor", r.contributor, "")
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

type ApiFindTopHighlightsRequest struct {
	ctx               context.Context
	ApiService        *HighlightsServiceAPIService
	page              *int32
	limit             *int32
	orderDirection    *OrderDirectionEnum
	range_            *int32
	prevDaysStartDate *int32
}

func (r ApiFindTopHighlightsRequest) Page(page int32) ApiFindTopHighlightsRequest {
	r.page = &page
	return r
}

func (r ApiFindTopHighlightsRequest) Limit(limit int32) ApiFindTopHighlightsRequest {
	r.limit = &limit
	return r
}

func (r ApiFindTopHighlightsRequest) OrderDirection(orderDirection OrderDirectionEnum) ApiFindTopHighlightsRequest {
	r.orderDirection = &orderDirection
	return r
}

// Range in days
func (r ApiFindTopHighlightsRequest) Range_(range_ int32) ApiFindTopHighlightsRequest {
	r.range_ = &range_
	return r
}

// Number of days in the past to start range block
func (r ApiFindTopHighlightsRequest) PrevDaysStartDate(prevDaysStartDate int32) ApiFindTopHighlightsRequest {
	r.prevDaysStartDate = &prevDaysStartDate
	return r
}

func (r ApiFindTopHighlightsRequest) Execute() (*FindAllHighlightsByUsername200Response, *http.Response, error) {
	return r.ApiService.FindTopHighlightsExecute(r)
}

/*
FindTopHighlights Finds top highlights for the day range and paginates them

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiFindTopHighlightsRequest
*/
func (a *HighlightsServiceAPIService) FindTopHighlights(ctx context.Context) ApiFindTopHighlightsRequest {
	return ApiFindTopHighlightsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return FindAllHighlightsByUsername200Response
func (a *HighlightsServiceAPIService) FindTopHighlightsExecute(r ApiFindTopHighlightsRequest) (*FindAllHighlightsByUsername200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *FindAllHighlightsByUsername200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "HighlightsServiceAPIService.FindTopHighlights")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/highlights/top"

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

type ApiGetAllHighlightReactionsRequest struct {
	ctx        context.Context
	ApiService *HighlightsServiceAPIService
	id         int32
}

func (r ApiGetAllHighlightReactionsRequest) Execute() (*DbUserHighlightReactionResponse, *http.Response, error) {
	return r.ApiService.GetAllHighlightReactionsExecute(r)
}

/*
GetAllHighlightReactions Retrieves total reactions for a highlight

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id
	@return ApiGetAllHighlightReactionsRequest
*/
func (a *HighlightsServiceAPIService) GetAllHighlightReactions(ctx context.Context, id int32) ApiGetAllHighlightReactionsRequest {
	return ApiGetAllHighlightReactionsRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
//
//	@return DbUserHighlightReactionResponse
func (a *HighlightsServiceAPIService) GetAllHighlightReactionsExecute(r ApiGetAllHighlightReactionsRequest) (*DbUserHighlightReactionResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *DbUserHighlightReactionResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "HighlightsServiceAPIService.GetAllHighlightReactions")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/highlights/{id}/reactions"
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

type ApiRemoveAFeaturedHighlightRequest struct {
	ctx        context.Context
	ApiService *HighlightsServiceAPIService
	id         int32
}

func (r ApiRemoveAFeaturedHighlightRequest) Execute() (*DbUserHighlight, *http.Response, error) {
	return r.ApiService.RemoveAFeaturedHighlightExecute(r)
}

/*
RemoveAFeaturedHighlight Remove a highlight from the featured list

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id
	@return ApiRemoveAFeaturedHighlightRequest
*/
func (a *HighlightsServiceAPIService) RemoveAFeaturedHighlight(ctx context.Context, id int32) ApiRemoveAFeaturedHighlightRequest {
	return ApiRemoveAFeaturedHighlightRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
//
//	@return DbUserHighlight
func (a *HighlightsServiceAPIService) RemoveAFeaturedHighlightExecute(r ApiRemoveAFeaturedHighlightRequest) (*DbUserHighlight, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodDelete
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *DbUserHighlight
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "HighlightsServiceAPIService.RemoveAFeaturedHighlight")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/highlights/{id}/featured"
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
