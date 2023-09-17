# \UserListsServiceAPI

All URIs are relative to *https://api.opensauced.pizza*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddListForUser**](UserListsServiceAPI.md#AddListForUser) | **Post** /v1/lists | Adds a new list for the authenticated user
[**DeleteListForUser**](UserListsServiceAPI.md#DeleteListForUser) | **Delete** /v1/lists/{id} | Deletes the list for the authenticated user
[**DeleteUserListContributor**](UserListsServiceAPI.md#DeleteUserListContributor) | **Delete** /v1/lists/{id}/contributors/{userListContributorId} | Delete contributor from an individual user list
[**GetContributors**](UserListsServiceAPI.md#GetContributors) | **Get** /v1/lists/contributors | Retrieves paginated contributors
[**GetListsForUser**](UserListsServiceAPI.md#GetListsForUser) | **Get** /v1/lists | Gets lists for the authenticated user
[**GetUserList**](UserListsServiceAPI.md#GetUserList) | **Get** /v1/lists/{id} | Retrieves an individual user list
[**GetUserListContributors**](UserListsServiceAPI.md#GetUserListContributors) | **Get** /v1/lists/{id}/contributors | Retrieves contributors for an individual user list
[**PostUserListContributors**](UserListsServiceAPI.md#PostUserListContributors) | **Post** /v1/lists/{id}/contributors | Add new contributors to an individual user list
[**UpdateListForUser**](UserListsServiceAPI.md#UpdateListForUser) | **Patch** /v1/lists/{id} | Updates the list for the authenticated user



## AddListForUser

> DbUserList AddListForUser(ctx).CreateUserListDto(createUserListDto).Execute()

Adds a new list for the authenticated user

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
    createUserListDto := *openapiclient.NewCreateUserListDto("My List", false, []int32{int32(123)}) // CreateUserListDto | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserListsServiceAPI.AddListForUser(context.Background()).CreateUserListDto(createUserListDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserListsServiceAPI.AddListForUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddListForUser`: DbUserList
    fmt.Fprintf(os.Stdout, "Response from `UserListsServiceAPI.AddListForUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddListForUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createUserListDto** | [**CreateUserListDto**](CreateUserListDto.md) |  | 

### Return type

[**DbUserList**](DbUserList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteListForUser

> DeleteListForUser(ctx, id).Execute()

Deletes the list for the authenticated user

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
    r, err := apiClient.UserListsServiceAPI.DeleteListForUser(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserListsServiceAPI.DeleteListForUser``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteListForUserRequest struct via the builder pattern


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


## DeleteUserListContributor

> DeleteUserListContributor(ctx, id, userListContributorId).Execute()

Delete contributor from an individual user list

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
    userListContributorId := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.UserListsServiceAPI.DeleteUserListContributor(context.Background(), id, userListContributorId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserListsServiceAPI.DeleteUserListContributor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**userListContributorId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUserListContributorRequest struct via the builder pattern


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


## GetContributors

> GetContributors200Response GetContributors(ctx).Page(page).Limit(limit).Location(location).Timezone(timezone).PrVelocity(prVelocity).Execute()

Retrieves paginated contributors

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
    location := "Denver, Colorado" // string |  (optional)
    timezone := "Mountain Standard Time" // string |  (optional)
    prVelocity := int32(2) // int32 | Less than or equal to the average number of days to merge a PR over the last 30 days (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserListsServiceAPI.GetContributors(context.Background()).Page(page).Limit(limit).Location(location).Timezone(timezone).PrVelocity(prVelocity).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserListsServiceAPI.GetContributors``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetContributors`: GetContributors200Response
    fmt.Fprintf(os.Stdout, "Response from `UserListsServiceAPI.GetContributors`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetContributorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** |  | [default to 1]
 **limit** | **int32** |  | [default to 10]
 **location** | **string** |  | 
 **timezone** | **string** |  | 
 **prVelocity** | **int32** | Less than or equal to the average number of days to merge a PR over the last 30 days | 

### Return type

[**GetContributors200Response**](GetContributors200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetListsForUser

> DbUserList GetListsForUser(ctx).Page(page).Limit(limit).OrderDirection(orderDirection).Range_(range_).Execute()

Gets lists for the authenticated user

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserListsServiceAPI.GetListsForUser(context.Background()).Page(page).Limit(limit).OrderDirection(orderDirection).Range_(range_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserListsServiceAPI.GetListsForUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetListsForUser`: DbUserList
    fmt.Fprintf(os.Stdout, "Response from `UserListsServiceAPI.GetListsForUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetListsForUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** |  | [default to 1]
 **limit** | **int32** |  | [default to 10]
 **orderDirection** | [**OrderDirectionEnum**](OrderDirectionEnum.md) |  | 
 **range_** | **int32** | Range in days | [default to 30]

### Return type

[**DbUserList**](DbUserList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserList

> DbUserList GetUserList(ctx, id).Execute()

Retrieves an individual user list

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
    resp, r, err := apiClient.UserListsServiceAPI.GetUserList(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserListsServiceAPI.GetUserList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUserList`: DbUserList
    fmt.Fprintf(os.Stdout, "Response from `UserListsServiceAPI.GetUserList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DbUserList**](DbUserList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserListContributors

> GetUserListContributors200Response GetUserListContributors(ctx, id).Page(page).Limit(limit).OrderDirection(orderDirection).Range_(range_).Execute()

Retrieves contributors for an individual user list

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserListsServiceAPI.GetUserListContributors(context.Background(), id).Page(page).Limit(limit).OrderDirection(orderDirection).Range_(range_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserListsServiceAPI.GetUserListContributors``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUserListContributors`: GetUserListContributors200Response
    fmt.Fprintf(os.Stdout, "Response from `UserListsServiceAPI.GetUserListContributors`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserListContributorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** |  | [default to 1]
 **limit** | **int32** |  | [default to 10]
 **orderDirection** | [**OrderDirectionEnum**](OrderDirectionEnum.md) |  | 
 **range_** | **int32** | Range in days | [default to 30]

### Return type

[**GetUserListContributors200Response**](GetUserListContributors200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostUserListContributors

> []DbUserListContributor PostUserListContributors(ctx, id).CollaboratorsDto(collaboratorsDto).Execute()

Add new contributors to an individual user list

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
    collaboratorsDto := *openapiclient.NewCollaboratorsDto([]int32{int32(123)}) // CollaboratorsDto | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserListsServiceAPI.PostUserListContributors(context.Background(), id).CollaboratorsDto(collaboratorsDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserListsServiceAPI.PostUserListContributors``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostUserListContributors`: []DbUserListContributor
    fmt.Fprintf(os.Stdout, "Response from `UserListsServiceAPI.PostUserListContributors`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostUserListContributorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **collaboratorsDto** | [**CollaboratorsDto**](CollaboratorsDto.md) |  | 

### Return type

[**[]DbUserListContributor**](DbUserListContributor.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateListForUser

> DbUserList UpdateListForUser(ctx, id).CreateUserListDto(createUserListDto).Execute()

Updates the list for the authenticated user

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
    createUserListDto := *openapiclient.NewCreateUserListDto("My List", false, []int32{int32(123)}) // CreateUserListDto | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserListsServiceAPI.UpdateListForUser(context.Background(), id).CreateUserListDto(createUserListDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserListsServiceAPI.UpdateListForUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateListForUser`: DbUserList
    fmt.Fprintf(os.Stdout, "Response from `UserListsServiceAPI.UpdateListForUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateListForUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createUserListDto** | [**CreateUserListDto**](CreateUserListDto.md) |  | 

### Return type

[**DbUserList**](DbUserList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

