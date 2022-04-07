package application_test

import (
	"testing"

	"github.com/guilhermereis14/go-hexagonal/application"
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
