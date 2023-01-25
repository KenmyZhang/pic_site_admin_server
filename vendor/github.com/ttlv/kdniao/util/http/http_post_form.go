package http

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/ttlv/kdniao"
)

func HttpPostForm(reqUrl string, params url.Values, logger kdniao.KdniaoLoggerInterface) (HttpResponse, error) {
	startTime := time.Now()
	var httpResp HttpResponse
	resp, err := http.PostForm(reqUrl, params)
	if err != nil {
		httpResp.SetStartTime(startTime)
		logger.Error(reqUrl, "|", params, "|", err)
		return httpResp, err
	}
	defer resp.Body.Close()

	var bodyBytes []byte
	if resp.Body != nil {
		bodyBytes, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			httpResp.SetStartTime(startTime)
			logger.Error(reqUrl, "|", params, "|", err)
			return httpResp, err
		}
	}

	httpResp.SetStartTime(startTime)
	httpResp.SetCode(resp.StatusCode)
	httpResp.SetBytes(bodyBytes)
	logger.Debug(reqUrl, "|", params, "|", resp.StatusCode, "|", httpResp.GetLatencyStr(), "|", string(bodyBytes))
	return httpResp, nil
}
