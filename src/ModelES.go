package common

import (
	"context"

	"github.com/olivere/elastic/v7"
)

// IModelMg 接口
type IModelEs interface {
	FindId(string, *elastic.Client, string) (*elastic.GetResult, error)
	IsExistsById(string, *elastic.Client, string) (bool, error)
	Search(string, *elastic.Client, *elastic.BoolQuery, *SearchPageCriteria) (*elastic.SearchResult, error)
}

// ModelMg 基于MongoDb的Model
type ModelEs struct {
	EsIndexName        string //索引名称
	Client             *elastic.Client
	SearchPageCriteria *SearchPageCriteria
}

// SearchPageCriteria
type SearchPageCriteria struct {
	Limit         int    `json:"limit"`
	Offset        int    `json:"offset"`
	OrderType     bool   `json:"order_type"`      //true 升序排序 false降序排序
	SortFieldName string `json:"sort_field_name"` //排序的字段
}

// GetExIndexName
func (ths *ModelEs) GetExIndexName() string {
	return ths.EsIndexName
}

// FindId 查询单条数据
func (ths *ModelEs) FindId(platform string, client *elastic.Client, id string) (*elastic.GetResult, error) {
	res, err := client.Get().Index(platform + "_" + ths.EsIndexName).Id(id).Do(context.Background())
	if err != nil {
		return nil, err
	}
	return res, nil
}

// IsExistsById
func (ths *ModelEs) IsExistsById(platform string, client *elastic.Client, id string) (bool, error) {
	exist, err := client.Exists().Index(platform + "_" + ths.EsIndexName).Id(id).Do(context.Background())
	if err != nil {
		return false, err
	}
	return exist, nil
}

// Search
func (ths *ModelEs) Search(platform string, client *elastic.Client, bQuery *elastic.BoolQuery, pageCriteria *SearchPageCriteria) (*elastic.SearchResult, error) {
	if pageCriteria == nil {
		pageCriteria = ths.SearchPageCriteria
	}
	esResp, err := client.Search().Index(platform+"_"+ths.EsIndexName).Query(bQuery).
		Sort(pageCriteria.SortFieldName, pageCriteria.OrderType).From(pageCriteria.Offset).
		Size(pageCriteria.Limit).Do(context.Background())
	return esResp, err
}
