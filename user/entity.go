package user

import "time"

type User struct {
	ID               int
	Name             string
	Occupation       string
	Email            string
	Password_hash    string
	Avatar_file_name string
	Role             string
	Token            string
	Created_at       time.Time
	Updated_at       time.Time
}
