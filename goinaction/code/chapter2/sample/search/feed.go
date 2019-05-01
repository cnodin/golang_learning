package search

import (
	"os"
	"encoding/json"
)

const dataFile = "data/data.json"

type Feed struct {
	Name string `json:"site"`
	URI  string `json:"link"`
	Type string `json:"type"`
}

func RetrieveFeeds() ([]*Feed, error)  {
	file, err := os.Open(dataFile)
	if err != nil {
		return nil, err
	}

	//当函数返回时，关闭文件；无论是否会出异常，此语句都会被调用，类似java的finally
	defer file.Close()

	//将文件解码到一个切片中，这个切片的每一项是一个指向一个Feed类型的指针
	var feeds []*Feed
	err = json.NewDecoder(file).Decode(&feeds)
	return feeds, err
}
