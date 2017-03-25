package models

import (
	"encoding/json"
	"time"

	"github.com/markbates/pop"
	"github.com/markbates/validate"
	"github.com/markbates/validate/validators"
	"github.com/satori/go.uuid"
)

type User struct {
	ID         uuid.UUID `json:"-" db:"id"`
	CreatedAt  time.Time `json:"-" db:"created_at"`
	UpdatedAt  time.Time `json:"-" db:"updated_at"`
	BloodGroup string    `json:"b" db:"blood_grp"`
	Dept       string    `json:"d" db:"department"`
	Gender     string    `json:"g" db:"gender"`
	Hall       string    `json:"h" db:"hall"`
	Hometown   string    `json:"a" db:"hometown"`
	Name       string    `json:"n" db:"name"`
	Program    string    `json:"p" db:"program"`
	Roll       string    `json:"i" db:"roll"`
	Room       string    `json:"r" db:"room"`
	Username   string    `json:"u" db:"username"`
}

// String is not required by pop and may be deleted
func (u User) String() string {
	ju, _ := json.Marshal(u)
	return string(ju)
}

// Users is not required by pop and may be deleted
type Users []User

// String is not required by pop and may be deleted
func (u Users) String() string {
	ju, _ := json.Marshal(u)
	return string(ju)
}

// Validate gets run everytime you call a "pop.Validate" method.
// This method is not required and may be deleted.
func (u *User) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.StringIsPresent{Field: u.Username, Name: "Username"},
	), nil
}

// ValidateSave gets run everytime you call "pop.ValidateSave" method.
// This method is not required and may be deleted.
func (u *User) ValidateSave(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run everytime you call "pop.ValidateUpdate" method.
// This method is not required and may be deleted.
func (u *User) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
