package wgrest

type InfoBody struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type AccountBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type InterfaceBody struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	Interface struct {
		Name        string `json:"name"`
		Description string `json:"description"`
	} `json:"interface"`
}
