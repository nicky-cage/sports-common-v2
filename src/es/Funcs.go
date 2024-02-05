package es

import (
	"context"
	"encoding/json"
	"log"
	"os"

	"github.com/olivere/elastic/v7"
)

// IsDir 判断所给路径是否为文件夹
func isDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

// 判断是否是文件
func isFile(path string) bool {
	return !isDir(path)
}

//通过id判断数据是否存在
func IsExistsById(id string, indexName string, client *elastic.Client) bool {
	exist, err := client.Exists().Index(indexName).Id(id).Do(context.Background())
	if err != nil {
		log.Println(err.Error())
	}
	return exist
}

//批量通过ids拿数据
func MgetByIds(myDocs []MyDocdument, indexName string, client *elastic.Client) map[string]int {
	mgetServ := client.MultiGet()
	esIdsMap := map[string]int{}
	for _, v := range myDocs {
		mgetServ = mgetServ.Add(elastic.NewMultiGetItem().Index(indexName).Id(v.Id))
		esIdsMap[v.Id] = 0 //默认都没有
	}
	retData, err := mgetServ.Do(context.TODO())
	if err != nil {
		log.Println(err.Error())
		return esIdsMap
	}
	if retData != nil && retData.Docs != nil && len(retData.Docs) > 0 {
		for _, v := range retData.Docs {
			if v.Found {
				if _, ok := esIdsMap[v.Id]; ok {
					esIdsMap[v.Id] = 1
				}
			}
		}
	}
	return esIdsMap
}

// EsAdd 添加
func EsAdd(id string, doc *MyDocdument, indexName string, client *elastic.Client) bool {
	//client, err := EsClient()
	//if err != nil {
	//	panic(err)
	//}
	_, err := client.Index().Index(indexName).Id(id).BodyJson(doc).Refresh("wait_for").Do(context.Background())
	if err != nil {
		log.Println(err.Error())
		return false
	}
	return true
}

// EsUpdate 修改
func EsUpdate(id string, docMap *map[string]interface{}, indexName string, client *elastic.Client) bool {
	_, err := client.Update().Index(indexName).Id(id).Doc(docMap).Refresh("wait_for").Do(context.Background())
	if err != nil {
		log.Println(err.Error())
		return false
	}
	return true
}

// EsDel 删除
func EsDel(id string, indexName string, client *elastic.Client) bool {
	_, err := client.Delete().Index(indexName).Id(id).Refresh("wait_for").Do(context.Background())
	if err != nil {
		log.Println(err.Error())
		return false
	}
	return true
}

// EsSearch 查找
func EsSearch(criteria *SearchCriteria, indexName string, client *elastic.Client) []MyDocdument {
	var mydocs []MyDocdument
	query := elastic.NewBoolQuery()
	query = query.Must(elastic.NewTermQuery("category", criteria.Category))
	query = query.Must(elastic.NewTermQuery("title", criteria.Title))
	query = query.Must(elastic.NewTermQuery("update_at", criteria.UpdatedAt))
	resp, err := client.Search().Index(indexName).Query(query).
		Sort(criteria.Sort, criteria.Order == "ASC").From(criteria.Offset).
		Size(criteria.Limit).Do(context.Background())
	if err != nil {
		log.Println(err.Error())
		return mydocs
	}
	if resp.Hits.TotalHits.Value > 0 {
		log.Printf("Found a total of %v Employee \n", resp.Hits.TotalHits)

		for _, v := range resp.Hits.Hits {

			var d MyDocdument
			err := json.Unmarshal(v.Source, &d) //另外一种取数据的方法
			if err != nil {
				log.Println("Deserialization failed")
			}
			mydocs = append(mydocs, d)
			log.Printf("Doc name %s : %s\n", d.Title, d.Category)
		}
	} else {
		log.Printf("Found no Doc \n")
	}
	return mydocs
}

// Aggregation 总的记录数
// es 			ES客户端连接对接
// esIndexName 	索引
// query 		查询query
// callback 	回调函数, 对于每个elastic.SearchResult 进行回调
// fields 		指定统计字段
func AggregationStat(client *elastic.Client, esIndexName string, query *elastic.BoolQuery, callback func(*elastic.SearchResult), fields ...string) {
	FetchAggreationsByPage(client, esIndexName, query, callback, 1, fields...)
}

// Aggregation 总的记录数
// es 			ES客户端连接对接
// esIndexName 	索引
// query 		查询query
// callback 	回调函数, 对于每个elastic.SearchResult 进行回调
// fields 		指定统计字段
func FetchAggreationsByPage(client *elastic.Client, esIndexName string, query *elastic.BoolQuery, callback func(*elastic.SearchResult), page int, fields ...string) {
	pageSize := 1024                       // 分页
	currentContext := context.Background() // 当前协程上下文
	fsc := elastic.NewFetchSourceContext(true).Include(fields...)
	builder := func() *elastic.SearchSource {
		if query != nil {
			return elastic.NewSearchSource().Query(query).FetchSourceContext(fsc)
		}
		return elastic.NewSearchSource().FetchSourceContext(fsc)
	}()
	cursor, err := client.Search(esIndexName).SearchSource(builder).From(pageSize*(page-1)).Size(pageSize).Sort("created_at", false).Do(currentContext)
	if err != nil {
		return
	}
	callback(cursor)
	totalRecords := int(cursor.TotalHits()) // 总计
	if page*pageSize > totalRecords {
		return
	}
	FetchAggreationsByPage(client, esIndexName, query, callback, page+1, fields...)
}
