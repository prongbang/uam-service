package core

type Filter struct {
	OffsetNo int64 `json:"page"`
	LimitNo  int64 `json:"limit"`
}
