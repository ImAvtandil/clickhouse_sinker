package parser

import (
	"strconv"

	"github.com/housepower/clickhouse_sinker/model"
)

type Parser interface {
	Parse(bs []byte) model.Metric
}

func NewParser(typ string) Parser {
	switch typ {
	case "json", "gjson":
		return &GjsonParser{}
	default:
		return &GjsonParser{}
	}
}

var json = jsoniter.ConfigCompatibleWithStandardLibrary

// JsonParser is replaced by GjsonParser
type JsonParser struct {
}

func (c *JsonParser) Parse(bs []byte) model.Metric {
	v := make(map[string]interface{})
	json.Unmarshal(bs, &v)
	return &JsonMetric{v}
}

type JsonMetric struct {
	mp map[string]interface{}
}

func (c *JsonMetric) Get(key string) interface{} {
	return c.mp[key]
}

func (c *JsonMetric) GetString(key string) string {
	//判断object
	val, _ := c.mp[key]
	if val == nil {
		return ""
	}
	switch val.(type) {
	case map[string]interface{}:
		return GetJsonShortStr(val.(map[string]interface{}))

	case string:
		return val.(string)
	}
	return ""
}

func (c *JsonMetric) GetFloat(key string) float64 {
	val, _ := c.mp[key]
	if val == nil {
		return 0
	}
	switch val.(type) {
	case float64:
		return val.(float64)

	case string:
		//这里要转为int ， fuck
		i, _ := strconv.ParseFloat(val.(string), 64)
		return i
	}
	return 0
}

func (c *JsonMetric) GetInt(key string) int64 {
	val, _ := c.mp[key]
	if val == nil {
		return 0
	}
	switch val.(type) {
	case float64:
		return int64(val.(float64))

	case string:
		//这里要转为int ， fuck
		i, _ := strconv.ParseInt(val.(string), 10, 64)
		return i
	}
	return 0
}

func (c *JsonMetric) GetArrayInt(key string) []int64 {
	return []int64{1}
}
func (c *JsonMetric) GetArrayString(key string) []string {
	return []string{""}
}

func GetJsonShortStr(v interface{}) string {
	bs, _ := json.Marshal(v)
	return string(bs)
}
