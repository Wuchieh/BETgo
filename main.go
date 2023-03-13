package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var (
	Meats      = make(chan Meat)       // 肉品用的channel
	meatList   []Meat                  // 肉品清單
	timeFormat = "2006-01-02 15:04:05" // 時間格式
)

func init() {
	rand.Seed(time.Now().UnixNano()) // 設置隨機種子

	// 初始化各種肉品
	b := &Beef{}
	c := &Chicken{}
	p := &Pork{}

	// 生成肉品清單
	meatList = append(meatList, b.Generator(10)...)
	meatList = append(meatList, p.Generator(7)...)
	meatList = append(meatList, c.Generator(5)...)

	// 打亂肉品清單
	rand.Shuffle(len(meatList), func(i, j int) {
		meatList[i], meatList[j] = meatList[j], meatList[i]
	})
}

func main() {
	// 生成員工
	A := NewPersonnel("A")
	B := NewPersonnel("B")
	C := NewPersonnel("C")
	D := NewPersonnel("D")
	E := NewPersonnel("E")

	var wg sync.WaitGroup // 用來等待所有員工goroutine結束
	wg.Add(5)

	go func() {
		for _, meat := range meatList {
			Meats <- meat // 發送肉品到channel
		}
		close(Meats) // 發送完畢後關閉channel
	}()

	// 讓 員工 去處理Meats channel的值 並傳入wg參數
	go A.Process(Meats, &wg)
	go B.Process(Meats, &wg)
	go C.Process(Meats, &wg)
	go D.Process(Meats, &wg)
	go E.Process(Meats, &wg)

	wg.Wait() // 等待所有員工goroutine結束
	fmt.Println("所有肉品都已處理完畢")
}
