# \UserRecommendationsServiceAPI

All URIs are relative to *https://api.opensauced.pizza*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FindUserOrgsRepoRecommendations**](UserRecommendationsServiceAPI.md#FindUserOrgsRepoRecommendations) | **Get** /v2/user/recommendations/orgs | Listing recommended repos for the authenticated user based on their orgs
[**FindUserRepoRecommendations**](UserRecommendationsServiceAPI.md#FindUserRepoRecommendations) | **Get** /v2/user/recommendations/repos | Listing recommended repos for the authenticated user



## FindUserOrgsRepoRecommendations

> FindAllTopReposByUsername200Response FindUserOrgsRepoRecommendations(ctx).Page(page).Limit(limit).OrderDirection(orderDirection).Range_(range_).PrevDaysStartDate(prevDaysStartDate).Execute()

Listing recommended repos for the authenticated user based on their orgs

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserRecommendationsServiceAPI.FindUserOrgsRepoRecommendations(context.Background()).Page(page).Limit(limit).OrderDirection(orderDirection).Range_(range_).PrevDaysStartDate(prevDaysStartDate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserRecommendationsServiceAPI.FindUserOrgsRepoRecommendations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindUserOrgsRepoRecommendations`: FindAllTopReposByUsername200Response
    fmt.Fprintf(os.Stdout, "Response from `UserRecommendationsServiceAPI.FindUserOrgsRepoRecommendations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFindUserOrgsRepoRecommendationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** |  | [default to 1]
 **limit** | **int32** |  | [default to 10]
 **orderDirection** | [**OrderDirectionEnum**](OrderDirectionEnum.md) |  | 
 **range_** | **int32** | Range in days | [default to 30]
 **prevDaysStartDate** | **int32** | Number of days in the past to start range block | [default to 0]

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


## FindUserRepoRecommendations

> FindAllTopReposByUsername200Response FindUserRepoRecommendations(ctx).Execute()

Listing recommended repos for the authenticated user

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserRecommendationsServiceAPI.FindUserRepoRecommendations(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserRecommendationsServiceAPI.FindUserRepoRecommendations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindUserRepoRecommendations`: FindAllTopReposByUsername200Response
    fmt.Fprintf(os.Stdout, "Response from `UserRecommendationsServiceAPI.FindUserRepoRecommendations`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiFindUserRepoRecommendationsRequest struct via the builder pattern


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

