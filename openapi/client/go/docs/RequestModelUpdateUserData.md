# RequestModelUpdateUserData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Encrypted** | Pointer to **int32** | Indicates that the data of the packet has been encrypted before transmission and placed in the data field. | [optional] 
**Data** | Pointer to [**RequestModelUpdateUserDataData**](RequestModelUpdateUserDataData.md) |  | [optional] 

## Methods

### NewRequestModelUpdateUserData

`func NewRequestModelUpdateUserData() *RequestModelUpdateUserData`

NewRequestModelUpdateUserData instantiates a new RequestModelUpdateUserData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestModelUpdateUserDataWithDefaults

`func NewRequestModelUpdateUserDataWithDefaults() *RequestModelUpdateUserData`

NewRequestModelUpdateUserDataWithDefaults instantiates a new RequestModelUpdateUserData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEncrypted

`func (o *RequestModelUpdateUserData) GetEncrypted() int32`

GetEncrypted returns the Encrypted field if non-nil, zero value otherwise.

### GetEncryptedOk

`func (o *RequestModelUpdateUserData) GetEncryptedOk() (*int32, bool)`

GetEncryptedOk returns a tuple with the Encrypted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncrypted

`func (o *RequestModelUpdateUserData) SetEncrypted(v int32)`

SetEncrypted sets Encrypted field to given value.

### HasEncrypted

`func (o *RequestModelUpdateUserData) HasEncrypted() bool`

HasEncrypted returns a boolean if a field has been set.

### GetData

`func (o *RequestModelUpdateUserData) GetData() RequestModelUpdateUserDataData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *RequestModelUpdateUserData) GetDataOk() (*RequestModelUpdateUserDataData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *RequestModelUpdateUserData) SetData(v RequestModelUpdateUserDataData)`

SetData sets Data field to given value.

### HasData

`func (o *RequestModelUpdateUserData) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


