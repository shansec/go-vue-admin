package request

type PageInfo struct {
	Page     int `json:"page" form:"page"`         // 页码
	PageSize int `json:"pageSize" form:"pageSize"` // 每页大小
}

type GetById struct {
	ID int `json:"id" form:"id"`
}

func (r *GetById) Uint() uint {
	return uint(r.ID)
}
