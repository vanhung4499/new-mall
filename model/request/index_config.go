package request

import (
	"new-mall/model"
	"new-mall/model/common"
)

type IndexConfigSearch struct {
	model.IndexConfig
	common.PageInfo
}

type IndexConfigAddParams struct {
	ConfigName  string `json:"configName"`
	ConfigType  int    `json:"configType"`
	ProductId   string `json:"productId"`
	RedirectUrl string `json:"redirectUrl"`
	ConfigRank  string `json:"configRank"`
}

type IndexConfigUpdateParams struct {
	ConfigId    int    `json:"configId"`
	ConfigName  string `json:"configName"`
	RedirectUrl string `json:"redirectUrl"`
	ConfigType  int    `json:"configType"`
	ProductId   int    `json:"productId"`
	ConfigRank  string `json:"configRank"`
}
