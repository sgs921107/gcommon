/*************************************************************************
	> File Name: test_funcs.go
	> Author: xiangcai
	> Mail: xiangcai@gmail.com
	> Created Time: 2020年12月10日 星期四 16时05分49秒
*************************************************************************/

package test

import (
	"go_common"
	"testing"
)

func TestTimeStamp(t *testing.T) {
	if ret := go_common.TimeStamp(4); ret != 0 {
		t.Errorf("TimeStamp(4) = %d, want 0", ret)
	}
}
