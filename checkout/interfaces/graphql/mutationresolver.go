package graphql

import (
	"context"
	cartApplication "flamingo.me/flamingo-commerce/v3/cart/application"
	"flamingo.me/flamingo-commerce/v3/cart/domain/decorator"
	"flamingo.me/flamingo-commerce/v3/checkout/application/placeorder"
	"flamingo.me/flamingo-commerce/v3/checkout/interfaces/graphql/dto"
	"flamingo.me/flamingo/v3/framework/flamingo"
	"flamingo.me/flamingo/v3/framework/web"
)

// CommerceCheckoutMutationResolver resolves graphql checkout mutations
type CommerceCheckoutMutationResolver struct {
	placeorderHandler    *placeorder.Handler
	cartService          *cartApplication.CartService
	logger               flamingo.Logger
	decoratedCartFactory *decorator.DecoratedCartFactory
}

// Inject dependencies
func (r *CommerceCheckoutMutationResolver) Inject(
	placeorderHandler *placeorder.Handler,
	cartService *cartApplication.CartService,
	decoratedCartFactory *decorator.DecoratedCartFactory,
	logger flamingo.Logger) {
	r.placeorderHandler = placeorderHandler
	r.decoratedCartFactory = decoratedCartFactory
	r.cartService = cartService
	r.logger = logger.WithField(flamingo.LogKeyModule, "checkout").WithField(flamingo.LogKeyCategory, "graphql")

}

//CommerceCheckoutPlaceOrderContext query
func (r *CommerceCheckoutMutationResolver) CommerceCheckoutRefreshPlaceOrder(ctx context.Context) (*dto.PlaceOrderContext, error) {

	poctx, err := r.placeorderHandler.RefreshPlaceOrder(ctx, placeorder.RefreshPlaceOrderCommand{})
	if err != nil {
		return nil, err
	}
	dc := r.decoratedCartFactory.Create(ctx, poctx.Cart)
	return &dto.PlaceOrderContext{
		Cart:       dc,
		OrderInfos: nil,
		State:      poctx.State,
		UUID:       poctx.UUID,
	}, nil
}

//CommerceCheckoutStartPlaceOrder starts a new process (if not running)
func (r *CommerceCheckoutMutationResolver) CommerceCheckoutStartPlaceOrder(ctx context.Context, returnURLRaw string) (*dto.StartPlaceOrderResult, error) {
	session := web.SessionFromContext(ctx)
	cart, err := r.cartService.GetCartReceiverService().ViewCart(ctx, session)
	if err != nil {
		return nil, err
	}
	startPlaceOrderCommand := placeorder.StartPlaceOrderCommand{Cart: *cart}
	pctx, err := r.placeorderHandler.StartPlaceOrder(ctx, startPlaceOrderCommand)
	if err == placeorder.ErrAnotherPlaceOrderProcessRunning {
		dtopctx, err := r.CommerceCheckoutRefreshPlaceOrder(ctx)
		if err != nil {
			return nil, err
		}
		return &dto.StartPlaceOrderResult{
			UUID: dtopctx.UUID,
		}, nil
	}
	if err != nil {
		return nil, err
	}
	return &dto.StartPlaceOrderResult{
		UUID: pctx.UUID,
	}, nil
}

//CommerceCheckoutCancelPlaceOrder starts a new process (if not running)
func (r *CommerceCheckoutMutationResolver) CommerceCheckoutCancelPlaceOrder(ctx context.Context) (bool, error) {
	return true, nil
}
