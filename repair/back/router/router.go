package router

import (
    "encoding/json"
    "strconv"
    "strings"
    "net/http"
	"repair/controllers/api"
	"repair/controllers/rest"
    "repair/models"
    "repair/models/repairlist"
    "repair/models/user"
	"github.com/gin-gonic/gin"
)

func getArrayCommal(name string) []int64 {
	values := strings.Split(name, ",")

	var items []int64
	for _, item := range values {
        n, _ := strconv.ParseInt(item, 10, 64)
		items = append(items, n)
	}

	return items
}

func getArrayCommai(name string) []int {
	values := strings.Split(name, ",")

	var items []int
	for _, item := range values {
        n, _ := strconv.Atoi(item)
		items = append(items, n)
	}

	return items
}

func SetRouter(r *gin.Engine) {

    r.GET("/api/jwt", func(c *gin.Context) {
		loginid := c.Query("loginid")
        passwd := c.Query("passwd")
        c.JSON(http.StatusOK, JwtAuth(c, loginid, passwd))
	})
	apiGroup := r.Group("/api")
	apiGroup.Use(JwtAuthRequired())
	{

		apiGroup.GET("/apt/search", func(c *gin.Context) {
			var controller api.AptController
			controller.Init(c)
			controller.Search()
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.POST("/aptdong/blueprint", func(c *gin.Context) {
			item_ := &models.Aptdongblueprint{}
			c.ShouldBindJSON(item_)
			var controller api.AptdongController
			controller.Init(c)
			controller.Blueprint(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.POST("/breakdown/deduplication", func(c *gin.Context) {
			var results map[string]any
			jsonData, _ := c.GetRawData()
			json.Unmarshal(jsonData, &results)
			var apt_ int64
			if v, flag := results["apt"]; flag {
				apt_ = int64(v.(float64))
			}
			var controller api.BreakdownController
			controller.Init(c)
			controller.Deduplication(apt_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.POST("/breakdown/lastdate", func(c *gin.Context) {
			var results map[string]any
			jsonData, _ := c.GetRawData()
			json.Unmarshal(jsonData, &results)
			var date_ int
			if v, flag := results["date"]; flag {
				date_ = int(v.(float64))
			}
			var ids_ []int64
			if v, flag := results["ids"]; flag {
				ids_= getArrayCommal(v.(string))
			}
			var controller api.BreakdownController
			controller.Init(c)
			controller.UpdateLastdate(date_, ids_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.POST("/breakdown/duedate", func(c *gin.Context) {
			var results map[string]any
			jsonData, _ := c.GetRawData()
			json.Unmarshal(jsonData, &results)
			var apt_ int64
			if v, flag := results["apt"]; flag {
				apt_ = int64(v.(float64))
			}
			var date_ int
			if v, flag := results["date"]; flag {
				date_ = int(v.(float64))
			}
			var ids_ []int64
			if v, flag := results["ids"]; flag {
				ids_= getArrayCommal(v.(string))
			}
			var controller api.BreakdownController
			controller.Init(c)
			controller.UpdateDuedate(apt_, date_, ids_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/breakdownhistory/auto/:id", func(c *gin.Context) {
			id_, _ := strconv.ParseInt(c.Param("id"), 10, 64)
			var controller api.BreakdownhistoryController
			controller.Init(c)
			controller.Auto(id_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.POST("/category/initdata", func(c *gin.Context) {
			var results map[string]any
			jsonData, _ := c.GetRawData()
			json.Unmarshal(jsonData, &results)
			var apt_ int64
			if v, flag := results["apt"]; flag {
				apt_ = int64(v.(float64))
			}
			var controller api.CategoryController
			controller.Init(c)
			controller.InitData(apt_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.POST("/category/duplicationdata", func(c *gin.Context) {
			var results map[string]any
			jsonData, _ := c.GetRawData()
			json.Unmarshal(jsonData, &results)
			var apt_ int64
			if v, flag := results["apt"]; flag {
				apt_ = int64(v.(float64))
			}
			var controller api.CategoryController
			controller.Init(c)
			controller.DuplicationData(apt_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.POST("/detail/duplication", func(c *gin.Context) {
			var results map[string]any
			jsonData, _ := c.GetRawData()
			json.Unmarshal(jsonData, &results)
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller api.DetailController
			controller.Init(c)
			controller.Duplication(id_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/download/file/:id", func(c *gin.Context) {
			id_, _ := strconv.ParseInt(c.Param("id"), 10, 64)
			var controller api.DownloadController
			controller.Init(c)
			controller.File(id_)
			controller.Close()
		})

		apiGroup.GET("/download/report/:id", func(c *gin.Context) {
			id_, _ := strconv.ParseInt(c.Param("id"), 10, 64)
			var controller api.DownloadController
			controller.Init(c)
			controller.Report(id_)
			controller.Close()
		})

		apiGroup.GET("/download/address", func(c *gin.Context) {
			var controller api.DownloadController
			controller.Init(c)
			controller.Address()
			controller.Close()
		})

		apiGroup.GET("/download/periodic0/:id", func(c *gin.Context) {
			id_, _ := strconv.ParseInt(c.Param("id"), 10, 64)
			var controller api.DownloadController
			controller.Init(c)
			controller.Periodic0(id_)
			controller.Close()
		})

		apiGroup.GET("/download/periodic1/:id", func(c *gin.Context) {
			id_, _ := strconv.ParseInt(c.Param("id"), 10, 64)
			var controller api.DownloadController
			controller.Init(c)
			controller.Periodic1(id_)
			controller.Close()
		})

		apiGroup.GET("/download/periodic2/:id", func(c *gin.Context) {
			id_, _ := strconv.ParseInt(c.Param("id"), 10, 64)
			var controller api.DownloadController
			controller.Init(c)
			controller.Periodic2(id_)
			controller.Close()
		})

		apiGroup.GET("/download/periodic3/:id", func(c *gin.Context) {
			id_, _ := strconv.ParseInt(c.Param("id"), 10, 64)
			var controller api.DownloadController
			controller.Init(c)
			controller.Periodic3(id_)
			controller.Close()
		})

		apiGroup.GET("/download/periodic4/:id", func(c *gin.Context) {
			id_, _ := strconv.ParseInt(c.Param("id"), 10, 64)
			var controller api.DownloadController
			controller.Init(c)
			controller.Periodic4(id_)
			controller.Close()
		})

		apiGroup.GET("/download/periodic5/:id", func(c *gin.Context) {
			id_, _ := strconv.ParseInt(c.Param("id"), 10, 64)
			var controller api.DownloadController
			controller.Init(c)
			controller.Periodic5(id_)
			controller.Close()
		})

		apiGroup.GET("/download/periodic/:id", func(c *gin.Context) {
			id_, _ := strconv.ParseInt(c.Param("id"), 10, 64)
			var controller api.DownloadController
			controller.Init(c)
			controller.Periodic(id_)
			controller.Close()
		})

		apiGroup.GET("/download/estimate/:id", func(c *gin.Context) {
			id_, _ := strconv.ParseInt(c.Param("id"), 10, 64)
			type_, _ := strconv.Atoi(c.Query("type"))
			var controller api.DownloadController
			controller.Init(c)
			controller.Estimate(id_, type_)
			controller.Close()
		})

		apiGroup.GET("/download/contract/:id", func(c *gin.Context) {
			id_, _ := strconv.ParseInt(c.Param("id"), 10, 64)
			type_, _ := strconv.Atoi(c.Query("type"))
			var controller api.DownloadController
			controller.Init(c)
			controller.Contract(id_, type_)
			controller.Close()
		})

		apiGroup.GET("/download/addressrepair", func(c *gin.Context) {
			var controller api.DownloadController
			controller.Init(c)
			controller.Addressrepair()
			controller.Close()
		})

		apiGroup.GET("/download/detail0/:id", func(c *gin.Context) {
			id_, _ := strconv.ParseInt(c.Param("id"), 10, 64)
			var controller api.DownloadController
			controller.Init(c)
			controller.Detail0(id_)
			controller.Close()
		})

		apiGroup.GET("/download/detail1/:id", func(c *gin.Context) {
			id_, _ := strconv.ParseInt(c.Param("id"), 10, 64)
			var controller api.DownloadController
			controller.Init(c)
			controller.Detail1(id_)
			controller.Close()
		})

		apiGroup.GET("/download/detail2/:id", func(c *gin.Context) {
			id_, _ := strconv.ParseInt(c.Param("id"), 10, 64)
			var controller api.DownloadController
			controller.Init(c)
			controller.Detail2(id_)
			controller.Close()
		})

		apiGroup.GET("/download/detail3/:id", func(c *gin.Context) {
			id_, _ := strconv.ParseInt(c.Param("id"), 10, 64)
			var controller api.DownloadController
			controller.Init(c)
			controller.Detail3(id_)
			controller.Close()
		})

		apiGroup.GET("/download/detail4/:id", func(c *gin.Context) {
			id_, _ := strconv.ParseInt(c.Param("id"), 10, 64)
			var controller api.DownloadController
			controller.Init(c)
			controller.Detail4(id_)
			controller.Close()
		})

		apiGroup.GET("/download/detail5/:id", func(c *gin.Context) {
			id_, _ := strconv.ParseInt(c.Param("id"), 10, 64)
			var controller api.DownloadController
			controller.Init(c)
			controller.Detail5(id_)
			controller.Close()
		})

		apiGroup.GET("/download/detail6/:id", func(c *gin.Context) {
			id_, _ := strconv.ParseInt(c.Param("id"), 10, 64)
			var controller api.DownloadController
			controller.Init(c)
			controller.Detail6(id_)
			controller.Close()
		})

		apiGroup.GET("/download/detail7/:id", func(c *gin.Context) {
			id_, _ := strconv.ParseInt(c.Param("id"), 10, 64)
			var controller api.DownloadController
			controller.Init(c)
			controller.Detail7(id_)
			controller.Close()
		})

		apiGroup.GET("/download/detail/:id", func(c *gin.Context) {
			id_, _ := strconv.ParseInt(c.Param("id"), 10, 64)
			var controller api.DownloadController
			controller.Init(c)
			controller.Detail(id_)
			controller.Close()
		})

		apiGroup.POST("/estimate", func(c *gin.Context) {
			item_ := &models.EstimateExtra{}
			c.ShouldBindJSON(item_)
			var controller api.EstimateController
			controller.Init(c)
			controller.Insert(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.PUT("/estimate", func(c *gin.Context) {
			item_ := &models.EstimateExtra{}
			c.ShouldBindJSON(item_)
			var controller api.EstimateController
			controller.Init(c)
			controller.Update(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.DELETE("/estimate", func(c *gin.Context) {
			item_ := &models.Estimate{}
			c.ShouldBindJSON(item_)
			var controller api.EstimateController
			controller.Init(c)
			controller.Delete(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/estimate", func(c *gin.Context) {
			page_, _ := strconv.Atoi(c.Query("page"))
			pagesize_, _ := strconv.Atoi(c.Query("pagesize"))
			var controller api.EstimateController
			controller.Init(c)
			controller.Index(page_, pagesize_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/image/size/:filename", func(c *gin.Context) {
			filename_ := c.Param("filename")
			var controller api.ImageController
			controller.Init(c)
			controller.Size(filename_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.POST("/managebook/process", func(c *gin.Context) {
			var results map[string]any
			jsonData, _ := c.GetRawData()
			json.Unmarshal(jsonData, &results)
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var name_ string
			if v, flag := results["name"]; flag {
				name_ = v.(string)
			}
			var order_ int
			if v, flag := results["order"]; flag {
				order_ = int(v.(float64))
			}
			var filename_ string
			if v, flag := results["filename"]; flag {
				filename_ = v.(string)
			}
			var controller api.ManagebookController
			controller.Init(c)
			controller.Process(id_, name_, order_, filename_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.POST("/managebook/multiprocess", func(c *gin.Context) {
			var results map[string]any
			jsonData, _ := c.GetRawData()
			json.Unmarshal(jsonData, &results)
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var filename_ string
			if v, flag := results["filename"]; flag {
				filename_ = v.(string)
			}
			var originalfilename_ string
			if v, flag := results["originalfilename"]; flag {
				originalfilename_ = v.(string)
			}
			var controller api.ManagebookController
			controller.Init(c)
			controller.Multiprocess(id_, filename_, originalfilename_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.POST("/oldapt/convert", func(c *gin.Context) {
			items_ := &models.ConvertOldapt{}
			c.ShouldBindJSON(items_)
			var controller api.OldaptController
			controller.Init(c)
			controller.Convert(items_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.POST("/patrol/upload", func(c *gin.Context) {
			var controller api.PatrolController
			controller.Init(c)
			controller.Upload()
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/periodic/data/:id", func(c *gin.Context) {
			id_, _ := strconv.ParseInt(c.Param("id"), 10, 64)
			var controller api.PeriodicController
			controller.Init(c)
			controller.Data(id_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.POST("/periodic/upload", func(c *gin.Context) {
			item_ := &models.Data{}
			c.ShouldBindJSON(item_)
			var controller api.PeriodicController
			controller.Init(c)
			controller.Upload(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.POST("/periodic/duplication", func(c *gin.Context) {
			var results map[string]any
			jsonData, _ := c.GetRawData()
			json.Unmarshal(jsonData, &results)
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller api.PeriodicController
			controller.Init(c)
			controller.Duplication(id_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/periodic/search", func(c *gin.Context) {
			var controller api.PeriodicController
			controller.Init(c)
			controller.Search()
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.POST("/periodicimage/process", func(c *gin.Context) {
			item_ := &models.Periodicimage{}
			c.ShouldBindJSON(item_)
			var controller api.PeriodicimageController
			controller.Init(c)
			controller.Process(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/repair/lastdate/:id", func(c *gin.Context) {
			id_, _ := strconv.ParseInt(c.Param("id"), 10, 64)
			var controller api.RepairController
			controller.Init(c)
			controller.Lastdate(id_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.POST("/repair/change", func(c *gin.Context) {
			var results map[string]any
			jsonData, _ := c.GetRawData()
			json.Unmarshal(jsonData, &results)
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller api.RepairController
			controller.Init(c)
			controller.Change(id_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/repairlist/search", func(c *gin.Context) {
			var controller api.RepairlistController
			controller.Init(c)
			controller.Search()
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/report/total/:id", func(c *gin.Context) {
			id_, _ := strconv.ParseInt(c.Param("id"), 10, 64)
			var controller api.ReportController
			controller.Init(c)
			controller.Total(id_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/report/summary/:id", func(c *gin.Context) {
			id_, _ := strconv.ParseInt(c.Param("id"), 10, 64)
			var controller api.ReportController
			controller.Init(c)
			controller.Summary(id_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/report/plan/:id", func(c *gin.Context) {
			id_, _ := strconv.ParseInt(c.Param("id"), 10, 64)
			var controller api.ReportController
			controller.Init(c)
			controller.Plan(id_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.POST("/standard/all", func(c *gin.Context) {
			item_ := &models.Standard{}
			c.ShouldBindJSON(item_)
			var controller api.StandardController
			controller.Init(c)
			controller.Insertall(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.PUT("/standard/all", func(c *gin.Context) {
			item_ := &models.Standard{}
			c.ShouldBindJSON(item_)
			var controller api.StandardController
			controller.Init(c)
			controller.Updateall(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.POST("/upload/index", func(c *gin.Context) {
			var controller api.UploadController
			controller.Init(c)
			controller.Index()
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.POST("/upload/periodic", func(c *gin.Context) {
			var controller api.UploadController
			controller.Init(c)
			controller.Periodic()
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/upload/diff/:id", func(c *gin.Context) {
			id_, _ := strconv.ParseInt(c.Param("id"), 10, 64)
			filename_ := c.Query("filename")
			var controller api.UploadController
			controller.Init(c)
			controller.Diff(id_, filename_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.POST("/upload/diffupdate", func(c *gin.Context) {
			diffs_ := &models.Diff{}
			c.ShouldBindJSON(diffs_)
			var controller api.UploadController
			controller.Init(c)
			controller.Diffupdate(diffs_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/upload/excel/:id", func(c *gin.Context) {
			id_, _ := strconv.ParseInt(c.Param("id"), 10, 64)
			filename_ := c.Query("filename")
			historydel_, _ := strconv.Atoi(c.Query("historydel"))
			breakdowndel_, _ := strconv.Atoi(c.Query("breakdowndel"))
			var controller api.UploadController
			controller.Init(c)
			controller.Excel(id_, filename_, historydel_, breakdowndel_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.POST("/upload/assistance", func(c *gin.Context) {
			var results map[string]any
			jsonData, _ := c.GetRawData()
			json.Unmarshal(jsonData, &results)
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var filenames_ []string
			if v, flag := results["filenames"]; flag {
			    strs := make([]string, 0)
			    for _, str := range v.([]any) {
			        strs = append(strs, str.(string))
			    }
				filenames_ = strs
			}
			var historydel_ int
			if v, flag := results["historydel"]; flag {
				historydel_ = int(v.(float64))
			}
			var controller api.UploadController
			controller.Init(c)
			controller.Assistance(id_, filenames_, historydel_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

	}

	{

		apiGroup.GET("/adjust/count/apt/:apt", func(c *gin.Context) {
			apt_, _ := strconv.ParseInt(c.Param("apt"), 10, 64)
			var controller rest.AdjustController
			controller.Init(c)
			controller.CountByApt(apt_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/adjust/find/apt/:apt", func(c *gin.Context) {
			apt_, _ := strconv.ParseInt(c.Param("apt"), 10, 64)
			var controller rest.AdjustController
			controller.Init(c)
			controller.FindByApt(apt_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.DELETE("/adjust/byapt", func(c *gin.Context) {
			item_ := &models.Adjust{}
			c.ShouldBindJSON(item_)
			var controller rest.AdjustController
			controller.Init(c)
			controller.DeleteByApt(item_.Apt)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/adjust/find/category/:category", func(c *gin.Context) {
			category_, _ := strconv.ParseInt(c.Param("category"), 10, 64)
			var controller rest.AdjustController
			controller.Init(c)
			controller.FindByCategory(category_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.DELETE("/adjust/bycategory", func(c *gin.Context) {
			item_ := &models.Adjust{}
			c.ShouldBindJSON(item_)
			var controller rest.AdjustController
			controller.Init(c)
			controller.DeleteByCategory(item_.Category)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.POST("/adjust", func(c *gin.Context) {
			item_ := &models.Adjust{}
			c.ShouldBindJSON(item_)
			var apicontroller api.AdjustController
			apicontroller.Init(c)
			var controller rest.AdjustController
			controller.Init(c)
			controller.Insert(item_)
			apicontroller.Post_Insert(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.POST("/adjust/batch", func(c *gin.Context) {
			item_ := &[]models.Adjust{}
			c.ShouldBindJSON(item_)
			var controller rest.AdjustController
			controller.Init(c)
			controller.Insertbatch(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.PUT("/adjust", func(c *gin.Context) {
			item_ := &models.Adjust{}
			c.ShouldBindJSON(item_)
			var apicontroller api.AdjustController
			apicontroller.Init(c)
			var controller rest.AdjustController
			controller.Init(c)
			controller.Update(item_)
			apicontroller.Post_Update(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.DELETE("/adjust", func(c *gin.Context) {
			item_ := &models.Adjust{}
			c.ShouldBindJSON(item_)
			var apicontroller api.AdjustController
			apicontroller.Init(c)
			var controller rest.AdjustController
			controller.Init(c)
			controller.Delete(item_)
			apicontroller.Post_Delete(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.DELETE("/adjust/batch", func(c *gin.Context) {
			item_ := &[]models.Adjust{}
			c.ShouldBindJSON(item_)
			var apicontroller api.AdjustController
			apicontroller.Init(c)
			var controller rest.AdjustController
			controller.Init(c)
			controller.Deletebatch(item_)
			apicontroller.Post_Deletebatch(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/adjust/count", func(c *gin.Context) {
			var controller rest.AdjustController
			controller.Init(c)
			controller.Count()
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/adjust/:id", func(c *gin.Context) {
			id_, _ := strconv.ParseInt(c.Param("id"), 10, 64)
			var controller rest.AdjustController
			controller.Init(c)
			controller.Read(id_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/adjust", func(c *gin.Context) {
			page_, _ := strconv.Atoi(c.Query("page"))
			pagesize_, _ := strconv.Atoi(c.Query("pagesize"))
			var controller rest.AdjustController
			controller.Init(c)
			controller.Index(page_, pagesize_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/advice/count/apt/:apt", func(c *gin.Context) {
			apt_, _ := strconv.ParseInt(c.Param("apt"), 10, 64)
			var controller rest.AdviceController
			controller.Init(c)
			controller.CountByApt(apt_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/advice/find/apt/:apt", func(c *gin.Context) {
			apt_, _ := strconv.ParseInt(c.Param("apt"), 10, 64)
			var controller rest.AdviceController
			controller.Init(c)
			controller.FindByApt(apt_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.POST("/advice", func(c *gin.Context) {
			item_ := &models.Advice{}
			c.ShouldBindJSON(item_)
			var controller rest.AdviceController
			controller.Init(c)
			controller.Insert(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.POST("/advice/batch", func(c *gin.Context) {
			item_ := &[]models.Advice{}
			c.ShouldBindJSON(item_)
			var controller rest.AdviceController
			controller.Init(c)
			controller.Insertbatch(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.PUT("/advice", func(c *gin.Context) {
			item_ := &models.Advice{}
			c.ShouldBindJSON(item_)
			var controller rest.AdviceController
			controller.Init(c)
			controller.Update(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.DELETE("/advice", func(c *gin.Context) {
			item_ := &models.Advice{}
			c.ShouldBindJSON(item_)
			var controller rest.AdviceController
			controller.Init(c)
			controller.Delete(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.DELETE("/advice/batch", func(c *gin.Context) {
			item_ := &[]models.Advice{}
			c.ShouldBindJSON(item_)
			var controller rest.AdviceController
			controller.Init(c)
			controller.Deletebatch(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/advice/count", func(c *gin.Context) {
			var controller rest.AdviceController
			controller.Init(c)
			controller.Count()
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/advice/:id", func(c *gin.Context) {
			id_, _ := strconv.ParseInt(c.Param("id"), 10, 64)
			var controller rest.AdviceController
			controller.Init(c)
			controller.Read(id_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/advice", func(c *gin.Context) {
			page_, _ := strconv.Atoi(c.Query("page"))
			pagesize_, _ := strconv.Atoi(c.Query("pagesize"))
			var controller rest.AdviceController
			controller.Init(c)
			controller.Index(page_, pagesize_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/approval/count/apt/:apt", func(c *gin.Context) {
			apt_, _ := strconv.ParseInt(c.Param("apt"), 10, 64)
			var controller rest.ApprovalController
			controller.Init(c)
			controller.CountByApt(apt_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/approval/find/apt/:apt", func(c *gin.Context) {
			apt_, _ := strconv.ParseInt(c.Param("apt"), 10, 64)
			var controller rest.ApprovalController
			controller.Init(c)
			controller.FindByApt(apt_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.DELETE("/approval/byapt", func(c *gin.Context) {
			item_ := &models.Approval{}
			c.ShouldBindJSON(item_)
			var controller rest.ApprovalController
			controller.Init(c)
			controller.DeleteByApt(item_.Apt)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.POST("/approval", func(c *gin.Context) {
			item_ := &models.Approval{}
			c.ShouldBindJSON(item_)
			var controller rest.ApprovalController
			controller.Init(c)
			controller.Insert(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.POST("/approval/batch", func(c *gin.Context) {
			item_ := &[]models.Approval{}
			c.ShouldBindJSON(item_)
			var controller rest.ApprovalController
			controller.Init(c)
			controller.Insertbatch(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.PUT("/approval", func(c *gin.Context) {
			item_ := &models.Approval{}
			c.ShouldBindJSON(item_)
			var controller rest.ApprovalController
			controller.Init(c)
			controller.Update(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.DELETE("/approval", func(c *gin.Context) {
			item_ := &models.Approval{}
			c.ShouldBindJSON(item_)
			var controller rest.ApprovalController
			controller.Init(c)
			controller.Delete(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.DELETE("/approval/batch", func(c *gin.Context) {
			item_ := &[]models.Approval{}
			c.ShouldBindJSON(item_)
			var controller rest.ApprovalController
			controller.Init(c)
			controller.Deletebatch(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/approval/count", func(c *gin.Context) {
			var controller rest.ApprovalController
			controller.Init(c)
			controller.Count()
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/approval/:id", func(c *gin.Context) {
			id_, _ := strconv.ParseInt(c.Param("id"), 10, 64)
			var controller rest.ApprovalController
			controller.Init(c)
			controller.Read(id_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/approval", func(c *gin.Context) {
			page_, _ := strconv.Atoi(c.Query("page"))
			pagesize_, _ := strconv.Atoi(c.Query("pagesize"))
			var controller rest.ApprovalController
			controller.Init(c)
			controller.Index(page_, pagesize_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/apt/count/namelike/:name", func(c *gin.Context) {
			name_ := c.Param("name")
			var controller rest.AptController
			controller.Init(c)
			controller.CountByNamelike(name_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/apt/find/namelike/:name", func(c *gin.Context) {
			name_ := c.Param("name")
			var controller rest.AptController
			controller.Init(c)
			controller.FindByNamelike(name_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/apt/count/emaillike/:email", func(c *gin.Context) {
			email_ := c.Param("email")
			var controller rest.AptController
			controller.Init(c)
			controller.CountByEmaillike(email_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/apt/find/emaillike/:email", func(c *gin.Context) {
			email_ := c.Param("email")
			var controller rest.AptController
			controller.Init(c)
			controller.FindByEmaillike(email_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.POST("/apt", func(c *gin.Context) {
			item_ := &models.Apt{}
			c.ShouldBindJSON(item_)
			var controller rest.AptController
			controller.Init(c)
			controller.Insert(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.POST("/apt/batch", func(c *gin.Context) {
			item_ := &[]models.Apt{}
			c.ShouldBindJSON(item_)
			var controller rest.AptController
			controller.Init(c)
			controller.Insertbatch(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.PUT("/apt", func(c *gin.Context) {
			item_ := &models.Apt{}
			c.ShouldBindJSON(item_)
			var controller rest.AptController
			controller.Init(c)
			controller.Update(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.DELETE("/apt", func(c *gin.Context) {
			item_ := &models.Apt{}
			c.ShouldBindJSON(item_)
			var controller rest.AptController
			controller.Init(c)
			controller.Delete(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.DELETE("/apt/batch", func(c *gin.Context) {
			item_ := &[]models.Apt{}
			c.ShouldBindJSON(item_)
			var controller rest.AptController
			controller.Init(c)
			controller.Deletebatch(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/apt/count", func(c *gin.Context) {
			var controller rest.AptController
			controller.Init(c)
			controller.Count()
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/apt/:id", func(c *gin.Context) {
			id_, _ := strconv.ParseInt(c.Param("id"), 10, 64)
			var controller rest.AptController
			controller.Init(c)
			controller.Read(id_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/apt", func(c *gin.Context) {
			page_, _ := strconv.Atoi(c.Query("page"))
			pagesize_, _ := strconv.Atoi(c.Query("pagesize"))
			var controller rest.AptController
			controller.Init(c)
			controller.Index(page_, pagesize_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.POST("/aptdetail", func(c *gin.Context) {
			item_ := &models.Aptdetail{}
			c.ShouldBindJSON(item_)
			var controller rest.AptdetailController
			controller.Init(c)
			controller.Insert(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.POST("/aptdetail/batch", func(c *gin.Context) {
			item_ := &[]models.Aptdetail{}
			c.ShouldBindJSON(item_)
			var controller rest.AptdetailController
			controller.Init(c)
			controller.Insertbatch(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.PUT("/aptdetail", func(c *gin.Context) {
			item_ := &models.Aptdetail{}
			c.ShouldBindJSON(item_)
			var controller rest.AptdetailController
			controller.Init(c)
			controller.Update(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.DELETE("/aptdetail", func(c *gin.Context) {
			item_ := &models.Aptdetail{}
			c.ShouldBindJSON(item_)
			var controller rest.AptdetailController
			controller.Init(c)
			controller.Delete(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.DELETE("/aptdetail/batch", func(c *gin.Context) {
			item_ := &[]models.Aptdetail{}
			c.ShouldBindJSON(item_)
			var controller rest.AptdetailController
			controller.Init(c)
			controller.Deletebatch(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/aptdetail/count", func(c *gin.Context) {
			var controller rest.AptdetailController
			controller.Init(c)
			controller.Count()
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/aptdetail/:id", func(c *gin.Context) {
			id_, _ := strconv.ParseInt(c.Param("id"), 10, 64)
			var controller rest.AptdetailController
			controller.Init(c)
			controller.Read(id_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/aptdetail", func(c *gin.Context) {
			page_, _ := strconv.Atoi(c.Query("page"))
			pagesize_, _ := strconv.Atoi(c.Query("pagesize"))
			var controller rest.AptdetailController
			controller.Init(c)
			controller.Index(page_, pagesize_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.POST("/aptdong", func(c *gin.Context) {
			item_ := &models.Aptdong{}
			c.ShouldBindJSON(item_)
			var apicontroller api.AptdongController
			apicontroller.Init(c)
			var controller rest.AptdongController
			controller.Init(c)
			controller.Insert(item_)
			apicontroller.Post_Insert(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.POST("/aptdong/batch", func(c *gin.Context) {
			item_ := &[]models.Aptdong{}
			c.ShouldBindJSON(item_)
			var controller rest.AptdongController
			controller.Init(c)
			controller.Insertbatch(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.PUT("/aptdong", func(c *gin.Context) {
			item_ := &models.Aptdong{}
			c.ShouldBindJSON(item_)
			var apicontroller api.AptdongController
			apicontroller.Init(c)
			var controller rest.AptdongController
			controller.Init(c)
			apicontroller.Pre_Update(item_)
			controller.Update(item_)
			apicontroller.Post_Update(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.DELETE("/aptdong", func(c *gin.Context) {
			item_ := &models.Aptdong{}
			c.ShouldBindJSON(item_)
			var apicontroller api.AptdongController
			apicontroller.Init(c)
			var controller rest.AptdongController
			controller.Init(c)
			controller.Delete(item_)
			apicontroller.Post_Delete(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.DELETE("/aptdong/batch", func(c *gin.Context) {
			item_ := &[]models.Aptdong{}
			c.ShouldBindJSON(item_)
			var apicontroller api.AptdongController
			apicontroller.Init(c)
			var controller rest.AptdongController
			controller.Init(c)
			controller.Deletebatch(item_)
			apicontroller.Post_Deletebatch(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/aptdong/count", func(c *gin.Context) {
			var controller rest.AptdongController
			controller.Init(c)
			controller.Count()
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/aptdong/:id", func(c *gin.Context) {
			id_, _ := strconv.ParseInt(c.Param("id"), 10, 64)
			var controller rest.AptdongController
			controller.Init(c)
			controller.Read(id_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/aptdong", func(c *gin.Context) {
			page_, _ := strconv.Atoi(c.Query("page"))
			pagesize_, _ := strconv.Atoi(c.Query("pagesize"))
			var controller rest.AptdongController
			controller.Init(c)
			controller.Index(page_, pagesize_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/aptdongetc/get/aptaptdongparentname/:apt", func(c *gin.Context) {
			apt_, _ := strconv.ParseInt(c.Param("apt"), 10, 64)
			aptdong_, _ := strconv.ParseInt(c.Query("aptdong"), 10, 64)
			parent_, _ := strconv.Atoi(c.Query("parent"))
			name_ := c.Query("name")
			var controller rest.AptdongetcController
			controller.Init(c)
			controller.GetByAptAptdongParentName(apt_, aptdong_, parent_, name_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/aptdongetc/count/aptdong/:aptdong", func(c *gin.Context) {
			aptdong_, _ := strconv.ParseInt(c.Param("aptdong"), 10, 64)
			var controller rest.AptdongetcController
			controller.Init(c)
			controller.CountByAptdong(aptdong_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.DELETE("/aptdongetc/byaptdong", func(c *gin.Context) {
			item_ := &models.Aptdongetc{}
			c.ShouldBindJSON(item_)
			var controller rest.AptdongetcController
			controller.Init(c)
			controller.DeleteByAptdong(item_.Aptdong)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.POST("/aptdongetc", func(c *gin.Context) {
			item_ := &models.Aptdongetc{}
			c.ShouldBindJSON(item_)
			var controller rest.AptdongetcController
			controller.Init(c)
			controller.Insert(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.POST("/aptdongetc/batch", func(c *gin.Context) {
			item_ := &[]models.Aptdongetc{}
			c.ShouldBindJSON(item_)
			var controller rest.AptdongetcController
			controller.Init(c)
			controller.Insertbatch(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.PUT("/aptdongetc", func(c *gin.Context) {
			item_ := &models.Aptdongetc{}
			c.ShouldBindJSON(item_)
			var controller rest.AptdongetcController
			controller.Init(c)
			controller.Update(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.DELETE("/aptdongetc", func(c *gin.Context) {
			item_ := &models.Aptdongetc{}
			c.ShouldBindJSON(item_)
			var controller rest.AptdongetcController
			controller.Init(c)
			controller.Delete(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.DELETE("/aptdongetc/batch", func(c *gin.Context) {
			item_ := &[]models.Aptdongetc{}
			c.ShouldBindJSON(item_)
			var controller rest.AptdongetcController
			controller.Init(c)
			controller.Deletebatch(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/aptdongetc/count", func(c *gin.Context) {
			var controller rest.AptdongetcController
			controller.Init(c)
			controller.Count()
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/aptdongetc/:id", func(c *gin.Context) {
			id_, _ := strconv.ParseInt(c.Param("id"), 10, 64)
			var controller rest.AptdongetcController
			controller.Init(c)
			controller.Read(id_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/aptdongetc", func(c *gin.Context) {
			page_, _ := strconv.Atoi(c.Query("page"))
			pagesize_, _ := strconv.Atoi(c.Query("pagesize"))
			var controller rest.AptdongetcController
			controller.Init(c)
			controller.Index(page_, pagesize_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/aptlist/count", func(c *gin.Context) {
			var controller rest.AptlistController
			controller.Init(c)
			controller.Count()
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/aptlist/:id", func(c *gin.Context) {
			id_, _ := strconv.ParseInt(c.Param("id"), 10, 64)
			var controller rest.AptlistController
			controller.Init(c)
			controller.Read(id_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/aptlist", func(c *gin.Context) {
			page_, _ := strconv.Atoi(c.Query("page"))
			pagesize_, _ := strconv.Atoi(c.Query("pagesize"))
			var controller rest.AptlistController
			controller.Init(c)
			controller.Index(page_, pagesize_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.POST("/aptperiodic", func(c *gin.Context) {
			item_ := &models.Aptperiodic{}
			c.ShouldBindJSON(item_)
			var controller rest.AptperiodicController
			controller.Init(c)
			controller.Insert(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.POST("/aptperiodic/batch", func(c *gin.Context) {
			item_ := &[]models.Aptperiodic{}
			c.ShouldBindJSON(item_)
			var controller rest.AptperiodicController
			controller.Init(c)
			controller.Insertbatch(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.PUT("/aptperiodic", func(c *gin.Context) {
			item_ := &models.Aptperiodic{}
			c.ShouldBindJSON(item_)
			var controller rest.AptperiodicController
			controller.Init(c)
			controller.Update(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.DELETE("/aptperiodic", func(c *gin.Context) {
			item_ := &models.Aptperiodic{}
			c.ShouldBindJSON(item_)
			var controller rest.AptperiodicController
			controller.Init(c)
			controller.Delete(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.DELETE("/aptperiodic/batch", func(c *gin.Context) {
			item_ := &[]models.Aptperiodic{}
			c.ShouldBindJSON(item_)
			var controller rest.AptperiodicController
			controller.Init(c)
			controller.Deletebatch(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/aptperiodic/count", func(c *gin.Context) {
			var controller rest.AptperiodicController
			controller.Init(c)
			controller.Count()
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/aptperiodic/:id", func(c *gin.Context) {
			id_, _ := strconv.ParseInt(c.Param("id"), 10, 64)
			var controller rest.AptperiodicController
			controller.Init(c)
			controller.Read(id_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/aptperiodic", func(c *gin.Context) {
			page_, _ := strconv.Atoi(c.Query("page"))
			pagesize_, _ := strconv.Atoi(c.Query("pagesize"))
			var controller rest.AptperiodicController
			controller.Init(c)
			controller.Index(page_, pagesize_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/aptrepairlist/count", func(c *gin.Context) {
			var controller rest.AptrepairlistController
			controller.Init(c)
			controller.Count()
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/aptrepairlist/:id", func(c *gin.Context) {
			id_, _ := strconv.ParseInt(c.Param("id"), 10, 64)
			var controller rest.AptrepairlistController
			controller.Init(c)
			controller.Read(id_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/aptrepairlist", func(c *gin.Context) {
			page_, _ := strconv.Atoi(c.Query("page"))
			pagesize_, _ := strconv.Atoi(c.Query("pagesize"))
			var controller rest.AptrepairlistController
			controller.Init(c)
			controller.Index(page_, pagesize_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/aptusagefloor/count/apt/:apt", func(c *gin.Context) {
			apt_, _ := strconv.ParseInt(c.Param("apt"), 10, 64)
			var controller rest.AptusagefloorController
			controller.Init(c)
			controller.CountByApt(apt_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/aptusagefloor/find/apt/:apt", func(c *gin.Context) {
			apt_, _ := strconv.ParseInt(c.Param("apt"), 10, 64)
			var controller rest.AptusagefloorController
			controller.Init(c)
			controller.FindByApt(apt_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.DELETE("/aptusagefloor/byapt", func(c *gin.Context) {
			item_ := &models.Aptusagefloor{}
			c.ShouldBindJSON(item_)
			var controller rest.AptusagefloorController
			controller.Init(c)
			controller.DeleteByApt(item_.Apt)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.POST("/aptusagefloor", func(c *gin.Context) {
			item_ := &models.Aptusagefloor{}
			c.ShouldBindJSON(item_)
			var controller rest.AptusagefloorController
			controller.Init(c)
			controller.Insert(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.POST("/aptusagefloor/batch", func(c *gin.Context) {
			item_ := &[]models.Aptusagefloor{}
			c.ShouldBindJSON(item_)
			var controller rest.AptusagefloorController
			controller.Init(c)
			controller.Insertbatch(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.PUT("/aptusagefloor", func(c *gin.Context) {
			item_ := &models.Aptusagefloor{}
			c.ShouldBindJSON(item_)
			var controller rest.AptusagefloorController
			controller.Init(c)
			controller.Update(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.DELETE("/aptusagefloor", func(c *gin.Context) {
			item_ := &models.Aptusagefloor{}
			c.ShouldBindJSON(item_)
			var controller rest.AptusagefloorController
			controller.Init(c)
			controller.Delete(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.DELETE("/aptusagefloor/batch", func(c *gin.Context) {
			item_ := &[]models.Aptusagefloor{}
			c.ShouldBindJSON(item_)
			var controller rest.AptusagefloorController
			controller.Init(c)
			controller.Deletebatch(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/aptusagefloor/count", func(c *gin.Context) {
			var controller rest.AptusagefloorController
			controller.Init(c)
			controller.Count()
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/aptusagefloor/:id", func(c *gin.Context) {
			id_, _ := strconv.ParseInt(c.Param("id"), 10, 64)
			var controller rest.AptusagefloorController
			controller.Init(c)
			controller.Read(id_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/aptusagefloor", func(c *gin.Context) {
			page_, _ := strconv.Atoi(c.Query("page"))
			pagesize_, _ := strconv.Atoi(c.Query("pagesize"))
			var controller rest.AptusagefloorController
			controller.Init(c)
			controller.Index(page_, pagesize_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/area/count/apt/:apt", func(c *gin.Context) {
			apt_, _ := strconv.ParseInt(c.Param("apt"), 10, 64)
			var controller rest.AreaController
			controller.Init(c)
			controller.CountByApt(apt_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/area/find/apt/:apt", func(c *gin.Context) {
			apt_, _ := strconv.ParseInt(c.Param("apt"), 10, 64)
			var controller rest.AreaController
			controller.Init(c)
			controller.FindByApt(apt_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.DELETE("/area/byapt", func(c *gin.Context) {
			item_ := &models.Area{}
			c.ShouldBindJSON(item_)
			var controller rest.AreaController
			controller.Init(c)
			controller.DeleteByApt(item_.Apt)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.POST("/area", func(c *gin.Context) {
			item_ := &models.Area{}
			c.ShouldBindJSON(item_)
			var controller rest.AreaController
			controller.Init(c)
			controller.Insert(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.POST("/area/batch", func(c *gin.Context) {
			item_ := &[]models.Area{}
			c.ShouldBindJSON(item_)
			var controller rest.AreaController
			controller.Init(c)
			controller.Insertbatch(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.PUT("/area", func(c *gin.Context) {
			item_ := &models.Area{}
			c.ShouldBindJSON(item_)
			var controller rest.AreaController
			controller.Init(c)
			controller.Update(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.DELETE("/area", func(c *gin.Context) {
			item_ := &models.Area{}
			c.ShouldBindJSON(item_)
			var controller rest.AreaController
			controller.Init(c)
			controller.Delete(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.DELETE("/area/batch", func(c *gin.Context) {
			item_ := &[]models.Area{}
			c.ShouldBindJSON(item_)
			var controller rest.AreaController
			controller.Init(c)
			controller.Deletebatch(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/area/count", func(c *gin.Context) {
			var controller rest.AreaController
			controller.Init(c)
			controller.Count()
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/area/:id", func(c *gin.Context) {
			id_, _ := strconv.ParseInt(c.Param("id"), 10, 64)
			var controller rest.AreaController
			controller.Init(c)
			controller.Read(id_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/area", func(c *gin.Context) {
			page_, _ := strconv.Atoi(c.Query("page"))
			pagesize_, _ := strconv.Atoi(c.Query("pagesize"))
			var controller rest.AreaController
			controller.Init(c)
			controller.Index(page_, pagesize_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/blueprint/count/apt/:apt", func(c *gin.Context) {
			apt_, _ := strconv.ParseInt(c.Param("apt"), 10, 64)
			var controller rest.BlueprintController
			controller.Init(c)
			controller.CountByApt(apt_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/blueprint/find/apt/:apt", func(c *gin.Context) {
			apt_, _ := strconv.ParseInt(c.Param("apt"), 10, 64)
			var controller rest.BlueprintController
			controller.Init(c)
			controller.FindByApt(apt_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.PUT("/blueprint/filenamebyid", func(c *gin.Context) {
			var results map[string]any
			jsonData, _ := c.GetRawData()
			json.Unmarshal(jsonData, &results)
			var filename_ string
			if v, flag := results["filename"]; flag {
				filename_ = v.(string)
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.BlueprintController
			controller.Init(c)
			controller.UpdateFilenameById(filename_, id_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.POST("/blueprint", func(c *gin.Context) {
			item_ := &models.Blueprint{}
			c.ShouldBindJSON(item_)
			var controller rest.BlueprintController
			controller.Init(c)
			controller.Insert(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.POST("/blueprint/batch", func(c *gin.Context) {
			item_ := &[]models.Blueprint{}
			c.ShouldBindJSON(item_)
			var controller rest.BlueprintController
			controller.Init(c)
			controller.Insertbatch(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.PUT("/blueprint", func(c *gin.Context) {
			item_ := &models.Blueprint{}
			c.ShouldBindJSON(item_)
			var controller rest.BlueprintController
			controller.Init(c)
			controller.Update(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.DELETE("/blueprint", func(c *gin.Context) {
			item_ := &models.Blueprint{}
			c.ShouldBindJSON(item_)
			var controller rest.BlueprintController
			controller.Init(c)
			controller.Delete(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.DELETE("/blueprint/batch", func(c *gin.Context) {
			item_ := &[]models.Blueprint{}
			c.ShouldBindJSON(item_)
			var controller rest.BlueprintController
			controller.Init(c)
			controller.Deletebatch(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/blueprint/count", func(c *gin.Context) {
			var controller rest.BlueprintController
			controller.Init(c)
			controller.Count()
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/blueprint/:id", func(c *gin.Context) {
			id_, _ := strconv.ParseInt(c.Param("id"), 10, 64)
			var controller rest.BlueprintController
			controller.Init(c)
			controller.Read(id_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/blueprint", func(c *gin.Context) {
			page_, _ := strconv.Atoi(c.Query("page"))
			pagesize_, _ := strconv.Atoi(c.Query("pagesize"))
			var controller rest.BlueprintController
			controller.Init(c)
			controller.Index(page_, pagesize_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/breakdown/count/apt/:apt", func(c *gin.Context) {
			apt_, _ := strconv.ParseInt(c.Param("apt"), 10, 64)
			var controller rest.BreakdownController
			controller.Init(c)
			controller.CountByApt(apt_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/breakdown/find/apt/:apt", func(c *gin.Context) {
			apt_, _ := strconv.ParseInt(c.Param("apt"), 10, 64)
			var controller rest.BreakdownController
			controller.Init(c)
			controller.FindByApt(apt_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.DELETE("/breakdown/byapt", func(c *gin.Context) {
			item_ := &models.Breakdown{}
			c.ShouldBindJSON(item_)
			var controller rest.BreakdownController
			controller.Init(c)
			controller.DeleteByApt(item_.Apt)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/breakdown/count/aptdong/:apt", func(c *gin.Context) {
			apt_, _ := strconv.ParseInt(c.Param("apt"), 10, 64)
			dong_, _ := strconv.ParseInt(c.Query("dong"), 10, 64)
			var controller rest.BreakdownController
			controller.Init(c)
			controller.CountByAptDong(apt_, dong_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/breakdown/count/aptstandard/:apt", func(c *gin.Context) {
			apt_, _ := strconv.ParseInt(c.Param("apt"), 10, 64)
			standard_, _ := strconv.ParseInt(c.Query("standard"), 10, 64)
			var controller rest.BreakdownController
			controller.Init(c)
			controller.CountByAptStandard(apt_, standard_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.PUT("/breakdown/duedatebyid", func(c *gin.Context) {
			var results map[string]any
			jsonData, _ := c.GetRawData()
			json.Unmarshal(jsonData, &results)
			var duedate_ int
			if v, flag := results["duedate"]; flag {
				duedate_ = int(v.(float64))
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var apicontroller api.BreakdownController
			apicontroller.Init(c)
			var controller rest.BreakdownController
			controller.Init(c)
			apicontroller.Pre_UpdateDuedateById(duedate_, id_)
			controller.UpdateDuedateById(duedate_, id_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.PUT("/breakdown/lastdatebyid", func(c *gin.Context) {
			var results map[string]any
			jsonData, _ := c.GetRawData()
			json.Unmarshal(jsonData, &results)
			var lastdate_ int
			if v, flag := results["lastdate"]; flag {
				lastdate_ = int(v.(float64))
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.BreakdownController
			controller.Init(c)
			controller.UpdateLastdateById(lastdate_, id_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/breakdown/find/category/:category", func(c *gin.Context) {
			category_, _ := strconv.ParseInt(c.Param("category"), 10, 64)
			var controller rest.BreakdownController
			controller.Init(c)
			controller.FindByCategory(category_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.DELETE("/breakdown/bycategory", func(c *gin.Context) {
			item_ := &models.Breakdown{}
			c.ShouldBindJSON(item_)
			var controller rest.BreakdownController
			controller.Init(c)
			controller.DeleteByCategory(item_.Category)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/breakdown/find/method/:method", func(c *gin.Context) {
			method_, _ := strconv.ParseInt(c.Param("method"), 10, 64)
			var controller rest.BreakdownController
			controller.Init(c)
			controller.FindByMethod(method_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.PUT("/breakdown/subcategorybycategory", func(c *gin.Context) {
			var results map[string]any
			jsonData, _ := c.GetRawData()
			json.Unmarshal(jsonData, &results)
			var subcategory_ int64
			if v, flag := results["subcategory"]; flag {
				subcategory_ = int64(v.(float64))
			}
			var category_ int64
			if v, flag := results["category"]; flag {
				category_ = int64(v.(float64))
			}
			var controller rest.BreakdownController
			controller.Init(c)
			controller.UpdateSubcategoryByCategory(subcategory_, category_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.POST("/breakdown", func(c *gin.Context) {
			item_ := &models.Breakdown{}
			c.ShouldBindJSON(item_)
			var apicontroller api.BreakdownController
			apicontroller.Init(c)
			var controller rest.BreakdownController
			controller.Init(c)
			apicontroller.Pre_Insert(item_)
			controller.Insert(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.POST("/breakdown/batch", func(c *gin.Context) {
			item_ := &[]models.Breakdown{}
			c.ShouldBindJSON(item_)
			var apicontroller api.BreakdownController
			apicontroller.Init(c)
			var controller rest.BreakdownController
			controller.Init(c)
			apicontroller.Pre_Insertbatch(item_)
			controller.Insertbatch(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.PUT("/breakdown", func(c *gin.Context) {
			item_ := &models.Breakdown{}
			c.ShouldBindJSON(item_)
			var apicontroller api.BreakdownController
			apicontroller.Init(c)
			var controller rest.BreakdownController
			controller.Init(c)
			apicontroller.Pre_Update(item_)
			controller.Update(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.DELETE("/breakdown", func(c *gin.Context) {
			item_ := &models.Breakdown{}
			c.ShouldBindJSON(item_)
			var apicontroller api.BreakdownController
			apicontroller.Init(c)
			var controller rest.BreakdownController
			controller.Init(c)
			apicontroller.Pre_Delete(item_)
			controller.Delete(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.DELETE("/breakdown/batch", func(c *gin.Context) {
			item_ := &[]models.Breakdown{}
			c.ShouldBindJSON(item_)
			var apicontroller api.BreakdownController
			apicontroller.Init(c)
			var controller rest.BreakdownController
			controller.Init(c)
			apicontroller.Pre_Deletebatch(item_)
			controller.Deletebatch(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/breakdown/count", func(c *gin.Context) {
			var controller rest.BreakdownController
			controller.Init(c)
			controller.Count()
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/breakdown/:id", func(c *gin.Context) {
			id_, _ := strconv.ParseInt(c.Param("id"), 10, 64)
			var controller rest.BreakdownController
			controller.Init(c)
			controller.Read(id_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/breakdown", func(c *gin.Context) {
			page_, _ := strconv.Atoi(c.Query("page"))
			pagesize_, _ := strconv.Atoi(c.Query("pagesize"))
			var controller rest.BreakdownController
			controller.Init(c)
			controller.Index(page_, pagesize_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/breakdown/sum", func(c *gin.Context) {
			var controller rest.BreakdownController
			controller.Init(c)
			controller.Sum()
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/breakdownhistory/count/apt/:apt", func(c *gin.Context) {
			apt_, _ := strconv.ParseInt(c.Param("apt"), 10, 64)
			var controller rest.BreakdownhistoryController
			controller.Init(c)
			controller.CountByApt(apt_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/breakdownhistory/find/apt/:apt", func(c *gin.Context) {
			apt_, _ := strconv.ParseInt(c.Param("apt"), 10, 64)
			var controller rest.BreakdownhistoryController
			controller.Init(c)
			controller.FindByApt(apt_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.DELETE("/breakdownhistory/byapt", func(c *gin.Context) {
			item_ := &models.Breakdownhistory{}
			c.ShouldBindJSON(item_)
			var controller rest.BreakdownhistoryController
			controller.Init(c)
			controller.DeleteByApt(item_.Apt)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/breakdownhistory/count/breakdown/:breakdown", func(c *gin.Context) {
			breakdown_, _ := strconv.ParseInt(c.Param("breakdown"), 10, 64)
			var controller rest.BreakdownhistoryController
			controller.Init(c)
			controller.CountByBreakdown(breakdown_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/breakdownhistory/get/breakdown/:breakdown", func(c *gin.Context) {
			breakdown_, _ := strconv.ParseInt(c.Param("breakdown"), 10, 64)
			var controller rest.BreakdownhistoryController
			controller.Init(c)
			controller.GetByBreakdown(breakdown_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/breakdownhistory/find/category/:category", func(c *gin.Context) {
			category_, _ := strconv.ParseInt(c.Param("category"), 10, 64)
			var controller rest.BreakdownhistoryController
			controller.Init(c)
			controller.FindByCategory(category_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.DELETE("/breakdownhistory/bycategory", func(c *gin.Context) {
			item_ := &models.Breakdownhistory{}
			c.ShouldBindJSON(item_)
			var controller rest.BreakdownhistoryController
			controller.Init(c)
			controller.DeleteByCategory(item_.Category)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/breakdownhistory/find/method/:method", func(c *gin.Context) {
			method_, _ := strconv.ParseInt(c.Param("method"), 10, 64)
			var controller rest.BreakdownhistoryController
			controller.Init(c)
			controller.FindByMethod(method_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.POST("/breakdownhistory", func(c *gin.Context) {
			item_ := &models.Breakdownhistory{}
			c.ShouldBindJSON(item_)
			var controller rest.BreakdownhistoryController
			controller.Init(c)
			controller.Insert(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.POST("/breakdownhistory/batch", func(c *gin.Context) {
			item_ := &[]models.Breakdownhistory{}
			c.ShouldBindJSON(item_)
			var controller rest.BreakdownhistoryController
			controller.Init(c)
			controller.Insertbatch(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.PUT("/breakdownhistory", func(c *gin.Context) {
			item_ := &models.Breakdownhistory{}
			c.ShouldBindJSON(item_)
			var controller rest.BreakdownhistoryController
			controller.Init(c)
			controller.Update(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.DELETE("/breakdownhistory", func(c *gin.Context) {
			item_ := &models.Breakdownhistory{}
			c.ShouldBindJSON(item_)
			var controller rest.BreakdownhistoryController
			controller.Init(c)
			controller.Delete(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.DELETE("/breakdownhistory/batch", func(c *gin.Context) {
			item_ := &[]models.Breakdownhistory{}
			c.ShouldBindJSON(item_)
			var controller rest.BreakdownhistoryController
			controller.Init(c)
			controller.Deletebatch(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/breakdownhistory/count", func(c *gin.Context) {
			var controller rest.BreakdownhistoryController
			controller.Init(c)
			controller.Count()
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/breakdownhistory/:id", func(c *gin.Context) {
			id_, _ := strconv.ParseInt(c.Param("id"), 10, 64)
			var controller rest.BreakdownhistoryController
			controller.Init(c)
			controller.Read(id_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/breakdownhistory", func(c *gin.Context) {
			page_, _ := strconv.Atoi(c.Query("page"))
			pagesize_, _ := strconv.Atoi(c.Query("pagesize"))
			var controller rest.BreakdownhistoryController
			controller.Init(c)
			controller.Index(page_, pagesize_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/breakdownhistory/sum", func(c *gin.Context) {
			var controller rest.BreakdownhistoryController
			controller.Init(c)
			controller.Sum()
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/category/get/levelname/:level", func(c *gin.Context) {
			level_, _ := strconv.Atoi(c.Param("level"))
			name_ := c.Query("name")
			var controller rest.CategoryController
			controller.Init(c)
			controller.GetByLevelName(level_, name_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/category/get/levelparentname/:level", func(c *gin.Context) {
			level_, _ := strconv.Atoi(c.Param("level"))
			parent_, _ := strconv.ParseInt(c.Query("parent"), 10, 64)
			name_ := c.Query("name")
			var controller rest.CategoryController
			controller.Init(c)
			controller.GetByLevelParentName(level_, parent_, name_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/category/count/apt/:apt", func(c *gin.Context) {
			apt_, _ := strconv.ParseInt(c.Param("apt"), 10, 64)
			var controller rest.CategoryController
			controller.Init(c)
			controller.CountByApt(apt_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/category/find/apt/:apt", func(c *gin.Context) {
			apt_, _ := strconv.ParseInt(c.Param("apt"), 10, 64)
			var controller rest.CategoryController
			controller.Init(c)
			controller.FindByApt(apt_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/category/count/aptlevel/:apt", func(c *gin.Context) {
			apt_, _ := strconv.ParseInt(c.Param("apt"), 10, 64)
			level_, _ := strconv.Atoi(c.Query("level"))
			var controller rest.CategoryController
			controller.Init(c)
			controller.CountByAptLevel(apt_, level_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/category/find/aptlevel/:apt", func(c *gin.Context) {
			apt_, _ := strconv.ParseInt(c.Param("apt"), 10, 64)
			level_, _ := strconv.Atoi(c.Query("level"))
			var controller rest.CategoryController
			controller.Init(c)
			controller.FindByAptLevel(apt_, level_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/category/get/aptlevelparentname/:apt", func(c *gin.Context) {
			apt_, _ := strconv.ParseInt(c.Param("apt"), 10, 64)
			level_, _ := strconv.Atoi(c.Query("level"))
			parent_, _ := strconv.ParseInt(c.Query("parent"), 10, 64)
			name_ := c.Query("name")
			var controller rest.CategoryController
			controller.Init(c)
			controller.GetByAptLevelParentName(apt_, level_, parent_, name_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/category/get/aptname/:apt", func(c *gin.Context) {
			apt_, _ := strconv.ParseInt(c.Param("apt"), 10, 64)
			name_ := c.Query("name")
			var controller rest.CategoryController
			controller.Init(c)
			controller.GetByAptName(apt_, name_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/category/count/aptparent/:apt", func(c *gin.Context) {
			apt_, _ := strconv.ParseInt(c.Param("apt"), 10, 64)
			parent_, _ := strconv.ParseInt(c.Query("parent"), 10, 64)
			var controller rest.CategoryController
			controller.Init(c)
			controller.CountByAptParent(apt_, parent_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/category/find/aptparent/:apt", func(c *gin.Context) {
			apt_, _ := strconv.ParseInt(c.Param("apt"), 10, 64)
			parent_, _ := strconv.ParseInt(c.Query("parent"), 10, 64)
			var controller rest.CategoryController
			controller.Init(c)
			controller.FindByAptParent(apt_, parent_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/category/find/aptorder/:apt", func(c *gin.Context) {
			apt_, _ := strconv.ParseInt(c.Param("apt"), 10, 64)
			order_, _ := strconv.Atoi(c.Query("order"))
			var controller rest.CategoryController
			controller.Init(c)
			controller.FindByAptOrder(apt_, order_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.POST("/category", func(c *gin.Context) {
			item_ := &models.Category{}
			c.ShouldBindJSON(item_)
			var controller rest.CategoryController
			controller.Init(c)
			controller.Insert(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.POST("/category/batch", func(c *gin.Context) {
			item_ := &[]models.Category{}
			c.ShouldBindJSON(item_)
			var controller rest.CategoryController
			controller.Init(c)
			controller.Insertbatch(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.PUT("/category", func(c *gin.Context) {
			item_ := &models.Category{}
			c.ShouldBindJSON(item_)
			var controller rest.CategoryController
			controller.Init(c)
			controller.Update(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.DELETE("/category", func(c *gin.Context) {
			item_ := &models.Category{}
			c.ShouldBindJSON(item_)
			var controller rest.CategoryController
			controller.Init(c)
			controller.Delete(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.DELETE("/category/batch", func(c *gin.Context) {
			item_ := &[]models.Category{}
			c.ShouldBindJSON(item_)
			var controller rest.CategoryController
			controller.Init(c)
			controller.Deletebatch(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/category/count", func(c *gin.Context) {
			var controller rest.CategoryController
			controller.Init(c)
			controller.Count()
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/category/:id", func(c *gin.Context) {
			id_, _ := strconv.ParseInt(c.Param("id"), 10, 64)
			var controller rest.CategoryController
			controller.Init(c)
			controller.Read(id_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/category", func(c *gin.Context) {
			page_, _ := strconv.Atoi(c.Query("page"))
			pagesize_, _ := strconv.Atoi(c.Query("pagesize"))
			var controller rest.CategoryController
			controller.Init(c)
			controller.Index(page_, pagesize_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.POST("/comparecompany", func(c *gin.Context) {
			item_ := &models.Comparecompany{}
			c.ShouldBindJSON(item_)
			var controller rest.ComparecompanyController
			controller.Init(c)
			controller.Insert(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.POST("/comparecompany/batch", func(c *gin.Context) {
			item_ := &[]models.Comparecompany{}
			c.ShouldBindJSON(item_)
			var controller rest.ComparecompanyController
			controller.Init(c)
			controller.Insertbatch(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.PUT("/comparecompany", func(c *gin.Context) {
			item_ := &models.Comparecompany{}
			c.ShouldBindJSON(item_)
			var controller rest.ComparecompanyController
			controller.Init(c)
			controller.Update(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.DELETE("/comparecompany", func(c *gin.Context) {
			item_ := &models.Comparecompany{}
			c.ShouldBindJSON(item_)
			var controller rest.ComparecompanyController
			controller.Init(c)
			controller.Delete(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.DELETE("/comparecompany/batch", func(c *gin.Context) {
			item_ := &[]models.Comparecompany{}
			c.ShouldBindJSON(item_)
			var controller rest.ComparecompanyController
			controller.Init(c)
			controller.Deletebatch(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/comparecompany/count", func(c *gin.Context) {
			var controller rest.ComparecompanyController
			controller.Init(c)
			controller.Count()
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/comparecompany/:id", func(c *gin.Context) {
			id_, _ := strconv.ParseInt(c.Param("id"), 10, 64)
			var controller rest.ComparecompanyController
			controller.Init(c)
			controller.Read(id_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/comparecompany", func(c *gin.Context) {
			page_, _ := strconv.Atoi(c.Query("page"))
			pagesize_, _ := strconv.Atoi(c.Query("pagesize"))
			var controller rest.ComparecompanyController
			controller.Init(c)
			controller.Index(page_, pagesize_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/compareestimate/find/estimate/:estimate", func(c *gin.Context) {
			estimate_, _ := strconv.ParseInt(c.Param("estimate"), 10, 64)
			var controller rest.CompareestimateController
			controller.Init(c)
			controller.FindByEstimate(estimate_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.DELETE("/compareestimate/byestimate", func(c *gin.Context) {
			item_ := &models.Compareestimate{}
			c.ShouldBindJSON(item_)
			var controller rest.CompareestimateController
			controller.Init(c)
			controller.DeleteByEstimate(item_.Estimate)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.POST("/compareestimate", func(c *gin.Context) {
			item_ := &models.Compareestimate{}
			c.ShouldBindJSON(item_)
			var controller rest.CompareestimateController
			controller.Init(c)
			controller.Insert(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.POST("/compareestimate/batch", func(c *gin.Context) {
			item_ := &[]models.Compareestimate{}
			c.ShouldBindJSON(item_)
			var controller rest.CompareestimateController
			controller.Init(c)
			controller.Insertbatch(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.PUT("/compareestimate", func(c *gin.Context) {
			item_ := &models.Compareestimate{}
			c.ShouldBindJSON(item_)
			var controller rest.CompareestimateController
			controller.Init(c)
			controller.Update(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.DELETE("/compareestimate", func(c *gin.Context) {
			item_ := &models.Compareestimate{}
			c.ShouldBindJSON(item_)
			var controller rest.CompareestimateController
			controller.Init(c)
			controller.Delete(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.DELETE("/compareestimate/batch", func(c *gin.Context) {
			item_ := &[]models.Compareestimate{}
			c.ShouldBindJSON(item_)
			var controller rest.CompareestimateController
			controller.Init(c)
			controller.Deletebatch(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/compareestimate/count", func(c *gin.Context) {
			var controller rest.CompareestimateController
			controller.Init(c)
			controller.Count()
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/compareestimate/:id", func(c *gin.Context) {
			id_, _ := strconv.ParseInt(c.Param("id"), 10, 64)
			var controller rest.CompareestimateController
			controller.Init(c)
			controller.Read(id_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/compareestimate", func(c *gin.Context) {
			page_, _ := strconv.Atoi(c.Query("page"))
			pagesize_, _ := strconv.Atoi(c.Query("pagesize"))
			var controller rest.CompareestimateController
			controller.Init(c)
			controller.Index(page_, pagesize_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/compareestimate/sum", func(c *gin.Context) {
			var controller rest.CompareestimateController
			controller.Init(c)
			controller.Sum()
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/contract/get/estimate/:estimate", func(c *gin.Context) {
			estimate_, _ := strconv.ParseInt(c.Param("estimate"), 10, 64)
			var controller rest.ContractController
			controller.Init(c)
			controller.GetByEstimate(estimate_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.POST("/contract", func(c *gin.Context) {
			item_ := &models.Contract{}
			c.ShouldBindJSON(item_)
			var apicontroller api.ContractController
			apicontroller.Init(c)
			var controller rest.ContractController
			controller.Init(c)
			controller.Insert(item_)
			apicontroller.Post_Insert(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.POST("/contract/batch", func(c *gin.Context) {
			item_ := &[]models.Contract{}
			c.ShouldBindJSON(item_)
			var controller rest.ContractController
			controller.Init(c)
			controller.Insertbatch(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.PUT("/contract", func(c *gin.Context) {
			item_ := &models.Contract{}
			c.ShouldBindJSON(item_)
			var apicontroller api.ContractController
			apicontroller.Init(c)
			var controller rest.ContractController
			controller.Init(c)
			controller.Update(item_)
			apicontroller.Post_Update(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.DELETE("/contract", func(c *gin.Context) {
			item_ := &models.Contract{}
			c.ShouldBindJSON(item_)
			var apicontroller api.ContractController
			apicontroller.Init(c)
			var controller rest.ContractController
			controller.Init(c)
			controller.Delete(item_)
			apicontroller.Post_Delete(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.DELETE("/contract/batch", func(c *gin.Context) {
			item_ := &[]models.Contract{}
			c.ShouldBindJSON(item_)
			var apicontroller api.ContractController
			apicontroller.Init(c)
			var controller rest.ContractController
			controller.Init(c)
			controller.Deletebatch(item_)
			apicontroller.Post_Deletebatch(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/contract/count", func(c *gin.Context) {
			var controller rest.ContractController
			controller.Init(c)
			controller.Count()
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/contract/:id", func(c *gin.Context) {
			id_, _ := strconv.ParseInt(c.Param("id"), 10, 64)
			var controller rest.ContractController
			controller.Init(c)
			controller.Read(id_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/contract", func(c *gin.Context) {
			page_, _ := strconv.Atoi(c.Query("page"))
			pagesize_, _ := strconv.Atoi(c.Query("pagesize"))
			var controller rest.ContractController
			controller.Init(c)
			controller.Index(page_, pagesize_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/contract/sum", func(c *gin.Context) {
			var controller rest.ContractController
			controller.Init(c)
			controller.Sum()
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.POST("/datacategory", func(c *gin.Context) {
			item_ := &models.Datacategory{}
			c.ShouldBindJSON(item_)
			var controller rest.DatacategoryController
			controller.Init(c)
			controller.Insert(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.POST("/datacategory/batch", func(c *gin.Context) {
			item_ := &[]models.Datacategory{}
			c.ShouldBindJSON(item_)
			var controller rest.DatacategoryController
			controller.Init(c)
			controller.Insertbatch(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.PUT("/datacategory", func(c *gin.Context) {
			item_ := &models.Datacategory{}
			c.ShouldBindJSON(item_)
			var controller rest.DatacategoryController
			controller.Init(c)
			controller.Update(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.DELETE("/datacategory", func(c *gin.Context) {
			item_ := &models.Datacategory{}
			c.ShouldBindJSON(item_)
			var controller rest.DatacategoryController
			controller.Init(c)
			controller.Delete(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.DELETE("/datacategory/batch", func(c *gin.Context) {
			item_ := &[]models.Datacategory{}
			c.ShouldBindJSON(item_)
			var controller rest.DatacategoryController
			controller.Init(c)
			controller.Deletebatch(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/datacategory/count", func(c *gin.Context) {
			var controller rest.DatacategoryController
			controller.Init(c)
			controller.Count()
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/datacategory/:id", func(c *gin.Context) {
			id_, _ := strconv.ParseInt(c.Param("id"), 10, 64)
			var controller rest.DatacategoryController
			controller.Init(c)
			controller.Read(id_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/datacategory", func(c *gin.Context) {
			page_, _ := strconv.Atoi(c.Query("page"))
			pagesize_, _ := strconv.Atoi(c.Query("pagesize"))
			var controller rest.DatacategoryController
			controller.Init(c)
			controller.Index(page_, pagesize_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.POST("/detail", func(c *gin.Context) {
			item_ := &models.Detail{}
			c.ShouldBindJSON(item_)
			var apicontroller api.DetailController
			apicontroller.Init(c)
			var controller rest.DetailController
			controller.Init(c)
			controller.Insert(item_)
			apicontroller.Post_Insert(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.POST("/detail/batch", func(c *gin.Context) {
			item_ := &[]models.Detail{}
			c.ShouldBindJSON(item_)
			var controller rest.DetailController
			controller.Init(c)
			controller.Insertbatch(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.PUT("/detail", func(c *gin.Context) {
			item_ := &models.Detail{}
			c.ShouldBindJSON(item_)
			var controller rest.DetailController
			controller.Init(c)
			controller.Update(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.DELETE("/detail", func(c *gin.Context) {
			item_ := &models.Detail{}
			c.ShouldBindJSON(item_)
			var controller rest.DetailController
			controller.Init(c)
			controller.Delete(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.DELETE("/detail/batch", func(c *gin.Context) {
			item_ := &[]models.Detail{}
			c.ShouldBindJSON(item_)
			var controller rest.DetailController
			controller.Init(c)
			controller.Deletebatch(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/detail/count", func(c *gin.Context) {
			var controller rest.DetailController
			controller.Init(c)
			controller.Count()
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/detail/:id", func(c *gin.Context) {
			id_, _ := strconv.ParseInt(c.Param("id"), 10, 64)
			var controller rest.DetailController
			controller.Init(c)
			controller.Read(id_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/detail", func(c *gin.Context) {
			page_, _ := strconv.Atoi(c.Query("page"))
			pagesize_, _ := strconv.Atoi(c.Query("pagesize"))
			var controller rest.DetailController
			controller.Init(c)
			controller.Index(page_, pagesize_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/detail/sum", func(c *gin.Context) {
			var controller rest.DetailController
			controller.Init(c)
			controller.Sum()
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/detailtechnician/count/detail/:detail", func(c *gin.Context) {
			detail_, _ := strconv.ParseInt(c.Param("detail"), 10, 64)
			var controller rest.DetailtechnicianController
			controller.Init(c)
			controller.CountByDetail(detail_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/detailtechnician/find/detail/:detail", func(c *gin.Context) {
			detail_, _ := strconv.ParseInt(c.Param("detail"), 10, 64)
			var controller rest.DetailtechnicianController
			controller.Init(c)
			controller.FindByDetail(detail_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.DELETE("/detailtechnician/bydetail", func(c *gin.Context) {
			item_ := &models.Detailtechnician{}
			c.ShouldBindJSON(item_)
			var controller rest.DetailtechnicianController
			controller.Init(c)
			controller.DeleteByDetail(item_.Detail)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.POST("/detailtechnician", func(c *gin.Context) {
			item_ := &models.Detailtechnician{}
			c.ShouldBindJSON(item_)
			var controller rest.DetailtechnicianController
			controller.Init(c)
			controller.Insert(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.POST("/detailtechnician/batch", func(c *gin.Context) {
			item_ := &[]models.Detailtechnician{}
			c.ShouldBindJSON(item_)
			var controller rest.DetailtechnicianController
			controller.Init(c)
			controller.Insertbatch(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.PUT("/detailtechnician", func(c *gin.Context) {
			item_ := &models.Detailtechnician{}
			c.ShouldBindJSON(item_)
			var controller rest.DetailtechnicianController
			controller.Init(c)
			controller.Update(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.DELETE("/detailtechnician", func(c *gin.Context) {
			item_ := &models.Detailtechnician{}
			c.ShouldBindJSON(item_)
			var controller rest.DetailtechnicianController
			controller.Init(c)
			controller.Delete(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.DELETE("/detailtechnician/batch", func(c *gin.Context) {
			item_ := &[]models.Detailtechnician{}
			c.ShouldBindJSON(item_)
			var controller rest.DetailtechnicianController
			controller.Init(c)
			controller.Deletebatch(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/detailtechnician/count", func(c *gin.Context) {
			var controller rest.DetailtechnicianController
			controller.Init(c)
			controller.Count()
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/detailtechnician/:id", func(c *gin.Context) {
			id_, _ := strconv.ParseInt(c.Param("id"), 10, 64)
			var controller rest.DetailtechnicianController
			controller.Init(c)
			controller.Read(id_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/detailtechnician", func(c *gin.Context) {
			page_, _ := strconv.Atoi(c.Query("page"))
			pagesize_, _ := strconv.Atoi(c.Query("pagesize"))
			var controller rest.DetailtechnicianController
			controller.Init(c)
			controller.Index(page_, pagesize_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/dong/count/apt/:apt", func(c *gin.Context) {
			apt_, _ := strconv.ParseInt(c.Param("apt"), 10, 64)
			var controller rest.DongController
			controller.Init(c)
			controller.CountByApt(apt_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/dong/find/apt/:apt", func(c *gin.Context) {
			apt_, _ := strconv.ParseInt(c.Param("apt"), 10, 64)
			var controller rest.DongController
			controller.Init(c)
			controller.FindByApt(apt_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.DELETE("/dong/byapt", func(c *gin.Context) {
			item_ := &models.Dong{}
			c.ShouldBindJSON(item_)
			var controller rest.DongController
			controller.Init(c)
			controller.DeleteByApt(item_.Apt)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.POST("/dong", func(c *gin.Context) {
			item_ := &models.Dong{}
			c.ShouldBindJSON(item_)
			var controller rest.DongController
			controller.Init(c)
			controller.Insert(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.POST("/dong/batch", func(c *gin.Context) {
			item_ := &[]models.Dong{}
			c.ShouldBindJSON(item_)
			var controller rest.DongController
			controller.Init(c)
			controller.Insertbatch(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.PUT("/dong", func(c *gin.Context) {
			item_ := &models.Dong{}
			c.ShouldBindJSON(item_)
			var controller rest.DongController
			controller.Init(c)
			controller.Update(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.DELETE("/dong", func(c *gin.Context) {
			item_ := &models.Dong{}
			c.ShouldBindJSON(item_)
			var apicontroller api.DongController
			apicontroller.Init(c)
			var controller rest.DongController
			controller.Init(c)
			controller.Delete(item_)
			apicontroller.Post_Delete(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.DELETE("/dong/batch", func(c *gin.Context) {
			item_ := &[]models.Dong{}
			c.ShouldBindJSON(item_)
			var apicontroller api.DongController
			apicontroller.Init(c)
			var controller rest.DongController
			controller.Init(c)
			controller.Deletebatch(item_)
			apicontroller.Post_Deletebatch(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/dong/count", func(c *gin.Context) {
			var controller rest.DongController
			controller.Init(c)
			controller.Count()
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/dong/:id", func(c *gin.Context) {
			id_, _ := strconv.ParseInt(c.Param("id"), 10, 64)
			var controller rest.DongController
			controller.Init(c)
			controller.Read(id_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/dong", func(c *gin.Context) {
			page_, _ := strconv.Atoi(c.Query("page"))
			pagesize_, _ := strconv.Atoi(c.Query("pagesize"))
			var controller rest.DongController
			controller.Init(c)
			controller.Index(page_, pagesize_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.POST("/estimate/batch", func(c *gin.Context) {
			item_ := &[]models.Estimate{}
			c.ShouldBindJSON(item_)
			var controller rest.EstimateController
			controller.Init(c)
			controller.Insertbatch(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.DELETE("/estimate/batch", func(c *gin.Context) {
			item_ := &[]models.Estimate{}
			c.ShouldBindJSON(item_)
			var controller rest.EstimateController
			controller.Init(c)
			controller.Deletebatch(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/estimate/count", func(c *gin.Context) {
			var controller rest.EstimateController
			controller.Init(c)
			controller.Count()
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/estimate/:id", func(c *gin.Context) {
			id_, _ := strconv.ParseInt(c.Param("id"), 10, 64)
			var controller rest.EstimateController
			controller.Init(c)
			controller.Read(id_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/estimate/sum", func(c *gin.Context) {
			var controller rest.EstimateController
			controller.Init(c)
			controller.Sum()
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.POST("/facilitycategory", func(c *gin.Context) {
			item_ := &models.Facilitycategory{}
			c.ShouldBindJSON(item_)
			var controller rest.FacilitycategoryController
			controller.Init(c)
			controller.Insert(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.POST("/facilitycategory/batch", func(c *gin.Context) {
			item_ := &[]models.Facilitycategory{}
			c.ShouldBindJSON(item_)
			var controller rest.FacilitycategoryController
			controller.Init(c)
			controller.Insertbatch(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.PUT("/facilitycategory", func(c *gin.Context) {
			item_ := &models.Facilitycategory{}
			c.ShouldBindJSON(item_)
			var controller rest.FacilitycategoryController
			controller.Init(c)
			controller.Update(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.DELETE("/facilitycategory", func(c *gin.Context) {
			item_ := &models.Facilitycategory{}
			c.ShouldBindJSON(item_)
			var controller rest.FacilitycategoryController
			controller.Init(c)
			controller.Delete(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.DELETE("/facilitycategory/batch", func(c *gin.Context) {
			item_ := &[]models.Facilitycategory{}
			c.ShouldBindJSON(item_)
			var controller rest.FacilitycategoryController
			controller.Init(c)
			controller.Deletebatch(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/facilitycategory/count", func(c *gin.Context) {
			var controller rest.FacilitycategoryController
			controller.Init(c)
			controller.Count()
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/facilitycategory/:id", func(c *gin.Context) {
			id_, _ := strconv.ParseInt(c.Param("id"), 10, 64)
			var controller rest.FacilitycategoryController
			controller.Init(c)
			controller.Read(id_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/facilitycategory", func(c *gin.Context) {
			page_, _ := strconv.Atoi(c.Query("page"))
			pagesize_, _ := strconv.Atoi(c.Query("pagesize"))
			var controller rest.FacilitycategoryController
			controller.Init(c)
			controller.Index(page_, pagesize_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/file/count/apt/:apt", func(c *gin.Context) {
			apt_, _ := strconv.ParseInt(c.Param("apt"), 10, 64)
			var controller rest.FileController
			controller.Init(c)
			controller.CountByApt(apt_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/file/find/apt/:apt", func(c *gin.Context) {
			apt_, _ := strconv.ParseInt(c.Param("apt"), 10, 64)
			var controller rest.FileController
			controller.Init(c)
			controller.FindByApt(apt_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.DELETE("/file/byapt", func(c *gin.Context) {
			item_ := &models.File{}
			c.ShouldBindJSON(item_)
			var controller rest.FileController
			controller.Init(c)
			controller.DeleteByApt(item_.Apt)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.POST("/file", func(c *gin.Context) {
			item_ := &models.File{}
			c.ShouldBindJSON(item_)
			var controller rest.FileController
			controller.Init(c)
			controller.Insert(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.POST("/file/batch", func(c *gin.Context) {
			item_ := &[]models.File{}
			c.ShouldBindJSON(item_)
			var controller rest.FileController
			controller.Init(c)
			controller.Insertbatch(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.PUT("/file", func(c *gin.Context) {
			item_ := &models.File{}
			c.ShouldBindJSON(item_)
			var controller rest.FileController
			controller.Init(c)
			controller.Update(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.DELETE("/file", func(c *gin.Context) {
			item_ := &models.File{}
			c.ShouldBindJSON(item_)
			var apicontroller api.FileController
			apicontroller.Init(c)
			var controller rest.FileController
			controller.Init(c)
			apicontroller.Pre_Delete(item_)
			controller.Delete(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.DELETE("/file/batch", func(c *gin.Context) {
			item_ := &[]models.File{}
			c.ShouldBindJSON(item_)
			var controller rest.FileController
			controller.Init(c)
			controller.Deletebatch(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/file/count", func(c *gin.Context) {
			var controller rest.FileController
			controller.Init(c)
			controller.Count()
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/file/:id", func(c *gin.Context) {
			id_, _ := strconv.ParseInt(c.Param("id"), 10, 64)
			var controller rest.FileController
			controller.Init(c)
			controller.Read(id_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/file", func(c *gin.Context) {
			page_, _ := strconv.Atoi(c.Query("page"))
			pagesize_, _ := strconv.Atoi(c.Query("pagesize"))
			var controller rest.FileController
			controller.Init(c)
			controller.Index(page_, pagesize_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/history/count/apt/:apt", func(c *gin.Context) {
			apt_, _ := strconv.ParseInt(c.Param("apt"), 10, 64)
			var controller rest.HistoryController
			controller.Init(c)
			controller.CountByApt(apt_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/history/find/apt/:apt", func(c *gin.Context) {
			apt_, _ := strconv.ParseInt(c.Param("apt"), 10, 64)
			var controller rest.HistoryController
			controller.Init(c)
			controller.FindByApt(apt_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.DELETE("/history/byapt", func(c *gin.Context) {
			item_ := &models.History{}
			c.ShouldBindJSON(item_)
			var controller rest.HistoryController
			controller.Init(c)
			controller.DeleteByApt(item_.Apt)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/history/find/category/:category", func(c *gin.Context) {
			category_, _ := strconv.ParseInt(c.Param("category"), 10, 64)
			var controller rest.HistoryController
			controller.Init(c)
			controller.FindByCategory(category_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.DELETE("/history/bycategory", func(c *gin.Context) {
			item_ := &models.History{}
			c.ShouldBindJSON(item_)
			var controller rest.HistoryController
			controller.Init(c)
			controller.DeleteByCategory(item_.Category)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.POST("/history", func(c *gin.Context) {
			item_ := &models.History{}
			c.ShouldBindJSON(item_)
			var controller rest.HistoryController
			controller.Init(c)
			controller.Insert(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.POST("/history/batch", func(c *gin.Context) {
			item_ := &[]models.History{}
			c.ShouldBindJSON(item_)
			var controller rest.HistoryController
			controller.Init(c)
			controller.Insertbatch(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.PUT("/history", func(c *gin.Context) {
			item_ := &models.History{}
			c.ShouldBindJSON(item_)
			var controller rest.HistoryController
			controller.Init(c)
			controller.Update(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.DELETE("/history", func(c *gin.Context) {
			item_ := &models.History{}
			c.ShouldBindJSON(item_)
			var controller rest.HistoryController
			controller.Init(c)
			controller.Delete(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.DELETE("/history/batch", func(c *gin.Context) {
			item_ := &[]models.History{}
			c.ShouldBindJSON(item_)
			var controller rest.HistoryController
			controller.Init(c)
			controller.Deletebatch(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/history/count", func(c *gin.Context) {
			var controller rest.HistoryController
			controller.Init(c)
			controller.Count()
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/history/:id", func(c *gin.Context) {
			id_, _ := strconv.ParseInt(c.Param("id"), 10, 64)
			var controller rest.HistoryController
			controller.Init(c)
			controller.Read(id_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/history", func(c *gin.Context) {
			page_, _ := strconv.Atoi(c.Query("page"))
			pagesize_, _ := strconv.Atoi(c.Query("pagesize"))
			var controller rest.HistoryController
			controller.Init(c)
			controller.Index(page_, pagesize_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/history/sum", func(c *gin.Context) {
			var controller rest.HistoryController
			controller.Init(c)
			controller.Sum()
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.POST("/inquiry", func(c *gin.Context) {
			item_ := &models.Inquiry{}
			c.ShouldBindJSON(item_)
			var controller rest.InquiryController
			controller.Init(c)
			controller.Insert(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.POST("/inquiry/batch", func(c *gin.Context) {
			item_ := &[]models.Inquiry{}
			c.ShouldBindJSON(item_)
			var controller rest.InquiryController
			controller.Init(c)
			controller.Insertbatch(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.PUT("/inquiry", func(c *gin.Context) {
			item_ := &models.Inquiry{}
			c.ShouldBindJSON(item_)
			var controller rest.InquiryController
			controller.Init(c)
			controller.Update(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.DELETE("/inquiry", func(c *gin.Context) {
			item_ := &models.Inquiry{}
			c.ShouldBindJSON(item_)
			var controller rest.InquiryController
			controller.Init(c)
			controller.Delete(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.DELETE("/inquiry/batch", func(c *gin.Context) {
			item_ := &[]models.Inquiry{}
			c.ShouldBindJSON(item_)
			var controller rest.InquiryController
			controller.Init(c)
			controller.Deletebatch(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/inquiry/count", func(c *gin.Context) {
			var controller rest.InquiryController
			controller.Init(c)
			controller.Count()
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/inquiry/:id", func(c *gin.Context) {
			id_, _ := strconv.ParseInt(c.Param("id"), 10, 64)
			var controller rest.InquiryController
			controller.Init(c)
			controller.Read(id_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/inquiry", func(c *gin.Context) {
			page_, _ := strconv.Atoi(c.Query("page"))
			pagesize_, _ := strconv.Atoi(c.Query("pagesize"))
			var controller rest.InquiryController
			controller.Init(c)
			controller.Index(page_, pagesize_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.DELETE("/managebook/byperiodic", func(c *gin.Context) {
			item_ := &models.Managebook{}
			c.ShouldBindJSON(item_)
			var controller rest.ManagebookController
			controller.Init(c)
			controller.DeleteByPeriodic(item_.Periodic)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.DELETE("/managebook/bymanagebookcategory", func(c *gin.Context) {
			item_ := &models.Managebook{}
			c.ShouldBindJSON(item_)
			var controller rest.ManagebookController
			controller.Init(c)
			controller.DeleteByManagebookcategory(item_.Managebookcategory)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.POST("/managebook", func(c *gin.Context) {
			item_ := &models.Managebook{}
			c.ShouldBindJSON(item_)
			var controller rest.ManagebookController
			controller.Init(c)
			controller.Insert(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.POST("/managebook/batch", func(c *gin.Context) {
			item_ := &[]models.Managebook{}
			c.ShouldBindJSON(item_)
			var controller rest.ManagebookController
			controller.Init(c)
			controller.Insertbatch(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.PUT("/managebook", func(c *gin.Context) {
			item_ := &models.Managebook{}
			c.ShouldBindJSON(item_)
			var controller rest.ManagebookController
			controller.Init(c)
			controller.Update(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.DELETE("/managebook", func(c *gin.Context) {
			item_ := &models.Managebook{}
			c.ShouldBindJSON(item_)
			var apicontroller api.ManagebookController
			apicontroller.Init(c)
			var controller rest.ManagebookController
			controller.Init(c)
			apicontroller.Pre_Delete(item_)
			controller.Delete(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.DELETE("/managebook/batch", func(c *gin.Context) {
			item_ := &[]models.Managebook{}
			c.ShouldBindJSON(item_)
			var apicontroller api.ManagebookController
			apicontroller.Init(c)
			var controller rest.ManagebookController
			controller.Init(c)
			apicontroller.Pre_Deletebatch(item_)
			controller.Deletebatch(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/managebook/count", func(c *gin.Context) {
			var controller rest.ManagebookController
			controller.Init(c)
			controller.Count()
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/managebook/:id", func(c *gin.Context) {
			id_, _ := strconv.ParseInt(c.Param("id"), 10, 64)
			var controller rest.ManagebookController
			controller.Init(c)
			controller.Read(id_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/managebook", func(c *gin.Context) {
			page_, _ := strconv.Atoi(c.Query("page"))
			pagesize_, _ := strconv.Atoi(c.Query("pagesize"))
			var controller rest.ManagebookController
			controller.Init(c)
			controller.Index(page_, pagesize_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.POST("/managebookcategory", func(c *gin.Context) {
			item_ := &models.Managebookcategory{}
			c.ShouldBindJSON(item_)
			var controller rest.ManagebookcategoryController
			controller.Init(c)
			controller.Insert(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.POST("/managebookcategory/batch", func(c *gin.Context) {
			item_ := &[]models.Managebookcategory{}
			c.ShouldBindJSON(item_)
			var controller rest.ManagebookcategoryController
			controller.Init(c)
			controller.Insertbatch(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.PUT("/managebookcategory", func(c *gin.Context) {
			item_ := &models.Managebookcategory{}
			c.ShouldBindJSON(item_)
			var controller rest.ManagebookcategoryController
			controller.Init(c)
			controller.Update(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.DELETE("/managebookcategory", func(c *gin.Context) {
			item_ := &models.Managebookcategory{}
			c.ShouldBindJSON(item_)
			var controller rest.ManagebookcategoryController
			controller.Init(c)
			controller.Delete(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.DELETE("/managebookcategory/batch", func(c *gin.Context) {
			item_ := &[]models.Managebookcategory{}
			c.ShouldBindJSON(item_)
			var controller rest.ManagebookcategoryController
			controller.Init(c)
			controller.Deletebatch(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/managebookcategory/count", func(c *gin.Context) {
			var controller rest.ManagebookcategoryController
			controller.Init(c)
			controller.Count()
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/managebookcategory/:id", func(c *gin.Context) {
			id_, _ := strconv.ParseInt(c.Param("id"), 10, 64)
			var controller rest.ManagebookcategoryController
			controller.Init(c)
			controller.Read(id_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/managebookcategory", func(c *gin.Context) {
			page_, _ := strconv.Atoi(c.Query("page"))
			pagesize_, _ := strconv.Atoi(c.Query("pagesize"))
			var controller rest.ManagebookcategoryController
			controller.Init(c)
			controller.Index(page_, pagesize_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.POST("/oldapt", func(c *gin.Context) {
			item_ := &models.Oldapt{}
			c.ShouldBindJSON(item_)
			var controller rest.OldaptController
			controller.Init(c)
			controller.Insert(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.POST("/oldapt/batch", func(c *gin.Context) {
			item_ := &[]models.Oldapt{}
			c.ShouldBindJSON(item_)
			var controller rest.OldaptController
			controller.Init(c)
			controller.Insertbatch(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.PUT("/oldapt", func(c *gin.Context) {
			item_ := &models.Oldapt{}
			c.ShouldBindJSON(item_)
			var controller rest.OldaptController
			controller.Init(c)
			controller.Update(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.DELETE("/oldapt", func(c *gin.Context) {
			item_ := &models.Oldapt{}
			c.ShouldBindJSON(item_)
			var controller rest.OldaptController
			controller.Init(c)
			controller.Delete(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.DELETE("/oldapt/batch", func(c *gin.Context) {
			item_ := &[]models.Oldapt{}
			c.ShouldBindJSON(item_)
			var controller rest.OldaptController
			controller.Init(c)
			controller.Deletebatch(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/oldapt/count", func(c *gin.Context) {
			var controller rest.OldaptController
			controller.Init(c)
			controller.Count()
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/oldapt/:id", func(c *gin.Context) {
			id_, _ := strconv.ParseInt(c.Param("id"), 10, 64)
			var controller rest.OldaptController
			controller.Init(c)
			controller.Read(id_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/oldapt", func(c *gin.Context) {
			page_, _ := strconv.Atoi(c.Query("page"))
			pagesize_, _ := strconv.Atoi(c.Query("pagesize"))
			var controller rest.OldaptController
			controller.Init(c)
			controller.Index(page_, pagesize_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.POST("/olddata", func(c *gin.Context) {
			item_ := &models.Olddata{}
			c.ShouldBindJSON(item_)
			var controller rest.OlddataController
			controller.Init(c)
			controller.Insert(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.POST("/olddata/batch", func(c *gin.Context) {
			item_ := &[]models.Olddata{}
			c.ShouldBindJSON(item_)
			var controller rest.OlddataController
			controller.Init(c)
			controller.Insertbatch(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.PUT("/olddata", func(c *gin.Context) {
			item_ := &models.Olddata{}
			c.ShouldBindJSON(item_)
			var controller rest.OlddataController
			controller.Init(c)
			controller.Update(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.DELETE("/olddata", func(c *gin.Context) {
			item_ := &models.Olddata{}
			c.ShouldBindJSON(item_)
			var controller rest.OlddataController
			controller.Init(c)
			controller.Delete(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.DELETE("/olddata/batch", func(c *gin.Context) {
			item_ := &[]models.Olddata{}
			c.ShouldBindJSON(item_)
			var controller rest.OlddataController
			controller.Init(c)
			controller.Deletebatch(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/olddata/count", func(c *gin.Context) {
			var controller rest.OlddataController
			controller.Init(c)
			controller.Count()
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/olddata/:id", func(c *gin.Context) {
			id_, _ := strconv.ParseInt(c.Param("id"), 10, 64)
			var controller rest.OlddataController
			controller.Init(c)
			controller.Read(id_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/olddata", func(c *gin.Context) {
			page_, _ := strconv.Atoi(c.Query("page"))
			pagesize_, _ := strconv.Atoi(c.Query("pagesize"))
			var controller rest.OlddataController
			controller.Init(c)
			controller.Index(page_, pagesize_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/olddata/sum", func(c *gin.Context) {
			var controller rest.OlddataController
			controller.Init(c)
			controller.Sum()
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.POST("/oldimage", func(c *gin.Context) {
			item_ := &models.Oldimage{}
			c.ShouldBindJSON(item_)
			var controller rest.OldimageController
			controller.Init(c)
			controller.Insert(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.POST("/oldimage/batch", func(c *gin.Context) {
			item_ := &[]models.Oldimage{}
			c.ShouldBindJSON(item_)
			var controller rest.OldimageController
			controller.Init(c)
			controller.Insertbatch(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.PUT("/oldimage", func(c *gin.Context) {
			item_ := &models.Oldimage{}
			c.ShouldBindJSON(item_)
			var controller rest.OldimageController
			controller.Init(c)
			controller.Update(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.DELETE("/oldimage", func(c *gin.Context) {
			item_ := &models.Oldimage{}
			c.ShouldBindJSON(item_)
			var controller rest.OldimageController
			controller.Init(c)
			controller.Delete(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.DELETE("/oldimage/batch", func(c *gin.Context) {
			item_ := &[]models.Oldimage{}
			c.ShouldBindJSON(item_)
			var controller rest.OldimageController
			controller.Init(c)
			controller.Deletebatch(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/oldimage/count", func(c *gin.Context) {
			var controller rest.OldimageController
			controller.Init(c)
			controller.Count()
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/oldimage/:id", func(c *gin.Context) {
			id_, _ := strconv.ParseInt(c.Param("id"), 10, 64)
			var controller rest.OldimageController
			controller.Init(c)
			controller.Read(id_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/oldimage", func(c *gin.Context) {
			page_, _ := strconv.Atoi(c.Query("page"))
			pagesize_, _ := strconv.Atoi(c.Query("pagesize"))
			var controller rest.OldimageController
			controller.Init(c)
			controller.Index(page_, pagesize_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.POST("/outline", func(c *gin.Context) {
			item_ := &models.Outline{}
			c.ShouldBindJSON(item_)
			var controller rest.OutlineController
			controller.Init(c)
			controller.Insert(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.POST("/outline/batch", func(c *gin.Context) {
			item_ := &[]models.Outline{}
			c.ShouldBindJSON(item_)
			var controller rest.OutlineController
			controller.Init(c)
			controller.Insertbatch(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.PUT("/outline", func(c *gin.Context) {
			item_ := &models.Outline{}
			c.ShouldBindJSON(item_)
			var controller rest.OutlineController
			controller.Init(c)
			controller.Update(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.DELETE("/outline", func(c *gin.Context) {
			item_ := &models.Outline{}
			c.ShouldBindJSON(item_)
			var controller rest.OutlineController
			controller.Init(c)
			controller.Delete(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.DELETE("/outline/batch", func(c *gin.Context) {
			item_ := &[]models.Outline{}
			c.ShouldBindJSON(item_)
			var controller rest.OutlineController
			controller.Init(c)
			controller.Deletebatch(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/outline/count", func(c *gin.Context) {
			var controller rest.OutlineController
			controller.Init(c)
			controller.Count()
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/outline/:id", func(c *gin.Context) {
			id_, _ := strconv.ParseInt(c.Param("id"), 10, 64)
			var controller rest.OutlineController
			controller.Init(c)
			controller.Read(id_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/outline", func(c *gin.Context) {
			page_, _ := strconv.Atoi(c.Query("page"))
			pagesize_, _ := strconv.Atoi(c.Query("pagesize"))
			var controller rest.OutlineController
			controller.Init(c)
			controller.Index(page_, pagesize_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/outline/sum", func(c *gin.Context) {
			var controller rest.OutlineController
			controller.Init(c)
			controller.Sum()
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.POST("/outlineplan", func(c *gin.Context) {
			item_ := &models.Outlineplan{}
			c.ShouldBindJSON(item_)
			var controller rest.OutlineplanController
			controller.Init(c)
			controller.Insert(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.POST("/outlineplan/batch", func(c *gin.Context) {
			item_ := &[]models.Outlineplan{}
			c.ShouldBindJSON(item_)
			var controller rest.OutlineplanController
			controller.Init(c)
			controller.Insertbatch(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.PUT("/outlineplan", func(c *gin.Context) {
			item_ := &models.Outlineplan{}
			c.ShouldBindJSON(item_)
			var controller rest.OutlineplanController
			controller.Init(c)
			controller.Update(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.DELETE("/outlineplan", func(c *gin.Context) {
			item_ := &models.Outlineplan{}
			c.ShouldBindJSON(item_)
			var controller rest.OutlineplanController
			controller.Init(c)
			controller.Delete(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.DELETE("/outlineplan/batch", func(c *gin.Context) {
			item_ := &[]models.Outlineplan{}
			c.ShouldBindJSON(item_)
			var controller rest.OutlineplanController
			controller.Init(c)
			controller.Deletebatch(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/outlineplan/count", func(c *gin.Context) {
			var controller rest.OutlineplanController
			controller.Init(c)
			controller.Count()
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/outlineplan/:id", func(c *gin.Context) {
			id_, _ := strconv.ParseInt(c.Param("id"), 10, 64)
			var controller rest.OutlineplanController
			controller.Init(c)
			controller.Read(id_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/outlineplan", func(c *gin.Context) {
			page_, _ := strconv.Atoi(c.Query("page"))
			pagesize_, _ := strconv.Atoi(c.Query("pagesize"))
			var controller rest.OutlineplanController
			controller.Init(c)
			controller.Index(page_, pagesize_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/outlineplan/sum", func(c *gin.Context) {
			var controller rest.OutlineplanController
			controller.Init(c)
			controller.Sum()
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/patrol/count/apt/:apt", func(c *gin.Context) {
			apt_, _ := strconv.ParseInt(c.Param("apt"), 10, 64)
			var controller rest.PatrolController
			controller.Init(c)
			controller.CountByApt(apt_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/patrol/find/apt/:apt", func(c *gin.Context) {
			apt_, _ := strconv.ParseInt(c.Param("apt"), 10, 64)
			var controller rest.PatrolController
			controller.Init(c)
			controller.FindByApt(apt_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.POST("/patrol", func(c *gin.Context) {
			item_ := &models.Patrol{}
			c.ShouldBindJSON(item_)
			var controller rest.PatrolController
			controller.Init(c)
			controller.Insert(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.POST("/patrol/batch", func(c *gin.Context) {
			item_ := &[]models.Patrol{}
			c.ShouldBindJSON(item_)
			var controller rest.PatrolController
			controller.Init(c)
			controller.Insertbatch(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.PUT("/patrol", func(c *gin.Context) {
			item_ := &models.Patrol{}
			c.ShouldBindJSON(item_)
			var controller rest.PatrolController
			controller.Init(c)
			controller.Update(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.DELETE("/patrol", func(c *gin.Context) {
			item_ := &models.Patrol{}
			c.ShouldBindJSON(item_)
			var controller rest.PatrolController
			controller.Init(c)
			controller.Delete(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.DELETE("/patrol/batch", func(c *gin.Context) {
			item_ := &[]models.Patrol{}
			c.ShouldBindJSON(item_)
			var controller rest.PatrolController
			controller.Init(c)
			controller.Deletebatch(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/patrol/count", func(c *gin.Context) {
			var controller rest.PatrolController
			controller.Init(c)
			controller.Count()
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/patrol/:id", func(c *gin.Context) {
			id_, _ := strconv.ParseInt(c.Param("id"), 10, 64)
			var controller rest.PatrolController
			controller.Init(c)
			controller.Read(id_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/patrol", func(c *gin.Context) {
			page_, _ := strconv.Atoi(c.Query("page"))
			pagesize_, _ := strconv.Atoi(c.Query("pagesize"))
			var controller rest.PatrolController
			controller.Init(c)
			controller.Index(page_, pagesize_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/patrolimage/count/apt/:apt", func(c *gin.Context) {
			apt_, _ := strconv.ParseInt(c.Param("apt"), 10, 64)
			var controller rest.PatrolimageController
			controller.Init(c)
			controller.CountByApt(apt_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/patrolimage/find/apt/:apt", func(c *gin.Context) {
			apt_, _ := strconv.ParseInt(c.Param("apt"), 10, 64)
			var controller rest.PatrolimageController
			controller.Init(c)
			controller.FindByApt(apt_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/patrolimage/count/patrol/:patrol", func(c *gin.Context) {
			patrol_, _ := strconv.ParseInt(c.Param("patrol"), 10, 64)
			var controller rest.PatrolimageController
			controller.Init(c)
			controller.CountByPatrol(patrol_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/patrolimage/find/patrol/:patrol", func(c *gin.Context) {
			patrol_, _ := strconv.ParseInt(c.Param("patrol"), 10, 64)
			var controller rest.PatrolimageController
			controller.Init(c)
			controller.FindByPatrol(patrol_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.POST("/patrolimage", func(c *gin.Context) {
			item_ := &models.Patrolimage{}
			c.ShouldBindJSON(item_)
			var controller rest.PatrolimageController
			controller.Init(c)
			controller.Insert(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.POST("/patrolimage/batch", func(c *gin.Context) {
			item_ := &[]models.Patrolimage{}
			c.ShouldBindJSON(item_)
			var controller rest.PatrolimageController
			controller.Init(c)
			controller.Insertbatch(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.PUT("/patrolimage", func(c *gin.Context) {
			item_ := &models.Patrolimage{}
			c.ShouldBindJSON(item_)
			var controller rest.PatrolimageController
			controller.Init(c)
			controller.Update(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.DELETE("/patrolimage", func(c *gin.Context) {
			item_ := &models.Patrolimage{}
			c.ShouldBindJSON(item_)
			var controller rest.PatrolimageController
			controller.Init(c)
			controller.Delete(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.DELETE("/patrolimage/batch", func(c *gin.Context) {
			item_ := &[]models.Patrolimage{}
			c.ShouldBindJSON(item_)
			var controller rest.PatrolimageController
			controller.Init(c)
			controller.Deletebatch(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/patrolimage/count", func(c *gin.Context) {
			var controller rest.PatrolimageController
			controller.Init(c)
			controller.Count()
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/patrolimage/:id", func(c *gin.Context) {
			id_, _ := strconv.ParseInt(c.Param("id"), 10, 64)
			var controller rest.PatrolimageController
			controller.Init(c)
			controller.Read(id_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/patrolimage", func(c *gin.Context) {
			page_, _ := strconv.Atoi(c.Query("page"))
			pagesize_, _ := strconv.Atoi(c.Query("pagesize"))
			var controller rest.PatrolimageController
			controller.Init(c)
			controller.Index(page_, pagesize_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.POST("/patrolimages", func(c *gin.Context) {
			item_ := &models.Patrolimages{}
			c.ShouldBindJSON(item_)
			var controller rest.PatrolimagesController
			controller.Init(c)
			controller.Insert(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.POST("/patrolimages/batch", func(c *gin.Context) {
			item_ := &[]models.Patrolimages{}
			c.ShouldBindJSON(item_)
			var controller rest.PatrolimagesController
			controller.Init(c)
			controller.Insertbatch(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.PUT("/patrolimages", func(c *gin.Context) {
			item_ := &models.Patrolimages{}
			c.ShouldBindJSON(item_)
			var controller rest.PatrolimagesController
			controller.Init(c)
			controller.Update(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.DELETE("/patrolimages", func(c *gin.Context) {
			item_ := &models.Patrolimages{}
			c.ShouldBindJSON(item_)
			var controller rest.PatrolimagesController
			controller.Init(c)
			controller.Delete(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.DELETE("/patrolimages/batch", func(c *gin.Context) {
			item_ := &[]models.Patrolimages{}
			c.ShouldBindJSON(item_)
			var controller rest.PatrolimagesController
			controller.Init(c)
			controller.Deletebatch(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/patrolimages/count", func(c *gin.Context) {
			var controller rest.PatrolimagesController
			controller.Init(c)
			controller.Count()
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/patrolimages/:id", func(c *gin.Context) {
			id_, _ := strconv.ParseInt(c.Param("id"), 10, 64)
			var controller rest.PatrolimagesController
			controller.Init(c)
			controller.Read(id_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/patrolimages", func(c *gin.Context) {
			page_, _ := strconv.Atoi(c.Query("page"))
			pagesize_, _ := strconv.Atoi(c.Query("pagesize"))
			var controller rest.PatrolimagesController
			controller.Init(c)
			controller.Index(page_, pagesize_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.POST("/periodic", func(c *gin.Context) {
			item_ := &models.Periodic{}
			c.ShouldBindJSON(item_)
			var apicontroller api.PeriodicController
			apicontroller.Init(c)
			var controller rest.PeriodicController
			controller.Init(c)
			apicontroller.Pre_Insert(item_)
			controller.Insert(item_)
			apicontroller.Post_Insert(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.POST("/periodic/batch", func(c *gin.Context) {
			item_ := &[]models.Periodic{}
			c.ShouldBindJSON(item_)
			var controller rest.PeriodicController
			controller.Init(c)
			controller.Insertbatch(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.PUT("/periodic", func(c *gin.Context) {
			item_ := &models.Periodic{}
			c.ShouldBindJSON(item_)
			var controller rest.PeriodicController
			controller.Init(c)
			controller.Update(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.DELETE("/periodic", func(c *gin.Context) {
			item_ := &models.Periodic{}
			c.ShouldBindJSON(item_)
			var controller rest.PeriodicController
			controller.Init(c)
			controller.Delete(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.DELETE("/periodic/batch", func(c *gin.Context) {
			item_ := &[]models.Periodic{}
			c.ShouldBindJSON(item_)
			var controller rest.PeriodicController
			controller.Init(c)
			controller.Deletebatch(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/periodic/count", func(c *gin.Context) {
			var controller rest.PeriodicController
			controller.Init(c)
			controller.Count()
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/periodic/:id", func(c *gin.Context) {
			id_, _ := strconv.ParseInt(c.Param("id"), 10, 64)
			var controller rest.PeriodicController
			controller.Init(c)
			controller.Read(id_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/periodic", func(c *gin.Context) {
			page_, _ := strconv.Atoi(c.Query("page"))
			pagesize_, _ := strconv.Atoi(c.Query("pagesize"))
			var controller rest.PeriodicController
			controller.Init(c)
			controller.Index(page_, pagesize_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/periodic/sum", func(c *gin.Context) {
			var controller rest.PeriodicController
			controller.Init(c)
			controller.Sum()
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/periodicblueprintzoom/get/periodicblueprint/:periodic", func(c *gin.Context) {
			periodic_, _ := strconv.ParseInt(c.Param("periodic"), 10, 64)
			blueprint_, _ := strconv.ParseInt(c.Query("blueprint"), 10, 64)
			var controller rest.PeriodicblueprintzoomController
			controller.Init(c)
			controller.GetByPeriodicBlueprint(periodic_, blueprint_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.PUT("/periodicblueprintzoom/statusbyperiodicblueprint", func(c *gin.Context) {
			var results map[string]any
			jsonData, _ := c.GetRawData()
			json.Unmarshal(jsonData, &results)
			var status_ int
			if v, flag := results["status"]; flag {
				status_ = int(v.(float64))
			}
			var periodic_ int64
			if v, flag := results["periodic"]; flag {
				periodic_ = int64(v.(float64))
			}
			var blueprint_ int64
			if v, flag := results["blueprint"]; flag {
				blueprint_ = int64(v.(float64))
			}
			var controller rest.PeriodicblueprintzoomController
			controller.Init(c)
			controller.UpdateStatusByPeriodicBlueprint(status_, periodic_, blueprint_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.DELETE("/periodicblueprintzoom/byperiodicblueprint", func(c *gin.Context) {
			item_ := &models.Periodicblueprintzoom{}
			c.ShouldBindJSON(item_)
			var controller rest.PeriodicblueprintzoomController
			controller.Init(c)
			controller.DeleteByPeriodicBlueprint(item_.Periodic, item_.Blueprint)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.POST("/periodicblueprintzoom", func(c *gin.Context) {
			item_ := &models.Periodicblueprintzoom{}
			c.ShouldBindJSON(item_)
			var controller rest.PeriodicblueprintzoomController
			controller.Init(c)
			controller.Insert(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.POST("/periodicblueprintzoom/batch", func(c *gin.Context) {
			item_ := &[]models.Periodicblueprintzoom{}
			c.ShouldBindJSON(item_)
			var controller rest.PeriodicblueprintzoomController
			controller.Init(c)
			controller.Insertbatch(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.PUT("/periodicblueprintzoom", func(c *gin.Context) {
			item_ := &models.Periodicblueprintzoom{}
			c.ShouldBindJSON(item_)
			var controller rest.PeriodicblueprintzoomController
			controller.Init(c)
			controller.Update(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.DELETE("/periodicblueprintzoom", func(c *gin.Context) {
			item_ := &models.Periodicblueprintzoom{}
			c.ShouldBindJSON(item_)
			var controller rest.PeriodicblueprintzoomController
			controller.Init(c)
			controller.Delete(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.DELETE("/periodicblueprintzoom/batch", func(c *gin.Context) {
			item_ := &[]models.Periodicblueprintzoom{}
			c.ShouldBindJSON(item_)
			var controller rest.PeriodicblueprintzoomController
			controller.Init(c)
			controller.Deletebatch(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/periodicblueprintzoom/count", func(c *gin.Context) {
			var controller rest.PeriodicblueprintzoomController
			controller.Init(c)
			controller.Count()
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/periodicblueprintzoom/:id", func(c *gin.Context) {
			id_, _ := strconv.ParseInt(c.Param("id"), 10, 64)
			var controller rest.PeriodicblueprintzoomController
			controller.Init(c)
			controller.Read(id_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/periodicblueprintzoom", func(c *gin.Context) {
			page_, _ := strconv.Atoi(c.Query("page"))
			pagesize_, _ := strconv.Atoi(c.Query("pagesize"))
			var controller rest.PeriodicblueprintzoomController
			controller.Init(c)
			controller.Index(page_, pagesize_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.POST("/periodicchange", func(c *gin.Context) {
			item_ := &models.Periodicchange{}
			c.ShouldBindJSON(item_)
			var controller rest.PeriodicchangeController
			controller.Init(c)
			controller.Insert(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.POST("/periodicchange/batch", func(c *gin.Context) {
			item_ := &[]models.Periodicchange{}
			c.ShouldBindJSON(item_)
			var controller rest.PeriodicchangeController
			controller.Init(c)
			controller.Insertbatch(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.PUT("/periodicchange", func(c *gin.Context) {
			item_ := &models.Periodicchange{}
			c.ShouldBindJSON(item_)
			var controller rest.PeriodicchangeController
			controller.Init(c)
			controller.Update(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.DELETE("/periodicchange", func(c *gin.Context) {
			item_ := &models.Periodicchange{}
			c.ShouldBindJSON(item_)
			var controller rest.PeriodicchangeController
			controller.Init(c)
			controller.Delete(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.DELETE("/periodicchange/batch", func(c *gin.Context) {
			item_ := &[]models.Periodicchange{}
			c.ShouldBindJSON(item_)
			var controller rest.PeriodicchangeController
			controller.Init(c)
			controller.Deletebatch(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/periodicchange/count", func(c *gin.Context) {
			var controller rest.PeriodicchangeController
			controller.Init(c)
			controller.Count()
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/periodicchange/:id", func(c *gin.Context) {
			id_, _ := strconv.ParseInt(c.Param("id"), 10, 64)
			var controller rest.PeriodicchangeController
			controller.Init(c)
			controller.Read(id_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/periodicchange", func(c *gin.Context) {
			page_, _ := strconv.Atoi(c.Query("page"))
			pagesize_, _ := strconv.Atoi(c.Query("pagesize"))
			var controller rest.PeriodicchangeController
			controller.Init(c)
			controller.Index(page_, pagesize_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.POST("/periodiccheck", func(c *gin.Context) {
			item_ := &models.Periodiccheck{}
			c.ShouldBindJSON(item_)
			var controller rest.PeriodiccheckController
			controller.Init(c)
			controller.Insert(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.POST("/periodiccheck/batch", func(c *gin.Context) {
			item_ := &[]models.Periodiccheck{}
			c.ShouldBindJSON(item_)
			var controller rest.PeriodiccheckController
			controller.Init(c)
			controller.Insertbatch(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.PUT("/periodiccheck", func(c *gin.Context) {
			item_ := &models.Periodiccheck{}
			c.ShouldBindJSON(item_)
			var controller rest.PeriodiccheckController
			controller.Init(c)
			controller.Update(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.DELETE("/periodiccheck", func(c *gin.Context) {
			item_ := &models.Periodiccheck{}
			c.ShouldBindJSON(item_)
			var controller rest.PeriodiccheckController
			controller.Init(c)
			controller.Delete(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.DELETE("/periodiccheck/batch", func(c *gin.Context) {
			item_ := &[]models.Periodiccheck{}
			c.ShouldBindJSON(item_)
			var controller rest.PeriodiccheckController
			controller.Init(c)
			controller.Deletebatch(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/periodiccheck/count", func(c *gin.Context) {
			var controller rest.PeriodiccheckController
			controller.Init(c)
			controller.Count()
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/periodiccheck/:id", func(c *gin.Context) {
			id_, _ := strconv.ParseInt(c.Param("id"), 10, 64)
			var controller rest.PeriodiccheckController
			controller.Init(c)
			controller.Read(id_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/periodiccheck", func(c *gin.Context) {
			page_, _ := strconv.Atoi(c.Query("page"))
			pagesize_, _ := strconv.Atoi(c.Query("pagesize"))
			var controller rest.PeriodiccheckController
			controller.Init(c)
			controller.Index(page_, pagesize_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.DELETE("/periodicdata/byperiodicblueprint", func(c *gin.Context) {
			item_ := &models.Periodicdata{}
			c.ShouldBindJSON(item_)
			var controller rest.PeriodicdataController
			controller.Init(c)
			controller.DeleteByPeriodicBlueprint(item_.Periodic, item_.Blueprint)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.POST("/periodicdata", func(c *gin.Context) {
			item_ := &models.Periodicdata{}
			c.ShouldBindJSON(item_)
			var controller rest.PeriodicdataController
			controller.Init(c)
			controller.Insert(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.POST("/periodicdata/batch", func(c *gin.Context) {
			item_ := &[]models.Periodicdata{}
			c.ShouldBindJSON(item_)
			var controller rest.PeriodicdataController
			controller.Init(c)
			controller.Insertbatch(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.PUT("/periodicdata", func(c *gin.Context) {
			item_ := &models.Periodicdata{}
			c.ShouldBindJSON(item_)
			var controller rest.PeriodicdataController
			controller.Init(c)
			controller.Update(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.DELETE("/periodicdata", func(c *gin.Context) {
			item_ := &models.Periodicdata{}
			c.ShouldBindJSON(item_)
			var controller rest.PeriodicdataController
			controller.Init(c)
			controller.Delete(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.DELETE("/periodicdata/batch", func(c *gin.Context) {
			item_ := &[]models.Periodicdata{}
			c.ShouldBindJSON(item_)
			var controller rest.PeriodicdataController
			controller.Init(c)
			controller.Deletebatch(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/periodicdata/count", func(c *gin.Context) {
			var controller rest.PeriodicdataController
			controller.Init(c)
			controller.Count()
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/periodicdata/:id", func(c *gin.Context) {
			id_, _ := strconv.ParseInt(c.Param("id"), 10, 64)
			var controller rest.PeriodicdataController
			controller.Init(c)
			controller.Read(id_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/periodicdata", func(c *gin.Context) {
			page_, _ := strconv.Atoi(c.Query("page"))
			pagesize_, _ := strconv.Atoi(c.Query("pagesize"))
			var apicontroller api.PeriodicdataController
			apicontroller.Init(c)
			var controller rest.PeriodicdataController
			controller.Init(c)
			controller.Index(page_, pagesize_)
			readItem := controller.Result["items"].([]models.Periodicdata)
			apicontroller.Post_Index(readItem)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/periodicdata/sum", func(c *gin.Context) {
			var controller rest.PeriodicdataController
			controller.Init(c)
			controller.Sum()
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/periodicdatabackup/count", func(c *gin.Context) {
			var controller rest.PeriodicdatabackupController
			controller.Init(c)
			controller.Count()
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/periodicdatabackup/:id", func(c *gin.Context) {
			id_, _ := strconv.ParseInt(c.Param("id"), 10, 64)
			var controller rest.PeriodicdatabackupController
			controller.Init(c)
			controller.Read(id_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/periodicdatabackup", func(c *gin.Context) {
			page_, _ := strconv.Atoi(c.Query("page"))
			pagesize_, _ := strconv.Atoi(c.Query("pagesize"))
			var controller rest.PeriodicdatabackupController
			controller.Init(c)
			controller.Index(page_, pagesize_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/periodicdatabackup/sum", func(c *gin.Context) {
			var controller rest.PeriodicdatabackupController
			controller.Init(c)
			controller.Sum()
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/periodicdataimage/get/periodic/:periodic", func(c *gin.Context) {
			periodic_, _ := strconv.ParseInt(c.Param("periodic"), 10, 64)
			var controller rest.PeriodicdataimageController
			controller.Init(c)
			controller.GetByPeriodic(periodic_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/periodicdataimage/count/offlinefilename/:offlinefilename", func(c *gin.Context) {
			offlinefilename_ := c.Param("offlinefilename")
			var controller rest.PeriodicdataimageController
			controller.Init(c)
			controller.CountByOfflinefilename(offlinefilename_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.DELETE("/periodicdataimage/byperiodicdata", func(c *gin.Context) {
			item_ := &models.Periodicdataimage{}
			c.ShouldBindJSON(item_)
			var controller rest.PeriodicdataimageController
			controller.Init(c)
			controller.DeleteByPeriodicdata(item_.Periodicdata)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.POST("/periodicdataimage", func(c *gin.Context) {
			item_ := &models.Periodicdataimage{}
			c.ShouldBindJSON(item_)
			var controller rest.PeriodicdataimageController
			controller.Init(c)
			controller.Insert(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.POST("/periodicdataimage/batch", func(c *gin.Context) {
			item_ := &[]models.Periodicdataimage{}
			c.ShouldBindJSON(item_)
			var controller rest.PeriodicdataimageController
			controller.Init(c)
			controller.Insertbatch(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.PUT("/periodicdataimage", func(c *gin.Context) {
			item_ := &models.Periodicdataimage{}
			c.ShouldBindJSON(item_)
			var controller rest.PeriodicdataimageController
			controller.Init(c)
			controller.Update(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.DELETE("/periodicdataimage", func(c *gin.Context) {
			item_ := &models.Periodicdataimage{}
			c.ShouldBindJSON(item_)
			var controller rest.PeriodicdataimageController
			controller.Init(c)
			controller.Delete(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.DELETE("/periodicdataimage/batch", func(c *gin.Context) {
			item_ := &[]models.Periodicdataimage{}
			c.ShouldBindJSON(item_)
			var controller rest.PeriodicdataimageController
			controller.Init(c)
			controller.Deletebatch(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/periodicdataimage/count", func(c *gin.Context) {
			var controller rest.PeriodicdataimageController
			controller.Init(c)
			controller.Count()
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/periodicdataimage/:id", func(c *gin.Context) {
			id_, _ := strconv.ParseInt(c.Param("id"), 10, 64)
			var controller rest.PeriodicdataimageController
			controller.Init(c)
			controller.Read(id_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/periodicdataimage", func(c *gin.Context) {
			page_, _ := strconv.Atoi(c.Query("page"))
			pagesize_, _ := strconv.Atoi(c.Query("pagesize"))
			var controller rest.PeriodicdataimageController
			controller.Init(c)
			controller.Index(page_, pagesize_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/periodicimage/get/periodic/:periodic", func(c *gin.Context) {
			periodic_, _ := strconv.ParseInt(c.Param("periodic"), 10, 64)
			var controller rest.PeriodicimageController
			controller.Init(c)
			controller.GetByPeriodic(periodic_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/periodicimage/count/offlinefilename/:offlinefilename", func(c *gin.Context) {
			offlinefilename_ := c.Param("offlinefilename")
			var controller rest.PeriodicimageController
			controller.Init(c)
			controller.CountByOfflinefilename(offlinefilename_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.POST("/periodicimage", func(c *gin.Context) {
			item_ := &models.Periodicimage{}
			c.ShouldBindJSON(item_)
			var controller rest.PeriodicimageController
			controller.Init(c)
			controller.Insert(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.POST("/periodicimage/batch", func(c *gin.Context) {
			item_ := &[]models.Periodicimage{}
			c.ShouldBindJSON(item_)
			var controller rest.PeriodicimageController
			controller.Init(c)
			controller.Insertbatch(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.PUT("/periodicimage", func(c *gin.Context) {
			item_ := &models.Periodicimage{}
			c.ShouldBindJSON(item_)
			var controller rest.PeriodicimageController
			controller.Init(c)
			controller.Update(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.DELETE("/periodicimage", func(c *gin.Context) {
			item_ := &models.Periodicimage{}
			c.ShouldBindJSON(item_)
			var apicontroller api.PeriodicimageController
			apicontroller.Init(c)
			var controller rest.PeriodicimageController
			controller.Init(c)
			apicontroller.Pre_Delete(item_)
			controller.Delete(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.DELETE("/periodicimage/batch", func(c *gin.Context) {
			item_ := &[]models.Periodicimage{}
			c.ShouldBindJSON(item_)
			var apicontroller api.PeriodicimageController
			apicontroller.Init(c)
			var controller rest.PeriodicimageController
			controller.Init(c)
			apicontroller.Pre_Deletebatch(item_)
			controller.Deletebatch(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/periodicimage/count", func(c *gin.Context) {
			var controller rest.PeriodicimageController
			controller.Init(c)
			controller.Count()
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/periodicimage/:id", func(c *gin.Context) {
			id_, _ := strconv.ParseInt(c.Param("id"), 10, 64)
			var controller rest.PeriodicimageController
			controller.Init(c)
			controller.Read(id_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/periodicimage", func(c *gin.Context) {
			page_, _ := strconv.Atoi(c.Query("page"))
			pagesize_, _ := strconv.Atoi(c.Query("pagesize"))
			var controller rest.PeriodicimageController
			controller.Init(c)
			controller.Index(page_, pagesize_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/periodicincidental/get/periodic/:periodic", func(c *gin.Context) {
			periodic_, _ := strconv.ParseInt(c.Param("periodic"), 10, 64)
			var controller rest.PeriodicincidentalController
			controller.Init(c)
			controller.GetByPeriodic(periodic_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.POST("/periodicincidental", func(c *gin.Context) {
			item_ := &models.Periodicincidental{}
			c.ShouldBindJSON(item_)
			var controller rest.PeriodicincidentalController
			controller.Init(c)
			controller.Insert(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.POST("/periodicincidental/batch", func(c *gin.Context) {
			item_ := &[]models.Periodicincidental{}
			c.ShouldBindJSON(item_)
			var controller rest.PeriodicincidentalController
			controller.Init(c)
			controller.Insertbatch(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.PUT("/periodicincidental", func(c *gin.Context) {
			item_ := &models.Periodicincidental{}
			c.ShouldBindJSON(item_)
			var controller rest.PeriodicincidentalController
			controller.Init(c)
			controller.Update(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.DELETE("/periodicincidental", func(c *gin.Context) {
			item_ := &models.Periodicincidental{}
			c.ShouldBindJSON(item_)
			var controller rest.PeriodicincidentalController
			controller.Init(c)
			controller.Delete(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.DELETE("/periodicincidental/batch", func(c *gin.Context) {
			item_ := &[]models.Periodicincidental{}
			c.ShouldBindJSON(item_)
			var controller rest.PeriodicincidentalController
			controller.Init(c)
			controller.Deletebatch(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/periodicincidental/count", func(c *gin.Context) {
			var controller rest.PeriodicincidentalController
			controller.Init(c)
			controller.Count()
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/periodicincidental/:id", func(c *gin.Context) {
			id_, _ := strconv.ParseInt(c.Param("id"), 10, 64)
			var controller rest.PeriodicincidentalController
			controller.Init(c)
			controller.Read(id_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/periodicincidental", func(c *gin.Context) {
			page_, _ := strconv.Atoi(c.Query("page"))
			pagesize_, _ := strconv.Atoi(c.Query("pagesize"))
			var controller rest.PeriodicincidentalController
			controller.Init(c)
			controller.Index(page_, pagesize_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/periodickeep/get/periodic/:periodic", func(c *gin.Context) {
			periodic_, _ := strconv.ParseInt(c.Param("periodic"), 10, 64)
			var controller rest.PeriodickeepController
			controller.Init(c)
			controller.GetByPeriodic(periodic_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.POST("/periodickeep", func(c *gin.Context) {
			item_ := &models.Periodickeep{}
			c.ShouldBindJSON(item_)
			var controller rest.PeriodickeepController
			controller.Init(c)
			controller.Insert(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.POST("/periodickeep/batch", func(c *gin.Context) {
			item_ := &[]models.Periodickeep{}
			c.ShouldBindJSON(item_)
			var controller rest.PeriodickeepController
			controller.Init(c)
			controller.Insertbatch(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.PUT("/periodickeep", func(c *gin.Context) {
			item_ := &models.Periodickeep{}
			c.ShouldBindJSON(item_)
			var controller rest.PeriodickeepController
			controller.Init(c)
			controller.Update(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.DELETE("/periodickeep", func(c *gin.Context) {
			item_ := &models.Periodickeep{}
			c.ShouldBindJSON(item_)
			var controller rest.PeriodickeepController
			controller.Init(c)
			controller.Delete(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.DELETE("/periodickeep/batch", func(c *gin.Context) {
			item_ := &[]models.Periodickeep{}
			c.ShouldBindJSON(item_)
			var controller rest.PeriodickeepController
			controller.Init(c)
			controller.Deletebatch(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/periodickeep/count", func(c *gin.Context) {
			var controller rest.PeriodickeepController
			controller.Init(c)
			controller.Count()
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/periodickeep/:id", func(c *gin.Context) {
			id_, _ := strconv.ParseInt(c.Param("id"), 10, 64)
			var controller rest.PeriodickeepController
			controller.Init(c)
			controller.Read(id_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/periodickeep", func(c *gin.Context) {
			page_, _ := strconv.Atoi(c.Query("page"))
			pagesize_, _ := strconv.Atoi(c.Query("pagesize"))
			var controller rest.PeriodickeepController
			controller.Init(c)
			controller.Index(page_, pagesize_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/periodicopinion/get/periodic/:periodic", func(c *gin.Context) {
			periodic_, _ := strconv.ParseInt(c.Param("periodic"), 10, 64)
			var controller rest.PeriodicopinionController
			controller.Init(c)
			controller.GetByPeriodic(periodic_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.POST("/periodicopinion", func(c *gin.Context) {
			item_ := &models.Periodicopinion{}
			c.ShouldBindJSON(item_)
			var controller rest.PeriodicopinionController
			controller.Init(c)
			controller.Insert(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.POST("/periodicopinion/batch", func(c *gin.Context) {
			item_ := &[]models.Periodicopinion{}
			c.ShouldBindJSON(item_)
			var controller rest.PeriodicopinionController
			controller.Init(c)
			controller.Insertbatch(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.PUT("/periodicopinion", func(c *gin.Context) {
			item_ := &models.Periodicopinion{}
			c.ShouldBindJSON(item_)
			var controller rest.PeriodicopinionController
			controller.Init(c)
			controller.Update(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.DELETE("/periodicopinion", func(c *gin.Context) {
			item_ := &models.Periodicopinion{}
			c.ShouldBindJSON(item_)
			var controller rest.PeriodicopinionController
			controller.Init(c)
			controller.Delete(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.DELETE("/periodicopinion/batch", func(c *gin.Context) {
			item_ := &[]models.Periodicopinion{}
			c.ShouldBindJSON(item_)
			var controller rest.PeriodicopinionController
			controller.Init(c)
			controller.Deletebatch(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/periodicopinion/count", func(c *gin.Context) {
			var controller rest.PeriodicopinionController
			controller.Init(c)
			controller.Count()
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/periodicopinion/:id", func(c *gin.Context) {
			id_, _ := strconv.ParseInt(c.Param("id"), 10, 64)
			var controller rest.PeriodicopinionController
			controller.Init(c)
			controller.Read(id_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/periodicopinion", func(c *gin.Context) {
			page_, _ := strconv.Atoi(c.Query("page"))
			pagesize_, _ := strconv.Atoi(c.Query("pagesize"))
			var controller rest.PeriodicopinionController
			controller.Init(c)
			controller.Index(page_, pagesize_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.POST("/periodicother", func(c *gin.Context) {
			item_ := &models.Periodicother{}
			c.ShouldBindJSON(item_)
			var controller rest.PeriodicotherController
			controller.Init(c)
			controller.Insert(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.POST("/periodicother/batch", func(c *gin.Context) {
			item_ := &[]models.Periodicother{}
			c.ShouldBindJSON(item_)
			var controller rest.PeriodicotherController
			controller.Init(c)
			controller.Insertbatch(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.PUT("/periodicother", func(c *gin.Context) {
			item_ := &models.Periodicother{}
			c.ShouldBindJSON(item_)
			var controller rest.PeriodicotherController
			controller.Init(c)
			controller.Update(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.DELETE("/periodicother", func(c *gin.Context) {
			item_ := &models.Periodicother{}
			c.ShouldBindJSON(item_)
			var controller rest.PeriodicotherController
			controller.Init(c)
			controller.Delete(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.DELETE("/periodicother/batch", func(c *gin.Context) {
			item_ := &[]models.Periodicother{}
			c.ShouldBindJSON(item_)
			var controller rest.PeriodicotherController
			controller.Init(c)
			controller.Deletebatch(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/periodicother/count", func(c *gin.Context) {
			var controller rest.PeriodicotherController
			controller.Init(c)
			controller.Count()
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/periodicother/:id", func(c *gin.Context) {
			id_, _ := strconv.ParseInt(c.Param("id"), 10, 64)
			var controller rest.PeriodicotherController
			controller.Init(c)
			controller.Read(id_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/periodicother", func(c *gin.Context) {
			page_, _ := strconv.Atoi(c.Query("page"))
			pagesize_, _ := strconv.Atoi(c.Query("pagesize"))
			var controller rest.PeriodicotherController
			controller.Init(c)
			controller.Index(page_, pagesize_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/periodicotheretc/get/periodic/:periodic", func(c *gin.Context) {
			periodic_, _ := strconv.ParseInt(c.Param("periodic"), 10, 64)
			var controller rest.PeriodicotheretcController
			controller.Init(c)
			controller.GetByPeriodic(periodic_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.POST("/periodicotheretc", func(c *gin.Context) {
			item_ := &models.Periodicotheretc{}
			c.ShouldBindJSON(item_)
			var controller rest.PeriodicotheretcController
			controller.Init(c)
			controller.Insert(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.POST("/periodicotheretc/batch", func(c *gin.Context) {
			item_ := &[]models.Periodicotheretc{}
			c.ShouldBindJSON(item_)
			var controller rest.PeriodicotheretcController
			controller.Init(c)
			controller.Insertbatch(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.PUT("/periodicotheretc", func(c *gin.Context) {
			item_ := &models.Periodicotheretc{}
			c.ShouldBindJSON(item_)
			var controller rest.PeriodicotheretcController
			controller.Init(c)
			controller.Update(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.DELETE("/periodicotheretc", func(c *gin.Context) {
			item_ := &models.Periodicotheretc{}
			c.ShouldBindJSON(item_)
			var controller rest.PeriodicotheretcController
			controller.Init(c)
			controller.Delete(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.DELETE("/periodicotheretc/batch", func(c *gin.Context) {
			item_ := &[]models.Periodicotheretc{}
			c.ShouldBindJSON(item_)
			var controller rest.PeriodicotheretcController
			controller.Init(c)
			controller.Deletebatch(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/periodicotheretc/count", func(c *gin.Context) {
			var controller rest.PeriodicotheretcController
			controller.Init(c)
			controller.Count()
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/periodicotheretc/:id", func(c *gin.Context) {
			id_, _ := strconv.ParseInt(c.Param("id"), 10, 64)
			var controller rest.PeriodicotheretcController
			controller.Init(c)
			controller.Read(id_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/periodicotheretc", func(c *gin.Context) {
			page_, _ := strconv.Atoi(c.Query("page"))
			pagesize_, _ := strconv.Atoi(c.Query("pagesize"))
			var controller rest.PeriodicotheretcController
			controller.Init(c)
			controller.Index(page_, pagesize_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/periodicouterwall/get/periodic/:periodic", func(c *gin.Context) {
			periodic_, _ := strconv.ParseInt(c.Param("periodic"), 10, 64)
			var controller rest.PeriodicouterwallController
			controller.Init(c)
			controller.GetByPeriodic(periodic_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.POST("/periodicouterwall", func(c *gin.Context) {
			item_ := &models.Periodicouterwall{}
			c.ShouldBindJSON(item_)
			var controller rest.PeriodicouterwallController
			controller.Init(c)
			controller.Insert(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.POST("/periodicouterwall/batch", func(c *gin.Context) {
			item_ := &[]models.Periodicouterwall{}
			c.ShouldBindJSON(item_)
			var controller rest.PeriodicouterwallController
			controller.Init(c)
			controller.Insertbatch(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.PUT("/periodicouterwall", func(c *gin.Context) {
			item_ := &models.Periodicouterwall{}
			c.ShouldBindJSON(item_)
			var controller rest.PeriodicouterwallController
			controller.Init(c)
			controller.Update(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.DELETE("/periodicouterwall", func(c *gin.Context) {
			item_ := &models.Periodicouterwall{}
			c.ShouldBindJSON(item_)
			var controller rest.PeriodicouterwallController
			controller.Init(c)
			controller.Delete(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.DELETE("/periodicouterwall/batch", func(c *gin.Context) {
			item_ := &[]models.Periodicouterwall{}
			c.ShouldBindJSON(item_)
			var controller rest.PeriodicouterwallController
			controller.Init(c)
			controller.Deletebatch(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/periodicouterwall/count", func(c *gin.Context) {
			var controller rest.PeriodicouterwallController
			controller.Init(c)
			controller.Count()
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/periodicouterwall/:id", func(c *gin.Context) {
			id_, _ := strconv.ParseInt(c.Param("id"), 10, 64)
			var controller rest.PeriodicouterwallController
			controller.Init(c)
			controller.Read(id_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/periodicouterwall", func(c *gin.Context) {
			page_, _ := strconv.Atoi(c.Query("page"))
			pagesize_, _ := strconv.Atoi(c.Query("pagesize"))
			var controller rest.PeriodicouterwallController
			controller.Init(c)
			controller.Index(page_, pagesize_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.POST("/periodicpast", func(c *gin.Context) {
			item_ := &models.Periodicpast{}
			c.ShouldBindJSON(item_)
			var controller rest.PeriodicpastController
			controller.Init(c)
			controller.Insert(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.POST("/periodicpast/batch", func(c *gin.Context) {
			item_ := &[]models.Periodicpast{}
			c.ShouldBindJSON(item_)
			var controller rest.PeriodicpastController
			controller.Init(c)
			controller.Insertbatch(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.PUT("/periodicpast", func(c *gin.Context) {
			item_ := &models.Periodicpast{}
			c.ShouldBindJSON(item_)
			var controller rest.PeriodicpastController
			controller.Init(c)
			controller.Update(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.DELETE("/periodicpast", func(c *gin.Context) {
			item_ := &models.Periodicpast{}
			c.ShouldBindJSON(item_)
			var controller rest.PeriodicpastController
			controller.Init(c)
			controller.Delete(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.DELETE("/periodicpast/batch", func(c *gin.Context) {
			item_ := &[]models.Periodicpast{}
			c.ShouldBindJSON(item_)
			var controller rest.PeriodicpastController
			controller.Init(c)
			controller.Deletebatch(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/periodicpast/count", func(c *gin.Context) {
			var controller rest.PeriodicpastController
			controller.Init(c)
			controller.Count()
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/periodicpast/:id", func(c *gin.Context) {
			id_, _ := strconv.ParseInt(c.Param("id"), 10, 64)
			var controller rest.PeriodicpastController
			controller.Init(c)
			controller.Read(id_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/periodicpast", func(c *gin.Context) {
			page_, _ := strconv.Atoi(c.Query("page"))
			pagesize_, _ := strconv.Atoi(c.Query("pagesize"))
			var controller rest.PeriodicpastController
			controller.Init(c)
			controller.Index(page_, pagesize_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.POST("/periodicresult", func(c *gin.Context) {
			item_ := &models.Periodicresult{}
			c.ShouldBindJSON(item_)
			var controller rest.PeriodicresultController
			controller.Init(c)
			controller.Insert(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.POST("/periodicresult/batch", func(c *gin.Context) {
			item_ := &[]models.Periodicresult{}
			c.ShouldBindJSON(item_)
			var controller rest.PeriodicresultController
			controller.Init(c)
			controller.Insertbatch(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.PUT("/periodicresult", func(c *gin.Context) {
			item_ := &models.Periodicresult{}
			c.ShouldBindJSON(item_)
			var controller rest.PeriodicresultController
			controller.Init(c)
			controller.Update(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.DELETE("/periodicresult", func(c *gin.Context) {
			item_ := &models.Periodicresult{}
			c.ShouldBindJSON(item_)
			var controller rest.PeriodicresultController
			controller.Init(c)
			controller.Delete(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.DELETE("/periodicresult/batch", func(c *gin.Context) {
			item_ := &[]models.Periodicresult{}
			c.ShouldBindJSON(item_)
			var controller rest.PeriodicresultController
			controller.Init(c)
			controller.Deletebatch(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/periodicresult/count", func(c *gin.Context) {
			var controller rest.PeriodicresultController
			controller.Init(c)
			controller.Count()
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/periodicresult/:id", func(c *gin.Context) {
			id_, _ := strconv.ParseInt(c.Param("id"), 10, 64)
			var controller rest.PeriodicresultController
			controller.Init(c)
			controller.Read(id_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/periodicresult", func(c *gin.Context) {
			page_, _ := strconv.Atoi(c.Query("page"))
			pagesize_, _ := strconv.Atoi(c.Query("pagesize"))
			var controller rest.PeriodicresultController
			controller.Init(c)
			controller.Index(page_, pagesize_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/periodictechnician/count/periodic/:periodic", func(c *gin.Context) {
			periodic_, _ := strconv.ParseInt(c.Param("periodic"), 10, 64)
			var controller rest.PeriodictechnicianController
			controller.Init(c)
			controller.CountByPeriodic(periodic_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/periodictechnician/find/periodic/:periodic", func(c *gin.Context) {
			periodic_, _ := strconv.ParseInt(c.Param("periodic"), 10, 64)
			var controller rest.PeriodictechnicianController
			controller.Init(c)
			controller.FindByPeriodic(periodic_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.DELETE("/periodictechnician/byperiodic", func(c *gin.Context) {
			item_ := &models.Periodictechnician{}
			c.ShouldBindJSON(item_)
			var controller rest.PeriodictechnicianController
			controller.Init(c)
			controller.DeleteByPeriodic(item_.Periodic)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.POST("/periodictechnician", func(c *gin.Context) {
			item_ := &models.Periodictechnician{}
			c.ShouldBindJSON(item_)
			var controller rest.PeriodictechnicianController
			controller.Init(c)
			controller.Insert(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.POST("/periodictechnician/batch", func(c *gin.Context) {
			item_ := &[]models.Periodictechnician{}
			c.ShouldBindJSON(item_)
			var controller rest.PeriodictechnicianController
			controller.Init(c)
			controller.Insertbatch(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.PUT("/periodictechnician", func(c *gin.Context) {
			item_ := &models.Periodictechnician{}
			c.ShouldBindJSON(item_)
			var controller rest.PeriodictechnicianController
			controller.Init(c)
			controller.Update(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.DELETE("/periodictechnician", func(c *gin.Context) {
			item_ := &models.Periodictechnician{}
			c.ShouldBindJSON(item_)
			var controller rest.PeriodictechnicianController
			controller.Init(c)
			controller.Delete(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.DELETE("/periodictechnician/batch", func(c *gin.Context) {
			item_ := &[]models.Periodictechnician{}
			c.ShouldBindJSON(item_)
			var controller rest.PeriodictechnicianController
			controller.Init(c)
			controller.Deletebatch(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/periodictechnician/count", func(c *gin.Context) {
			var controller rest.PeriodictechnicianController
			controller.Init(c)
			controller.Count()
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/periodictechnician/:id", func(c *gin.Context) {
			id_, _ := strconv.ParseInt(c.Param("id"), 10, 64)
			var controller rest.PeriodictechnicianController
			controller.Init(c)
			controller.Read(id_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/periodictechnician", func(c *gin.Context) {
			page_, _ := strconv.Atoi(c.Query("page"))
			pagesize_, _ := strconv.Atoi(c.Query("pagesize"))
			var controller rest.PeriodictechnicianController
			controller.Init(c)
			controller.Index(page_, pagesize_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.POST("/program", func(c *gin.Context) {
			item_ := &models.Program{}
			c.ShouldBindJSON(item_)
			var controller rest.ProgramController
			controller.Init(c)
			controller.Insert(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.POST("/program/batch", func(c *gin.Context) {
			item_ := &[]models.Program{}
			c.ShouldBindJSON(item_)
			var controller rest.ProgramController
			controller.Init(c)
			controller.Insertbatch(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.PUT("/program", func(c *gin.Context) {
			item_ := &models.Program{}
			c.ShouldBindJSON(item_)
			var controller rest.ProgramController
			controller.Init(c)
			controller.Update(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.DELETE("/program", func(c *gin.Context) {
			item_ := &models.Program{}
			c.ShouldBindJSON(item_)
			var controller rest.ProgramController
			controller.Init(c)
			controller.Delete(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.DELETE("/program/batch", func(c *gin.Context) {
			item_ := &[]models.Program{}
			c.ShouldBindJSON(item_)
			var controller rest.ProgramController
			controller.Init(c)
			controller.Deletebatch(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/program/count", func(c *gin.Context) {
			var controller rest.ProgramController
			controller.Init(c)
			controller.Count()
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/program/:id", func(c *gin.Context) {
			id_, _ := strconv.ParseInt(c.Param("id"), 10, 64)
			var controller rest.ProgramController
			controller.Init(c)
			controller.Read(id_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/program", func(c *gin.Context) {
			page_, _ := strconv.Atoi(c.Query("page"))
			pagesize_, _ := strconv.Atoi(c.Query("pagesize"))
			var controller rest.ProgramController
			controller.Init(c)
			controller.Index(page_, pagesize_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.PUT("/repair/statusbyid", func(c *gin.Context) {
			var results map[string]any
			jsonData, _ := c.GetRawData()
			json.Unmarshal(jsonData, &results)
			var status_ int
			if v, flag := results["status"]; flag {
				status_ = int(v.(float64))
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.RepairController
			controller.Init(c)
			controller.UpdateStatusById(status_, id_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.POST("/repair", func(c *gin.Context) {
			item_ := &models.Repair{}
			c.ShouldBindJSON(item_)
			var controller rest.RepairController
			controller.Init(c)
			controller.Insert(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.POST("/repair/batch", func(c *gin.Context) {
			item_ := &[]models.Repair{}
			c.ShouldBindJSON(item_)
			var controller rest.RepairController
			controller.Init(c)
			controller.Insertbatch(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.PUT("/repair", func(c *gin.Context) {
			item_ := &models.Repair{}
			c.ShouldBindJSON(item_)
			var apicontroller api.RepairController
			apicontroller.Init(c)
			var controller rest.RepairController
			controller.Init(c)
			apicontroller.Pre_Update(item_)
			controller.Update(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.DELETE("/repair", func(c *gin.Context) {
			item_ := &models.Repair{}
			c.ShouldBindJSON(item_)
			var controller rest.RepairController
			controller.Init(c)
			controller.Delete(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.DELETE("/repair/batch", func(c *gin.Context) {
			item_ := &[]models.Repair{}
			c.ShouldBindJSON(item_)
			var controller rest.RepairController
			controller.Init(c)
			controller.Deletebatch(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/repair/count", func(c *gin.Context) {
			var controller rest.RepairController
			controller.Init(c)
			controller.Count()
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/repair/:id", func(c *gin.Context) {
			id_, _ := strconv.ParseInt(c.Param("id"), 10, 64)
			var controller rest.RepairController
			controller.Init(c)
			controller.Read(id_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/repair", func(c *gin.Context) {
			page_, _ := strconv.Atoi(c.Query("page"))
			pagesize_, _ := strconv.Atoi(c.Query("pagesize"))
			var controller rest.RepairController
			controller.Init(c)
			controller.Index(page_, pagesize_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.POST("/repairarea", func(c *gin.Context) {
			item_ := &models.Repairarea{}
			c.ShouldBindJSON(item_)
			var controller rest.RepairareaController
			controller.Init(c)
			controller.Insert(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.POST("/repairarea/batch", func(c *gin.Context) {
			item_ := &[]models.Repairarea{}
			c.ShouldBindJSON(item_)
			var controller rest.RepairareaController
			controller.Init(c)
			controller.Insertbatch(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.PUT("/repairarea", func(c *gin.Context) {
			item_ := &models.Repairarea{}
			c.ShouldBindJSON(item_)
			var controller rest.RepairareaController
			controller.Init(c)
			controller.Update(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.DELETE("/repairarea", func(c *gin.Context) {
			item_ := &models.Repairarea{}
			c.ShouldBindJSON(item_)
			var controller rest.RepairareaController
			controller.Init(c)
			controller.Delete(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.DELETE("/repairarea/batch", func(c *gin.Context) {
			item_ := &[]models.Repairarea{}
			c.ShouldBindJSON(item_)
			var controller rest.RepairareaController
			controller.Init(c)
			controller.Deletebatch(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/repairarea/count", func(c *gin.Context) {
			var controller rest.RepairareaController
			controller.Init(c)
			controller.Count()
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/repairarea/:id", func(c *gin.Context) {
			id_, _ := strconv.ParseInt(c.Param("id"), 10, 64)
			var controller rest.RepairareaController
			controller.Init(c)
			controller.Read(id_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/repairarea", func(c *gin.Context) {
			page_, _ := strconv.Atoi(c.Query("page"))
			pagesize_, _ := strconv.Atoi(c.Query("pagesize"))
			var controller rest.RepairareaController
			controller.Init(c)
			controller.Index(page_, pagesize_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.PUT("/repairlist/statusbyid", func(c *gin.Context) {
			var results map[string]any
			jsonData, _ := c.GetRawData()
			json.Unmarshal(jsonData, &results)
			var status_ int
			if v, flag := results["status"]; flag {
				status_ = int(v.(float64))
			}
			var id_ int64
			if v, flag := results["id"]; flag {
				id_ = int64(v.(float64))
			}
			var controller rest.RepairlistController
			controller.Init(c)
			controller.UpdateStatusById(status_, id_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/repairlist/count/repairtypes", func(c *gin.Context) {
			repairtype_ := repairlist.ConvertRepairtype(getArrayCommai(c.Query("repairtype")))
			var controller rest.RepairlistController
			controller.Init(c)
			controller.CountByRepairtypes(repairtype_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/repairlist/find/repairtypes", func(c *gin.Context) {
			repairtype_ := repairlist.ConvertRepairtype(getArrayCommai(c.Query("repairtype")))
			var controller rest.RepairlistController
			controller.Init(c)
			controller.FindByRepairtypes(repairtype_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/repairlist/count", func(c *gin.Context) {
			var controller rest.RepairlistController
			controller.Init(c)
			controller.Count()
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/repairlist/:id", func(c *gin.Context) {
			id_, _ := strconv.ParseInt(c.Param("id"), 10, 64)
			var controller rest.RepairlistController
			controller.Init(c)
			controller.Read(id_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/repairlist", func(c *gin.Context) {
			page_, _ := strconv.Atoi(c.Query("page"))
			pagesize_, _ := strconv.Atoi(c.Query("pagesize"))
			var controller rest.RepairlistController
			controller.Init(c)
			controller.Index(page_, pagesize_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/review/count/apt/:apt", func(c *gin.Context) {
			apt_, _ := strconv.ParseInt(c.Param("apt"), 10, 64)
			var controller rest.ReviewController
			controller.Init(c)
			controller.CountByApt(apt_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/review/find/apt/:apt", func(c *gin.Context) {
			apt_, _ := strconv.ParseInt(c.Param("apt"), 10, 64)
			var controller rest.ReviewController
			controller.Init(c)
			controller.FindByApt(apt_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/review/find/category/:category", func(c *gin.Context) {
			category_, _ := strconv.ParseInt(c.Param("category"), 10, 64)
			var controller rest.ReviewController
			controller.Init(c)
			controller.FindByCategory(category_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.DELETE("/review/bycategory", func(c *gin.Context) {
			item_ := &models.Review{}
			c.ShouldBindJSON(item_)
			var controller rest.ReviewController
			controller.Init(c)
			controller.DeleteByCategory(item_.Category)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/review/find/method/:method", func(c *gin.Context) {
			method_, _ := strconv.ParseInt(c.Param("method"), 10, 64)
			var controller rest.ReviewController
			controller.Init(c)
			controller.FindByMethod(method_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.POST("/review", func(c *gin.Context) {
			item_ := &models.Review{}
			c.ShouldBindJSON(item_)
			var controller rest.ReviewController
			controller.Init(c)
			controller.Insert(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.POST("/review/batch", func(c *gin.Context) {
			item_ := &[]models.Review{}
			c.ShouldBindJSON(item_)
			var controller rest.ReviewController
			controller.Init(c)
			controller.Insertbatch(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.PUT("/review", func(c *gin.Context) {
			item_ := &models.Review{}
			c.ShouldBindJSON(item_)
			var controller rest.ReviewController
			controller.Init(c)
			controller.Update(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.DELETE("/review", func(c *gin.Context) {
			item_ := &models.Review{}
			c.ShouldBindJSON(item_)
			var controller rest.ReviewController
			controller.Init(c)
			controller.Delete(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.DELETE("/review/batch", func(c *gin.Context) {
			item_ := &[]models.Review{}
			c.ShouldBindJSON(item_)
			var controller rest.ReviewController
			controller.Init(c)
			controller.Deletebatch(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/review/count", func(c *gin.Context) {
			var controller rest.ReviewController
			controller.Init(c)
			controller.Count()
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/review/:id", func(c *gin.Context) {
			id_, _ := strconv.ParseInt(c.Param("id"), 10, 64)
			var controller rest.ReviewController
			controller.Init(c)
			controller.Read(id_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/review", func(c *gin.Context) {
			page_, _ := strconv.Atoi(c.Query("page"))
			pagesize_, _ := strconv.Atoi(c.Query("pagesize"))
			var controller rest.ReviewController
			controller.Init(c)
			controller.Index(page_, pagesize_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/review/sum", func(c *gin.Context) {
			var controller rest.ReviewController
			controller.Init(c)
			controller.Sum()
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.POST("/reviewbasic", func(c *gin.Context) {
			item_ := &models.Reviewbasic{}
			c.ShouldBindJSON(item_)
			var controller rest.ReviewbasicController
			controller.Init(c)
			controller.Insert(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.POST("/reviewbasic/batch", func(c *gin.Context) {
			item_ := &[]models.Reviewbasic{}
			c.ShouldBindJSON(item_)
			var controller rest.ReviewbasicController
			controller.Init(c)
			controller.Insertbatch(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.PUT("/reviewbasic", func(c *gin.Context) {
			item_ := &models.Reviewbasic{}
			c.ShouldBindJSON(item_)
			var controller rest.ReviewbasicController
			controller.Init(c)
			controller.Update(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.DELETE("/reviewbasic", func(c *gin.Context) {
			item_ := &models.Reviewbasic{}
			c.ShouldBindJSON(item_)
			var controller rest.ReviewbasicController
			controller.Init(c)
			controller.Delete(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.DELETE("/reviewbasic/batch", func(c *gin.Context) {
			item_ := &[]models.Reviewbasic{}
			c.ShouldBindJSON(item_)
			var controller rest.ReviewbasicController
			controller.Init(c)
			controller.Deletebatch(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/reviewbasic/count", func(c *gin.Context) {
			var controller rest.ReviewbasicController
			controller.Init(c)
			controller.Count()
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/reviewbasic/:id", func(c *gin.Context) {
			id_, _ := strconv.ParseInt(c.Param("id"), 10, 64)
			var controller rest.ReviewbasicController
			controller.Init(c)
			controller.Read(id_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/reviewbasic", func(c *gin.Context) {
			page_, _ := strconv.Atoi(c.Query("page"))
			pagesize_, _ := strconv.Atoi(c.Query("pagesize"))
			var controller rest.ReviewbasicController
			controller.Init(c)
			controller.Index(page_, pagesize_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/reviewcontent/count/apt/:apt", func(c *gin.Context) {
			apt_, _ := strconv.ParseInt(c.Param("apt"), 10, 64)
			var controller rest.ReviewcontentController
			controller.Init(c)
			controller.CountByApt(apt_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/reviewcontent/find/apt/:apt", func(c *gin.Context) {
			apt_, _ := strconv.ParseInt(c.Param("apt"), 10, 64)
			var controller rest.ReviewcontentController
			controller.Init(c)
			controller.FindByApt(apt_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.POST("/reviewcontent", func(c *gin.Context) {
			item_ := &models.Reviewcontent{}
			c.ShouldBindJSON(item_)
			var controller rest.ReviewcontentController
			controller.Init(c)
			controller.Insert(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.POST("/reviewcontent/batch", func(c *gin.Context) {
			item_ := &[]models.Reviewcontent{}
			c.ShouldBindJSON(item_)
			var controller rest.ReviewcontentController
			controller.Init(c)
			controller.Insertbatch(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.PUT("/reviewcontent", func(c *gin.Context) {
			item_ := &models.Reviewcontent{}
			c.ShouldBindJSON(item_)
			var controller rest.ReviewcontentController
			controller.Init(c)
			controller.Update(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.DELETE("/reviewcontent", func(c *gin.Context) {
			item_ := &models.Reviewcontent{}
			c.ShouldBindJSON(item_)
			var controller rest.ReviewcontentController
			controller.Init(c)
			controller.Delete(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.DELETE("/reviewcontent/batch", func(c *gin.Context) {
			item_ := &[]models.Reviewcontent{}
			c.ShouldBindJSON(item_)
			var controller rest.ReviewcontentController
			controller.Init(c)
			controller.Deletebatch(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/reviewcontent/count", func(c *gin.Context) {
			var controller rest.ReviewcontentController
			controller.Init(c)
			controller.Count()
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/reviewcontent/:id", func(c *gin.Context) {
			id_, _ := strconv.ParseInt(c.Param("id"), 10, 64)
			var controller rest.ReviewcontentController
			controller.Init(c)
			controller.Read(id_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/reviewcontent", func(c *gin.Context) {
			page_, _ := strconv.Atoi(c.Query("page"))
			pagesize_, _ := strconv.Atoi(c.Query("pagesize"))
			var controller rest.ReviewcontentController
			controller.Init(c)
			controller.Index(page_, pagesize_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.POST("/reviewcontentbasic", func(c *gin.Context) {
			item_ := &models.Reviewcontentbasic{}
			c.ShouldBindJSON(item_)
			var controller rest.ReviewcontentbasicController
			controller.Init(c)
			controller.Insert(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.POST("/reviewcontentbasic/batch", func(c *gin.Context) {
			item_ := &[]models.Reviewcontentbasic{}
			c.ShouldBindJSON(item_)
			var controller rest.ReviewcontentbasicController
			controller.Init(c)
			controller.Insertbatch(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.PUT("/reviewcontentbasic", func(c *gin.Context) {
			item_ := &models.Reviewcontentbasic{}
			c.ShouldBindJSON(item_)
			var controller rest.ReviewcontentbasicController
			controller.Init(c)
			controller.Update(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.DELETE("/reviewcontentbasic", func(c *gin.Context) {
			item_ := &models.Reviewcontentbasic{}
			c.ShouldBindJSON(item_)
			var controller rest.ReviewcontentbasicController
			controller.Init(c)
			controller.Delete(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.DELETE("/reviewcontentbasic/batch", func(c *gin.Context) {
			item_ := &[]models.Reviewcontentbasic{}
			c.ShouldBindJSON(item_)
			var controller rest.ReviewcontentbasicController
			controller.Init(c)
			controller.Deletebatch(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/reviewcontentbasic/count", func(c *gin.Context) {
			var controller rest.ReviewcontentbasicController
			controller.Init(c)
			controller.Count()
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/reviewcontentbasic/:id", func(c *gin.Context) {
			id_, _ := strconv.ParseInt(c.Param("id"), 10, 64)
			var controller rest.ReviewcontentbasicController
			controller.Init(c)
			controller.Read(id_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/reviewcontentbasic", func(c *gin.Context) {
			page_, _ := strconv.Atoi(c.Query("page"))
			pagesize_, _ := strconv.Atoi(c.Query("pagesize"))
			var controller rest.ReviewcontentbasicController
			controller.Init(c)
			controller.Index(page_, pagesize_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.POST("/reviewdate", func(c *gin.Context) {
			item_ := &models.Reviewdate{}
			c.ShouldBindJSON(item_)
			var controller rest.ReviewdateController
			controller.Init(c)
			controller.Insert(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.POST("/reviewdate/batch", func(c *gin.Context) {
			item_ := &[]models.Reviewdate{}
			c.ShouldBindJSON(item_)
			var controller rest.ReviewdateController
			controller.Init(c)
			controller.Insertbatch(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.PUT("/reviewdate", func(c *gin.Context) {
			item_ := &models.Reviewdate{}
			c.ShouldBindJSON(item_)
			var controller rest.ReviewdateController
			controller.Init(c)
			controller.Update(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.DELETE("/reviewdate", func(c *gin.Context) {
			item_ := &models.Reviewdate{}
			c.ShouldBindJSON(item_)
			var controller rest.ReviewdateController
			controller.Init(c)
			controller.Delete(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.DELETE("/reviewdate/batch", func(c *gin.Context) {
			item_ := &[]models.Reviewdate{}
			c.ShouldBindJSON(item_)
			var controller rest.ReviewdateController
			controller.Init(c)
			controller.Deletebatch(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/reviewdate/count", func(c *gin.Context) {
			var controller rest.ReviewdateController
			controller.Init(c)
			controller.Count()
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/reviewdate/:id", func(c *gin.Context) {
			id_, _ := strconv.ParseInt(c.Param("id"), 10, 64)
			var controller rest.ReviewdateController
			controller.Init(c)
			controller.Read(id_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/reviewdate", func(c *gin.Context) {
			page_, _ := strconv.Atoi(c.Query("page"))
			pagesize_, _ := strconv.Atoi(c.Query("pagesize"))
			var controller rest.ReviewdateController
			controller.Init(c)
			controller.Index(page_, pagesize_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/saving/count/apt/:apt", func(c *gin.Context) {
			apt_, _ := strconv.ParseInt(c.Param("apt"), 10, 64)
			var controller rest.SavingController
			controller.Init(c)
			controller.CountByApt(apt_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/saving/find/apt/:apt", func(c *gin.Context) {
			apt_, _ := strconv.ParseInt(c.Param("apt"), 10, 64)
			var controller rest.SavingController
			controller.Init(c)
			controller.FindByApt(apt_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.DELETE("/saving/byapt", func(c *gin.Context) {
			item_ := &models.Saving{}
			c.ShouldBindJSON(item_)
			var controller rest.SavingController
			controller.Init(c)
			controller.DeleteByApt(item_.Apt)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.POST("/saving", func(c *gin.Context) {
			item_ := &models.Saving{}
			c.ShouldBindJSON(item_)
			var controller rest.SavingController
			controller.Init(c)
			controller.Insert(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.POST("/saving/batch", func(c *gin.Context) {
			item_ := &[]models.Saving{}
			c.ShouldBindJSON(item_)
			var controller rest.SavingController
			controller.Init(c)
			controller.Insertbatch(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.PUT("/saving", func(c *gin.Context) {
			item_ := &models.Saving{}
			c.ShouldBindJSON(item_)
			var controller rest.SavingController
			controller.Init(c)
			controller.Update(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.DELETE("/saving", func(c *gin.Context) {
			item_ := &models.Saving{}
			c.ShouldBindJSON(item_)
			var controller rest.SavingController
			controller.Init(c)
			controller.Delete(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.DELETE("/saving/batch", func(c *gin.Context) {
			item_ := &[]models.Saving{}
			c.ShouldBindJSON(item_)
			var controller rest.SavingController
			controller.Init(c)
			controller.Deletebatch(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/saving/count", func(c *gin.Context) {
			var controller rest.SavingController
			controller.Init(c)
			controller.Count()
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/saving/:id", func(c *gin.Context) {
			id_, _ := strconv.ParseInt(c.Param("id"), 10, 64)
			var controller rest.SavingController
			controller.Init(c)
			controller.Read(id_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/saving", func(c *gin.Context) {
			page_, _ := strconv.Atoi(c.Query("page"))
			pagesize_, _ := strconv.Atoi(c.Query("pagesize"))
			var controller rest.SavingController
			controller.Init(c)
			controller.Index(page_, pagesize_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/standard/count/apt/:apt", func(c *gin.Context) {
			apt_, _ := strconv.ParseInt(c.Param("apt"), 10, 64)
			var controller rest.StandardController
			controller.Init(c)
			controller.CountByApt(apt_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/standard/find/apt/:apt", func(c *gin.Context) {
			apt_, _ := strconv.ParseInt(c.Param("apt"), 10, 64)
			var controller rest.StandardController
			controller.Init(c)
			controller.FindByApt(apt_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/standard/get/aptcategoryname/:apt", func(c *gin.Context) {
			apt_, _ := strconv.ParseInt(c.Param("apt"), 10, 64)
			category_, _ := strconv.ParseInt(c.Query("category"), 10, 64)
			name_ := c.Query("name")
			var controller rest.StandardController
			controller.Init(c)
			controller.GetByAptCategoryName(apt_, category_, name_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/standard/find/category/:category", func(c *gin.Context) {
			category_, _ := strconv.ParseInt(c.Param("category"), 10, 64)
			var controller rest.StandardController
			controller.Init(c)
			controller.FindByCategory(category_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.DELETE("/standard/bycategory", func(c *gin.Context) {
			item_ := &models.Standard{}
			c.ShouldBindJSON(item_)
			var controller rest.StandardController
			controller.Init(c)
			controller.DeleteByCategory(item_.Category)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.POST("/standard", func(c *gin.Context) {
			item_ := &models.Standard{}
			c.ShouldBindJSON(item_)
			var controller rest.StandardController
			controller.Init(c)
			controller.Insert(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.POST("/standard/batch", func(c *gin.Context) {
			item_ := &[]models.Standard{}
			c.ShouldBindJSON(item_)
			var controller rest.StandardController
			controller.Init(c)
			controller.Insertbatch(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.PUT("/standard", func(c *gin.Context) {
			item_ := &models.Standard{}
			c.ShouldBindJSON(item_)
			var apicontroller api.StandardController
			apicontroller.Init(c)
			var controller rest.StandardController
			controller.Init(c)
			apicontroller.Pre_Update(item_)
			controller.Update(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.DELETE("/standard", func(c *gin.Context) {
			item_ := &models.Standard{}
			c.ShouldBindJSON(item_)
			var apicontroller api.StandardController
			apicontroller.Init(c)
			var controller rest.StandardController
			controller.Init(c)
			apicontroller.Pre_Delete(item_)
			controller.Delete(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.DELETE("/standard/batch", func(c *gin.Context) {
			item_ := &[]models.Standard{}
			c.ShouldBindJSON(item_)
			var apicontroller api.StandardController
			apicontroller.Init(c)
			var controller rest.StandardController
			controller.Init(c)
			apicontroller.Pre_Deletebatch(item_)
			controller.Deletebatch(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/standard/count", func(c *gin.Context) {
			var controller rest.StandardController
			controller.Init(c)
			controller.Count()
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/standard/:id", func(c *gin.Context) {
			id_, _ := strconv.ParseInt(c.Param("id"), 10, 64)
			var controller rest.StandardController
			controller.Init(c)
			controller.Read(id_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/standard", func(c *gin.Context) {
			page_, _ := strconv.Atoi(c.Query("page"))
			pagesize_, _ := strconv.Atoi(c.Query("pagesize"))
			var controller rest.StandardController
			controller.Init(c)
			controller.Index(page_, pagesize_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.POST("/standardbackup", func(c *gin.Context) {
			item_ := &models.Standardbackup{}
			c.ShouldBindJSON(item_)
			var controller rest.StandardbackupController
			controller.Init(c)
			controller.Insert(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.POST("/standardbackup/batch", func(c *gin.Context) {
			item_ := &[]models.Standardbackup{}
			c.ShouldBindJSON(item_)
			var controller rest.StandardbackupController
			controller.Init(c)
			controller.Insertbatch(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.PUT("/standardbackup", func(c *gin.Context) {
			item_ := &models.Standardbackup{}
			c.ShouldBindJSON(item_)
			var controller rest.StandardbackupController
			controller.Init(c)
			controller.Update(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.DELETE("/standardbackup", func(c *gin.Context) {
			item_ := &models.Standardbackup{}
			c.ShouldBindJSON(item_)
			var controller rest.StandardbackupController
			controller.Init(c)
			controller.Delete(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.DELETE("/standardbackup/batch", func(c *gin.Context) {
			item_ := &[]models.Standardbackup{}
			c.ShouldBindJSON(item_)
			var controller rest.StandardbackupController
			controller.Init(c)
			controller.Deletebatch(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/standardbackup/count", func(c *gin.Context) {
			var controller rest.StandardbackupController
			controller.Init(c)
			controller.Count()
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/standardbackup/:id", func(c *gin.Context) {
			id_, _ := strconv.ParseInt(c.Param("id"), 10, 64)
			var controller rest.StandardbackupController
			controller.Init(c)
			controller.Read(id_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/standardbackup", func(c *gin.Context) {
			page_, _ := strconv.Atoi(c.Query("page"))
			pagesize_, _ := strconv.Atoi(c.Query("pagesize"))
			var controller rest.StandardbackupController
			controller.Init(c)
			controller.Index(page_, pagesize_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/standardhistory/count/apt/:apt", func(c *gin.Context) {
			apt_, _ := strconv.ParseInt(c.Param("apt"), 10, 64)
			var controller rest.StandardhistoryController
			controller.Init(c)
			controller.CountByApt(apt_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/standardhistory/find/apt/:apt", func(c *gin.Context) {
			apt_, _ := strconv.ParseInt(c.Param("apt"), 10, 64)
			var controller rest.StandardhistoryController
			controller.Init(c)
			controller.FindByApt(apt_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.DELETE("/standardhistory/byapt", func(c *gin.Context) {
			item_ := &models.Standardhistory{}
			c.ShouldBindJSON(item_)
			var controller rest.StandardhistoryController
			controller.Init(c)
			controller.DeleteByApt(item_.Apt)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/standardhistory/count/standard/:standard", func(c *gin.Context) {
			standard_, _ := strconv.ParseInt(c.Param("standard"), 10, 64)
			var controller rest.StandardhistoryController
			controller.Init(c)
			controller.CountByStandard(standard_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/standardhistory/get/standard/:standard", func(c *gin.Context) {
			standard_, _ := strconv.ParseInt(c.Param("standard"), 10, 64)
			var controller rest.StandardhistoryController
			controller.Init(c)
			controller.GetByStandard(standard_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.POST("/standardhistory", func(c *gin.Context) {
			item_ := &models.Standardhistory{}
			c.ShouldBindJSON(item_)
			var controller rest.StandardhistoryController
			controller.Init(c)
			controller.Insert(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.POST("/standardhistory/batch", func(c *gin.Context) {
			item_ := &[]models.Standardhistory{}
			c.ShouldBindJSON(item_)
			var controller rest.StandardhistoryController
			controller.Init(c)
			controller.Insertbatch(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.PUT("/standardhistory", func(c *gin.Context) {
			item_ := &models.Standardhistory{}
			c.ShouldBindJSON(item_)
			var controller rest.StandardhistoryController
			controller.Init(c)
			controller.Update(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.DELETE("/standardhistory", func(c *gin.Context) {
			item_ := &models.Standardhistory{}
			c.ShouldBindJSON(item_)
			var controller rest.StandardhistoryController
			controller.Init(c)
			controller.Delete(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.DELETE("/standardhistory/batch", func(c *gin.Context) {
			item_ := &[]models.Standardhistory{}
			c.ShouldBindJSON(item_)
			var controller rest.StandardhistoryController
			controller.Init(c)
			controller.Deletebatch(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/standardhistory/count", func(c *gin.Context) {
			var controller rest.StandardhistoryController
			controller.Init(c)
			controller.Count()
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/standardhistory/:id", func(c *gin.Context) {
			id_, _ := strconv.ParseInt(c.Param("id"), 10, 64)
			var controller rest.StandardhistoryController
			controller.Init(c)
			controller.Read(id_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/standardhistory", func(c *gin.Context) {
			page_, _ := strconv.Atoi(c.Query("page"))
			pagesize_, _ := strconv.Atoi(c.Query("pagesize"))
			var controller rest.StandardhistoryController
			controller.Init(c)
			controller.Index(page_, pagesize_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/standardlist/count/apt/:apt", func(c *gin.Context) {
			apt_, _ := strconv.ParseInt(c.Param("apt"), 10, 64)
			var controller rest.StandardlistController
			controller.Init(c)
			controller.CountByApt(apt_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/standardlist/find/apt/:apt", func(c *gin.Context) {
			apt_, _ := strconv.ParseInt(c.Param("apt"), 10, 64)
			var controller rest.StandardlistController
			controller.Init(c)
			controller.FindByApt(apt_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/standardlist/count", func(c *gin.Context) {
			var controller rest.StandardlistController
			controller.Init(c)
			controller.Count()
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/standardlist/:id", func(c *gin.Context) {
			id_, _ := strconv.ParseInt(c.Param("id"), 10, 64)
			var controller rest.StandardlistController
			controller.Init(c)
			controller.Read(id_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/standardlist", func(c *gin.Context) {
			page_, _ := strconv.Atoi(c.Query("page"))
			pagesize_, _ := strconv.Atoi(c.Query("pagesize"))
			var controller rest.StandardlistController
			controller.Init(c)
			controller.Index(page_, pagesize_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.POST("/standardwage", func(c *gin.Context) {
			item_ := &models.Standardwage{}
			c.ShouldBindJSON(item_)
			var controller rest.StandardwageController
			controller.Init(c)
			controller.Insert(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.POST("/standardwage/batch", func(c *gin.Context) {
			item_ := &[]models.Standardwage{}
			c.ShouldBindJSON(item_)
			var controller rest.StandardwageController
			controller.Init(c)
			controller.Insertbatch(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.PUT("/standardwage", func(c *gin.Context) {
			item_ := &models.Standardwage{}
			c.ShouldBindJSON(item_)
			var controller rest.StandardwageController
			controller.Init(c)
			controller.Update(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.DELETE("/standardwage", func(c *gin.Context) {
			item_ := &models.Standardwage{}
			c.ShouldBindJSON(item_)
			var controller rest.StandardwageController
			controller.Init(c)
			controller.Delete(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.DELETE("/standardwage/batch", func(c *gin.Context) {
			item_ := &[]models.Standardwage{}
			c.ShouldBindJSON(item_)
			var controller rest.StandardwageController
			controller.Init(c)
			controller.Deletebatch(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/standardwage/count", func(c *gin.Context) {
			var controller rest.StandardwageController
			controller.Init(c)
			controller.Count()
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/standardwage/:id", func(c *gin.Context) {
			id_, _ := strconv.ParseInt(c.Param("id"), 10, 64)
			var controller rest.StandardwageController
			controller.Init(c)
			controller.Read(id_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/standardwage", func(c *gin.Context) {
			page_, _ := strconv.Atoi(c.Query("page"))
			pagesize_, _ := strconv.Atoi(c.Query("pagesize"))
			var controller rest.StandardwageController
			controller.Init(c)
			controller.Index(page_, pagesize_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.POST("/technician", func(c *gin.Context) {
			item_ := &models.Technician{}
			c.ShouldBindJSON(item_)
			var controller rest.TechnicianController
			controller.Init(c)
			controller.Insert(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.POST("/technician/batch", func(c *gin.Context) {
			item_ := &[]models.Technician{}
			c.ShouldBindJSON(item_)
			var controller rest.TechnicianController
			controller.Init(c)
			controller.Insertbatch(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.PUT("/technician", func(c *gin.Context) {
			item_ := &models.Technician{}
			c.ShouldBindJSON(item_)
			var controller rest.TechnicianController
			controller.Init(c)
			controller.Update(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.DELETE("/technician", func(c *gin.Context) {
			item_ := &models.Technician{}
			c.ShouldBindJSON(item_)
			var controller rest.TechnicianController
			controller.Init(c)
			controller.Delete(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.DELETE("/technician/batch", func(c *gin.Context) {
			item_ := &[]models.Technician{}
			c.ShouldBindJSON(item_)
			var controller rest.TechnicianController
			controller.Init(c)
			controller.Deletebatch(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/technician/count", func(c *gin.Context) {
			var controller rest.TechnicianController
			controller.Init(c)
			controller.Count()
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/technician/:id", func(c *gin.Context) {
			id_, _ := strconv.ParseInt(c.Param("id"), 10, 64)
			var controller rest.TechnicianController
			controller.Init(c)
			controller.Read(id_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/technician", func(c *gin.Context) {
			page_, _ := strconv.Atoi(c.Query("page"))
			pagesize_, _ := strconv.Atoi(c.Query("pagesize"))
			var controller rest.TechnicianController
			controller.Init(c)
			controller.Index(page_, pagesize_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/totalreport/count/apt/:apt", func(c *gin.Context) {
			apt_, _ := strconv.ParseInt(c.Param("apt"), 10, 64)
			var controller rest.TotalreportController
			controller.Init(c)
			controller.CountByApt(apt_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/totalreport/find/apt/:apt", func(c *gin.Context) {
			apt_, _ := strconv.ParseInt(c.Param("apt"), 10, 64)
			var controller rest.TotalreportController
			controller.Init(c)
			controller.FindByApt(apt_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/totalreport/count", func(c *gin.Context) {
			var controller rest.TotalreportController
			controller.Init(c)
			controller.Count()
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/totalreport/:id", func(c *gin.Context) {
			id_, _ := strconv.ParseInt(c.Param("id"), 10, 64)
			var controller rest.TotalreportController
			controller.Init(c)
			controller.Read(id_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/totalreport", func(c *gin.Context) {
			page_, _ := strconv.Atoi(c.Query("page"))
			pagesize_, _ := strconv.Atoi(c.Query("pagesize"))
			var controller rest.TotalreportController
			controller.Init(c)
			controller.Index(page_, pagesize_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/totalreport/sum", func(c *gin.Context) {
			var controller rest.TotalreportController
			controller.Init(c)
			controller.Sum()
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/totalyearreport/count/apt/:apt", func(c *gin.Context) {
			apt_, _ := strconv.ParseInt(c.Param("apt"), 10, 64)
			var controller rest.TotalyearreportController
			controller.Init(c)
			controller.CountByApt(apt_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/totalyearreport/find/apt/:apt", func(c *gin.Context) {
			apt_, _ := strconv.ParseInt(c.Param("apt"), 10, 64)
			var controller rest.TotalyearreportController
			controller.Init(c)
			controller.FindByApt(apt_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/totalyearreport/count", func(c *gin.Context) {
			var controller rest.TotalyearreportController
			controller.Init(c)
			controller.Count()
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/totalyearreport/:id", func(c *gin.Context) {
			id_, _ := strconv.ParseInt(c.Param("id"), 10, 64)
			var controller rest.TotalyearreportController
			controller.Init(c)
			controller.Read(id_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/totalyearreport", func(c *gin.Context) {
			page_, _ := strconv.Atoi(c.Query("page"))
			pagesize_, _ := strconv.Atoi(c.Query("pagesize"))
			var controller rest.TotalyearreportController
			controller.Init(c)
			controller.Index(page_, pagesize_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/totalyearreport/sum", func(c *gin.Context) {
			var controller rest.TotalyearreportController
			controller.Init(c)
			controller.Sum()
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/user/count/level/:level", func(c *gin.Context) {
			var level_ user.Level
			level__, _ := strconv.Atoi(c.Param("level"))
			level_ = user.Level(level__)
			var controller rest.UserController
			controller.Init(c)
			controller.CountByLevel(level_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/user/find/level/:level", func(c *gin.Context) {
			var level_ user.Level
			level__, _ := strconv.Atoi(c.Param("level"))
			level_ = user.Level(level__)
			var controller rest.UserController
			controller.Init(c)
			controller.FindByLevel(level_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/user/count/apt/:apt", func(c *gin.Context) {
			apt_, _ := strconv.ParseInt(c.Param("apt"), 10, 64)
			var controller rest.UserController
			controller.Init(c)
			controller.CountByApt(apt_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/user/find/apt/:apt", func(c *gin.Context) {
			apt_, _ := strconv.ParseInt(c.Param("apt"), 10, 64)
			var controller rest.UserController
			controller.Init(c)
			controller.FindByApt(apt_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/user/count/aptlevel/:apt", func(c *gin.Context) {
			apt_, _ := strconv.ParseInt(c.Param("apt"), 10, 64)
			var level_ user.Level
			level__, _ := strconv.Atoi(c.Query("level"))
			level_ = user.Level(level__)
			var controller rest.UserController
			controller.Init(c)
			controller.CountByAptLevel(apt_, level_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/user/find/aptlevel/:apt", func(c *gin.Context) {
			apt_, _ := strconv.ParseInt(c.Param("apt"), 10, 64)
			var level_ user.Level
			level__, _ := strconv.Atoi(c.Query("level"))
			level_ = user.Level(level__)
			var controller rest.UserController
			controller.Init(c)
			controller.FindByAptLevel(apt_, level_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/user/get/email/:email", func(c *gin.Context) {
			email_ := c.Param("email")
			var controller rest.UserController
			controller.Init(c)
			controller.GetByEmail(email_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/user/count/email/:email", func(c *gin.Context) {
			email_ := c.Param("email")
			var controller rest.UserController
			controller.Init(c)
			controller.CountByEmail(email_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/user/get/loginid/:loginid", func(c *gin.Context) {
			loginid_ := c.Param("loginid")
			var controller rest.UserController
			controller.Init(c)
			controller.GetByLoginid(loginid_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/user/count/loginid/:loginid", func(c *gin.Context) {
			loginid_ := c.Param("loginid")
			var controller rest.UserController
			controller.Init(c)
			controller.CountByLoginid(loginid_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.POST("/user", func(c *gin.Context) {
			item_ := &models.UserUpdate{}
			c.ShouldBindJSON(item_)
			var controller rest.UserController
			controller.Init(c)
			controller.Insert(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.POST("/user/batch", func(c *gin.Context) {
			item_ := &[]models.UserUpdate{}
			c.ShouldBindJSON(item_)
			var controller rest.UserController
			controller.Init(c)
			controller.Insertbatch(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.PUT("/user", func(c *gin.Context) {
			item_ := &models.UserUpdate{}
			c.ShouldBindJSON(item_)
			var controller rest.UserController
			controller.Init(c)
			controller.Update(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.DELETE("/user", func(c *gin.Context) {
			item_ := &models.User{}
			c.ShouldBindJSON(item_)
			var controller rest.UserController
			controller.Init(c)
			controller.Delete(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.DELETE("/user/batch", func(c *gin.Context) {
			item_ := &[]models.User{}
			c.ShouldBindJSON(item_)
			var controller rest.UserController
			controller.Init(c)
			controller.Deletebatch(item_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/user/count", func(c *gin.Context) {
			var controller rest.UserController
			controller.Init(c)
			controller.Count()
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/user/:id", func(c *gin.Context) {
			id_, _ := strconv.ParseInt(c.Param("id"), 10, 64)
			var controller rest.UserController
			controller.Init(c)
			controller.Read(id_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/user", func(c *gin.Context) {
			page_, _ := strconv.Atoi(c.Query("page"))
			pagesize_, _ := strconv.Atoi(c.Query("pagesize"))
			var controller rest.UserController
			controller.Init(c)
			controller.Index(page_, pagesize_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/yearreport/count/apt/:apt", func(c *gin.Context) {
			apt_, _ := strconv.ParseInt(c.Param("apt"), 10, 64)
			var controller rest.YearreportController
			controller.Init(c)
			controller.CountByApt(apt_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/yearreport/find/apt/:apt", func(c *gin.Context) {
			apt_, _ := strconv.ParseInt(c.Param("apt"), 10, 64)
			var controller rest.YearreportController
			controller.Init(c)
			controller.FindByApt(apt_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/yearreport/count", func(c *gin.Context) {
			var controller rest.YearreportController
			controller.Init(c)
			controller.Count()
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/yearreport/:id", func(c *gin.Context) {
			id_, _ := strconv.ParseInt(c.Param("id"), 10, 64)
			var controller rest.YearreportController
			controller.Init(c)
			controller.Read(id_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/yearreport", func(c *gin.Context) {
			page_, _ := strconv.Atoi(c.Query("page"))
			pagesize_, _ := strconv.Atoi(c.Query("pagesize"))
			var controller rest.YearreportController
			controller.Init(c)
			controller.Index(page_, pagesize_)
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

		apiGroup.GET("/yearreport/sum", func(c *gin.Context) {
			var controller rest.YearreportController
			controller.Init(c)
			controller.Sum()
			controller.Close()
			c.JSON(controller.Code, controller.Result)
		})

	}

}