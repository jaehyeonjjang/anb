package rest


import (
	"repair/controllers"
	"repair/models"

    "strings"
)

type RepairController struct {
	controllers.Controller
}


// @Put()
func (c *RepairController) UpdateStatusById(status int ,id int64) {
    
    conn := c.NewConnection()

	_manager := models.NewRepairManager(conn)
    
    err := _manager.UpdateStatusById(status, id)
    if err != nil {
        c.Set("code", "error")    
        c.Set("error", err)
        return
    }
    
}


func (c *RepairController) Insert(item *models.Repair) {
    
    
	conn := c.NewConnection()
    
	manager := models.NewRepairManager(conn)
	err := manager.Insert(item)
    if err != nil {
        c.Set("code", "error")    
        c.Set("error", err)
        return
    }

    id := manager.GetIdentity()
    c.Result["id"] = id
    item.Id = id
}

func (c *RepairController) Insertbatch(item *[]models.Repair) {  
    if item == nil || len(*item) == 0 {
        return
    }

    rows := len(*item)
    
    
    
	conn := c.NewConnection()
    
	manager := models.NewRepairManager(conn)

    for i := 0; i < rows; i++ {
	    err := manager.Insert(&((*item)[i]))
        if err != nil {
            c.Set("code", "error")    
            c.Set("error", err)
            return
        }
    }
}

func (c *RepairController) Update(item *models.Repair) {
    
    
	conn := c.NewConnection()

	manager := models.NewRepairManager(conn)
    err := manager.Update(item)
    if err != nil {
        c.Set("code", "error")    
        c.Set("error", err)
        return
    }
}

func (c *RepairController) Delete(item *models.Repair) {
    
    
    conn := c.NewConnection()

	manager := models.NewRepairManager(conn)

    
	err := manager.Delete(item.Id)
    if err != nil {
        c.Set("code", "error")    
        c.Set("error", err)
    }
}

func (c *RepairController) Deletebatch(item *[]models.Repair) {
    
    
    conn := c.NewConnection()

	manager := models.NewRepairManager(conn)

    for _, v := range *item {
        
    
	    err := manager.Delete(v.Id)
        if err != nil {
            c.Set("code", "error")    
            c.Set("error", err)
            return
        }
    }
}

func (c *RepairController) Count() {
    
    
	conn := c.NewConnection()

	manager := models.NewRepairManager(conn)

    var args []interface{}
    
    _type := c.Geti("type")
    if _type != 0 {
        args = append(args, models.Where{Column:"type", Value:_type, Compare:"="})    
    }
    _status := c.Geti("status")
    if _status != 0 {
        args = append(args, models.Where{Column:"status", Value:_status, Compare:"="})    
    }
    _calculatetype := c.Geti("calculatetype")
    if _calculatetype != 0 {
        args = append(args, models.Where{Column:"calculatetype", Value:_calculatetype, Compare:"="})    
    }
    _provision := c.Geti("provision")
    if _provision != 0 {
        args = append(args, models.Where{Column:"provision", Value:_provision, Compare:"="})    
    }
    _complex1 := c.Get("complex1")
    if _complex1 != "" {
        args = append(args, models.Where{Column:"complex1", Value:_complex1, Compare:"="})
    }
    _complex2 := c.Get("complex2")
    if _complex2 != "" {
        args = append(args, models.Where{Column:"complex2", Value:_complex2, Compare:"="})
    }
    _completionyear := c.Geti("completionyear")
    if _completionyear != 0 {
        args = append(args, models.Where{Column:"completionyear", Value:_completionyear, Compare:"="})    
    }
    _completionmonth := c.Geti("completionmonth")
    if _completionmonth != 0 {
        args = append(args, models.Where{Column:"completionmonth", Value:_completionmonth, Compare:"="})    
    }
    _completionday := c.Geti("completionday")
    if _completionday != 0 {
        args = append(args, models.Where{Column:"completionday", Value:_completionday, Compare:"="})    
    }
    _parcelrate := c.Geti("parcelrate")
    if _parcelrate != 0 {
        args = append(args, models.Where{Column:"parcelrate", Value:_parcelrate, Compare:"="})    
    }
    _planyears := c.Geti("planyears")
    if _planyears != 0 {
        args = append(args, models.Where{Column:"planyears", Value:_planyears, Compare:"="})    
    }
    _info1 := c.Get("info1")
    if _info1 != "" {
        args = append(args, models.Where{Column:"info1", Value:_info1, Compare:"="})
    }
    _info2 := c.Get("info2")
    if _info2 != "" {
        args = append(args, models.Where{Column:"info2", Value:_info2, Compare:"="})
    }
    _info3 := c.Get("info3")
    if _info3 != "" {
        args = append(args, models.Where{Column:"info3", Value:_info3, Compare:"="})
    }
    _info4 := c.Get("info4")
    if _info4 != "" {
        args = append(args, models.Where{Column:"info4", Value:_info4, Compare:"="})
    }
    _info5 := c.Get("info5")
    if _info5 != "" {
        args = append(args, models.Where{Column:"info5", Value:_info5, Compare:"="})
    }
    _info6 := c.Get("info6")
    if _info6 != "" {
        args = append(args, models.Where{Column:"info6", Value:_info6, Compare:"="})
    }
    _info7 := c.Get("info7")
    if _info7 != "" {
        args = append(args, models.Where{Column:"info7", Value:_info7, Compare:"="})
    }
    _info8 := c.Get("info8")
    if _info8 != "" {
        args = append(args, models.Where{Column:"info8", Value:_info8, Compare:"="})
    }
    _info9 := c.Get("info9")
    if _info9 != "" {
        args = append(args, models.Where{Column:"info9", Value:_info9, Compare:"="})
    }
    _info10 := c.Get("info10")
    if _info10 != "" {
        args = append(args, models.Where{Column:"info10", Value:_info10, Compare:"="})
    }
    _info11 := c.Get("info11")
    if _info11 != "" {
        args = append(args, models.Where{Column:"info11", Value:_info11, Compare:"="})
    }
    _structure1 := c.Get("structure1")
    if _structure1 != "" {
        args = append(args, models.Where{Column:"structure1", Value:_structure1, Compare:"="})
    }
    _structure2 := c.Get("structure2")
    if _structure2 != "" {
        args = append(args, models.Where{Column:"structure2", Value:_structure2, Compare:"="})
    }
    _structure3 := c.Get("structure3")
    if _structure3 != "" {
        args = append(args, models.Where{Column:"structure3", Value:_structure3, Compare:"="})
    }
    _structure4 := c.Get("structure4")
    if _structure4 != "" {
        args = append(args, models.Where{Column:"structure4", Value:_structure4, Compare:"="})
    }
    _structure5 := c.Get("structure5")
    if _structure5 != "" {
        args = append(args, models.Where{Column:"structure5", Value:_structure5, Compare:"="})
    }
    _structure6 := c.Get("structure6")
    if _structure6 != "" {
        args = append(args, models.Where{Column:"structure6", Value:_structure6, Compare:"="})
    }
    _structure7 := c.Get("structure7")
    if _structure7 != "" {
        args = append(args, models.Where{Column:"structure7", Value:_structure7, Compare:"="})
    }
    _structure8 := c.Get("structure8")
    if _structure8 != "" {
        args = append(args, models.Where{Column:"structure8", Value:_structure8, Compare:"="})
    }
    _structure9 := c.Get("structure9")
    if _structure9 != "" {
        args = append(args, models.Where{Column:"structure9", Value:_structure9, Compare:"="})
    }
    _structure10 := c.Get("structure10")
    if _structure10 != "" {
        args = append(args, models.Where{Column:"structure10", Value:_structure10, Compare:"="})
    }
    _structure11 := c.Get("structure11")
    if _structure11 != "" {
        args = append(args, models.Where{Column:"structure11", Value:_structure11, Compare:"="})
    }
    _structure12 := c.Get("structure12")
    if _structure12 != "" {
        args = append(args, models.Where{Column:"structure12", Value:_structure12, Compare:"="})
    }
    _structure13 := c.Get("structure13")
    if _structure13 != "" {
        args = append(args, models.Where{Column:"structure13", Value:_structure13, Compare:"="})
    }
    _structure14 := c.Get("structure14")
    if _structure14 != "" {
        args = append(args, models.Where{Column:"structure14", Value:_structure14, Compare:"="})
    }
    _reviewcontent1 := c.Get("reviewcontent1")
    if _reviewcontent1 != "" {
        args = append(args, models.Where{Column:"reviewcontent1", Value:_reviewcontent1, Compare:"="})
    }
    _reviewcontent2 := c.Get("reviewcontent2")
    if _reviewcontent2 != "" {
        args = append(args, models.Where{Column:"reviewcontent2", Value:_reviewcontent2, Compare:"="})
    }
    _reviewcontent3 := c.Get("reviewcontent3")
    if _reviewcontent3 != "" {
        args = append(args, models.Where{Column:"reviewcontent3", Value:_reviewcontent3, Compare:"="})
    }
    _reviewcontent4 := c.Get("reviewcontent4")
    if _reviewcontent4 != "" {
        args = append(args, models.Where{Column:"reviewcontent4", Value:_reviewcontent4, Compare:"="})
    }
    _reviewcontent5 := c.Get("reviewcontent5")
    if _reviewcontent5 != "" {
        args = append(args, models.Where{Column:"reviewcontent5", Value:_reviewcontent5, Compare:"="})
    }
    _reviewcontent6 := c.Get("reviewcontent6")
    if _reviewcontent6 != "" {
        args = append(args, models.Where{Column:"reviewcontent6", Value:_reviewcontent6, Compare:"="})
    }
    _reviewcontent7 := c.Get("reviewcontent7")
    if _reviewcontent7 != "" {
        args = append(args, models.Where{Column:"reviewcontent7", Value:_reviewcontent7, Compare:"="})
    }
    _savingprice := c.Geti("savingprice")
    if _savingprice != 0 {
        args = append(args, models.Where{Column:"savingprice", Value:_savingprice, Compare:"="})    
    }
    _price1 := c.Get("price1")
    if _price1 != "" {
        args = append(args, models.Where{Column:"price1", Value:_price1, Compare:"="})
    }
    _price2 := c.Get("price2")
    if _price2 != "" {
        args = append(args, models.Where{Column:"price2", Value:_price2, Compare:"="})
    }
    _price3 := c.Get("price3")
    if _price3 != "" {
        args = append(args, models.Where{Column:"price3", Value:_price3, Compare:"="})
    }
    _price4 := c.Get("price4")
    if _price4 != "" {
        args = append(args, models.Where{Column:"price4", Value:_price4, Compare:"="})
    }
    _price5 := c.Get("price5")
    if _price5 != "" {
        args = append(args, models.Where{Column:"price5", Value:_price5, Compare:"="})
    }
    _reportdate := c.Get("reportdate")
    if _reportdate != "" {
        args = append(args, models.Where{Column:"reportdate", Value:_reportdate, Compare:"="})
    }
    _content1 := c.Get("content1")
    if _content1 != "" {
        args = append(args, models.Where{Column:"content1", Value:_content1, Compare:"="})
    }
    _content2 := c.Get("content2")
    if _content2 != "" {
        args = append(args, models.Where{Column:"content2", Value:_content2, Compare:"="})
    }
    _content3 := c.Get("content3")
    if _content3 != "" {
        args = append(args, models.Where{Column:"content3", Value:_content3, Compare:"="})
    }
    _content4 := c.Get("content4")
    if _content4 != "" {
        args = append(args, models.Where{Column:"content4", Value:_content4, Compare:"="})
    }
    _content5 := c.Get("content5")
    if _content5 != "" {
        args = append(args, models.Where{Column:"content5", Value:_content5, Compare:"="})
    }
    _content6 := c.Get("content6")
    if _content6 != "" {
        args = append(args, models.Where{Column:"content6", Value:_content6, Compare:"="})
    }
    _content7 := c.Get("content7")
    if _content7 != "" {
        args = append(args, models.Where{Column:"content7", Value:_content7, Compare:"="})
    }
    _content8 := c.Get("content8")
    if _content8 != "" {
        args = append(args, models.Where{Column:"content8", Value:_content8, Compare:"="})
    }
    _content9 := c.Get("content9")
    if _content9 != "" {
        args = append(args, models.Where{Column:"content9", Value:_content9, Compare:"="})
    }
    _content10 := c.Get("content10")
    if _content10 != "" {
        args = append(args, models.Where{Column:"content10", Value:_content10, Compare:"="})
    }
    _content11 := c.Get("content11")
    if _content11 != "" {
        args = append(args, models.Where{Column:"content11", Value:_content11, Compare:"="})
    }
    _periodtype := c.Geti("periodtype")
    if _periodtype != 0 {
        args = append(args, models.Where{Column:"periodtype", Value:_periodtype, Compare:"="})    
    }
    _remark := c.Get("remark")
    if _remark != "" {
        args = append(args, models.Where{Column:"remark", Value:_remark, Compare:"="})
    }
    _apt := c.Geti64("apt")
    if _apt != 0 {
        args = append(args, models.Where{Column:"apt", Value:_apt, Compare:"="})    
    }
    _startdate := c.Get("startdate")
    _enddate := c.Get("enddate")

    if _startdate != "" && _enddate != "" {        
        var v [2]string
        v[0] = _startdate
        v[1] = _enddate  
        args = append(args, models.Where{Column:"date", Value:v, Compare:"between"})    
    } else if  _startdate != "" {          
        args = append(args, models.Where{Column:"date", Value:_startdate, Compare:">="})
    } else if  _enddate != "" {          
        args = append(args, models.Where{Column:"date", Value:_enddate, Compare:"<="})            
    }
    

    
    
    total := manager.Count(args)
	c.Set("total", total)
}


func (c *RepairController) Read(id int64) {
    
    
	conn := c.NewConnection()

	manager := models.NewRepairManager(conn)
	item := manager.Get(id)

    
    
    c.Set("item", item)
}

func (c *RepairController) Index(page int, pagesize int) {
    
    
	conn := c.NewConnection()

	manager := models.NewRepairManager(conn)

    var args []interface{}
    
    _type := c.Geti("type")
    if _type != 0 {
        args = append(args, models.Where{Column:"type", Value:_type, Compare:"="})    
    }
    _status := c.Geti("status")
    if _status != 0 {
        args = append(args, models.Where{Column:"status", Value:_status, Compare:"="})    
    }
    _calculatetype := c.Geti("calculatetype")
    if _calculatetype != 0 {
        args = append(args, models.Where{Column:"calculatetype", Value:_calculatetype, Compare:"="})    
    }
    _provision := c.Geti("provision")
    if _provision != 0 {
        args = append(args, models.Where{Column:"provision", Value:_provision, Compare:"="})    
    }
    _complex1 := c.Get("complex1")
    if _complex1 != "" {
        args = append(args, models.Where{Column:"complex1", Value:_complex1, Compare:"="})
    }
    _complex2 := c.Get("complex2")
    if _complex2 != "" {
        args = append(args, models.Where{Column:"complex2", Value:_complex2, Compare:"="})
    }
    _completionyear := c.Geti("completionyear")
    if _completionyear != 0 {
        args = append(args, models.Where{Column:"completionyear", Value:_completionyear, Compare:"="})    
    }
    _completionmonth := c.Geti("completionmonth")
    if _completionmonth != 0 {
        args = append(args, models.Where{Column:"completionmonth", Value:_completionmonth, Compare:"="})    
    }
    _completionday := c.Geti("completionday")
    if _completionday != 0 {
        args = append(args, models.Where{Column:"completionday", Value:_completionday, Compare:"="})    
    }
    _parcelrate := c.Geti("parcelrate")
    if _parcelrate != 0 {
        args = append(args, models.Where{Column:"parcelrate", Value:_parcelrate, Compare:"="})    
    }
    _planyears := c.Geti("planyears")
    if _planyears != 0 {
        args = append(args, models.Where{Column:"planyears", Value:_planyears, Compare:"="})    
    }
    _info1 := c.Get("info1")
    if _info1 != "" {
        args = append(args, models.Where{Column:"info1", Value:_info1, Compare:"="})
    }
    _info2 := c.Get("info2")
    if _info2 != "" {
        args = append(args, models.Where{Column:"info2", Value:_info2, Compare:"="})
    }
    _info3 := c.Get("info3")
    if _info3 != "" {
        args = append(args, models.Where{Column:"info3", Value:_info3, Compare:"="})
    }
    _info4 := c.Get("info4")
    if _info4 != "" {
        args = append(args, models.Where{Column:"info4", Value:_info4, Compare:"="})
    }
    _info5 := c.Get("info5")
    if _info5 != "" {
        args = append(args, models.Where{Column:"info5", Value:_info5, Compare:"="})
    }
    _info6 := c.Get("info6")
    if _info6 != "" {
        args = append(args, models.Where{Column:"info6", Value:_info6, Compare:"="})
    }
    _info7 := c.Get("info7")
    if _info7 != "" {
        args = append(args, models.Where{Column:"info7", Value:_info7, Compare:"="})
    }
    _info8 := c.Get("info8")
    if _info8 != "" {
        args = append(args, models.Where{Column:"info8", Value:_info8, Compare:"="})
    }
    _info9 := c.Get("info9")
    if _info9 != "" {
        args = append(args, models.Where{Column:"info9", Value:_info9, Compare:"="})
    }
    _info10 := c.Get("info10")
    if _info10 != "" {
        args = append(args, models.Where{Column:"info10", Value:_info10, Compare:"="})
    }
    _info11 := c.Get("info11")
    if _info11 != "" {
        args = append(args, models.Where{Column:"info11", Value:_info11, Compare:"="})
    }
    _structure1 := c.Get("structure1")
    if _structure1 != "" {
        args = append(args, models.Where{Column:"structure1", Value:_structure1, Compare:"="})
    }
    _structure2 := c.Get("structure2")
    if _structure2 != "" {
        args = append(args, models.Where{Column:"structure2", Value:_structure2, Compare:"="})
    }
    _structure3 := c.Get("structure3")
    if _structure3 != "" {
        args = append(args, models.Where{Column:"structure3", Value:_structure3, Compare:"="})
    }
    _structure4 := c.Get("structure4")
    if _structure4 != "" {
        args = append(args, models.Where{Column:"structure4", Value:_structure4, Compare:"="})
    }
    _structure5 := c.Get("structure5")
    if _structure5 != "" {
        args = append(args, models.Where{Column:"structure5", Value:_structure5, Compare:"="})
    }
    _structure6 := c.Get("structure6")
    if _structure6 != "" {
        args = append(args, models.Where{Column:"structure6", Value:_structure6, Compare:"="})
    }
    _structure7 := c.Get("structure7")
    if _structure7 != "" {
        args = append(args, models.Where{Column:"structure7", Value:_structure7, Compare:"="})
    }
    _structure8 := c.Get("structure8")
    if _structure8 != "" {
        args = append(args, models.Where{Column:"structure8", Value:_structure8, Compare:"="})
    }
    _structure9 := c.Get("structure9")
    if _structure9 != "" {
        args = append(args, models.Where{Column:"structure9", Value:_structure9, Compare:"="})
    }
    _structure10 := c.Get("structure10")
    if _structure10 != "" {
        args = append(args, models.Where{Column:"structure10", Value:_structure10, Compare:"="})
    }
    _structure11 := c.Get("structure11")
    if _structure11 != "" {
        args = append(args, models.Where{Column:"structure11", Value:_structure11, Compare:"="})
    }
    _structure12 := c.Get("structure12")
    if _structure12 != "" {
        args = append(args, models.Where{Column:"structure12", Value:_structure12, Compare:"="})
    }
    _structure13 := c.Get("structure13")
    if _structure13 != "" {
        args = append(args, models.Where{Column:"structure13", Value:_structure13, Compare:"="})
    }
    _structure14 := c.Get("structure14")
    if _structure14 != "" {
        args = append(args, models.Where{Column:"structure14", Value:_structure14, Compare:"="})
    }
    _reviewcontent1 := c.Get("reviewcontent1")
    if _reviewcontent1 != "" {
        args = append(args, models.Where{Column:"reviewcontent1", Value:_reviewcontent1, Compare:"="})
    }
    _reviewcontent2 := c.Get("reviewcontent2")
    if _reviewcontent2 != "" {
        args = append(args, models.Where{Column:"reviewcontent2", Value:_reviewcontent2, Compare:"="})
    }
    _reviewcontent3 := c.Get("reviewcontent3")
    if _reviewcontent3 != "" {
        args = append(args, models.Where{Column:"reviewcontent3", Value:_reviewcontent3, Compare:"="})
    }
    _reviewcontent4 := c.Get("reviewcontent4")
    if _reviewcontent4 != "" {
        args = append(args, models.Where{Column:"reviewcontent4", Value:_reviewcontent4, Compare:"="})
    }
    _reviewcontent5 := c.Get("reviewcontent5")
    if _reviewcontent5 != "" {
        args = append(args, models.Where{Column:"reviewcontent5", Value:_reviewcontent5, Compare:"="})
    }
    _reviewcontent6 := c.Get("reviewcontent6")
    if _reviewcontent6 != "" {
        args = append(args, models.Where{Column:"reviewcontent6", Value:_reviewcontent6, Compare:"="})
    }
    _reviewcontent7 := c.Get("reviewcontent7")
    if _reviewcontent7 != "" {
        args = append(args, models.Where{Column:"reviewcontent7", Value:_reviewcontent7, Compare:"="})
    }
    _savingprice := c.Geti("savingprice")
    if _savingprice != 0 {
        args = append(args, models.Where{Column:"savingprice", Value:_savingprice, Compare:"="})    
    }
    _price1 := c.Get("price1")
    if _price1 != "" {
        args = append(args, models.Where{Column:"price1", Value:_price1, Compare:"="})
    }
    _price2 := c.Get("price2")
    if _price2 != "" {
        args = append(args, models.Where{Column:"price2", Value:_price2, Compare:"="})
    }
    _price3 := c.Get("price3")
    if _price3 != "" {
        args = append(args, models.Where{Column:"price3", Value:_price3, Compare:"="})
    }
    _price4 := c.Get("price4")
    if _price4 != "" {
        args = append(args, models.Where{Column:"price4", Value:_price4, Compare:"="})
    }
    _price5 := c.Get("price5")
    if _price5 != "" {
        args = append(args, models.Where{Column:"price5", Value:_price5, Compare:"="})
    }
    _reportdate := c.Get("reportdate")
    if _reportdate != "" {
        args = append(args, models.Where{Column:"reportdate", Value:_reportdate, Compare:"="})
    }
    _content1 := c.Get("content1")
    if _content1 != "" {
        args = append(args, models.Where{Column:"content1", Value:_content1, Compare:"="})
    }
    _content2 := c.Get("content2")
    if _content2 != "" {
        args = append(args, models.Where{Column:"content2", Value:_content2, Compare:"="})
    }
    _content3 := c.Get("content3")
    if _content3 != "" {
        args = append(args, models.Where{Column:"content3", Value:_content3, Compare:"="})
    }
    _content4 := c.Get("content4")
    if _content4 != "" {
        args = append(args, models.Where{Column:"content4", Value:_content4, Compare:"="})
    }
    _content5 := c.Get("content5")
    if _content5 != "" {
        args = append(args, models.Where{Column:"content5", Value:_content5, Compare:"="})
    }
    _content6 := c.Get("content6")
    if _content6 != "" {
        args = append(args, models.Where{Column:"content6", Value:_content6, Compare:"="})
    }
    _content7 := c.Get("content7")
    if _content7 != "" {
        args = append(args, models.Where{Column:"content7", Value:_content7, Compare:"="})
    }
    _content8 := c.Get("content8")
    if _content8 != "" {
        args = append(args, models.Where{Column:"content8", Value:_content8, Compare:"="})
    }
    _content9 := c.Get("content9")
    if _content9 != "" {
        args = append(args, models.Where{Column:"content9", Value:_content9, Compare:"="})
    }
    _content10 := c.Get("content10")
    if _content10 != "" {
        args = append(args, models.Where{Column:"content10", Value:_content10, Compare:"="})
    }
    _content11 := c.Get("content11")
    if _content11 != "" {
        args = append(args, models.Where{Column:"content11", Value:_content11, Compare:"="})
    }
    _periodtype := c.Geti("periodtype")
    if _periodtype != 0 {
        args = append(args, models.Where{Column:"periodtype", Value:_periodtype, Compare:"="})    
    }
    _remark := c.Get("remark")
    if _remark != "" {
        args = append(args, models.Where{Column:"remark", Value:_remark, Compare:"="})
    }
    _apt := c.Geti64("apt")
    if _apt != 0 {
        args = append(args, models.Where{Column:"apt", Value:_apt, Compare:"="})    
    }
    _startdate := c.Get("startdate")
    _enddate := c.Get("enddate")
    if _startdate != "" && _enddate != "" {        
        var v [2]string
        v[0] = _startdate
        v[1] = _enddate  
        args = append(args, models.Where{Column:"date", Value:v, Compare:"between"})    
    } else if  _startdate != "" {          
        args = append(args, models.Where{Column:"date", Value:_startdate, Compare:">="})
    } else if  _enddate != "" {          
        args = append(args, models.Where{Column:"date", Value:_enddate, Compare:"<="})            
    }
    

    
    
    if page != 0 && pagesize != 0 {
        args = append(args, models.Paging(page, pagesize))
    }
    
    orderby := c.Get("orderby")
    if orderby == "" {
        if page != 0 && pagesize != 0 {
            orderby = "id desc"
            args = append(args, models.Ordering(orderby))
        }
    } else {
        orderbys := strings.Split(orderby, ",")

        str := ""
        for i, v := range orderbys {
            if i == 0 {
                str += v
            } else {
                if strings.Contains(v, "_") {                   
                    str += ", " + strings.Trim(v, " ")
                } else {
                    str += ", r_" + strings.Trim(v, " ")                
                }
            }
        }
        
        args = append(args, models.Ordering(str))
    }
    
	items := manager.Find(args)
	c.Set("items", items)

    if page == 1 {
       total := manager.Count(args)
	   c.Set("total", total)
    }
}





