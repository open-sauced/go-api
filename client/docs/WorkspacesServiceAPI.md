# \WorkspacesServiceAPI

All URIs are relative to *https://api.opensauced.pizza*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateWorkspaceForUser**](WorkspacesServiceAPI.md#CreateWorkspaceForUser) | **Post** /v2/workspaces | Create a new workspace for the authenticated user
[**DeleteWorkspaceForUser**](WorkspacesServiceAPI.md#DeleteWorkspaceForUser) | **Delete** /v2/workspaces/{id} | Deletes a workspace for the authenticated user
[**GetWorkspaceByIdForUser**](WorkspacesServiceAPI.md#GetWorkspaceByIdForUser) | **Get** /v2/workspaces/{id} | Gets a workspace for the authenticated user
[**GetWorkspaceForUser**](WorkspacesServiceAPI.md#GetWorkspaceForUser) | **Get** /v2/workspaces | Gets workspaces for the authenticated user
[**UpdateWorkspaceForUser**](WorkspacesServiceAPI.md#UpdateWorkspaceForUser) | **Patch** /v2/workspaces/{id} | Updates a workspace for the authenticated user



## CreateWorkspaceForUser

> DbWorkspace CreateWorkspaceForUser(ctx).CreateWorkspaceDto(createWorkspaceDto).Execute()

Create a new workspace for the authenticated user

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
    createWorkspaceDto := *openapiclient.NewCreateWorkspaceDto("My Workspace", "A workspace for my OSS collaborators", []interface{}{"TODO"}, []interface{}{"TODO"}) // CreateWorkspaceDto | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkspacesServiceAPI.CreateWorkspaceForUser(context.Background()).CreateWorkspaceDto(createWorkspaceDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkspacesServiceAPI.CreateWorkspaceForUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWorkspaceForUser`: DbWorkspace
    fmt.Fprintf(os.Stdout, "Response from `WorkspacesServiceAPI.CreateWorkspaceForUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateWorkspaceForUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createWorkspaceDto** | [**CreateWorkspaceDto**](CreateWorkspaceDto.md) |  | 

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


## DeleteWorkspaceForUser

> DeleteWorkspaceForUser(ctx, id).Execute()

Deletes a workspace for the authenticated user

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.WorkspacesServiceAPI.DeleteWorkspaceForUser(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkspacesServiceAPI.DeleteWorkspaceForUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteWorkspaceForUserRequest struct via the builder pattern


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


## GetWorkspaceByIdForUser

> DbWorkspace GetWorkspaceByIdForUser(ctx, id).Execute()

Gets a workspace for the authenticated user

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkspacesServiceAPI.GetWorkspaceByIdForUser(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkspacesServiceAPI.GetWorkspaceByIdForUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWorkspaceByIdForUser`: DbWorkspace
    fmt.Fprintf(os.Stdout, "Response from `WorkspacesServiceAPI.GetWorkspaceByIdForUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorkspaceByIdForUserRequest struct via the builder pattern


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


## GetWorkspaceForUser

> DbWorkspace GetWorkspaceForUser(ctx).Page(page).Limit(limit).OrderDirection(orderDirection).Range_(range_).PrevDaysStartDate(prevDaysStartDate).Execute()

Gets workspaces for the authenticated user

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
    resp, r, err := apiClient.WorkspacesServiceAPI.GetWorkspaceForUser(context.Background()).Page(page).Limit(limit).OrderDirection(orderDirection).Range_(range_).PrevDaysStartDate(prevDaysStartDate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkspacesServiceAPI.GetWorkspaceForUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWorkspaceForUser`: DbWorkspace
    fmt.Fprintf(os.Stdout, "Response from `WorkspacesServiceAPI.GetWorkspaceForUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetWorkspaceForUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** |  | [default to 1]
 **limit** | **int32** |  | [default to 10]
 **orderDirection** | [**OrderDirectionEnum**](OrderDirectionEnum.md) |  | 
 **range_** | **int32** | Range in days | [default to 30]
 **prevDaysStartDate** | **int32** | Number of days in the past to start range block | [default to 0]

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


## UpdateWorkspaceForUser

> DbWorkspace UpdateWorkspaceForUser(ctx, id).UpdateWorkspaceDto(updateWorkspaceDto).Execute()

Updates a workspace for the authenticated user

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
    updateWorkspaceDto := *openapiclient.NewUpdateWorkspaceDto("My Workspace", "A workspace for my OSS collaborators") // UpdateWorkspaceDto | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkspacesServiceAPI.UpdateWorkspaceForUser(context.Background(), id).UpdateWorkspaceDto(updateWorkspaceDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkspacesServiceAPI.UpdateWorkspaceForUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateWorkspaceForUser`: DbWorkspace
    fmt.Fprintf(os.Stdout, "Response from `WorkspacesServiceAPI.UpdateWorkspaceForUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWorkspaceForUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateWorkspaceDto** | [**UpdateWorkspaceDto**](UpdateWorkspaceDto.md) |  | 

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

