# \WorkspaceContributorsServiceAPI

All URIs are relative to *https://api.opensauced.pizza*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddOneWorkspaceContributorForUser**](WorkspaceContributorsServiceAPI.md#AddOneWorkspaceContributorForUser) | **Post** /v2/workspaces/{id}/contributors/{contributorId} | Updates a workspace contributor for the authenticated user
[**AddWorkspaceContributorsForUser**](WorkspaceContributorsServiceAPI.md#AddWorkspaceContributorsForUser) | **Post** /v2/workspaces/{id}/contributors | Updates workspace contributor for the authenticated user
[**DeleteOneWorkspaceContributorForUser**](WorkspaceContributorsServiceAPI.md#DeleteOneWorkspaceContributorForUser) | **Delete** /v2/workspaces/{id}/contributors/{contributorId} | Delete a workspace contributors for the authenticated user
[**DeleteWorkspaceContributorsForUser**](WorkspaceContributorsServiceAPI.md#DeleteWorkspaceContributorsForUser) | **Delete** /v2/workspaces/{id}/contributors | Deletes workspace contributors for the authenticated user
[**GetWorkspaceContributorsForUser**](WorkspaceContributorsServiceAPI.md#GetWorkspaceContributorsForUser) | **Get** /v2/workspaces/{id}/contributors | Gets workspace contributors for the authenticated user



## AddOneWorkspaceContributorForUser

> DbWorkspaceContributor AddOneWorkspaceContributorForUser(ctx, id, contributorId).Execute()

Updates a workspace contributor for the authenticated user

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
    contributorId := "contributorId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkspaceContributorsServiceAPI.AddOneWorkspaceContributorForUser(context.Background(), id, contributorId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkspaceContributorsServiceAPI.AddOneWorkspaceContributorForUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddOneWorkspaceContributorForUser`: DbWorkspaceContributor
    fmt.Fprintf(os.Stdout, "Response from `WorkspaceContributorsServiceAPI.AddOneWorkspaceContributorForUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**contributorId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddOneWorkspaceContributorForUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**DbWorkspaceContributor**](DbWorkspaceContributor.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddWorkspaceContributorsForUser

> DbWorkspaceContributor AddWorkspaceContributorsForUser(ctx, id).UpdateWorkspaceContributorsDto(updateWorkspaceContributorsDto).Execute()

Updates workspace contributor for the authenticated user

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
    updateWorkspaceContributorsDto := *openapiclient.NewUpdateWorkspaceContributorsDto([]interface{}{"TODO"}) // UpdateWorkspaceContributorsDto | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkspaceContributorsServiceAPI.AddWorkspaceContributorsForUser(context.Background(), id).UpdateWorkspaceContributorsDto(updateWorkspaceContributorsDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkspaceContributorsServiceAPI.AddWorkspaceContributorsForUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddWorkspaceContributorsForUser`: DbWorkspaceContributor
    fmt.Fprintf(os.Stdout, "Response from `WorkspaceContributorsServiceAPI.AddWorkspaceContributorsForUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddWorkspaceContributorsForUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateWorkspaceContributorsDto** | [**UpdateWorkspaceContributorsDto**](UpdateWorkspaceContributorsDto.md) |  | 

### Return type

[**DbWorkspaceContributor**](DbWorkspaceContributor.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOneWorkspaceContributorForUser

> DbWorkspace DeleteOneWorkspaceContributorForUser(ctx, id, contributorId).Execute()

Delete a workspace contributors for the authenticated user

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
    contributorId := "contributorId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkspaceContributorsServiceAPI.DeleteOneWorkspaceContributorForUser(context.Background(), id, contributorId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkspaceContributorsServiceAPI.DeleteOneWorkspaceContributorForUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteOneWorkspaceContributorForUser`: DbWorkspace
    fmt.Fprintf(os.Stdout, "Response from `WorkspaceContributorsServiceAPI.DeleteOneWorkspaceContributorForUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**contributorId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOneWorkspaceContributorForUserRequest struct via the builder pattern


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


## DeleteWorkspaceContributorsForUser

> DbWorkspace DeleteWorkspaceContributorsForUser(ctx, id).DeleteWorkspaceContributorsDto(deleteWorkspaceContributorsDto).Execute()

Deletes workspace contributors for the authenticated user

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
    deleteWorkspaceContributorsDto := *openapiclient.NewDeleteWorkspaceContributorsDto([]interface{}{"TODO"}) // DeleteWorkspaceContributorsDto | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkspaceContributorsServiceAPI.DeleteWorkspaceContributorsForUser(context.Background(), id).DeleteWorkspaceContributorsDto(deleteWorkspaceContributorsDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkspaceContributorsServiceAPI.DeleteWorkspaceContributorsForUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteWorkspaceContributorsForUser`: DbWorkspace
    fmt.Fprintf(os.Stdout, "Response from `WorkspaceContributorsServiceAPI.DeleteWorkspaceContributorsForUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteWorkspaceContributorsForUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **deleteWorkspaceContributorsDto** | [**DeleteWorkspaceContributorsDto**](DeleteWorkspaceContributorsDto.md) |  | 

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


## GetWorkspaceContributorsForUser

> DbWorkspaceContributor GetWorkspaceContributorsForUser(ctx, id).Page(page).Limit(limit).OrderDirection(orderDirection).Range_(range_).PrevDaysStartDate(prevDaysStartDate).Execute()

Gets workspace contributors for the authenticated user

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
    resp, r, err := apiClient.WorkspaceContributorsServiceAPI.GetWorkspaceContributorsForUser(context.Background(), id).Page(page).Limit(limit).OrderDirection(orderDirection).Range_(range_).PrevDaysStartDate(prevDaysStartDate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkspaceContributorsServiceAPI.GetWorkspaceContributorsForUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWorkspaceContributorsForUser`: DbWorkspaceContributor
    fmt.Fprintf(os.Stdout, "Response from `WorkspaceContributorsServiceAPI.GetWorkspaceContributorsForUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorkspaceContributorsForUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** |  | [default to 1]
 **limit** | **int32** |  | [default to 10]
 **orderDirection** | [**OrderDirectionEnum**](OrderDirectionEnum.md) |  | 
 **range_** | **int32** | Range in days | [default to 30]
 **prevDaysStartDate** | **int32** | Number of days in the past to start range block | [default to 0]

### Return type

[**DbWorkspaceContributor**](DbWorkspaceContributor.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

