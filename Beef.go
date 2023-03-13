package main

import (
	"fmt"
	"time"
)

type Beef struct {
}

func (b *Beef) Processing(name string) {
	fmt.Println(name, "在", time.Now().Format(timeFormat), "取得牛肉")
	time.Sleep(time.Second) // 處理時間
	fmt.Println(name, "在", time.Now().Format(timeFormat), "處理完牛肉")
}

func (b *Beef) Generator(count int) []Meat {
	var ms []Meat
	for i := 0; i < count; i++ {
		nb := &Beef{}
		ms = append(ms, nb)
		//meatsCount++
	}
	return ms
}

func (b *Beef) Name() string {
	return "牛肉"
}
