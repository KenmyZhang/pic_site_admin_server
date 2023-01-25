package sdk

import (
	"encoding/json"
	"github.com/ttlv/kdniao"
	"github.com/ttlv/kdniao/enum"
	"github.com/ttlv/kdniao/request"
	"github.com/ttlv/kdniao/response"
	"github.com/ttlv/kdniao/util"
	"github.com/ttlv/kdniao/util/http"
	"strconv"
)

func NewApiRecognise(config kdniao.KdniaoConfig, logger kdniao.KdniaoLoggerInterface) ApiRecognise {
	return ApiRecognise{config, logger}
}

type ApiRecognise struct {
	config kdniao.KdniaoConfig
	logger kdniao.KdniaoLoggerInterface
}

func (obj ApiRecognise) GetRequest(logisticCode string) request.RecogniseRequest {
	req := request.NewRecogniseRequest()
	req.SetConfig(obj.config)
	req.SetLogisticCode(logisticCode)
	return req
}

func (obj ApiRecognise) GetResponse(req request.RecogniseRequest) (response.RecogniseResponse, error) {
	url := enum.GATEWAY + enum.URI_BUSINESS

	req.UpdateRequestData()
	var resp response.RecogniseResponse
	httpResp, err := http.HttpPostForm(url, req.ToValues(), obj.logger)
	if err != nil {
		return resp, util.ErrorWrap(err, "ApiRecognise,http fail")
	} else if !httpResp.IsOk() {
		return resp, util.ErrorNew("ApiRecognise,code:" + strconv.Itoa(httpResp.GetCode()))
	}
	err = json.Unmarshal(httpResp.GetBytes(), &resp)
	if err != nil {
		return resp, util.ErrorWrap(err, "ApiRecognise,json decode fail")
	}
	if !resp.IsSuccess() {
		return resp, util.ErrorNew("ApiRecognise," + resp.GetError())
	}
	return resp, nil
}
