/*************************************************************************
	> File Name: encrypt_test.go
	> Author: xiangcai
	> Mail: xiangcai@gmail.com
	> Created Time: 2021年06月11日 星期五 15时58分20秒
*************************************************************************/

package test

import (
	"github.com/sgs921107/gcommon"
	"testing"
)

func TestMD5(t *testing.T) {
	data := map[string]string{
		"":                                 "d41d8cd98f00b204e9800998ecf8427e",
		"123456":                           "e10adc3949ba59abbe56e057f20f883e",
		"d41d8cd98f00b204e9800998ecf8427e": "74be16979710d4c4e7c6647856088456",
		"e10adc3949ba59abbe56e057f20f883e": "14e1b600b1fd579f47433b88e8d85291",
	}
	for text, md5String := range data {
		if ret := gcommon.EncryptMD5(text); ret != md5String {
			t.Errorf("gcommon.EncryptMD5(%s) == %s, want %s", text, ret, md5String)
		}
	}
}
