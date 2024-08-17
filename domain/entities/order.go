package entities

type Order struct {
	BaseEntity
	AuditableEntity
	userId   uint64
	statusId uint64
	status   OrderStatus
}
