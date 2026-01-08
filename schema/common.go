package schema

import "github.com/google/uuid"

type Pagination struct {
	TotalItems  int64 `json:"total_items"`
	TotalPages  int   `json:"total_pages"`
	Limit       int   `json:"limit"`
	Page        int   `json:"page"`
	HasNextPage bool  `json:"has_next_page"`
	HasPrevPage bool  `json:"has_prev_page"`
}

type RegularIDs uint
type LargeIDs uuid.UUID
