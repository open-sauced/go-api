# \ContributorsServiceAPI

All URIs are relative to *https://api.opensauced.pizza*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FindAllChurnPullRequestContributors**](ContributorsServiceAPI.md#FindAllChurnPullRequestContributors) | **Get** /v1/contributors/insights/churn | Gets all recent churned contributors for the last 30 days based on repo IDs
[**FindAllRecentPullRequestContributors**](ContributorsServiceAPI.md#FindAllRecentPullRequestContributors) | **Get** /v1/contributors/insights/recent | Gets all recent contributors for the last 30 days based on repo IDs
[**FindAllRepeatPullRequestContributors**](ContributorsServiceAPI.md#FindAllRepeatPullRequestContributors) | **Get** /v1/contributors/insights/repeat | Gets all recent repeat contributors for the last 30 days based on repo IDs
[**NewPullRequestContributors**](ContributorsServiceAPI.md#NewPullRequestContributors) | **Get** /v1/contributors/insights/new | Gets new contributors given a date range for repo IDs
[**SearchAllPullRequestContributors**](ContributorsServiceAPI.md#SearchAllPullRequestContributors) | **Get** /v1/contributors/search | Searches contributors from pull requests using filters and paginates them



## FindAllChurnPullRequestContributors

> SearchAllPullRequestContributors200Response FindAllChurnPullRequestContributors(ctx).Page(page).Limit(limit).OrderDirection(orderDirection).Range_(range_).RepoIds(repoIds).Execute()

Gets all recent churned contributors for the last 30 days based on repo IDs

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
    repoIds := "repoIds_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContributorsServiceAPI.FindAllChurnPullRequestContributors(context.Background()).Page(page).Limit(limit).OrderDirection(orderDirection).Range_(range_).RepoIds(repoIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContributorsServiceAPI.FindAllChurnPullRequestContributors``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindAllChurnPullRequestContributors`: SearchAllPullRequestContributors200Response
    fmt.Fprintf(os.Stdout, "Response from `ContributorsServiceAPI.FindAllChurnPullRequestContributors`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFindAllChurnPullRequestContributorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **float32** |  | [default to 1]
 **limit** | **float32** |  | [default to 10]
 **orderDirection** | [**OrderDirectionEnum**](OrderDirectionEnum.md) |  | 
 **range_** | **float32** | Range in days | [default to 30]
 **repoIds** | **string** |  | 

### Return type

[**SearchAllPullRequestContributors200Response**](SearchAllPullRequestContributors200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindAllRecentPullRequestContributors

> SearchAllPullRequestContributors200Response FindAllRecentPullRequestContributors(ctx).Page(page).Limit(limit).OrderDirection(orderDirection).Range_(range_).RepoIds(repoIds).Execute()

Gets all recent contributors for the last 30 days based on repo IDs

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
    repoIds := "repoIds_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContributorsServiceAPI.FindAllRecentPullRequestContributors(context.Background()).Page(page).Limit(limit).OrderDirection(orderDirection).Range_(range_).RepoIds(repoIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContributorsServiceAPI.FindAllRecentPullRequestContributors``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindAllRecentPullRequestContributors`: SearchAllPullRequestContributors200Response
    fmt.Fprintf(os.Stdout, "Response from `ContributorsServiceAPI.FindAllRecentPullRequestContributors`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFindAllRecentPullRequestContributorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **float32** |  | [default to 1]
 **limit** | **float32** |  | [default to 10]
 **orderDirection** | [**OrderDirectionEnum**](OrderDirectionEnum.md) |  | 
 **range_** | **float32** | Range in days | [default to 30]
 **repoIds** | **string** |  | 

### Return type

[**SearchAllPullRequestContributors200Response**](SearchAllPullRequestContributors200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindAllRepeatPullRequestContributors

> SearchAllPullRequestContributors200Response FindAllRepeatPullRequestContributors(ctx).Page(page).Limit(limit).OrderDirection(orderDirection).Range_(range_).RepoIds(repoIds).Execute()

Gets all recent repeat contributors for the last 30 days based on repo IDs

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
    repoIds := "repoIds_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContributorsServiceAPI.FindAllRepeatPullRequestContributors(context.Background()).Page(page).Limit(limit).OrderDirection(orderDirection).Range_(range_).RepoIds(repoIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContributorsServiceAPI.FindAllRepeatPullRequestContributors``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindAllRepeatPullRequestContributors`: SearchAllPullRequestContributors200Response
    fmt.Fprintf(os.Stdout, "Response from `ContributorsServiceAPI.FindAllRepeatPullRequestContributors`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFindAllRepeatPullRequestContributorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **float32** |  | [default to 1]
 **limit** | **float32** |  | [default to 10]
 **orderDirection** | [**OrderDirectionEnum**](OrderDirectionEnum.md) |  | 
 **range_** | **float32** | Range in days | [default to 30]
 **repoIds** | **string** |  | 

### Return type

[**SearchAllPullRequestContributors200Response**](SearchAllPullRequestContributors200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NewPullRequestContributors

> SearchAllPullRequestContributors200Response NewPullRequestContributors(ctx).Page(page).Limit(limit).OrderDirection(orderDirection).Range_(range_).Filter(filter).Topic(topic).Repo(repo).RepoIds(repoIds).Execute()

Gets new contributors given a date range for repo IDs

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
    filter := openapiclient.InsightFilterFieldsEnum("recent") // InsightFilterFieldsEnum |  (optional)
    topic := "javascript" // string |  (optional)
    repo := "open-sauced/insights" // string |  (optional)
    repoIds := "repoIds_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContributorsServiceAPI.NewPullRequestContributors(context.Background()).Page(page).Limit(limit).OrderDirection(orderDirection).Range_(range_).Filter(filter).Topic(topic).Repo(repo).RepoIds(repoIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContributorsServiceAPI.NewPullRequestContributors``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NewPullRequestContributors`: SearchAllPullRequestContributors200Response
    fmt.Fprintf(os.Stdout, "Response from `ContributorsServiceAPI.NewPullRequestContributors`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNewPullRequestContributorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **float32** |  | [default to 1]
 **limit** | **float32** |  | [default to 10]
 **orderDirection** | [**OrderDirectionEnum**](OrderDirectionEnum.md) |  | 
 **range_** | **float32** | Range in days | [default to 30]
 **filter** | [**InsightFilterFieldsEnum**](InsightFilterFieldsEnum.md) |  | 
 **topic** | **string** |  | 
 **repo** | **string** |  | 
 **repoIds** | **string** |  | 

### Return type

[**SearchAllPullRequestContributors200Response**](SearchAllPullRequestContributors200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchAllPullRequestContributors

> SearchAllPullRequestContributors200Response SearchAllPullRequestContributors(ctx).Page(page).Limit(limit).OrderDirection(orderDirection).Range_(range_).Filter(filter).Topic(topic).Repo(repo).RepoIds(repoIds).Execute()

Searches contributors from pull requests using filters and paginates them

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
    filter := openapiclient.InsightFilterFieldsEnum("recent") // InsightFilterFieldsEnum |  (optional)
    topic := "javascript" // string |  (optional)
    repo := "open-sauced/insights" // string |  (optional)
    repoIds := "repoIds_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContributorsServiceAPI.SearchAllPullRequestContributors(context.Background()).Page(page).Limit(limit).OrderDirection(orderDirection).Range_(range_).Filter(filter).Topic(topic).Repo(repo).RepoIds(repoIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContributorsServiceAPI.SearchAllPullRequestContributors``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchAllPullRequestContributors`: SearchAllPullRequestContributors200Response
    fmt.Fprintf(os.Stdout, "Response from `ContributorsServiceAPI.SearchAllPullRequestContributors`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchAllPullRequestContributorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **float32** |  | [default to 1]
 **limit** | **float32** |  | [default to 10]
 **orderDirection** | [**OrderDirectionEnum**](OrderDirectionEnum.md) |  | 
 **range_** | **float32** | Range in days | [default to 30]
 **filter** | [**InsightFilterFieldsEnum**](InsightFilterFieldsEnum.md) |  | 
 **topic** | **string** |  | 
 **repo** | **string** |  | 
 **repoIds** | **string** |  | 

### Return type

[**SearchAllPullRequestContributors200Response**](SearchAllPullRequestContributors200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

