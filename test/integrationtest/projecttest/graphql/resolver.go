// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

//+build !graphql

package graphql

import (
	"context"

	"flamingo.me/flamingo-commerce/v3/cart/domain/cart"
	"flamingo.me/flamingo-commerce/v3/cart/domain/validation"
	"flamingo.me/flamingo-commerce/v3/cart/interfaces/controller/forms"
	graphql1 "flamingo.me/flamingo-commerce/v3/cart/interfaces/graphql"
	"flamingo.me/flamingo-commerce/v3/cart/interfaces/graphql/dto"
	domain1 "flamingo.me/flamingo-commerce/v3/category/domain"
	graphql7 "flamingo.me/flamingo-commerce/v3/category/interfaces/graphql"
	"flamingo.me/flamingo-commerce/v3/category/interfaces/graphql/categorydto"
	graphql4 "flamingo.me/flamingo-commerce/v3/checkout/interfaces/graphql"
	dto1 "flamingo.me/flamingo-commerce/v3/checkout/interfaces/graphql/dto"
	graphql6 "flamingo.me/flamingo-commerce/v3/customer/interfaces/graphql"
	"flamingo.me/flamingo-commerce/v3/customer/interfaces/graphql/dtocustomer"
	graphql5 "flamingo.me/flamingo-commerce/v3/product/interfaces/graphql"
	graphqlproductdto "flamingo.me/flamingo-commerce/v3/product/interfaces/graphql/product/dto"
	"flamingo.me/flamingo-commerce/v3/search/domain"
	graphql2 "flamingo.me/flamingo-commerce/v3/search/interfaces/graphql"
	"flamingo.me/flamingo-commerce/v3/search/interfaces/graphql/searchdto"
	graphql3 "flamingo.me/graphql"
)

var _ ResolverRoot = new(rootResolver)

type rootResolver struct {
	rootResolverCommerce_Cart                         *rootResolverCommerce_Cart
	rootResolverCommerce_CartItem                     *rootResolverCommerce_CartItem
	rootResolverCommerce_CartShippingItem             *rootResolverCommerce_CartShippingItem
	rootResolverCommerce_Cart_DefaultPaymentSelection *rootResolverCommerce_Cart_DefaultPaymentSelection
	rootResolverCommerce_Search_Meta                  *rootResolverCommerce_Search_Meta
	rootResolverMutation                              *rootResolverMutation
	rootResolverQuery                                 *rootResolverQuery
}

func (r *rootResolver) Inject(
	rootResolverCommerce_Cart *rootResolverCommerce_Cart,
	rootResolverCommerce_CartItem *rootResolverCommerce_CartItem,
	rootResolverCommerce_CartShippingItem *rootResolverCommerce_CartShippingItem,
	rootResolverCommerce_Cart_DefaultPaymentSelection *rootResolverCommerce_Cart_DefaultPaymentSelection,
	rootResolverCommerce_Search_Meta *rootResolverCommerce_Search_Meta,
	rootResolverMutation *rootResolverMutation,
	rootResolverQuery *rootResolverQuery,
) {
	r.rootResolverCommerce_Cart = rootResolverCommerce_Cart
	r.rootResolverCommerce_CartItem = rootResolverCommerce_CartItem
	r.rootResolverCommerce_CartShippingItem = rootResolverCommerce_CartShippingItem
	r.rootResolverCommerce_Cart_DefaultPaymentSelection = rootResolverCommerce_Cart_DefaultPaymentSelection
	r.rootResolverCommerce_Search_Meta = rootResolverCommerce_Search_Meta
	r.rootResolverMutation = rootResolverMutation
	r.rootResolverQuery = rootResolverQuery
}

func (r *rootResolver) Commerce_Cart() Commerce_CartResolver {
	return r.rootResolverCommerce_Cart
}
func (r *rootResolver) Commerce_CartItem() Commerce_CartItemResolver {
	return r.rootResolverCommerce_CartItem
}
func (r *rootResolver) Commerce_CartShippingItem() Commerce_CartShippingItemResolver {
	return r.rootResolverCommerce_CartShippingItem
}
func (r *rootResolver) Commerce_Cart_DefaultPaymentSelection() Commerce_Cart_DefaultPaymentSelectionResolver {
	return r.rootResolverCommerce_Cart_DefaultPaymentSelection
}
func (r *rootResolver) Commerce_Search_Meta() Commerce_Search_MetaResolver {
	return r.rootResolverCommerce_Search_Meta
}
func (r *rootResolver) Mutation() MutationResolver {
	return r.rootResolverMutation
}
func (r *rootResolver) Query() QueryResolver {
	return r.rootResolverQuery
}

type rootResolverCommerce_Cart struct {
	resolveGetDeliveryByCode func(ctx context.Context, obj *cart.Cart, deliveryCode string) (*cart.Delivery, error)
}

func (r *rootResolverCommerce_Cart) Inject(
	commerce_CartGetDeliveryByCode *graphql1.Resolver,
) {
	r.resolveGetDeliveryByCode = commerce_CartGetDeliveryByCode.GetDeliveryByCodeWithoutBool
}

func (r *rootResolverCommerce_Cart) GetDeliveryByCode(ctx context.Context, obj *cart.Cart, deliveryCode string) (*cart.Delivery, error) {
	return r.resolveGetDeliveryByCode(ctx, obj, deliveryCode)
}

type rootResolverCommerce_CartItem struct {
	resolveAppliedDiscounts func(ctx context.Context, obj *cart.Item) (*dto.CartAppliedDiscounts, error)
}

func (r *rootResolverCommerce_CartItem) Inject(
	commerce_CartItemAppliedDiscounts *dto.CartAppliedDiscountsResolver,
) {
	r.resolveAppliedDiscounts = commerce_CartItemAppliedDiscounts.ForItem
}

func (r *rootResolverCommerce_CartItem) AppliedDiscounts(ctx context.Context, obj *cart.Item) (*dto.CartAppliedDiscounts, error) {
	return r.resolveAppliedDiscounts(ctx, obj)
}

type rootResolverCommerce_CartShippingItem struct {
	resolveAppliedDiscounts func(ctx context.Context, obj *cart.ShippingItem) (*dto.CartAppliedDiscounts, error)
}

func (r *rootResolverCommerce_CartShippingItem) Inject(
	commerce_CartShippingItemAppliedDiscounts *dto.CartAppliedDiscountsResolver,
) {
	r.resolveAppliedDiscounts = commerce_CartShippingItemAppliedDiscounts.ForShippingItem
}

func (r *rootResolverCommerce_CartShippingItem) AppliedDiscounts(ctx context.Context, obj *cart.ShippingItem) (*dto.CartAppliedDiscounts, error) {
	return r.resolveAppliedDiscounts(ctx, obj)
}

type rootResolverCommerce_Cart_DefaultPaymentSelection struct {
	resolveCartSplit func(ctx context.Context, obj *cart.DefaultPaymentSelection) ([]*dto.PaymentSelectionSplit, error)
}

func (r *rootResolverCommerce_Cart_DefaultPaymentSelection) Inject(
	commerce_Cart_DefaultPaymentSelectionCartSplit *graphql1.CommerceCartQueryResolver,
) {
	r.resolveCartSplit = commerce_Cart_DefaultPaymentSelectionCartSplit.CartSplit
}

func (r *rootResolverCommerce_Cart_DefaultPaymentSelection) CartSplit(ctx context.Context, obj *cart.DefaultPaymentSelection) ([]*dto.PaymentSelectionSplit, error) {
	return r.resolveCartSplit(ctx, obj)
}

type rootResolverCommerce_Search_Meta struct {
	resolveSortOptions func(ctx context.Context, obj *domain.SearchMeta) ([]*searchdto.CommerceSearchSortOption, error)
}

func (r *rootResolverCommerce_Search_Meta) Inject(
	commerce_Search_MetaSortOptions *graphql2.CommerceSearchQueryResolver,
) {
	r.resolveSortOptions = commerce_Search_MetaSortOptions.SortOptions
}

func (r *rootResolverCommerce_Search_Meta) SortOptions(ctx context.Context, obj *domain.SearchMeta) ([]*searchdto.CommerceSearchSortOption, error) {
	return r.resolveSortOptions(ctx, obj)
}

type rootResolverMutation struct {
	resolveFlamingo                                  func(ctx context.Context) (*string, error)
	resolveCommerceAddToCart                         func(ctx context.Context, marketplaceCode string, qty int, deliveryCode string) (*dto.DecoratedCart, error)
	resolveCommerceDeleteCartDelivery                func(ctx context.Context, deliveryCode string) (*dto.DecoratedCart, error)
	resolveCommerceDeleteItem                        func(ctx context.Context, itemID string, deliveryCode string) (*dto.DecoratedCart, error)
	resolveCommerceUpdateItemQty                     func(ctx context.Context, itemID string, deliveryCode string, qty int) (*dto.DecoratedCart, error)
	resolveCommerceCartUpdateBillingAddress          func(ctx context.Context, addressForm *forms.AddressForm) (*dto.BillingAddressForm, error)
	resolveCommerceCartUpdateSelectedPayment         func(ctx context.Context, gateway string, method string) (*dto.SelectedPaymentResult, error)
	resolveCommerceCartApplyCouponCodeOrGiftCard     func(ctx context.Context, code string) (*dto.DecoratedCart, error)
	resolveCommerceCartRemoveGiftCard                func(ctx context.Context, giftCardCode string) (*dto.DecoratedCart, error)
	resolveCommerceCartRemoveCouponCode              func(ctx context.Context, couponCode string) (*dto.DecoratedCart, error)
	resolveCommerceCartUpdateDeliveryAddresses       func(ctx context.Context, deliveryAdresses []*forms.DeliveryForm) ([]*dto.DeliveryAddressForm, error)
	resolveCommerceCartUpdateDeliveryShippingOptions func(ctx context.Context, shippingOptions []*dto.DeliveryShippingOption) ([]*dto.DeliveryAddressForm, error)
	resolveCommerceCartClean                         func(ctx context.Context) (bool, error)
	resolveCommerceCheckoutStartPlaceOrder           func(ctx context.Context, returnURL string) (*dto1.StartPlaceOrderResult, error)
	resolveCommerceCheckoutCancelPlaceOrder          func(ctx context.Context) (bool, error)
	resolveCommerceCheckoutClearPlaceOrder           func(ctx context.Context) (bool, error)
	resolveCommerceCheckoutRefreshPlaceOrder         func(ctx context.Context) (*dto1.PlaceOrderContext, error)
	resolveCommerceCheckoutRefreshPlaceOrderBlocking func(ctx context.Context) (*dto1.PlaceOrderContext, error)
}

func (r *rootResolverMutation) Inject(
	mutationFlamingo *graphql3.FlamingoQueryResolver,
	mutationCommerceAddToCart *graphql1.CommerceCartMutationResolver,
	mutationCommerceDeleteCartDelivery *graphql1.CommerceCartMutationResolver,
	mutationCommerceDeleteItem *graphql1.CommerceCartMutationResolver,
	mutationCommerceUpdateItemQty *graphql1.CommerceCartMutationResolver,
	mutationCommerceCartUpdateBillingAddress *graphql1.CommerceCartMutationResolver,
	mutationCommerceCartUpdateSelectedPayment *graphql1.CommerceCartMutationResolver,
	mutationCommerceCartApplyCouponCodeOrGiftCard *graphql1.CommerceCartMutationResolver,
	mutationCommerceCartRemoveGiftCard *graphql1.CommerceCartMutationResolver,
	mutationCommerceCartRemoveCouponCode *graphql1.CommerceCartMutationResolver,
	mutationCommerceCartUpdateDeliveryAddresses *graphql1.CommerceCartMutationResolver,
	mutationCommerceCartUpdateDeliveryShippingOptions *graphql1.CommerceCartMutationResolver,
	mutationCommerceCartClean *graphql1.CommerceCartMutationResolver,
	mutationCommerceCheckoutStartPlaceOrder *graphql4.CommerceCheckoutMutationResolver,
	mutationCommerceCheckoutCancelPlaceOrder *graphql4.CommerceCheckoutMutationResolver,
	mutationCommerceCheckoutClearPlaceOrder *graphql4.CommerceCheckoutMutationResolver,
	mutationCommerceCheckoutRefreshPlaceOrder *graphql4.CommerceCheckoutMutationResolver,
	mutationCommerceCheckoutRefreshPlaceOrderBlocking *graphql4.CommerceCheckoutMutationResolver,
) {
	r.resolveFlamingo = mutationFlamingo.Flamingo
	r.resolveCommerceAddToCart = mutationCommerceAddToCart.CommerceAddToCart
	r.resolveCommerceDeleteCartDelivery = mutationCommerceDeleteCartDelivery.CommerceDeleteCartDelivery
	r.resolveCommerceDeleteItem = mutationCommerceDeleteItem.CommerceDeleteItem
	r.resolveCommerceUpdateItemQty = mutationCommerceUpdateItemQty.CommerceUpdateItemQty
	r.resolveCommerceCartUpdateBillingAddress = mutationCommerceCartUpdateBillingAddress.CommerceCartUpdateBillingAddress
	r.resolveCommerceCartUpdateSelectedPayment = mutationCommerceCartUpdateSelectedPayment.CommerceCartUpdateSelectedPayment
	r.resolveCommerceCartApplyCouponCodeOrGiftCard = mutationCommerceCartApplyCouponCodeOrGiftCard.CommerceCartApplyCouponCodeOrGiftCard
	r.resolveCommerceCartRemoveGiftCard = mutationCommerceCartRemoveGiftCard.CommerceCartRemoveGiftCard
	r.resolveCommerceCartRemoveCouponCode = mutationCommerceCartRemoveCouponCode.CommerceCartRemoveCouponCode
	r.resolveCommerceCartUpdateDeliveryAddresses = mutationCommerceCartUpdateDeliveryAddresses.CommerceCartUpdateDeliveryAddresses
	r.resolveCommerceCartUpdateDeliveryShippingOptions = mutationCommerceCartUpdateDeliveryShippingOptions.CommerceCartUpdateDeliveryShippingOptions
	r.resolveCommerceCartClean = mutationCommerceCartClean.CartClean
	r.resolveCommerceCheckoutStartPlaceOrder = mutationCommerceCheckoutStartPlaceOrder.CommerceCheckoutStartPlaceOrder
	r.resolveCommerceCheckoutCancelPlaceOrder = mutationCommerceCheckoutCancelPlaceOrder.CommerceCheckoutCancelPlaceOrder
	r.resolveCommerceCheckoutClearPlaceOrder = mutationCommerceCheckoutClearPlaceOrder.CommerceCheckoutClearPlaceOrder
	r.resolveCommerceCheckoutRefreshPlaceOrder = mutationCommerceCheckoutRefreshPlaceOrder.CommerceCheckoutRefreshPlaceOrder
	r.resolveCommerceCheckoutRefreshPlaceOrderBlocking = mutationCommerceCheckoutRefreshPlaceOrderBlocking.CommerceCheckoutRefreshPlaceOrderBlocking
}

func (r *rootResolverMutation) Flamingo(ctx context.Context) (*string, error) {
	return r.resolveFlamingo(ctx)
}
func (r *rootResolverMutation) CommerceAddToCart(ctx context.Context, marketplaceCode string, qty int, deliveryCode string) (*dto.DecoratedCart, error) {
	return r.resolveCommerceAddToCart(ctx, marketplaceCode, qty, deliveryCode)
}
func (r *rootResolverMutation) CommerceDeleteCartDelivery(ctx context.Context, deliveryCode string) (*dto.DecoratedCart, error) {
	return r.resolveCommerceDeleteCartDelivery(ctx, deliveryCode)
}
func (r *rootResolverMutation) CommerceDeleteItem(ctx context.Context, itemID string, deliveryCode string) (*dto.DecoratedCart, error) {
	return r.resolveCommerceDeleteItem(ctx, itemID, deliveryCode)
}
func (r *rootResolverMutation) CommerceUpdateItemQty(ctx context.Context, itemID string, deliveryCode string, qty int) (*dto.DecoratedCart, error) {
	return r.resolveCommerceUpdateItemQty(ctx, itemID, deliveryCode, qty)
}
func (r *rootResolverMutation) CommerceCartUpdateBillingAddress(ctx context.Context, addressForm *forms.AddressForm) (*dto.BillingAddressForm, error) {
	return r.resolveCommerceCartUpdateBillingAddress(ctx, addressForm)
}
func (r *rootResolverMutation) CommerceCartUpdateSelectedPayment(ctx context.Context, gateway string, method string) (*dto.SelectedPaymentResult, error) {
	return r.resolveCommerceCartUpdateSelectedPayment(ctx, gateway, method)
}
func (r *rootResolverMutation) CommerceCartApplyCouponCodeOrGiftCard(ctx context.Context, code string) (*dto.DecoratedCart, error) {
	return r.resolveCommerceCartApplyCouponCodeOrGiftCard(ctx, code)
}
func (r *rootResolverMutation) CommerceCartRemoveGiftCard(ctx context.Context, giftCardCode string) (*dto.DecoratedCart, error) {
	return r.resolveCommerceCartRemoveGiftCard(ctx, giftCardCode)
}
func (r *rootResolverMutation) CommerceCartRemoveCouponCode(ctx context.Context, couponCode string) (*dto.DecoratedCart, error) {
	return r.resolveCommerceCartRemoveCouponCode(ctx, couponCode)
}
func (r *rootResolverMutation) CommerceCartUpdateDeliveryAddresses(ctx context.Context, deliveryAdresses []*forms.DeliveryForm) ([]*dto.DeliveryAddressForm, error) {
	return r.resolveCommerceCartUpdateDeliveryAddresses(ctx, deliveryAdresses)
}
func (r *rootResolverMutation) CommerceCartUpdateDeliveryShippingOptions(ctx context.Context, shippingOptions []*dto.DeliveryShippingOption) ([]*dto.DeliveryAddressForm, error) {
	return r.resolveCommerceCartUpdateDeliveryShippingOptions(ctx, shippingOptions)
}
func (r *rootResolverMutation) CommerceCartClean(ctx context.Context) (bool, error) {
	return r.resolveCommerceCartClean(ctx)
}
func (r *rootResolverMutation) CommerceCheckoutStartPlaceOrder(ctx context.Context, returnURL string) (*dto1.StartPlaceOrderResult, error) {
	return r.resolveCommerceCheckoutStartPlaceOrder(ctx, returnURL)
}
func (r *rootResolverMutation) CommerceCheckoutCancelPlaceOrder(ctx context.Context) (bool, error) {
	return r.resolveCommerceCheckoutCancelPlaceOrder(ctx)
}
func (r *rootResolverMutation) CommerceCheckoutClearPlaceOrder(ctx context.Context) (bool, error) {
	return r.resolveCommerceCheckoutClearPlaceOrder(ctx)
}
func (r *rootResolverMutation) CommerceCheckoutRefreshPlaceOrder(ctx context.Context) (*dto1.PlaceOrderContext, error) {
	return r.resolveCommerceCheckoutRefreshPlaceOrder(ctx)
}
func (r *rootResolverMutation) CommerceCheckoutRefreshPlaceOrderBlocking(ctx context.Context) (*dto1.PlaceOrderContext, error) {
	return r.resolveCommerceCheckoutRefreshPlaceOrderBlocking(ctx)
}

type rootResolverQuery struct {
	resolveFlamingo                         func(ctx context.Context) (*string, error)
	resolveCommerceProduct                  func(ctx context.Context, marketPlaceCode string, variantMarketPlaceCode *string) (graphqlproductdto.Product, error)
	resolveCommerceProductSearch            func(ctx context.Context, searchRequest *searchdto.CommerceSearchRequest) (*graphql5.SearchResultDTO, error)
	resolveCommerceCustomerStatus           func(ctx context.Context) (*dtocustomer.CustomerStatusResult, error)
	resolveCommerceCustomer                 func(ctx context.Context) (*dtocustomer.CustomerResult, error)
	resolveCommerceCart                     func(ctx context.Context) (*dto.DecoratedCart, error)
	resolveCommerceCartValidator            func(ctx context.Context) (*validation.Result, error)
	resolveCommerceCartQtyRestriction       func(ctx context.Context, marketplaceCode string, variantCode *string, deliveryCode string) (*validation.RestrictionResult, error)
	resolveCommerceCheckoutActivePlaceOrder func(ctx context.Context) (bool, error)
	resolveCommerceCheckoutCurrentContext   func(ctx context.Context) (*dto1.PlaceOrderContext, error)
	resolveCommerceCategoryTree             func(ctx context.Context, activeCategoryCode string) (domain1.Tree, error)
	resolveCommerceCategory                 func(ctx context.Context, categoryCode string, categorySearchRequest *searchdto.CommerceSearchRequest) (*categorydto.CategorySearchResult, error)
}

func (r *rootResolverQuery) Inject(
	queryFlamingo *graphql3.FlamingoQueryResolver,
	queryCommerceProduct *graphql5.CommerceProductQueryResolver,
	queryCommerceProductSearch *graphql5.CommerceProductQueryResolver,
	queryCommerceCustomerStatus *graphql6.CustomerResolver,
	queryCommerceCustomer *graphql6.CustomerResolver,
	queryCommerceCart *graphql1.CommerceCartQueryResolver,
	queryCommerceCartValidator *graphql1.CommerceCartQueryResolver,
	queryCommerceCartQtyRestriction *graphql1.CommerceCartQueryResolver,
	queryCommerceCheckoutActivePlaceOrder *graphql4.CommerceCheckoutQueryResolver,
	queryCommerceCheckoutCurrentContext *graphql4.CommerceCheckoutQueryResolver,
	queryCommerceCategoryTree *graphql7.CommerceCategoryQueryResolver,
	queryCommerceCategory *graphql7.CommerceCategoryQueryResolver,
) {
	r.resolveFlamingo = queryFlamingo.Flamingo
	r.resolveCommerceProduct = queryCommerceProduct.CommerceProduct
	r.resolveCommerceProductSearch = queryCommerceProductSearch.CommerceProductSearch
	r.resolveCommerceCustomerStatus = queryCommerceCustomerStatus.CommerceCustomerStatus
	r.resolveCommerceCustomer = queryCommerceCustomer.CommerceCustomer
	r.resolveCommerceCart = queryCommerceCart.CommerceCart
	r.resolveCommerceCartValidator = queryCommerceCartValidator.CommerceCartValidator
	r.resolveCommerceCartQtyRestriction = queryCommerceCartQtyRestriction.CommerceCartQtyRestriction
	r.resolveCommerceCheckoutActivePlaceOrder = queryCommerceCheckoutActivePlaceOrder.CommerceCheckoutActivePlaceOrder
	r.resolveCommerceCheckoutCurrentContext = queryCommerceCheckoutCurrentContext.CommerceCheckoutCurrentContext
	r.resolveCommerceCategoryTree = queryCommerceCategoryTree.CommerceCategoryTree
	r.resolveCommerceCategory = queryCommerceCategory.CommerceCategory
}

func (r *rootResolverQuery) Flamingo(ctx context.Context) (*string, error) {
	return r.resolveFlamingo(ctx)
}
func (r *rootResolverQuery) CommerceProduct(ctx context.Context, marketPlaceCode string, variantMarketPlaceCode *string) (graphqlproductdto.Product, error) {
	return r.resolveCommerceProduct(ctx, marketPlaceCode, variantMarketPlaceCode)
}
func (r *rootResolverQuery) CommerceProductSearch(ctx context.Context, searchRequest *searchdto.CommerceSearchRequest) (*graphql5.SearchResultDTO, error) {
	return r.resolveCommerceProductSearch(ctx, searchRequest)
}
func (r *rootResolverQuery) CommerceCustomerStatus(ctx context.Context) (*dtocustomer.CustomerStatusResult, error) {
	return r.resolveCommerceCustomerStatus(ctx)
}
func (r *rootResolverQuery) CommerceCustomer(ctx context.Context) (*dtocustomer.CustomerResult, error) {
	return r.resolveCommerceCustomer(ctx)
}
func (r *rootResolverQuery) CommerceCart(ctx context.Context) (*dto.DecoratedCart, error) {
	return r.resolveCommerceCart(ctx)
}
func (r *rootResolverQuery) CommerceCartValidator(ctx context.Context) (*validation.Result, error) {
	return r.resolveCommerceCartValidator(ctx)
}
func (r *rootResolverQuery) CommerceCartQtyRestriction(ctx context.Context, marketplaceCode string, variantCode *string, deliveryCode string) (*validation.RestrictionResult, error) {
	return r.resolveCommerceCartQtyRestriction(ctx, marketplaceCode, variantCode, deliveryCode)
}
func (r *rootResolverQuery) CommerceCheckoutActivePlaceOrder(ctx context.Context) (bool, error) {
	return r.resolveCommerceCheckoutActivePlaceOrder(ctx)
}
func (r *rootResolverQuery) CommerceCheckoutCurrentContext(ctx context.Context) (*dto1.PlaceOrderContext, error) {
	return r.resolveCommerceCheckoutCurrentContext(ctx)
}
func (r *rootResolverQuery) CommerceCategoryTree(ctx context.Context, activeCategoryCode string) (domain1.Tree, error) {
	return r.resolveCommerceCategoryTree(ctx, activeCategoryCode)
}
func (r *rootResolverQuery) CommerceCategory(ctx context.Context, categoryCode string, categorySearchRequest *searchdto.CommerceSearchRequest) (*categorydto.CategorySearchResult, error) {
	return r.resolveCommerceCategory(ctx, categoryCode, categorySearchRequest)
}
