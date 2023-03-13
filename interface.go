package main

// Meat 肉品接口
type Meat interface {
	Processing(string)    // 處理 ［肉］事件
	Generator(int) []Meat // 用來產生該類陣列
	Name() string         // 測試用
}
