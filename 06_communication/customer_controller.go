package _6_communication

import (
	__styles "sergio/unit-testing/01_styles"
)

type Customer struct {
	Email string
}

func (c Customer) Purchase(store __styles.Store, product __styles.Product, quantity __styles.Quantity) bool {
	if !store.HasEnoughInventory(product, quantity) {
		return false
	}
	_ = store.RemoveInventory(product, quantity)
	return true
}

func NewCustomerController(emailGateway IEmailGateway) CustomerController {
	store := __styles.NewStore()
	store.AddInventory(__styles.Shampoo, 100)
	store.AddInventory(__styles.Book, 100)
	return CustomerController{mainStore: store, emailGateway: emailGateway}
}

type CustomerController struct {
	mainStore    __styles.Store
	emailGateway IEmailGateway
}

func (c CustomerController) Purchase(customerId int, productId int, quantity __styles.Quantity) bool {
	customer := GetCustomerById(customerId)
	product := GetProductById(productId)
	isSuccess := customer.Purchase(
		c.mainStore, product, quantity)
	if isSuccess {
		c.emailGateway.SendReceipt(
			customer.Email, product, quantity)
	}
	return isSuccess
}

func GetCustomerById(customerId int) Customer {
	return Customer{Email: "sergioabreu@ntopus.com.br"}
}

func GetProductById(productId int) __styles.Product {
	return __styles.Shampoo
}
