# DbContributionsProjects

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrgId** | **string** | The org name of the repo | 
**ProjectId** | **string** | The project name of the repo | 
**RepoId** | **int32** | The ID of the associated repo in the OpenSauced API context | 
**Contributions** | **int32** | The number of contributions associated with a org/repo | 

## Methods

### NewDbContributionsProjects

`func NewDbContributionsProjects(orgId string, projectId string, repoId int32, contributions int32, ) *DbContributionsProjects`

NewDbContributionsProjects instantiates a new DbContributionsProjects object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDbContributionsProjectsWithDefaults

`func NewDbContributionsProjectsWithDefaults() *DbContributionsProjects`

NewDbContributionsProjectsWithDefaults instantiates a new DbContributionsProjects object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrgId

`func (o *DbContributionsProjects) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *DbContributionsProjects) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *DbContributionsProjects) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.


### GetProjectId

`func (o *DbContributionsProjects) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *DbContributionsProjects) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *DbContributionsProjects) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetRepoId

`func (o *DbContributionsProjects) GetRepoId() int32`

GetRepoId returns the RepoId field if non-nil, zero value otherwise.

### GetRepoIdOk

`func (o *DbContributionsProjects) GetRepoIdOk() (*int32, bool)`

GetRepoIdOk returns a tuple with the RepoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepoId

`func (o *DbContributionsProjects) SetRepoId(v int32)`

SetRepoId sets RepoId field to given value.


### GetContributions

`func (o *DbContributionsProjects) GetContributions() int32`

GetContributions returns the Contributions field if non-nil, zero value otherwise.

### GetContributionsOk

`func (o *DbContributionsProjects) GetContributionsOk() (*int32, bool)`

GetContributionsOk returns a tuple with the Contributions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContributions

`func (o *DbContributionsProjects) SetContributions(v int32)`

SetContributions sets Contributions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


