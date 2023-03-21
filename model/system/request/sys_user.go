package request

type Register struct {
	Username  string `json:"username"`                                                                // 用户登录名
	Password  string `json:"password"`                                                                // 用户登录密码
	NickName  string `json:"nickName" gorm:"default:系统用户"`                                            // 用户昵称 	// 用户侧边主题
	HeaderImg string `json:"headerImg" gorm:"default:https://qmplusimg.henrongyi.top/gva_header.jpg"` // 用户头像 	// 用户角色ID
	Phone     string `json:"phone"`                                                                   // 用户手机号
	Email     string `json:"email"`
}
