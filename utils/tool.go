package utils

import (
	"context"
	"time"
)

func SetTimeout(f func(), timeout int) context.CancelFunc {
	ctx, cancelFunc := context.WithCancel(context.Background())
	go func() {
		select {
		case <-ctx.Done():
		case <-time.After(time.Duration(timeout) * time.Second):
			f()
		}
	}()
	return cancelFunc
}

func SetInterval(f func(), timeout int) context.CancelFunc {
	ctx, cancelFunc := context.WithCancel(context.Background())
	go func() {
		for {
			time.Sleep(time.Duration(timeout) * time.Second)
			select {
			case <-ctx.Done():
				return
			default:
				f()
			}
		}
	}()
	return cancelFunc
}

// 防抖函数,单位时间内连续触发的事件,但是只在最后一次出发后开始计算执行一次
func Debounce(f func(), wait int) func() {
	var cf context.CancelFunc
	return func() {
		if cf != nil {
			cf()
		}
		cf = SetTimeout(f, wait)
	}
}

// 节流函数,一段时间内无论触发多少次到时间后只执行一次
func Throttled(f func(), wait int) func() {
	var cf context.CancelFunc
	return func() {
		if cf == nil {
			cf = SetTimeout(f, wait)
		}
	}
}
