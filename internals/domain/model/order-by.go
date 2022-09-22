package model

type OrderBy bool

const (
	OrderByAsc  = OrderBy(true)  // ASC
	OrderByDesc = OrderBy(false) // DESC
)

func OrderByFromBool(isAsc bool) OrderBy {
	if isAsc {
		return OrderByAsc
	}
	return OrderByDesc
}

func (o OrderBy) Int() int {
	if o {
		return 1
	}
	return -1
}
