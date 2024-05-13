package wgrest

type InfoBody struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type BasicAuthFields struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type NewAccountBody struct {
	Auth    BasicAuthFields `json:"auth"`
	Account struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Role     string `json:"role"`
	}
}

type InterfaceBody struct {
	Auth      BasicAuthFields `json:"auth"`
	Interface struct {
		Name        string `json:"name"`
		Address     string `json:"address"`
		Description string `json:"description"`
	} `json:"interface"`
}
