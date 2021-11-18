package test

import (
	"fmt"
	"funtester/funtester"
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
				fmt.Fprintf(w, "这是net/http创建的server第一种方式")
				return
			}
			fmt.Fprintf(w, funtester.FunTester)
			return
		}),
	}
	server.ListenAndServe()
	log.Println("开始创建HTTP服务")

}

type Argsss struct {
	Code int
	msg  string
}

func TestHttpServer4(t *testing.T) {
	router := gin.New()

	api := router.Group("/test")
	{
		api.GET("/fun", gin.HandlerFunc(func(context *gin.Context) {
			context.JSON(http.StatusOK, funtester.FunTester)
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
	http.Handle("/test", &indexHandler{content: "这是net/http第二种创建服务语法"})
	http.Handle("/", &indexHandler{content: funtester.FunTester})
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
		projectGroup := app.Group("/test")
		projectGroup.GET("/", PropertyAddHandler)
	}
	app.Server.Addr = ":8001"
	gracehttp.Serve(app.Server)

}

func PropertyAddHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"msg":  "这是net/http第三种创建服务的语法",
		"code": 888,
	})
}
func TestFastSer(t *testing.T) {
	address := ":8001"
	handler := func(ctx *fasthttp.RequestCtx) {
		path := string(ctx.Path())
		switch path {
		case "/test":
			ctx.SetBody([]byte("这是fasthttp创建服务的第一种语法"))
		default:
			ctx.SetBody([]byte(funtester.FunTester))
		}
	}
	s := &fasthttp.Server{
		Handler: handler,
		Name:    "FunTester server",
	}

	if err := s.ListenAndServe(address); err != nil {
		log.Fatal("error in ListenAndServe", err.Error())
	}

}
func TestFastSer2(t *testing.T) {
	address := ":8001"

	router := fasthttprouter.New()
	router.GET("/test", func(ctx *fasthttp.RequestCtx) {
		ctx.Response.SetBody([]byte("这是fasthttp创建server的第二种语法"))
	})
	router.GET("/", func(ctx *fasthttp.RequestCtx) {
		ctx.Response.SetBody([]byte(funtester.FunTester))
	})
	fasthttp.ListenAndServe(address, router.Handler)
}
