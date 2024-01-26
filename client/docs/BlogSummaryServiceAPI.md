# \BlogSummaryServiceAPI

All URIs are relative to *https://api.opensauced.pizza*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GenerateBlogSummary**](BlogSummaryServiceAPI.md#GenerateBlogSummary) | **Post** /v2/blogs/summary/generate | Generate a summary for a markdown supported blog



## GenerateBlogSummary

> GenerateBlogSummary(ctx).CreateBlogSummaryDto(createBlogSummaryDto).Execute()

Generate a summary for a markdown supported blog

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
    createBlogSummaryDto := *openapiclient.NewCreateBlogSummaryDto(int32(250), int32(7), "formal", "english", "BlogTitle_example", "BlogMarkdown_example") // CreateBlogSummaryDto | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.BlogSummaryServiceAPI.GenerateBlogSummary(context.Background()).CreateBlogSummaryDto(createBlogSummaryDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BlogSummaryServiceAPI.GenerateBlogSummary``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGenerateBlogSummaryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createBlogSummaryDto** | [**CreateBlogSummaryDto**](CreateBlogSummaryDto.md) |  | 

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

