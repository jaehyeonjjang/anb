package rest


import (
	"repair/controllers"
	"repair/models"

    "strings"
)

type PeriodicouterwallController struct {
	controllers.Controller
}



func (c *PeriodicouterwallController) GetByPeriodic(periodic int64) *models.Periodicouterwall {
    
    conn := c.NewConnection()

	_manager := models.NewPeriodicouterwallManager(conn)
    
    item := _manager.GetByPeriodic(periodic)
    
    c.Set("item", item)
    
    
    
    return item
    
}


func (c *PeriodicouterwallController) Insert(item *models.Periodicouterwall) {
    
    
	conn := c.NewConnection()
    
	manager := models.NewPeriodicouterwallManager(conn)
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

func (c *PeriodicouterwallController) Insertbatch(item *[]models.Periodicouterwall) {  
    if item == nil || len(*item) == 0 {
        return
    }

    rows := len(*item)
    
    
    
	conn := c.NewConnection()
    
	manager := models.NewPeriodicouterwallManager(conn)

    for i := 0; i < rows; i++ {
	    err := manager.Insert(&((*item)[i]))
        if err != nil {
            c.Set("code", "error")    
            c.Set("error", err)
            return
        }
    }
}

func (c *PeriodicouterwallController) Update(item *models.Periodicouterwall) {
    
    
	conn := c.NewConnection()

	manager := models.NewPeriodicouterwallManager(conn)
    err := manager.Update(item)
    if err != nil {
        c.Set("code", "error")    
        c.Set("error", err)
        return
    }
}

func (c *PeriodicouterwallController) Delete(item *models.Periodicouterwall) {
    
    
    conn := c.NewConnection()

	manager := models.NewPeriodicouterwallManager(conn)

    
	err := manager.Delete(item.Id)
    if err != nil {
        c.Set("code", "error")    
        c.Set("error", err)
    }
}

func (c *PeriodicouterwallController) Deletebatch(item *[]models.Periodicouterwall) {
    
    
    conn := c.NewConnection()

	manager := models.NewPeriodicouterwallManager(conn)

    for _, v := range *item {
        
    
	    err := manager.Delete(v.Id)
        if err != nil {
            c.Set("code", "error")    
            c.Set("error", err)
            return
        }
    }
}

func (c *PeriodicouterwallController) Count() {
    
    
	conn := c.NewConnection()

	manager := models.NewPeriodicouterwallManager(conn)

    var args []interface{}
    
    _result1 := c.Get("result1")
    if _result1 != "" {
        args = append(args, models.Where{Column:"result1", Value:_result1, Compare:"="})
    }
    _result2 := c.Get("result2")
    if _result2 != "" {
        args = append(args, models.Where{Column:"result2", Value:_result2, Compare:"="})
    }
    _result3 := c.Get("result3")
    if _result3 != "" {
        args = append(args, models.Where{Column:"result3", Value:_result3, Compare:"="})
    }
    _result4 := c.Get("result4")
    if _result4 != "" {
        args = append(args, models.Where{Column:"result4", Value:_result4, Compare:"="})
    }
    _result5 := c.Get("result5")
    if _result5 != "" {
        args = append(args, models.Where{Column:"result5", Value:_result5, Compare:"="})
    }
    _result6 := c.Get("result6")
    if _result6 != "" {
        args = append(args, models.Where{Column:"result6", Value:_result6, Compare:"="})
    }
    _status1 := c.Get("status1")
    if _status1 != "" {
        args = append(args, models.Where{Column:"status1", Value:_status1, Compare:"="})
    }
    _status2 := c.Get("status2")
    if _status2 != "" {
        args = append(args, models.Where{Column:"status2", Value:_status2, Compare:"="})
    }
    _status3 := c.Get("status3")
    if _status3 != "" {
        args = append(args, models.Where{Column:"status3", Value:_status3, Compare:"="})
    }
    _status4 := c.Get("status4")
    if _status4 != "" {
        args = append(args, models.Where{Column:"status4", Value:_status4, Compare:"="})
    }
    _status5 := c.Get("status5")
    if _status5 != "" {
        args = append(args, models.Where{Column:"status5", Value:_status5, Compare:"="})
    }
    _status6 := c.Get("status6")
    if _status6 != "" {
        args = append(args, models.Where{Column:"status6", Value:_status6, Compare:"="})
    }
    _position1 := c.Get("position1")
    if _position1 != "" {
        args = append(args, models.Where{Column:"position1", Value:_position1, Compare:"="})
    }
    _position2 := c.Get("position2")
    if _position2 != "" {
        args = append(args, models.Where{Column:"position2", Value:_position2, Compare:"="})
    }
    _position3 := c.Get("position3")
    if _position3 != "" {
        args = append(args, models.Where{Column:"position3", Value:_position3, Compare:"="})
    }
    _position4 := c.Get("position4")
    if _position4 != "" {
        args = append(args, models.Where{Column:"position4", Value:_position4, Compare:"="})
    }
    _position5 := c.Get("position5")
    if _position5 != "" {
        args = append(args, models.Where{Column:"position5", Value:_position5, Compare:"="})
    }
    _position6 := c.Get("position6")
    if _position6 != "" {
        args = append(args, models.Where{Column:"position6", Value:_position6, Compare:"="})
    }
    _content := c.Get("content")
    if _content != "" {
        args = append(args, models.Where{Column:"content", Value:_content, Compare:"="})
        
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


func (c *PeriodicouterwallController) Read(id int64) {
    
    
	conn := c.NewConnection()

	manager := models.NewPeriodicouterwallManager(conn)
	item := manager.Get(id)

    
    
    c.Set("item", item)
}

func (c *PeriodicouterwallController) Index(page int, pagesize int) {
    
    
	conn := c.NewConnection()

	manager := models.NewPeriodicouterwallManager(conn)

    var args []interface{}
    
    _result1 := c.Get("result1")
    if _result1 != "" {
        args = append(args, models.Where{Column:"result1", Value:_result1, Compare:"="})
    }
    _result2 := c.Get("result2")
    if _result2 != "" {
        args = append(args, models.Where{Column:"result2", Value:_result2, Compare:"="})
    }
    _result3 := c.Get("result3")
    if _result3 != "" {
        args = append(args, models.Where{Column:"result3", Value:_result3, Compare:"="})
    }
    _result4 := c.Get("result4")
    if _result4 != "" {
        args = append(args, models.Where{Column:"result4", Value:_result4, Compare:"="})
    }
    _result5 := c.Get("result5")
    if _result5 != "" {
        args = append(args, models.Where{Column:"result5", Value:_result5, Compare:"="})
    }
    _result6 := c.Get("result6")
    if _result6 != "" {
        args = append(args, models.Where{Column:"result6", Value:_result6, Compare:"="})
    }
    _status1 := c.Get("status1")
    if _status1 != "" {
        args = append(args, models.Where{Column:"status1", Value:_status1, Compare:"="})
    }
    _status2 := c.Get("status2")
    if _status2 != "" {
        args = append(args, models.Where{Column:"status2", Value:_status2, Compare:"="})
    }
    _status3 := c.Get("status3")
    if _status3 != "" {
        args = append(args, models.Where{Column:"status3", Value:_status3, Compare:"="})
    }
    _status4 := c.Get("status4")
    if _status4 != "" {
        args = append(args, models.Where{Column:"status4", Value:_status4, Compare:"="})
    }
    _status5 := c.Get("status5")
    if _status5 != "" {
        args = append(args, models.Where{Column:"status5", Value:_status5, Compare:"="})
    }
    _status6 := c.Get("status6")
    if _status6 != "" {
        args = append(args, models.Where{Column:"status6", Value:_status6, Compare:"="})
    }
    _position1 := c.Get("position1")
    if _position1 != "" {
        args = append(args, models.Where{Column:"position1", Value:_position1, Compare:"="})
    }
    _position2 := c.Get("position2")
    if _position2 != "" {
        args = append(args, models.Where{Column:"position2", Value:_position2, Compare:"="})
    }
    _position3 := c.Get("position3")
    if _position3 != "" {
        args = append(args, models.Where{Column:"position3", Value:_position3, Compare:"="})
    }
    _position4 := c.Get("position4")
    if _position4 != "" {
        args = append(args, models.Where{Column:"position4", Value:_position4, Compare:"="})
    }
    _position5 := c.Get("position5")
    if _position5 != "" {
        args = append(args, models.Where{Column:"position5", Value:_position5, Compare:"="})
    }
    _position6 := c.Get("position6")
    if _position6 != "" {
        args = append(args, models.Where{Column:"position6", Value:_position6, Compare:"="})
    }
    _content := c.Get("content")
    if _content != "" {
        args = append(args, models.Where{Column:"content", Value:_content, Compare:"="})
        
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
                    str += ", po_" + strings.Trim(v, " ")                
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





