# \WorkspaceMembersServiceAPI

All URIs are relative to *https://api.opensauced.pizza*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteOneWorkspaceMemberForUser**](WorkspaceMembersServiceAPI.md#DeleteOneWorkspaceMemberForUser) | **Delete** /v2/workspaces/{id}/members/{memberId} | Delete a workspace member for the authenticated user
[**DeleteWorkspaceMembersForUser**](WorkspaceMembersServiceAPI.md#DeleteWorkspaceMembersForUser) | **Delete** /v2/workspaces/{id}/members | Deletes workspace members for the authenticated user
[**GetWorkspaceMembersForUser**](WorkspaceMembersServiceAPI.md#GetWorkspaceMembersForUser) | **Get** /v2/workspaces/{id}/members | Gets workspace members for the authenticated user
[**UpdateOneWorkspaceMemberForUser**](WorkspaceMembersServiceAPI.md#UpdateOneWorkspaceMemberForUser) | **Patch** /v2/workspaces/{id}/members/{memberId} | Updates a workspace member for the authenticated user
[**UpdateWorkspaceMembersForUser**](WorkspaceMembersServiceAPI.md#UpdateWorkspaceMembersForUser) | **Post** /v2/workspaces/{id}/members | Updates workspace members for the authenticated user



## DeleteOneWorkspaceMemberForUser

> DbWorkspace DeleteOneWorkspaceMemberForUser(ctx, id, memberId).Execute()

Delete a workspace member for the authenticated user

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
    id := "id_example" // string | 
    memberId := "memberId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkspaceMembersServiceAPI.DeleteOneWorkspaceMemberForUser(context.Background(), id, memberId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkspaceMembersServiceAPI.DeleteOneWorkspaceMemberForUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteOneWorkspaceMemberForUser`: DbWorkspace
    fmt.Fprintf(os.Stdout, "Response from `WorkspaceMembersServiceAPI.DeleteOneWorkspaceMemberForUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**memberId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOneWorkspaceMemberForUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**DbWorkspace**](DbWorkspace.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteWorkspaceMembersForUser

> DbWorkspace DeleteWorkspaceMembersForUser(ctx, id).DeleteWorkspaceMembersDto(deleteWorkspaceMembersDto).Execute()

Deletes workspace members for the authenticated user

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
    id := "id_example" // string | 
    deleteWorkspaceMembersDto := *openapiclient.NewDeleteWorkspaceMembersDto([]interface{}{"TODO"}) // DeleteWorkspaceMembersDto | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkspaceMembersServiceAPI.DeleteWorkspaceMembersForUser(context.Background(), id).DeleteWorkspaceMembersDto(deleteWorkspaceMembersDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkspaceMembersServiceAPI.DeleteWorkspaceMembersForUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteWorkspaceMembersForUser`: DbWorkspace
    fmt.Fprintf(os.Stdout, "Response from `WorkspaceMembersServiceAPI.DeleteWorkspaceMembersForUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteWorkspaceMembersForUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **deleteWorkspaceMembersDto** | [**DeleteWorkspaceMembersDto**](DeleteWorkspaceMembersDto.md) |  | 

### Return type

[**DbWorkspace**](DbWorkspace.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWorkspaceMembersForUser

> DbWorkspaceMember GetWorkspaceMembersForUser(ctx, id).Page(page).Limit(limit).OrderDirection(orderDirection).Range_(range_).PrevDaysStartDate(prevDaysStartDate).Execute()

Gets workspace members for the authenticated user

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
    id := "id_example" // string | 
    page := int32(56) // int32 |  (optional) (default to 1)
    limit := int32(56) // int32 |  (optional) (default to 10)
    orderDirection := openapiclient.OrderDirectionEnum("ASC") // OrderDirectionEnum |  (optional)
    range_ := int32(56) // int32 | Range in days (optional) (default to 30)
    prevDaysStartDate := int32(56) // int32 | Number of days in the past to start range block (optional) (default to 0)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkspaceMembersServiceAPI.GetWorkspaceMembersForUser(context.Background(), id).Page(page).Limit(limit).OrderDirection(orderDirection).Range_(range_).PrevDaysStartDate(prevDaysStartDate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkspaceMembersServiceAPI.GetWorkspaceMembersForUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWorkspaceMembersForUser`: DbWorkspaceMember
    fmt.Fprintf(os.Stdout, "Response from `WorkspaceMembersServiceAPI.GetWorkspaceMembersForUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorkspaceMembersForUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** |  | [default to 1]
 **limit** | **int32** |  | [default to 10]
 **orderDirection** | [**OrderDirectionEnum**](OrderDirectionEnum.md) |  | 
 **range_** | **int32** | Range in days | [default to 30]
 **prevDaysStartDate** | **int32** | Number of days in the past to start range block | [default to 0]

### Return type

[**DbWorkspaceMember**](DbWorkspaceMember.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOneWorkspaceMemberForUser

> DbWorkspaceMember UpdateOneWorkspaceMemberForUser(ctx, id, memberId).UpdateWorkspaceMemberDto(updateWorkspaceMemberDto).Execute()

Updates a workspace member for the authenticated user

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
    id := "id_example" // string | 
    memberId := "memberId_example" // string | 
    updateWorkspaceMemberDto := *openapiclient.NewUpdateWorkspaceMemberDto() // UpdateWorkspaceMemberDto | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkspaceMembersServiceAPI.UpdateOneWorkspaceMemberForUser(context.Background(), id, memberId).UpdateWorkspaceMemberDto(updateWorkspaceMemberDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkspaceMembersServiceAPI.UpdateOneWorkspaceMemberForUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOneWorkspaceMemberForUser`: DbWorkspaceMember
    fmt.Fprintf(os.Stdout, "Response from `WorkspaceMembersServiceAPI.UpdateOneWorkspaceMemberForUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**memberId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOneWorkspaceMemberForUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateWorkspaceMemberDto** | [**UpdateWorkspaceMemberDto**](UpdateWorkspaceMemberDto.md) |  | 

### Return type

[**DbWorkspaceMember**](DbWorkspaceMember.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateWorkspaceMembersForUser

> DbWorkspaceMember UpdateWorkspaceMembersForUser(ctx, id).UpdateWorkspaceMembersDto(updateWorkspaceMembersDto).Execute()

Updates workspace members for the authenticated user

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
    id := "id_example" // string | 
    updateWorkspaceMembersDto := *openapiclient.NewUpdateWorkspaceMembersDto([]interface{}{"TODO"}) // UpdateWorkspaceMembersDto | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkspaceMembersServiceAPI.UpdateWorkspaceMembersForUser(context.Background(), id).UpdateWorkspaceMembersDto(updateWorkspaceMembersDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkspaceMembersServiceAPI.UpdateWorkspaceMembersForUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateWorkspaceMembersForUser`: DbWorkspaceMember
    fmt.Fprintf(os.Stdout, "Response from `WorkspaceMembersServiceAPI.UpdateWorkspaceMembersForUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWorkspaceMembersForUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateWorkspaceMembersDto** | [**UpdateWorkspaceMembersDto**](UpdateWorkspaceMembersDto.md) |  | 

### Return type

[**DbWorkspaceMember**](DbWorkspaceMember.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

