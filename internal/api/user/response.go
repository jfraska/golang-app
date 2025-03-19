package user

type OauthUserResponse struct {
	Id            string `json:"id"`
	Name          string `json:"name"`
	Email         string `json:"email"`
	Image         string `json:"picture"`
	EmailVerified bool   `json:"verified_email"`
	State         string
}
