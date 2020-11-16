package entity

type Follow struct {
	BaseEntity
	Following    User
	FollowingID  string `sql:"type:uuid"`
	FollowedBy   User
	FollowedByID string `sql:"type:uuid"`
}
