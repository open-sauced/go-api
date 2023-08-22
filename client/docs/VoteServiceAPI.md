# \VoteServiceAPI

All URIs are relative to *https://api.opensauced.pizza*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DownVoteOneById**](VoteServiceAPI.md#DownVoteOneById) | **Delete** /v1/repos/{id}/vote | Finds a repo by :id and removes existing vote
[**DownVoteOneByOwnerAndRepo**](VoteServiceAPI.md#DownVoteOneByOwnerAndRepo) | **Delete** /v1/repos/{owner}/{repo}/vote | Finds a repo by :owner and :repo and removes existing vote
[**FindAllUserVoted**](VoteServiceAPI.md#FindAllUserVoted) | **Get** /v1/repos/listUserVoted | Finds all repos voted by authenticated user and paginates them
[**FindOneByRepoId**](VoteServiceAPI.md#FindOneByRepoId) | **Get** /v1/repos/{repoId}/vote | Finds a repo by :repoId and returns if authenticated user has voted for it
[**VoteOneById**](VoteServiceAPI.md#VoteOneById) | **Put** /v1/repos/{id}/vote | Finds a repo by :id and adds a vote
[**VoteOneByOwnerAndRepo**](VoteServiceAPI.md#VoteOneByOwnerAndRepo) | **Put** /v1/repos/{owner}/{repo}/vote | Finds a repo by :owner and :repo and adds a vote



## DownVoteOneById

> DbRepoToUserVotes DownVoteOneById(ctx, id).Execute()

Finds a repo by :id and removes existing vote

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
    id := float32(8.14) // float32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VoteServiceAPI.DownVoteOneById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VoteServiceAPI.DownVoteOneById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DownVoteOneById`: DbRepoToUserVotes
    fmt.Fprintf(os.Stdout, "Response from `VoteServiceAPI.DownVoteOneById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownVoteOneByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DbRepoToUserVotes**](DbRepoToUserVotes.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DownVoteOneByOwnerAndRepo

> DbRepoToUserVotes DownVoteOneByOwnerAndRepo(ctx, owner, repo).Execute()

Finds a repo by :owner and :repo and removes existing vote

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VoteServiceAPI.DownVoteOneByOwnerAndRepo(context.Background(), owner, repo).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VoteServiceAPI.DownVoteOneByOwnerAndRepo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DownVoteOneByOwnerAndRepo`: DbRepoToUserVotes
    fmt.Fprintf(os.Stdout, "Response from `VoteServiceAPI.DownVoteOneByOwnerAndRepo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**repo** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownVoteOneByOwnerAndRepoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**DbRepoToUserVotes**](DbRepoToUserVotes.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindAllUserVoted

> FindAllTopReposByUsername200Response FindAllUserVoted(ctx).Page(page).Limit(limit).OrderDirection(orderDirection).Range_(range_).OrderBy(orderBy).Execute()

Finds all repos voted by authenticated user and paginates them

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
    orderBy := openapiclient.RepoOrderFieldsEnum("issues") // RepoOrderFieldsEnum |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VoteServiceAPI.FindAllUserVoted(context.Background()).Page(page).Limit(limit).OrderDirection(orderDirection).Range_(range_).OrderBy(orderBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VoteServiceAPI.FindAllUserVoted``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindAllUserVoted`: FindAllTopReposByUsername200Response
    fmt.Fprintf(os.Stdout, "Response from `VoteServiceAPI.FindAllUserVoted`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFindAllUserVotedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **float32** |  | [default to 1]
 **limit** | **float32** |  | [default to 10]
 **orderDirection** | [**OrderDirectionEnum**](OrderDirectionEnum.md) |  | 
 **range_** | **float32** | Range in days | [default to 30]
 **orderBy** | [**RepoOrderFieldsEnum**](RepoOrderFieldsEnum.md) |  | 

### Return type

[**FindAllTopReposByUsername200Response**](FindAllTopReposByUsername200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindOneByRepoId

> VotedRepoDto FindOneByRepoId(ctx, repoId).Execute()

Finds a repo by :repoId and returns if authenticated user has voted for it

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
    repoId := float32(8.14) // float32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VoteServiceAPI.FindOneByRepoId(context.Background(), repoId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VoteServiceAPI.FindOneByRepoId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindOneByRepoId`: VotedRepoDto
    fmt.Fprintf(os.Stdout, "Response from `VoteServiceAPI.FindOneByRepoId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repoId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindOneByRepoIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**VotedRepoDto**](VotedRepoDto.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VoteOneById

> DbRepoToUserVotes VoteOneById(ctx, id).Execute()

Finds a repo by :id and adds a vote

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
    id := float32(8.14) // float32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VoteServiceAPI.VoteOneById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VoteServiceAPI.VoteOneById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `VoteOneById`: DbRepoToUserVotes
    fmt.Fprintf(os.Stdout, "Response from `VoteServiceAPI.VoteOneById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiVoteOneByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DbRepoToUserVotes**](DbRepoToUserVotes.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VoteOneByOwnerAndRepo

> DbRepoToUserVotes VoteOneByOwnerAndRepo(ctx, owner, repo).Execute()

Finds a repo by :owner and :repo and adds a vote

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VoteServiceAPI.VoteOneByOwnerAndRepo(context.Background(), owner, repo).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VoteServiceAPI.VoteOneByOwnerAndRepo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `VoteOneByOwnerAndRepo`: DbRepoToUserVotes
    fmt.Fprintf(os.Stdout, "Response from `VoteServiceAPI.VoteOneByOwnerAndRepo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**repo** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiVoteOneByOwnerAndRepoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**DbRepoToUserVotes**](DbRepoToUserVotes.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

