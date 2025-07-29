package model

import "github.com/google/uuid"

type (
	User struct {
		ID uuid.UUID
	}

	UserSource struct {
		User
		IDext  string
		Source string
	}
)
