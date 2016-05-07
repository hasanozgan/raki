package context

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hasanozgan/vodka"
	"github.com/hasanozgan/vodka/examples/basic/services"
	"golang.org/x/net/context"
)

type AppContext struct {
	ctx               *gin.Context
	dbClientFactory   *vodka.DbClientFactory
	httpClientFactory *vodka.HttpClientFactory
	loggerFactory     *vodka.LoggerFactory
}

func NewAppContext(ctx *gin.Context) context.Context {
	fields := map[string]interface{}{}
	fields["X-Correlation-ID"] = ctx.Request.Header["X-Correlation-ID"]

	loggerFactory := vodka.NewLoggerFactoryWithFields(fields)
	dbClientFactory := vodka.NewDbClientFactory(loggerFactory)
	httpClientFactory := vodka.NewHttpClientFactory(loggerFactory)

	appContext := AppContext{ctx, dbClientFactory, httpClientFactory, loggerFactory}
	return appContext
}

func (a *AppContext) IsLoggedIn() bool {
	return false
}

func (a *AppContext) UserInfo() *UserInfo {
	return &UserInfo{}
}

func (a *AppContext) Services() *services.Services {
	return services.New(a.dbClientFactory, a.httpClientFactory, a.loggerFactory)
}

func (a *AppContext) NewLogger(name string) vodka.Logger {
	return a.loggerFactory.NewLogger(name)
}

/************************************/
/***** GOLANG.ORG/X/NET/CONTEXT *****/
/************************************/

func (c AppContext) Deadline() (deadline time.Time, ok bool) {
	return
}

func (c AppContext) Done() <-chan struct{} {
	return nil
}

func (c AppContext) Err() error {
	return nil
}

func (c AppContext) Value(key interface{}) interface{} {
	return nil
}
