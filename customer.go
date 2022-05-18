package woocommerce

import (
	"fmt"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/google/go-querystring/query"
	"github.com/hiscaler/woocommerce-go/entity"
	"github.com/hiscaler/woocommerce-go/entity/customer"
	jsoniter "github.com/json-iterator/go"
)

type CustomerService service

type CustomersQueryParams struct {
	QueryParams
	Search  string `url:"search,omitempty"`
	Exclude []int  `url:"exclude,omitempty"`
	Include []int  `url:"include,omitempty"`
	Email   string `url:"email,omitempty"`
	Role    string `url:"role,omitempty"`
}

func (m CustomersQueryParams) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.OrderBy, validation.When(m.OrderBy != "", validation.In("id", "include", "name", "registered_date").Error("无效的排序字段"))),
		validation.Field(&m.Email, validation.When(m.Email != "", is.EmailFormat.Error("无效的邮箱"))),
		validation.Field(&m.Role, validation.When(m.Role != "", validation.In("all", "administrator", "editor", "author", "contributor", "subscriber", "shop_manager").Error("无效的角色"))),
	)
}

func (s CustomerService) All(params CustomersQueryParams) (items []customer.Customer, isLastPage bool, err error) {
	if err = params.Validate(); err != nil {
		return
	}

	params.TidyVars()
	urlValues, _ := query.Values(params)
	var res []customer.Customer
	resp, err := s.httpClient.R().SetQueryParamsFromValues(urlValues).Get("/customers")
	if err != nil {
		return
	}

	if resp.IsSuccess() {
		if err = jsoniter.Unmarshal(resp.Body(), &res); err == nil {
			items = res
		}
	}
	return
}

func (s CustomerService) One(id int) (item customer.Customer, err error) {
	resp, err := s.httpClient.R().Get(fmt.Sprintf("/customers/%d", id))
	if err != nil {
		return
	}

	if resp.IsSuccess() {
		err = jsoniter.Unmarshal(resp.Body(), &item)
	}
	return
}

// CreateCustomerRequest Create customer request
type CreateCustomerRequest struct {
	Email     string            `json:"email,omitempty"`
	FirstName string            `json:"first_name,omitempty"`
	LastName  string            `json:"last_name,omitempty"`
	Username  string            `json:"username,omitempty"`
	Password  string            `json:"password,omitempty"`
	Billing   *entity.Billing   `json:"billing,omitempty"`
	Shipping  *entity.Shipping  `json:"shipping,omitempty"`
	MetaData  []entity.MetaData `json:"meta_data,omitempty"`
}

func (m CreateCustomerRequest) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.Email, validation.Required.Error("邮箱不能为空"), is.EmailFormat.Error("无效的邮箱")),
		validation.Field(&m.FirstName, validation.Required.Error("姓不能为空")),
		validation.Field(&m.LastName, validation.Required.Error("名不能为空")),
		validation.Field(&m.Username, validation.Required.Error("登录帐号不能为空")),
		validation.Field(&m.Password, validation.Required.Error("登录密码不能为空")),
		validation.Field(&m.Billing),
	)
}

func (s CustomerService) Create(req CreateCustomerRequest) (item customer.Customer, err error) {
	if err = req.Validate(); err != nil {
		return
	}

	resp, err := s.httpClient.R().SetBody(req).Post("/customers")
	if err != nil {
		return
	}

	if resp.IsSuccess() {
		err = jsoniter.Unmarshal(resp.Body(), &item)
	}
	return
}

// Update customer

type UpdateCustomerRequest struct {
	Email     string            `json:"email,omitempty"`
	FirstName string            `json:"first_name,omitempty"`
	LastName  string            `json:"last_name,omitempty"`
	Billing   *entity.Billing   `json:"billing,omitempty"`
	Shipping  *entity.Shipping  `json:"shipping,omitempty"`
	MetaData  []entity.MetaData `json:"meta_data,omitempty"`
}

func (m UpdateCustomerRequest) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.Email, validation.When(m.Email != "", is.EmailFormat.Error("无效的邮箱"))),
		validation.Field(&m.Billing),
	)
}

func (s CustomerService) Update(id int, req UpdateCustomerRequest) (item customer.Customer, err error) {
	if err = req.Validate(); err != nil {
		return
	}

	resp, err := s.httpClient.R().SetBody(req).Put(fmt.Sprintf("/customers/%d", id))
	if err != nil {
		return
	}

	if resp.IsSuccess() {
		err = jsoniter.Unmarshal(resp.Body(), &item)
	}
	return
}

// Delete Delete a customer
func (s CustomerService) Delete(id int) (item customer.Customer, err error) {
	resp, err := s.httpClient.R().Delete(fmt.Sprintf("/customers/%d", id))
	if err != nil {
		return
	}

	if resp.IsSuccess() {
		err = jsoniter.Unmarshal(resp.Body(), &item)
	}
	return
}