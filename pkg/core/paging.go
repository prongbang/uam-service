package core

import (
	"math"
)

const PagingLimitDefault = 20

type Paging struct {
	List  any   `json:"list"`
	Page  int   `json:"page"`
	Limit int   `json:"limit"`
	Count int64 `json:"count"`
	Total int64 `json:"total"`
	Start int64 `json:"start"`
	End   int64 `json:"end"`
}

func Offset(pageNo int, limitNo int) int {
	return (limitNo * pageNo) - limitNo
}

func Pagination(pageNo int, limitNo int, getCount func() int64, getData func(limit int, offset int) any) Paging {
	total := getCount()
	pageCount := math.Ceil(float64(total) / float64(limitNo))
	pageCountInt := int64(pageCount)
	if pageNo <= 0 {
		pageNo = 1
	}
	offset := Offset(pageNo, limitNo)
	startRow := (pageNo - 1) * limitNo
	endRow := startRow + limitNo - 1

	data := getData(limitNo, offset)

	return Paging{
		List:  data,
		Page:  pageNo,
		Limit: limitNo,
		Count: pageCountInt,
		Total: total,
		Start: int64(startRow),
		End:   int64(endRow),
	}
}
