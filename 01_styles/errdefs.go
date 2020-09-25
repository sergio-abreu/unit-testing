package _1_styles

import "fmt"

func NewNotEnoughInventoryErr(product Product, currentQuantity, desiredQuantity Quantity) NotEnoughInventoryErr {
	return NotEnoughInventoryErr{product: product, currentQuantity: currentQuantity, desiredQuantity: desiredQuantity}
}

type NotEnoughInventoryErr struct {
	product         Product
	currentQuantity Quantity
	desiredQuantity Quantity
}

func (err NotEnoughInventoryErr) Error() string {
	return fmt.Sprintf("not enough inventory for product %d: current inventory is %d, but desired inventory is %d", err.product, err.currentQuantity, err.desiredQuantity)
}
