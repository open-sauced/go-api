# \EndorsementsServiceAPI

All URIs are relative to *https://api.opensauced.pizza*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateEndorsement**](EndorsementsServiceAPI.md#CreateEndorsement) | **Post** /v1/endorsements | Creates a new endorsement record
[**DeleteEndoresementById**](EndorsementsServiceAPI.md#DeleteEndoresementById) | **Delete** /v1/endorsements/{id} | Finds ands deletes the endorsement by ID
[**FindAllByRepoOwnerOrUsername**](EndorsementsServiceAPI.md#FindAllByRepoOwnerOrUsername) | **Get** /v1/endorsements/repos/{repoOwnerOrUser} | Finds all endorsements by repo org or username and paginates them
[**FindAllEndorsements**](EndorsementsServiceAPI.md#FindAllEndorsements) | **Get** /v1/endorsements | Finds all endorsements and paginates them
[**FindAllEndorsementsByRepo**](EndorsementsServiceAPI.md#FindAllEndorsementsByRepo) | **Get** /v1/endorsements/repos/{owner}/{repo} | Finds all endorsements by repo owner or username and paginates them
[**FindAllUserCreatedEndorsements**](EndorsementsServiceAPI.md#FindAllUserCreatedEndorsements) | **Get** /v1/user/endorsements/created | Finds all endorsements created by the authenticated user and paginates them
[**FindAllUserReceivedEndorsements**](EndorsementsServiceAPI.md#FindAllUserReceivedEndorsements) | **Get** /v1/user/endorsements/received | Finds all endorsements received by the authenticated user and paginates them
[**FindAllUserReceivedEndorsements_0**](EndorsementsServiceAPI.md#FindAllUserReceivedEndorsements_0) | **Get** /v1/endorsements/user/{username}/created | Finds all endorsements received by the user and paginates them
[**FindAllUserReceivedEndorsements_1**](EndorsementsServiceAPI.md#FindAllUserReceivedEndorsements_1) | **Get** /v1/endorsements/user/{username}/received | Finds all endorsements received by the user and paginates them
[**FindEndorsementById**](EndorsementsServiceAPI.md#FindEndorsementById) | **Get** /v1/endorsements/{id} | Retrieves the endorsement based on ID



## CreateEndorsement

> DbEndorsement CreateEndorsement(ctx).XOpenSaucedToken(xOpenSaucedToken).CreateEndorsementDto(createEndorsementDto).Execute()

Creates a new endorsement record

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
    xOpenSaucedToken := "xOpenSaucedToken_example" // string | 
    createEndorsementDto := *openapiclient.NewCreateEndorsementDto(float32(42211), float32(5736810), float32(78133), "https://github.com/open-sauced/insights/pulls", "doc") // CreateEndorsementDto | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EndorsementsServiceAPI.CreateEndorsement(context.Background()).XOpenSaucedToken(xOpenSaucedToken).CreateEndorsementDto(createEndorsementDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EndorsementsServiceAPI.CreateEndorsement``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateEndorsement`: DbEndorsement
    fmt.Fprintf(os.Stdout, "Response from `EndorsementsServiceAPI.CreateEndorsement`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateEndorsementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xOpenSaucedToken** | **string** |  | 
 **createEndorsementDto** | [**CreateEndorsementDto**](CreateEndorsementDto.md) |  | 

### Return type

[**DbEndorsement**](DbEndorsement.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteEndoresementById

> DeleteEndoresementById(ctx, id).Execute()

Finds ands deletes the endorsement by ID

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
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.EndorsementsServiceAPI.DeleteEndoresementById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EndorsementsServiceAPI.DeleteEndoresementById``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteEndoresementByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindAllByRepoOwnerOrUsername

> FindAllUserCreatedEndorsements200Response FindAllByRepoOwnerOrUsername(ctx, repoOwnerOrUser).Page(page).Limit(limit).OrderDirection(orderDirection).Range_(range_).Execute()

Finds all endorsements by repo org or username and paginates them

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
    repoOwnerOrUser := "repoOwnerOrUser_example" // string | 
    page := float32(8.14) // float32 |  (optional) (default to 1)
    limit := float32(8.14) // float32 |  (optional) (default to 10)
    orderDirection := openapiclient.OrderDirectionEnum("ASC") // OrderDirectionEnum |  (optional)
    range_ := float32(8.14) // float32 | Range in days (optional) (default to 30)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EndorsementsServiceAPI.FindAllByRepoOwnerOrUsername(context.Background(), repoOwnerOrUser).Page(page).Limit(limit).OrderDirection(orderDirection).Range_(range_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EndorsementsServiceAPI.FindAllByRepoOwnerOrUsername``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindAllByRepoOwnerOrUsername`: FindAllUserCreatedEndorsements200Response
    fmt.Fprintf(os.Stdout, "Response from `EndorsementsServiceAPI.FindAllByRepoOwnerOrUsername`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repoOwnerOrUser** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindAllByRepoOwnerOrUsernameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **float32** |  | [default to 1]
 **limit** | **float32** |  | [default to 10]
 **orderDirection** | [**OrderDirectionEnum**](OrderDirectionEnum.md) |  | 
 **range_** | **float32** | Range in days | [default to 30]

### Return type

[**FindAllUserCreatedEndorsements200Response**](FindAllUserCreatedEndorsements200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindAllEndorsements

> FindAllUserCreatedEndorsements200Response FindAllEndorsements(ctx).Page(page).Limit(limit).OrderDirection(orderDirection).Range_(range_).Execute()

Finds all endorsements and paginates them

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
    resp, r, err := apiClient.EndorsementsServiceAPI.FindAllEndorsements(context.Background()).Page(page).Limit(limit).OrderDirection(orderDirection).Range_(range_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EndorsementsServiceAPI.FindAllEndorsements``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindAllEndorsements`: FindAllUserCreatedEndorsements200Response
    fmt.Fprintf(os.Stdout, "Response from `EndorsementsServiceAPI.FindAllEndorsements`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFindAllEndorsementsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **float32** |  | [default to 1]
 **limit** | **float32** |  | [default to 10]
 **orderDirection** | [**OrderDirectionEnum**](OrderDirectionEnum.md) |  | 
 **range_** | **float32** | Range in days | [default to 30]

### Return type

[**FindAllUserCreatedEndorsements200Response**](FindAllUserCreatedEndorsements200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindAllEndorsementsByRepo

> FindAllUserCreatedEndorsements200Response FindAllEndorsementsByRepo(ctx, owner, repo).Page(page).Limit(limit).OrderDirection(orderDirection).Range_(range_).Execute()

Finds all endorsements by repo owner or username and paginates them

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
    owner := "owner_example" // string | 
    repo := "repo_example" // string | 
    page := float32(8.14) // float32 |  (optional) (default to 1)
    limit := float32(8.14) // float32 |  (optional) (default to 10)
    orderDirection := openapiclient.OrderDirectionEnum("ASC") // OrderDirectionEnum |  (optional)
    range_ := float32(8.14) // float32 | Range in days (optional) (default to 30)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EndorsementsServiceAPI.FindAllEndorsementsByRepo(context.Background(), owner, repo).Page(page).Limit(limit).OrderDirection(orderDirection).Range_(range_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EndorsementsServiceAPI.FindAllEndorsementsByRepo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindAllEndorsementsByRepo`: FindAllUserCreatedEndorsements200Response
    fmt.Fprintf(os.Stdout, "Response from `EndorsementsServiceAPI.FindAllEndorsementsByRepo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**repo** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindAllEndorsementsByRepoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **float32** |  | [default to 1]
 **limit** | **float32** |  | [default to 10]
 **orderDirection** | [**OrderDirectionEnum**](OrderDirectionEnum.md) |  | 
 **range_** | **float32** | Range in days | [default to 30]

### Return type

[**FindAllUserCreatedEndorsements200Response**](FindAllUserCreatedEndorsements200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindAllUserCreatedEndorsements

> FindAllUserCreatedEndorsements200Response FindAllUserCreatedEndorsements(ctx).Page(page).Limit(limit).OrderDirection(orderDirection).Range_(range_).Execute()

Finds all endorsements created by the authenticated user and paginates them

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
    resp, r, err := apiClient.EndorsementsServiceAPI.FindAllUserCreatedEndorsements(context.Background()).Page(page).Limit(limit).OrderDirection(orderDirection).Range_(range_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EndorsementsServiceAPI.FindAllUserCreatedEndorsements``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindAllUserCreatedEndorsements`: FindAllUserCreatedEndorsements200Response
    fmt.Fprintf(os.Stdout, "Response from `EndorsementsServiceAPI.FindAllUserCreatedEndorsements`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFindAllUserCreatedEndorsementsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **float32** |  | [default to 1]
 **limit** | **float32** |  | [default to 10]
 **orderDirection** | [**OrderDirectionEnum**](OrderDirectionEnum.md) |  | 
 **range_** | **float32** | Range in days | [default to 30]

### Return type

[**FindAllUserCreatedEndorsements200Response**](FindAllUserCreatedEndorsements200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindAllUserReceivedEndorsements

> FindAllUserCreatedEndorsements200Response FindAllUserReceivedEndorsements(ctx).Page(page).Limit(limit).OrderDirection(orderDirection).Range_(range_).Execute()

Finds all endorsements received by the authenticated user and paginates them

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
    resp, r, err := apiClient.EndorsementsServiceAPI.FindAllUserReceivedEndorsements(context.Background()).Page(page).Limit(limit).OrderDirection(orderDirection).Range_(range_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EndorsementsServiceAPI.FindAllUserReceivedEndorsements``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindAllUserReceivedEndorsements`: FindAllUserCreatedEndorsements200Response
    fmt.Fprintf(os.Stdout, "Response from `EndorsementsServiceAPI.FindAllUserReceivedEndorsements`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFindAllUserReceivedEndorsementsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **float32** |  | [default to 1]
 **limit** | **float32** |  | [default to 10]
 **orderDirection** | [**OrderDirectionEnum**](OrderDirectionEnum.md) |  | 
 **range_** | **float32** | Range in days | [default to 30]

### Return type

[**FindAllUserCreatedEndorsements200Response**](FindAllUserCreatedEndorsements200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindAllUserReceivedEndorsements_0

> FindAllUserCreatedEndorsements200Response FindAllUserReceivedEndorsements_0(ctx, username).Page(page).Limit(limit).OrderDirection(orderDirection).Range_(range_).Execute()

Finds all endorsements received by the user and paginates them

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
    username := "username_example" // string | 
    page := float32(8.14) // float32 |  (optional) (default to 1)
    limit := float32(8.14) // float32 |  (optional) (default to 10)
    orderDirection := openapiclient.OrderDirectionEnum("ASC") // OrderDirectionEnum |  (optional)
    range_ := float32(8.14) // float32 | Range in days (optional) (default to 30)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EndorsementsServiceAPI.FindAllUserReceivedEndorsements_0(context.Background(), username).Page(page).Limit(limit).OrderDirection(orderDirection).Range_(range_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EndorsementsServiceAPI.FindAllUserReceivedEndorsements_0``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindAllUserReceivedEndorsements_0`: FindAllUserCreatedEndorsements200Response
    fmt.Fprintf(os.Stdout, "Response from `EndorsementsServiceAPI.FindAllUserReceivedEndorsements_0`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindAllUserReceivedEndorsements_1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **float32** |  | [default to 1]
 **limit** | **float32** |  | [default to 10]
 **orderDirection** | [**OrderDirectionEnum**](OrderDirectionEnum.md) |  | 
 **range_** | **float32** | Range in days | [default to 30]

### Return type

[**FindAllUserCreatedEndorsements200Response**](FindAllUserCreatedEndorsements200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindAllUserReceivedEndorsements_1

> FindAllUserCreatedEndorsements200Response FindAllUserReceivedEndorsements_1(ctx, username).Page(page).Limit(limit).OrderDirection(orderDirection).Range_(range_).Execute()

Finds all endorsements received by the user and paginates them

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
    username := "username_example" // string | 
    page := float32(8.14) // float32 |  (optional) (default to 1)
    limit := float32(8.14) // float32 |  (optional) (default to 10)
    orderDirection := openapiclient.OrderDirectionEnum("ASC") // OrderDirectionEnum |  (optional)
    range_ := float32(8.14) // float32 | Range in days (optional) (default to 30)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EndorsementsServiceAPI.FindAllUserReceivedEndorsements_1(context.Background(), username).Page(page).Limit(limit).OrderDirection(orderDirection).Range_(range_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EndorsementsServiceAPI.FindAllUserReceivedEndorsements_1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindAllUserReceivedEndorsements_1`: FindAllUserCreatedEndorsements200Response
    fmt.Fprintf(os.Stdout, "Response from `EndorsementsServiceAPI.FindAllUserReceivedEndorsements_1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindAllUserReceivedEndorsements_2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **float32** |  | [default to 1]
 **limit** | **float32** |  | [default to 10]
 **orderDirection** | [**OrderDirectionEnum**](OrderDirectionEnum.md) |  | 
 **range_** | **float32** | Range in days | [default to 30]

### Return type

[**FindAllUserCreatedEndorsements200Response**](FindAllUserCreatedEndorsements200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindEndorsementById

> DbEndorsement FindEndorsementById(ctx, id).Execute()

Retrieves the endorsement based on ID



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
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EndorsementsServiceAPI.FindEndorsementById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EndorsementsServiceAPI.FindEndorsementById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindEndorsementById`: DbEndorsement
    fmt.Fprintf(os.Stdout, "Response from `EndorsementsServiceAPI.FindEndorsementById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindEndorsementByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DbEndorsement**](DbEndorsement.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

