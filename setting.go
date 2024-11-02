package woocommerce

import (
	"github.com/ViktorLi1988/woocommerce-go/entity"
	jsoniter "github.com/json-iterator/go"
)

type settingService service

func (s settingService) Groups() (items []entity.SettingGroup, err error) {
	resp, err := s.httpClient.R().Get("/settings")
	if err != nil {
		return
	}

	if resp.IsSuccess() {
		err = jsoniter.Unmarshal(resp.Body(), &items)
	}
	return
}
