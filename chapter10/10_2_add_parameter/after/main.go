package main

import (
	"fmt"
	"strings"
)

func main() {
	city := &CompCity{
		Volatility: 1,
		Comps: CompSet{
			{Typ: Simple, Price: 2000},
			{Typ: Simple, Price: 3000},
			{Typ: HighEnd, Price: 8000},
		},
	}
	// 初始
	fmt.Printf("%v\n", city)
	// 第一天
	city.Volatility = 5
	city.RefreshAllComp()
	fmt.Printf("%v\n", city)
	// 第二天
	city.Volatility = 2
	city.RefreshAllComp()
	fmt.Printf("%v\n", city)
}

/*
故事: 电脑城卖电脑
描述:
受到市场的变化, 每天电脑的价格会发生变化, 由Volatility参数影响;
不同电脑, 有不同的型号, 所以自身的计算方式也是不同; 所以一下栗子就是讲解, 如果刷新每一款电脑价格
*/

const (
	_ = iota
	Simple
	HighEnd
)

type (
	// CompCity 电脑城
	CompCity struct {
		Volatility int32
		Comps      CompSet
	}
	// CompSet 电脑集
	CompSet []*Comp
	// 电脑
	Comp struct {
		Typ   int32
		Price int32
	}
)

// RefreshAllComp 刷新所有电脑价格
func (cc *CompCity) RefreshAllComp() {
	// 此处将函数作用域分散, 层层调用, 外部就可以获取单个对象, 进行刷新
	cc.Comps.refresh(cc.Volatility)
}

func (cc *CompCity) String() string {
	builder := strings.Builder{}
	for _, comp := range cc.Comps {
		builder.WriteString(fmt.Sprintf("comp:%+v", comp))
	}
	return builder.String()
}

func (set CompSet) refresh(volatility int32) {
	for _, comp := range set {
		comp.refresh(volatility)
	}
}

func (c *Comp) refresh(volatility int32) {
	switch c.Typ {
	case Simple:
		c.Price *= volatility * 2
	case HighEnd:
		c.Price *= volatility * 3
	}
}
