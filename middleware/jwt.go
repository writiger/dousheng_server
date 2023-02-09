package middleware

import (
	"context"
	"dousheng_server/rpc"
	"dousheng_server/user_service/dal/model"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/hertz-contrib/jwt"
	"net/http"
	"time"
)

var JwtMiddleware *jwt.HertzJWTMiddleware

func init() {
	var err error
	JwtMiddleware, err = jwt.New(&jwt.HertzJWTMiddleware{
		Key:         []byte("revres_gnehsoud"),
		Timeout:     time.Hour * 7 * 24,
		MaxRefresh:  time.Hour * 7 * 24,
		TokenLookup: "query: token, header: Authorization, cookie: jwt",
		// 解析
		IdentityHandler: func(ctx context.Context, c *app.RequestContext) interface{} {
			claims := jwt.ExtractClaims(ctx, c)
			return &model.User{
				UUID: int64(claims[jwt.IdentityKey].(float64)),
			}
		},
		// 负载信息
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(int64); ok {
				return jwt.MapClaims{
					jwt.IdentityKey: v,
				}
			}
			return jwt.MapClaims{}
		},
		// 验证
		Authenticator: func(ctx context.Context, c *app.RequestContext) (interface{}, error) {
			username := c.Query("username")
			password := c.Query("password")
			if len(username) == 0 || len(password) == 0 || len(username) > 32 || len(password) > 32 {
				return "", jwt.ErrMissingLoginValues
			}
			uuid, err := rpc.LoginByPassword(username, password)
			c.Set("uuid", uuid)
			return uuid, err
		},
		LoginResponse: func(ctx context.Context, c *app.RequestContext, code int, token string, expire time.Time) {
			uuid, _ := c.Get("uuid")
			c.JSON(http.StatusOK, utils.H{
				"status_code": 0,
				"status_msg":  "success",
				"user_id":     uuid,
				"token":       token,
			})
		},
		Unauthorized: func(ctx context.Context, c *app.RequestContext, code int, message string) {
			c.JSON(http.StatusOK, utils.H{
				"code":    -1,
				"message": message,
			})
		},
	})
	if err != nil {
		panic("init jwt failed")
	}
}
