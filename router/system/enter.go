package system

type RouterGroup struct {
	UserRouter
	BaseRouter
	DeptRouter
	SystemConfigRouter
	AutoCodeRouter
	InitRouter
	ApiRouter
	RoleRouter
}
