package user

type UserinfoResponse struct {
	Email         string `json:"email"`
	Head          string `json:"head"`
	Name          string `json:"name"`
	Sex           uint8  `json:"sex"`
	VipLevel      uint8  `json:"vipLevel"`
	VipExpireTime string `json:"VipExpireTime"`
	Secret        string `json:"secret"`
	Status        bool   `json:"status"`
	Token         string `json:"token"`
}
