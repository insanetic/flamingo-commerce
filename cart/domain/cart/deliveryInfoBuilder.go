package cart

import (
	"strings"

	"flamingo.me/flamingo/framework/flamingo"
)

type (

	//DeliveryInfoBuilder - Factory
	DeliveryInfoBuilder interface {
		BuildByDeliveryCode(deliveryCode string) (DeliveryInfo, error)
		//BuildDeliveryInfoUpdateCommand(ctx web.Context, decoratedCart *DecoratedCart) ([]DeliveryInfoUpdateCommand, error)
	}

	// DefaultDeliveryInfoBuilder defines the default delivery info builder used
	DefaultDeliveryInfoBuilder struct {
		logger flamingo.Logger
	}
)

// Inject dependencies
func (b *DefaultDeliveryInfoBuilder) Inject(
	logger flamingo.Logger,
) {
	b.logger = logger
}

// BuildByDeliveryCode builds a DeliveryInfo by deliveryCode
func (b *DefaultDeliveryInfoBuilder) BuildByDeliveryCode(deliverycode string) (DeliveryInfo, error) {
	if deliverycode == "" {
		b.logger.WithField("category", "cart").WithField("subcategory", "DefaultDeliveryInfoBuilder").Warn("Empty deliverycode")
		return DeliveryInfo{
			Code:   deliverycode,
			Method: DELIVERY_METHOD_UNSPECIFIED,
		}, nil
	}
	if deliverycode == DELIVERY_METHOD_DELIVERY {
		return DeliveryInfo{
			Code:   deliverycode,
			Method: DELIVERY_METHOD_DELIVERY,
		}, nil
	}

	if deliverycode == "pickup_store" {
		return DeliveryInfo{
			Code:   deliverycode,
			Method: DELIVERY_METHOD_PICKUP,
			DeliveryLocation: DeliveryLocation{
				Type: DELIVERYLOCATION_TYPE_STORE,
			},
		}, nil
	}

	intentParts := strings.SplitN(deliverycode, "_", 3)
	if len(intentParts) != 3 {
		b.logger.WithField("category", "cart").WithField("subcategory", "DefaultDeliveryInfoBuilder").Warn("Unknown deliverycode", deliverycode)
		return DeliveryInfo{
			Code:   deliverycode,
			Method: DELIVERY_METHOD_UNSPECIFIED,
		}, nil
	}
	if intentParts[0] == DELIVERY_METHOD_PICKUP || intentParts[0] == DELIVERY_METHOD_DELIVERY {
		if intentParts[1] == DELIVERYLOCATION_TYPE_STORE {
			return DeliveryInfo{
				Code:   deliverycode,
				Method: intentParts[0],
				DeliveryLocation: DeliveryLocation{
					Code: intentParts[2],
					Type: DELIVERYLOCATION_TYPE_STORE,
				},
			}, nil
		} else if intentParts[1] == DELIVERYLOCATION_TYPE_COLLECTIONPOINT {
			return DeliveryInfo{
				Code:   deliverycode,
				Method: intentParts[0],
				DeliveryLocation: DeliveryLocation{
					Code: intentParts[2],
					Type: DELIVERYLOCATION_TYPE_COLLECTIONPOINT,
				},
			}, nil
		} else {
			return DeliveryInfo{
				Code:   deliverycode,
				Method: intentParts[0],
				DeliveryLocation: DeliveryLocation{
					Code: intentParts[2],
					Type: intentParts[1],
				},
			}, nil
		}
	}
	b.logger.WithField("category", "cart").WithField("subcategory", "DefaultDeliveryInfoBuilder").Warn("Unknown IntentString", deliverycode)
	return DeliveryInfo{
		Code:   deliverycode,
		Method: DELIVERY_METHOD_UNSPECIFIED,
	}, nil
}
