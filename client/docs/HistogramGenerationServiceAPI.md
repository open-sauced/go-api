# \HistogramGenerationServiceAPI

All URIs are relative to *https://api.opensauced.pizza*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ForksHistogram**](HistogramGenerationServiceAPI.md#ForksHistogram) | **Get** /v2/histogram/forks | Generates a forks histogram based on 1 day time buckets
[**IssueCommentsHistogram**](HistogramGenerationServiceAPI.md#IssueCommentsHistogram) | **Get** /v2/histogram/issue-comments | Generates a issue comments created histogram based on 1 day time buckets
[**IssuesHistogram**](HistogramGenerationServiceAPI.md#IssuesHistogram) | **Get** /v2/histogram/issues | Generates a issues created histogram based on 1 day time buckets
[**PrReviewsHistogram**](HistogramGenerationServiceAPI.md#PrReviewsHistogram) | **Get** /v2/histogram/pull-request-reviews | Generates a request reviews histogram based on 1 day time buckets
[**PrsHistogram**](HistogramGenerationServiceAPI.md#PrsHistogram) | **Get** /v2/histogram/pull-requests | Generates a pull request created histogram based on 1 day time buckets
[**PushesHistogram**](HistogramGenerationServiceAPI.md#PushesHistogram) | **Get** /v2/histogram/pushes | Generates a pushes histogram based on 1 day time buckets
[**StarsHistogram**](HistogramGenerationServiceAPI.md#StarsHistogram) | **Get** /v2/histogram/stars | Generates a stars histogram based on 1 day time buckets



## ForksHistogram

> []DbPushGitHubEventsHistogram ForksHistogram(ctx).Range_(range_).PrevDaysStartDate(prevDaysStartDate).Width(width).Repo(repo).Contributor(contributor).RepoIds(repoIds).OrderDirection(orderDirection).Filter(filter).Topic(topic).Execute()

Generates a forks histogram based on 1 day time buckets

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
    range_ := int32(56) // int32 | Range in days (optional) (default to 30)
    prevDaysStartDate := int32(56) // int32 | Number of days in the past to start range block (optional) (default to 0)
    width := int32(56) // int32 | Day width of histogram buckets (optional) (default to 1)
    repo := "open-sauced/app" // string | Repo name (optional)
    contributor := "bdougie" // string |  (optional)
    repoIds := "repoIds_example" // string |  (optional)
    orderDirection := openapiclient.OrderDirectionEnum("ASC") // OrderDirectionEnum |  (optional)
    filter := openapiclient.InsightFilterFieldsEnum("recent") // InsightFilterFieldsEnum |  (optional)
    topic := "topic_example" // string |  (optional) (default to "")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HistogramGenerationServiceAPI.ForksHistogram(context.Background()).Range_(range_).PrevDaysStartDate(prevDaysStartDate).Width(width).Repo(repo).Contributor(contributor).RepoIds(repoIds).OrderDirection(orderDirection).Filter(filter).Topic(topic).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HistogramGenerationServiceAPI.ForksHistogram``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ForksHistogram`: []DbPushGitHubEventsHistogram
    fmt.Fprintf(os.Stdout, "Response from `HistogramGenerationServiceAPI.ForksHistogram`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiForksHistogramRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **range_** | **int32** | Range in days | [default to 30]
 **prevDaysStartDate** | **int32** | Number of days in the past to start range block | [default to 0]
 **width** | **int32** | Day width of histogram buckets | [default to 1]
 **repo** | **string** | Repo name | 
 **contributor** | **string** |  | 
 **repoIds** | **string** |  | 
 **orderDirection** | [**OrderDirectionEnum**](OrderDirectionEnum.md) |  | 
 **filter** | [**InsightFilterFieldsEnum**](InsightFilterFieldsEnum.md) |  | 
 **topic** | **string** |  | [default to &quot;&quot;]

### Return type

[**[]DbPushGitHubEventsHistogram**](DbPushGitHubEventsHistogram.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IssueCommentsHistogram

> []DbPushGitHubEventsHistogram IssueCommentsHistogram(ctx).Range_(range_).PrevDaysStartDate(prevDaysStartDate).Width(width).Repo(repo).Contributor(contributor).RepoIds(repoIds).OrderDirection(orderDirection).Filter(filter).Topic(topic).Execute()

Generates a issue comments created histogram based on 1 day time buckets

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
    range_ := int32(56) // int32 | Range in days (optional) (default to 30)
    prevDaysStartDate := int32(56) // int32 | Number of days in the past to start range block (optional) (default to 0)
    width := int32(56) // int32 | Day width of histogram buckets (optional) (default to 1)
    repo := "open-sauced/app" // string | Repo name (optional)
    contributor := "bdougie" // string |  (optional)
    repoIds := "repoIds_example" // string |  (optional)
    orderDirection := openapiclient.OrderDirectionEnum("ASC") // OrderDirectionEnum |  (optional)
    filter := openapiclient.InsightFilterFieldsEnum("recent") // InsightFilterFieldsEnum |  (optional)
    topic := "topic_example" // string |  (optional) (default to "")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HistogramGenerationServiceAPI.IssueCommentsHistogram(context.Background()).Range_(range_).PrevDaysStartDate(prevDaysStartDate).Width(width).Repo(repo).Contributor(contributor).RepoIds(repoIds).OrderDirection(orderDirection).Filter(filter).Topic(topic).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HistogramGenerationServiceAPI.IssueCommentsHistogram``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IssueCommentsHistogram`: []DbPushGitHubEventsHistogram
    fmt.Fprintf(os.Stdout, "Response from `HistogramGenerationServiceAPI.IssueCommentsHistogram`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIssueCommentsHistogramRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **range_** | **int32** | Range in days | [default to 30]
 **prevDaysStartDate** | **int32** | Number of days in the past to start range block | [default to 0]
 **width** | **int32** | Day width of histogram buckets | [default to 1]
 **repo** | **string** | Repo name | 
 **contributor** | **string** |  | 
 **repoIds** | **string** |  | 
 **orderDirection** | [**OrderDirectionEnum**](OrderDirectionEnum.md) |  | 
 **filter** | [**InsightFilterFieldsEnum**](InsightFilterFieldsEnum.md) |  | 
 **topic** | **string** |  | [default to &quot;&quot;]

### Return type

[**[]DbPushGitHubEventsHistogram**](DbPushGitHubEventsHistogram.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IssuesHistogram

> []DbPushGitHubEventsHistogram IssuesHistogram(ctx).Range_(range_).PrevDaysStartDate(prevDaysStartDate).Width(width).Repo(repo).Contributor(contributor).RepoIds(repoIds).OrderDirection(orderDirection).Filter(filter).Topic(topic).Execute()

Generates a issues created histogram based on 1 day time buckets

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
    range_ := int32(56) // int32 | Range in days (optional) (default to 30)
    prevDaysStartDate := int32(56) // int32 | Number of days in the past to start range block (optional) (default to 0)
    width := int32(56) // int32 | Day width of histogram buckets (optional) (default to 1)
    repo := "open-sauced/app" // string | Repo name (optional)
    contributor := "bdougie" // string |  (optional)
    repoIds := "repoIds_example" // string |  (optional)
    orderDirection := openapiclient.OrderDirectionEnum("ASC") // OrderDirectionEnum |  (optional)
    filter := openapiclient.InsightFilterFieldsEnum("recent") // InsightFilterFieldsEnum |  (optional)
    topic := "topic_example" // string |  (optional) (default to "")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HistogramGenerationServiceAPI.IssuesHistogram(context.Background()).Range_(range_).PrevDaysStartDate(prevDaysStartDate).Width(width).Repo(repo).Contributor(contributor).RepoIds(repoIds).OrderDirection(orderDirection).Filter(filter).Topic(topic).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HistogramGenerationServiceAPI.IssuesHistogram``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IssuesHistogram`: []DbPushGitHubEventsHistogram
    fmt.Fprintf(os.Stdout, "Response from `HistogramGenerationServiceAPI.IssuesHistogram`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIssuesHistogramRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **range_** | **int32** | Range in days | [default to 30]
 **prevDaysStartDate** | **int32** | Number of days in the past to start range block | [default to 0]
 **width** | **int32** | Day width of histogram buckets | [default to 1]
 **repo** | **string** | Repo name | 
 **contributor** | **string** |  | 
 **repoIds** | **string** |  | 
 **orderDirection** | [**OrderDirectionEnum**](OrderDirectionEnum.md) |  | 
 **filter** | [**InsightFilterFieldsEnum**](InsightFilterFieldsEnum.md) |  | 
 **topic** | **string** |  | [default to &quot;&quot;]

### Return type

[**[]DbPushGitHubEventsHistogram**](DbPushGitHubEventsHistogram.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PrReviewsHistogram

> []DbPullRequestGitHubEventsHistogram PrReviewsHistogram(ctx).Range_(range_).PrevDaysStartDate(prevDaysStartDate).Width(width).Repo(repo).Contributor(contributor).RepoIds(repoIds).OrderDirection(orderDirection).Filter(filter).Topic(topic).Execute()

Generates a request reviews histogram based on 1 day time buckets

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
    range_ := int32(56) // int32 | Range in days (optional) (default to 30)
    prevDaysStartDate := int32(56) // int32 | Number of days in the past to start range block (optional) (default to 0)
    width := int32(56) // int32 | Day width of histogram buckets (optional) (default to 1)
    repo := "open-sauced/app" // string | Repo name (optional)
    contributor := "bdougie" // string |  (optional)
    repoIds := "repoIds_example" // string |  (optional)
    orderDirection := openapiclient.OrderDirectionEnum("ASC") // OrderDirectionEnum |  (optional)
    filter := openapiclient.InsightFilterFieldsEnum("recent") // InsightFilterFieldsEnum |  (optional)
    topic := "topic_example" // string |  (optional) (default to "")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HistogramGenerationServiceAPI.PrReviewsHistogram(context.Background()).Range_(range_).PrevDaysStartDate(prevDaysStartDate).Width(width).Repo(repo).Contributor(contributor).RepoIds(repoIds).OrderDirection(orderDirection).Filter(filter).Topic(topic).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HistogramGenerationServiceAPI.PrReviewsHistogram``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PrReviewsHistogram`: []DbPullRequestGitHubEventsHistogram
    fmt.Fprintf(os.Stdout, "Response from `HistogramGenerationServiceAPI.PrReviewsHistogram`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPrReviewsHistogramRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **range_** | **int32** | Range in days | [default to 30]
 **prevDaysStartDate** | **int32** | Number of days in the past to start range block | [default to 0]
 **width** | **int32** | Day width of histogram buckets | [default to 1]
 **repo** | **string** | Repo name | 
 **contributor** | **string** |  | 
 **repoIds** | **string** |  | 
 **orderDirection** | [**OrderDirectionEnum**](OrderDirectionEnum.md) |  | 
 **filter** | [**InsightFilterFieldsEnum**](InsightFilterFieldsEnum.md) |  | 
 **topic** | **string** |  | [default to &quot;&quot;]

### Return type

[**[]DbPullRequestGitHubEventsHistogram**](DbPullRequestGitHubEventsHistogram.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PrsHistogram

> []DbPullRequestGitHubEventsHistogram PrsHistogram(ctx).Range_(range_).PrevDaysStartDate(prevDaysStartDate).Width(width).Repo(repo).Contributor(contributor).RepoIds(repoIds).OrderDirection(orderDirection).Filter(filter).Topic(topic).Execute()

Generates a pull request created histogram based on 1 day time buckets

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
    range_ := int32(56) // int32 | Range in days (optional) (default to 30)
    prevDaysStartDate := int32(56) // int32 | Number of days in the past to start range block (optional) (default to 0)
    width := int32(56) // int32 | Day width of histogram buckets (optional) (default to 1)
    repo := "open-sauced/app" // string | Repo name (optional)
    contributor := "bdougie" // string |  (optional)
    repoIds := "repoIds_example" // string |  (optional)
    orderDirection := openapiclient.OrderDirectionEnum("ASC") // OrderDirectionEnum |  (optional)
    filter := openapiclient.InsightFilterFieldsEnum("recent") // InsightFilterFieldsEnum |  (optional)
    topic := "topic_example" // string |  (optional) (default to "")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HistogramGenerationServiceAPI.PrsHistogram(context.Background()).Range_(range_).PrevDaysStartDate(prevDaysStartDate).Width(width).Repo(repo).Contributor(contributor).RepoIds(repoIds).OrderDirection(orderDirection).Filter(filter).Topic(topic).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HistogramGenerationServiceAPI.PrsHistogram``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PrsHistogram`: []DbPullRequestGitHubEventsHistogram
    fmt.Fprintf(os.Stdout, "Response from `HistogramGenerationServiceAPI.PrsHistogram`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPrsHistogramRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **range_** | **int32** | Range in days | [default to 30]
 **prevDaysStartDate** | **int32** | Number of days in the past to start range block | [default to 0]
 **width** | **int32** | Day width of histogram buckets | [default to 1]
 **repo** | **string** | Repo name | 
 **contributor** | **string** |  | 
 **repoIds** | **string** |  | 
 **orderDirection** | [**OrderDirectionEnum**](OrderDirectionEnum.md) |  | 
 **filter** | [**InsightFilterFieldsEnum**](InsightFilterFieldsEnum.md) |  | 
 **topic** | **string** |  | [default to &quot;&quot;]

### Return type

[**[]DbPullRequestGitHubEventsHistogram**](DbPullRequestGitHubEventsHistogram.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PushesHistogram

> []DbPushGitHubEventsHistogram PushesHistogram(ctx).Range_(range_).PrevDaysStartDate(prevDaysStartDate).Width(width).Repo(repo).Contributor(contributor).RepoIds(repoIds).OrderDirection(orderDirection).Filter(filter).Topic(topic).Execute()

Generates a pushes histogram based on 1 day time buckets

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
    range_ := int32(56) // int32 | Range in days (optional) (default to 30)
    prevDaysStartDate := int32(56) // int32 | Number of days in the past to start range block (optional) (default to 0)
    width := int32(56) // int32 | Day width of histogram buckets (optional) (default to 1)
    repo := "open-sauced/app" // string | Repo name (optional)
    contributor := "bdougie" // string |  (optional)
    repoIds := "repoIds_example" // string |  (optional)
    orderDirection := openapiclient.OrderDirectionEnum("ASC") // OrderDirectionEnum |  (optional)
    filter := openapiclient.InsightFilterFieldsEnum("recent") // InsightFilterFieldsEnum |  (optional)
    topic := "topic_example" // string |  (optional) (default to "")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HistogramGenerationServiceAPI.PushesHistogram(context.Background()).Range_(range_).PrevDaysStartDate(prevDaysStartDate).Width(width).Repo(repo).Contributor(contributor).RepoIds(repoIds).OrderDirection(orderDirection).Filter(filter).Topic(topic).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HistogramGenerationServiceAPI.PushesHistogram``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PushesHistogram`: []DbPushGitHubEventsHistogram
    fmt.Fprintf(os.Stdout, "Response from `HistogramGenerationServiceAPI.PushesHistogram`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPushesHistogramRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **range_** | **int32** | Range in days | [default to 30]
 **prevDaysStartDate** | **int32** | Number of days in the past to start range block | [default to 0]
 **width** | **int32** | Day width of histogram buckets | [default to 1]
 **repo** | **string** | Repo name | 
 **contributor** | **string** |  | 
 **repoIds** | **string** |  | 
 **orderDirection** | [**OrderDirectionEnum**](OrderDirectionEnum.md) |  | 
 **filter** | [**InsightFilterFieldsEnum**](InsightFilterFieldsEnum.md) |  | 
 **topic** | **string** |  | [default to &quot;&quot;]

### Return type

[**[]DbPushGitHubEventsHistogram**](DbPushGitHubEventsHistogram.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StarsHistogram

> []DbWatchGitHubEventsHistogram StarsHistogram(ctx).Range_(range_).PrevDaysStartDate(prevDaysStartDate).Width(width).Repo(repo).Contributor(contributor).RepoIds(repoIds).OrderDirection(orderDirection).Filter(filter).Topic(topic).Execute()

Generates a stars histogram based on 1 day time buckets

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
    range_ := int32(56) // int32 | Range in days (optional) (default to 30)
    prevDaysStartDate := int32(56) // int32 | Number of days in the past to start range block (optional) (default to 0)
    width := int32(56) // int32 | Day width of histogram buckets (optional) (default to 1)
    repo := "open-sauced/app" // string | Repo name (optional)
    contributor := "bdougie" // string |  (optional)
    repoIds := "repoIds_example" // string |  (optional)
    orderDirection := openapiclient.OrderDirectionEnum("ASC") // OrderDirectionEnum |  (optional)
    filter := openapiclient.InsightFilterFieldsEnum("recent") // InsightFilterFieldsEnum |  (optional)
    topic := "topic_example" // string |  (optional) (default to "")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HistogramGenerationServiceAPI.StarsHistogram(context.Background()).Range_(range_).PrevDaysStartDate(prevDaysStartDate).Width(width).Repo(repo).Contributor(contributor).RepoIds(repoIds).OrderDirection(orderDirection).Filter(filter).Topic(topic).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HistogramGenerationServiceAPI.StarsHistogram``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StarsHistogram`: []DbWatchGitHubEventsHistogram
    fmt.Fprintf(os.Stdout, "Response from `HistogramGenerationServiceAPI.StarsHistogram`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStarsHistogramRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **range_** | **int32** | Range in days | [default to 30]
 **prevDaysStartDate** | **int32** | Number of days in the past to start range block | [default to 0]
 **width** | **int32** | Day width of histogram buckets | [default to 1]
 **repo** | **string** | Repo name | 
 **contributor** | **string** |  | 
 **repoIds** | **string** |  | 
 **orderDirection** | [**OrderDirectionEnum**](OrderDirectionEnum.md) |  | 
 **filter** | [**InsightFilterFieldsEnum**](InsightFilterFieldsEnum.md) |  | 
 **topic** | **string** |  | [default to &quot;&quot;]

### Return type

[**[]DbWatchGitHubEventsHistogram**](DbWatchGitHubEventsHistogram.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

