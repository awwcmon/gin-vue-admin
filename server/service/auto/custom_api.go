package auto

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/auto"
	autoReq "github.com/flipped-aurora/gin-vue-admin/server/model/auto/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"io"
	"net/http"
)

type CustomApiService struct {
}

// CreateCustomApi 创建CustomApi记录
// Author [piexlmax](https://github.com/piexlmax)
func (customApiService *CustomApiService) CreateCustomApi(customApi *auto.CustomApi) (err error) {
	err = global.GVA_DB.Create(customApi).Error
	return err
}

// DeleteCustomApi 删除CustomApi记录
// Author [piexlmax](https://github.com/piexlmax)
func (customApiService *CustomApiService) DeleteCustomApi(ID string) (err error) {
	err = global.GVA_DB.Delete(&auto.CustomApi{}, "id = ?", ID).Error
	return err
}

// DeleteCustomApiByIds 批量删除CustomApi记录
// Author [piexlmax](https://github.com/piexlmax)
func (customApiService *CustomApiService) DeleteCustomApiByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]auto.CustomApi{}, "id in ?", IDs).Error
	return err
}

// UpdateCustomApi 更新CustomApi记录
// Author [piexlmax](https://github.com/piexlmax)
func (customApiService *CustomApiService) UpdateCustomApi(customApi auto.CustomApi) (err error) {
	err = global.GVA_DB.Save(&customApi).Error
	return err
}

// GetCustomApi 根据ID获取CustomApi记录
// Author [piexlmax](https://github.com/piexlmax)
func (customApiService *CustomApiService) GetCustomApi(ID string) (customApi auto.CustomApi, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&customApi).Error
	return
}

// GetCustomApiInfoList 分页获取CustomApi记录
// Author [piexlmax](https://github.com/piexlmax)
func (customApiService *CustomApiService) GetCustomApiInfoList(info autoReq.CustomApiSearch) (list []auto.CustomApi, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&auto.CustomApi{})
	var customApis []auto.CustomApi
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.ApiName != "" {
		db = db.Where("api_name = ?", info.ApiName)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	var OrderStr string
	orderMap := make(map[string]bool)
	orderMap["api_name"] = true
	if orderMap[info.Sort] {
		OrderStr = info.Sort
		if info.Order == "descending" {
			OrderStr = OrderStr + " desc"
		}
		db = db.Order(OrderStr)
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&customApis).Error
	return customApis, total, err
}

func (customApiService *CustomApiService) Execute(customApi *auto.CustomApi, token string) (resp response.Response, err error) {
	//return response.Response{}
	var body map[string]interface{}
	var bodybytes []byte
	if customApi.ParamConfig != "" {
		err := json.Unmarshal([]byte(customApi.ParamConfig), &body)
		if err != nil {

			return response.Response{}, err
		}
		bodybytes, err = json.Marshal(body)
		if err != nil {
			return response.Response{}, err
		}
	}
	url := global.GVA_CONFIG.System.Domain + customApi.Api
	fmt.Println(url)
	req, err := http.NewRequest(customApi.Method, url, bytes.NewReader(bodybytes))
	if err != nil {
		return response.Response{}, err
	}
	req.Header.Add("x-token", token)
	req.Header.Add("content-type", "application/json")
	respd, err := http.DefaultClient.Do(req)
	if err != nil {
		return response.Response{}, err
	}
	defer respd.Body.Close()
	bodyBytes, _ := io.ReadAll(respd.Body)
	var res response.Response
	err = json.Unmarshal(bodyBytes, &res)
	if err != nil {
		return response.Response{}, err
	}
	return res, err
}
