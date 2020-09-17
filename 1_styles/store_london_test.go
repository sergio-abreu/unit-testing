package __styles_test

import (
	"github.com/golang/mock/gomock"
	. "github.com/onsi/gomega"
	. "sergio/unit-testing/1_styles"
	"sergio/unit-testing/1_styles/mocks"
	"testing"
)

func Test_London_Purchase_succeeds_when_enough_inventory(t *testing.T) {
	// Arrange
	g := NewGomegaWithT(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	store := mocks.NewMockIStore(ctrl)
	store.EXPECT().
		HasEnoughInventory(gomock.Any(), gomock.Any()).
		Return(true)
	customer := NewCustomer()

	store.EXPECT().
		RemoveInventory(Shampoo, Quantity(5)).
		Times(1)

	// Act
	err := customer.Purchase(store, Shampoo, 5)

	// Assert
	g.Expect(err).Should(
		Not(HaveOccurred()))
}

func Test_London_Purchase_fails_when_not_enough_inventory(t *testing.T) {
	// Arrange
	g := NewGomegaWithT(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	store := mocks.NewMockIStore(ctrl)
	store.EXPECT().
		HasEnoughInventory(gomock.Any(), gomock.Any()).
		Return(false)
	store.EXPECT().
		GetInventory(gomock.Any()).
		Return(Quantity(4))
	customer := NewCustomer()

	// Act
	err := customer.Purchase(store, Shampoo, 5)

	// Assert
	g.Expect(err).Should(
		MatchError("not enough inventory for product 0: current inventory is 4, but desired inventory is 5"))
}
