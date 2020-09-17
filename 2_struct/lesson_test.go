package __struct

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "sergio/unit-testing/1_styles"
	"sergio/unit-testing/Assert"
	"testing"
	"time"
)

//Padrão AAA
func Test_Sum_of_two_numbers(t *testing.T) {
	// Arrange
	first := 10.0
	second := 20.0
	var calculator = Calculator{}

	// Act
	result := calculator.Sum(first, second)

	// Assert
	Assert.Equal(30.0, result)
}

// Um conceito por teste
func Test_Add_months(t *testing.T) {
	// Arrange
	date_1 := time.Date(2004, 5, 31, 0, 0, 0, 0, time.UTC)

	// Act
	date_2 := date_1.AddDate(0, 1, 0)

	// Assert
	Assert.Equal(1, date_2.Day())
	Assert.Equal(time.Month(7), date_2.Month())
	Assert.Equal(2004, date_2.Year())

	// Arrange
	// Act
	date_3 := date_1.AddDate(0, 2, 0)

	// Assert
	Assert.Equal(31, date_3.Day())
	Assert.Equal(time.Month(7), date_3.Month())
	Assert.Equal(2004, date_3.Year())

	// Arrange
	// Act
	date_4 := date_1.AddDate(0, 1, 0).AddDate(0, 1, 0)

	// Assert
	Assert.Equal(1, date_4.Day())
	Assert.Equal(time.Month(8), date_4.Month())
	Assert.Equal(2004, date_4.Year())
}

// Uma linha na seção act
func Test_Classical_Purchase_succeeds_when_enough_inventory(t *testing.T) {
	// Arrange
	g := NewGomegaWithT(t)
	store := NewStore()
	store.AddInventory(Shampoo, 10)
	customer := NewCustomer2()

	// Act
	err := customer.Purchase(store, Shampoo, 5)
	_ = store.RemoveInventory(Shampoo, 5)

	// Assert
	g.Expect(err).Should(
		Not(HaveOccurred()))
	g.Expect(store.GetInventory(Shampoo)).Should(
		BeEquivalentTo(5))
}

// Remover os comentários, identificando o que está sendo testado
func Test___Sum_of_two_numbers(t *testing.T) {
	first := 10.0
	second := 20.0
	var sut = Calculator{}

	result := sut.Sum(first, second)

	Assert.Equal(30.0, result)
}

// Reutilizando arrange section
func Test_Calculator_Sum(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Calculator Suite")
}

var _ = Describe("Calculator suite", func() {
	var store Store
	var sut Customer

	Context("Using before each", func() {
		BeforeEach(func() {
			store = NewStore()
			store.AddInventory(Shampoo, 10)
			sut = NewCustomer()
		})

		It("Purchase succeeds when enough inventory", func() {
			err := sut.Purchase(store, Shampoo, 5)

			Expect(err).Should(
				Not(HaveOccurred()))
			Expect(store.GetInventory(Shampoo)).Should(
				BeEquivalentTo(5))
		})

		It("Purchase fails when not enough inventory", func() {
			err := sut.Purchase(store, Shampoo, 15)

			Expect(err).Should(
				MatchError("not enough inventory for product 0: current inventory is 10, but desired inventory is 15"))
			Expect(store.GetInventory(Shampoo)).Should(
				BeEquivalentTo(10))
		})
	})

	Context("Not using before each", func() {

		It("Purchase succeeds when enough inventory", func() {
			store := CreateStoreWithInventory(Shampoo, 10)
			sut := CreateCustomer()

			err := sut.Purchase(store, Shampoo, 5)

			Expect(err).Should(
				Not(HaveOccurred()))
			Expect(store.GetInventory(Shampoo)).Should(
				BeEquivalentTo(5))
		})

		It("Purchase fails when not enough inventory", func() {
			store := CreateStoreWithInventory(Shampoo, 10)
			sut := CreateCustomer()

			err := sut.Purchase(store, Shampoo, 15)

			Expect(err).Should(
				MatchError("not enough inventory for product 0: current inventory is 10, but desired inventory is 15"))
			Expect(store.GetInventory(Shampoo)).Should(
				BeEquivalentTo(10))
		})

		BeforeEach(func() {
			// inicializar conexão com banco de dados
		})

		AfterEach(func() {
			// fechar conexão com o banco de dados
		})

		It("Purchase succeeds when enough inventory", func() {
			// uso do banco de dados
		})
	})

})

func CreateStoreWithInventory(product Product, quantity Quantity) Store {
	store := NewStore()
	store.AddInventory(product, quantity)
	return store
}

func CreateCustomer() Customer {
	return NewCustomer()
}

// Nomeando o teste
// Existe alguns padrões rígidos de nomeclatura: Um exemplo de padrão
// 		segue a linha [MethodUnderTest]_[Scenario]_[ExpectedResult]
func Test__Sum_TwoNumbers_ReturnsSum(t *testing.T) {}
func Test__Sum_of_two_numbers(t *testing.T)        {}

// Nome com padrão rígido
func Test_IsDeliveryValid_InvalidDate_ReturnsFalse(t *testing.T) {
	var sut DeliveryService
	pastDate := time.Now().AddDate(0, 0, -1)
	delivery := Delivery{Date: pastDate}

	isValid := sut.IsDeliveryValid(delivery)

	Assert.Equal(false, isValid)
}

// Refatorando para um nome mais expressivo
func Test_Delivery_with_invalid_date_should_be_considered_invalid(t *testing.T) {}

// Removendo conceitos genéricos
func Test_Delivery_with_past_date_should_be_considered_invalid(t *testing.T) {}

// Removendo palavras desnecessárias
func Test_Delivery_with_past_date_should_be_invalid(t *testing.T) {}

// Afirmando o que o teste irá fazer
func Test_Delivery_with_past_date_is_invalid(t *testing.T) {}

// Adicionando artigos e palavras para dar melhor sentido na sentença
func Test_Delivery_with_a_past_date_is_invalid(t *testing.T)           {}
func Test_Delivery_for_today_is_invalid(t *testing.T)                  {}
func Test_Delivery_for_tomorrow_is_invalid(t *testing.T)               {}
func Test_The_soonest_delivery_date_is_two_days_from_now(t *testing.T) {}

type testCase struct {
	deliveryDate time.Time
	expected     bool
}

func Test_Can_detect_an_invalid_delivery_date(t *testing.T) {
	testCases := []testCase{
		{time.Now().AddDate(0, 0, -1), false},
		{time.Now().AddDate(0, 0, 0), false},
		{time.Now().AddDate(0, 0, 1), false},
		{time.Now().AddDate(0, 0, 2), true},
	}
	for _, _case := range testCases {
		t.Run("Can detect an invalid delivery date", func(t *testing.T) {
			sut := DeliveryService{}
			delivery := Delivery{Date: _case.deliveryDate}

			isValid := sut.IsDeliveryValid(delivery)

			Assert.Equal(_case.expected, isValid)
		})
	}
}

func Test__Detects_an_invalid_delivery_date(t *testing.T) {
	for range GetDeliveryDates() {
		t.Run("Detects an invalid delivery date", func(t *testing.T) {
			/* ... */
		})
	}
}

func Test__The_soonest_delivery_date_is_two_days_from_now(t *testing.T) {}

func GetDeliveryDates() []time.Time {
	deliveryDates := []time.Time{
		time.Now().AddDate(0, 0, -1),
		time.Now().AddDate(0, 0, 0),
		time.Now().AddDate(0, 0, 1),
	}
	return deliveryDates
}

// Fluent Assertion
func Test_Sum_of_two_numbers_(t *testing.T) {
	g := NewGomegaWithT(t)
	first := 10.0
	second := 20.0
	var sut = Calculator{}

	result := sut.Sum(first, second)

	g.Expect(result).Should(
		BeEquivalentTo(30.0))
}
