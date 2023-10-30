package request

import (
	"new-mall/model"
	"new-mall/model/common"
)

type CarouselSearch struct {
	model.Carousel
	common.PageInfo
}

type CarouselAddParam struct {
	CarouselUrl  string `json:"carouselUrl"`
	RedirectUrl  string `json:"redirectUrl"`
	CarouselRank string `json:"carouselRank"`
}

type CarouselUpdateParam struct {
	CarouselId   int    `json:"carouselId"`
	CarouselUrl  string `json:"carouselUrl"`
	RedirectUrl  string `json:"redirectUrl"`
	CarouselRank string `json:"carouselRank" `
}
