package rest


import (
	"repair/controllers"
	"repair/models"

    "strings"
)

type PeriodickeepController struct {
	controllers.Controller
}



func (c *PeriodickeepController) GetByPeriodic(periodic int64) *models.Periodickeep {
    
    conn := c.NewConnection()

	_manager := models.NewPeriodickeepManager(conn)
    
    item := _manager.GetByPeriodic(periodic)
    
    c.Set("item", item)
    
    
    
    return item
    
}


func (c *PeriodickeepController) Insert(item *models.Periodickeep) {
    
    
	conn := c.NewConnection()
    
	manager := models.NewPeriodickeepManager(conn)
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

func (c *PeriodickeepController) Insertbatch(item *[]models.Periodickeep) {  
    if item == nil || len(*item) == 0 {
        return
    }

    rows := len(*item)
    
    
    
	conn := c.NewConnection()
    
	manager := models.NewPeriodickeepManager(conn)

    for i := 0; i < rows; i++ {
	    err := manager.Insert(&((*item)[i]))
        if err != nil {
            c.Set("code", "error")    
            c.Set("error", err)
            return
        }
    }
}

func (c *PeriodickeepController) Update(item *models.Periodickeep) {
    
    
	conn := c.NewConnection()

	manager := models.NewPeriodickeepManager(conn)
    err := manager.Update(item)
    if err != nil {
        c.Set("code", "error")    
        c.Set("error", err)
        return
    }
}

func (c *PeriodickeepController) Delete(item *models.Periodickeep) {
    
    
    conn := c.NewConnection()

	manager := models.NewPeriodickeepManager(conn)

    
	err := manager.Delete(item.Id)
    if err != nil {
        c.Set("code", "error")    
        c.Set("error", err)
    }
}

func (c *PeriodickeepController) Deletebatch(item *[]models.Periodickeep) {
    
    
    conn := c.NewConnection()

	manager := models.NewPeriodickeepManager(conn)

    for _, v := range *item {
        
    
	    err := manager.Delete(v.Id)
        if err != nil {
            c.Set("code", "error")    
            c.Set("error", err)
            return
        }
    }
}

func (c *PeriodickeepController) Count() {
    
    
	conn := c.NewConnection()

	manager := models.NewPeriodickeepManager(conn)

    var args []interface{}
    
    _status1 := c.Geti("status1")
    if _status1 != 0 {
        args = append(args, models.Where{Column:"status1", Value:_status1, Compare:"="})    
    }
    _status2 := c.Geti("status2")
    if _status2 != 0 {
        args = append(args, models.Where{Column:"status2", Value:_status2, Compare:"="})    
    }
    _status3 := c.Geti("status3")
    if _status3 != 0 {
        args = append(args, models.Where{Column:"status3", Value:_status3, Compare:"="})    
    }
    _status4 := c.Geti("status4")
    if _status4 != 0 {
        args = append(args, models.Where{Column:"status4", Value:_status4, Compare:"="})    
    }
    _status5 := c.Geti("status5")
    if _status5 != 0 {
        args = append(args, models.Where{Column:"status5", Value:_status5, Compare:"="})    
    }
    _status6 := c.Geti("status6")
    if _status6 != 0 {
        args = append(args, models.Where{Column:"status6", Value:_status6, Compare:"="})    
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
    _remark1 := c.Get("remark1")
    if _remark1 != "" {
        args = append(args, models.Where{Column:"remark1", Value:_remark1, Compare:"="})
    }
    _remark2 := c.Get("remark2")
    if _remark2 != "" {
        args = append(args, models.Where{Column:"remark2", Value:_remark2, Compare:"="})
    }
    _remark3 := c.Get("remark3")
    if _remark3 != "" {
        args = append(args, models.Where{Column:"remark3", Value:_remark3, Compare:"="})
    }
    _remark4 := c.Get("remark4")
    if _remark4 != "" {
        args = append(args, models.Where{Column:"remark4", Value:_remark4, Compare:"="})
    }
    _remark5 := c.Get("remark5")
    if _remark5 != "" {
        args = append(args, models.Where{Column:"remark5", Value:_remark5, Compare:"="})
    }
    _remark6 := c.Get("remark6")
    if _remark6 != "" {
        args = append(args, models.Where{Column:"remark6", Value:_remark6, Compare:"="})
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


func (c *PeriodickeepController) Read(id int64) {
    
    
	conn := c.NewConnection()

	manager := models.NewPeriodickeepManager(conn)
	item := manager.Get(id)

    
    
    c.Set("item", item)
}

func (c *PeriodickeepController) Index(page int, pagesize int) {
    
    
	conn := c.NewConnection()

	manager := models.NewPeriodickeepManager(conn)

    var args []interface{}
    
    _status1 := c.Geti("status1")
    if _status1 != 0 {
        args = append(args, models.Where{Column:"status1", Value:_status1, Compare:"="})    
    }
    _status2 := c.Geti("status2")
    if _status2 != 0 {
        args = append(args, models.Where{Column:"status2", Value:_status2, Compare:"="})    
    }
    _status3 := c.Geti("status3")
    if _status3 != 0 {
        args = append(args, models.Where{Column:"status3", Value:_status3, Compare:"="})    
    }
    _status4 := c.Geti("status4")
    if _status4 != 0 {
        args = append(args, models.Where{Column:"status4", Value:_status4, Compare:"="})    
    }
    _status5 := c.Geti("status5")
    if _status5 != 0 {
        args = append(args, models.Where{Column:"status5", Value:_status5, Compare:"="})    
    }
    _status6 := c.Geti("status6")
    if _status6 != 0 {
        args = append(args, models.Where{Column:"status6", Value:_status6, Compare:"="})    
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
    _remark1 := c.Get("remark1")
    if _remark1 != "" {
        args = append(args, models.Where{Column:"remark1", Value:_remark1, Compare:"="})
    }
    _remark2 := c.Get("remark2")
    if _remark2 != "" {
        args = append(args, models.Where{Column:"remark2", Value:_remark2, Compare:"="})
    }
    _remark3 := c.Get("remark3")
    if _remark3 != "" {
        args = append(args, models.Where{Column:"remark3", Value:_remark3, Compare:"="})
    }
    _remark4 := c.Get("remark4")
    if _remark4 != "" {
        args = append(args, models.Where{Column:"remark4", Value:_remark4, Compare:"="})
    }
    _remark5 := c.Get("remark5")
    if _remark5 != "" {
        args = append(args, models.Where{Column:"remark5", Value:_remark5, Compare:"="})
    }
    _remark6 := c.Get("remark6")
    if _remark6 != "" {
        args = append(args, models.Where{Column:"remark6", Value:_remark6, Compare:"="})
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
                    str += ", pk_" + strings.Trim(v, " ")                
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





