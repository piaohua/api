# 飞语语音

Go语言飞语语音会议模式后台实现

Backend implementation of feiyu conference in golang 


# usage

```go
//初始化
cfg := &feiyu.NewFeiyuConfig{
}

fy, err := wxapi.NewFeiyuApi(cfg)
if err != nil {
	panic(err)
}

// 往飞语服务器添加飞语云账户的接口，只有添加过的账户才能使用SDK
addAccountResult, err := fy.AddAccount(accId, phone)
if err != nil {
	panic(err)
}
fmt.Println(addAccountResult)

// 查看终端SDK账户
getAccountResult, err := fy.GetAccount(infoType, info)
if err != nil {
	panic(err)
}
fmt.Println(getAccountResult)

//此接口用于离开会议以及要取消整个会议室的时候调用
resultCode := fy.HangUp(accId, cId, hType)
fmt.Println(resultCode)

//此接口用于将用户设置静音
resultCode := fy.Mute(accId, cId, hType)
fmt.Println(resultCode)

//此接口用于将用户设置取消静音
resultCode := fy.Unmute(accId, cId, hType)
fmt.Println(resultCode)

```

# document

