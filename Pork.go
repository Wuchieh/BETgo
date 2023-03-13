package main

import (
	"fmt"
	"time"
)

type Pork struct {
}

func (p *Pork) Processing(name string) {
	fmt.Println(name, "在", time.Now().Format(timeFormat), "取得豬肉")
	time.Sleep(2 * time.Second) // 處理時間
	fmt.Println(name, "在", time.Now().Format(timeFormat), "處理完豬肉")
}

func (p *Pork) Generator(count int) []Meat {
	var ms []Meat
	for i := 0; i < count; i++ {
		nb := &Pork{}
		ms = append(ms, nb)
		//meatsCount++
	}
	return ms
}

func (p *Pork) Name() string {
	return "豬肉"
}
