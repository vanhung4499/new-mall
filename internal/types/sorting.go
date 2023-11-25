package types

import (
	"errors"
	"strings"
)

const (
	ASC  = "asc"
	DESC = "desc"
)

type Sorting struct {
	OrderField string `json:"order_field" form:"order_field"`
	SortType   string `json:"sort_type" form:"sort_type"`
}

func (o *Sorting) Fulfill() {
	if o.OrderField == "" {
		o.OrderField = "id"

	}

	if o.SortType == "" {

		o.SortType = DESC

	}

}

func (o Sorting) Validate() error {
	o.SortType = strings.TrimSpace(o.SortType)

	if o.SortType != ASC && o.SortType != DESC {
		return errors.New("sortType should be asc or desc")
	}

	return nil

}
