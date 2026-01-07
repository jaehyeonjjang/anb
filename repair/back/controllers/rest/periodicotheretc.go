package rest


import (
	"repair/controllers"
	"repair/models"

    "strings"
)

type PeriodicotheretcController struct {
	controllers.Controller
}



func (c *PeriodicotheretcController) GetByPeriodic(periodic int64) *models.Periodicotheretc {
    
    conn := c.NewConnection()

	_manager := models.NewPeriodicotheretcManager(conn)
    
    item := _manager.GetByPeriodic(periodic)
    
    c.Set("item", item)
    
    
    
    return item
    
}


func (c *PeriodicotheretcController) Insert(item *models.Periodicotheretc) {
    
    
	conn := c.NewConnection()
    
	manager := models.NewPeriodicotheretcManager(conn)
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

func (c *PeriodicotheretcController) Insertbatch(item *[]models.Periodicotheretc) {  
    if item == nil || len(*item) == 0 {
        return
    }

    rows := len(*item)
    
    
    
	conn := c.NewConnection()
    
	manager := models.NewPeriodicotheretcManager(conn)

    for i := 0; i < rows; i++ {
	    err := manager.Insert(&((*item)[i]))
        if err != nil {
            c.Set("code", "error")    
            c.Set("error", err)
            return
        }
    }
}

func (c *PeriodicotheretcController) Update(item *models.Periodicotheretc) {
    
    
	conn := c.NewConnection()

	manager := models.NewPeriodicotheretcManager(conn)
    err := manager.Update(item)
    if err != nil {
        c.Set("code", "error")    
        c.Set("error", err)
        return
    }
}

func (c *PeriodicotheretcController) Delete(item *models.Periodicotheretc) {
    
    
    conn := c.NewConnection()

	manager := models.NewPeriodicotheretcManager(conn)

    
	err := manager.Delete(item.Id)
    if err != nil {
        c.Set("code", "error")    
        c.Set("error", err)
    }
}

func (c *PeriodicotheretcController) Deletebatch(item *[]models.Periodicotheretc) {
    
    
    conn := c.NewConnection()

	manager := models.NewPeriodicotheretcManager(conn)

    for _, v := range *item {
        
    
	    err := manager.Delete(v.Id)
        if err != nil {
            c.Set("code", "error")    
            c.Set("error", err)
            return
        }
    }
}

func (c *PeriodicotheretcController) Count() {
    
    
	conn := c.NewConnection()

	manager := models.NewPeriodicotheretcManager(conn)

    var args []interface{}
    
    _content1 := c.Get("content1")
    if _content1 != "" {
        args = append(args, models.Where{Column:"content1", Value:_content1, Compare:"="})
    }
    _content2 := c.Get("content2")
    if _content2 != "" {
        args = append(args, models.Where{Column:"content2", Value:_content2, Compare:"="})
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


func (c *PeriodicotheretcController) Read(id int64) {
    
    
	conn := c.NewConnection()

	manager := models.NewPeriodicotheretcManager(conn)
	item := manager.Get(id)

    
    
    c.Set("item", item)
}

func (c *PeriodicotheretcController) Index(page int, pagesize int) {
    
    
	conn := c.NewConnection()

	manager := models.NewPeriodicotheretcManager(conn)

    var args []interface{}
    
    _content1 := c.Get("content1")
    if _content1 != "" {
        args = append(args, models.Where{Column:"content1", Value:_content1, Compare:"="})
    }
    _content2 := c.Get("content2")
    if _content2 != "" {
        args = append(args, models.Where{Column:"content2", Value:_content2, Compare:"="})
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
                    str += ", pe_" + strings.Trim(v, " ")                
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





