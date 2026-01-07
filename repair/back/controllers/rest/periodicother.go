package rest


import (
	"repair/controllers"
	"repair/models"

    "strings"
)

type PeriodicotherController struct {
	controllers.Controller
}



func (c *PeriodicotherController) Insert(item *models.Periodicother) {
    
    
	conn := c.NewConnection()
    
	manager := models.NewPeriodicotherManager(conn)
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

func (c *PeriodicotherController) Insertbatch(item *[]models.Periodicother) {  
    if item == nil || len(*item) == 0 {
        return
    }

    rows := len(*item)
    
    
    
	conn := c.NewConnection()
    
	manager := models.NewPeriodicotherManager(conn)

    for i := 0; i < rows; i++ {
	    err := manager.Insert(&((*item)[i]))
        if err != nil {
            c.Set("code", "error")    
            c.Set("error", err)
            return
        }
    }
}

func (c *PeriodicotherController) Update(item *models.Periodicother) {
    
    
	conn := c.NewConnection()

	manager := models.NewPeriodicotherManager(conn)
    err := manager.Update(item)
    if err != nil {
        c.Set("code", "error")    
        c.Set("error", err)
        return
    }
}

func (c *PeriodicotherController) Delete(item *models.Periodicother) {
    
    
    conn := c.NewConnection()

	manager := models.NewPeriodicotherManager(conn)

    
	err := manager.Delete(item.Id)
    if err != nil {
        c.Set("code", "error")    
        c.Set("error", err)
    }
}

func (c *PeriodicotherController) Deletebatch(item *[]models.Periodicother) {
    
    
    conn := c.NewConnection()

	manager := models.NewPeriodicotherManager(conn)

    for _, v := range *item {
        
    
	    err := manager.Delete(v.Id)
        if err != nil {
            c.Set("code", "error")    
            c.Set("error", err)
            return
        }
    }
}

func (c *PeriodicotherController) Count() {
    
    
	conn := c.NewConnection()

	manager := models.NewPeriodicotherManager(conn)

    var args []interface{}
    
    _name := c.Get("name")
    if _name != "" {
        args = append(args, models.Where{Column:"name", Value:_name, Compare:"="})
        
    }
    _type := c.Geti("type")
    if _type != 0 {
        args = append(args, models.Where{Column:"type", Value:_type, Compare:"="})    
    }
    _result := c.Geti("result")
    if _result != 0 {
        args = append(args, models.Where{Column:"result", Value:_result, Compare:"="})    
    }
    _status := c.Get("status")
    if _status != "" {
        args = append(args, models.Where{Column:"status", Value:_status, Compare:"="})
    }
    _position := c.Get("position")
    if _position != "" {
        args = append(args, models.Where{Column:"position", Value:_position, Compare:"="})
    }
    _filename := c.Get("filename")
    if _filename != "" {
        args = append(args, models.Where{Column:"filename", Value:_filename, Compare:"="})
    }
    _offlinefilename := c.Get("offlinefilename")
    if _offlinefilename != "" {
        args = append(args, models.Where{Column:"offlinefilename", Value:_offlinefilename, Compare:"="})
    }
    _change := c.Geti("change")
    if _change != 0 {
        args = append(args, models.Where{Column:"change", Value:_change, Compare:"="})    
    }
    _category := c.Geti("category")
    if _category != 0 {
        args = append(args, models.Where{Column:"category", Value:_category, Compare:"="})    
    }
    _order := c.Geti("order")
    if _order != 0 {
        args = append(args, models.Where{Column:"order", Value:_order, Compare:"="})    
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


func (c *PeriodicotherController) Read(id int64) {
    
    
	conn := c.NewConnection()

	manager := models.NewPeriodicotherManager(conn)
	item := manager.Get(id)

    
    
    c.Set("item", item)
}

func (c *PeriodicotherController) Index(page int, pagesize int) {
    
    
	conn := c.NewConnection()

	manager := models.NewPeriodicotherManager(conn)

    var args []interface{}
    
    _name := c.Get("name")
    if _name != "" {
        args = append(args, models.Where{Column:"name", Value:_name, Compare:"="})
        
    }
    _type := c.Geti("type")
    if _type != 0 {
        args = append(args, models.Where{Column:"type", Value:_type, Compare:"="})    
    }
    _result := c.Geti("result")
    if _result != 0 {
        args = append(args, models.Where{Column:"result", Value:_result, Compare:"="})    
    }
    _status := c.Get("status")
    if _status != "" {
        args = append(args, models.Where{Column:"status", Value:_status, Compare:"="})
    }
    _position := c.Get("position")
    if _position != "" {
        args = append(args, models.Where{Column:"position", Value:_position, Compare:"="})
    }
    _filename := c.Get("filename")
    if _filename != "" {
        args = append(args, models.Where{Column:"filename", Value:_filename, Compare:"="})
    }
    _offlinefilename := c.Get("offlinefilename")
    if _offlinefilename != "" {
        args = append(args, models.Where{Column:"offlinefilename", Value:_offlinefilename, Compare:"="})
    }
    _change := c.Geti("change")
    if _change != 0 {
        args = append(args, models.Where{Column:"change", Value:_change, Compare:"="})    
    }
    _category := c.Geti("category")
    if _category != 0 {
        args = append(args, models.Where{Column:"category", Value:_category, Compare:"="})    
    }
    _order := c.Geti("order")
    if _order != 0 {
        args = append(args, models.Where{Column:"order", Value:_order, Compare:"="})    
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





