package placeorder

import (
	"context"
	"errors"
	"math/rand"
	"strconv"
	"sync"

	authDomain "flamingo.me/flamingo/v3/core/oauth/domain"

	cartDomain "flamingo.me/flamingo-commerce/v3/cart/domain/cart"
	"flamingo.me/flamingo-commerce/v3/cart/domain/placeorder"
)

// AttributeErrorKey is used to store a forced test error to the cart's additional attributes
const AttributeErrorKey = "test-error"

type (
	// FakeAdapter provides fake place order adapter
	FakeAdapter struct {
		locker       sync.Locker
		placedOrders map[string]placeorder.PlacedOrderInfos
	}
)

var (
	_ placeorder.Service = new(FakeAdapter)
	// NextCancelFails can be set to let the next call to any FakeAdapter's instance fail
	NextCancelFails bool
)

// Inject dependencies
func (f *FakeAdapter) Inject() *FakeAdapter {
	f.placedOrders = make(map[string]placeorder.PlacedOrderInfos)
	f.locker = &sync.Mutex{}
	return f
}

// PlaceGuestCart places a guest cart order
func (f *FakeAdapter) PlaceGuestCart(ctx context.Context, cart *cartDomain.Cart, payment *placeorder.Payment) (placeorder.PlacedOrderInfos, error) {
	return f.placeCart(cart)
}

// PlaceCustomerCart places a customer cart
func (f *FakeAdapter) PlaceCustomerCart(ctx context.Context, auth authDomain.Auth, cart *cartDomain.Cart, payment *placeorder.Payment) (placeorder.PlacedOrderInfos, error) {
	return f.placeCart(cart)
}

func (f *FakeAdapter) placeCart(cart *cartDomain.Cart) (placeorder.PlacedOrderInfos, error) {
	f.locker.Lock()
	defer f.locker.Unlock()
	forcedError := cart.AdditionalData.CustomAttributes[AttributeErrorKey]
	if forcedError != "" {
		return nil, errors.New(forcedError)
	}

	reservedID := cart.AdditionalData.ReservedOrderID

	_, found := f.placedOrders[reservedID]

	if found {
		return nil, errors.New("Another order with #" + reservedID + " already placed")
	}

	var placedOrders placeorder.PlacedOrderInfos
	placedOrders = append(placedOrders, placeorder.PlacedOrderInfo{
		OrderNumber: reservedID,
	})

	f.placedOrders[reservedID] = placedOrders

	return placedOrders, nil
}

// ReserveOrderID returns the reserved order id
func (f *FakeAdapter) ReserveOrderID(_ context.Context, cart *cartDomain.Cart) (string, error) {
	forcedError := cart.AdditionalData.CustomAttributes[AttributeErrorKey]
	if forcedError != "" {
		return "", errors.New(forcedError)
	}
	return cart.ID + "-" + strconv.Itoa(rand.Int()), nil
}

// CancelGuestOrder cancels a guest order
func (f *FakeAdapter) CancelGuestOrder(ctx context.Context, orderInfos placeorder.PlacedOrderInfos) error {
	return f.cancelOrder(orderInfos)
}

// CancelCustomerOrder cancels a customer order
func (f *FakeAdapter) CancelCustomerOrder(ctx context.Context, orderInfos placeorder.PlacedOrderInfos, auth authDomain.Auth) error {
	return f.cancelOrder(orderInfos)
}

func (f *FakeAdapter) cancelOrder(orderInfos placeorder.PlacedOrderInfos) error {
	f.locker.Lock()
	defer f.locker.Unlock()
	if NextCancelFails {
		NextCancelFails = false
		return errors.New("test")
	}

	var toDelete []string
	for _, order := range orderInfos {
		_, found := f.placedOrders[order.OrderNumber]

		if !found {
			return errors.New("Order cancel not possible order #" + order.OrderNumber + " wasn't placed")
		}

		toDelete = append(toDelete, order.OrderNumber)
	}

	for _, orderNumber := range toDelete {
		delete(f.placedOrders, orderNumber)
	}

	return nil
}
