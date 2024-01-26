# \IssueSummaryServiceAPI

All URIs are relative to *https://api.opensauced.pizza*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GenerateIssueSummary**](IssueSummaryServiceAPI.md#GenerateIssueSummary) | **Post** /v2/issues/summary/generate | Generate a summary for an issue



## GenerateIssueSummary

> GenerateIssueSummary(ctx).CreateIssueSummaryDto(createIssueSummaryDto).Execute()

Generate a summary for an issue

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
    createIssueSummaryDto := *openapiclient.NewCreateIssueSummaryDto(int32(250), int32(7), "formal", "english", "IssueTitle_example", "IssueDescription_example", "IssueComments_example") // CreateIssueSummaryDto | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.IssueSummaryServiceAPI.GenerateIssueSummary(context.Background()).CreateIssueSummaryDto(createIssueSummaryDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IssueSummaryServiceAPI.GenerateIssueSummary``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGenerateIssueSummaryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createIssueSummaryDto** | [**CreateIssueSummaryDto**](CreateIssueSummaryDto.md) |  | 

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

