package time

import (
	"fmt"
	"testing"
)

func TestTime(t *testing.T) {
	fmt.Printf("当前时间：%v \n",CurrentTime())

	fmt.Printf("当前年：%v \n",CurrentYear())
	fmt.Printf("当前月:%v \n",CurrentMonth())
	fmt.Printf("当前日:%v \n",CurrentDay())
	fmt.Printf("当前时:%v \n",CurrentHour())
	fmt.Printf("当前分:%v \n",CurrentMinute())
	fmt.Printf("当前时间戳(秒秒):%v \n",CurrentSecond())
	fmt.Printf("当前时间戳(毫秒):%v \n",CurrentMillisecond())
	fmt.Printf("当前时间戳(微妙):%v \n",CurrentNanosecond())

	fmt.Printf("日期:%v 转为时间戳 %v \n",CurrentTime(),TimeToSecond(CurrentTime()))
	//second,_ := strconv.ParseInt(,10,64)
	fmt.Printf("时间戳格式化为日期%v \n", SecondToTimer(CurrentSecond()))
	//fmt.Println(timers.Unix())
	//fmt.Println(CurrentShortTime())
	//time.Unix()
	//fmt.Println(SecondToTimer(CurrentTime()))
}
