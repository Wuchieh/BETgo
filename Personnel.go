package main

import "sync"

type Personnel struct {
	Name string
}

// NewPersonnel 輸入名稱 生成員工
func NewPersonnel(name string) *Personnel {
	return &Personnel{
		Name: name,
	}
}

// Process 員工工作區
func (p *Personnel) Process(meats chan Meat, s *sync.WaitGroup) {
	// 從channel中取出肉品 並處理
	for i := range meats {
		i.Processing(p.Name)
	}

	s.Done() // channel已關閉時 且工作已經完成時 向WaitGroup發送完成訊號
}
