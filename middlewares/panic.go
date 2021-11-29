package middlewares

import (
	"fmt"
	"log"
	"os"
	"reflect"
	"runtime/debug"

	"github.com/getsentry/sentry-go"
	"github.com/gin-gonic/gin"
)

func init() {
	if err := sentry.Init(sentry.ClientOptions{
		Dsn:              os.Getenv("SENTRY"),
		Environment:      os.Getenv("ENV"),
		Release:          os.Getenv("RELEASE"),
		TracesSampleRate: 1.0,
	}); err != nil {
		log.Println("Setup Sentry Failed")
	}
}

func Panic(c *gin.Context) {
	if r := recover(); r != nil {
		if hub := sentry.GetHubFromContext(c.Request.Context()); hub != nil {
			fmt.Printf("PANIC ERR: %+v\n", r)
			debug.PrintStack()
			switch v := r.(type) {
			case error:
				sentry.CaptureException(v)
			case string:
				sentry.CaptureMessage(v)
			default:
				fmt.Printf("PANIC TYPE: %+v\n", reflect.TypeOf(r))
			}
		}
	}
}
