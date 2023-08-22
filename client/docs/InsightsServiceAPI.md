# \InsightsServiceAPI

All URIs are relative to *https://api.opensauced.pizza*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddInsightForUser**](InsightsServiceAPI.md#AddInsightForUser) | **Post** /v1/user/insights | Adds a new insight page for the authenticated user
[**AddMemberForInsight**](InsightsServiceAPI.md#AddMemberForInsight) | **Post** /v1/user/insights/{id}/members | Adds a new member for the insight
[**FindAllInsightMembers**](InsightsServiceAPI.md#FindAllInsightMembers) | **Get** /v1/user/insights/{id}/members | Listing all members for an insight and paginate them
[**FindAllInsightsByUserId**](InsightsServiceAPI.md#FindAllInsightsByUserId) | **Get** /v1/user/insights | Listing all insights for a user and paginate them
[**FindInsightPageById**](InsightsServiceAPI.md#FindInsightPageById) | **Get** /v1/insights/{id} | Finds a insight page by :id
[**RemoveInsightForUser**](InsightsServiceAPI.md#RemoveInsightForUser) | **Delete** /v1/insights/{id} | Removes an insight page for the authenticated user
[**RemoveInsightMemberById**](InsightsServiceAPI.md#RemoveInsightMemberById) | **Delete** /v1/user/insights/{id}/members/{memberId} | Removes a member from an insight
[**UpdateInsightForUser**](InsightsServiceAPI.md#UpdateInsightForUser) | **Patch** /v1/user/insights/{id} | Updates an insight page for the authenticated user
[**UpdateInsightMember**](InsightsServiceAPI.md#UpdateInsightMember) | **Patch** /v1/user/insights/{id}/members/{memberId} | Updates an insight member information



## AddInsightForUser

> DbInsight AddInsightForUser(ctx).CreateInsightDto(createInsightDto).Execute()

Adds a new insight page for the authenticated user

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
    createInsightDto := *openapiclient.NewCreateInsightDto("My Team", false, []openapiclient.RepoInfo{*openapiclient.NewRepoInfo(float32(234343), "open-sauced/insights")}) // CreateInsightDto | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InsightsServiceAPI.AddInsightForUser(context.Background()).CreateInsightDto(createInsightDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InsightsServiceAPI.AddInsightForUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddInsightForUser`: DbInsight
    fmt.Fprintf(os.Stdout, "Response from `InsightsServiceAPI.AddInsightForUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddInsightForUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createInsightDto** | [**CreateInsightDto**](CreateInsightDto.md) |  | 

### Return type

[**DbInsight**](DbInsight.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddMemberForInsight

> DbInsightMember AddMemberForInsight(ctx, id).CreateInsightMemberDto(createInsightMemberDto).Execute()

Adds a new member for the insight

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
    createInsightMemberDto := *openapiclient.NewCreateInsightMemberDto("hello@opensauced.pizza") // CreateInsightMemberDto | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InsightsServiceAPI.AddMemberForInsight(context.Background(), id).CreateInsightMemberDto(createInsightMemberDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InsightsServiceAPI.AddMemberForInsight``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddMemberForInsight`: DbInsightMember
    fmt.Fprintf(os.Stdout, "Response from `InsightsServiceAPI.AddMemberForInsight`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddMemberForInsightRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createInsightMemberDto** | [**CreateInsightMemberDto**](CreateInsightMemberDto.md) |  | 

### Return type

[**DbInsightMember**](DbInsightMember.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindAllInsightMembers

> FindAllInsightMembers200Response FindAllInsightMembers(ctx, id).Page(page).Limit(limit).OrderDirection(orderDirection).Range_(range_).Execute()

Listing all members for an insight and paginate them

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InsightsServiceAPI.FindAllInsightMembers(context.Background(), id).Page(page).Limit(limit).OrderDirection(orderDirection).Range_(range_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InsightsServiceAPI.FindAllInsightMembers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindAllInsightMembers`: FindAllInsightMembers200Response
    fmt.Fprintf(os.Stdout, "Response from `InsightsServiceAPI.FindAllInsightMembers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindAllInsightMembersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **float32** |  | [default to 1]
 **limit** | **float32** |  | [default to 10]
 **orderDirection** | [**OrderDirectionEnum**](OrderDirectionEnum.md) |  | 
 **range_** | **float32** | Range in days | [default to 30]

### Return type

[**FindAllInsightMembers200Response**](FindAllInsightMembers200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindAllInsightsByUserId

> FindAllInsightsByUserId200Response FindAllInsightsByUserId(ctx).Page(page).Limit(limit).OrderDirection(orderDirection).Range_(range_).Execute()

Listing all insights for a user and paginate them

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InsightsServiceAPI.FindAllInsightsByUserId(context.Background()).Page(page).Limit(limit).OrderDirection(orderDirection).Range_(range_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InsightsServiceAPI.FindAllInsightsByUserId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindAllInsightsByUserId`: FindAllInsightsByUserId200Response
    fmt.Fprintf(os.Stdout, "Response from `InsightsServiceAPI.FindAllInsightsByUserId`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFindAllInsightsByUserIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **float32** |  | [default to 1]
 **limit** | **float32** |  | [default to 10]
 **orderDirection** | [**OrderDirectionEnum**](OrderDirectionEnum.md) |  | 
 **range_** | **float32** | Range in days | [default to 30]

### Return type

[**FindAllInsightsByUserId200Response**](FindAllInsightsByUserId200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindInsightPageById

> DbInsight FindInsightPageById(ctx, id).Execute()

Finds a insight page by :id

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
    resp, r, err := apiClient.InsightsServiceAPI.FindInsightPageById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InsightsServiceAPI.FindInsightPageById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindInsightPageById`: DbInsight
    fmt.Fprintf(os.Stdout, "Response from `InsightsServiceAPI.FindInsightPageById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindInsightPageByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DbInsight**](DbInsight.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveInsightForUser

> DbInsight RemoveInsightForUser(ctx, id).UpdateInsightDto(updateInsightDto).Execute()

Removes an insight page for the authenticated user

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
    updateInsightDto := *openapiclient.NewUpdateInsightDto("My Team", false, []openapiclient.RepoInfo{*openapiclient.NewRepoInfo(float32(234343), "open-sauced/insights")}) // UpdateInsightDto | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InsightsServiceAPI.RemoveInsightForUser(context.Background(), id).UpdateInsightDto(updateInsightDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InsightsServiceAPI.RemoveInsightForUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemoveInsightForUser`: DbInsight
    fmt.Fprintf(os.Stdout, "Response from `InsightsServiceAPI.RemoveInsightForUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveInsightForUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateInsightDto** | [**UpdateInsightDto**](UpdateInsightDto.md) |  | 

### Return type

[**DbInsight**](DbInsight.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveInsightMemberById

> DbInsight RemoveInsightMemberById(ctx, id, memberId).Execute()

Removes a member from an insight

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
    memberId := "memberId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InsightsServiceAPI.RemoveInsightMemberById(context.Background(), id, memberId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InsightsServiceAPI.RemoveInsightMemberById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemoveInsightMemberById`: DbInsight
    fmt.Fprintf(os.Stdout, "Response from `InsightsServiceAPI.RemoveInsightMemberById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **float32** |  | 
**memberId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveInsightMemberByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**DbInsight**](DbInsight.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateInsightForUser

> DbInsight UpdateInsightForUser(ctx, id).UpdateInsightDto(updateInsightDto).Execute()

Updates an insight page for the authenticated user

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
    updateInsightDto := *openapiclient.NewUpdateInsightDto("My Team", false, []openapiclient.RepoInfo{*openapiclient.NewRepoInfo(float32(234343), "open-sauced/insights")}) // UpdateInsightDto | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InsightsServiceAPI.UpdateInsightForUser(context.Background(), id).UpdateInsightDto(updateInsightDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InsightsServiceAPI.UpdateInsightForUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateInsightForUser`: DbInsight
    fmt.Fprintf(os.Stdout, "Response from `InsightsServiceAPI.UpdateInsightForUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateInsightForUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateInsightDto** | [**UpdateInsightDto**](UpdateInsightDto.md) |  | 

### Return type

[**DbInsight**](DbInsight.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateInsightMember

> DbInsight UpdateInsightMember(ctx, id, memberId).UpdateInsightMemberDto(updateInsightMemberDto).Execute()

Updates an insight member information

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
    memberId := "memberId_example" // string | 
    updateInsightMemberDto := *openapiclient.NewUpdateInsightMemberDto("view") // UpdateInsightMemberDto | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InsightsServiceAPI.UpdateInsightMember(context.Background(), id, memberId).UpdateInsightMemberDto(updateInsightMemberDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InsightsServiceAPI.UpdateInsightMember``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateInsightMember`: DbInsight
    fmt.Fprintf(os.Stdout, "Response from `InsightsServiceAPI.UpdateInsightMember`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **float32** |  | 
**memberId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateInsightMemberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateInsightMemberDto** | [**UpdateInsightMemberDto**](UpdateInsightMemberDto.md) |  | 

### Return type

[**DbInsight**](DbInsight.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

