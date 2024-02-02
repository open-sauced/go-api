# \UserListsServiceAPI

All URIs are relative to *https://api.opensauced.pizza*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddListForUser**](UserListsServiceAPI.md#AddListForUser) | **Post** /v2/lists | Adds a new list for the authenticated user
[**DeleteListForUser**](UserListsServiceAPI.md#DeleteListForUser) | **Delete** /v2/lists/{id} | Deletes the list for the authenticated user
[**DeleteUserListContributor**](UserListsServiceAPI.md#DeleteUserListContributor) | **Delete** /v2/lists/{id}/contributors/{userListContributorId} | Delete contributor from an individual user list
[**GetContributionsByProject**](UserListsServiceAPI.md#GetContributionsByProject) | **Get** /v2/lists/{id}/stats/contributions-by-project | Gets contributions in a given timeframe
[**GetContributionsByTimeFrame**](UserListsServiceAPI.md#GetContributionsByTimeFrame) | **Get** /v2/lists/{id}/stats/contributions-evolution-by-type | Gets contributions in a given timeframe
[**GetContributorContributionsByProject**](UserListsServiceAPI.md#GetContributorContributionsByProject) | **Get** /v2/lists/{id}/stats/top-project-contributions-by-contributor | Gets top 20 contributor stats in a list by a given project
[**GetContributors**](UserListsServiceAPI.md#GetContributors) | **Get** /v2/lists/contributors | Retrieves paginated contributors
[**GetContributorsByTimeframe**](UserListsServiceAPI.md#GetContributorsByTimeframe) | **Get** /v2/lists/{id}/stats/contributions-evolution-by-contributor-type | Gets contributions by category within timeframe
[**GetFeaturedLists**](UserListsServiceAPI.md#GetFeaturedLists) | **Get** /v2/lists/featured | Gets public featured lists
[**GetListsForUser**](UserListsServiceAPI.md#GetListsForUser) | **Get** /v2/lists | Gets lists for the authenticated user
[**GetMostActiveContributorsV2**](UserListsServiceAPI.md#GetMostActiveContributorsV2) | **Get** /v2/lists/{id}/stats/most-active-contributors | Gets most active contributors for a given list
[**GetTimezones**](UserListsServiceAPI.md#GetTimezones) | **Get** /v2/lists/timezones | Retrieves all users timezones
[**GetUserList**](UserListsServiceAPI.md#GetUserList) | **Get** /v2/lists/{id} | Retrieves an individual user list
[**GetUserListContributorHighlightedRepos**](UserListsServiceAPI.md#GetUserListContributorHighlightedRepos) | **Get** /v2/lists/{id}/contributors/highlights/tagged-repos | Retrieves highlighted repos for contributors for an individual user list
[**GetUserListContributorHighlights**](UserListsServiceAPI.md#GetUserListContributorHighlights) | **Get** /v2/lists/{id}/contributors/highlights | Retrieves highlights for contributors for an individual user list
[**GetUserListContributors**](UserListsServiceAPI.md#GetUserListContributors) | **Get** /v2/lists/{id}/contributors | Retrieves contributors for an individual user list
[**PostUserListContributors**](UserListsServiceAPI.md#PostUserListContributors) | **Post** /v2/lists/{id}/contributors | Add new contributors to an individual user list
[**UpdateListForUser**](UserListsServiceAPI.md#UpdateListForUser) | **Patch** /v2/lists/{id} | Updates the list for the authenticated user



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
    createUserListDto := *openapiclient.NewCreateUserListDto("My List", false, []interface{}{"TODO"}) // CreateUserListDto | 

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


## GetContributionsByProject

> DbContributionsProjects GetContributionsByProject(ctx, id).Range_(range_).Execute()

Gets contributions in a given timeframe

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserListsServiceAPI.GetContributionsByProject(context.Background(), id).Range_(range_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserListsServiceAPI.GetContributionsByProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetContributionsByProject`: DbContributionsProjects
    fmt.Fprintf(os.Stdout, "Response from `UserListsServiceAPI.GetContributionsByProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetContributionsByProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **range_** | **int32** | Range in days | [default to 30]

### Return type

[**DbContributionsProjects**](DbContributionsProjects.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContributionsByTimeFrame

> DbContributionStatTimeframe GetContributionsByTimeFrame(ctx, id).Range_(range_).ContributorType(contributorType).Execute()

Gets contributions in a given timeframe

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
    contributorType := openapiclient.UserListContributorStatsTypeEnum("all") // UserListContributorStatsTypeEnum |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserListsServiceAPI.GetContributionsByTimeFrame(context.Background(), id).Range_(range_).ContributorType(contributorType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserListsServiceAPI.GetContributionsByTimeFrame``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetContributionsByTimeFrame`: DbContributionStatTimeframe
    fmt.Fprintf(os.Stdout, "Response from `UserListsServiceAPI.GetContributionsByTimeFrame`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetContributionsByTimeFrameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **range_** | **int32** | Range in days | [default to 30]
 **contributorType** | [**UserListContributorStatsTypeEnum**](UserListContributorStatsTypeEnum.md) |  | 

### Return type

[**DbContributionStatTimeframe**](DbContributionStatTimeframe.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContributorContributionsByProject

> DbUserListContributorStat GetContributorContributionsByProject(ctx, id).RepoName(repoName).Range_(range_).Execute()

Gets top 20 contributor stats in a list by a given project

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
    repoName := "open-sauced/api" // string | Repo full name
    range_ := int32(56) // int32 | Range in days (optional) (default to 30)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserListsServiceAPI.GetContributorContributionsByProject(context.Background(), id).RepoName(repoName).Range_(range_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserListsServiceAPI.GetContributorContributionsByProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetContributorContributionsByProject`: DbUserListContributorStat
    fmt.Fprintf(os.Stdout, "Response from `UserListsServiceAPI.GetContributorContributionsByProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetContributorContributionsByProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **repoName** | **string** | Repo full name | 
 **range_** | **int32** | Range in days | [default to 30]

### Return type

[**DbUserListContributorStat**](DbUserListContributorStat.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContributors

> GetContributors200Response GetContributors(ctx).Page(page).Limit(limit).Location(location).Contributor(contributor).Timezone(timezone).PrVelocity(prVelocity).Execute()

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
    location := []string{"Inner_example"} // []string |  (optional)
    contributor := "bdougie" // string |  (optional)
    timezone := []string{"Inner_example"} // []string |  (optional)
    prVelocity := int32(2) // int32 | Less than or equal to the average number of days to merge a PR over the last 30 days (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserListsServiceAPI.GetContributors(context.Background()).Page(page).Limit(limit).Location(location).Contributor(contributor).Timezone(timezone).PrVelocity(prVelocity).Execute()
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
 **location** | **[]string** |  | 
 **contributor** | **string** |  | 
 **timezone** | **[]string** |  | 
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


## GetContributorsByTimeframe

> DbContributorCategoryTimeframe GetContributorsByTimeframe(ctx, id).Range_(range_).ContributorType(contributorType).Execute()

Gets contributions by category within timeframe

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
    contributorType := openapiclient.UserListContributorStatsTypeEnum("all") // UserListContributorStatsTypeEnum |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserListsServiceAPI.GetContributorsByTimeframe(context.Background(), id).Range_(range_).ContributorType(contributorType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserListsServiceAPI.GetContributorsByTimeframe``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetContributorsByTimeframe`: DbContributorCategoryTimeframe
    fmt.Fprintf(os.Stdout, "Response from `UserListsServiceAPI.GetContributorsByTimeframe`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetContributorsByTimeframeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **range_** | **int32** | Range in days | [default to 30]
 **contributorType** | [**UserListContributorStatsTypeEnum**](UserListContributorStatsTypeEnum.md) |  | 

### Return type

[**DbContributorCategoryTimeframe**](DbContributorCategoryTimeframe.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFeaturedLists

> DbUserList GetFeaturedLists(ctx).Page(page).Limit(limit).OrderDirection(orderDirection).Range_(range_).PrevDaysStartDate(prevDaysStartDate).Execute()

Gets public featured lists

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
    resp, r, err := apiClient.UserListsServiceAPI.GetFeaturedLists(context.Background()).Page(page).Limit(limit).OrderDirection(orderDirection).Range_(range_).PrevDaysStartDate(prevDaysStartDate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserListsServiceAPI.GetFeaturedLists``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFeaturedLists`: DbUserList
    fmt.Fprintf(os.Stdout, "Response from `UserListsServiceAPI.GetFeaturedLists`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFeaturedListsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** |  | [default to 1]
 **limit** | **int32** |  | [default to 10]
 **orderDirection** | [**OrderDirectionEnum**](OrderDirectionEnum.md) |  | 
 **range_** | **int32** | Range in days | [default to 30]
 **prevDaysStartDate** | **int32** | Number of days in the past to start range block | [default to 0]

### Return type

[**DbUserList**](DbUserList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetListsForUser

> DbUserList GetListsForUser(ctx).Page(page).Limit(limit).OrderDirection(orderDirection).Range_(range_).PrevDaysStartDate(prevDaysStartDate).Execute()

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
    prevDaysStartDate := int32(56) // int32 | Number of days in the past to start range block (optional) (default to 0)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserListsServiceAPI.GetListsForUser(context.Background()).Page(page).Limit(limit).OrderDirection(orderDirection).Range_(range_).PrevDaysStartDate(prevDaysStartDate).Execute()
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
 **prevDaysStartDate** | **int32** | Number of days in the past to start range block | [default to 0]

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


## GetMostActiveContributorsV2

> GetMostActiveContributorsV2200Response GetMostActiveContributorsV2(ctx, id).Page(page).Limit(limit).OrderDirection(orderDirection).Range_(range_).PrevDaysStartDate(prevDaysStartDate).ContributorType(contributorType).OrderBy(orderBy).Execute()

Gets most active contributors for a given list

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
    contributorType := openapiclient.UserListContributorStatsTypeEnum("all") // UserListContributorStatsTypeEnum |  (optional)
    orderBy := openapiclient.UserListContributorStatsOrderEnum("commits") // UserListContributorStatsOrderEnum |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserListsServiceAPI.GetMostActiveContributorsV2(context.Background(), id).Page(page).Limit(limit).OrderDirection(orderDirection).Range_(range_).PrevDaysStartDate(prevDaysStartDate).ContributorType(contributorType).OrderBy(orderBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserListsServiceAPI.GetMostActiveContributorsV2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMostActiveContributorsV2`: GetMostActiveContributorsV2200Response
    fmt.Fprintf(os.Stdout, "Response from `UserListsServiceAPI.GetMostActiveContributorsV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMostActiveContributorsV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** |  | [default to 1]
 **limit** | **int32** |  | [default to 10]
 **orderDirection** | [**OrderDirectionEnum**](OrderDirectionEnum.md) |  | 
 **range_** | **int32** | Range in days | [default to 30]
 **prevDaysStartDate** | **int32** | Number of days in the past to start range block | [default to 0]
 **contributorType** | [**UserListContributorStatsTypeEnum**](UserListContributorStatsTypeEnum.md) |  | 
 **orderBy** | [**UserListContributorStatsOrderEnum**](UserListContributorStatsOrderEnum.md) |  | 

### Return type

[**GetMostActiveContributorsV2200Response**](GetMostActiveContributorsV2200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTimezones

> DbTimezone GetTimezones(ctx).Execute()

Retrieves all users timezones

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserListsServiceAPI.GetTimezones(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserListsServiceAPI.GetTimezones``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTimezones`: DbTimezone
    fmt.Fprintf(os.Stdout, "Response from `UserListsServiceAPI.GetTimezones`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetTimezonesRequest struct via the builder pattern


### Return type

[**DbTimezone**](DbTimezone.md)

### Authorization

No authorization required

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

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserListContributorHighlightedRepos

> GetUserListContributorHighlightedRepos200Response GetUserListContributorHighlightedRepos(ctx, id).Page(page).Limit(limit).OrderDirection(orderDirection).Range_(range_).PrevDaysStartDate(prevDaysStartDate).Execute()

Retrieves highlighted repos for contributors for an individual user list

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
    resp, r, err := apiClient.UserListsServiceAPI.GetUserListContributorHighlightedRepos(context.Background(), id).Page(page).Limit(limit).OrderDirection(orderDirection).Range_(range_).PrevDaysStartDate(prevDaysStartDate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserListsServiceAPI.GetUserListContributorHighlightedRepos``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUserListContributorHighlightedRepos`: GetUserListContributorHighlightedRepos200Response
    fmt.Fprintf(os.Stdout, "Response from `UserListsServiceAPI.GetUserListContributorHighlightedRepos`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserListContributorHighlightedReposRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** |  | [default to 1]
 **limit** | **int32** |  | [default to 10]
 **orderDirection** | [**OrderDirectionEnum**](OrderDirectionEnum.md) |  | 
 **range_** | **int32** | Range in days | [default to 30]
 **prevDaysStartDate** | **int32** | Number of days in the past to start range block | [default to 0]

### Return type

[**GetUserListContributorHighlightedRepos200Response**](GetUserListContributorHighlightedRepos200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserListContributorHighlights

> FindAllHighlightsByUsername200Response GetUserListContributorHighlights(ctx, id).Page(page).Limit(limit).OrderDirection(orderDirection).Range_(range_).PrevDaysStartDate(prevDaysStartDate).Repo(repo).Contributor(contributor).Execute()

Retrieves highlights for contributors for an individual user list

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
    repo := "open-sauced/insights" // string | Highlight Repo Filter (optional)
    contributor := "RitaDee" // string | Highlight User Filter (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserListsServiceAPI.GetUserListContributorHighlights(context.Background(), id).Page(page).Limit(limit).OrderDirection(orderDirection).Range_(range_).PrevDaysStartDate(prevDaysStartDate).Repo(repo).Contributor(contributor).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserListsServiceAPI.GetUserListContributorHighlights``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUserListContributorHighlights`: FindAllHighlightsByUsername200Response
    fmt.Fprintf(os.Stdout, "Response from `UserListsServiceAPI.GetUserListContributorHighlights`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserListContributorHighlightsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** |  | [default to 1]
 **limit** | **int32** |  | [default to 10]
 **orderDirection** | [**OrderDirectionEnum**](OrderDirectionEnum.md) |  | 
 **range_** | **int32** | Range in days | [default to 30]
 **prevDaysStartDate** | **int32** | Number of days in the past to start range block | [default to 0]
 **repo** | **string** | Highlight Repo Filter | 
 **contributor** | **string** | Highlight User Filter | 

### Return type

[**FindAllHighlightsByUsername200Response**](FindAllHighlightsByUsername200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserListContributors

> GetUserListContributors200Response GetUserListContributors(ctx, id).Page(page).Limit(limit).Location(location).Contributor(contributor).Timezone(timezone).PrVelocity(prVelocity).Execute()

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
    location := []string{"Inner_example"} // []string |  (optional)
    contributor := "bdougie" // string |  (optional)
    timezone := []string{"Inner_example"} // []string |  (optional)
    prVelocity := int32(2) // int32 | Less than or equal to the average number of days to merge a PR over the last 30 days (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserListsServiceAPI.GetUserListContributors(context.Background(), id).Page(page).Limit(limit).Location(location).Contributor(contributor).Timezone(timezone).PrVelocity(prVelocity).Execute()
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
 **location** | **[]string** |  | 
 **contributor** | **string** |  | 
 **timezone** | **[]string** |  | 
 **prVelocity** | **int32** | Less than or equal to the average number of days to merge a PR over the last 30 days | 

### Return type

[**GetUserListContributors200Response**](GetUserListContributors200Response.md)

### Authorization

No authorization required

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
    createUserListDto := *openapiclient.NewCreateUserListDto("My List", false, []interface{}{"TODO"}) // CreateUserListDto | 

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

