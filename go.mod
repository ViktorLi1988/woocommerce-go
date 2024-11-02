module github.com/ViktorLi1988/woocommerce-go

go 1.21

toolchain go1.21.6

require (
	github.com/araddon/dateparse v0.0.0-20210429162001-6b43995a97de
	github.com/brianvoe/gofakeit/v6 v6.28.0
	github.com/go-ozzo/ozzo-validation/v4 v4.3.0
	github.com/go-resty/resty/v2 v2.15.3
	github.com/google/go-querystring v1.1.0
	github.com/hiscaler/gox v0.0.0-20240824093629-11ab6f7acd0b
	github.com/hiscaler/woocommerce-go v1.0.3
	github.com/json-iterator/go v1.1.12
	github.com/stretchr/testify v1.9.0
)

require (
	github.com/asaskevich/govalidator v0.0.0-20230301143203-a9d515a09cc2 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	golang.org/x/net v0.30.0 // indirect
	golang.org/x/text v0.19.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/hiscaler/woocommerce-go => ../woocommerce-go
