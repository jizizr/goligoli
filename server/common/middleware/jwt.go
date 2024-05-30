package middleware

import (
	"context"
	"errors"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/golang-jwt/jwt/v4"
	"github.com/jizizr/goligoli/server/common/consts"
	"github.com/jizizr/goligoli/server/common/tools"
	"github.com/jizizr/goligoli/server/service/api/biz/errno"
	"strings"
)

func JWTAuth() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			errno.SendResponse(c, consts.CodeBadRequest, nil)
			c.Abort()
			return
		}
		tokenArr := strings.SplitN(token, " ", 2)
		if len(tokenArr) != 2 || tokenArr[0] != "Bearer" {
			errno.SendResponse(c, consts.CodeNeedLogin, nil)
			c.Abort()
			return
		}

		claims, err := tools.ParseToken(tokenArr[1])
		if err != nil {
			var ve *jwt.ValidationError
			if errors.As(err, &ve) && ve.Errors == jwt.ValidationErrorExpired {
				errno.SendResponse(c, consts.CodeTokenExpired, nil)
			} else {
				errno.SendResponse(c, consts.CodeTokenInvalid, nil)
			}
			c.Abort()
			return
		}
		c.Set("UID", claims.Uid)
		c.Next(ctx)
	}
}
