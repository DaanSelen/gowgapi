package wgrest

type infoBody struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type interfaceBody struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	Interface struct {
		State       string `json:"state"`
		Name        string `json:"name"`
		Description string `json:"description"`
	} `json:"interface"`
}
