package graphql

import (
	"flamingo.me/flamingo-commerce/v3/checkout/application"
	"flamingo.me/flamingo-commerce/v3/checkout/domain/placeorder/process"
	"flamingo.me/flamingo-commerce/v3/checkout/interfaces/graphql/dto"
	"flamingo.me/graphql"
	"github.com/99designs/gqlgen/codegen/config"
)

//go:generate go run github.com/go-bindata/go-bindata/go-bindata -nometadata -o fs.go -pkg graphql schema.graphql

// Service is the Graphql-Service of this module
type Service struct{}

// Schema returns graphql schema of this module
func (*Service) Schema() []byte {
	return MustAsset("schema.graphql")
}

// Models return the 'Schema name' => 'Go model' mapping of this module
func (*Service) Models() map[string]config.TypeMapEntry {
	return graphql.ModelMap{
		"Commerce_Checkout_PlaceOrderContext":                                            dto.PlaceOrderContext{},
		"Commerce_Checkout_StartPlaceOrder_Result":                                       dto.StartPlaceOrderResult{},
		"Commerce_Checkout_PlacedOrderInfos":                                             dto.PlacedOrderInfos{},
		"Commerce_Checkout_PlaceOrderPaymentInfo":                                        application.PlaceOrderPaymentInfo{},
		"Commerce_Checkout_PlaceOrderState_State":                                        new(dto.State),
		"Commerce_Checkout_PlaceOrderState_State_Wait":                                   dto.Wait{},
		"Commerce_Checkout_PlaceOrderState_State_WaitForCustomer":                        dto.WaitForCustomer{},
		"Commerce_Checkout_PlaceOrderState_State_Success":                                dto.Success{},
		"Commerce_Checkout_PlaceOrderState_State_Failed":                                 dto.Failed{},
		"Commerce_Checkout_PlaceOrderState_State_ShowIframe":                             dto.ShowIframe{},
		"Commerce_Checkout_PlaceOrderState_State_ShowHTML":                               dto.ShowHTML{},
		"Commerce_Checkout_PlaceOrderState_State_Redirect":                               dto.Redirect{},
		"Commerce_Checkout_PlaceOrderState_State_PostRedirect":                           dto.PostRedirect{},
		"Commerce_Checkout_PlaceOrderState_Form_Parameter":                               dto.FormParameter{},
		"Commerce_Checkout_PlaceOrderState_State_FailedReason":                           new(process.FailedReason),
		"Commerce_Checkout_PlaceOrderState_State_FailedReason_Error":                     process.ErrorOccurredReason{},
		"Commerce_Checkout_PlaceOrderState_State_FailedReason_PaymentError":              process.PaymentErrorOccurredReason{},
		"Commerce_Checkout_PlaceOrderState_State_FailedReason_CartValidationError":       process.CartValidationErrorReason{},
		"Commerce_Checkout_PlaceOrderState_State_FailedReason_CanceledByCustomer":        process.CanceledByCustomerReason{},
		"Commerce_Checkout_PlaceOrderState_State_FailedReason_PaymentCanceledByCustomer": process.PaymentCanceledByCustomerReason{},
	}.Models()
}
