# BetaTesterBuildsLinkagesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]AppEncryptionDeclarationRelationshipsBuildsDataInner**](AppEncryptionDeclarationRelationshipsBuildsDataInner.md) |  | 
**Links** | [**PagedDocumentLinks**](PagedDocumentLinks.md) |  | 
**Meta** | Pointer to [**PagingInformation**](PagingInformation.md) |  | [optional] 

## Methods

### NewBetaTesterBuildsLinkagesResponse

`func NewBetaTesterBuildsLinkagesResponse(data []AppEncryptionDeclarationRelationshipsBuildsDataInner, links PagedDocumentLinks, ) *BetaTesterBuildsLinkagesResponse`

NewBetaTesterBuildsLinkagesResponse instantiates a new BetaTesterBuildsLinkagesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBetaTesterBuildsLinkagesResponseWithDefaults

`func NewBetaTesterBuildsLinkagesResponseWithDefaults() *BetaTesterBuildsLinkagesResponse`

NewBetaTesterBuildsLinkagesResponseWithDefaults instantiates a new BetaTesterBuildsLinkagesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *BetaTesterBuildsLinkagesResponse) GetData() []AppEncryptionDeclarationRelationshipsBuildsDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *BetaTesterBuildsLinkagesResponse) GetDataOk() (*[]AppEncryptionDeclarationRelationshipsBuildsDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *BetaTesterBuildsLinkagesResponse) SetData(v []AppEncryptionDeclarationRelationshipsBuildsDataInner)`

SetData sets Data field to given value.


### GetLinks

`func (o *BetaTesterBuildsLinkagesResponse) GetLinks() PagedDocumentLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *BetaTesterBuildsLinkagesResponse) GetLinksOk() (*PagedDocumentLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *BetaTesterBuildsLinkagesResponse) SetLinks(v PagedDocumentLinks)`

SetLinks sets Links field to given value.


### GetMeta

`func (o *BetaTesterBuildsLinkagesResponse) GetMeta() PagingInformation`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *BetaTesterBuildsLinkagesResponse) GetMetaOk() (*PagingInformation, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *BetaTesterBuildsLinkagesResponse) SetMeta(v PagingInformation)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *BetaTesterBuildsLinkagesResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


