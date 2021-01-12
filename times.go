package time

import (
	"fmt"
	"time"
)

var (
	cstSh *time.Location
	timeLayout string
	timers time.Time
)
//设置时区
func init()  {
	cstSh, _ = time.LoadLocation("Asia/Shanghai") //设置时区：上海
	timeLayout = "2006-01-02 15:04:05"  //时间模板
	timers = time.Now().In(cstSh)
}

func CurrentTimes() {
	//Unix()秒数
	//UnixNano()秒数和纳秒数
	fmt.Println("当前unix时间（秒）:",timers.Unix())
	fmt.Println("毫秒",int(timers.UnixNano()/1000000))
	fmt.Println("微妙",timers.UnixNano()/1000)
	fmt.Println("纳秒",timers.UnixNano())

	year := timers.Format("2006")
	month := timers.Format("01")
	day := timers.Format("02")
	hour := timers.Format("15")
	min := timers.Format("04")
	second := timers.Format("05")
	dateline := string(year+"年-"+month+"月-"+day+"日 "+hour+"时:"+min+"分:"+second+"秒")
	fmt.Println("获取当前时间的年、月、日 时、分、秒:",dateline)


	//获取当前日期
	date := timers.Format(timeLayout) //返回当地时间
	fmt.Println("获取当前时间：年-月-日 时:分:秒:",date)

	//日期转时间戳
	fmt.Println("当前时间戳(秒)",timers.Unix())
	unixTime,_ := time.ParseInLocation("2006-01-02 15:04:05",date,cstSh)
	sec := unixTime.Unix()
	fmt.Println("年-月-日 时:分:秒:",sec)

	//时间戳转日期
	dataTimeStr := time.Unix(sec, 0).Format(timeLayout)
	fmt.Println("当前时间戳：",dataTimeStr)
}
//2006-01-02 15:04:05.999999999 -0700 MST
//获取当前时间:年-月-日 时:分:秒
func CurrentTime() string {
	return timers.Format(timeLayout)
}

func CurrentShortTime() (year int, month time.Month, day int) {
	return timers.Date()
}

//当前年
func CurrentYear() string {
	return timers.Format("2006")
}

//当前月
func CurrentMonth()  string {
	return timers.Format("01")
}

//当前日
func CurrentDay()  string {
	return timers.Format("02")
}

//当前时
func CurrentHour() string {
	return timers.Format("15")
}

//当前分
func CurrentMinute() string {
	return timers.Format("06")
}

//当前时间秒
func CurrentSecond() int64 {
	return timers.Unix()
}

//当前毫秒
func CurrentMillisecond()  int64 {
	return timers.UnixNano()/1000000
}

//当前微妙
func CurrentNanosecond()  int64 {
	return timers.UnixNano()
}

//日期转时间戳
func TimeToSecond(date string) int64 {
	//fmt.Println("当前时间戳(秒)",timers.Unix())
	//unixTime,_ := time.ParseInLocation("2006-01-02 15:04:05",date,cstSh)
	//sec := unixTime.Unix()
	//fmt.Println(sec)
	fmt.Println("日期转时间戳：")
	unixTime,err := time.ParseInLocation(timeLayout,date,cstSh)
	if err != nil {
		fmt.Printf("时间:%v 转码失败\n",date)
	}
	return unixTime.Unix()
}

//时间戳转日期
func SecondToTimer(timestamp int64) string {
	return time.Unix(timestamp,0).Format(timeLayout)
}
