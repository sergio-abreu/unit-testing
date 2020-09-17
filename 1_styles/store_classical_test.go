package __styles_test

import (
	. "github.com/onsi/gomega"
	. "sergio/unit-testing/1_styles"
	"testing"
)

func Test_Classical_Purchase_succeeds_when_enough_inventory(t *testing.T) {
	// Arrange
	g := NewGomegaWithT(t)
	store := NewStore()
	store.AddInventory(Shampoo, 10)
	customer := NewCustomer()

	// Act
	err := customer.Purchase(store, Shampoo, 5)

	// Assert
	g.Expect(err).Should(
		Not(HaveOccurred()))
	g.Expect(store.GetInventory(Shampoo)).Should(
		BeEquivalentTo(5))
}

func Test_Classical_Purchase_fails_when_not_enough_inventory(t *testing.T) {
	// Arrange
	g := NewGomegaWithT(t)
	store := NewStore()
	store.AddInventory(Shampoo, 10)
	customer := NewCustomer()

	// Act
	err := customer.Purchase(store, Shampoo, 15)

	// Assert
	g.Expect(err).Should(
		MatchError("not enough inventory for product 0: current inventory is 10, but desired inventory is 15"))
	g.Expect(store.GetInventory(Shampoo)).Should(
		BeEquivalentTo(10))
}
