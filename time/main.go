package main

import (
	"fmt"
	"reflect"
	"strconv"
	"time"
)

func printType(obj interface{}, name string) {
	fmt.Printf("obj <%s> value: %s and type: %s\n", name, obj, reflect.TypeOf(obj))
}

func main() {

	// 打印日期格式
	fmt.Println("")
	now := time.Now() // time.Time
	// fmt.Printf("time now is: %s and type: <%s>", now, reflect.TypeOf(now))
	printType(now, "now := time.Now()") // 2018-06-26 20:19:27.6787991 +0800 CST m=+0.004009401
	// time.Time 分解
	printType(now.Year(), "now.Year()")             // 2018, int
	printType(now.Month(), "now.Month()")           // June, time.Month
	printType(int(now.Month()), "int(now.Month())") // 6, int
	printType(now.Day(), "now.Day()")               // 26, int
	printType(now.Hour(), "now.Hour()")             // 20, int
	printType(now.Minute(), "now.Minute()")         // 19, int
	printType(now.Second(), "now.Second()")         // 27, int
	printType(now.Nanosecond(), "now.Nanosecond()") // 678799100, int

	// from int to
	fmt.Println("")
	timestampInt := now.Unix()                            // int
	printType(timestampInt, "timestampInt := now.Unix()") // 1530014826

	// string date format type
	fmt.Println("")
	printType(time.RFC1123, "time.RFC1123")         // Mon, 02 Jan 2006 15:04:05 MST, string
	printType(time.RFC1123Z, "time.RFC1123Z")       // Mon, 02 Jan 2006 15:04:05 -0700
	printType(time.RFC3339, "time.RFC3339")         // 2006-01-02T15:04:05Z07:00
	printType(time.RFC3339Nano, "time.RFC3339Nano") // 2006-01-02T15:04:05.999999999Z07:00
	printType(time.RFC822, "time.RFC822")           // 02 Jan 06 15:04 MST
	printType(time.RFC850, "time.RFC850")           // Monday, 02-Jan-06 15:04:05 MST
	printType(time.RubyDate, "time.RubyDate")       // Mon Jan 02 15:04:05 -0700 2006
	printType(time.UnixDate, "time.UnixDate")       // Mon Jan _2 15:04:05 MST 2006       // 2016-01-02 15:04:05
	s, _ := time.Parse("2016-10-01 00:00:00", "Tue, 26 Jun 2018 20:19:27 CST")
	printType(s, "s, _ := time.Parse(\"2016-10-01 00:00:00\", \"Tue, 26 Jun 2018 20:19:27 CST\")") //  0001-01-01 00:00:00 +0000 UTC and type: time.Time, GO里面的时间格式是一个非常奇怪的东西

	// format datetime
	fmt.Println("")
	printType(now.Format(time.RFC1123), "now.Format(time.RFC1123)") // Tue, 26 Jun 2018 20:19:27 CST, string

	// timestamp int to
	fmt.Println("")
	printType(time.Unix(timestampInt, 0), "time.Unix(timestampInt, 0)") // 2018-06-26 20:19:27 +0800 CST

	// timestamp string to int to time.Time
	fmt.Println("")
	t, _ := strconv.ParseInt("1530014826", 10, 64)
	printType(time.Unix(t, 0), "time.Unix(t, _ := strconv.ParseInt(\"1530014826\", 10, 64), 0)") // 2018-06-26 20:07:06 +0800 CST

	// convert string into time.Time
	fmt.Println("")
	str2Time, _ := time.Parse(time.RFC1123, "Tue, 26 Jun 2018 20:19:27 CST")
	printType(str2Time, "str2Time, _ := time.Parse(time.RFC1123, \"Tue, 26 Jun 2018 20:19:27 CST\")")

	// time.Duration
	fmt.Println("")
	printType(time.Duration(40), "time.Duration(60)")                         // 40ns
	printType(time.Duration(40*time.Second), "time.Duration(60*time.Second)") // 40s
	printType(time.After(5*time.Second), "time.After(5*time.Second)")         // value: <-chan time.Time=0xc04209a000 and type: <-chan time.Time

	// More date format, should look for other packages
	// Example: https: //github.com/araddon/dateparse
	//

	// 另外，golang中奇怪的日期其实点, 2016, 1, 2, 3pm, 4, 5
	fmt.Println("")
	d1, _ := time.Parse("01/02/2006", "10/01/2018") // 奇怪的日期format
	printType(d1, "time.Parse(\"01/02/2006\", \"10/01/2018\")")
	fmt.Println("")
	d2, _ := time.Parse("2016-01-02", "2018-06-27")
	printType(d2, "time.Parse(\"2016-01-02\", \"2018-06-27\")  // can't parse right")

	// how about add dates?
	fmt.Println("")
	printType(d1.AddDate(-1, -2, -3), "d1.AddDate(-1, -2, -3) // d1 = 2018-10-01 ")

	fmt.Println("")
	printType(d1.Add(time.Duration(30*time.Hour)), "d1.Add(time.Duration(30 * time.Hour))")

	// 计算两个日期之间的差距, add or sub一个time.Time返回一个time.Duration
	fmt.Println("")
	printType(d1.Sub(time.Now()), "d1.Sub(time.Now())")

}
