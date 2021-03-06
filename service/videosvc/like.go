package videosvc

import (
	"errors"

	"github.com/2103561941/douyin/repository"
)

type Like struct {
	UserId     uint64
	VideoId    uint64
	ActionType int
}

func (action *Like) Like() error {

	vidinfo := &repository.Video{
		Id:     action.VideoId,
		UserId: action.UserId,
	}

	if err := vidinfo.GetLikeInfo(); err != nil {
		return err
	}

	addlikeinfo := &repository.LikeTable{
		UserId:  action.UserId,
		VideoId: vidinfo.Id,
	}

	// 获取like表的信息
	if err := addlikeinfo.GetLikeInfoinLike(); err != nil {
		return err
	}

	if action.ActionType == 1 { //点赞
		if addlikeinfo.ActionType == 1 {
			return errors.New("you can not like it again")
		}
		if err := vidinfo.Like(vidinfo); err != nil {
			return err
		}
	} else if action.ActionType == 2 { //取消点赞
		if addlikeinfo.ActionType == 0 {
			return errors.New("you can not unlike it again")
		}
		if err := vidinfo.UnLike(vidinfo); err != nil {
			return err
		}
	} else {
		return errors.New("invalid like action")
	}

	if err := addlikeinfo.UpdateLike(action.ActionType); err != nil {
		return err
	}
	//请查看UpdateLike函数的注释
	return nil
}
