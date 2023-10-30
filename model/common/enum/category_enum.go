package enum

type CategoryLevel int8

const (
	Default    CategoryLevel = 0
	LevelOne   CategoryLevel = 1
	LevelTwo   CategoryLevel = 2
	LevelThree CategoryLevel = 3
)

func (g CategoryLevel) Info() (int, string) {
	switch g {
	case LevelOne:
		return 1, "First level classification"
	case LevelTwo:
		return 2, "Second level classification"
	case LevelThree:
		return 3, "Third level classification"
	default:
		return 0, "error"
	}
}

func (g CategoryLevel) Code() int {
	switch g {
	case LevelOne:
		return 1
	case LevelTwo:
		return 2
	case LevelThree:
		return 3
	default:
		return 0
	}
}
