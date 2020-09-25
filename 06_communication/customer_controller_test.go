package _6_communication

import (
	"github.com/golang/mock/gomock"
	"github.com/onsi/gomega"
	__styles "sergio/unit-testing/01_styles"
	"sergio/unit-testing/06_communication/mocks"
	"testing"
)

func Test_Successful_purchase(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mock := mocks.NewMockIEmailGateway(ctrl)
	sut := NewCustomerController(mock)

	mock.EXPECT().
		SendReceipt("sergioabreu@ntopus.com.br", __styles.Shampoo, __styles.Quantity(5)).
		Times(1)

	isSuccess := sut.Purchase(1, 2, 5)

	g.Expect(isSuccess).Should(
		gomega.BeTrue())
}
