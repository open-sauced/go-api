# \UserCollaborationsServiceAPI

All URIs are relative to *https://api.opensauced.pizza*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddUserCollaboration**](UserCollaborationsServiceAPI.md#AddUserCollaboration) | **Post** /v2/user/collaborations | Adds a new collaboration request for the user
[**FindAllUserCollaborations**](UserCollaborationsServiceAPI.md#FindAllUserCollaborations) | **Get** /v2/user/collaborations | Listing all collaborations for the authenticated user
[**RemoveUserCollaborationById**](UserCollaborationsServiceAPI.md#RemoveUserCollaborationById) | **Delete** /v2/user/collaborations/{id} | Removes the user collaboration request
[**UpdateUserCollaboration**](UserCollaborationsServiceAPI.md#UpdateUserCollaboration) | **Patch** /v2/user/collaborations/{id} | Updates a user collaboration



## AddUserCollaboration

> DbUserCollaboration AddUserCollaboration(ctx).CreateUserCollaborationDto(createUserCollaborationDto).Execute()

Adds a new collaboration request for the user

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
    createUserCollaborationDto := *openapiclient.NewCreateUserCollaborationDto("bdougie", "Come collaborate on a cool project") // CreateUserCollaborationDto | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserCollaborationsServiceAPI.AddUserCollaboration(context.Background()).CreateUserCollaborationDto(createUserCollaborationDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserCollaborationsServiceAPI.AddUserCollaboration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddUserCollaboration`: DbUserCollaboration
    fmt.Fprintf(os.Stdout, "Response from `UserCollaborationsServiceAPI.AddUserCollaboration`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddUserCollaborationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createUserCollaborationDto** | [**CreateUserCollaborationDto**](CreateUserCollaborationDto.md) |  | 

### Return type

[**DbUserCollaboration**](DbUserCollaboration.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindAllUserCollaborations

> FindAllUserCollaborations200Response FindAllUserCollaborations(ctx).Page(page).Limit(limit).OrderDirection(orderDirection).Range_(range_).PrevDaysStartDate(prevDaysStartDate).Execute()

Listing all collaborations for the authenticated user

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
    resp, r, err := apiClient.UserCollaborationsServiceAPI.FindAllUserCollaborations(context.Background()).Page(page).Limit(limit).OrderDirection(orderDirection).Range_(range_).PrevDaysStartDate(prevDaysStartDate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserCollaborationsServiceAPI.FindAllUserCollaborations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindAllUserCollaborations`: FindAllUserCollaborations200Response
    fmt.Fprintf(os.Stdout, "Response from `UserCollaborationsServiceAPI.FindAllUserCollaborations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFindAllUserCollaborationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** |  | [default to 1]
 **limit** | **int32** |  | [default to 10]
 **orderDirection** | [**OrderDirectionEnum**](OrderDirectionEnum.md) |  | 
 **range_** | **int32** | Range in days | [default to 30]
 **prevDaysStartDate** | **int32** | Number of days in the past to start range block | [default to 0]

### Return type

[**FindAllUserCollaborations200Response**](FindAllUserCollaborations200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveUserCollaborationById

> RemoveUserCollaborationById(ctx, id).Execute()

Removes the user collaboration request

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
    r, err := apiClient.UserCollaborationsServiceAPI.RemoveUserCollaborationById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserCollaborationsServiceAPI.RemoveUserCollaborationById``: %v\n", err)
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

Other parameters are passed through a pointer to a apiRemoveUserCollaborationByIdRequest struct via the builder pattern


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


## UpdateUserCollaboration

> DbUserCollaboration UpdateUserCollaboration(ctx, id).UpdateUserCollaborationDto(updateUserCollaborationDto).Execute()

Updates a user collaboration

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
    updateUserCollaborationDto := *openapiclient.NewUpdateUserCollaborationDto("accept") // UpdateUserCollaborationDto | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserCollaborationsServiceAPI.UpdateUserCollaboration(context.Background(), id).UpdateUserCollaborationDto(updateUserCollaborationDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserCollaborationsServiceAPI.UpdateUserCollaboration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateUserCollaboration`: DbUserCollaboration
    fmt.Fprintf(os.Stdout, "Response from `UserCollaborationsServiceAPI.UpdateUserCollaboration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUserCollaborationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateUserCollaborationDto** | [**UpdateUserCollaborationDto**](UpdateUserCollaborationDto.md) |  | 

### Return type

[**DbUserCollaboration**](DbUserCollaboration.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

