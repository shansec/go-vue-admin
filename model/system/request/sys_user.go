package request

// Register 用户注册
type Register struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	NickName  string `json:"nickName" gorm:"default:'PlusUser'"`
	HeaderImg string `json:"headerImg" gorm:"default:'1111'"`
}

// Login 用户登录
type Login struct {
	// 用户名
	Username string `json:"username"`
	// 密码
	Password string `json:"password"`
}
