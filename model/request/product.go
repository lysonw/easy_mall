package request

type ProductDetailReq struct {
	Pid uint64 `form:"pid" json:"pid" binding:"require"`
}
