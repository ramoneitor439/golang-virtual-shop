package entities

type CartProduct struct {
	BaseEntity
	AuditableEntity
	productId uint64
	product   Product
	userId    uint64
	amount    uint64
}
