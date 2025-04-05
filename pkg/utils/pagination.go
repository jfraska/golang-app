package utils

import (
	"go.mongodb.org/mongo-driver/bson"
)

type Pagination struct {
	Limit      int    `json:"limit"`
	Page       int    `json:"page"`
	Sort       string `json:"sort,omitempty"`
	TotalRows  int    `json:"total_rows"`
	TotalPages int    `json:"total_pages"`
	TotalSize  int    `json:"total_size,omitempty"`
}

type PaginationRequestPayload struct {
	Limit int    `query:"limit" json:"limit"`
	Page  int    `query:"page" json:"page"`
	Sort  string `query:"sort" json:"sort"`
}

func NewPaginationFromPaginationRequest(req PaginationRequestPayload) Pagination {

	if req.Limit < 1 {
		req.Limit = 5
	}

	if req.Page < 1 {
		req.Page = 1
	}

	return Pagination{
		Limit: req.Limit,
		Page:  req.Page,
		Sort:  req.Sort,
	}
}

func NewPaginationFromModel(result bson.M) Pagination {

	return Pagination{
		Limit:      int(result["limit"].(int32)),
		Page:       int(result["page"].(int32)),
		TotalRows:  int(result["total_rows"].(int32)),
		TotalPages: int(result["total_pages"].(float64)),
		TotalSize:  int(result["total_size"].(int64)),
	}

}
