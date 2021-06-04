package validate

type Password struct {
	Password string `json:"password" validate:"required,min=6,max=120" label:"密码"`
}
