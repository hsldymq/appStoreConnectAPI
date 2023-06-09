# BuildUpdateRequestData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Id** | **string** |  | 
**Attributes** | Pointer to [**BuildUpdateRequestDataAttributes**](BuildUpdateRequestDataAttributes.md) |  | [optional] 
**Relationships** | Pointer to [**BuildUpdateRequestDataRelationships**](BuildUpdateRequestDataRelationships.md) |  | [optional] 

## Methods

### NewBuildUpdateRequestData

`func NewBuildUpdateRequestData(type_ string, id string, ) *BuildUpdateRequestData`

NewBuildUpdateRequestData instantiates a new BuildUpdateRequestData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBuildUpdateRequestDataWithDefaults

`func NewBuildUpdateRequestDataWithDefaults() *BuildUpdateRequestData`

NewBuildUpdateRequestDataWithDefaults instantiates a new BuildUpdateRequestData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *BuildUpdateRequestData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BuildUpdateRequestData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BuildUpdateRequestData) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *BuildUpdateRequestData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BuildUpdateRequestData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BuildUpdateRequestData) SetId(v string)`

SetId sets Id field to given value.


### GetAttributes

`func (o *BuildUpdateRequestData) GetAttributes() BuildUpdateRequestDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *BuildUpdateRequestData) GetAttributesOk() (*BuildUpdateRequestDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *BuildUpdateRequestData) SetAttributes(v BuildUpdateRequestDataAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *BuildUpdateRequestData) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetRelationships

`func (o *BuildUpdateRequestData) GetRelationships() BuildUpdateRequestDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *BuildUpdateRequestData) GetRelationshipsOk() (*BuildUpdateRequestDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *BuildUpdateRequestData) SetRelationships(v BuildUpdateRequestDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *BuildUpdateRequestData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


