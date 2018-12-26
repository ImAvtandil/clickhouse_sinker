package model

type Metric interface {
	Get(key string) interface{}
	GetString(key string) string
	GetFloat(key string) float64
	GetInt(key string) int64
	GetArrayInt(key string) []int64
	GetArrayString(key string) []string
}

type DimMetrics struct {
	Dims   []*ColumnWithType
	Fields []*ColumnWithType
}

type ColumnWithType struct {
	Name string
	Type string
}
