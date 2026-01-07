package router

import (
    "net/http"
    _ "strconv"
    "anb/controllers/api"

    _ "anb/models"
    "github.com/gin-gonic/gin"
)

func SetRouter(r *gin.Engine) {
    apiGroup := r.Group("/api")
    //apiGroup.Use(TokenRequired())

    {
        apiGroup.Any("/picture/list", func(c *gin.Context) {
            var controller api.PictureController

            controller.Init(c)

            controller.AjaxList()
            controller.Close()
            c.JSON(http.StatusOK, controller.Result)
        })

        apiGroup.Any("/picture/upload", func(c *gin.Context) {
            var controller api.PictureController

            controller.Init(c)

            controller.AjaxUpload()
            controller.Close()
            c.JSON(http.StatusOK, controller.Result)
        })

        apiGroup.Any("/picture/insert", func(c *gin.Context) {
            var controller api.PictureController

            controller.Init(c)

            controller.AjaxInsert()
            controller.Close()
            c.JSON(http.StatusOK, controller.Result)
        })

        apiGroup.Any("/sync/insert", func(c *gin.Context) {
            var controller api.SyncController

            controller.Init(c)

            controller.AjaxInsert()
            controller.Close()
            c.JSON(http.StatusOK, controller.Result)
        })

        apiGroup.Any("/sync/sync", func(c *gin.Context) {
            var controller api.SyncController

            controller.Init(c)

            controller.AjaxSync()
            controller.Close()
            c.JSON(http.StatusOK, controller.Result)
        })

        apiGroup.Any("/sync/download", func(c *gin.Context) {
            var controller api.SyncController

            controller.Init(c)

            controller.AjaxDownload()
            controller.Close()
            c.JSON(http.StatusOK, controller.Result)
        })

        apiGroup.Any("/sync/upload", func(c *gin.Context) {
            var controller api.SyncController

            controller.Init(c)

            controller.AjaxUpload()
            controller.Close()
            c.JSON(http.StatusOK, controller.Result)
        })

        apiGroup.Any("/sync/complete", func(c *gin.Context) {
            var controller api.SyncController

            controller.Init(c)

            controller.AjaxComplete()
            controller.Close()
            c.JSON(http.StatusOK, controller.Result)
        })

        apiGroup.Any("/imagefloor/get", func(c *gin.Context) {
            var controller api.ImagefloorController

            controller.Init(c)

            controller.AjaxGet()
            controller.Close()
            c.JSON(http.StatusOK, controller.Result)
        })

        apiGroup.Any("/apt/list", func(c *gin.Context) {
            var controller api.AptController

            controller.Init(c)

            controller.AjaxList()
            controller.Close()
            c.JSON(http.StatusOK, controller.Result)
        })

        apiGroup.Any("/data/list", func(c *gin.Context) {
            var controller api.DataController

            controller.Init(c)

            controller.AjaxList()
            controller.Close()
            c.JSON(http.StatusOK, controller.Result)
        })

        apiGroup.Any("/data/delete", func(c *gin.Context) {
            var controller api.DataController

            controller.Init(c)

            controller.AjaxDelete()
            controller.Close()
            c.JSON(http.StatusOK, controller.Result)
        })

        apiGroup.Any("/data/insert", func(c *gin.Context) {
            var controller api.DataController

            controller.Init(c)

            controller.AjaxInsert()
            controller.Close()
            c.JSON(http.StatusOK, controller.Result)
        })

        apiGroup.Any("/data/image", func(c *gin.Context) {
            var controller api.DataController

            controller.Init(c)

            controller.AjaxImage()
            controller.Close()
            c.JSON(http.StatusOK, controller.Result)
        })

        apiGroup.Any("/data/multiinsert", func(c *gin.Context) {
            var controller api.DataController

            controller.Init(c)

            controller.AjaxMultiinsert()
            controller.Close()
            c.JSON(http.StatusOK, controller.Result)
        })

        apiGroup.Any("/status/list", func(c *gin.Context) {
            var controller api.StatusController

            controller.Init(c)

            controller.AjaxList()
            controller.Close()
            c.JSON(http.StatusOK, controller.Result)
        })

        apiGroup.Any("/image/list", func(c *gin.Context) {
            var controller api.ImageController

            controller.Init(c)

            controller.AjaxList()
            controller.Close()
            c.JSON(http.StatusOK, controller.Result)
        })

        apiGroup.Any("/image/insert", func(c *gin.Context) {
            var controller api.ImageController

            controller.Init(c)

            controller.AjaxInsert()
            controller.Close()
            c.JSON(http.StatusOK, controller.Result)
        })

        apiGroup.Any("/image/delete", func(c *gin.Context) {
            var controller api.ImageController

            controller.Init(c)

            controller.AjaxDelete()
            controller.Close()
            c.JSON(http.StatusOK, controller.Result)
        })

        apiGroup.Any("/image/view", func(c *gin.Context) {
            var controller api.ImageController

            controller.Init(c)

            controller.AjaxView()
            controller.Close()
            c.JSON(http.StatusOK, controller.Result)
        })

        apiGroup.Any("/image/upload", func(c *gin.Context) {
            var controller api.ImageController

            controller.Init(c)

            controller.AjaxUpload()
            controller.Close()
            c.JSON(http.StatusOK, controller.Result)
        })

        apiGroup.Any("/image/upload2", func(c *gin.Context) {
            var controller api.ImageController

            controller.Init(c)

            controller.AjaxUpload2()
            controller.Close()
            c.JSON(http.StatusOK, controller.Result)
        })

        apiGroup.Any("/image/uploadandroidbase64", func(c *gin.Context) {
            var controller api.ImageController

            controller.Init(c)

            controller.AjaxUploadandroidbase64()
            controller.Close()
            c.JSON(http.StatusOK, controller.Result)
        })

        apiGroup.Any("/image/uploadandroid", func(c *gin.Context) {
            var controller api.ImageController

            controller.Init(c)

            controller.AjaxUploadandroid()
            controller.Close()
            c.JSON(http.StatusOK, controller.Result)
        })

        apiGroup.GET("/report", func(c *gin.Context) {
            var controller api.ReportController

            controller.Init(c)
            controller.Set("current", "/api/report")
            controller.Current = "/api/report"
            controller.Index()


            controller.Set("self", "report")
            controller.Display("/api/report.html")
            controller.Close()
        })

        apiGroup.GET("/report/download/hwp", func(c *gin.Context) {
            var controller api.ReportController
            controller.Init(c)

            controller.DownloadHwp()

            controller.Set("self", "report")
            controller.Close()
        })

        apiGroup.Any("/report/summary", func(c *gin.Context) {
            var controller api.ReportController

            controller.Init(c)

            controller.AjaxSummary()
            controller.Close()
            c.JSON(http.StatusOK, controller.Result)
        })

        apiGroup.Any("/json/list", func(c *gin.Context) {
            var controller api.JsonController

            controller.Init(c)

            controller.AjaxList()
            controller.Close()
            c.JSON(http.StatusOK, controller.Result)
        })

        apiGroup.Any("/json/data", func(c *gin.Context) {
            var controller api.JsonController

            controller.Init(c)

            controller.AjaxData()
            controller.Close()
            c.JSON(http.StatusOK, controller.Result)
        })

        apiGroup.Any("/json/check", func(c *gin.Context) {
            var controller api.JsonController

            controller.Init(c)

            controller.AjaxCheck()
            controller.Close()
            c.JSON(http.StatusOK, controller.Result)
        })

        apiGroup.Any("/aptgroup/get", func(c *gin.Context) {
            var controller api.AptgroupController

            controller.Init(c)

            controller.AjaxGet()
            controller.Close()
            c.JSON(http.StatusOK, controller.Result)
        })

        apiGroup.Any("/login/login", func(c *gin.Context) {
            var controller api.LoginController

            controller.Init(c)

            controller.AjaxLogin()
            controller.Close()
            c.JSON(http.StatusOK, controller.Result)
        })
    }
}
