package config

import (
	"github.com/actliboy/hoper/server/go/lib/utils/conctrl"
	"log"
	"os/exec"
	"path/filepath"
	"tools/bilibili/rpc"
)

type Customize struct {
	PageBegin         int `init:flag:p`
	PageEnd           int
	FavId             int
	StopTime          int
	WorkCount         uint
	DownloadPath      string
	DownloadVideoPath string
	DownloadPicPath   string
	Cookie            string
	SkipKind          []conctrl.Kind
	FFmpegPath        string
}

type config struct {
	Bilibili Customize
	//Log      log.LogConfig
}

func (c *config) Init() {
	if c.Bilibili.PageEnd == 0 {
		c.Bilibili.PageEnd = 1
	}
	if c.Bilibili.PageBegin == 0 {
		c.Bilibili.PageBegin = c.Bilibili.PageEnd
	}
	if c.Bilibili.WorkCount == 0 {
		c.Bilibili.WorkCount = 5
	}
	if c.Bilibili.StopTime == 0 {
		c.Bilibili.StopTime = 1
	}
	rpc.Cookie = c.Bilibili.Cookie
	c.Bilibili.DownloadPath, _ = filepath.Abs(c.Bilibili.DownloadPath)
	c.Bilibili.DownloadVideoPath = c.Bilibili.DownloadPath + "/video"
	c.Bilibili.DownloadPicPath = c.Bilibili.DownloadPath + "/pic"

	path, err := exec.LookPath("ffmpeg")
	if err != nil {
		log.Println(err)
	} else {
		c.Bilibili.FFmpegPath = path
	}
}

var Conf = &config{}
