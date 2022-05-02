package application_test

import (
	"testing"

	"github.com/guilhermereis14/go-hexagonal/application"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

func Test_givenProduct_whenCallEnable_thenShouldReturnNil(t *testing.T) {
	product := application.Product{}
	product.Price = 10.00
	err := product.Enable()
	require.Nil(t, err)
}
func Test_givenProduct_whenCallEnableAndPriceEqualsZero_thenShouldReturnErrorMessage(t *testing.T) {
	product := application.Product{}
	product.Price = 0.00
	err := product.Enable()
	require.Equal(t, "The price must be greater than zero to enable the product", err.Error())
}
func Test_givenProduct_whenCallDisable_thenShouldReturnNil(t *testing.T) {
	product := application.Product{}
	product.Price = 0.00
	err := product.Disable()
	require.Nil(t, err)
}
func Test_givenProduct_whenCallDisableAndPriceNotEqualsZero_thenShouldReturnErrorMessage(t *testing.T) {
	product := application.Product{}
	product.Price = 10.00
	err := product.Disable()
	require.Equal(t, "The price must be zero in order to have the product disabled", err.Error())
}

func Test_givenProduct_whenCallIsValidFunction_thenShouldValidateRulesProduct(t *testing.T) {
	product := application.Product{}
	product.ID = uuid.NewV4().String()
	product.Name = "hello"
	product.Status = application.DISABLED
	product.Price = 10

	_, err := product.IsValid()
	require.Nil(t, err)

	product.Status = "INVALID"
	_, err = product.IsValid()
	require.Equal(t, "the status must be enabled or desabled", err.Error())

	product.Status = application.ENABLED
	_, err = product.IsValid()
	require.Nil(t, err)

	product.Price = -10
	_, err = product.IsValid()
	require.Equal(t, "the price must be greater or equal zero", err.Error())
}
