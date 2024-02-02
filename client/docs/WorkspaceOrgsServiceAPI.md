# \WorkspaceOrgsServiceAPI

All URIs are relative to *https://api.opensauced.pizza*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddOneWorkspaceOrgForUser**](WorkspaceOrgsServiceAPI.md#AddOneWorkspaceOrgForUser) | **Post** /v2/workspaces/{id}/orgs/{orgUserId} | Adds one workspace org for the authenticated user
[**AddWorkspaceOrgsForUser**](WorkspaceOrgsServiceAPI.md#AddWorkspaceOrgsForUser) | **Post** /v2/workspaces/{id}/orgs | Adds workspace orgs for the authenticated user
[**DeleteOneWorkspaceOrgForUser**](WorkspaceOrgsServiceAPI.md#DeleteOneWorkspaceOrgForUser) | **Delete** /v2/workspaces/{id}/orgs/{orgUserId} | Deletes a workspace org for the authenticated user
[**DeleteWorkspaceOrgsForUser**](WorkspaceOrgsServiceAPI.md#DeleteWorkspaceOrgsForUser) | **Delete** /v2/workspaces/{id}/orgs | Deletes workspace orgs for the authenticated user
[**GetWorkspaceOrgsForUser**](WorkspaceOrgsServiceAPI.md#GetWorkspaceOrgsForUser) | **Get** /v2/workspaces/{id}/orgs | Gets workspace orgs for the authenticated user



## AddOneWorkspaceOrgForUser

> DbWorkspaceOrg AddOneWorkspaceOrgForUser(ctx, id, orgUserId).Execute()

Adds one workspace org for the authenticated user

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
    orgUserId := float32(8.14) // float32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkspaceOrgsServiceAPI.AddOneWorkspaceOrgForUser(context.Background(), id, orgUserId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkspaceOrgsServiceAPI.AddOneWorkspaceOrgForUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddOneWorkspaceOrgForUser`: DbWorkspaceOrg
    fmt.Fprintf(os.Stdout, "Response from `WorkspaceOrgsServiceAPI.AddOneWorkspaceOrgForUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**orgUserId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddOneWorkspaceOrgForUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**DbWorkspaceOrg**](DbWorkspaceOrg.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddWorkspaceOrgsForUser

> DbWorkspaceOrg AddWorkspaceOrgsForUser(ctx, id).UpdateWorkspaceOrgsDto(updateWorkspaceOrgsDto).Execute()

Adds workspace orgs for the authenticated user

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
    updateWorkspaceOrgsDto := *openapiclient.NewUpdateWorkspaceOrgsDto([]interface{}{"TODO"}) // UpdateWorkspaceOrgsDto | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkspaceOrgsServiceAPI.AddWorkspaceOrgsForUser(context.Background(), id).UpdateWorkspaceOrgsDto(updateWorkspaceOrgsDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkspaceOrgsServiceAPI.AddWorkspaceOrgsForUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddWorkspaceOrgsForUser`: DbWorkspaceOrg
    fmt.Fprintf(os.Stdout, "Response from `WorkspaceOrgsServiceAPI.AddWorkspaceOrgsForUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddWorkspaceOrgsForUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateWorkspaceOrgsDto** | [**UpdateWorkspaceOrgsDto**](UpdateWorkspaceOrgsDto.md) |  | 

### Return type

[**DbWorkspaceOrg**](DbWorkspaceOrg.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOneWorkspaceOrgForUser

> DbWorkspace DeleteOneWorkspaceOrgForUser(ctx, id, orgUserId).Execute()

Deletes a workspace org for the authenticated user

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
    orgUserId := float32(8.14) // float32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkspaceOrgsServiceAPI.DeleteOneWorkspaceOrgForUser(context.Background(), id, orgUserId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkspaceOrgsServiceAPI.DeleteOneWorkspaceOrgForUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteOneWorkspaceOrgForUser`: DbWorkspace
    fmt.Fprintf(os.Stdout, "Response from `WorkspaceOrgsServiceAPI.DeleteOneWorkspaceOrgForUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**orgUserId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOneWorkspaceOrgForUserRequest struct via the builder pattern


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


## DeleteWorkspaceOrgsForUser

> DbWorkspace DeleteWorkspaceOrgsForUser(ctx, id).DeleteWorkspaceOrgsDto(deleteWorkspaceOrgsDto).Execute()

Deletes workspace orgs for the authenticated user

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
    deleteWorkspaceOrgsDto := *openapiclient.NewDeleteWorkspaceOrgsDto([]interface{}{"TODO"}) // DeleteWorkspaceOrgsDto | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkspaceOrgsServiceAPI.DeleteWorkspaceOrgsForUser(context.Background(), id).DeleteWorkspaceOrgsDto(deleteWorkspaceOrgsDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkspaceOrgsServiceAPI.DeleteWorkspaceOrgsForUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteWorkspaceOrgsForUser`: DbWorkspace
    fmt.Fprintf(os.Stdout, "Response from `WorkspaceOrgsServiceAPI.DeleteWorkspaceOrgsForUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteWorkspaceOrgsForUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **deleteWorkspaceOrgsDto** | [**DeleteWorkspaceOrgsDto**](DeleteWorkspaceOrgsDto.md) |  | 

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


## GetWorkspaceOrgsForUser

> DbWorkspaceOrg GetWorkspaceOrgsForUser(ctx, id).Page(page).Limit(limit).OrderDirection(orderDirection).Range_(range_).PrevDaysStartDate(prevDaysStartDate).Execute()

Gets workspace orgs for the authenticated user

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
    resp, r, err := apiClient.WorkspaceOrgsServiceAPI.GetWorkspaceOrgsForUser(context.Background(), id).Page(page).Limit(limit).OrderDirection(orderDirection).Range_(range_).PrevDaysStartDate(prevDaysStartDate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkspaceOrgsServiceAPI.GetWorkspaceOrgsForUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWorkspaceOrgsForUser`: DbWorkspaceOrg
    fmt.Fprintf(os.Stdout, "Response from `WorkspaceOrgsServiceAPI.GetWorkspaceOrgsForUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorkspaceOrgsForUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** |  | [default to 1]
 **limit** | **int32** |  | [default to 10]
 **orderDirection** | [**OrderDirectionEnum**](OrderDirectionEnum.md) |  | 
 **range_** | **int32** | Range in days | [default to 30]
 **prevDaysStartDate** | **int32** | Number of days in the past to start range block | [default to 0]

### Return type

[**DbWorkspaceOrg**](DbWorkspaceOrg.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

