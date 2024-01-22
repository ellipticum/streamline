package models

import "time"

type User struct {
	ID          int64    `db:"id" json:"id"`
	Username    string   `db:"username" json:"username"`
	Email       string   `db:"email" json:"email"`
	FirstName   string   `db:"first_name" json:"firstName"`
	LastName    string   `db:"last_name" json:"lastName"`
	Description string   `db:"description" json:"description"`
	Contacts    []string `db:"contacts" json:"contacts"`

	DateOfBirth     time.Time `db:"date_of_birth" json:"dateOfBirth"`
	SubscriptionEnd time.Time `db:"subscription_end" json:"subscriptionEnd"`
}
