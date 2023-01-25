package request

import (
	`encoding/json`

	"github.com/ttlv/kdniao/enum"
)

func NewExpressQueryRequest() ExpressQueryRequest {
	req := ExpressQueryRequest{}
	req.SetRequestType(enum.REQUEST_TYPE_EXPRESS_REQUEST)
	req.SetDataType(enum.DATA_TYPE_JSON)
	return req
}

type ExpressQueryRequest struct {
	KdniaoRequest
	OrderCode    string `json:"OrderCode"`    // 订单编号
	ShipperCode  string `json:"ShipperCode"`  // 快递公司编码
	CustomerName string `json:"CustomerName"` // ShipperCode 为 SF 时必填，对应寄件人 /收件人手机号后四位；ShipperCode 为其他快递时，可不填或保 留字段，不可传值
	Sort         int    `json:"Sort"`         // 轨迹排序，0-升序，1-降序，默认 0-升序
	LogisticCode string `json:"LogisticCode"` // 物流单号
}

func (req *ExpressQueryRequest) SetLogisticCode(logisticCode string) *ExpressQueryRequest {
	req.LogisticCode = logisticCode
	return req
}

func (req ExpressQueryRequest) GetLogisticCode() string {
	return req.LogisticCode
}

func (req *ExpressQueryRequest) UpdateRequestData() *ExpressQueryRequest {
	req.requestData = req.ToJson()
	return req
}

func (req ExpressQueryRequest) ToJson() string {
	str, err := json.Marshal(req)
	if err != nil {
		return ""
	}
	return string(str)
}