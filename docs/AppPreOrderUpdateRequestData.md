# AppPreOrderUpdateRequestData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Id** | **string** |  | 
**Attributes** | Pointer to [**AppPreOrderCreateRequestDataAttributes**](AppPreOrderCreateRequestDataAttributes.md) |  | [optional] 

## Methods

### NewAppPreOrderUpdateRequestData

`func NewAppPreOrderUpdateRequestData(type_ string, id string, ) *AppPreOrderUpdateRequestData`

NewAppPreOrderUpdateRequestData instantiates a new AppPreOrderUpdateRequestData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppPreOrderUpdateRequestDataWithDefaults

`func NewAppPreOrderUpdateRequestDataWithDefaults() *AppPreOrderUpdateRequestData`

NewAppPreOrderUpdateRequestDataWithDefaults instantiates a new AppPreOrderUpdateRequestData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AppPreOrderUpdateRequestData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AppPreOrderUpdateRequestData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AppPreOrderUpdateRequestData) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *AppPreOrderUpdateRequestData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AppPreOrderUpdateRequestData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AppPreOrderUpdateRequestData) SetId(v string)`

SetId sets Id field to given value.


### GetAttributes

`func (o *AppPreOrderUpdateRequestData) GetAttributes() AppPreOrderCreateRequestDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *AppPreOrderUpdateRequestData) GetAttributesOk() (*AppPreOrderCreateRequestDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *AppPreOrderUpdateRequestData) SetAttributes(v AppPreOrderCreateRequestDataAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *AppPreOrderUpdateRequestData) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


