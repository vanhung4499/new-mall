package response

type UserDetailResponse struct {
	NickName      string `json:"nickName"`
	LoginName     string `json:"loginName"`
	IntroduceSign string `json:"introduceSign"`
}
