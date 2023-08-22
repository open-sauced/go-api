# \PullRequestsServiceAPI

All URIs are relative to *http://localhost:3001*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GenerateCodeExplanation**](PullRequestsServiceAPI.md#GenerateCodeExplanation) | **Post** /v1/prs/explanation/generate | Generates an explanation for the provided code
[**GenerateCodeRefactor**](PullRequestsServiceAPI.md#GenerateCodeRefactor) | **Post** /v1/prs/suggestion/generate | Generates a refactor suggestion based on the provided code
[**GenerateCodeTest**](PullRequestsServiceAPI.md#GenerateCodeTest) | **Post** /v1/prs/test/generate | Generates a test for the provided code
[**GeneratePRDescription**](PullRequestsServiceAPI.md#GeneratePRDescription) | **Post** /v1/prs/description/generate | Generates a PR description based on the provided information
[**GetPullRequestInsights**](PullRequestsServiceAPI.md#GetPullRequestInsights) | **Get** /v1/prs/insights | Find pull request insights over the last 2 months
[**ListAllPullRequests**](PullRequestsServiceAPI.md#ListAllPullRequests) | **Get** /v1/prs/list | Finds all pull requests and paginates them
[**SearchAllPullRequests**](PullRequestsServiceAPI.md#SearchAllPullRequests) | **Get** /v1/prs/search | Searches pull requests using filters and paginates them



## GenerateCodeExplanation

> GenerateCodeExplanation(ctx).GenerateCodeExplanationDto(generateCodeExplanationDto).Execute()

Generates an explanation for the provided code

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
    generateCodeExplanationDto := *openapiclient.NewGenerateCodeExplanationDto(float32(250), float32(7), "english", "Code_example") // GenerateCodeExplanationDto | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PullRequestsServiceAPI.GenerateCodeExplanation(context.Background()).GenerateCodeExplanationDto(generateCodeExplanationDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PullRequestsServiceAPI.GenerateCodeExplanation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGenerateCodeExplanationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **generateCodeExplanationDto** | [**GenerateCodeExplanationDto**](GenerateCodeExplanationDto.md) |  | 

### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GenerateCodeRefactor

> GenerateCodeRefactor(ctx).GenerateCodeRefactorSuggestionDto(generateCodeRefactorSuggestionDto).Execute()

Generates a refactor suggestion based on the provided code

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
    generateCodeRefactorSuggestionDto := *openapiclient.NewGenerateCodeRefactorSuggestionDto(float32(250), float32(7), "english", "Code_example") // GenerateCodeRefactorSuggestionDto | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PullRequestsServiceAPI.GenerateCodeRefactor(context.Background()).GenerateCodeRefactorSuggestionDto(generateCodeRefactorSuggestionDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PullRequestsServiceAPI.GenerateCodeRefactor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGenerateCodeRefactorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **generateCodeRefactorSuggestionDto** | [**GenerateCodeRefactorSuggestionDto**](GenerateCodeRefactorSuggestionDto.md) |  | 

### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GenerateCodeTest

> GenerateCodeTest(ctx).GenerateCodeTestSuggestionDto(generateCodeTestSuggestionDto).Execute()

Generates a test for the provided code

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
    generateCodeTestSuggestionDto := *openapiclient.NewGenerateCodeTestSuggestionDto(float32(250), float32(7), "Code_example") // GenerateCodeTestSuggestionDto | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PullRequestsServiceAPI.GenerateCodeTest(context.Background()).GenerateCodeTestSuggestionDto(generateCodeTestSuggestionDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PullRequestsServiceAPI.GenerateCodeTest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGenerateCodeTestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **generateCodeTestSuggestionDto** | [**GenerateCodeTestSuggestionDto**](GenerateCodeTestSuggestionDto.md) |  | 

### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GeneratePRDescription

> GeneratePRDescription(ctx).GeneratePullRequestDescriptionDto(generatePullRequestDescriptionDto).Execute()

Generates a PR description based on the provided information

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
    generatePullRequestDescriptionDto := *openapiclient.NewGeneratePullRequestDescriptionDto(float32(250), float32(7), "formal", "english") // GeneratePullRequestDescriptionDto | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PullRequestsServiceAPI.GeneratePRDescription(context.Background()).GeneratePullRequestDescriptionDto(generatePullRequestDescriptionDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PullRequestsServiceAPI.GeneratePRDescription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGeneratePRDescriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **generatePullRequestDescriptionDto** | [**GeneratePullRequestDescriptionDto**](GeneratePullRequestDescriptionDto.md) |  | 

### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPullRequestInsights

> []DbPRInsight GetPullRequestInsights(ctx).Page(page).Limit(limit).OrderDirection(orderDirection).Range_(range_).Filter(filter).Topic(topic).Repo(repo).RepoIds(repoIds).Execute()

Find pull request insights over the last 2 months

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
    filter := openapiclient.InsightFilterFieldsEnum("recent") // InsightFilterFieldsEnum |  (optional)
    topic := "topic_example" // string |  (optional) (default to "")
    repo := "open-sauced/insights" // string |  (optional)
    repoIds := "repoIds_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PullRequestsServiceAPI.GetPullRequestInsights(context.Background()).Page(page).Limit(limit).OrderDirection(orderDirection).Range_(range_).Filter(filter).Topic(topic).Repo(repo).RepoIds(repoIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PullRequestsServiceAPI.GetPullRequestInsights``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPullRequestInsights`: []DbPRInsight
    fmt.Fprintf(os.Stdout, "Response from `PullRequestsServiceAPI.GetPullRequestInsights`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPullRequestInsightsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **float32** |  | [default to 1]
 **limit** | **float32** |  | [default to 10]
 **orderDirection** | [**OrderDirectionEnum**](OrderDirectionEnum.md) |  | 
 **range_** | **float32** | Range in days | [default to 30]
 **filter** | [**InsightFilterFieldsEnum**](InsightFilterFieldsEnum.md) |  | 
 **topic** | **string** |  | [default to &quot;&quot;]
 **repo** | **string** |  | 
 **repoIds** | **string** |  | 

### Return type

[**[]DbPRInsight**](DbPRInsight.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAllPullRequests

> FindContributorPullRequests200Response ListAllPullRequests(ctx).Page(page).Limit(limit).OrderDirection(orderDirection).Range_(range_).Execute()

Finds all pull requests and paginates them

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
    resp, r, err := apiClient.PullRequestsServiceAPI.ListAllPullRequests(context.Background()).Page(page).Limit(limit).OrderDirection(orderDirection).Range_(range_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PullRequestsServiceAPI.ListAllPullRequests``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAllPullRequests`: FindContributorPullRequests200Response
    fmt.Fprintf(os.Stdout, "Response from `PullRequestsServiceAPI.ListAllPullRequests`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAllPullRequestsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **float32** |  | [default to 1]
 **limit** | **float32** |  | [default to 10]
 **orderDirection** | [**OrderDirectionEnum**](OrderDirectionEnum.md) |  | 
 **range_** | **float32** | Range in days | [default to 30]

### Return type

[**FindContributorPullRequests200Response**](FindContributorPullRequests200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchAllPullRequests

> FindContributorPullRequests200Response SearchAllPullRequests(ctx).Page(page).Limit(limit).OrderDirection(orderDirection).Range_(range_).OrderBy(orderBy).Filter(filter).Topic(topic).Repo(repo).RepoIds(repoIds).Status(status).Contributor(contributor).Execute()

Searches pull requests using filters and paginates them

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
    orderBy := openapiclient.PullRequestOrderFieldsEnum("created_at") // PullRequestOrderFieldsEnum |  (optional)
    filter := openapiclient.InsightFilterFieldsEnum("recent") // InsightFilterFieldsEnum |  (optional)
    topic := "javascript" // string |  (optional)
    repo := "open-sauced/insights" // string |  (optional)
    repoIds := "repoIds_example" // string |  (optional)
    status := openapiclient.PullRequestStatusEnum("open") // PullRequestStatusEnum |  (optional)
    contributor := "bdougie" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PullRequestsServiceAPI.SearchAllPullRequests(context.Background()).Page(page).Limit(limit).OrderDirection(orderDirection).Range_(range_).OrderBy(orderBy).Filter(filter).Topic(topic).Repo(repo).RepoIds(repoIds).Status(status).Contributor(contributor).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PullRequestsServiceAPI.SearchAllPullRequests``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchAllPullRequests`: FindContributorPullRequests200Response
    fmt.Fprintf(os.Stdout, "Response from `PullRequestsServiceAPI.SearchAllPullRequests`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchAllPullRequestsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **float32** |  | [default to 1]
 **limit** | **float32** |  | [default to 10]
 **orderDirection** | [**OrderDirectionEnum**](OrderDirectionEnum.md) |  | 
 **range_** | **float32** | Range in days | [default to 30]
 **orderBy** | [**PullRequestOrderFieldsEnum**](PullRequestOrderFieldsEnum.md) |  | 
 **filter** | [**InsightFilterFieldsEnum**](InsightFilterFieldsEnum.md) |  | 
 **topic** | **string** |  | 
 **repo** | **string** |  | 
 **repoIds** | **string** |  | 
 **status** | [**PullRequestStatusEnum**](PullRequestStatusEnum.md) |  | 
 **contributor** | **string** |  | 

### Return type

[**FindContributorPullRequests200Response**](FindContributorPullRequests200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

