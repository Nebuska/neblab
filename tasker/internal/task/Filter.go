package task

type Filter struct {
	Search  string
	BoardId []uint
	Status  []string

	SortBy       string
	ReversedSort bool
	Before       uint
	After        uint

	PageSize   int
	PageNumber int

	AccessibleByUserId uint
}

func (receiver Filter) OrderBy() string {
	if receiver.ReversedSort {
		return receiver.SortBy + " DESC"
	}
	return receiver.SortBy
}
