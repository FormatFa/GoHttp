package main

import (
	"fmt"
	"net/url"
	"testing"
)

func TestAdd2(t *testing.T) {
	str := "grant_type=password&username=1758@qq.com&password=123456&client_id=client&client_secret=client_secret&scope=projects user_info issues notes"
	scapedPath := url.PathEscape(str)
	fmt.Println("Escaped Path:", scapedPath)

	// 使用EscapeQuery函数对查询参数进行编码转义
	escapedQuery := url.QueryEscape(str)
	fmt.Println("Escaped Query:", escapedQuery)
}
