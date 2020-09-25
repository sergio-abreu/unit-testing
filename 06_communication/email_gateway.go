package _6_communication

import __styles "sergio/unit-testing/01_styles"

type IEmailGateway interface {
	SendReceipt(email string, product __styles.Product, quantity __styles.Quantity)
}
