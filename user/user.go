package user

import "time"

type User struct {
	ID        uint64
	Name      string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (u User) Create() error {
	return nil
}
