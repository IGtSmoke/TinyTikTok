package dto

type Follow struct {
	ID         int64
	UserID     int64
	FollowerID int64
	Cancel     int8
}
