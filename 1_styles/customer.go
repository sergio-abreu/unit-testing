package __styles

func NewCustomer() Customer {
	return Customer{}
}

type Customer struct{}

func (customer Customer) Purchase(store IStore, product Product, quantity Quantity) error {
	if !store.HasEnoughInventory(product, quantity) {
		return NewNotEnoughInventoryErr(Shampoo, store.GetInventory(product), quantity)
	}
	return store.RemoveInventory(product, quantity)
}
