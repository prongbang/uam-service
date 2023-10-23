package core

type Params struct {
	OffsetNo int `json:"page"`
	LimitNo  int `json:"limit"`
}

func (p Params) IgnorePaging() {
	p.LimitNo = 0
	p.OffsetNo = -1
}
