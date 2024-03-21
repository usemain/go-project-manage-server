package user

type UserinfoResponse struct {
	Email  string `json:"email"`
	Head   string `json:"head"`
	Name   string `json:"name"`
	Sex    uint8  `json:"sex"`
	Level  uint8  `json:"level"`
	Status bool   `json:"status"`
	Token  string `json:"token"`
}
