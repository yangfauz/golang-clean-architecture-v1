package user

import "time"

type User struct {
	ID        int
	Name      string
	Role      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
