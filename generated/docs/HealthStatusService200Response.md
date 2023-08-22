# HealthStatusService200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** |  | [optional] 
**Info** | Pointer to [**map[string]HealthStatusService200ResponseInfoValue**](HealthStatusService200ResponseInfoValue.md) |  | [optional] 
**Error** | Pointer to [**map[string]HealthStatusService200ResponseInfoValue**](HealthStatusService200ResponseInfoValue.md) |  | [optional] 
**Details** | Pointer to [**map[string]HealthStatusService200ResponseInfoValue**](HealthStatusService200ResponseInfoValue.md) |  | [optional] 

## Methods

### NewHealthStatusService200Response

`func NewHealthStatusService200Response() *HealthStatusService200Response`

NewHealthStatusService200Response instantiates a new HealthStatusService200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHealthStatusService200ResponseWithDefaults

`func NewHealthStatusService200ResponseWithDefaults() *HealthStatusService200Response`

NewHealthStatusService200ResponseWithDefaults instantiates a new HealthStatusService200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *HealthStatusService200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *HealthStatusService200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *HealthStatusService200Response) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *HealthStatusService200Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetInfo

`func (o *HealthStatusService200Response) GetInfo() map[string]HealthStatusService200ResponseInfoValue`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *HealthStatusService200Response) GetInfoOk() (*map[string]HealthStatusService200ResponseInfoValue, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *HealthStatusService200Response) SetInfo(v map[string]HealthStatusService200ResponseInfoValue)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *HealthStatusService200Response) HasInfo() bool`

HasInfo returns a boolean if a field has been set.

### SetInfoNil

`func (o *HealthStatusService200Response) SetInfoNil(b bool)`

 SetInfoNil sets the value for Info to be an explicit nil

### UnsetInfo
`func (o *HealthStatusService200Response) UnsetInfo()`

UnsetInfo ensures that no value is present for Info, not even an explicit nil
### GetError

`func (o *HealthStatusService200Response) GetError() map[string]HealthStatusService200ResponseInfoValue`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *HealthStatusService200Response) GetErrorOk() (*map[string]HealthStatusService200ResponseInfoValue, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *HealthStatusService200Response) SetError(v map[string]HealthStatusService200ResponseInfoValue)`

SetError sets Error field to given value.

### HasError

`func (o *HealthStatusService200Response) HasError() bool`

HasError returns a boolean if a field has been set.

### SetErrorNil

`func (o *HealthStatusService200Response) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *HealthStatusService200Response) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil
### GetDetails

`func (o *HealthStatusService200Response) GetDetails() map[string]HealthStatusService200ResponseInfoValue`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *HealthStatusService200Response) GetDetailsOk() (*map[string]HealthStatusService200ResponseInfoValue, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *HealthStatusService200Response) SetDetails(v map[string]HealthStatusService200ResponseInfoValue)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *HealthStatusService200Response) HasDetails() bool`

HasDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


