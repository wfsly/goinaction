package search

const dataFile = "data/data.json"

// Feed 包含我们需要处理的数据源的信息
type Feed struct {
	Name string `json:"site"`
	URI  string `json:"link"`
	Type string `json:"type"`
}
