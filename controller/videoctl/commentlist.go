package videoctl

import (
	"github.com/2103561941/douyin/controller/commonctl"
	"github.com/2103561941/douyin/service/videosvc"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

type Commentlistresponse struct {
	commonctl.Response
	Infos []*videosvc.CommentResponseWrapper `json:"comment_list"`
}

func GetCommentList(c *gin.Context) {

	videoIDraw, err := strconv.Atoi(c.Query("video_id"))
	if err != nil {
		c.JSON(http.StatusOK, commonctl.Response{
			Status_code: -1,
			Status_msg:  "video_id is undefined",
		})
		return
	}
	videoID := uint64(videoIDraw) //视频ID

	testcal, boolen := c.Get("middleware_geted_user_id")
	if boolen == false {
		log.Println("user_page didn't get")
	}

	userId := testcal.(uint64)

	//userId := commonctl.UserLoginMap[token].Id // 主动去访问的用户id

	list := videosvc.CommentList{
		VideoID: videoID,
		UserID:  userId,
	}
	if err := list.GetCommentList(); err != nil {
		c.JSON(http.StatusOK, commonctl.Response{
			Status_code: -1,
			Status_msg:  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, Commentlistresponse{
		Response: commonctl.Response{Status_code: 0},
		Infos:    list.Videos,
	})

}
