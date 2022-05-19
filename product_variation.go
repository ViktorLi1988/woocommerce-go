package woocommerce

import (
	"errors"
	"fmt"
	"github.com/google/go-querystring/query"
	"github.com/hiscaler/woocommerce-go/entity"
	jsoniter "github.com/json-iterator/go"
)

type productVariationService service

// Product variations

type ProductVariationsQueryParams struct {
	queryParams
	Search string `json:"search,omitempty"`
}

func (m ProductVariationsQueryParams) Validate() error {
	return nil
}

// All List all product variations
func (s productVariationService) All(productId int, params ProductVariationsQueryParams) (items []entity.ProductVariation, isLastPage bool, err error) {
	if err = params.Validate(); err != nil {
		return
	}

	params.TidyVars()
	urlValues, _ := query.Values(params)
	var res []entity.ProductVariation
	resp, err := s.httpClient.R().SetQueryParamsFromValues(urlValues).Get(fmt.Sprintf("/products/%d/variations", productId))
	if err != nil {
		return
	}

	if resp.IsSuccess() {
		if err = jsoniter.Unmarshal(resp.Body(), &res); err == nil {
			items = res
		}
	} else {
		err = ErrorWrap(resp.StatusCode(), "")
	}
	return
}

// One retrieve a product variation
func (s productVariationService) One(productId, variationId int) (item entity.ProductVariation, err error) {
	resp, err := s.httpClient.R().Get(fmt.Sprintf("/products/%d/variations/%d", productId, variationId))
	if err != nil {
		return
	}

	if resp.IsSuccess() {
		err = jsoniter.Unmarshal(resp.Body(), &item)
	} else {
		err = ErrorWrap(resp.StatusCode(), "")
	}
	return
}

// Create

type CreateProductVariationRequest struct {
	Description    string                             `json:"description,omitempty"`
	SKU            string                             `json:"sku,omitempty"`
	RegularPrice   string                             `json:"regular_price,omitempty"`
	SalePrice      string                             `json:"sale_price,omitempty"`
	Status         string                             `json:"status,omitempty"`
	Virtual        bool                               `json:"virtual,omitempty"`
	Downloadable   bool                               `json:"downloadable,omitempty"`
	Downloads      []entity.ProductDownload           `json:"downloads,omitempty"`
	DownloadLimit  int                                `json:"download_limit,omitempty"`
	DownloadExpiry int                                `json:"download_expiry,omitempty"`
	TaxStatus      string                             `json:"tax_status,omitempty"`
	TaxClass       string                             `json:"tax_class,omitempty"`
	ManageStock    bool                               `json:"manage_stock,omitempty"`
	StockQuantity  int                                `json:"stock_quantity,omitempty"`
	StockStatus    string                             `json:"stock_status,omitempty"`
	Backorders     string                             `json:"backorders,omitempty"`
	Weight         string                             `json:"weight,omitempty"`
	Dimension      *entity.ProductDimension           `json:"dimensions,omitempty"`
	ShippingClass  string                             `json:"shipping_class,omitempty"`
	Image          *entity.ProductImage               `json:"image,omitempty"`
	Attributes     []entity.ProductVariationAttribute `json:"attributes,omitempty"`
	MenuOrder      int                                `json:"menu_order,omitempty"`
	MetaData       []entity.Meta                      `json:"meta_data,omitempty"`
}

func (m CreateProductVariationRequest) Validate() error {
	return nil
}

func (s productVariationService) Create(productId int, req CreateProductVariationRequest) (item entity.ProductVariation, err error) {
	if err = req.Validate(); err != nil {
		return
	}

	resp, err := s.httpClient.R().
		SetBody(req).
		Post(fmt.Sprintf("/products/%d/variations", productId))
	if err != nil {
		return
	}

	if resp.IsSuccess() {
		err = jsoniter.Unmarshal(resp.Body(), &item)
	} else {
		err = ErrorWrap(resp.StatusCode(), "")
	}
	return
}

// Update

type UpdateProductVariationRequest = CreateProductVariationRequest

func (s productVariationService) Update(productId int, req UpdateProductVariationRequest) (item entity.ProductVariation, err error) {
	if err = req.Validate(); err != nil {
		return
	}

	resp, err := s.httpClient.R().
		SetBody(req).
		Put(fmt.Sprintf("/products/%d/variations", productId))
	if err != nil {
		return
	}

	if resp.IsSuccess() {
		err = jsoniter.Unmarshal(resp.Body(), &item)
	} else {
		err = ErrorWrap(resp.StatusCode(), "")
	}
	return
}

// Delete

func (s productVariationService) Delete(productId, variationId int, force bool) (item entity.ProductVariation, err error) {
	resp, err := s.httpClient.R().
		SetBody(map[string]bool{"force": force}).
		Delete(fmt.Sprintf("/products/%d/variations/%d", productId, variationId))
	if err != nil {
		return
	}

	if resp.IsSuccess() {
		err = jsoniter.Unmarshal(resp.Body(), &item)
	} else {
		err = ErrorWrap(resp.StatusCode(), "")
	}
	return
}

// Batch Update

type BatchProductVariationsCreateItem = CreateProductVariationRequest
type BatchProductVariationsUpdateItem struct {
	ID int `json:"id"`
	CreateProductVariationRequest
}

type BatchProductVariationsRequest struct {
	Create []BatchProductVariationsCreateItem `json:"create,omitempty"`
	Update []BatchProductVariationsUpdateItem `json:"update,omitempty"`
	Delete []int                              `json:"delete,omitempty"`
}

func (m BatchProductVariationsRequest) Validate() error {
	if len(m.Create) == 0 && len(m.Update) == 0 && len(m.Delete) == 0 {
		return errors.New("无效的请求数据")
	}
	return nil
}

type BatchProductVariationsResult struct {
	Create []entity.ProductVariation `json:"create"`
	Update []entity.ProductVariation `json:"update"`
	Delete []entity.ProductVariation `json:"delete"`
}

func (s productVariationService) Batch(req BatchProductVariationsRequest) (res BatchProductVariationsResult, err error) {
	if err = req.Validate(); err != nil {
		return
	}

	resp, err := s.httpClient.R().SetBody(req).Post("/products/variations/batch")
	if err != nil {
		return
	}

	if resp.IsSuccess() {
		err = jsoniter.Unmarshal(resp.Body(), &res)
	}
	return
}
