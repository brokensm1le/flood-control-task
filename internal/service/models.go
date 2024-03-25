package service

type ResponseModel struct {
	Result bool   `json:"result"`
	Err    string `json:"err"`
}
