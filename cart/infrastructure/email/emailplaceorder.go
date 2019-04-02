package email

import (
	"context"

	cartDomain "flamingo.me/flamingo-commerce/v3/cart/domain/cart"
	authDomain "flamingo.me/flamingo/v3/core/oauth/domain"
	"flamingo.me/flamingo/v3/framework/flamingo"
)

type (
	// PlaceOrderServiceAdapter provides an implementation of the PlaceOrderService as email adpater
	//  TODO - this example adapter need to be implemented
	PlaceOrderServiceAdapter struct {
		emailAddress string
		logger       flamingo.Logger
	}
)

var (
	_ cartDomain.PlaceOrderService = new(PlaceOrderServiceAdapter)
)

// Inject dependencies
func (e *PlaceOrderServiceAdapter) Inject(logger flamingo.Logger, config *struct {
	EmailAddress string `inject:"config:commerce.cart.emailAdapter.emailAddress"`
}) {
	e.emailAddress = config.EmailAddress
	e.logger = logger.WithField("module", "cart").WithField("category", "emailAdapter")
}

// PlaceGuestCart places a guest cart as order email
func (e *PlaceOrderServiceAdapter) PlaceGuestCart(ctx context.Context, cart *cartDomain.Cart, payment *cartDomain.Payment) (cartDomain.PlacedOrderInfos, error) {
	var placedOrders cartDomain.PlacedOrderInfos
	placedOrders = append(placedOrders, cartDomain.PlacedOrderInfo{
		OrderNumber: cart.ID,
	})

	return placedOrders, nil
}

// PlaceCustomerCart places a customer cart as order email
func (e *PlaceOrderServiceAdapter) PlaceCustomerCart(ctx context.Context, auth authDomain.Auth, cart *cartDomain.Cart, payment *cartDomain.Payment) (cartDomain.PlacedOrderInfos, error) {
	var placedOrders cartDomain.PlacedOrderInfos
	placedOrders = append(placedOrders, cartDomain.PlacedOrderInfo{
		OrderNumber: cart.ID,
	})

	return placedOrders, nil
}
