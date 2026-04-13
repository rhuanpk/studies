package model

import (
	"database/sql"
	"math"
	"regexp"
)

// Pagination is the struct for pagination.
type Pagination struct {
	PageNumber int `json:"page_number" form:"page_number" binding:"required_with=PageSize,numeric"`
	PageSize   int `json:"page_size" form:"page_size" binding:"required_with=PageNumber,numeric"`
	LastPage   int `json:"last_page"`
	Total      int `json:"total"`
}

// SetInfos sets items total and last page in pagination.
//
// If overrise param is false, inform the original query.
// This func will trim prefix select clause and suffix
// limit and offset automatic.
//
// If override param is true, inform the custom query to get the total.
// The select clause must contain some "as total".
func (p *Pagination) SetInfos(db *sql.DB, query string, override bool, args ...any) error {
	// parse query if no override from outside
	if !override {
		query = `select count(*) total from` +
			regexp.
				MustCompile(`(?is)(^(\s+)?select.*from|(limit|offset)\s+\d+)`).
				ReplaceAllString(query, "")
	}

	// exec total query
	row := db.QueryRow(query, args...)
	if row.Err() != nil {
		return row.Err()
	}

	// scan total value
	err := row.Scan(&p.Total)
	if err != nil {
		return err
	}

	// set last page
	p.LastPage = int(math.Round(math.Ceil(float64(p.Total) / float64(p.PageSize))))

	// no errors
	return nil
}

// Offset return the offset of pagination.
func (p Pagination) Offset() int {
	return p.PageSize * (p.PageNumber - 1)
}

// IsUsing return if pagination is using.
func (p Pagination) IsUsing() bool {
	if p.PageNumber >= 1 && p.PageSize >= 1 {
		return true
	}
	return false
}
