package handler

import (
	"context"
	"dousheng_server/rpc"
	usermodel "dousheng_server/user_service/dal/model"
	"dousheng_server/video_service/dal/model"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"io"
	"mime/multipart"
	"os"
	"strings"
)

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
			"status_msg":  "wong param",
		})
		return
	}
	// TODO 保证一致性
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
	if errRpc != nil {
		c.JSON(consts.StatusBadRequest, utils.H{
			"status_code": -1,
			"status_msg":  "rpc.PublishVideo wrong",
		})
		return
	}
	// 3. 保存视频
	path := fmt.Sprintf("%d.%s", uuid, suffix)
	file, errSave := os.OpenFile("static/videos/"+path, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 066)
	defer file.Close()
	fileSaved, _ := req.f.Open()
	io.Copy(file, fileSaved)
	// 4. 失败一个全部回调
	if errSave != nil || errRpc != nil {
		_ = os.Remove(path)
		_ = rpc.DeleteVideo(uuid)
		c.JSON(consts.StatusBadRequest, utils.H{
			"status_code": -1,
			"status_msg":  "save action failed , roll back",
		})
	}
	// 5. 返回
	c.JSON(consts.StatusOK, utils.H{
		"status_code": 0,
		"status_msg":  "success",
	})
}
