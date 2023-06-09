# BetaAppClipInvocationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**BetaAppClipInvocation**](BetaAppClipInvocation.md) |  | 
**Included** | Pointer to [**[]BetaAppClipInvocationLocalization**](BetaAppClipInvocationLocalization.md) |  | [optional] 
**Links** | [**DocumentLinks**](DocumentLinks.md) |  | 

## Methods

### NewBetaAppClipInvocationResponse

`func NewBetaAppClipInvocationResponse(data BetaAppClipInvocation, links DocumentLinks, ) *BetaAppClipInvocationResponse`

NewBetaAppClipInvocationResponse instantiates a new BetaAppClipInvocationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBetaAppClipInvocationResponseWithDefaults

`func NewBetaAppClipInvocationResponseWithDefaults() *BetaAppClipInvocationResponse`

NewBetaAppClipInvocationResponseWithDefaults instantiates a new BetaAppClipInvocationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *BetaAppClipInvocationResponse) GetData() BetaAppClipInvocation`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *BetaAppClipInvocationResponse) GetDataOk() (*BetaAppClipInvocation, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *BetaAppClipInvocationResponse) SetData(v BetaAppClipInvocation)`

SetData sets Data field to given value.


### GetIncluded

`func (o *BetaAppClipInvocationResponse) GetIncluded() []BetaAppClipInvocationLocalization`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *BetaAppClipInvocationResponse) GetIncludedOk() (*[]BetaAppClipInvocationLocalization, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *BetaAppClipInvocationResponse) SetIncluded(v []BetaAppClipInvocationLocalization)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *BetaAppClipInvocationResponse) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.

### GetLinks

`func (o *BetaAppClipInvocationResponse) GetLinks() DocumentLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *BetaAppClipInvocationResponse) GetLinksOk() (*DocumentLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *BetaAppClipInvocationResponse) SetLinks(v DocumentLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


