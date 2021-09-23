package middleware

import (
	"database/sql"
	"fmt"
	"sandbox-go-api-sqlboiler-rest-auth/internal/config"
	"sandbox-go-api-sqlboiler-rest-auth/models"
	"time"

	"github.com/gorilla/securecookie"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/labstack/echo/v4"
)

func SetCustomContext(cfg *config.Config, l *zap.SugaredLogger, db *sql.DB, sc *securecookie.SecureCookie) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := NewCustomContext(c, cfg, l, db, sc)
			return next(cc)
		}
	}
}

func SessionRestorer(db *sql.DB, logger *zap.SugaredLogger, sc *securecookie.SecureCookie) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := c.(*CustomContext)
			cv, err := c.Cookie("session")
			if err != nil {
				return next(c)
			}

			var dv string
			err = sc.Decode("session", cv.Value, &dv)
			if err != nil {
				return echo.NewHTTPError(500, "cannot decode cookie", err)
			}
			logger.Debug("got cookie(session id): ", dv)

			sess, err := models.FindSession(c.Request().Context(), db, dv)
			if err != nil {
				// maybe wrong cookie id?
				return echo.NewHTTPError(500, "cannot get cookie, but got session id", dv, err)
			}
			user, err := sess.User().One(c.Request().Context(), db)
			if err != nil {
				return echo.NewHTTPError(500, "cannod find user from session relation", dv, err)
			}
			logger.Debug("got user in middleware", user)
			cc.CurrentUser = user
			return next(c)
		}
	}
}

func RequestZapLogger(log *zap.Logger) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			start := time.Now()

			err := next(c)
			if err != nil {
				c.Error(err)
			}

			req := c.Request()
			res := c.Response()

			fields := []zapcore.Field{
				zap.Int("status", res.Status),
				zap.String("request", fmt.Sprintf("%s %s", req.Method, req.RequestURI)),
				zap.String("remote_ip", c.RealIP()),
				zap.String("latency", time.Since(start).String()),
				zap.String("host", req.Host),
				zap.Int64("size", res.Size),
				zap.String("user_agent", req.UserAgent()),
			}

			id := req.Header.Get(echo.HeaderXRequestID)
			if id == "" {
				id = res.Header().Get(echo.HeaderXRequestID)
				fields = append(fields, zap.String("request_id", id))
			}

			n := res.Status
			switch {
			case n >= 500:
				log.With(zap.Error(err)).Error("Server error", fields...)
			case n >= 400:
				log.Info("Client error", fields...)
			case n >= 300:
				log.Info("Redirection", fields...)
			default:
				log.Info("Success", fields...)
			}

			return nil
		}
	}
}
