/*
@open-sauced/api.opensauced.pizza

 ## Swagger-UI API Documentation  This REST API can be used to create, read, update or delete data from the Open Sauced community platform. The Swagger-UI provides useful information to get started and an overview of all available resources. Each API route is clickable and has their own detailed description on how to use it. The base URL for the API is [api.opensauced.pizza](https://api.opensauced.pizza).  [comment]: # (TODO: add bearer auth information)  ## Rate limiting  Every IP address is allowed to perform 5000 requests per hour. This is measured by saving the date of the initial request and counting all requests in the next hour. When an IP address goes over the limit, HTTP status code 429 is returned. The returned HTTP headers of any API request show the current rate limit status:  header | description --- | --- `X-RateLimit-Limit` | The maximum number of requests allowed per hour `X-RateLimit-Remaining` | The number of requests remaining in the current rate limit window `X-RateLimit-Reset` | The date and time at which the current rate limit window resets in [UTC epoch seconds](https://en.wikipedia.org/wiki/Unix_time)  [comment]: # (TODO: add pagination information)  ## Common response codes  Each route shows for each method which data they expect and which they will respond when the call succeeds. The table below shows most common response codes you can receive from our endpoints.  code | condition --- | --- [`200`](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/200) | The [`GET`](https://developer.mozilla.org/en-US/docs/Web/HTTP/Methods/GET) request was handled successfully. The response provides the requested data. [`201`](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/201) | The [`POST`](https://developer.mozilla.org/en-US/docs/Web/HTTP/Methods/POST) request was handled successfully. The response provides the created data. [`204`](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/204) | The [`PATCH`](https://developer.mozilla.org/en-US/docs/Web/HTTP/Methods/PATCH) or [`DELETE`](https://developer.mozilla.org/en-US/docs/Web/HTTP/Methods/DELETE) request was handled successfully. The response provides no data, generally. [`400`](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/400) | The server will not process the request due to something that is perceived to be a client error. Check the provided error for mote information. [`401`](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/401) | The request requires user authentication. Check the provided error for more information. [`403`](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/403) | The request was valid, but the server is refusing user access. Check the provided error for more information. [`404`](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/404) | The requested resource could not be found. Check the provided error for more information. [`429`](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/429) | The current API Key made too many requests in the last hour. Check [Rate limiting](#ratelimiting) for more information.  ## Additional links

API version: 1
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

// ContributionServiceAPIService ContributionServiceAPI service
type ContributionServiceAPIService service

type ApiFindAllByOwnerAndRepoRequest struct {
	ctx            context.Context
	ApiService     *ContributionServiceAPIService
	owner          string
	repo           string
	page           *int32
	limit          *int32
	orderDirection *OrderDirectionEnum
	range_         *int32
	orderBy        *RepoOrderFieldsEnum
}

func (r ApiFindAllByOwnerAndRepoRequest) Page(page int32) ApiFindAllByOwnerAndRepoRequest {
	r.page = &page
	return r
}

func (r ApiFindAllByOwnerAndRepoRequest) Limit(limit int32) ApiFindAllByOwnerAndRepoRequest {
	r.limit = &limit
	return r
}

func (r ApiFindAllByOwnerAndRepoRequest) OrderDirection(orderDirection OrderDirectionEnum) ApiFindAllByOwnerAndRepoRequest {
	r.orderDirection = &orderDirection
	return r
}

// Range in days
func (r ApiFindAllByOwnerAndRepoRequest) Range_(range_ int32) ApiFindAllByOwnerAndRepoRequest {
	r.range_ = &range_
	return r
}

func (r ApiFindAllByOwnerAndRepoRequest) OrderBy(orderBy RepoOrderFieldsEnum) ApiFindAllByOwnerAndRepoRequest {
	r.orderBy = &orderBy
	return r
}

func (r ApiFindAllByOwnerAndRepoRequest) Execute() (*FindAllByRepoId200Response, *http.Response, error) {
	return r.ApiService.FindAllByOwnerAndRepoExecute(r)
}

/*
FindAllByOwnerAndRepo Finds a repo by :owner and :repo listing all contributions and paginating them

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param owner
	@param repo
	@return ApiFindAllByOwnerAndRepoRequest
*/
func (a *ContributionServiceAPIService) FindAllByOwnerAndRepo(ctx context.Context, owner string, repo string) ApiFindAllByOwnerAndRepoRequest {
	return ApiFindAllByOwnerAndRepoRequest{
		ApiService: a,
		ctx:        ctx,
		owner:      owner,
		repo:       repo,
	}
}

// Execute executes the request
//
//	@return FindAllByRepoId200Response
func (a *ContributionServiceAPIService) FindAllByOwnerAndRepoExecute(r ApiFindAllByOwnerAndRepoRequest) (*FindAllByRepoId200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *FindAllByRepoId200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContributionServiceAPIService.FindAllByOwnerAndRepo")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/repos/{owner}/{repo}/contributions"
	localVarPath = strings.Replace(localVarPath, "{"+"owner"+"}", url.PathEscape(parameterValueToString(r.owner, "owner")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"repo"+"}", url.PathEscape(parameterValueToString(r.repo, "repo")), -1)

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
	if r.orderBy != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "orderBy", r.orderBy, "")
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

type ApiFindAllByRepoIdRequest struct {
	ctx            context.Context
	ApiService     *ContributionServiceAPIService
	id             int32
	page           *int32
	limit          *int32
	orderDirection *OrderDirectionEnum
	range_         *int32
	orderBy        *RepoOrderFieldsEnum
}

func (r ApiFindAllByRepoIdRequest) Page(page int32) ApiFindAllByRepoIdRequest {
	r.page = &page
	return r
}

func (r ApiFindAllByRepoIdRequest) Limit(limit int32) ApiFindAllByRepoIdRequest {
	r.limit = &limit
	return r
}

func (r ApiFindAllByRepoIdRequest) OrderDirection(orderDirection OrderDirectionEnum) ApiFindAllByRepoIdRequest {
	r.orderDirection = &orderDirection
	return r
}

// Range in days
func (r ApiFindAllByRepoIdRequest) Range_(range_ int32) ApiFindAllByRepoIdRequest {
	r.range_ = &range_
	return r
}

func (r ApiFindAllByRepoIdRequest) OrderBy(orderBy RepoOrderFieldsEnum) ApiFindAllByRepoIdRequest {
	r.orderBy = &orderBy
	return r
}

func (r ApiFindAllByRepoIdRequest) Execute() (*FindAllByRepoId200Response, *http.Response, error) {
	return r.ApiService.FindAllByRepoIdExecute(r)
}

/*
FindAllByRepoId Find a repo by :id listing all contributions and paginating them

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id
	@return ApiFindAllByRepoIdRequest
*/
func (a *ContributionServiceAPIService) FindAllByRepoId(ctx context.Context, id int32) ApiFindAllByRepoIdRequest {
	return ApiFindAllByRepoIdRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
//
//	@return FindAllByRepoId200Response
func (a *ContributionServiceAPIService) FindAllByRepoIdExecute(r ApiFindAllByRepoIdRequest) (*FindAllByRepoId200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *FindAllByRepoId200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContributionServiceAPIService.FindAllByRepoId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/repos/{id}/contributions"
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
	if r.orderBy != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "orderBy", r.orderBy, "")
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
