package dto

import "github.com/Nebuska/task-tracker/internal/task"

type TaskQuery struct {
	Search  string   `json:"search"`
	BoardId []uint   `json:"board_id"`
	Status  []string `json:"status"`

	SortBy       string `json:"sort_by"`
	ReversedSort bool   `json:"reversed_sort"`
	Before       uint   `json:"before"`
	After        uint   `json:"after"`

	PageSize   int `json:"page_size"`
	PageNumber int `json:"page_number"`
}

func (receiver TaskQuery) ToFilter() task.Filter {
	return task.Filter{
		Search:       receiver.Search,
		BoardId:      receiver.BoardId,
		Status:       receiver.Status,
		SortBy:       receiver.SortBy,
		ReversedSort: receiver.ReversedSort,
		Before:       receiver.Before,
		After:        receiver.After,
		PageSize:     receiver.PageSize,
		PageNumber:   receiver.PageNumber,
	}
}
