/*
 * @Author: xiangcai
 * @Date: 2021-08-18 11:04:25
 * @LastEditors: xiangcai
 * @LastEditTime: 2021-08-18 11:32:27
 * @Description: file content
 */
package test

import (
	"testing"
	"time"

	"github.com/sgs921107/gcommon"
)

func TestGroupSize(t *testing.T) {
	size := 50
	wg := gcommon.NewThreadGroup(size)
	want := 10000
	for i := 0; i < want; i++ {
		wg.Add(1)
		go func(wg gcommon.Group) {
			length := wg.Len()
			if length > size {
				t.Errorf("want length <= %d, want %d", size, length)
			}
			time.Sleep(time.Millisecond)
			wg.Done()
		}(wg)
	}
}
