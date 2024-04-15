package middlewares

import (
	"log"
	"net/http"
	"strconv"
	"time"

	repostirories "github.com/fentezi/session-auth/internal/repositories"
	"github.com/fentezi/session-auth/pkg"
	"github.com/gin-gonic/gin"
)

func SessionMiddleware(repo *repostirories.Repositories) gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, err := c.Request.Cookie("session_id")
		if err != nil || cookie.Value == "" {
			c.Status(http.StatusUnauthorized)
			c.Abort()
			return
		}
		log.Println(cookie.MaxAge)
		if shouldExtendSession(cookie.MaxAge) {
			uuid := cookie.Value
			userID, err := repo.Redis.GetSession(uuid)
			if err != nil {
				c.Status(http.StatusUnauthorized)
				c.Abort()
				return
			}
			if err = repo.Redis.DeleteSession(uuid); err != nil {
				c.Status(http.StatusInternalServerError)
				c.Abort()
				return
			}
			uuid = pkg.GenerateUUID()
			uIntVal, _ := strconv.ParseUint(userID, 10, 0)
			err = repo.Redis.CreateSession(uuid, uint(uIntVal))
			if err != nil {
				c.Status(http.StatusInternalServerError)
				c.Abort()
				return
			}
			c.SetCookie("session_id", uuid, 600, "/", "localhost", false, true)
		}
		c.Next()
	}
}

func shouldExtendSession(maxAge int) bool {
	return 2*time.Minute > time.Duration(maxAge*int(time.Second))
}
