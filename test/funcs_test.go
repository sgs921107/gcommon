/*************************************************************************
	> File Name: test_funcs.go
	> Author: xiangcai
	> Mail: xiangcai@gmail.com
	> Created Time: 2020年12月10日 星期四 16时05分49秒
*************************************************************************/

package test

import (
	"github.com/sgs921107/gcommon"
	"math/rand"
	"testing"
	"time"
)

func TestTimeStamp(t *testing.T) {
	if ret := gcommon.TimeStamp(4); ret != 0 {
		t.Errorf("TimeStamp(4) == %d, want 0", ret)
	}
}

// 测试将时间转化为int秒数的func
func TestDurationToIntSecond(t *testing.T) {
	// 指定数据测试
	if gcommon.DurationToIntSecond(time.Second*10) != 10 {
		t.Error("DurationToIntSecond(time.Second * 10) != 10, want 10")
	}
	// 小于1秒的数据
	if gcommon.DurationToIntSecond(10) != 0 {
		t.Error("DurationToIntSecond(10) != 0, want 0")
	}
	// 随机生成1000条case进行测试
	for i := 0; i < 1000; i++ {
		num := rand.Int()
		duration := time.Duration(num)
		want := num / 1e9
		if gcommon.DurationToIntSecond(duration) != want {
			t.Errorf("DurationToIntSecond(%d) != %d, want %d", duration, want, want)
		}
	}
}

func TestFetchUrlHost(t *testing.T) {
	if ret, err := gcommon.FetchUrlHost("http://www.baidu.com/q=hello"); ret != "www.baidu.com" {
		t.Errorf(`gcommon.FetchUrlHost("http://www.baidu.com/q=hello") == (%v, %v), want ("www.baidu.com, nil")`, ret, err)
	}
	if ret, err := gcommon.FetchUrlHost("http://mail.baidu.com?user=go"); ret != "mail.baidu.com" {
		t.Errorf(`gcommon.FetchUrlHost("http://www.baidu.com/q=hello") == (%v, %v), want ("mail.baidu.com, nil")`, ret, err)
	}
}
