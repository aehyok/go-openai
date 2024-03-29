package service

import (
	"encoding/json"
	"errors"
	"geekdemo/middleware"
	"net/http"

	"github.com/mitchellh/mapstructure"
)

type PointRequest struct {
	Points []Point `json:"points"`
}

type Point struct {
	ID      string      `json:"id"`
	Payload interface{} `json:"payload"`
	Vector  []float64   `json:"vector"`
}

type PointResponse struct {
	Result interface{} `json:"result"`
	Status interface{} `json:"status"`
	Time   float64     `json:"time"`
}

type PointSearchRequest struct {
	Params      map[string]interface{} `json:"params"`
	Vector      []float64              `json:"vector"`
	Limit       int                    `json:"limit"`
	WithPayload bool                   `json:"with_payload"`
	WithVector  bool                   `json:"with_vector"`
}
type Match struct {
	Value string `json:"value"`
}
type Must struct {
	Key   string `json:"key"`
	Match Match  `json:"match"`
}

type SearchResult struct {
	ID      string      `json:"id"`
	Version int         `json:"version"`
	Score   float64     `json:"score"`
	Payload interface{} `json:"payload"`
	Vector  []float64   `json:"vector,omitempty"`
}

const (
	collectionApi = "/collections"
	pointsApi     = "/points"
	searchApi     = "/search"
)

func CreatePoints(collectionName string, pointRequest PointRequest) (err error) {
	response := &CommonResponse{}
	var reqBytes []byte
	reqBytes, err = json.Marshal(pointRequest)
	if err != nil {
		// common.Logger.Error(err.Error())
		return
	}

	body, err := middleware.Send(http.MethodPut, collectionApi+"/"+collectionName+pointsApi+"?wait=true", reqBytes)
	if err != nil {
		// common.Logger.Error(err.Error())
		return
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		return
	}
	if response.Result == nil {
		return errors.New(response.Status.(map[string]interface{})["error"].(string))
	}
	return

}

// SearchPoints 搜索点
func SearchPoints(collectionName string, pointSearchRequest PointSearchRequest) (res []SearchResult, err error) {
	response := &CommonResponse{}
	var reqBytes []byte
	reqBytes, err = json.Marshal(pointSearchRequest)
	if err != nil {
		// common.Logger.Error(err.Error())
		return
	}

	// 发送请求
	body, err := middleware.Send(http.MethodPost, collectionApi+"/"+collectionName+pointsApi+searchApi, reqBytes)
	if err != nil {
		// common.Logger.Error(err.Error())
		return
	}

	// 解析响应
	err = json.Unmarshal(body, &response)
	if err != nil {
		// common.Logger.Error(err.Error())
		return
	}
	if response.Result == nil {
		return res, errors.New(response.Status.(map[string]interface{})["error"].(string))
	}
	list := response.Result.([]interface{})
	for _, v := range list {
		sr := SearchResult{}
		err = mapstructure.Decode(v, &sr)
		if err != nil {
			// common.Logger.Error(err.Error())
			return
		}
		res = append(res, sr)
	}
	return

}

type CommonResponse struct {
	Result interface{} `json:"result"`
	Status interface{} `json:"status"`
	Time   float64     `json:"time"`
}
