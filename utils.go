/**
 * @Author: Jacky
 * @Description:
 * @File: utils
 * @Version: 1.0.0
 * @Date: 2022/4/27 15:48
 */
package submail

import (
	"encoding/json"
	"sort"
	"strings"
)

func SortAndJoin(data map[string]string, sep string) string {
	// 获取 key
	keys := make([]string, 0, len(data))
	for key, _ := range data {
		keys = append(keys, key)
	}

	// 排序
	sort.Strings(keys)

	// 组合
	kv := make([]string, 0, len(data))
	for _, key := range keys {
		kv = append(kv, key+"="+data[key])
	}

	return strings.Join(kv, sep)
}

func JsonContentType() (string, string) {
	return "Content-Type", "application/json;charset=UTF-8"
}

func JsonMarshal(v interface{}) string {
	str, _ := json.Marshal(v)
	return string(str)
}
