package wgrest

type InfoBody struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type BasicFields struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type NewAccountBody struct {
	Auth    BasicFields `json:"auth"`
	Account struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Role     string `json:"role"`
	}
}

type InterfaceBody struct {
	Auth      BasicFields `json:"auth"`
	Interface struct {
		Name        string `json:"name"`
		Description string `json:"description"`
	} `json:"interface"`
}
