package meta

import (
	"fmt"
	"os"
	"strconv"
)

type Meta struct {
	Page       int `json:"page"`
	PerPage    int `json:"per_page"`
	PageCount  int `json:"page_count"`
	TotalCount int `json:"total_count"`
}

// explicar por que poner New en lugar de NewMeta
func New(page, perPage, total int) (*Meta, error) {

	// si perPage es menor o igual a cero, le meto el valor por defecto
	if perPage <= 0 {
		var err error
		perPage, err = strconv.Atoi(os.Getenv("PAGINATOR_LIMIT_DEFAULT"))
		if err != nil {
			return nil, err
		}

	}

	pageCount := 0
	if total >= 0 {
		// total 75, per page 25
		// sin el -1 me va a mostrar 4 paginas en lugar de 3
		pageCount = (total + perPage - 1) / perPage
		if page > pageCount {
			page = pageCount
		}
	}

	if page < 1 {
		page = 1
	}

	return &Meta{
		Page:       page,
		PerPage:    perPage,
		TotalCount: total,
		PageCount:  pageCount,
	}, nil
}

func (p *Meta) Offset() int {
	fmt.Println(p)
	return (p.Page - 1) * p.PerPage
}

func (p *Meta) Limit() int {
	return p.PerPage
}
