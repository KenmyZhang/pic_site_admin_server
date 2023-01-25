package sdk

import (
	`encoding/json`
	`strconv`

	`github.com/ttlv/kdniao`
	`github.com/ttlv/kdniao/enum`
	`github.com/ttlv/kdniao/request`
	`github.com/ttlv/kdniao/response`
	`github.com/ttlv/kdniao/util`
	`github.com/ttlv/kdniao/util/http`
)

func NewExpressQuery(config kdniao.KdniaoConfig, logger kdniao.KdniaoLoggerInterface) ApiExpressQuery {
	return ApiExpressQuery{config, logger}
}

type ApiExpressQuery struct {
	config kdniao.KdniaoConfig
	logger kdniao.KdniaoLoggerInterface
}

func (obj ApiExpressQuery) GetRequest(logisticCode string) request.ExpressQueryRequest {
	req := request.NewExpressQueryRequest()
	req.SetConfig(obj.config)
	req.SetLogisticCode(logisticCode)
	return req
}

func (obj ApiExpressQuery) GetResponse(req request.ExpressQueryRequest) (response.ExpressQueryResponse, error) {
	url := enum.GATEWAY + enum.URI_BUSINESS

	req.UpdateRequestData()
	var resp response.ExpressQueryResponse
	httpResp, err := http.HttpPostForm(url, req.ToValues(), obj.logger)
	if err != nil {
		return resp, util.ErrorWrap(err, "ApiExpressQuery,http fail")
	} else if !httpResp.IsOk() {
		return resp, util.ErrorNew("ApiExpressQuery,code:" + strconv.Itoa(httpResp.GetCode()))
	}
	err = json.Unmarshal(httpResp.GetBytes(), &resp)
	if err != nil {
		return resp, util.ErrorWrap(err, "ApiExpressQuery,json decode fail")
	}
	if !resp.IsSuccess() {
		return resp, util.ErrorNew("ApiExpressQuery," + resp.GetError())
	}
	return resp, nil
}
