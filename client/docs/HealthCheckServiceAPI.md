# \HealthCheckServiceAPI

All URIs are relative to *https://api.opensauced.pizza*

Method | HTTP request | Description
------------- | ------------- | -------------
[**HealthStatusService**](HealthCheckServiceAPI.md#HealthStatusService) | **Get** /v2/health/service | Check the health of Open Sauced service endpoints
[**HealthStatusWeb**](HealthCheckServiceAPI.md#HealthStatusWeb) | **Get** /v2/health/web | Check the health of Open Sauced web endpoints



## HealthStatusService

> HealthStatusService200Response HealthStatusService(ctx).Execute()

Check the health of Open Sauced service endpoints

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
    resp, r, err := apiClient.HealthCheckServiceAPI.HealthStatusService(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HealthCheckServiceAPI.HealthStatusService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `HealthStatusService`: HealthStatusService200Response
    fmt.Fprintf(os.Stdout, "Response from `HealthCheckServiceAPI.HealthStatusService`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiHealthStatusServiceRequest struct via the builder pattern


### Return type

[**HealthStatusService200Response**](HealthStatusService200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HealthStatusWeb

> HealthStatusService200Response HealthStatusWeb(ctx).Execute()

Check the health of Open Sauced web endpoints

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
    resp, r, err := apiClient.HealthCheckServiceAPI.HealthStatusWeb(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HealthCheckServiceAPI.HealthStatusWeb``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `HealthStatusWeb`: HealthStatusService200Response
    fmt.Fprintf(os.Stdout, "Response from `HealthCheckServiceAPI.HealthStatusWeb`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiHealthStatusWebRequest struct via the builder pattern


### Return type

[**HealthStatusService200Response**](HealthStatusService200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

