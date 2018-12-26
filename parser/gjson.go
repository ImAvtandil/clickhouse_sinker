package parser

import (
	"github.com/housepower/clickhouse_sinker/model"

	"github.com/tidwall/gjson"
)

type GjsonParser struct {
}

func (c *GjsonParser) Parse(bs []byte) model.Metric {
	return &GjsonMetric{string(bs)}
}

type GjsonMetric struct {
	raw string
}

func (c *GjsonMetric) Get(key string) interface{} {
	return gjson.Get(c.raw, key).Value()
}

func (c *GjsonMetric) GetString(key string) string {
	return gjson.Get(c.raw, key).String()
}

func (c *GjsonMetric) GetFloat(key string) float64 {
	return gjson.Get(c.raw, key).Float()
}

func (c *GjsonMetric) GetInt(key string) int64 {
	return gjson.Get(c.raw, key).Int()
}

func (c *GjsonMetric) GetArrayInt(key string) []int64 {
	val := gjson.Get(c.raw, key)
	var ran = val.Array()
	var args = make([]int64, len(ran))
	for i, name := range val.Array() {
		args[i] = name.Int()
	}
	return args
}

func (c *GjsonMetric) GetArrayString(key string) []string {
	val := gjson.Get(c.raw, key)
	var ran = val.Array()
	var args = make([]string, len(ran))
	for i, name := range val.Array() {
		args[i] = name.String()
	}
	return args
}
