# \SubmitServiceAPI

All URIs are relative to *https://api.opensauced.pizza*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DownSubmitOneById**](SubmitServiceAPI.md#DownSubmitOneById) | **Delete** /v2/repos/{id}/submit | Finds a repo by :id and removes existing submission
[**DownSubmitOneByOwnerAndRepo**](SubmitServiceAPI.md#DownSubmitOneByOwnerAndRepo) | **Delete** /v2/repos/{owner}/{repo}/submit | Finds a repo by :owner and :repo and removes existing submission
[**FindAllUserSubmitted**](SubmitServiceAPI.md#FindAllUserSubmitted) | **Get** /v2/repos/listUserSubmitted | Finds all repos submitted by authenticated user and paginates them
[**SubmitOneById**](SubmitServiceAPI.md#SubmitOneById) | **Put** /v2/repos/{id}/submit | Finds a repo by :id and adds a submission
[**SubmitOneByOwnerAndRepo**](SubmitServiceAPI.md#SubmitOneByOwnerAndRepo) | **Put** /v2/repos/{owner}/{repo}/submit | Finds a repo by :owner and :repo and adds a submission



## DownSubmitOneById

> DbRepoToUserSubmissions DownSubmitOneById(ctx, id).Execute()

Finds a repo by :id and removes existing submission

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/open-sauced/go-api"
)

func main() {
    id := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SubmitServiceAPI.DownSubmitOneById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubmitServiceAPI.DownSubmitOneById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DownSubmitOneById`: DbRepoToUserSubmissions
    fmt.Fprintf(os.Stdout, "Response from `SubmitServiceAPI.DownSubmitOneById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownSubmitOneByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DbRepoToUserSubmissions**](DbRepoToUserSubmissions.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DownSubmitOneByOwnerAndRepo

> DbRepoToUserSubmissions DownSubmitOneByOwnerAndRepo(ctx, owner, repo).Execute()

Finds a repo by :owner and :repo and removes existing submission

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/open-sauced/go-api"
)

func main() {
    owner := "owner_example" // string | 
    repo := "repo_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SubmitServiceAPI.DownSubmitOneByOwnerAndRepo(context.Background(), owner, repo).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubmitServiceAPI.DownSubmitOneByOwnerAndRepo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DownSubmitOneByOwnerAndRepo`: DbRepoToUserSubmissions
    fmt.Fprintf(os.Stdout, "Response from `SubmitServiceAPI.DownSubmitOneByOwnerAndRepo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**repo** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownSubmitOneByOwnerAndRepoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**DbRepoToUserSubmissions**](DbRepoToUserSubmissions.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindAllUserSubmitted

> FindAllTopReposByUsername200Response FindAllUserSubmitted(ctx).Page(page).Limit(limit).OrderDirection(orderDirection).Range_(range_).PrevDaysStartDate(prevDaysStartDate).OrderBy(orderBy).Execute()

Finds all repos submitted by authenticated user and paginates them

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/open-sauced/go-api"
)

func main() {
    page := int32(56) // int32 |  (optional) (default to 1)
    limit := int32(56) // int32 |  (optional) (default to 10)
    orderDirection := openapiclient.OrderDirectionEnum("ASC") // OrderDirectionEnum |  (optional)
    range_ := int32(56) // int32 | Range in days (optional) (default to 30)
    prevDaysStartDate := int32(56) // int32 | Number of days in the past to start range block (optional) (default to 0)
    orderBy := openapiclient.RepoOrderFieldsEnum("issues") // RepoOrderFieldsEnum |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SubmitServiceAPI.FindAllUserSubmitted(context.Background()).Page(page).Limit(limit).OrderDirection(orderDirection).Range_(range_).PrevDaysStartDate(prevDaysStartDate).OrderBy(orderBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubmitServiceAPI.FindAllUserSubmitted``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindAllUserSubmitted`: FindAllTopReposByUsername200Response
    fmt.Fprintf(os.Stdout, "Response from `SubmitServiceAPI.FindAllUserSubmitted`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFindAllUserSubmittedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** |  | [default to 1]
 **limit** | **int32** |  | [default to 10]
 **orderDirection** | [**OrderDirectionEnum**](OrderDirectionEnum.md) |  | 
 **range_** | **int32** | Range in days | [default to 30]
 **prevDaysStartDate** | **int32** | Number of days in the past to start range block | [default to 0]
 **orderBy** | [**RepoOrderFieldsEnum**](RepoOrderFieldsEnum.md) |  | 

### Return type

[**FindAllTopReposByUsername200Response**](FindAllTopReposByUsername200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubmitOneById

> DbRepoToUserSubmissions SubmitOneById(ctx, id).Execute()

Finds a repo by :id and adds a submission

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/open-sauced/go-api"
)

func main() {
    id := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SubmitServiceAPI.SubmitOneById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubmitServiceAPI.SubmitOneById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SubmitOneById`: DbRepoToUserSubmissions
    fmt.Fprintf(os.Stdout, "Response from `SubmitServiceAPI.SubmitOneById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubmitOneByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DbRepoToUserSubmissions**](DbRepoToUserSubmissions.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubmitOneByOwnerAndRepo

> DbRepoToUserSubmissions SubmitOneByOwnerAndRepo(ctx, owner, repo).Execute()

Finds a repo by :owner and :repo and adds a submission

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/open-sauced/go-api"
)

func main() {
    owner := "owner_example" // string | 
    repo := "repo_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SubmitServiceAPI.SubmitOneByOwnerAndRepo(context.Background(), owner, repo).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubmitServiceAPI.SubmitOneByOwnerAndRepo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SubmitOneByOwnerAndRepo`: DbRepoToUserSubmissions
    fmt.Fprintf(os.Stdout, "Response from `SubmitServiceAPI.SubmitOneByOwnerAndRepo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**repo** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubmitOneByOwnerAndRepoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**DbRepoToUserSubmissions**](DbRepoToUserSubmissions.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

