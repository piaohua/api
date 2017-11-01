package xfapi

type XfapiConfig struct {
	AppId              string //应用唯一标识
	AppSecret          string //应用密钥
	AppToken           string //应用级别token
	GetTokenUrl        string //app级别token获取地址
	GetUserTokenUrl    string //用户token获取地址
	UpdateUserTokenUrl string //用户token更新地址
	//用户系统对接
	UserImportUrl string //导入用户地址
	UserUpdateUrl string //用户信息更新地址
	UserDeleteUrl string //用户信息删除地址
	//群组（讨论组）
	GroupCreateUrl     string //创建群组地址
	GroupAddUrl        string //添加群组（讨论组）成员地址
	GroupJoinUrl       string //加入群组(仅适用于群组)地址
	GroupExitUrl       string //退出群组地址
	GroupKickUrl       string //踢除群组（讨论组）成员地址
	GroupRemoveUrl     string //解散群组地址
	GroupInfoUrl       string //获取群组信息地址
	GroupAppointmgrUrl string //任命管理员地址
	GroupRevokemgrUrl  string //移除管理员地址
	GroupListUrl       string //获取群组列表地址
	GroupVerifyUrl     string //群组验证地址
	GroupEditUrl       string //编辑群组资料地址
	GroupTransferUrl   string //移交群主地址
	GroupSearchUrl     string //群组搜索地址
}
