// Pagination struct
package db

import "math"

// meta struct. Not using sort this time
type Meta struct {
	Limit      int   `json:"limit,omitempty"`
	Page       int   `json:"page,omitempty"`
	TotalRows  int64 `json:"total_rows,omitempty"`
	TotalPages int   `json:"total_pages,omitempty"`
	//SortBy     string `json:"sort_by,omitempty;query:sort_by"`
	//SortOrder  string `json:"sort_order,omitempty;query:sort_order"`
}

// pagination struct
type Pagination struct {
	Meta    `json:"meta,omitempty"`
	Results interface{} `json:"results"`
}

func (p *Pagination) SetLimit(limit int) {
	switch {
	case limit > 100:
		limit = 100
	case limit <= 0:
		limit = 10
	}

	p.Limit = limit
}

func (p *Pagination) SetPage(page int) {
	p.Page = page
}

func (p *Pagination) SetTotalRows(total_rows int64) {
	p.TotalRows = total_rows
}

func (p *Pagination) SetTotalPages() {
	p.TotalPages = p.GetTotalPage()
}

func (p *Pagination) GetLimit() int {
	if p.Limit == 0 {
		p.Limit = 10
	}
	return p.Limit
}

func (p *Pagination) GetPage() int {
	if p.Page == 0 {
		p.Page = 1
	}
	return p.Page
}

func (p *Pagination) GetTotalRows() int64 {
	return p.TotalRows
}

func (p *Pagination) GetOffset() int {
	return (p.GetPage() - 1) * p.GetLimit()
}

func (p *Pagination) GetTotalPage() int {
	return int(math.Ceil(float64(p.GetTotalRows()) / float64(p.GetLimit())))
}
