package middlewares

import (
	"net/http"
	"strconv"
	"time"

	repostirories "github.com/fentezi/session-auth/internal/repositories"
	"github.com/fentezi/session-auth/pkg"
	"github.com/gin-gonic/gin"
)

func SessionMiddleware(repo *repostirories.Repositories) gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, err := c.Cookie("session_id")
		if err != nil || cookie == "" {
			c.Status(http.StatusUnauthorized)
			c.Abort()
			return
		}
		timeSession, err := repo.Redis.GetTimeKey(cookie)
		if err != nil || timeSession == -2 {
			c.Status(http.StatusUnauthorized)
			c.Abort()
			return
		}
		if shouldExtendSession(timeSession) {
			uuid := cookie
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

func shouldExtendSession(timeSession time.Duration) bool {
	return timeSession < 2*time.Minute
}
