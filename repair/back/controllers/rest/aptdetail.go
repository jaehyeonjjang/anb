package rest


import (
	"repair/controllers"
	"repair/models"

    "strings"
)

type AptdetailController struct {
	controllers.Controller
}



func (c *AptdetailController) Insert(item *models.Aptdetail) {
    
    
	conn := c.NewConnection()
    
	manager := models.NewAptdetailManager(conn)
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

func (c *AptdetailController) Insertbatch(item *[]models.Aptdetail) {  
    if item == nil || len(*item) == 0 {
        return
    }

    rows := len(*item)
    
    
    
	conn := c.NewConnection()
    
	manager := models.NewAptdetailManager(conn)

    for i := 0; i < rows; i++ {
	    err := manager.Insert(&((*item)[i]))
        if err != nil {
            c.Set("code", "error")    
            c.Set("error", err)
            return
        }
    }
}

func (c *AptdetailController) Update(item *models.Aptdetail) {
    
    
	conn := c.NewConnection()

	manager := models.NewAptdetailManager(conn)
    err := manager.Update(item)
    if err != nil {
        c.Set("code", "error")    
        c.Set("error", err)
        return
    }
}

func (c *AptdetailController) Delete(item *models.Aptdetail) {
    
    
    conn := c.NewConnection()

	manager := models.NewAptdetailManager(conn)

    
	err := manager.Delete(item.Id)
    if err != nil {
        c.Set("code", "error")    
        c.Set("error", err)
    }
}

func (c *AptdetailController) Deletebatch(item *[]models.Aptdetail) {
    
    
    conn := c.NewConnection()

	manager := models.NewAptdetailManager(conn)

    for _, v := range *item {
        
    
	    err := manager.Delete(v.Id)
        if err != nil {
            c.Set("code", "error")    
            c.Set("error", err)
            return
        }
    }
}

func (c *AptdetailController) Count() {
    
    
	conn := c.NewConnection()

	manager := models.NewAptdetailManager(conn)

    var args []interface{}
    
    _current1 := c.Get("current1")
    if _current1 != "" {
        args = append(args, models.Where{Column:"current1", Value:_current1, Compare:"="})
    }
    _current2 := c.Get("current2")
    if _current2 != "" {
        args = append(args, models.Where{Column:"current2", Value:_current2, Compare:"="})
    }
    _current3 := c.Get("current3")
    if _current3 != "" {
        args = append(args, models.Where{Column:"current3", Value:_current3, Compare:"="})
    }
    _current4 := c.Get("current4")
    if _current4 != "" {
        args = append(args, models.Where{Column:"current4", Value:_current4, Compare:"="})
    }
    _current5 := c.Get("current5")
    if _current5 != "" {
        args = append(args, models.Where{Column:"current5", Value:_current5, Compare:"="})
    }
    _current6 := c.Geti("current6")
    if _current6 != 0 {
        args = append(args, models.Where{Column:"current6", Value:_current6, Compare:"="})    
    }
    _current7 := c.Geti("current7")
    if _current7 != 0 {
        args = append(args, models.Where{Column:"current7", Value:_current7, Compare:"="})    
    }
    _current8 := c.Geti("current8")
    if _current8 != 0 {
        args = append(args, models.Where{Column:"current8", Value:_current8, Compare:"="})    
    }
    _current9 := c.Get("current9")
    if _current9 != "" {
        args = append(args, models.Where{Column:"current9", Value:_current9, Compare:"="})
    }
    _current10 := c.Get("current10")
    if _current10 != "" {
        args = append(args, models.Where{Column:"current10", Value:_current10, Compare:"="})
    }
    _current11 := c.Geti("current11")
    if _current11 != 0 {
        args = append(args, models.Where{Column:"current11", Value:_current11, Compare:"="})    
    }
    _current12 := c.Geti("current12")
    if _current12 != 0 {
        args = append(args, models.Where{Column:"current12", Value:_current12, Compare:"="})    
    }
    _current13 := c.Geti("current13")
    if _current13 != 0 {
        args = append(args, models.Where{Column:"current13", Value:_current13, Compare:"="})    
    }
    _current14 := c.Geti("current14")
    if _current14 != 0 {
        args = append(args, models.Where{Column:"current14", Value:_current14, Compare:"="})    
    }
    _current15 := c.Get("current15")
    if _current15 != "" {
        args = append(args, models.Where{Column:"current15", Value:_current15, Compare:"="})
    }
    _current16 := c.Get("current16")
    if _current16 != "" {
        args = append(args, models.Where{Column:"current16", Value:_current16, Compare:"="})
    }
    _current17 := c.Get("current17")
    if _current17 != "" {
        args = append(args, models.Where{Column:"current17", Value:_current17, Compare:"="})
    }
    _current18 := c.Get("current18")
    if _current18 != "" {
        args = append(args, models.Where{Column:"current18", Value:_current18, Compare:"="})
    }
    _current19 := c.Get("current19")
    if _current19 != "" {
        args = append(args, models.Where{Column:"current19", Value:_current19, Compare:"="})
    }
    _current20 := c.Get("current20")
    if _current20 != "" {
        args = append(args, models.Where{Column:"current20", Value:_current20, Compare:"="})
    }
    _current21 := c.Get("current21")
    if _current21 != "" {
        args = append(args, models.Where{Column:"current21", Value:_current21, Compare:"="})
    }
    _current22 := c.Geti("current22")
    if _current22 != 0 {
        args = append(args, models.Where{Column:"current22", Value:_current22, Compare:"="})    
    }
    _current23 := c.Geti("current23")
    if _current23 != 0 {
        args = append(args, models.Where{Column:"current23", Value:_current23, Compare:"="})    
    }
    _outline1 := c.Geti("outline1")
    if _outline1 != 0 {
        args = append(args, models.Where{Column:"outline1", Value:_outline1, Compare:"="})    
    }
    _outline2 := c.Geti("outline2")
    if _outline2 != 0 {
        args = append(args, models.Where{Column:"outline2", Value:_outline2, Compare:"="})    
    }
    _outline3 := c.Geti("outline3")
    if _outline3 != 0 {
        args = append(args, models.Where{Column:"outline3", Value:_outline3, Compare:"="})    
    }
    _outline4 := c.Geti("outline4")
    if _outline4 != 0 {
        args = append(args, models.Where{Column:"outline4", Value:_outline4, Compare:"="})    
    }
    _outline5 := c.Geti("outline5")
    if _outline5 != 0 {
        args = append(args, models.Where{Column:"outline5", Value:_outline5, Compare:"="})    
    }
    _outline6 := c.Geti("outline6")
    if _outline6 != 0 {
        args = append(args, models.Where{Column:"outline6", Value:_outline6, Compare:"="})    
    }
    _outline7 := c.Geti("outline7")
    if _outline7 != 0 {
        args = append(args, models.Where{Column:"outline7", Value:_outline7, Compare:"="})    
    }
    _outline8 := c.Geti("outline8")
    if _outline8 != 0 {
        args = append(args, models.Where{Column:"outline8", Value:_outline8, Compare:"="})    
    }
    _outline9 := c.Geti("outline9")
    if _outline9 != 0 {
        args = append(args, models.Where{Column:"outline9", Value:_outline9, Compare:"="})    
    }
    _record1 := c.Get("record1")
    if _record1 != "" {
        args = append(args, models.Where{Column:"record1", Value:_record1, Compare:"="})
    }
    _record2 := c.Get("record2")
    if _record2 != "" {
        args = append(args, models.Where{Column:"record2", Value:_record2, Compare:"="})
    }
    _record3 := c.Get("record3")
    if _record3 != "" {
        args = append(args, models.Where{Column:"record3", Value:_record3, Compare:"="})
    }
    _startrecord4 := c.Get("startrecord4")
    _endrecord4 := c.Get("endrecord4")

    if _startrecord4 != "" && _endrecord4 != "" {        
        var v [2]string
        v[0] = _startrecord4
        v[1] = _endrecord4  
        args = append(args, models.Where{Column:"record4", Value:v, Compare:"between"})    
    } else if  _startrecord4 != "" {          
        args = append(args, models.Where{Column:"record4", Value:_startrecord4, Compare:">="})
    } else if  _endrecord4 != "" {          
        args = append(args, models.Where{Column:"record4", Value:_endrecord4, Compare:"<="})            
    }
    _startrecord5 := c.Get("startrecord5")
    _endrecord5 := c.Get("endrecord5")

    if _startrecord5 != "" && _endrecord5 != "" {        
        var v [2]string
        v[0] = _startrecord5
        v[1] = _endrecord5  
        args = append(args, models.Where{Column:"record5", Value:v, Compare:"between"})    
    } else if  _startrecord5 != "" {          
        args = append(args, models.Where{Column:"record5", Value:_startrecord5, Compare:">="})
    } else if  _endrecord5 != "" {          
        args = append(args, models.Where{Column:"record5", Value:_endrecord5, Compare:"<="})            
    }
    _deligate := c.Get("deligate")
    if _deligate != "" {
        args = append(args, models.Where{Column:"deligate", Value:_deligate, Compare:"="})
    }
    _facilitydivision := c.Geti("facilitydivision")
    if _facilitydivision != 0 {
        args = append(args, models.Where{Column:"facilitydivision", Value:_facilitydivision, Compare:"="})    
    }
    _facilitytype := c.Geti("facilitytype")
    if _facilitytype != 0 {
        args = append(args, models.Where{Column:"facilitytype", Value:_facilitytype, Compare:"="})    
    }
    _facilitycategory := c.Geti("facilitycategory")
    if _facilitycategory != 0 {
        args = append(args, models.Where{Column:"facilitycategory", Value:_facilitycategory, Compare:"="})    
    }
    _struct1 := c.Get("struct1")
    if _struct1 != "" {
        args = append(args, models.Where{Column:"struct1", Value:_struct1, Compare:"="})
    }
    _struct2 := c.Get("struct2")
    if _struct2 != "" {
        args = append(args, models.Where{Column:"struct2", Value:_struct2, Compare:"="})
    }
    _struct3 := c.Get("struct3")
    if _struct3 != "" {
        args = append(args, models.Where{Column:"struct3", Value:_struct3, Compare:"="})
    }
    _struct4 := c.Get("struct4")
    if _struct4 != "" {
        args = append(args, models.Where{Column:"struct4", Value:_struct4, Compare:"="})
    }
    _struct5 := c.Geti("struct5")
    if _struct5 != 0 {
        args = append(args, models.Where{Column:"struct5", Value:_struct5, Compare:"="})    
    }
    _struct6 := c.Get("struct6")
    if _struct6 != "" {
        args = append(args, models.Where{Column:"struct6", Value:_struct6, Compare:"="})
    }
    _struct7 := c.Get("struct7")
    if _struct7 != "" {
        args = append(args, models.Where{Column:"struct7", Value:_struct7, Compare:"="})
    }
    _struct8 := c.Get("struct8")
    if _struct8 != "" {
        args = append(args, models.Where{Column:"struct8", Value:_struct8, Compare:"="})
    }
    _struct9 := c.Get("struct9")
    if _struct9 != "" {
        args = append(args, models.Where{Column:"struct9", Value:_struct9, Compare:"="})
    }
    _struct10 := c.Get("struct10")
    if _struct10 != "" {
        args = append(args, models.Where{Column:"struct10", Value:_struct10, Compare:"="})
    }
    _struct11 := c.Get("struct11")
    if _struct11 != "" {
        args = append(args, models.Where{Column:"struct11", Value:_struct11, Compare:"="})
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


func (c *AptdetailController) Read(id int64) {
    
    
	conn := c.NewConnection()

	manager := models.NewAptdetailManager(conn)
	item := manager.Get(id)

    
    
    c.Set("item", item)
}

func (c *AptdetailController) Index(page int, pagesize int) {
    
    
	conn := c.NewConnection()

	manager := models.NewAptdetailManager(conn)

    var args []interface{}
    
    _current1 := c.Get("current1")
    if _current1 != "" {
        args = append(args, models.Where{Column:"current1", Value:_current1, Compare:"="})
    }
    _current2 := c.Get("current2")
    if _current2 != "" {
        args = append(args, models.Where{Column:"current2", Value:_current2, Compare:"="})
    }
    _current3 := c.Get("current3")
    if _current3 != "" {
        args = append(args, models.Where{Column:"current3", Value:_current3, Compare:"="})
    }
    _current4 := c.Get("current4")
    if _current4 != "" {
        args = append(args, models.Where{Column:"current4", Value:_current4, Compare:"="})
    }
    _current5 := c.Get("current5")
    if _current5 != "" {
        args = append(args, models.Where{Column:"current5", Value:_current5, Compare:"="})
    }
    _current6 := c.Geti("current6")
    if _current6 != 0 {
        args = append(args, models.Where{Column:"current6", Value:_current6, Compare:"="})    
    }
    _current7 := c.Geti("current7")
    if _current7 != 0 {
        args = append(args, models.Where{Column:"current7", Value:_current7, Compare:"="})    
    }
    _current8 := c.Geti("current8")
    if _current8 != 0 {
        args = append(args, models.Where{Column:"current8", Value:_current8, Compare:"="})    
    }
    _current9 := c.Get("current9")
    if _current9 != "" {
        args = append(args, models.Where{Column:"current9", Value:_current9, Compare:"="})
    }
    _current10 := c.Get("current10")
    if _current10 != "" {
        args = append(args, models.Where{Column:"current10", Value:_current10, Compare:"="})
    }
    _current11 := c.Geti("current11")
    if _current11 != 0 {
        args = append(args, models.Where{Column:"current11", Value:_current11, Compare:"="})    
    }
    _current12 := c.Geti("current12")
    if _current12 != 0 {
        args = append(args, models.Where{Column:"current12", Value:_current12, Compare:"="})    
    }
    _current13 := c.Geti("current13")
    if _current13 != 0 {
        args = append(args, models.Where{Column:"current13", Value:_current13, Compare:"="})    
    }
    _current14 := c.Geti("current14")
    if _current14 != 0 {
        args = append(args, models.Where{Column:"current14", Value:_current14, Compare:"="})    
    }
    _current15 := c.Get("current15")
    if _current15 != "" {
        args = append(args, models.Where{Column:"current15", Value:_current15, Compare:"="})
    }
    _current16 := c.Get("current16")
    if _current16 != "" {
        args = append(args, models.Where{Column:"current16", Value:_current16, Compare:"="})
    }
    _current17 := c.Get("current17")
    if _current17 != "" {
        args = append(args, models.Where{Column:"current17", Value:_current17, Compare:"="})
    }
    _current18 := c.Get("current18")
    if _current18 != "" {
        args = append(args, models.Where{Column:"current18", Value:_current18, Compare:"="})
    }
    _current19 := c.Get("current19")
    if _current19 != "" {
        args = append(args, models.Where{Column:"current19", Value:_current19, Compare:"="})
    }
    _current20 := c.Get("current20")
    if _current20 != "" {
        args = append(args, models.Where{Column:"current20", Value:_current20, Compare:"="})
    }
    _current21 := c.Get("current21")
    if _current21 != "" {
        args = append(args, models.Where{Column:"current21", Value:_current21, Compare:"="})
    }
    _current22 := c.Geti("current22")
    if _current22 != 0 {
        args = append(args, models.Where{Column:"current22", Value:_current22, Compare:"="})    
    }
    _current23 := c.Geti("current23")
    if _current23 != 0 {
        args = append(args, models.Where{Column:"current23", Value:_current23, Compare:"="})    
    }
    _outline1 := c.Geti("outline1")
    if _outline1 != 0 {
        args = append(args, models.Where{Column:"outline1", Value:_outline1, Compare:"="})    
    }
    _outline2 := c.Geti("outline2")
    if _outline2 != 0 {
        args = append(args, models.Where{Column:"outline2", Value:_outline2, Compare:"="})    
    }
    _outline3 := c.Geti("outline3")
    if _outline3 != 0 {
        args = append(args, models.Where{Column:"outline3", Value:_outline3, Compare:"="})    
    }
    _outline4 := c.Geti("outline4")
    if _outline4 != 0 {
        args = append(args, models.Where{Column:"outline4", Value:_outline4, Compare:"="})    
    }
    _outline5 := c.Geti("outline5")
    if _outline5 != 0 {
        args = append(args, models.Where{Column:"outline5", Value:_outline5, Compare:"="})    
    }
    _outline6 := c.Geti("outline6")
    if _outline6 != 0 {
        args = append(args, models.Where{Column:"outline6", Value:_outline6, Compare:"="})    
    }
    _outline7 := c.Geti("outline7")
    if _outline7 != 0 {
        args = append(args, models.Where{Column:"outline7", Value:_outline7, Compare:"="})    
    }
    _outline8 := c.Geti("outline8")
    if _outline8 != 0 {
        args = append(args, models.Where{Column:"outline8", Value:_outline8, Compare:"="})    
    }
    _outline9 := c.Geti("outline9")
    if _outline9 != 0 {
        args = append(args, models.Where{Column:"outline9", Value:_outline9, Compare:"="})    
    }
    _record1 := c.Get("record1")
    if _record1 != "" {
        args = append(args, models.Where{Column:"record1", Value:_record1, Compare:"="})
    }
    _record2 := c.Get("record2")
    if _record2 != "" {
        args = append(args, models.Where{Column:"record2", Value:_record2, Compare:"="})
    }
    _record3 := c.Get("record3")
    if _record3 != "" {
        args = append(args, models.Where{Column:"record3", Value:_record3, Compare:"="})
    }
    _startrecord4 := c.Get("startrecord4")
    _endrecord4 := c.Get("endrecord4")
    if _startrecord4 != "" && _endrecord4 != "" {        
        var v [2]string
        v[0] = _startrecord4
        v[1] = _endrecord4  
        args = append(args, models.Where{Column:"record4", Value:v, Compare:"between"})    
    } else if  _startrecord4 != "" {          
        args = append(args, models.Where{Column:"record4", Value:_startrecord4, Compare:">="})
    } else if  _endrecord4 != "" {          
        args = append(args, models.Where{Column:"record4", Value:_endrecord4, Compare:"<="})            
    }
    _startrecord5 := c.Get("startrecord5")
    _endrecord5 := c.Get("endrecord5")
    if _startrecord5 != "" && _endrecord5 != "" {        
        var v [2]string
        v[0] = _startrecord5
        v[1] = _endrecord5  
        args = append(args, models.Where{Column:"record5", Value:v, Compare:"between"})    
    } else if  _startrecord5 != "" {          
        args = append(args, models.Where{Column:"record5", Value:_startrecord5, Compare:">="})
    } else if  _endrecord5 != "" {          
        args = append(args, models.Where{Column:"record5", Value:_endrecord5, Compare:"<="})            
    }
    _deligate := c.Get("deligate")
    if _deligate != "" {
        args = append(args, models.Where{Column:"deligate", Value:_deligate, Compare:"="})
    }
    _facilitydivision := c.Geti("facilitydivision")
    if _facilitydivision != 0 {
        args = append(args, models.Where{Column:"facilitydivision", Value:_facilitydivision, Compare:"="})    
    }
    _facilitytype := c.Geti("facilitytype")
    if _facilitytype != 0 {
        args = append(args, models.Where{Column:"facilitytype", Value:_facilitytype, Compare:"="})    
    }
    _facilitycategory := c.Geti("facilitycategory")
    if _facilitycategory != 0 {
        args = append(args, models.Where{Column:"facilitycategory", Value:_facilitycategory, Compare:"="})    
    }
    _struct1 := c.Get("struct1")
    if _struct1 != "" {
        args = append(args, models.Where{Column:"struct1", Value:_struct1, Compare:"="})
    }
    _struct2 := c.Get("struct2")
    if _struct2 != "" {
        args = append(args, models.Where{Column:"struct2", Value:_struct2, Compare:"="})
    }
    _struct3 := c.Get("struct3")
    if _struct3 != "" {
        args = append(args, models.Where{Column:"struct3", Value:_struct3, Compare:"="})
    }
    _struct4 := c.Get("struct4")
    if _struct4 != "" {
        args = append(args, models.Where{Column:"struct4", Value:_struct4, Compare:"="})
    }
    _struct5 := c.Geti("struct5")
    if _struct5 != 0 {
        args = append(args, models.Where{Column:"struct5", Value:_struct5, Compare:"="})    
    }
    _struct6 := c.Get("struct6")
    if _struct6 != "" {
        args = append(args, models.Where{Column:"struct6", Value:_struct6, Compare:"="})
    }
    _struct7 := c.Get("struct7")
    if _struct7 != "" {
        args = append(args, models.Where{Column:"struct7", Value:_struct7, Compare:"="})
    }
    _struct8 := c.Get("struct8")
    if _struct8 != "" {
        args = append(args, models.Where{Column:"struct8", Value:_struct8, Compare:"="})
    }
    _struct9 := c.Get("struct9")
    if _struct9 != "" {
        args = append(args, models.Where{Column:"struct9", Value:_struct9, Compare:"="})
    }
    _struct10 := c.Get("struct10")
    if _struct10 != "" {
        args = append(args, models.Where{Column:"struct10", Value:_struct10, Compare:"="})
    }
    _struct11 := c.Get("struct11")
    if _struct11 != "" {
        args = append(args, models.Where{Column:"struct11", Value:_struct11, Compare:"="})
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
                    str += ", ad_" + strings.Trim(v, " ")                
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





