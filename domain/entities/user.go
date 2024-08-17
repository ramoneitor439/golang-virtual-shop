package entities

import "golang.org/x/crypto/bcrypt"

type User struct {
	BaseEntity
	AuditableEntity
	Email        string
	PasswordHash string
	FirstName    string
	LastName     string
	IsActive     bool
	Roles        []*Role
}

func (user *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	return err == nil
}

func (user *User) SetPassword(password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.PasswordHash = string(hashedPassword)

	return nil
}

func (user *User) AddRoles(roles ...*Role) {
	user.Roles = append(user.Roles, roles...)
}
