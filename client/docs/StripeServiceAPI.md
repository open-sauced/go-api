# \StripeServiceAPI

All URIs are relative to *https://api.opensauced.pizza*

Method | HTTP request | Description
------------- | ------------- | -------------
[**HandleStripeWebhook**](StripeServiceAPI.md#HandleStripeWebhook) | **Post** /v1/stripe/webhooks | 



## HandleStripeWebhook

> HandleStripeWebhook(ctx).Execute()



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.StripeServiceAPI.HandleStripeWebhook(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StripeServiceAPI.HandleStripeWebhook``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiHandleStripeWebhookRequest struct via the builder pattern


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

