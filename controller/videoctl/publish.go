package videoctl

import (
	"fmt"
	"github.com/2103561941/douyin/service/videosvc"
	"net/http"
	"path/filepath"

	"github.com/2103561941/douyin/controller/commonctl"
	"github.com/gin-gonic/gin"
)

func Publish(c *gin.Context) {
	token := c.PostForm("token")
	if _, ok := commonctl.UserLoginMap[token]; !ok {
		c.JSON(http.StatusOK, commonctl.Response{
			Status_code: -1,
			Status_msg:  "user is not login",
		})
		return
	}
	data, err := c.FormFile("data")
	if err != nil {
		c.JSON(http.StatusOK, commonctl.Response{
			Status_code: -1,
			Status_msg:  err.Error(),
		})
		return
	}
	localhost := "http://localhost:9999/douyin/index/"
	filename := filepath.Base(data.Filename)
	user := commonctl.UserLoginMap[token]
	title := c.PostForm("title")

	finalName := fmt.Sprintf("%d_%s", user.Id, filename)
	//需要判断同一用户上传同一个文件两次的情况
	saveFile := filepath.Join("./video_content/", finalName)
	if err := c.SaveUploadedFile(data, saveFile); err != nil {
		c.JSON(http.StatusOK, commonctl.Response{
			Status_code: -1,
			Status_msg:  err.Error(),
		})
		return
	}
	videoinfo := &videosvc.Publish_video{
		UserID:  user.Id,
		PlayURL: localhost + finalName,
		Title:   title,
	}

	if err := videoinfo.PublishVideo(); err != nil {
		c.JSON(http.StatusOK, commonctl.Response{
			Status_code: -1,
			Status_msg:  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, commonctl.Response{Status_code: 0})
	return
}