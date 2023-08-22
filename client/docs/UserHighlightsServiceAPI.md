# \UserHighlightsServiceAPI

All URIs are relative to *https://api.opensauced.pizza*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddHighlightForUser**](UserHighlightsServiceAPI.md#AddHighlightForUser) | **Post** /v1/user/highlights | Adds a new highlight for the authenticated user
[**AddHighlightReactionForUser**](UserHighlightsServiceAPI.md#AddHighlightReactionForUser) | **Post** /v1/user/highlights/{id}/reactions/{emojiId} | Adds a new highlight reaction for the authenticated user
[**DeleteHighlightForUser**](UserHighlightsServiceAPI.md#DeleteHighlightForUser) | **Delete** /v1/user/highlights/{id} | Deletes the highlight for the authenticated user
[**DeleteHighlightReactionForUser**](UserHighlightsServiceAPI.md#DeleteHighlightReactionForUser) | **Delete** /v1/user/highlights/{id}/reactions/{emojiId} | Deletes the highlight reaction for the authenticated user
[**GetAllHighlightUserReactions**](UserHighlightsServiceAPI.md#GetAllHighlightUserReactions) | **Get** /v1/user/highlights/{id}/reactions | Retrieves a highlight reactions for the authenticated user
[**GetFollowingHighlightRepos**](UserHighlightsServiceAPI.md#GetFollowingHighlightRepos) | **Get** /v1/user/highlights/following/repos | Fetches highlight repos for users the authenticated user follows
[**GetFollowingHighlights**](UserHighlightsServiceAPI.md#GetFollowingHighlights) | **Get** /v1/user/highlights/following | Fetches highlights for users the authenticated user follows
[**GetUserHighlight**](UserHighlightsServiceAPI.md#GetUserHighlight) | **Get** /v1/user/highlights/{id} | Retrieves a highlight
[**UpdateHighlightForUser**](UserHighlightsServiceAPI.md#UpdateHighlightForUser) | **Patch** /v1/user/highlights/{id} | Updates the highlight for the authenticated user



## AddHighlightForUser

> OmitTypeClass AddHighlightForUser(ctx).CreateUserHighlightDto(createUserHighlightDto).Execute()

Adds a new highlight for the authenticated user

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
    createUserHighlightDto := *openapiclient.NewCreateUserHighlightDto("github.com/open-sauced/insights/pull/1", "My first PR to Open Sauced!", "pull_request", []Array{"TODO"}) // CreateUserHighlightDto | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserHighlightsServiceAPI.AddHighlightForUser(context.Background()).CreateUserHighlightDto(createUserHighlightDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserHighlightsServiceAPI.AddHighlightForUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddHighlightForUser`: OmitTypeClass
    fmt.Fprintf(os.Stdout, "Response from `UserHighlightsServiceAPI.AddHighlightForUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddHighlightForUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createUserHighlightDto** | [**CreateUserHighlightDto**](CreateUserHighlightDto.md) |  | 

### Return type

[**OmitTypeClass**](OmitTypeClass.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddHighlightReactionForUser

> DbUserHighlightReaction AddHighlightReactionForUser(ctx, id, emojiId).Execute()

Adds a new highlight reaction for the authenticated user

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
    emojiId := "emojiId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserHighlightsServiceAPI.AddHighlightReactionForUser(context.Background(), id, emojiId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserHighlightsServiceAPI.AddHighlightReactionForUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddHighlightReactionForUser`: DbUserHighlightReaction
    fmt.Fprintf(os.Stdout, "Response from `UserHighlightsServiceAPI.AddHighlightReactionForUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **float32** |  | 
**emojiId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddHighlightReactionForUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**DbUserHighlightReaction**](DbUserHighlightReaction.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteHighlightForUser

> DeleteHighlightForUser(ctx, id).Execute()

Deletes the highlight for the authenticated user

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
    r, err := apiClient.UserHighlightsServiceAPI.DeleteHighlightForUser(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserHighlightsServiceAPI.DeleteHighlightForUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteHighlightForUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteHighlightReactionForUser

> DeleteHighlightReactionForUser(ctx, id, emojiId).Execute()

Deletes the highlight reaction for the authenticated user

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
    emojiId := "emojiId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.UserHighlightsServiceAPI.DeleteHighlightReactionForUser(context.Background(), id, emojiId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserHighlightsServiceAPI.DeleteHighlightReactionForUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **float32** |  | 
**emojiId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteHighlightReactionForUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllHighlightUserReactions

> DbUserHighlightReactionResponse GetAllHighlightUserReactions(ctx, id).Execute()

Retrieves a highlight reactions for the authenticated user

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
    resp, r, err := apiClient.UserHighlightsServiceAPI.GetAllHighlightUserReactions(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserHighlightsServiceAPI.GetAllHighlightUserReactions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllHighlightUserReactions`: DbUserHighlightReactionResponse
    fmt.Fprintf(os.Stdout, "Response from `UserHighlightsServiceAPI.GetAllHighlightUserReactions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllHighlightUserReactionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DbUserHighlightReactionResponse**](DbUserHighlightReactionResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFollowingHighlightRepos

> DbUserHighlightRepo GetFollowingHighlightRepos(ctx).Page(page).Limit(limit).OrderDirection(orderDirection).Range_(range_).Repo(repo).Execute()

Fetches highlight repos for users the authenticated user follows

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
    repo := "open-sauced/insights" // string | Highlight Repo Filter (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserHighlightsServiceAPI.GetFollowingHighlightRepos(context.Background()).Page(page).Limit(limit).OrderDirection(orderDirection).Range_(range_).Repo(repo).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserHighlightsServiceAPI.GetFollowingHighlightRepos``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFollowingHighlightRepos`: DbUserHighlightRepo
    fmt.Fprintf(os.Stdout, "Response from `UserHighlightsServiceAPI.GetFollowingHighlightRepos`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFollowingHighlightReposRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **float32** |  | [default to 1]
 **limit** | **float32** |  | [default to 10]
 **orderDirection** | [**OrderDirectionEnum**](OrderDirectionEnum.md) |  | 
 **range_** | **float32** | Range in days | [default to 30]
 **repo** | **string** | Highlight Repo Filter | 

### Return type

[**DbUserHighlightRepo**](DbUserHighlightRepo.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFollowingHighlights

> DbUserHighlight GetFollowingHighlights(ctx).Page(page).Limit(limit).OrderDirection(orderDirection).Range_(range_).Repo(repo).Execute()

Fetches highlights for users the authenticated user follows

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
    repo := "open-sauced/insights" // string | Highlight Repo Filter (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserHighlightsServiceAPI.GetFollowingHighlights(context.Background()).Page(page).Limit(limit).OrderDirection(orderDirection).Range_(range_).Repo(repo).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserHighlightsServiceAPI.GetFollowingHighlights``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFollowingHighlights`: DbUserHighlight
    fmt.Fprintf(os.Stdout, "Response from `UserHighlightsServiceAPI.GetFollowingHighlights`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFollowingHighlightsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **float32** |  | [default to 1]
 **limit** | **float32** |  | [default to 10]
 **orderDirection** | [**OrderDirectionEnum**](OrderDirectionEnum.md) |  | 
 **range_** | **float32** | Range in days | [default to 30]
 **repo** | **string** | Highlight Repo Filter | 

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


## GetUserHighlight

> DbUserHighlight GetUserHighlight(ctx, id).Execute()

Retrieves a highlight

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
    resp, r, err := apiClient.UserHighlightsServiceAPI.GetUserHighlight(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserHighlightsServiceAPI.GetUserHighlight``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUserHighlight`: DbUserHighlight
    fmt.Fprintf(os.Stdout, "Response from `UserHighlightsServiceAPI.GetUserHighlight`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserHighlightRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DbUserHighlight**](DbUserHighlight.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateHighlightForUser

> DbUserHighlight UpdateHighlightForUser(ctx, id).CreateUserHighlightDto(createUserHighlightDto).Execute()

Updates the highlight for the authenticated user

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
    createUserHighlightDto := *openapiclient.NewCreateUserHighlightDto("github.com/open-sauced/insights/pull/1", "My first PR to Open Sauced!", "pull_request", []Array{"TODO"}) // CreateUserHighlightDto | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserHighlightsServiceAPI.UpdateHighlightForUser(context.Background(), id).CreateUserHighlightDto(createUserHighlightDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserHighlightsServiceAPI.UpdateHighlightForUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateHighlightForUser`: DbUserHighlight
    fmt.Fprintf(os.Stdout, "Response from `UserHighlightsServiceAPI.UpdateHighlightForUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateHighlightForUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createUserHighlightDto** | [**CreateUserHighlightDto**](CreateUserHighlightDto.md) |  | 

### Return type

[**DbUserHighlight**](DbUserHighlight.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

