package server

import (
	"facegram_file_server/api/clientApiV1"
	"facegram_file_server/api/middleware"
	"facegram_file_server/pkg/utility"
	"strconv"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func RunApiServer() {

	setAppMode()

	srv := gin.Default()

	srv.Static("/storage/css", "./templates/resources/css")
	srv.Static("/storage/js", "./templates/resources/js")
	srv.Static("/storage/fonts", "./templates/resources/fonts")

	srv.LoadHTMLGlob("./templates/view/*.html")

	apiV1 := srv.Group("/v1")

	apiV1.Use(middleware.AllRouteRequestLimiter) //request limiter for all route

	//https://domain.com/v1/upload
	up := apiV1.Group("/upload")
	{
		up.Use(middleware.UploadRouteCheckAccess)
		up.POST("/new", clientApiV1.NewUpload)
	}

	//https://domain.com/v1/download
	dl := apiV1.Group("/download")
	{
		dl.POST("/link", clientApiV1.GetLinks)      // get direct-url file from storage
		dl.POST("/preview", clientApiV1.GetPreview) // get file preview
		dl.POST("/archive", clientApiV1.GetArchive) // get user file list
	}

	////https://domain.com/v1/documents
	docs := apiV1.Group("/documents")
	{
		docs.Use(middleware.DocumentRouteCheckAccess)
		docs.GET("/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	//employee authentication for document access
	//https://domain.com/v1/auth
	auth := apiV1.Group("/auth")
	{
		auth.GET("/", clientApiV1.LoginForm)
		auth.POST("/", clientApiV1.LoginHandler)
	}

	port := utility.GetEnv("APP_PORT", ":80")

	err := srv.Run(*port)
	if err != nil {
		panic(err)
	}
}

func setAppMode() {

	debug, err := strconv.ParseBool(*utility.GetEnv("APP_DEBUG", "false"))

	if err != nil || !debug {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	return
}
