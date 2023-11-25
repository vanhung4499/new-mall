package common

const (
	sortColumnDefault    = "id"
	sortDirectionDefault = "ASC"
)

type Sorting struct {
	SortColumn    string `json:"sort_column,omitempty" form:"sort_column" binding:"omitempty"`
	SortDirection string `json:"sort_direction,omitempty" form:"sort_direction" binding:"omitempty"`
}

func (s *Sorting) FillDefault() {
	if s.SortColumn == "" {
		s.SortColumn = sortColumnDefault
	}

	if s.SortDirection == "" {
		s.SortDirection = sortDirectionDefault
	}
}
