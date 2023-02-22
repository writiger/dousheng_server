package handler

import (
	"context"
	zaplog "dousheng_server/deploy/log"
	"dousheng_server/middleware"
	"dousheng_server/middleware/covermaker"
	"dousheng_server/rpc"
	usermodel "dousheng_server/user_service/dal/model"
	"dousheng_server/video_service/dal/model"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/hertz-contrib/jwt"
	"io"
	"mime/multipart"
	"os"
	"strconv"
	"strings"
	"time"
)

// Publish .
func Publish(ctx context.Context, c *app.RequestContext) {
	userGet, _ := c.Get("identity")
	user := userGet.(*usermodel.User)
	type fileParam struct {
		f     *multipart.FileHeader `form:"data"`
		token string                `form:"token"`
		title string                `form:"title"`
	}
	// 1. 验证参数
	req := fileParam{}
	err := c.Bind(&req)
	if err != nil {
		c.JSON(consts.StatusBadRequest, utils.H{
			"status_code": -1,
			"status_msg":  "wong param" + err.Error(),
		})
		return
	}
	// 获取文件后缀
	filenameArr := strings.Split(req.f.Filename, ".")
	suffix := filenameArr[len(filenameArr)-1]

	// 2. 调用服务
	uuid, errRpc := rpc.PublishVideo(&model.Video{
		UserID: user.UUID,
		// 向服务传递文件后缀
		// 通过UUID+后缀保存文件
		// 防止文件冲突
		PlayURL:       suffix,
		CoverURL:      "",
		FavoriteCount: 0,
		CommentCount:  0,
		Title:         req.title,
	})

	// 3.1 保存视频
	videoPath := fmt.Sprintf("static/videos/%d.%s", uuid, suffix)
	coverPath := fmt.Sprintf("static/covers/%d", uuid)
	file, errSave := os.OpenFile(videoPath, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 066)
	defer file.Close()
	fileSaved, _ := req.f.Open()
	io.Copy(file, fileSaved)
	// 3.2 保存封面
	_, errCover := covermaker.GetSnapshot(videoPath, coverPath, 1)
	// 4. 失败一个全部回调
	if errSave != nil || errRpc != nil || errCover != nil {
		_ = os.Remove(videoPath)
		_ = os.Remove(coverPath + ".png")
		_ = rpc.DeleteVideo(uuid)
		c.JSON(consts.StatusBadRequest, utils.H{
			"status_code": -1,
			"status_msg":  "rpc.PublishVideo wrong",
		})
		zaplog.ZapLogger.Errorf("save video failed errSave:%v, errRpc:%v, errCover:%v", errSave, errRpc, errCover)
		return
	}
	// 5. 返回
	c.JSON(consts.StatusOK, utils.H{
		"status_code": 0,
		"status_msg":  "success",
	})
}

// Feed .
func Feed(ctx context.Context, c *app.RequestContext) {
	// 1. 初始化可选参数
	var uuid int64
	lastTimeStamp := time.Now().UnixMilli()
	// 2. 通过时间戳生成time
	lastTimeStr := c.Query("latest_time")
	if lastTimeStr != "" {
		lastTimeStamp, _ = strconv.ParseInt(lastTimeStr, 10, 64)
		//lastTimeStamp = lastTimeStamp * 100
	}
	// 4. 通过token字符串解析uuid
	tokenString := c.Query("token")
	if tokenString != "" {
		requesterToken, _ := middleware.JwtMiddleware.ParseTokenString(tokenString)
		claim := jwt.ExtractClaimsFromToken(requesterToken)
		uuid = int64(claim["identity"].(float64))
	}
	// 5. 返回
	videos, nextTime, err := rpc.Feed(lastTimeStamp, uuid)
	if len(videos) < 3 {
		// 时间设置为当前时间即可完成循环
		nextTime = time.Now().UnixMilli()
	}
	if err != nil {
		c.JSON(consts.StatusBadRequest, utils.H{
			"status_code": -1,
			"status_msg":  err.Error(),
		})
		zaplog.ZapLogger.Errorf("rpc.Feed failed to visit err:%v", err)
		return
	}

	c.JSON(consts.StatusOK, utils.H{
		"status_code": 0,
		"status_msg":  "success",
		"next_time":   nextTime,
		"video_list":  videos,
	})
}

// VideoList .
func VideoList(ctx context.Context, c *app.RequestContext) {
	userIdStr := c.Query("user_id")
	userId, err := strconv.ParseInt(userIdStr, 10, 64)
	if err != nil {
		c.JSON(consts.StatusBadRequest, utils.H{
			"status_code": -1,
			"status_msg":  "wrong param" + err.Error(),
		})
		return
	}
	videos, err := rpc.VideoList(userId)
	if err != nil {
		c.JSON(consts.StatusBadRequest, utils.H{
			"status_code": -1,
			"status_msg":  err.Error(),
		})
		zaplog.ZapLogger.Errorf("rpc.VideoList failed to visit err:%v", err)
		return
	}
	c.JSON(consts.StatusOK, utils.H{
		"status_code": 0,
		"status_msg":  "success",
		"video_list":  videos,
	})
}

// VideoLike 点赞或取消点赞
func VideoLike(ctx context.Context, c *app.RequestContext) {
	// 1. 验证参数
	userGet, _ := c.Get("identity")
	videoGet := c.Query("video_id")
	actionTypeGet := c.Query("action_type")
	userId := userGet.(*usermodel.User).UUID
	videoId, err := strconv.ParseInt(videoGet, 10, 64)
	if err != nil {
		c.JSON(consts.StatusBadRequest, utils.H{
			"status_code": -1,
			"status_msg":  "wrong param : video_id" + err.Error(),
		})
		return
	}
	actionType, err := strconv.ParseInt(actionTypeGet, 10, 32)
	if err != nil || (actionType != 1 && actionType != 2) {
		c.JSON(consts.StatusBadRequest, utils.H{
			"status_code": -1,
			"status_msg":  "wrong param : action_type",
		})
		return
	}
	// 2. 调用rpc服务
	err = rpc.LikeVideo(userId, videoId, int32(actionType))
	if err != nil {
		c.JSON(consts.StatusBadRequest, utils.H{
			"status_code": -1,
			"status_msg":  "rpc.LikeVideo wrong" + err.Error(),
		})
		zaplog.ZapLogger.Errorf("rpc.LikeVideo failed to visit err:%v", err)
		return
	}
	// 3. 返回
	c.JSON(consts.StatusOK, utils.H{
		"status_code": 0,
		"status_msg":  "success",
	})
}

// FavoriteList .
func FavoriteList(ctx context.Context, c *app.RequestContext) {
	// 1. 验证参数
	userIdStr := c.Query("user_id")
	requester, _ := c.Get("identity")

	userId, err := strconv.ParseInt(userIdStr, 10, 64)
	tokenId := requester.(*usermodel.User).UUID
	if err != nil {
		c.JSON(consts.StatusBadRequest, utils.H{
			"status_code": -1,
			"status_msg":  "wrong param : video_id" + err.Error(),
		})
		return
	}
	// 2. 请求服务
	videos, err := rpc.FavoriteVideoList(tokenId, userId)
	if err != nil {
		c.JSON(consts.StatusBadRequest, utils.H{
			"status_code": -1,
			"status_msg":  "rpc.FavoriteVideoList wrong" + err.Error(),
		})
		zaplog.ZapLogger.Errorf("rpc.FavoriteVideoList failed to visit err:%v", err)
		return
	}
	c.JSON(consts.StatusOK, utils.H{
		"status_code": 0,
		"status_msg":  "success",
		"video_list":  videos,
	})
}

// CommentAction .
func CommentAction(ctx context.Context, c *app.RequestContext) {
	// 1. 验证参数
	userGet, _ := c.Get("identity")
	user := userGet.(*usermodel.User)
	videoGet := c.Query("video_id")
	videoId, err := strconv.ParseInt(videoGet, 10, 64)
	tokenId := user.UUID
	if err != nil {
		c.JSON(consts.StatusBadRequest, utils.H{
			"status_code": -1,
			"status_msg":  "wong param video_id" + err.Error(),
		})
		return
	}
	actionType := c.Query("action_type")

	switch actionType {
	case "1": // 2.1 发布
		// 2.2 验证评论文本
		content := c.Query("comment_text")
		if content == "" {
			c.JSON(consts.StatusBadRequest, utils.H{
				"status_code": -1,
				"status_msg":  "content is empty" + err.Error(),
			})
			return
		}
		// 2.3 调用rpc
		comment, err := rpc.PostComment(tokenId, model.Comment{
			UserId:  user.UUID,
			VideoId: videoId,
			Content: content,
		})
		// 2.4 返回
		if err != nil {
			c.JSON(consts.StatusBadRequest, utils.H{
				"status_code": -1,
				"status_msg":  "rpc.PostComment wrong" + err.Error(),
			})
			zaplog.ZapLogger.Errorf("rpc.postComment failed to visit err:%v", err)
			return
		}

		c.JSON(consts.StatusOK, utils.H{
			"status_code": 0,
			"status_msg":  "success",
			"comment":     comment,
		})

	case "2": // 2.1 删除
		// 2.2 验证评论uuid
		uuidStr := c.Query("comment_id")
		uuid, err := strconv.ParseInt(uuidStr, 10, 64)
		// 2.3 调用rpc
		err = rpc.DeleteComment(uuid)
		if err != nil {
			c.JSON(consts.StatusBadRequest, utils.H{
				"status_code": -1,
				"status_msg":  "rpc.DeleteComment wrong",
			})
			return
		}
		c.JSON(consts.StatusOK, utils.H{
			"status_code": 0,
			"status_msg":  "success",
			"comment":     nil,
		})

	default:
		c.JSON(consts.StatusBadRequest, utils.H{
			"status_code": -1,
			"status_msg":  "wong param action_type",
		})
		return
	}
}

// GetComment .
func GetComment(ctx context.Context, c *app.RequestContext) {
	// 1. 验证参数
	videoIdStr := c.Query("video_id")
	requester, _ := c.Get("identity")
	videoId, err := strconv.ParseInt(videoIdStr, 10, 64)
	tokenId := requester.(*usermodel.User).UUID
	if err != nil {
		c.JSON(consts.StatusBadRequest, utils.H{
			"status_code": -1,
			"status_msg":  "wrong param : video_id",
		})
		return
	}
	// 2. 调用服务
	comments, err := rpc.GetComment(tokenId, videoId)
	if err != nil {
		c.JSON(consts.StatusBadRequest, utils.H{
			"status_code": -1,
			"status_msg":  "rpc.GetComment wrong" + err.Error(),
		})
		zaplog.ZapLogger.Errorf("rpc.GetComment failed to visit err:%v", err)
		return
	}
	// 3. 返回
	c.JSON(consts.StatusOK, utils.H{
		"status_code":  0,
		"status_msg":   "success",
		"comment_list": comments,
	})
}
