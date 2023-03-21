package request

type PageInfo struct {
	// 页码
	Page int `json:"page" form:"page"`
	// 每页大小
	PagSize int `json:"pagSize" form:"pageSize"`
}

type GetById struct {
	ID int `json:"id" form:"id"`
}

func (r *GetById) Uint() uint {
	return uint(r.ID)
}
