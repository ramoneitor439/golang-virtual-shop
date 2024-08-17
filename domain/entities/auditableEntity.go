package entities

import "time"

type AuditableEntity struct {
	CreatedAt time.Time
	UpdatedAt *time.Time
}
