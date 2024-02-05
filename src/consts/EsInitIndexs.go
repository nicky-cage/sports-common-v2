package consts

var EsServersUrl []string
var EsPlatformName string
var EsLoginUser string
var EsLoginPass string

// EsInitMappings 默认配置动态模板
var EsInitMappings = map[string]map[string]string{
	"default": {
		"wagers": `{
				"mappings": {
				"dynamic_templates":[
						{
							"all_to_double":{
								"match_mapping_type":"double",
								"mapping":{
									"type":"double"
								}
							}
						}],
					"properties": {
						"Id": {
							"type": "keyword",
							"index": true
						},
						"username": {
							"type": "keyword",
							"index": true
						},
						"playname": {
							"type": "keyword",
							"index": true
						},
						"user_id": {
							"type": "integer",
							"index": true
						},
						"is_agent": {
								"type": "byte"
						},
						"top_name": {
								"type": "keyword"
						},
						"top_code": {
								"type": "keyword",
								"index": true
						},
						"top_id": {
								"type": "integer"
						},
						"forefather_id": {
								"type": "integer"
						},
						"game_code": {
								"type": "keyword",
								"index": true
						},
						"game_type": {
								"type": "short"
						},
						"game_code_type": {
								"type": "keyword",
								"index": true
						},
						"bill_no": {
								"type": "keyword",
								"index": true
						},
						"bet_money": {
								"type": "double",
								"index": false
						},
						"return_money": {
								"type": "double",
								"index": false
						},
						"net_money": {
								"type": "double",
								"index": false
						},
						"valid_money": {
								"type": "double",
								"index": false
						},
						"valid_money_ext": {
								"type": "double",
								"index": false
						},
                        "rebate_money": {
								"type": "double",
								"index": false
						},
 						"rebate_ratio": {
								"type": "double",
								"index": false
						},
						"status": {
								"type": "byte"
						},
						"extend": {
							"type": "text"
						},
						"created_at": {
									"type": "date",
								  "format": "strict_date_optional_time||epoch_second"
						},
						"updated_at": {
									"type": "date",
								  "format": "strict_date_optional_time||epoch_second"
						}
					}
				}
			}`,
		"login_logs": `{
				"mappings": {
					"properties": {
						"Id": {
							"type": "keyword"
						},
						"username": {
							"type": "keyword"
						},
						"user_id": {
								"type": "integer"
						},
						"top_name": {
								"type": "keyword"
						},
						"login_ip": {
								"type": "keyword"
						},
						"last_ip": {
								"type": "keyword"
						},
						"login_type": {
								"type": "keyword"
						},
						"device_no": {
								"type": "keyword"
						},
						"note": {
								"type": "keyword"
						},
						"created_at": {
								"type": "date",
	  							  "format": "strict_date_optional_time||epoch_second"
						},
						"updated_at": {
								"type": "date",
								  "format": "strict_date_optional_time||epoch_second"
						}
					}
				}
			}`,
		"admin_logs": `{
				"mappings": {
					"properties": {
						"Id": {
							"type": "keyword"
						},
						"admin_name": {
							"type": "keyword"
						},
						"admin_id": {
								"type": "integer"
						},
						"login_ip": {
								"type": "ip"
						},
						"menu_name": {
								"type": "keyword"
						},
						"menu_id": {
								"type": "integer"
						},
						"menu_name_no": {
								"type": "short"
						},
						"request_url_path": {
								"type": "keyword"
						},
						"get_data": {
								"type": "text"
						},
						"post_data": {
								"type": "text"
						},
						"note": {
								"type": "text"
						},
						"created_at": {
								"type": "date",
	  							  "format": "strict_date_optional_time||epoch_second"
						},
						"updated_at": {
								"type": "date",
								  "format": "strict_date_optional_time||epoch_second"
						}
					}
				}
			}`,
	},
}

/*
"shipu_tests": `{
		"mappings": {
			"properties": {
				"id": {
					"type": "long"
				},
				"stype": {
					"type": "byte"
				},
				"title": {
					"type": "keyword"
				},
				"content": {
						"type": "text"
				},
				"author": {
						"type": "keyword"
				},
				"money": {
						"type": "double"
				},
				"addtime": {
						"type": "date",
						"format": "yyyy-MM-dd HH:mm:ss||yyyy-MM-dd||epoch_millis"
				}
			}
		}
	}`,
*/
