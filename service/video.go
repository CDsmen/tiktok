package service

import (
	"bytes"
	"fmt"
	"io"
	"math/rand"
	"mime/multipart"
	"os"
	"strconv"
	"tiktok/dao"
	"tiktok/myRedis"
	"tiktok/myjwt"
	"time"

	ffmpeg "github.com/u2takey/ffmpeg-go"
)

// GenerateVideoCover 获取封面
func GenerateVideoCover(inFileName string, frameNum int, coverName string) string {
	buf := bytes.NewBuffer(nil)
	err := ffmpeg.Input(inFileName).
		Filter("select", ffmpeg.Args{fmt.Sprintf("gte(n,%d)", frameNum)}).
		Output("pipe:", ffmpeg.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}).
		WithOutput(buf, os.Stdout).
		Run()
	if err != nil {
		panic(err)
	}

	filePath := "public/video_cover/" + coverName + ".jpg"
	outFile, err := os.Create(filePath)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	defer outFile.Close()

	_, err = io.Copy(outFile, buf)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	SeverIp := "http://192.168.2.91"
	coverCul := SeverIp + ":8080" + "/static/video_cover/" + coverName + ".jpg"
	return coverCul
}
func AddVideo(title string, token string, file *multipart.FileHeader) error {
	// token不存在
	err := myjwt.FindToken(token)
	if err != nil {
		return err
	}
	// 解析token
	claim, err := myjwt.VerifyAction(token)
	if err != nil {
		return err
	}

	// 读取文件数据
	fileBytes, err := file.Open()
	if err != nil {
		return err
	}
	defer fileBytes.Close()

	// 存储到本地
	filename := fmt.Sprintf("%v", claim.UserID) + fmt.Sprintf("%v", rand.Int63())
	filePath := "public/video/" + filename + ".mp4"
	SeverIp := "http://192.168.2.91"
	playUrl := SeverIp + ":8080" + "/static/video/" + filename + ".mp4"
	outFile, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer outFile.Close()

	_, err = io.Copy(outFile, fileBytes)
	if err != nil {
		return err
	}

	coverUrl := GenerateVideoCover(filePath, 1, filename)
	if coverUrl == "" {
		return err
	}

	err = dao.Video_add(claim.UserID, title, playUrl, coverUrl)
	if err != nil {
		return err
	}

	return nil
}

func Feed(latestTime string, token string, videoList *[]Video) (int64, error) {
	// token不存在
	err := myjwt.FindToken(token)
	if err != nil {
		return 0, err
	}
	// 解析token
	_, err = myjwt.VerifyAction(token)
	if err != nil {
		return 0, err
	}

	if latestTime == "0" || latestTime == "" {
		latestTime = strconv.FormatInt(time.Now().Unix(), 10)
	}
	if len([]rune(latestTime)) > 11 {
		latestTime = latestTime[0:10]
	}

	err = dao.Feed(latestTime, videoList)
	if err != nil {
		return 0, err
	}

	nextTime := time.Now().Unix()
	for id := range *videoList {
		err = FullVideo(&(*videoList)[id], token)
		if err != nil {
			return 0, err
		}
		if (*videoList)[id].CreateTime < nextTime {
			nextTime = (*videoList)[id].CreateTime
		}
	}
	return nextTime, nil
}

func ListPublish(user_id int64, token string, videoList *[]Video) error {
	// token不存在
	err := myjwt.FindToken(token)
	if err != nil {
		return err
	}

	// 解析token
	_, err = myjwt.VerifyAction(token)
	if err != nil {
		return err
	}

	err = dao.Video_list(user_id, videoList)
	if err != nil {
		return err
	}

	for id := range *videoList {
		err = FullVideo(&(*videoList)[id], token)
		if err != nil {
			return err
		}
	}
	return nil
}

func FullVideo(video *Video, token string) error {
	err := UserInfo(strconv.FormatInt(video.Userid, 10), token, &(*video).Author)
	if err != nil {
		return err
	}

	// 补充IsFavorite
	video.IsFavorite, err = dao.Video_IsFavorite(video.Userid, video.Id)
	if err != nil {
		return err
	}

	// 补充FavoriteCount
	if n, err := myRedis.RdbVsF.Exists(myRedis.Ctx, strconv.FormatInt(video.Id, 10)).Result(); n > 0 {
		if err != nil {
			return err
		}
		video.FavoriteCount, err = myRedis.RdbVsF.Get(myRedis.Ctx, strconv.FormatInt(video.Id, 10)).Int64()
		if err != nil {
			return err
		}
	} else {
		video.FavoriteCount, err = dao.Video_VsF(video.Id)
		if err != nil {
			return err
		}
		myRedis.RdbVsF.Set(myRedis.Ctx, strconv.FormatInt(video.Id, 10), video.FavoriteCount, 0)
	}

	// 补充CommentCount
	if n, err := myRedis.RdbVsC.Exists(myRedis.Ctx, strconv.FormatInt(video.Id, 10)).Result(); n > 0 {
		if err != nil {
			return err
		}
		video.CommentCount, err = myRedis.RdbVsC.Get(myRedis.Ctx, strconv.FormatInt(video.Id, 10)).Int64()
		if err != nil {
			return err
		}
	} else {
		video.CommentCount, err = dao.Video_VsC(video.Id)
		if err != nil {
			return err
		}
		myRedis.RdbVsC.Set(myRedis.Ctx, strconv.FormatInt(video.Id, 10), video.CommentCount, 0)
	}

	return nil
}
