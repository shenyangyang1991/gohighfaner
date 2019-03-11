// json 高性能工具集
// 时间 2019年3月9号
// 作者 申杨杨

package json

import "github.com/json-iterator/go"

// json 标准
var json = jsoniter.ConfigCompatibleWithStandardLibrary

// Marshal is interface to []byte
func Marshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

// Unmarshal is []byte to interface
func Unmarshal(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}
