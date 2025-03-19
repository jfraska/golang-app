package user

type RegisterRequestPayload struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

type LoginRequestPayload struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

type OauthRequestPayload struct {
	State string `form:"state" binding:"required"`
}

type OauthCallbackRequestPayload struct {
	State string `form:"state" binding:"required"`
	Code  string `form:"code" binding:"required"`
}

type OauthUserRequestPayload struct {
	Name    string `json:"name" binding:"required"`
	Email   string `json:"email" binding:"required,email"`
	Picture string `json:"picture" binding:"required"`
}
