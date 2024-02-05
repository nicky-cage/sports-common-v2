#!/bin/bash

curl -X GET "localhost:9200/xingkong_wagers/_search?pretty" \
    -H 'Content-Type: application/json' \
    -d'
{
    "query": { 
        "bool": { 
            "must": [
                { 
                    "match": { 
                        "bill_no": "Search" 
                    }
                }
            ]
        }
    }
}
'

#curl -X GET "localhost:9200/_search" \
#    -H 'Content-Type: application/json' \
#    -d'
#{
#    "query": { 
#    "bool": { 
#        "must": [
#            { 
#                "match": { "title":   "Search" }
#            }, 
#            { "match": { "content": "Elasticsearch" }}  
#        ],
#        "filter": [ 
#        { "term":  { "status": "published" }}, 
#            { "range": { "publish_date": { "gte": "2015-01-01" }}} ]
#        }
#    }
#}
#'
