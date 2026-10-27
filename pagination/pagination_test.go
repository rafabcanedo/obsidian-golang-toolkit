package pagination_test

import (
	"testing"

	"github.com/rafabcanedo/obsidian-golang-toolkit/pagination"
	"github.com/rafabcanedo/obsidian-golang-toolkit/testutil"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name     string
		page     int
		pageSize int
		wantPage int
		wantSize int
	}{
		{"normal values", 2, 20, 2, 20},
		{"page less than 1", 0, 10, 1, 10},
		{"pageSize less than 1", 1, 0, 1, 10},
		{"pageSize greater than 100", 1, 200, 1, 100},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := pagination.New(tt.page, tt.pageSize)
			testutil.Equal(t, p.Page, tt.wantPage)
			testutil.Equal(t, p.PageSize, tt.wantSize)
		})
	}
}

func TestOffset(t *testing.T) {
	tests := []struct {
		name     string
		page     int
		pageSize int
		want     int
	}{
		{"page 1", 1, 10, 0},
		{"page 2", 2, 10, 10},
		{"page 3 with pageSize 5", 3, 5, 10},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := pagination.New(tt.page, tt.pageSize)
			testutil.Equal(t, p.Offset(), tt.want)
		})
	}
}

func TestNewMeta(t *testing.T) {
	tests := []struct {
		name           string
		total          int
		pageSize       int
		wantTotalPages int
	}{
		{"25 items pageSize 10", 25, 10, 3},
		{"20 items pageSize 10", 20, 10, 2},
		{"0 items", 0, 10, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := pagination.New(1, tt.pageSize)
			meta := p.NewMeta(tt.total)
			testutil.Equal(t, meta.TotalPages, tt.wantTotalPages)
			testutil.Equal(t, meta.Total, tt.total)
			testutil.Equal(t, meta.CurrentPage, 1)
			testutil.Equal(t, meta.PageSize, tt.pageSize)
		})
	}
}
