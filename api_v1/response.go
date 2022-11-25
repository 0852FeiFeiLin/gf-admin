package api_v1

import "github.com/gogf/gf/v2/frame/g"

type BoolRes bool
type IntRes int
type Int64Res int64
type Int64ArrRes []int64
type MapRes g.Map
type ListRes g.List
type ArrayRes g.Array

type DataRes interface{}

// JsonRes 数据返回通用JSON数据结构
type JsonRes struct {
	Code    int         `json:"code" dc:"错误码((0:成功, 1:失败, >1:错误码)"`
	Message string      `json:"message" dc:"提示信息"`
	Data    interface{} `json:"data" dc:"返回数据(业务接口定义具体数据结构)"`
	Time    string      `json:"time" dc:"响应时间"`
}
