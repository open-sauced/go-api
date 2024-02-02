# \WorkspacesStatsServiceAPI

All URIs are relative to *https://api.opensauced.pizza*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetWorkspaceStatsForUser**](WorkspacesStatsServiceAPI.md#GetWorkspaceStatsForUser) | **Get** /v2/workspaces/{id}/stats | Gets workspace stats for the authenticated user



## GetWorkspaceStatsForUser

> DbWorkspaceStats GetWorkspaceStatsForUser(ctx, id).Range_(range_).PrevDaysStartDate(prevDaysStartDate).Execute()

Gets workspace stats for the authenticated user

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
    range_ := int32(56) // int32 | Range in days (optional) (default to 30)
    prevDaysStartDate := int32(56) // int32 | Number of days in the past to start range block (optional) (default to 0)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkspacesStatsServiceAPI.GetWorkspaceStatsForUser(context.Background(), id).Range_(range_).PrevDaysStartDate(prevDaysStartDate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkspacesStatsServiceAPI.GetWorkspaceStatsForUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWorkspaceStatsForUser`: DbWorkspaceStats
    fmt.Fprintf(os.Stdout, "Response from `WorkspacesStatsServiceAPI.GetWorkspaceStatsForUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorkspaceStatsForUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **range_** | **int32** | Range in days | [default to 30]
 **prevDaysStartDate** | **int32** | Number of days in the past to start range block | [default to 0]

### Return type

[**DbWorkspaceStats**](DbWorkspaceStats.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

