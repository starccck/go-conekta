package conekta_test

import (
	"testing"

	"github.com/sait/go-conekta/conekta"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestOrder(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Handling order")
}

var _ = Describe("Handle order", func() {
	//Testing key
	conekta.ApiKey = "key_eYvWV7gSDkNYXsmr"
	Context("Create order", func() {
		It("Should response 200", func() {
			//New Order
			order := new(conekta.Order)
			item := conekta.LineItem{
				Name:      "Tacos",
				UnitPrice: 1000,
				Quantity:  12,
			}
			order.LineItems = append(order.LineItems, item)
			shipping := conekta.ShippingLine{
				Amunt:   1500,
				Carrier: "FEDEX",
			}
			order.ShippingLines = append(order.ShippingLines, shipping)
			order.Currency = "MXN"
			//testing customer id
			order.CustomerInfo.CustomerID = "cus_2fkJPFjQKABcmiZWz"
			order.ShippingContact = conekta.ShippingContact{
				Address: conekta.Address{
					Street1:    "Calle 123, int 2",
					PostalCode: "06100",
					Country:    "MX",
				},
			}
			//Adding some metadata
			order.Metadata = conekta.Metadata{
				"reference": "12987324097",
				"more_info": "lalalalala",
				"hello":     "world",
			}
			charge := conekta.Charge{
				PaymentMethod: conekta.PaymentMethod{
					Type: "default",
				},
			}
			order.Charges = append(order.Charges, charge)
			//Send to conekta
			statusCode, _ := order.Create()
			Expect(statusCode).Should(Equal(200))
		})
	})
	Context("Update order", func() {
		It("Should response 200", func() {
			order := new(conekta.Order)
			order.ID = "ord_2fw8EWJusiRrxdPzT"
			order.Currency = "MXN"
			statusCode, _ := order.Update()
			Expect(statusCode).Should(Equal(200))
		})
	})
	Context("Capture order", func() {
		It("Should response 200", func() {
		})
	})
	Context("Refound order", func() {
		It("Should response 200", func() {
		})
	})
	Context("Find order", func() {
		It("Should response 200", func() {
		})
	})
})