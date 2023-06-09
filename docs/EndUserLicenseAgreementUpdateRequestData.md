# EndUserLicenseAgreementUpdateRequestData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Id** | **string** |  | 
**Attributes** | Pointer to [**BetaLicenseAgreementAttributes**](BetaLicenseAgreementAttributes.md) |  | [optional] 
**Relationships** | Pointer to [**EndUserLicenseAgreementUpdateRequestDataRelationships**](EndUserLicenseAgreementUpdateRequestDataRelationships.md) |  | [optional] 

## Methods

### NewEndUserLicenseAgreementUpdateRequestData

`func NewEndUserLicenseAgreementUpdateRequestData(type_ string, id string, ) *EndUserLicenseAgreementUpdateRequestData`

NewEndUserLicenseAgreementUpdateRequestData instantiates a new EndUserLicenseAgreementUpdateRequestData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEndUserLicenseAgreementUpdateRequestDataWithDefaults

`func NewEndUserLicenseAgreementUpdateRequestDataWithDefaults() *EndUserLicenseAgreementUpdateRequestData`

NewEndUserLicenseAgreementUpdateRequestDataWithDefaults instantiates a new EndUserLicenseAgreementUpdateRequestData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *EndUserLicenseAgreementUpdateRequestData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EndUserLicenseAgreementUpdateRequestData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EndUserLicenseAgreementUpdateRequestData) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *EndUserLicenseAgreementUpdateRequestData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EndUserLicenseAgreementUpdateRequestData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EndUserLicenseAgreementUpdateRequestData) SetId(v string)`

SetId sets Id field to given value.


### GetAttributes

`func (o *EndUserLicenseAgreementUpdateRequestData) GetAttributes() BetaLicenseAgreementAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *EndUserLicenseAgreementUpdateRequestData) GetAttributesOk() (*BetaLicenseAgreementAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *EndUserLicenseAgreementUpdateRequestData) SetAttributes(v BetaLicenseAgreementAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *EndUserLicenseAgreementUpdateRequestData) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetRelationships

`func (o *EndUserLicenseAgreementUpdateRequestData) GetRelationships() EndUserLicenseAgreementUpdateRequestDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *EndUserLicenseAgreementUpdateRequestData) GetRelationshipsOk() (*EndUserLicenseAgreementUpdateRequestDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *EndUserLicenseAgreementUpdateRequestData) SetRelationships(v EndUserLicenseAgreementUpdateRequestDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *EndUserLicenseAgreementUpdateRequestData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


