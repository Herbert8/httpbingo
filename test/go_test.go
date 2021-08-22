package tests

import (
	"net/url"
	"strings"
	"testing"
	"unicode/utf8"
)

func TestTime(t *testing.T) {
	//t.Log("abc")
	//t.Log(time.Unix(10, 5*1000000000))
	//var ck http.Cookie
	////ck.MaxAge

}

func TestStr(t *testing.T) {
	s := `abc
x	xxx
2	22`

	t.Log(len(s))
}



//func TestSplit(t *testing.T) {
//	s := "/base/aa/bb/cc///"
//
//	arr := parsePathParams(s, "/base")
//
//	t.Log(arr)
//}

func TestReplace(t *testing.T) {
	s := "abcmmmabcnnnabcxxx"
	r := "abc"
	ret := strings.ReplaceAll(s, r, "")
	t.Log(ret)
}

func TestEncode(t *testing.T) {
	t.Log(url.QueryEscape("a/b?c&d%e"))
	t.Log(url.PathEscape("a/b?c&d%e"))
}

func TestLen(t *testing.T) {
	s := "测试中文123"
	t.Log(utf8.RuneCountInString(s))
}

func TestArrLen(t *testing.T) {
	//var arr [5]int
	arr := [...]int{1, 2, 3}
	println(len(arr))
}