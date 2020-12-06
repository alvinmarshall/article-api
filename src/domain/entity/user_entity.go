package entity

import (
	"article_api/src/utils/encryption"
	"article_api/src/vo"
	uuid "github.com/satori/go.uuid"
)

type User struct {
	BaseEntity
	Username string `gorm:"type:varchar(30);not null;unique;"`
	Email    string `gorm:"type:varchar(100);not null;"`
	Password string `gorm:"type:varchar(200);not null;"`
	//@One to one(Profile.UserID,owner)
	Profile Profile
}

type Users []User

type Profile struct {
	Name      string `gorm:"type:varchar(100);"`
	Bio       string `gorm:"type:varchar(200);"`
	AvatarUrl string `gorm:"type:varchar(200);"`
	//@One to one (User.Profile)
	UserID uuid.UUID `sql:"type:uuid"`
}

type Follow struct {
	BaseEntity
	Following User
	//@one to one belongs to ^
	FollowingID string `sql:"type:uuid"`
	FollowedBy  User
	//@one to one belongs to ^
	FollowedByID string `sql:"type:uuid"`
}

func (user User) UserOutput() *vo.UserResponse {
	return &vo.UserResponse{
		Username: user.Username,
		Email:    user.Email,
		Profile: vo.ProfileResponse{
			Name:      user.Profile.Name,
			Bio:       user.Profile.Bio,
			AvatarUrl: user.Profile.AvatarUrl,
		},
	}
}
func (users Users) UsersOutput() []vo.UserResponse {
	var output []vo.UserResponse
	for _, data := range users {
		output = append(output, *data.UserOutput())
	}
	return output
}

func (user *User) BeforeSave() error {
	err := user.getHashPassword(user.Password)
	if err != nil {
		return err
	}
	return nil
}

func (user *User) getHashPassword(password string) error {
	hashPassword, err := encryption.HashPassword(password)
	if err != nil {
		return err
	}
	user.Password = string(hashPassword)
	return nil
}
