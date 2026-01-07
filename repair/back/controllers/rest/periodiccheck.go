package rest


import (
	"repair/controllers"
	"repair/models"

    "strings"
)

type PeriodiccheckController struct {
	controllers.Controller
}



func (c *PeriodiccheckController) Insert(item *models.Periodiccheck) {
    
    
	conn := c.NewConnection()
    
	manager := models.NewPeriodiccheckManager(conn)
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

func (c *PeriodiccheckController) Insertbatch(item *[]models.Periodiccheck) {  
    if item == nil || len(*item) == 0 {
        return
    }

    rows := len(*item)
    
    
    
	conn := c.NewConnection()
    
	manager := models.NewPeriodiccheckManager(conn)

    for i := 0; i < rows; i++ {
	    err := manager.Insert(&((*item)[i]))
        if err != nil {
            c.Set("code", "error")    
            c.Set("error", err)
            return
        }
    }
}

func (c *PeriodiccheckController) Update(item *models.Periodiccheck) {
    
    
	conn := c.NewConnection()

	manager := models.NewPeriodiccheckManager(conn)
    err := manager.Update(item)
    if err != nil {
        c.Set("code", "error")    
        c.Set("error", err)
        return
    }
}

func (c *PeriodiccheckController) Delete(item *models.Periodiccheck) {
    
    
    conn := c.NewConnection()

	manager := models.NewPeriodiccheckManager(conn)

    
	err := manager.Delete(item.Id)
    if err != nil {
        c.Set("code", "error")    
        c.Set("error", err)
    }
}

func (c *PeriodiccheckController) Deletebatch(item *[]models.Periodiccheck) {
    
    
    conn := c.NewConnection()

	manager := models.NewPeriodiccheckManager(conn)

    for _, v := range *item {
        
    
	    err := manager.Delete(v.Id)
        if err != nil {
            c.Set("code", "error")    
            c.Set("error", err)
            return
        }
    }
}

func (c *PeriodiccheckController) Count() {
    
    
	conn := c.NewConnection()

	manager := models.NewPeriodiccheckManager(conn)

    var args []interface{}
    
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
    _content12 := c.Get("content12")
    if _content12 != "" {
        args = append(args, models.Where{Column:"content12", Value:_content12, Compare:"="})
    }
    _content13 := c.Get("content13")
    if _content13 != "" {
        args = append(args, models.Where{Column:"content13", Value:_content13, Compare:"="})
    }
    _content14 := c.Get("content14")
    if _content14 != "" {
        args = append(args, models.Where{Column:"content14", Value:_content14, Compare:"="})
    }
    _content15 := c.Get("content15")
    if _content15 != "" {
        args = append(args, models.Where{Column:"content15", Value:_content15, Compare:"="})
    }
    _content16 := c.Get("content16")
    if _content16 != "" {
        args = append(args, models.Where{Column:"content16", Value:_content16, Compare:"="})
    }
    _use1 := c.Geti("use1")
    if _use1 != 0 {
        args = append(args, models.Where{Column:"use1", Value:_use1, Compare:"="})    
    }
    _use2 := c.Geti("use2")
    if _use2 != 0 {
        args = append(args, models.Where{Column:"use2", Value:_use2, Compare:"="})    
    }
    _use3 := c.Geti("use3")
    if _use3 != 0 {
        args = append(args, models.Where{Column:"use3", Value:_use3, Compare:"="})    
    }
    _use4 := c.Geti("use4")
    if _use4 != 0 {
        args = append(args, models.Where{Column:"use4", Value:_use4, Compare:"="})    
    }
    _need1 := c.Geti("need1")
    if _need1 != 0 {
        args = append(args, models.Where{Column:"need1", Value:_need1, Compare:"="})    
    }
    _need2 := c.Geti("need2")
    if _need2 != 0 {
        args = append(args, models.Where{Column:"need2", Value:_need2, Compare:"="})    
    }
    _need3 := c.Geti("need3")
    if _need3 != 0 {
        args = append(args, models.Where{Column:"need3", Value:_need3, Compare:"="})    
    }
    _need4 := c.Geti("need4")
    if _need4 != 0 {
        args = append(args, models.Where{Column:"need4", Value:_need4, Compare:"="})    
    }
    _aptdong := c.Geti64("aptdong")
    if _aptdong != 0 {
        args = append(args, models.Where{Column:"aptdong", Value:_aptdong, Compare:"="})    
    }
    _periodic := c.Geti64("periodic")
    if _periodic != 0 {
        args = append(args, models.Where{Column:"periodic", Value:_periodic, Compare:"="})    
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


func (c *PeriodiccheckController) Read(id int64) {
    
    
	conn := c.NewConnection()

	manager := models.NewPeriodiccheckManager(conn)
	item := manager.Get(id)

    
    
    c.Set("item", item)
}

func (c *PeriodiccheckController) Index(page int, pagesize int) {
    
    
	conn := c.NewConnection()

	manager := models.NewPeriodiccheckManager(conn)

    var args []interface{}
    
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
    _content12 := c.Get("content12")
    if _content12 != "" {
        args = append(args, models.Where{Column:"content12", Value:_content12, Compare:"="})
    }
    _content13 := c.Get("content13")
    if _content13 != "" {
        args = append(args, models.Where{Column:"content13", Value:_content13, Compare:"="})
    }
    _content14 := c.Get("content14")
    if _content14 != "" {
        args = append(args, models.Where{Column:"content14", Value:_content14, Compare:"="})
    }
    _content15 := c.Get("content15")
    if _content15 != "" {
        args = append(args, models.Where{Column:"content15", Value:_content15, Compare:"="})
    }
    _content16 := c.Get("content16")
    if _content16 != "" {
        args = append(args, models.Where{Column:"content16", Value:_content16, Compare:"="})
    }
    _use1 := c.Geti("use1")
    if _use1 != 0 {
        args = append(args, models.Where{Column:"use1", Value:_use1, Compare:"="})    
    }
    _use2 := c.Geti("use2")
    if _use2 != 0 {
        args = append(args, models.Where{Column:"use2", Value:_use2, Compare:"="})    
    }
    _use3 := c.Geti("use3")
    if _use3 != 0 {
        args = append(args, models.Where{Column:"use3", Value:_use3, Compare:"="})    
    }
    _use4 := c.Geti("use4")
    if _use4 != 0 {
        args = append(args, models.Where{Column:"use4", Value:_use4, Compare:"="})    
    }
    _need1 := c.Geti("need1")
    if _need1 != 0 {
        args = append(args, models.Where{Column:"need1", Value:_need1, Compare:"="})    
    }
    _need2 := c.Geti("need2")
    if _need2 != 0 {
        args = append(args, models.Where{Column:"need2", Value:_need2, Compare:"="})    
    }
    _need3 := c.Geti("need3")
    if _need3 != 0 {
        args = append(args, models.Where{Column:"need3", Value:_need3, Compare:"="})    
    }
    _need4 := c.Geti("need4")
    if _need4 != 0 {
        args = append(args, models.Where{Column:"need4", Value:_need4, Compare:"="})    
    }
    _aptdong := c.Geti64("aptdong")
    if _aptdong != 0 {
        args = append(args, models.Where{Column:"aptdong", Value:_aptdong, Compare:"="})    
    }
    _periodic := c.Geti64("periodic")
    if _periodic != 0 {
        args = append(args, models.Where{Column:"periodic", Value:_periodic, Compare:"="})    
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
                    str += ", pc_" + strings.Trim(v, " ")                
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





