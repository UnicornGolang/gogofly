package middleware

import (
	"gogofly/api"
	global "gogofly/global/constants"
	"gogofly/model"
	"gogofly/service"
	"gogofly/utils"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	ERR_CODE_INVALID_TOKEN   = 10401
	ERR_CODE_TOKEN_PARSE     = 10402
	ERR_CODE_TOKEN_NOT_MATCH = 10403
	ERR_CODE_TOKEN_EXPIRED   = 10404
	ERR_CODE_TOKEN_RENEW     = 10405
	TOKEN_NAME               = "Authorization"
	TOKEN_PREFIX             = "Bearer "
	RENEW_TOKEN_DURATION     = 10 * 60 * time.Second
)

// 自定义一个鉴权的中间件
func Auth() func(c *gin.Context) {
	return func(c *gin.Context) {
		token := c.GetHeader(TOKEN_NAME)
		// Token 不存在直接返回
		if token == "" || !strings.HasPrefix(token, TOKEN_PREFIX) {
			global.Log.Error("token pattern error", token)
			tokenErr(c, ERR_CODE_INVALID_TOKEN)
			return
		}
		// 校验 token 的格式与签名没有被篡改
		token = token[len(TOKEN_PREFIX):]
		jwtClaims, err := utils.ParseToken(token)
		uid := jwtClaims.Id
		if err != nil || uid == 0 {
			tokenErr(c, ERR_CODE_TOKEN_PARSE)
			return
		}
		// 校验是否与 redis 中存储的是否一致
		rkey := strings.ReplaceAll(global.LOGIN_USRE_TOKEN_REDIS_PREFIX, "{ID}", strconv.Itoa(int(uid)))
		storeToken, err := global.RDB.Get(rkey)

		if err != nil || token != storeToken {
			tokenErr(c, ERR_CODE_TOKEN_NOT_MATCH)
			return
		}
		// 校验设置的 token 是否过期
		t, err := global.RDB.GetKeyTTL(rkey)
		if err != nil || t <= 0 {
			tokenErr(c, ERR_CODE_TOKEN_EXPIRED)
			return
		}
		if t.Seconds() < RENEW_TOKEN_DURATION.Seconds() {
			token, err = service.GenerateAndCacheLoginUserToken(uid, jwtClaims.Name)
			if err != nil {
				tokenErr(c, ERR_CODE_TOKEN_RENEW)
				return
			}
			c.Header("token", token)
		}

		// 使用数据库中查询对象，由于鉴权频繁操作，不适用
		// -----------------------------------------------
		// user, err := dao.NewUserDao().GetUserById(uid)
		// if err != nil {
		//   tokenErr(c)
		//   return
		// }
		// c.Set(global.LOGIN_USER, user)

		// 将登录的用户信息存储到上下文中，方便程序中其他地方使用
		c.Set(global.LOGIN_USER, model.LoginUser{
			Id:   uid,
			Name: jwtClaims.Name,
		})

		// 程序继续向下执行
		c.Next()
	}
}

func tokenErr(c *gin.Context, code int) {
	api.Fail(c, api.ResponseJson{
		Status: http.StatusUnauthorized,
		Code:   code,
		Msg:    "invalid token",
	})

}
