package __communication

import __styles "sergio/unit-testing/1_styles"

type IEmailGateway interface {
	SendReceipt(email string, product __styles.Product, quantity __styles.Quantity)
}
