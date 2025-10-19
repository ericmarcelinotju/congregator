package concern

import (
	"gorm.io/gorm"
)

type PaginationReq struct {
	Limit *int `uri:"limit" binding:"omitempty,min=1"`
	Page  *int `uri:"page" binding:"omitempty,min=1"`
}

func (p PaginationReq) Scope(db *gorm.DB) *gorm.DB {
	if p.Limit != nil && p.Page != nil {

		offset := *p.Limit * (*p.Page - 1)
		if *p.Limit > 0 && offset >= 0 {
			return db.
				Limit(*p.Limit).
				Offset(offset)
		}
	}
	return db
}

type PaginationResp[T any] struct {
	Data  []T   `json:"data"`
	Total int64 `json:"total"`
}
