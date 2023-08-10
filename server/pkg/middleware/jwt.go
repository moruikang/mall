package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"mall.com/config/global"
	"mall.com/pkg/response"
	"mall.com/store/models"
	"net/http"
	"strconv"

	jwt "github.com/appleboy/gin-jwt/v2"
	//jwtauth "github.com/golang-jwt/jwt"
	"mall.com/pkg/common"
	"mall.com/service"
	"time"
)

// JwtAuth JWT认证中间件
func JwtAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token == "" {
			response.Failed("未登录或非法访问", c)
			c.Abort()
			return
		}
		if err := common.VerifyToken(token); err != nil {
			response.Failed("登录已过期，请重新登录", c)
			c.Abort()
			return
		}
		c.Next()
	}
}


func NewGinJwtMiddlewares() (*jwt.GinJWTMiddleware, error) {
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "mall",
		Key:         []byte("demo"),
		Timeout:     10 * time.Hour,
		MaxRefresh:  10 * time.Hour,
		IdentityKey: global.Config.Jwt.SigningKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*models.User); ok {
				return jwt.MapClaims{
					//"id":   v.Id,
					"username": v.Username,
				}
			}
			return jwt.MapClaims{}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {

			var param models.WebUserLoginParam
			if err := c.ShouldBind(&param); err != nil {
				return "", jwt.ErrMissingLoginValues
			}
			// 检查验证码
			if !common.VerifyCaptcha(param.CaptchaId, param.CaptchaValue) {
				return "", errors.New("验证码错误")
			}
			// 校验用户信息是否正确

			userinfo, err := service.GetUser(param)
			if err != nil  {
				return nil, jwt.ErrFailedAuthentication
			}

			c.Request.Header.Set("userID", strconv.FormatUint(userinfo.Id, 10))
			return userinfo, nil
		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			if _, ok := data.(*models.User); ok {
				return true
			}
			return true
		},
		Unauthorized: func(c *gin.Context, code int, message string) {

			c.JSON(
				http.StatusOK,
				response.Response{
					400,
					message,
					code
				},
				)
			//c.JSON(code, gin.H{
			//	"code":    code,
			//	"message": message,
			//})
		},
		TokenLookup: "header: Authorization, query: token, cookie: jwt",
		TokenHeadName: "Bearer",
		LoginResponse: func(c *gin.Context, code int, message string, time time.Time) {
			if code == 200 {
				c.JSONP(code, gin.H{
					"code": code,
					"data": map[string]interface{}{
						"uid":   c.Request.Header.Get("userID"),
						"token": message,
					},
					"message":"登录成功",
				})
			} else {
				c.JSONP(code, gin.H{
					"code":code,
					"data":         "",
					"message": "登录失败",
				})
			}

		},
		SendCookie: true,
	})
	return authMiddleware, err
}

