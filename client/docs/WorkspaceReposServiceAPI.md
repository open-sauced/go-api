# \WorkspaceReposServiceAPI

All URIs are relative to *https://api.opensauced.pizza*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddOneWorkspaceRepoForUser**](WorkspaceReposServiceAPI.md#AddOneWorkspaceRepoForUser) | **Post** /v2/workspaces/{id}/repos/{owner}/{repo} | Adds workspace repos for the authenticated user
[**AddWorkspaceReposForUser**](WorkspaceReposServiceAPI.md#AddWorkspaceReposForUser) | **Post** /v2/workspaces/{id}/repos | Adds workspace repos for the authenticated user
[**DeleteOneWorkspaceRepoForUser**](WorkspaceReposServiceAPI.md#DeleteOneWorkspaceRepoForUser) | **Delete** /v2/workspaces/{id}/repos/{owner}/{repo} | Delete a workspace repos for the authenticated user
[**DeleteWorkspaceReposForUser**](WorkspaceReposServiceAPI.md#DeleteWorkspaceReposForUser) | **Delete** /v2/workspaces/{id}/repos | Deletes workspace repos for the authenticated user
[**GetWorkspaceReposForUser**](WorkspaceReposServiceAPI.md#GetWorkspaceReposForUser) | **Get** /v2/workspaces/{id}/repos | Gets workspace repos for the authenticated user



## AddOneWorkspaceRepoForUser

> DbWorkspaceRepo AddOneWorkspaceRepoForUser(ctx, id, owner, repo).Execute()

Adds workspace repos for the authenticated user

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
    owner := "owner_example" // string | 
    repo := "repo_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkspaceReposServiceAPI.AddOneWorkspaceRepoForUser(context.Background(), id, owner, repo).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkspaceReposServiceAPI.AddOneWorkspaceRepoForUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddOneWorkspaceRepoForUser`: DbWorkspaceRepo
    fmt.Fprintf(os.Stdout, "Response from `WorkspaceReposServiceAPI.AddOneWorkspaceRepoForUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**owner** | **string** |  | 
**repo** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddOneWorkspaceRepoForUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**DbWorkspaceRepo**](DbWorkspaceRepo.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddWorkspaceReposForUser

> DbWorkspaceRepo AddWorkspaceReposForUser(ctx, id).UpdateWorkspaceReposDto(updateWorkspaceReposDto).Execute()

Adds workspace repos for the authenticated user

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
    updateWorkspaceReposDto := *openapiclient.NewUpdateWorkspaceReposDto([]interface{}{"TODO"}) // UpdateWorkspaceReposDto | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkspaceReposServiceAPI.AddWorkspaceReposForUser(context.Background(), id).UpdateWorkspaceReposDto(updateWorkspaceReposDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkspaceReposServiceAPI.AddWorkspaceReposForUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddWorkspaceReposForUser`: DbWorkspaceRepo
    fmt.Fprintf(os.Stdout, "Response from `WorkspaceReposServiceAPI.AddWorkspaceReposForUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddWorkspaceReposForUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateWorkspaceReposDto** | [**UpdateWorkspaceReposDto**](UpdateWorkspaceReposDto.md) |  | 

### Return type

[**DbWorkspaceRepo**](DbWorkspaceRepo.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOneWorkspaceRepoForUser

> DbWorkspace DeleteOneWorkspaceRepoForUser(ctx, id, owner, repo).Execute()

Delete a workspace repos for the authenticated user

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
    owner := "owner_example" // string | 
    repo := "repo_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkspaceReposServiceAPI.DeleteOneWorkspaceRepoForUser(context.Background(), id, owner, repo).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkspaceReposServiceAPI.DeleteOneWorkspaceRepoForUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteOneWorkspaceRepoForUser`: DbWorkspace
    fmt.Fprintf(os.Stdout, "Response from `WorkspaceReposServiceAPI.DeleteOneWorkspaceRepoForUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**owner** | **string** |  | 
**repo** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOneWorkspaceRepoForUserRequest struct via the builder pattern


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


## DeleteWorkspaceReposForUser

> DbWorkspace DeleteWorkspaceReposForUser(ctx, id).DeleteWorkspaceReposDto(deleteWorkspaceReposDto).Execute()

Deletes workspace repos for the authenticated user

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
    deleteWorkspaceReposDto := *openapiclient.NewDeleteWorkspaceReposDto([]interface{}{"TODO"}) // DeleteWorkspaceReposDto | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkspaceReposServiceAPI.DeleteWorkspaceReposForUser(context.Background(), id).DeleteWorkspaceReposDto(deleteWorkspaceReposDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkspaceReposServiceAPI.DeleteWorkspaceReposForUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteWorkspaceReposForUser`: DbWorkspace
    fmt.Fprintf(os.Stdout, "Response from `WorkspaceReposServiceAPI.DeleteWorkspaceReposForUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteWorkspaceReposForUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **deleteWorkspaceReposDto** | [**DeleteWorkspaceReposDto**](DeleteWorkspaceReposDto.md) |  | 

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


## GetWorkspaceReposForUser

> DbWorkspaceRepo GetWorkspaceReposForUser(ctx, id).Page(page).Limit(limit).OrderDirection(orderDirection).Range_(range_).PrevDaysStartDate(prevDaysStartDate).Execute()

Gets workspace repos for the authenticated user

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
    resp, r, err := apiClient.WorkspaceReposServiceAPI.GetWorkspaceReposForUser(context.Background(), id).Page(page).Limit(limit).OrderDirection(orderDirection).Range_(range_).PrevDaysStartDate(prevDaysStartDate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkspaceReposServiceAPI.GetWorkspaceReposForUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWorkspaceReposForUser`: DbWorkspaceRepo
    fmt.Fprintf(os.Stdout, "Response from `WorkspaceReposServiceAPI.GetWorkspaceReposForUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorkspaceReposForUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** |  | [default to 1]
 **limit** | **int32** |  | [default to 10]
 **orderDirection** | [**OrderDirectionEnum**](OrderDirectionEnum.md) |  | 
 **range_** | **int32** | Range in days | [default to 30]
 **prevDaysStartDate** | **int32** | Number of days in the past to start range block | [default to 0]

### Return type

[**DbWorkspaceRepo**](DbWorkspaceRepo.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

