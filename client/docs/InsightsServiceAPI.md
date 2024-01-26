# \InsightsServiceAPI

All URIs are relative to *https://api.opensauced.pizza*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddInsightForUser**](InsightsServiceAPI.md#AddInsightForUser) | **Post** /v2/user/insights | Adds a new insight page for the authenticated user
[**AddMemberForInsight**](InsightsServiceAPI.md#AddMemberForInsight) | **Post** /v2/user/insights/{id}/members | Adds a new member for the insight
[**FindAllInsightMembers**](InsightsServiceAPI.md#FindAllInsightMembers) | **Get** /v2/user/insights/{id}/members | Listing all members for an insight and paginate them
[**FindAllInsightsByUserId**](InsightsServiceAPI.md#FindAllInsightsByUserId) | **Get** /v2/user/insights | Listing all insights for a user and paginate them
[**FindFeaturedInsights**](InsightsServiceAPI.md#FindFeaturedInsights) | **Get** /v2/insights/featured | Finds featured insights
[**FindInsightPageById**](InsightsServiceAPI.md#FindInsightPageById) | **Get** /v2/insights/{id} | Finds a insight page by :id
[**FindInsightReposById**](InsightsServiceAPI.md#FindInsightReposById) | **Get** /v2/insights/{id}/repos | Finds a insight page repositories by :id
[**RemoveInsightForUser**](InsightsServiceAPI.md#RemoveInsightForUser) | **Delete** /v2/insights/{id} | Removes an insight page for the authenticated user
[**RemoveInsightMemberById**](InsightsServiceAPI.md#RemoveInsightMemberById) | **Delete** /v2/user/insights/{id}/members/{memberId} | Removes a member from an insight
[**UpdateInsightForUser**](InsightsServiceAPI.md#UpdateInsightForUser) | **Patch** /v2/user/insights/{id} | Updates an insight page for the authenticated user
[**UpdateInsightMember**](InsightsServiceAPI.md#UpdateInsightMember) | **Patch** /v2/user/insights/{id}/members/{memberId} | Updates an insight member information



## AddInsightForUser

> DbInsight AddInsightForUser(ctx).CreateInsightDto(createInsightDto).Execute()

Adds a new insight page for the authenticated user

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
    createInsightDto := *openapiclient.NewCreateInsightDto("My Team", false, []openapiclient.RepoInfo{*openapiclient.NewRepoInfo(int32(234343), "open-sauced/insights")}) // CreateInsightDto | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InsightsServiceAPI.AddInsightForUser(context.Background()).CreateInsightDto(createInsightDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InsightsServiceAPI.AddInsightForUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddInsightForUser`: DbInsight
    fmt.Fprintf(os.Stdout, "Response from `InsightsServiceAPI.AddInsightForUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddInsightForUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createInsightDto** | [**CreateInsightDto**](CreateInsightDto.md) |  | 

### Return type

[**DbInsight**](DbInsight.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddMemberForInsight

> DbInsightMember AddMemberForInsight(ctx, id).CreateInsightMemberDto(createInsightMemberDto).Execute()

Adds a new member for the insight

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
    createInsightMemberDto := *openapiclient.NewCreateInsightMemberDto("hello@opensauced.pizza") // CreateInsightMemberDto | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InsightsServiceAPI.AddMemberForInsight(context.Background(), id).CreateInsightMemberDto(createInsightMemberDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InsightsServiceAPI.AddMemberForInsight``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddMemberForInsight`: DbInsightMember
    fmt.Fprintf(os.Stdout, "Response from `InsightsServiceAPI.AddMemberForInsight`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddMemberForInsightRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createInsightMemberDto** | [**CreateInsightMemberDto**](CreateInsightMemberDto.md) |  | 

### Return type

[**DbInsightMember**](DbInsightMember.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindAllInsightMembers

> FindAllInsightMembers200Response FindAllInsightMembers(ctx, id).Page(page).Limit(limit).OrderDirection(orderDirection).Range_(range_).PrevDaysStartDate(prevDaysStartDate).Execute()

Listing all members for an insight and paginate them

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InsightsServiceAPI.FindAllInsightMembers(context.Background(), id).Page(page).Limit(limit).OrderDirection(orderDirection).Range_(range_).PrevDaysStartDate(prevDaysStartDate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InsightsServiceAPI.FindAllInsightMembers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindAllInsightMembers`: FindAllInsightMembers200Response
    fmt.Fprintf(os.Stdout, "Response from `InsightsServiceAPI.FindAllInsightMembers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindAllInsightMembersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** |  | [default to 1]
 **limit** | **int32** |  | [default to 10]
 **orderDirection** | [**OrderDirectionEnum**](OrderDirectionEnum.md) |  | 
 **range_** | **int32** | Range in days | [default to 30]
 **prevDaysStartDate** | **int32** | Number of days in the past to start range block | [default to 0]

### Return type

[**FindAllInsightMembers200Response**](FindAllInsightMembers200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindAllInsightsByUserId

> FindAllInsightsByUserId200Response FindAllInsightsByUserId(ctx).Page(page).Limit(limit).OrderDirection(orderDirection).Range_(range_).PrevDaysStartDate(prevDaysStartDate).Execute()

Listing all insights for a user and paginate them

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
    resp, r, err := apiClient.InsightsServiceAPI.FindAllInsightsByUserId(context.Background()).Page(page).Limit(limit).OrderDirection(orderDirection).Range_(range_).PrevDaysStartDate(prevDaysStartDate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InsightsServiceAPI.FindAllInsightsByUserId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindAllInsightsByUserId`: FindAllInsightsByUserId200Response
    fmt.Fprintf(os.Stdout, "Response from `InsightsServiceAPI.FindAllInsightsByUserId`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFindAllInsightsByUserIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** |  | [default to 1]
 **limit** | **int32** |  | [default to 10]
 **orderDirection** | [**OrderDirectionEnum**](OrderDirectionEnum.md) |  | 
 **range_** | **int32** | Range in days | [default to 30]
 **prevDaysStartDate** | **int32** | Number of days in the past to start range block | [default to 0]

### Return type

[**FindAllInsightsByUserId200Response**](FindAllInsightsByUserId200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindFeaturedInsights

> DbInsight FindFeaturedInsights(ctx).Page(page).Limit(limit).OrderDirection(orderDirection).Range_(range_).PrevDaysStartDate(prevDaysStartDate).Execute()

Finds featured insights

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
    resp, r, err := apiClient.InsightsServiceAPI.FindFeaturedInsights(context.Background()).Page(page).Limit(limit).OrderDirection(orderDirection).Range_(range_).PrevDaysStartDate(prevDaysStartDate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InsightsServiceAPI.FindFeaturedInsights``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindFeaturedInsights`: DbInsight
    fmt.Fprintf(os.Stdout, "Response from `InsightsServiceAPI.FindFeaturedInsights`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFindFeaturedInsightsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** |  | [default to 1]
 **limit** | **int32** |  | [default to 10]
 **orderDirection** | [**OrderDirectionEnum**](OrderDirectionEnum.md) |  | 
 **range_** | **int32** | Range in days | [default to 30]
 **prevDaysStartDate** | **int32** | Number of days in the past to start range block | [default to 0]

### Return type

[**DbInsight**](DbInsight.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindInsightPageById

> DbInsight FindInsightPageById(ctx, id).Include(include).Execute()

Finds a insight page by :id

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
    include := "all" // string | Include all data (optional) (default to "all")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InsightsServiceAPI.FindInsightPageById(context.Background(), id).Include(include).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InsightsServiceAPI.FindInsightPageById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindInsightPageById`: DbInsight
    fmt.Fprintf(os.Stdout, "Response from `InsightsServiceAPI.FindInsightPageById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindInsightPageByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **include** | **string** | Include all data | [default to &quot;all&quot;]

### Return type

[**DbInsight**](DbInsight.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindInsightReposById

> DbInsightRepo FindInsightReposById(ctx, id).Execute()

Finds a insight page repositories by :id

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InsightsServiceAPI.FindInsightReposById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InsightsServiceAPI.FindInsightReposById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindInsightReposById`: DbInsightRepo
    fmt.Fprintf(os.Stdout, "Response from `InsightsServiceAPI.FindInsightReposById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindInsightReposByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DbInsightRepo**](DbInsightRepo.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveInsightForUser

> DbInsight RemoveInsightForUser(ctx, id).UpdateInsightDto(updateInsightDto).Execute()

Removes an insight page for the authenticated user

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
    updateInsightDto := *openapiclient.NewUpdateInsightDto("My Team", false, []openapiclient.RepoInfo{*openapiclient.NewRepoInfo(int32(234343), "open-sauced/insights")}) // UpdateInsightDto | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InsightsServiceAPI.RemoveInsightForUser(context.Background(), id).UpdateInsightDto(updateInsightDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InsightsServiceAPI.RemoveInsightForUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemoveInsightForUser`: DbInsight
    fmt.Fprintf(os.Stdout, "Response from `InsightsServiceAPI.RemoveInsightForUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveInsightForUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateInsightDto** | [**UpdateInsightDto**](UpdateInsightDto.md) |  | 

### Return type

[**DbInsight**](DbInsight.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveInsightMemberById

> DbInsight RemoveInsightMemberById(ctx, id, memberId).Execute()

Removes a member from an insight

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
    memberId := "memberId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InsightsServiceAPI.RemoveInsightMemberById(context.Background(), id, memberId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InsightsServiceAPI.RemoveInsightMemberById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemoveInsightMemberById`: DbInsight
    fmt.Fprintf(os.Stdout, "Response from `InsightsServiceAPI.RemoveInsightMemberById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 
**memberId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveInsightMemberByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**DbInsight**](DbInsight.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateInsightForUser

> DbInsight UpdateInsightForUser(ctx, id).UpdateInsightDto(updateInsightDto).Execute()

Updates an insight page for the authenticated user

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
    updateInsightDto := *openapiclient.NewUpdateInsightDto("My Team", false, []openapiclient.RepoInfo{*openapiclient.NewRepoInfo(int32(234343), "open-sauced/insights")}) // UpdateInsightDto | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InsightsServiceAPI.UpdateInsightForUser(context.Background(), id).UpdateInsightDto(updateInsightDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InsightsServiceAPI.UpdateInsightForUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateInsightForUser`: DbInsight
    fmt.Fprintf(os.Stdout, "Response from `InsightsServiceAPI.UpdateInsightForUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateInsightForUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateInsightDto** | [**UpdateInsightDto**](UpdateInsightDto.md) |  | 

### Return type

[**DbInsight**](DbInsight.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateInsightMember

> DbInsight UpdateInsightMember(ctx, id, memberId).UpdateInsightMemberDto(updateInsightMemberDto).Execute()

Updates an insight member information

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
    memberId := "memberId_example" // string | 
    updateInsightMemberDto := *openapiclient.NewUpdateInsightMemberDto("view") // UpdateInsightMemberDto | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InsightsServiceAPI.UpdateInsightMember(context.Background(), id, memberId).UpdateInsightMemberDto(updateInsightMemberDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InsightsServiceAPI.UpdateInsightMember``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateInsightMember`: DbInsight
    fmt.Fprintf(os.Stdout, "Response from `InsightsServiceAPI.UpdateInsightMember`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 
**memberId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateInsightMemberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateInsightMemberDto** | [**UpdateInsightMemberDto**](UpdateInsightMemberDto.md) |  | 

### Return type

[**DbInsight**](DbInsight.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

