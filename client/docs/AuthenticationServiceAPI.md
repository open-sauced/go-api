# \AuthenticationServiceAPI

All URIs are relative to *http://localhost:3001*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CheckAuthSession**](AuthenticationServiceAPI.md#CheckAuthSession) | **Get** /v1/auth/session | Get authenticated session information
[**PostCreateCheckoutSession**](AuthenticationServiceAPI.md#PostCreateCheckoutSession) | **Post** /v1/auth/checkout/session | Creates a new checkout session for the user
[**PostOnboarding**](AuthenticationServiceAPI.md#PostOnboarding) | **Post** /v1/auth/onboarding | Updates onboarding information for user
[**PostWaitlist**](AuthenticationServiceAPI.md#PostWaitlist) | **Post** /v1/auth/waitlist | Updates waitlist information for user
[**UpdateEmailPreferencesForUserProfile**](AuthenticationServiceAPI.md#UpdateEmailPreferencesForUserProfile) | **Patch** /v1/auth/profile/email | Updates the email preferences for the authenticated user profile
[**UpdateInterestsForUserProfile**](AuthenticationServiceAPI.md#UpdateInterestsForUserProfile) | **Patch** /v1/auth/profile/interests | Updates the interests for the authenticated user profile
[**UpdateProfileForUser**](AuthenticationServiceAPI.md#UpdateProfileForUser) | **Patch** /v1/auth/profile | Updates the profile for the authenticated user



## CheckAuthSession

> SupabaseAuthDto CheckAuthSession(ctx).Execute()

Get authenticated session information

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
    resp, r, err := apiClient.AuthenticationServiceAPI.CheckAuthSession(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationServiceAPI.CheckAuthSession``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CheckAuthSession`: SupabaseAuthDto
    fmt.Fprintf(os.Stdout, "Response from `AuthenticationServiceAPI.CheckAuthSession`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCheckAuthSessionRequest struct via the builder pattern


### Return type

[**SupabaseAuthDto**](SupabaseAuthDto.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCreateCheckoutSession

> SupabaseAuthDto PostCreateCheckoutSession(ctx).Execute()

Creates a new checkout session for the user

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
    resp, r, err := apiClient.AuthenticationServiceAPI.PostCreateCheckoutSession(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationServiceAPI.PostCreateCheckoutSession``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostCreateCheckoutSession`: SupabaseAuthDto
    fmt.Fprintf(os.Stdout, "Response from `AuthenticationServiceAPI.PostCreateCheckoutSession`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostCreateCheckoutSessionRequest struct via the builder pattern


### Return type

[**SupabaseAuthDto**](SupabaseAuthDto.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostOnboarding

> SupabaseAuthDto PostOnboarding(ctx).UserOnboardingDto(userOnboardingDto).Execute()

Updates onboarding information for user

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
    userOnboardingDto := *openapiclient.NewUserOnboardingDto([]string{"Interests_example"}, "UTC-5") // UserOnboardingDto | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthenticationServiceAPI.PostOnboarding(context.Background()).UserOnboardingDto(userOnboardingDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationServiceAPI.PostOnboarding``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostOnboarding`: SupabaseAuthDto
    fmt.Fprintf(os.Stdout, "Response from `AuthenticationServiceAPI.PostOnboarding`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostOnboardingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userOnboardingDto** | [**UserOnboardingDto**](UserOnboardingDto.md) |  | 

### Return type

[**SupabaseAuthDto**](SupabaseAuthDto.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostWaitlist

> SupabaseAuthDto PostWaitlist(ctx).Execute()

Updates waitlist information for user

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
    resp, r, err := apiClient.AuthenticationServiceAPI.PostWaitlist(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationServiceAPI.PostWaitlist``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostWaitlist`: SupabaseAuthDto
    fmt.Fprintf(os.Stdout, "Response from `AuthenticationServiceAPI.PostWaitlist`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostWaitlistRequest struct via the builder pattern


### Return type

[**SupabaseAuthDto**](SupabaseAuthDto.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateEmailPreferencesForUserProfile

> DbUser UpdateEmailPreferencesForUserProfile(ctx).UpdateUserEmailPreferencesDto(updateUserEmailPreferencesDto).Execute()

Updates the email preferences for the authenticated user profile

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
    updateUserEmailPreferencesDto := *openapiclient.NewUpdateUserEmailPreferencesDto(false, false) // UpdateUserEmailPreferencesDto | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthenticationServiceAPI.UpdateEmailPreferencesForUserProfile(context.Background()).UpdateUserEmailPreferencesDto(updateUserEmailPreferencesDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationServiceAPI.UpdateEmailPreferencesForUserProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateEmailPreferencesForUserProfile`: DbUser
    fmt.Fprintf(os.Stdout, "Response from `AuthenticationServiceAPI.UpdateEmailPreferencesForUserProfile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateEmailPreferencesForUserProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateUserEmailPreferencesDto** | [**UpdateUserEmailPreferencesDto**](UpdateUserEmailPreferencesDto.md) |  | 

### Return type

[**DbUser**](DbUser.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateInterestsForUserProfile

> DbUser UpdateInterestsForUserProfile(ctx).UpdateUserProfileInterestsDto(updateUserProfileInterestsDto).Execute()

Updates the interests for the authenticated user profile

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
    updateUserProfileInterestsDto := *openapiclient.NewUpdateUserProfileInterestsDto([]string{"Interests_example"}) // UpdateUserProfileInterestsDto | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthenticationServiceAPI.UpdateInterestsForUserProfile(context.Background()).UpdateUserProfileInterestsDto(updateUserProfileInterestsDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationServiceAPI.UpdateInterestsForUserProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateInterestsForUserProfile`: DbUser
    fmt.Fprintf(os.Stdout, "Response from `AuthenticationServiceAPI.UpdateInterestsForUserProfile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateInterestsForUserProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateUserProfileInterestsDto** | [**UpdateUserProfileInterestsDto**](UpdateUserProfileInterestsDto.md) |  | 

### Return type

[**DbUser**](DbUser.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateProfileForUser

> DbUser UpdateProfileForUser(ctx).UpdateUserDto(updateUserDto).Execute()

Updates the profile for the authenticated user

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
    updateUserDto := *openapiclient.NewUpdateUserDto("Pizza Maker", "hello@opensauced.pizza") // UpdateUserDto | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthenticationServiceAPI.UpdateProfileForUser(context.Background()).UpdateUserDto(updateUserDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationServiceAPI.UpdateProfileForUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateProfileForUser`: DbUser
    fmt.Fprintf(os.Stdout, "Response from `AuthenticationServiceAPI.UpdateProfileForUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateProfileForUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateUserDto** | [**UpdateUserDto**](UpdateUserDto.md) |  | 

### Return type

[**DbUser**](DbUser.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

