# BetaAppReviewDetailUpdateRequestData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Id** | **string** |  | 
**Attributes** | Pointer to [**AppStoreReviewDetailAttributes**](AppStoreReviewDetailAttributes.md) |  | [optional] 

## Methods

### NewBetaAppReviewDetailUpdateRequestData

`func NewBetaAppReviewDetailUpdateRequestData(type_ string, id string, ) *BetaAppReviewDetailUpdateRequestData`

NewBetaAppReviewDetailUpdateRequestData instantiates a new BetaAppReviewDetailUpdateRequestData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBetaAppReviewDetailUpdateRequestDataWithDefaults

`func NewBetaAppReviewDetailUpdateRequestDataWithDefaults() *BetaAppReviewDetailUpdateRequestData`

NewBetaAppReviewDetailUpdateRequestDataWithDefaults instantiates a new BetaAppReviewDetailUpdateRequestData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *BetaAppReviewDetailUpdateRequestData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BetaAppReviewDetailUpdateRequestData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BetaAppReviewDetailUpdateRequestData) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *BetaAppReviewDetailUpdateRequestData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BetaAppReviewDetailUpdateRequestData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BetaAppReviewDetailUpdateRequestData) SetId(v string)`

SetId sets Id field to given value.


### GetAttributes

`func (o *BetaAppReviewDetailUpdateRequestData) GetAttributes() AppStoreReviewDetailAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *BetaAppReviewDetailUpdateRequestData) GetAttributesOk() (*AppStoreReviewDetailAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *BetaAppReviewDetailUpdateRequestData) SetAttributes(v AppStoreReviewDetailAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *BetaAppReviewDetailUpdateRequestData) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


