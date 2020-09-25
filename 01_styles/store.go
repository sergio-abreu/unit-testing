package _1_styles

type (
	Product   int
	Quantity  uint
	Inventory map[Product]Quantity
	IStore    interface {
		AddInventory(product Product, addedQuantity Quantity)
		HasEnoughInventory(product Product, desiredQuantity Quantity) bool
		RemoveInventory(product Product, quantity Quantity) error
		GetInventory(product Product) Quantity
	}
)

const (
	Shampoo Product = iota
	Book
)

func NewStore() Store {
	return Store{inventory: Inventory{}}
}

type Store struct {
	inventory Inventory
}

func (store Store) AddInventory(product Product, addedQuantity Quantity) {
	store.inventory[product] += addedQuantity
}

func (store Store) HasEnoughInventory(product Product, desiredQuantity Quantity) bool {
	currentQuantity := store.inventory[product]
	return currentQuantity >= desiredQuantity
}

func (store Store) RemoveInventory(product Product, desiredQuantity Quantity) error {
	if !store.HasEnoughInventory(product, desiredQuantity) {
		return NewNotEnoughInventoryErr(Shampoo, store.GetInventory(product), desiredQuantity)
	}
	store.inventory[product] -= desiredQuantity
	return nil
}

func (store Store) GetInventory(product Product) Quantity {
	return store.inventory[product]
}
