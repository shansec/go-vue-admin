package request

type PageInfo struct {
	Page     int `json:"page" form:"page"`
	PageSize int `json:"pageSize" form:"pageSize"`
}

type GetById struct {
	ID int `json:"id" form:"id"`
}

func (r *GetById) Uint() uint {
	return uint(r.ID)
}

type IdsReq struct {
	Ids []int `json:"ids" form:"ids"`
}

type GetAuthorityId struct {
	AuthorityId string `json:"authorityId" form:"authorityId"`
}

type Empty struct{}
