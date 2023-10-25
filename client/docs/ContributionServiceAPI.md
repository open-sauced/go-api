# \ContributionServiceAPI

All URIs are relative to *https://api.opensauced.pizza*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FindAllByOwnerAndRepo**](ContributionServiceAPI.md#FindAllByOwnerAndRepo) | **Get** /v1/repos/{owner}/{repo}/contributions | Finds a repo by :owner and :repo listing all contributions and paginating them
[**FindAllByOwnerRepoAndContributorLogin**](ContributionServiceAPI.md#FindAllByOwnerRepoAndContributorLogin) | **Get** /v1/repos/{owner}/{repo}/{login}/contributions | Finds a repo by :owner and :repo listing all contributions for a given :login and paginating them
[**FindAllByRepoId**](ContributionServiceAPI.md#FindAllByRepoId) | **Get** /v1/repos/{id}/contributions | Find a repo by :id listing all contributions and paginating them
[**FindAllContributorsByRepoId**](ContributionServiceAPI.md#FindAllContributorsByRepoId) | **Get** /v1/repos/{owner}/{repo}/contributions/contributors | Finds a repo by :owner and :repo listing all contributions by their user login



## FindAllByOwnerAndRepo

> FindAllByRepoId200Response FindAllByOwnerAndRepo(ctx, owner, repo).Page(page).Limit(limit).OrderDirection(orderDirection).Range_(range_).PrevDaysStartDate(prevDaysStartDate).OrderBy(orderBy).Execute()

Finds a repo by :owner and :repo listing all contributions and paginating them

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
    owner := "owner_example" // string | 
    repo := "repo_example" // string | 
    page := int32(56) // int32 |  (optional) (default to 1)
    limit := int32(56) // int32 |  (optional) (default to 10)
    orderDirection := openapiclient.OrderDirectionEnum("ASC") // OrderDirectionEnum |  (optional)
    range_ := int32(56) // int32 | Range in days (optional) (default to 30)
    prevDaysStartDate := int32(56) // int32 | Number of days in the past to start range block (optional) (default to 0)
    orderBy := openapiclient.RepoOrderFieldsEnum("issues") // RepoOrderFieldsEnum |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContributionServiceAPI.FindAllByOwnerAndRepo(context.Background(), owner, repo).Page(page).Limit(limit).OrderDirection(orderDirection).Range_(range_).PrevDaysStartDate(prevDaysStartDate).OrderBy(orderBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContributionServiceAPI.FindAllByOwnerAndRepo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindAllByOwnerAndRepo`: FindAllByRepoId200Response
    fmt.Fprintf(os.Stdout, "Response from `ContributionServiceAPI.FindAllByOwnerAndRepo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**repo** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindAllByOwnerAndRepoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int32** |  | [default to 1]
 **limit** | **int32** |  | [default to 10]
 **orderDirection** | [**OrderDirectionEnum**](OrderDirectionEnum.md) |  | 
 **range_** | **int32** | Range in days | [default to 30]
 **prevDaysStartDate** | **int32** | Number of days in the past to start range block | [default to 0]
 **orderBy** | [**RepoOrderFieldsEnum**](RepoOrderFieldsEnum.md) |  | 

### Return type

[**FindAllByRepoId200Response**](FindAllByRepoId200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindAllByOwnerRepoAndContributorLogin

> DbRepoLoginContributions FindAllByOwnerRepoAndContributorLogin(ctx, owner, repo, login).PrevDaysStartDate(prevDaysStartDate).Range_(range_).Execute()

Finds a repo by :owner and :repo listing all contributions for a given :login and paginating them

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
    owner := "owner_example" // string | 
    repo := "repo_example" // string | 
    login := "login_example" // string | 
    prevDaysStartDate := int32(56) // int32 | Previous number of days to go back to start date range (optional)
    range_ := int32(56) // int32 | Range in days (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContributionServiceAPI.FindAllByOwnerRepoAndContributorLogin(context.Background(), owner, repo, login).PrevDaysStartDate(prevDaysStartDate).Range_(range_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContributionServiceAPI.FindAllByOwnerRepoAndContributorLogin``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindAllByOwnerRepoAndContributorLogin`: DbRepoLoginContributions
    fmt.Fprintf(os.Stdout, "Response from `ContributionServiceAPI.FindAllByOwnerRepoAndContributorLogin`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**repo** | **string** |  | 
**login** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindAllByOwnerRepoAndContributorLoginRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **prevDaysStartDate** | **int32** | Previous number of days to go back to start date range | 
 **range_** | **int32** | Range in days | 

### Return type

[**DbRepoLoginContributions**](DbRepoLoginContributions.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindAllByRepoId

> FindAllByRepoId200Response FindAllByRepoId(ctx, id).Page(page).Limit(limit).OrderDirection(orderDirection).Range_(range_).PrevDaysStartDate(prevDaysStartDate).OrderBy(orderBy).Execute()

Find a repo by :id listing all contributions and paginating them

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
    id := int32(56) // int32 | 
    page := int32(56) // int32 |  (optional) (default to 1)
    limit := int32(56) // int32 |  (optional) (default to 10)
    orderDirection := openapiclient.OrderDirectionEnum("ASC") // OrderDirectionEnum |  (optional)
    range_ := int32(56) // int32 | Range in days (optional) (default to 30)
    prevDaysStartDate := int32(56) // int32 | Number of days in the past to start range block (optional) (default to 0)
    orderBy := openapiclient.RepoOrderFieldsEnum("issues") // RepoOrderFieldsEnum |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContributionServiceAPI.FindAllByRepoId(context.Background(), id).Page(page).Limit(limit).OrderDirection(orderDirection).Range_(range_).PrevDaysStartDate(prevDaysStartDate).OrderBy(orderBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContributionServiceAPI.FindAllByRepoId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindAllByRepoId`: FindAllByRepoId200Response
    fmt.Fprintf(os.Stdout, "Response from `ContributionServiceAPI.FindAllByRepoId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindAllByRepoIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** |  | [default to 1]
 **limit** | **int32** |  | [default to 10]
 **orderDirection** | [**OrderDirectionEnum**](OrderDirectionEnum.md) |  | 
 **range_** | **int32** | Range in days | [default to 30]
 **prevDaysStartDate** | **int32** | Number of days in the past to start range block | [default to 0]
 **orderBy** | [**RepoOrderFieldsEnum**](RepoOrderFieldsEnum.md) |  | 

### Return type

[**FindAllByRepoId200Response**](FindAllByRepoId200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindAllContributorsByRepoId

> []DbRepoLoginContributions FindAllContributorsByRepoId(ctx, owner, repo).PrevDaysStartDate(prevDaysStartDate).Range_(range_).Execute()

Finds a repo by :owner and :repo listing all contributions by their user login

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
    owner := "owner_example" // string | 
    repo := "repo_example" // string | 
    prevDaysStartDate := int32(56) // int32 | Previous number of days to go back to start date range (optional)
    range_ := int32(56) // int32 | Range in days (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContributionServiceAPI.FindAllContributorsByRepoId(context.Background(), owner, repo).PrevDaysStartDate(prevDaysStartDate).Range_(range_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContributionServiceAPI.FindAllContributorsByRepoId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindAllContributorsByRepoId`: []DbRepoLoginContributions
    fmt.Fprintf(os.Stdout, "Response from `ContributionServiceAPI.FindAllContributorsByRepoId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**repo** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindAllContributorsByRepoIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **prevDaysStartDate** | **int32** | Previous number of days to go back to start date range | 
 **range_** | **int32** | Range in days | 

### Return type

[**[]DbRepoLoginContributions**](DbRepoLoginContributions.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

