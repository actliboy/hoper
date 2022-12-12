package stringsi

import (
	"log"
	"testing"
)

func TestFormatLen(t *testing.T) {
	s := "post"
	log.Println(FormatLen(s, 10), "test")
	s = "AutoCommit"
	log.Println(CamelToSnake(s))
}

func TestReplaceRuneEmpty(t *testing.T) {
	s := "p我o爱s中t"
	log.Println(ReplaceRuneEmpty(s, []rune{'o'}))
	log.Println(ReplaceRuneEmpty(s, []rune{'o', 's'}))
	log.Println(ReplaceRuneEmpty(s, []rune{'o', 't'}))
	log.Println(ReplaceRuneEmpty(s, []rune{'中', 't'}))
}

func TestCountdownCutoff(t *testing.T) {
	log.Println(CountdownCutoff("https://video.weibo.com/media/play?livephoto=https%3A%2F%2Flivephoto.us.sinaimg.cn%2F002OnXdGgx07YpcajtkH0f0f0100gv8Q0k01.mov", "%2F"))
	log.Println(CountdownCutoff("https://wx1.sinaimg.cn/orj360/6ebedee6ly1h566bbzyc6j20n00cuabd.jpg", "/"))
	log.Println(Cutoff("https://wx1.sinaimg.cn/orj360/6ebedee6ly1h566bbzyc6j20n00cuabd.jpg", "wx1"))
	log.Println(CountdownCutoff(CutoffContain("https://f.video.weibocdn.com/o0/F9Nmm1ZJlx080UxqxlJK010412004rJS0E010.mp4?label=mp4_hd&template=540x960.24.0&ori=0&ps=1CwnkDw1GXwCQx&Expires=1670569613&ssig=fAQcBh4HGt&KID=unistore,video", "mp4"), "/"))
}
