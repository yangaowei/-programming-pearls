package main

import (
	"fmt"
	"reflect"
	"time"
)

func main() {
	dur := time.Minute
	fmt.Println(time.Minute.Minutes())
	fmt.Println(reflect.TypeOf(time.Minute))
	fmt.Println(int64(dur / time.Second))
	utime := time.Now().Unix()

	fmt.Println(utime)
	udur := int64(dur / time.Second)
	slot := utime / udur
	delay := time.Duration((slot+1)*udur-utime) * time.Second

	fmt.Println(delay)
	fmt.Println(time.Second)

	// c := time.Tick(1 * time.Second)
	// for now := range c {
	// 	fmt.Printf("%v %v\n", now.UnixNano(), reflect.TypeOf(now))
	// }
	t := time.NewTicker(1 * time.Second)

	for now := range t.C {
		fmt.Println(now)
		t.Stop()
	}
}
