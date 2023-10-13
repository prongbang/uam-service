package core

type Params struct {
	OffsetNo int64 `json:"page"`
	LimitNo  int64 `json:"limit"`
}
