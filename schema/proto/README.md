# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [user.proto](#user.proto)
    - [CreateUserReq](#user.CreateUserReq)
    - [CreateUserRes](#user.CreateUserRes)
    - [GetUserByEmailAndPasswordReq](#user.GetUserByEmailAndPasswordReq)
    - [GetUserReq](#user.GetUserReq)
    - [GetUserRes](#user.GetUserRes)
    - [User](#user.User)
  
    - [UserService](#user.UserService)
  
- [Scalar Value Types](#scalar-value-types)



<a name="user.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## user.proto



<a name="user.CreateUserReq"></a>

### CreateUserReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| email | [string](#string) |  |  |
| password | [string](#string) |  |  |
| name | [string](#string) |  |  |






<a name="user.CreateUserRes"></a>

### CreateUserRes



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user | [User](#user.User) |  |  |






<a name="user.GetUserByEmailAndPasswordReq"></a>

### GetUserByEmailAndPasswordReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| email | [string](#string) |  |  |
| password | [string](#string) |  |  |






<a name="user.GetUserReq"></a>

### GetUserReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  |  |






<a name="user.GetUserRes"></a>

### GetUserRes



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user | [User](#user.User) |  |  |






<a name="user.User"></a>

### User



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  |  |
| email | [string](#string) |  |  |
| name | [string](#string) |  |  |
| created_at | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| updated_at | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |





 

 

 


<a name="user.UserService"></a>

### UserService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| GetUser | [GetUserReq](#user.GetUserReq) | [GetUserRes](#user.GetUserRes) |  |
| GetUserByEmailAndPassword | [GetUserByEmailAndPasswordReq](#user.GetUserByEmailAndPasswordReq) | [GetUserRes](#user.GetUserRes) |  |
| CreateUser | [CreateUserReq](#user.CreateUserReq) | [CreateUserRes](#user.CreateUserRes) |  |

 



## Scalar Value Types

| .proto Type | Notes | C++ | Java | Python | Go | C# | PHP | Ruby |
| ----------- | ----- | --- | ---- | ------ | -- | -- | --- | ---- |
| <a name="double" /> double |  | double | double | float | float64 | double | float | Float |
| <a name="float" /> float |  | float | float | float | float32 | float | float | Float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum or Fixnum (as required) |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="bool" /> bool |  | bool | boolean | boolean | bool | bool | boolean | TrueClass/FalseClass |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode | string | string | string | String (UTF-8) |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str | []byte | ByteString | string | String (ASCII-8BIT) |

