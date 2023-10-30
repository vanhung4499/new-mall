package response

// Homepage classified data VO (third level)
type ThirdLevelCategoryVO struct {
	CategoryId    int    `json:"categoryId"`
	CategoryLevel int    `json:"categoryLevel"`
	CategoryName  string `json:"categoryName" `
}

type SecondLevelCategoryVO struct {
	CategoryId            int                    `json:"categoryId"`
	ParentId              int                    `json:"parentId"`
	CategoryLevel         int                    `json:"categoryLevel"`
	CategoryName          string                 `json:"categoryName" `
	ThirdLevelCategoryVOS []ThirdLevelCategoryVO `json:"thirdLevelCategoryVOS"`
}

type MallIndexCategoryVO struct {
	CategoryId int `json:"categoryId"`
	//ParentId               int                      `json:"parentId"`
	CategoryLevel          int                     `json:"categoryLevel"`
	CategoryName           string                  `json:"categoryName" `
	SecondLevelCategoryVOS []SecondLevelCategoryVO `json:"secondLevelCategoryVOS"`
}
