package vo

type UserResponse struct {
	Username string          `json:"username"`
	Email    string          `json:"email"`
	Profile  ProfileResponse `json:"profile"`
}

type ProfileResponse struct {
	Name      string `json:"name"`
	Bio       string `json:"bio"`
	AvatarUrl string `json:"avatar_url"`
}
