package _2_struct

import (
	__styles "sergio/unit-testing/01_styles"
)

func NewCustomer2() Customer2 {
	return Customer2{}
}

type Customer2 struct{}

func (customer Customer2) Purchase(store __styles.IStore, product __styles.Product, quantity __styles.Quantity) error {
	if !store.HasEnoughInventory(product, quantity) {
		return __styles.NewNotEnoughInventoryErr(__styles.Shampoo, store.GetInventory(product), quantity)
	}
	return nil
}
