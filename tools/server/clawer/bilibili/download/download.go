package download

import (
	"context"
	"fmt"
	"github.com/liov/hoper/server/go/lib/utils/fs"
	"github.com/liov/hoper/server/go/lib/utils/net/http/client"
	timei "github.com/liov/hoper/server/go/lib/utils/time"
	"github.com/liov/hoper/server/go/lib_v2/utils/net/http/client/crawler"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	claweri "tools/clawer"
	"tools/clawer/bilibili/config"
	"tools/clawer/bilibili/dao"
	"tools/clawer/bilibili/rpc"
)

func (video *Video) DownloadVideoReq(typ string, order int, url string) *crawler.Request {
	return &crawler.Request{
		TaskMeta: crawler.TaskMeta{BaseTaskMeta: crawler.BaseTaskMeta{Key: "下载视频：" + strconv.Itoa(video.Cid) + typ}, Kind: KindDownloadVideo},
		TaskFunc: func(ctx context.Context) ([]*crawler.Request, error) {
			return video.DownloadVideo(typ, order, url)
		},
	}
}

func (video *Video) DownloadVideo(typ string, order int, url string) ([]*crawler.Request, error) {

	var filename string

	video.Part = fs.FileNameClean(video.Part)
	video.Title = fs.FileNameClean(video.Title)
	if strings.HasSuffix(video.Title, video.Part) {
		video.Part = PartEqTitle
	}

	if video.CodecId == VideoTypeFlv {
		pubAt := video.PubAt.Format(timei.TimeFormatCompact)
		dir := config.Conf.Bilibili.DownloadPath + fs.PathSeparator + strconv.Itoa(video.Uid) + fs.PathSeparator + pubAt[:4]
		filename = fmt.Sprintf("%s_%d_%d_%d_%s_%s_%d_%d.flv.downloading", pubAt, video.Uid, video.Aid, video.Cid, video.Title, video.Part, order, video.Quality)
		filename = dir + fs.PathSeparator + filename

	} else {
		filename = fmt.Sprintf("%d_%d_%d.m4s.%s.downloading", video.Uid, video.Aid, video.Cid, typ)
		filename = config.Conf.Bilibili.DownloadTmpPath + fs.PathSeparator + filename
	}

	newname := filename[:len(filename)-len(DownloadingExt)]

	_, err := os.Stat(newname)
	if os.IsNotExist(err) {
		referer := rpc.GetViewUrl(video.Aid)
		referer = referer + fmt.Sprintf("/?p=%d", video.Page)

		c := http.Client{CheckRedirect: genCheckRedirectfun(referer)}

		request, err := http.NewRequest(http.MethodGet, url, nil)
		if err != nil {
			return nil, err
		}
		request.Header.Set("User-Agent", client.UserAgent1)
		request.Header.Set("Accept", "*/*")
		request.Header.Set("Accept-Language", "en-US,en;q=0.5")
		request.Header.Set("Accept-Encoding", "gzip, deflate, br")
		request.Header.Set("Range", "bytes=0-")
		request.Header.Set("Referer", referer)
		request.Header.Set("Origin", "https://www.bilibili.com")
		request.Header.Set("Connection", "keep-alive")
		request.Header.Set("Cookie", rpc.Cookie)

		resp, err := c.Do(request)
		if err != nil {
			log.Printf("下载 %d 时出错, 错误信息：%s", video.Cid, err)
			return nil, err
		}

		if resp.StatusCode != http.StatusPartialContent {
			log.Printf("下载 %d 时出错, 错误码：%d", video.Cid, resp.StatusCode)
			return nil, fmt.Errorf("错误码： %d", resp.StatusCode)
		}
		defer resp.Body.Close()

		file, err := fs.Create(filename)
		if err != nil {
			log.Println("错误信息：", err)
			return nil, err
		}

		log.Println("正在下载："+filename, "质量：", video.Quality)
		_, err = io.Copy(file, resp.Body)
		if err != nil {
			os.Remove(filename)
			log.Printf("下载失败 filename: %s", filename)
			log.Println("错误信息：", err)

			// request again
			//go requestLater(file, resp, video)
			return nil, err
		}
		file.Close()
		err = os.Rename(filename, newname)
		if err != nil {
			return nil, err
		}
	}
	log.Println("下载完成：" + newname)

	if video.CodecId == VideoTypeFlv {
		dao.Dao.Hoper.Table(dao.TableNameVideo).Where("cid = ?", video.Cid).Update("record", 3)
		dir := claweri.Dir{
			Platform: 3,
			UserId:   video.Uid,
			KeyId:    video.Cid,
			KeyIdStr: fmt.Sprintf("%d_%d", video.Aid, video.Cid),
			BaseUrl:  fmt.Sprintf("%s_%s_%d_%d.flv", video.Title, video.Part, order, video.Quality),
			Type:     3,
			PubAt:    video.PubAt,
		}
		dao.Dao.Hoper.Create(&dir)
	}

	if video.CodecId == VideoTypeM4sCodec12 || video.CodecId == VideoTypeM4sCodec7 {
		merge.Add(video)
	}
	return nil, nil
}

func genCheckRedirectfun(referer string) func(req *http.Request, via []*http.Request) error {
	return func(req *http.Request, via []*http.Request) error {
		req.Header.Set("Referer", referer)
		return nil
	}
}
