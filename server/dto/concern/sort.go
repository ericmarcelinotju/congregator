package concern

import (
	"fmt"
	"strings"

	"gorm.io/gorm"
)

type SortReq struct {
	Sort *string `json:"sort" form:"sort" uri:"sort"`
}

func (s SortReq) Scope(db *gorm.DB) *gorm.DB {
	sorts := strings.Split(*s.Sort, ":")
	if len(sorts) == 2 {
		column := sorts[0]
		order := sorts[1]

		if len(column) > 0 && len(order) > 0 {
			order := fmt.Sprintf("%s %s", column, order)
			return db.Order(order)
		}
	}
	return db.Order("created_at DESC")
}
