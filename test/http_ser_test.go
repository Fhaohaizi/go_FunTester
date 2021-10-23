package test

import (
	"fmt"
	"funtester/task"
	"github.com/buaazp/fasthttprouter"
	"github.com/facebookgo/grace/gracehttp"
	"github.com/gin-gonic/gin"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/valyala/fasthttp"
	"log"
	"net/http"
	"strings"
	"testing"
	"time"
)

type indexHandler struct {
	content string
}

func (ih *indexHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, ih.content)
}

func TestHttpSer(t *testing.T) {

	server := http.Server{
		Addr: ":8001",
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if strings.Index(r.URL.String(), "test") > 0 {
				fmt.Fprintf(w, "测试请求")
				return
			}
			fmt.Fprintf(w, task.FunTester)
			return
		}),
	}
	server.ListenAndServe()
	log.Println("开始创建HTTP服务")

}

func TestHttpServer23(t *testing.T) {
	router := gin.New()

	api := router.Group("/okreplay/api")
	{
		api.POST("/submit", gin.HandlerFunc(func(context *gin.Context) {
			context.ShouldBindJSON(map[string]interface{}{
				"code": 3,
				"msg":  "FunTester",
			})
		}))

	}
	s := &http.Server{
		Addr:           ":8001",
		Handler:        router,
		ReadTimeout:    1000 * time.Second,
		WriteTimeout:   1000 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
func TestHttpSer2(t *testing.T) {

	http.Handle("/test/", &indexHandler{content: "hello world!"})
	http.Handle("/", &indexHandler{content: task.FunTester})
	http.ListenAndServe(":8001", nil)
}
func TestHttpSer3(t *testing.T) {
	app := echo.New()
	app.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.DELETE, echo.POST, echo.OPTIONS, echo.PUT, echo.HEAD},
		AllowHeaders: []string{echo.HeaderContentType, echo.HeaderAuthorization},
	}))
	app.Group("/test")
	{
		projectGroup := app.Group("/property")
		projectGroup.POST("/create", PropertyAddHandler)
	}
	gracehttp.Serve()

}

func PropertyAddHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"msg": "FunTester",
	})
}
func TestFastSer(t *testing.T) {
	address := "127.0.0.1:3001"
	s := &fasthttp.Server{
		Handler: func(ctx *fasthttp.RequestCtx) {
			ctx.Response.SetBody([]byte(task.FunTester))
		}, // 注意这里
		Name: "FunTester server", // 服务器名称
	}

	router := fasthttprouter.New()
	router.GET("/test", func(ctx *fasthttp.RequestCtx) {
		ctx.Response.SetBody([]byte("get"))
	})
	//fasthttp.ListenAndServe(":12345",func(ctx *fasthttp.RequestCtx){
	//	ctx.Response.SetBody([]byte(FunTester))
	//})
	fasthttp.ListenAndServe(address, router.Handler)
	if err := s.ListenAndServe(address); err != nil {
		log.Fatal("error in ListenAndServe", err.Error())
	}

}
