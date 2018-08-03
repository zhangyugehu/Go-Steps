package main

import (
	"study/lottery/engine"
	"study/lottery/parser"
)

func main() {
	e:=engine.SimpleEngine{

	}
	lotteryParser := parser.Lottery360{

	}
	e.Run(engine.Request{
		"http://chart.cp.360.cn/kaijiang/ssq?lotId=220051&chartType=undefined&spanType=0&span=3000&r=0.11565802984942652#roll_0",
		lotteryParser.Ssq,
	})
	e.Run(engine.Request{
		"https://chart.cp.360.cn/kaijiang/sd?lotId=210053&chartType=undefined&spanType=0&span=2000&r=0.9942958945259759#roll_0",
		lotteryParser.Fc3d,
	})
	e.Run(engine.Request{
		"https://chart.cp.360.cn/zst/k3jl?lotId=258203&chartType=dst&spanType=0&span=2000&r=0.8753756247257816#roll_0",
		lotteryParser.Q3,
	})
	e.Run(engine.Request{
		"https://chart.cp.360.cn/zst/k3jl?lotId=258203&chartType=dst&spanType=0&span=2000&r=0.8753756247257816#roll_0",
		lotteryParser.Q3,
	})
	e.Run(engine.Request{
		"",
		lotteryParser.Dlt,
	})
}
