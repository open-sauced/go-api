# \RepositoryServiceAPI

All URIs are relative to *https://api.opensauced.pizza*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FindAll**](RepositoryServiceAPI.md#FindAll) | **Get** /v1/repos/list | Finds all repos and paginates them
[**FindAllByOwnerAndRepo**](RepositoryServiceAPI.md#FindAllByOwnerAndRepo) | **Get** /v1/repos/{owner}/{repo}/contributions | Finds a repo by :owner and :repo listing all contributions and paginating them
[**FindAllByRepoId**](RepositoryServiceAPI.md#FindAllByRepoId) | **Get** /v1/repos/{id}/contributions | Find a repo by :id listing all contributions and paginating them
[**FindAllReposWithFilters**](RepositoryServiceAPI.md#FindAllReposWithFilters) | **Get** /v1/repos/search | Finds all repos using filters and paginates them
[**FindOneById**](RepositoryServiceAPI.md#FindOneById) | **Get** /v1/repos/{id} | Finds a repo by :id
[**FindOneByOwnerAndRepo**](RepositoryServiceAPI.md#FindOneByOwnerAndRepo) | **Get** /v1/repos/{owner}/{repo} | Finds a repo by :owner and :repo



## FindAll

> FindAllTopReposByUsername200Response FindAll(ctx).Page(page).Limit(limit).OrderDirection(orderDirection).Range_(range_).OrderBy(orderBy).Execute()

Finds all repos and paginates them

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    page := float32(8.14) // float32 |  (optional) (default to 1)
    limit := float32(8.14) // float32 |  (optional) (default to 10)
    orderDirection := openapiclient.OrderDirectionEnum("ASC") // OrderDirectionEnum |  (optional)
    range_ := float32(8.14) // float32 | Range in days (optional) (default to 30)
    orderBy := openapiclient.RepoOrderFieldsEnum("issues") // RepoOrderFieldsEnum |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RepositoryServiceAPI.FindAll(context.Background()).Page(page).Limit(limit).OrderDirection(orderDirection).Range_(range_).OrderBy(orderBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RepositoryServiceAPI.FindAll``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindAll`: FindAllTopReposByUsername200Response
    fmt.Fprintf(os.Stdout, "Response from `RepositoryServiceAPI.FindAll`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFindAllRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **float32** |  | [default to 1]
 **limit** | **float32** |  | [default to 10]
 **orderDirection** | [**OrderDirectionEnum**](OrderDirectionEnum.md) |  | 
 **range_** | **float32** | Range in days | [default to 30]
 **orderBy** | [**RepoOrderFieldsEnum**](RepoOrderFieldsEnum.md) |  | 

### Return type

[**FindAllTopReposByUsername200Response**](FindAllTopReposByUsername200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindAllByOwnerAndRepo

> FindAllByRepoId200Response FindAllByOwnerAndRepo(ctx, owner, repo).Page(page).Limit(limit).OrderDirection(orderDirection).Range_(range_).OrderBy(orderBy).Execute()

Finds a repo by :owner and :repo listing all contributions and paginating them

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    owner := "owner_example" // string | 
    repo := "repo_example" // string | 
    page := float32(8.14) // float32 |  (optional) (default to 1)
    limit := float32(8.14) // float32 |  (optional) (default to 10)
    orderDirection := openapiclient.OrderDirectionEnum("ASC") // OrderDirectionEnum |  (optional)
    range_ := float32(8.14) // float32 | Range in days (optional) (default to 30)
    orderBy := openapiclient.RepoOrderFieldsEnum("issues") // RepoOrderFieldsEnum |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RepositoryServiceAPI.FindAllByOwnerAndRepo(context.Background(), owner, repo).Page(page).Limit(limit).OrderDirection(orderDirection).Range_(range_).OrderBy(orderBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RepositoryServiceAPI.FindAllByOwnerAndRepo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindAllByOwnerAndRepo`: FindAllByRepoId200Response
    fmt.Fprintf(os.Stdout, "Response from `RepositoryServiceAPI.FindAllByOwnerAndRepo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**repo** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindAllByOwnerAndRepoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **float32** |  | [default to 1]
 **limit** | **float32** |  | [default to 10]
 **orderDirection** | [**OrderDirectionEnum**](OrderDirectionEnum.md) |  | 
 **range_** | **float32** | Range in days | [default to 30]
 **orderBy** | [**RepoOrderFieldsEnum**](RepoOrderFieldsEnum.md) |  | 

### Return type

[**FindAllByRepoId200Response**](FindAllByRepoId200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindAllByRepoId

> FindAllByRepoId200Response FindAllByRepoId(ctx, id).Page(page).Limit(limit).OrderDirection(orderDirection).Range_(range_).OrderBy(orderBy).Execute()

Find a repo by :id listing all contributions and paginating them

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    id := float32(8.14) // float32 | 
    page := float32(8.14) // float32 |  (optional) (default to 1)
    limit := float32(8.14) // float32 |  (optional) (default to 10)
    orderDirection := openapiclient.OrderDirectionEnum("ASC") // OrderDirectionEnum |  (optional)
    range_ := float32(8.14) // float32 | Range in days (optional) (default to 30)
    orderBy := openapiclient.RepoOrderFieldsEnum("issues") // RepoOrderFieldsEnum |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RepositoryServiceAPI.FindAllByRepoId(context.Background(), id).Page(page).Limit(limit).OrderDirection(orderDirection).Range_(range_).OrderBy(orderBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RepositoryServiceAPI.FindAllByRepoId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindAllByRepoId`: FindAllByRepoId200Response
    fmt.Fprintf(os.Stdout, "Response from `RepositoryServiceAPI.FindAllByRepoId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindAllByRepoIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **float32** |  | [default to 1]
 **limit** | **float32** |  | [default to 10]
 **orderDirection** | [**OrderDirectionEnum**](OrderDirectionEnum.md) |  | 
 **range_** | **float32** | Range in days | [default to 30]
 **orderBy** | [**RepoOrderFieldsEnum**](RepoOrderFieldsEnum.md) |  | 

### Return type

[**FindAllByRepoId200Response**](FindAllByRepoId200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindAllReposWithFilters

> FindAllTopReposByUsername200Response FindAllReposWithFilters(ctx).Page(page).Limit(limit).OrderDirection(orderDirection).Range_(range_).OrderBy(orderBy).Filter(filter).Repo(repo).Topic(topic).RepoIds(repoIds).Execute()

Finds all repos using filters and paginates them

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    page := float32(8.14) // float32 |  (optional) (default to 1)
    limit := float32(8.14) // float32 |  (optional) (default to 10)
    orderDirection := openapiclient.OrderDirectionEnum("ASC") // OrderDirectionEnum |  (optional)
    range_ := float32(8.14) // float32 | Range in days (optional) (default to 30)
    orderBy := openapiclient.RepoOrderFieldsEnum("issues") // RepoOrderFieldsEnum |  (optional)
    filter := openapiclient.InsightFilterFieldsEnum("recent") // InsightFilterFieldsEnum |  (optional)
    repo := "repo_example" // string |  (optional)
    topic := "topic_example" // string |  (optional) (default to "")
    repoIds := "repoIds_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RepositoryServiceAPI.FindAllReposWithFilters(context.Background()).Page(page).Limit(limit).OrderDirection(orderDirection).Range_(range_).OrderBy(orderBy).Filter(filter).Repo(repo).Topic(topic).RepoIds(repoIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RepositoryServiceAPI.FindAllReposWithFilters``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindAllReposWithFilters`: FindAllTopReposByUsername200Response
    fmt.Fprintf(os.Stdout, "Response from `RepositoryServiceAPI.FindAllReposWithFilters`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFindAllReposWithFiltersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **float32** |  | [default to 1]
 **limit** | **float32** |  | [default to 10]
 **orderDirection** | [**OrderDirectionEnum**](OrderDirectionEnum.md) |  | 
 **range_** | **float32** | Range in days | [default to 30]
 **orderBy** | [**RepoOrderFieldsEnum**](RepoOrderFieldsEnum.md) |  | 
 **filter** | [**InsightFilterFieldsEnum**](InsightFilterFieldsEnum.md) |  | 
 **repo** | **string** |  | 
 **topic** | **string** |  | [default to &quot;&quot;]
 **repoIds** | **string** |  | 

### Return type

[**FindAllTopReposByUsername200Response**](FindAllTopReposByUsername200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindOneById

> DbRepo FindOneById(ctx, id).Execute()

Finds a repo by :id

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    id := float32(8.14) // float32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RepositoryServiceAPI.FindOneById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RepositoryServiceAPI.FindOneById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindOneById`: DbRepo
    fmt.Fprintf(os.Stdout, "Response from `RepositoryServiceAPI.FindOneById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindOneByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DbRepo**](DbRepo.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindOneByOwnerAndRepo

> DbRepo FindOneByOwnerAndRepo(ctx, owner, repo).Execute()

Finds a repo by :owner and :repo

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    owner := "owner_example" // string | 
    repo := "repo_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RepositoryServiceAPI.FindOneByOwnerAndRepo(context.Background(), owner, repo).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RepositoryServiceAPI.FindOneByOwnerAndRepo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindOneByOwnerAndRepo`: DbRepo
    fmt.Fprintf(os.Stdout, "Response from `RepositoryServiceAPI.FindOneByOwnerAndRepo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**repo** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindOneByOwnerAndRepoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**DbRepo**](DbRepo.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

