package __mocks

import (
	"github.com/golang/mock/gomock"
	. "github.com/onsi/gomega"
	__styles "sergio/unit-testing/1_styles"
	mocks2 "sergio/unit-testing/1_styles/mocks"
	"sergio/unit-testing/4_mocks/mocks"
	"testing"
)

// Mock
func Test_Sending_a_greeting_email(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mock := mocks.NewMockIEmailGateway(ctrl)
	sut := NewController(mock)

	mock.EXPECT().
		SendGreetingsEmail("user@email.com").
		Times(1)

	sut.GreetUser("user@email.com")
}

// Stub
func Test_Creating_a_report(t *testing.T) {
	g := NewGomegaWithT(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	stub := mocks.NewMockIDatabase(ctrl)
	sut := NewReportController(stub)
	stub.EXPECT().
		GetNumberOfUsers().
		Return(10, nil)

	report, err := sut.CreteReport()

	g.Expect(err).Should(
		Not(HaveOccurred()))
	g.Expect(report.NumberOfUsers).Should(
		BeEquivalentTo(10))
}

// Mock & Stubs não fazer asserts em stubs
func Test_London_Purchase_fails_when_not_enough_inventory(t *testing.T) {
	// Arrange
	g := NewGomegaWithT(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	store := mocks2.NewMockIStore(ctrl)
	store.EXPECT().
		HasEnoughInventory(gomock.Any(), gomock.Any()).
		Return(false)
	store.EXPECT().
		GetInventory(gomock.Any()).
		Return(__styles.Quantity(4))
	customer := __styles.NewCustomer()

	// Assert
	store.EXPECT().
		RemoveInventory(gomock.Any(), gomock.Any()).
		Times(0)

	// Act
	err := customer.Purchase(store, __styles.Shampoo, 5)

	// Assert
	g.Expect(err).Should(
		MatchError("not enough inventory for product 0: current inventory is 4, but desired inventory is 5"))
}
