# \HighlightsServiceAPI

All URIs are relative to *https://api.opensauced.pizza*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddAFeaturedHighlight**](HighlightsServiceAPI.md#AddAFeaturedHighlight) | **Post** /v1/highlights/{id}/featured | Add a highlight to the featured list
[**FindAllFeaturedHighlights**](HighlightsServiceAPI.md#FindAllFeaturedHighlights) | **Get** /v1/highlights/featured | Finds all featured highlights and paginates them
[**FindAllHighlightRepos**](HighlightsServiceAPI.md#FindAllHighlightRepos) | **Get** /v1/highlights/repos/list | Finds all highlight repos and paginates them
[**FindAllHighlights**](HighlightsServiceAPI.md#FindAllHighlights) | **Get** /v1/highlights/list | Finds all highlights and paginates them
[**FindTopHighlights**](HighlightsServiceAPI.md#FindTopHighlights) | **Get** /v1/highlights/top | Finds top highlights for the day range and paginates them
[**GetAllHighlightReactions**](HighlightsServiceAPI.md#GetAllHighlightReactions) | **Get** /v1/highlights/{id}/reactions | Retrieves total reactions for a highlight
[**RemoveAFeaturedHighlight**](HighlightsServiceAPI.md#RemoveAFeaturedHighlight) | **Delete** /v1/highlights/{id}/featured | Remove a highlight from the featured list



## AddAFeaturedHighlight

> DbUserHighlight AddAFeaturedHighlight(ctx, id).Execute()

Add a highlight to the featured list

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
    resp, r, err := apiClient.HighlightsServiceAPI.AddAFeaturedHighlight(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HighlightsServiceAPI.AddAFeaturedHighlight``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddAFeaturedHighlight`: DbUserHighlight
    fmt.Fprintf(os.Stdout, "Response from `HighlightsServiceAPI.AddAFeaturedHighlight`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddAFeaturedHighlightRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DbUserHighlight**](DbUserHighlight.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindAllFeaturedHighlights

> FindAllHighlightsByUsername200Response FindAllFeaturedHighlights(ctx).Page(page).Limit(limit).OrderDirection(orderDirection).Range_(range_).Repo(repo).Execute()

Finds all featured highlights and paginates them

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
    repo := "open-sauced/insights" // string | Highlight Repo Filter (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HighlightsServiceAPI.FindAllFeaturedHighlights(context.Background()).Page(page).Limit(limit).OrderDirection(orderDirection).Range_(range_).Repo(repo).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HighlightsServiceAPI.FindAllFeaturedHighlights``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindAllFeaturedHighlights`: FindAllHighlightsByUsername200Response
    fmt.Fprintf(os.Stdout, "Response from `HighlightsServiceAPI.FindAllFeaturedHighlights`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFindAllFeaturedHighlightsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** |  | [default to 1]
 **limit** | **int32** |  | [default to 10]
 **orderDirection** | [**OrderDirectionEnum**](OrderDirectionEnum.md) |  | 
 **range_** | **int32** | Range in days | [default to 30]
 **repo** | **string** | Highlight Repo Filter | 

### Return type

[**FindAllHighlightsByUsername200Response**](FindAllHighlightsByUsername200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindAllHighlightRepos

> FindAllHighlightRepos200Response FindAllHighlightRepos(ctx).Page(page).Limit(limit).OrderDirection(orderDirection).Range_(range_).Execute()

Finds all highlight repos and paginates them

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HighlightsServiceAPI.FindAllHighlightRepos(context.Background()).Page(page).Limit(limit).OrderDirection(orderDirection).Range_(range_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HighlightsServiceAPI.FindAllHighlightRepos``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindAllHighlightRepos`: FindAllHighlightRepos200Response
    fmt.Fprintf(os.Stdout, "Response from `HighlightsServiceAPI.FindAllHighlightRepos`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFindAllHighlightReposRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** |  | [default to 1]
 **limit** | **int32** |  | [default to 10]
 **orderDirection** | [**OrderDirectionEnum**](OrderDirectionEnum.md) |  | 
 **range_** | **int32** | Range in days | [default to 30]

### Return type

[**FindAllHighlightRepos200Response**](FindAllHighlightRepos200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindAllHighlights

> FindAllHighlightsByUsername200Response FindAllHighlights(ctx).Page(page).Limit(limit).OrderDirection(orderDirection).Range_(range_).Repo(repo).Execute()

Finds all highlights and paginates them

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
    repo := "open-sauced/insights" // string | Highlight Repo Filter (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HighlightsServiceAPI.FindAllHighlights(context.Background()).Page(page).Limit(limit).OrderDirection(orderDirection).Range_(range_).Repo(repo).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HighlightsServiceAPI.FindAllHighlights``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindAllHighlights`: FindAllHighlightsByUsername200Response
    fmt.Fprintf(os.Stdout, "Response from `HighlightsServiceAPI.FindAllHighlights`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFindAllHighlightsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** |  | [default to 1]
 **limit** | **int32** |  | [default to 10]
 **orderDirection** | [**OrderDirectionEnum**](OrderDirectionEnum.md) |  | 
 **range_** | **int32** | Range in days | [default to 30]
 **repo** | **string** | Highlight Repo Filter | 

### Return type

[**FindAllHighlightsByUsername200Response**](FindAllHighlightsByUsername200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindTopHighlights

> FindAllHighlightsByUsername200Response FindTopHighlights(ctx).Page(page).Limit(limit).OrderDirection(orderDirection).Range_(range_).Execute()

Finds top highlights for the day range and paginates them

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HighlightsServiceAPI.FindTopHighlights(context.Background()).Page(page).Limit(limit).OrderDirection(orderDirection).Range_(range_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HighlightsServiceAPI.FindTopHighlights``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindTopHighlights`: FindAllHighlightsByUsername200Response
    fmt.Fprintf(os.Stdout, "Response from `HighlightsServiceAPI.FindTopHighlights`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFindTopHighlightsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** |  | [default to 1]
 **limit** | **int32** |  | [default to 10]
 **orderDirection** | [**OrderDirectionEnum**](OrderDirectionEnum.md) |  | 
 **range_** | **int32** | Range in days | [default to 30]

### Return type

[**FindAllHighlightsByUsername200Response**](FindAllHighlightsByUsername200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllHighlightReactions

> DbUserHighlightReactionResponse GetAllHighlightReactions(ctx, id).Execute()

Retrieves total reactions for a highlight

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
    resp, r, err := apiClient.HighlightsServiceAPI.GetAllHighlightReactions(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HighlightsServiceAPI.GetAllHighlightReactions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllHighlightReactions`: DbUserHighlightReactionResponse
    fmt.Fprintf(os.Stdout, "Response from `HighlightsServiceAPI.GetAllHighlightReactions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllHighlightReactionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DbUserHighlightReactionResponse**](DbUserHighlightReactionResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveAFeaturedHighlight

> DbUserHighlight RemoveAFeaturedHighlight(ctx, id).Execute()

Remove a highlight from the featured list

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
    resp, r, err := apiClient.HighlightsServiceAPI.RemoveAFeaturedHighlight(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HighlightsServiceAPI.RemoveAFeaturedHighlight``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemoveAFeaturedHighlight`: DbUserHighlight
    fmt.Fprintf(os.Stdout, "Response from `HighlightsServiceAPI.RemoveAFeaturedHighlight`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveAFeaturedHighlightRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DbUserHighlight**](DbUserHighlight.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

