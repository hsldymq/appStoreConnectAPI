# AppInfoLocalizationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**AppInfoLocalization**](AppInfoLocalization.md) |  | 
**Included** | Pointer to [**[]AppInfo**](AppInfo.md) |  | [optional] 
**Links** | [**DocumentLinks**](DocumentLinks.md) |  | 

## Methods

### NewAppInfoLocalizationResponse

`func NewAppInfoLocalizationResponse(data AppInfoLocalization, links DocumentLinks, ) *AppInfoLocalizationResponse`

NewAppInfoLocalizationResponse instantiates a new AppInfoLocalizationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppInfoLocalizationResponseWithDefaults

`func NewAppInfoLocalizationResponseWithDefaults() *AppInfoLocalizationResponse`

NewAppInfoLocalizationResponseWithDefaults instantiates a new AppInfoLocalizationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *AppInfoLocalizationResponse) GetData() AppInfoLocalization`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AppInfoLocalizationResponse) GetDataOk() (*AppInfoLocalization, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AppInfoLocalizationResponse) SetData(v AppInfoLocalization)`

SetData sets Data field to given value.


### GetIncluded

`func (o *AppInfoLocalizationResponse) GetIncluded() []AppInfo`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *AppInfoLocalizationResponse) GetIncludedOk() (*[]AppInfo, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *AppInfoLocalizationResponse) SetIncluded(v []AppInfo)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *AppInfoLocalizationResponse) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.

### GetLinks

`func (o *AppInfoLocalizationResponse) GetLinks() DocumentLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *AppInfoLocalizationResponse) GetLinksOk() (*DocumentLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *AppInfoLocalizationResponse) SetLinks(v DocumentLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


