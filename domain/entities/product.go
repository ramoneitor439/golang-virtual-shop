package entities

type Product struct {
	BaseEntity
	AuditableEntity
	name        string
	price       float64
	description string
}
