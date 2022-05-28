package usersvc

import (
	"errors"
	"github.com/2103561941/douyin/repository"
)

type UserFollow struct {
	Id          string
	User_id     string
	To_user_id  string
	Action_type string
}

func (user *UserFollow) Follow() error {
	status := &repository.Follow{
		UserId:     user.User_id,
		FollowerId: user.To_user_id,
		Status:     user.Action_type,
	}

	if err := status.CheckStatus(); err != nil {
		//报错应该是因为没有找到。所以建立对应行 不知道是这个意思不
		/*
			INSERT INTO follow(user_id, follower_id, status)
			VALUES
			(user.User_id, user.To_user_id, user.Action_type);
		*/

		if err := status.Insert(); err != nil {
			return err
		}

	}

	switch status.Status {
	case "0":
		// not followed
		/*
			UPDATE follow
			SET status = "1"
			WHERE User_id = user.User_id;
		*/
		if user.Action_type == "1" {
			if err := status.UpdateStatus("1"); err != nil {
				return err
			}
		}
		if user.Action_type == "2" {
			return errors.New("you didn't follow this user")
		}

	case "1":
		// already followed this user. 已经关注该用户
		if user.Action_type == "1" {
			return errors.New("you already followed this user")
		}
		if user.Action_type == "2" {
			if err := status.UpdateStatus("0"); err != nil {
				return err
			}
		}
	case "2":
		// already followed by this user 已经被该用户关注
		/*
			UPDATE follow
			SET status = "3"
			WHERE User_id = user.User_id;
		*/
		if user.Action_type == "1" {
			if err := status.UpdateStatus("3"); err != nil {
				return err
			}
		}

		if user.Action_type == "2" {
			return errors.New("you didn't follow this user")
		}

	case "3":
		// already mutual following 已经互粉
		if user.Action_type == "1" {
			return errors.New("you already mutual following with this user")
		}
		if user.Action_type == "2" {
			if err := status.UpdateStatus("2"); err != nil {
				return err
			}
		}
	default:
	}

	return nil
}