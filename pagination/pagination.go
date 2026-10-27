package pagination

type Params struct {
    Page     int
    PageSize int
}

type Meta struct {
    Total       int `json:"total"`
    TotalPages  int `json:"total_pages"`
    CurrentPage int `json:"current_page"`
    PageSize    int `json:"page_size"`
}

func New(page, pageSize int) Params {

	if page < 1 {
		page = 1
	}

	if pageSize < 1 {
		pageSize = 10
	}

	if pageSize > 100 {
		pageSize = 100
	}

	return Params{Page: page, PageSize: pageSize}
}

func (p Params) Offset() int {
	return (p.Page - 1) * p.PageSize
}

func (p Params) NewMeta(total int) Meta {
	totalPages := (total + p.PageSize - 1) / p.PageSize

	return Meta{
		Total:       total,
        TotalPages:  totalPages,
        CurrentPage: p.Page,
        PageSize:    p.PageSize,
	}
}
