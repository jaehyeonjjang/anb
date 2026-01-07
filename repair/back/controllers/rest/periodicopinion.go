package rest


import (
	"repair/controllers"
	"repair/models"

    "strings"
)

type PeriodicopinionController struct {
	controllers.Controller
}



func (c *PeriodicopinionController) GetByPeriodic(periodic int64) *models.Periodicopinion {
    
    conn := c.NewConnection()

	_manager := models.NewPeriodicopinionManager(conn)
    
    item := _manager.GetByPeriodic(periodic)
    
    c.Set("item", item)
    
    
    
    return item
    
}


func (c *PeriodicopinionController) Insert(item *models.Periodicopinion) {
    
    
	conn := c.NewConnection()
    
	manager := models.NewPeriodicopinionManager(conn)
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

func (c *PeriodicopinionController) Insertbatch(item *[]models.Periodicopinion) {  
    if item == nil || len(*item) == 0 {
        return
    }

    rows := len(*item)
    
    
    
	conn := c.NewConnection()
    
	manager := models.NewPeriodicopinionManager(conn)

    for i := 0; i < rows; i++ {
	    err := manager.Insert(&((*item)[i]))
        if err != nil {
            c.Set("code", "error")    
            c.Set("error", err)
            return
        }
    }
}

func (c *PeriodicopinionController) Update(item *models.Periodicopinion) {
    
    
	conn := c.NewConnection()

	manager := models.NewPeriodicopinionManager(conn)
    err := manager.Update(item)
    if err != nil {
        c.Set("code", "error")    
        c.Set("error", err)
        return
    }
}

func (c *PeriodicopinionController) Delete(item *models.Periodicopinion) {
    
    
    conn := c.NewConnection()

	manager := models.NewPeriodicopinionManager(conn)

    
	err := manager.Delete(item.Id)
    if err != nil {
        c.Set("code", "error")    
        c.Set("error", err)
    }
}

func (c *PeriodicopinionController) Deletebatch(item *[]models.Periodicopinion) {
    
    
    conn := c.NewConnection()

	manager := models.NewPeriodicopinionManager(conn)

    for _, v := range *item {
        
    
	    err := manager.Delete(v.Id)
        if err != nil {
            c.Set("code", "error")    
            c.Set("error", err)
            return
        }
    }
}

func (c *PeriodicopinionController) Count() {
    
    
	conn := c.NewConnection()

	manager := models.NewPeriodicopinionManager(conn)

    var args []interface{}
    
    _grade := c.Geti("grade")
    if _grade != 0 {
        args = append(args, models.Where{Column:"grade", Value:_grade, Compare:"="})    
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
    _cause1 := c.Get("cause1")
    if _cause1 != "" {
        args = append(args, models.Where{Column:"cause1", Value:_cause1, Compare:"="})
    }
    _cause2 := c.Get("cause2")
    if _cause2 != "" {
        args = append(args, models.Where{Column:"cause2", Value:_cause2, Compare:"="})
    }
    _cause3 := c.Get("cause3")
    if _cause3 != "" {
        args = append(args, models.Where{Column:"cause3", Value:_cause3, Compare:"="})
    }
    _cause4 := c.Get("cause4")
    if _cause4 != "" {
        args = append(args, models.Where{Column:"cause4", Value:_cause4, Compare:"="})
    }
    _cause5 := c.Get("cause5")
    if _cause5 != "" {
        args = append(args, models.Where{Column:"cause5", Value:_cause5, Compare:"="})
    }
    _cause6 := c.Get("cause6")
    if _cause6 != "" {
        args = append(args, models.Where{Column:"cause6", Value:_cause6, Compare:"="})
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


func (c *PeriodicopinionController) Read(id int64) {
    
    
	conn := c.NewConnection()

	manager := models.NewPeriodicopinionManager(conn)
	item := manager.Get(id)

    
    
    c.Set("item", item)
}

func (c *PeriodicopinionController) Index(page int, pagesize int) {
    
    
	conn := c.NewConnection()

	manager := models.NewPeriodicopinionManager(conn)

    var args []interface{}
    
    _grade := c.Geti("grade")
    if _grade != 0 {
        args = append(args, models.Where{Column:"grade", Value:_grade, Compare:"="})    
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
    _cause1 := c.Get("cause1")
    if _cause1 != "" {
        args = append(args, models.Where{Column:"cause1", Value:_cause1, Compare:"="})
    }
    _cause2 := c.Get("cause2")
    if _cause2 != "" {
        args = append(args, models.Where{Column:"cause2", Value:_cause2, Compare:"="})
    }
    _cause3 := c.Get("cause3")
    if _cause3 != "" {
        args = append(args, models.Where{Column:"cause3", Value:_cause3, Compare:"="})
    }
    _cause4 := c.Get("cause4")
    if _cause4 != "" {
        args = append(args, models.Where{Column:"cause4", Value:_cause4, Compare:"="})
    }
    _cause5 := c.Get("cause5")
    if _cause5 != "" {
        args = append(args, models.Where{Column:"cause5", Value:_cause5, Compare:"="})
    }
    _cause6 := c.Get("cause6")
    if _cause6 != "" {
        args = append(args, models.Where{Column:"cause6", Value:_cause6, Compare:"="})
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





