package enum

type IndexConfigEnum int8

// Home page configuration items 1-Search box hot search 2-Search drop-down box hot search 3-(Home page) Hot-selling products 4-(Home page) New products online 5-(Home page) Recommended for you
const (
	IndexSearchHots       IndexConfigEnum = 1
	IndexSearchDownHots   IndexConfigEnum = 2
	IndexProductHot       IndexConfigEnum = 3
	IndexProductNew       IndexConfigEnum = 4
	IndexProductRecommend IndexConfigEnum = 5
)

func (i IndexConfigEnum) Info() (int, string) {
	switch i {
	case IndexSearchHots:
		return 1, "INDEX_SEARCH_HOTS"
	case IndexSearchDownHots:
		return 2, "Second lvel classification"
	case IndexProductHot:
		return 3, "Third level classification"
	case IndexProductNew:
		return 4, "Third level classification"
	case IndexProductRecommend:
		return 5, "Third level classification"
	default:
		return 0, "DEFAULT"
	}
}

func (i IndexConfigEnum) Code() int {
	switch i {
	case IndexSearchHots:
		return 1
	case IndexSearchDownHots:
		return 2
	case IndexProductHot:
		return 3
	case IndexProductNew:
		return 4
	case IndexProductRecommend:
		return 5
	default:
		return 0
	}
}
