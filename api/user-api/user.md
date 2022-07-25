### 1. "发送验证码"

1. route definition

- Url: /user/v1/sent-phone-code/:phone
- Method: GET
- Request: `sentPhoneCodeReq`
- Response: `sentPhoneCodeResp`

2. request definition



```golang
type SentPhoneCodeReq struct {
	Phone string `path:"phone"`
}
```


3. response definition



```golang
type SentPhoneCodeResp struct {
	Code int32 `json:"code"`
	Msg string `json:"msg"`
	Msg_id uint64 `json:"msg_id"`
}
```

### 2. "注册"

1. route definition

- Url: /user/v1/signup
- Method: POST
- Request: `signUpReq`
- Response: `signUpResp`

2. request definition



```golang
type SignUpReq struct {
	Phone string `json:"phone"`
	Phone_code string `json:"phone_code"`
	Password string `json:"password"`
	Msg_id uint64 `json:"msg_id"`
}
```


3. response definition



```golang
type SignUpResp struct {
	Code int32 `json:"code"`
	Msg string `json:"msg"`
	User_id int64 `json:"user_id"`
	Token string `json:"token"`
}
```

### 3. "登录"

1. route definition

- Url: /user/v1/signin
- Method: POST
- Request: `signInReq`
- Response: `signInResp`

2. request definition



```golang
type SignInReq struct {
	Phone string `json:"phone"`
	Password string `json:"password"`
}
```


3. response definition



```golang
type SignInResp struct {
	Code int32 `json:"code"`
	Msg string `json:"msg"`
	Token string `json:"token"`
}
```

### 4. "获取当前用户信息 jwt"

1. route definition

- Url: /account/v1/me
- Method: GET
- Request: `-`
- Response: `meResp`

2. request definition



3. response definition



```golang
type MeResp struct {
	Code int32 `json:"code"`
	Msg string `json:"msg"`
	UserId string `json:"user_id"`
	Name string `json:"name"`
	Avatar string `json:"avatar"`
	Buid int64 `json:"buid"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}
```

### 5. "设置用户信息"

1. route definition

- Url: /account/v1/setInfo
- Method: POST
- Request: `setInfoReq`
- Response: `setInfoResp`

2. request definition



```golang
type SetInfoReq struct {
	Name string `json:"name"`
	Avatar string `json:"avatar"`
	Buid string `json:"buid"`
}
```


3. response definition



```golang
type SetInfoResp struct {
	Code int32 `json:"code"`
	Msg string `json:"msg"`
}
```

### 6. "根据 buid 获取 user_id 仅测试用"

1. route definition

- Url: /app/v1/get-userid-by-buid
- Method: GET
- Request: `getUserIdByBuidReq`
- Response: `getUserIdByBuidResp`

2. request definition



```golang
type GetUserIdByBuidReq struct {
	Buid int64 `json:"buid"`
}
```


3. response definition



```golang
type GetUserIdByBuidResp struct {
	Code int32 `json:"code"`
	Msg string `json:"msg"`
	UserId string `json:"user_id"`
}
```

