package main

import (
	"fmt"
	"time"
)

type Chicken struct {
}

func (c *Chicken) Processing(name string) {
	fmt.Println(name, "在", time.Now().Format(timeFormat), "取得雞肉")
	time.Sleep(3 * time.Second) // 處理時間
	fmt.Println(name, "在", time.Now().Format(timeFormat), "處理完雞肉")
}

func (c *Chicken) Generator(count int) []Meat {
	var ms []Meat
	for i := 0; i < count; i++ {
		nb := &Chicken{}
		ms = append(ms, nb)
		//meatsCount++
	}
	return ms
}

func (c *Chicken) Name() string {
	return "雞肉"
}
