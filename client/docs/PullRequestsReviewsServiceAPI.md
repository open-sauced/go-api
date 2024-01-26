# \PullRequestsReviewsServiceAPI

All URIs are relative to *https://api.opensauced.pizza*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SearchAllPullRequestsReviews**](PullRequestsReviewsServiceAPI.md#SearchAllPullRequestsReviews) | **Get** /v2/prs/reviews/search | Searches pull requests reviews using filters and paginates them



## SearchAllPullRequestsReviews

> SearchAllPullRequestsReviews200Response SearchAllPullRequestsReviews(ctx).Page(page).Limit(limit).OrderDirection(orderDirection).Range_(range_).PrevDaysStartDate(prevDaysStartDate).OrderBy(orderBy).Filter(filter).Topic(topic).Repo(repo).RepoIds(repoIds).Status(status).Contributor(contributor).ListId(listId).DistinctAuthors(distinctAuthors).Execute()

Searches pull requests reviews using filters and paginates them

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
    orderBy := openapiclient.PullRequestOrderFieldsEnum("created_at") // PullRequestOrderFieldsEnum |  (optional)
    filter := openapiclient.InsightFilterFieldsEnum("recent") // InsightFilterFieldsEnum |  (optional)
    topic := "javascript" // string |  (optional)
    repo := "open-sauced/insights" // string |  (optional)
    repoIds := "repoIds_example" // string |  (optional)
    status := openapiclient.PullRequestStatusEnum("open") // PullRequestStatusEnum |  (optional)
    contributor := "bdougie" // string |  (optional)
    listId := "uuid-v4" // string |  (optional)
    distinctAuthors := "true" // string |  (optional) (default to "false")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PullRequestsReviewsServiceAPI.SearchAllPullRequestsReviews(context.Background()).Page(page).Limit(limit).OrderDirection(orderDirection).Range_(range_).PrevDaysStartDate(prevDaysStartDate).OrderBy(orderBy).Filter(filter).Topic(topic).Repo(repo).RepoIds(repoIds).Status(status).Contributor(contributor).ListId(listId).DistinctAuthors(distinctAuthors).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PullRequestsReviewsServiceAPI.SearchAllPullRequestsReviews``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchAllPullRequestsReviews`: SearchAllPullRequestsReviews200Response
    fmt.Fprintf(os.Stdout, "Response from `PullRequestsReviewsServiceAPI.SearchAllPullRequestsReviews`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchAllPullRequestsReviewsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** |  | [default to 1]
 **limit** | **int32** |  | [default to 10]
 **orderDirection** | [**OrderDirectionEnum**](OrderDirectionEnum.md) |  | 
 **range_** | **int32** | Range in days | [default to 30]
 **prevDaysStartDate** | **int32** | Number of days in the past to start range block | [default to 0]
 **orderBy** | [**PullRequestOrderFieldsEnum**](PullRequestOrderFieldsEnum.md) |  | 
 **filter** | [**InsightFilterFieldsEnum**](InsightFilterFieldsEnum.md) |  | 
 **topic** | **string** |  | 
 **repo** | **string** |  | 
 **repoIds** | **string** |  | 
 **status** | [**PullRequestStatusEnum**](PullRequestStatusEnum.md) |  | 
 **contributor** | **string** |  | 
 **listId** | **string** |  | 
 **distinctAuthors** | **string** |  | [default to &quot;false&quot;]

### Return type

[**SearchAllPullRequestsReviews200Response**](SearchAllPullRequestsReviews200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

