package middlewares

import (
	"net/http"

	repostirories "github.com/fentezi/session-auth/internal/repositories"
	"github.com/gin-gonic/gin"
)

func SignInMiddleware(repo *repostirories.Repositories) gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, err := c.Cookie("session_id")
		if err != nil || cookie == "" {
			c.Next()
			return
		}

		userID, err := repo.Redis.GetSession(cookie)
		if err != nil || userID == "" {
			c.Next()
			return
		}
		c.Redirect(http.StatusFound, "/")
		c.Abort()
	}
}
