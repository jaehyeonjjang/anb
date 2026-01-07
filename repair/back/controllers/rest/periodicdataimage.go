package rest


import (
	"repair/controllers"
	"repair/models"

    "strings"
)

type PeriodicdataimageController struct {
	controllers.Controller
}



func (c *PeriodicdataimageController) GetByPeriodic(periodic int64) *models.Periodicdataimage {
    
    conn := c.NewConnection()

	_manager := models.NewPeriodicdataimageManager(conn)
    
    item := _manager.GetByPeriodic(periodic)
    
    c.Set("item", item)
    
    
    
    return item
    
}


func (c *PeriodicdataimageController) CountByOfflinefilename(offlinefilename string) int {
    
    conn := c.NewConnection()

	_manager := models.NewPeriodicdataimageManager(conn)
    
    item := _manager.CountByOfflinefilename(offlinefilename)
    
    
    
    c.Set("count", item)
    
    return item
    
}

// @Delete()
func (c *PeriodicdataimageController) DeleteByPeriodicdata(periodicdata int64) {
    
    conn := c.NewConnection()

	_manager := models.NewPeriodicdataimageManager(conn)
    
    err := _manager.DeleteByPeriodicdata(periodicdata)
    if err != nil {
        c.Set("code", "error")    
        c.Set("error", err)
        return
    }
    
}


func (c *PeriodicdataimageController) Insert(item *models.Periodicdataimage) {
    
    
	conn := c.NewConnection()
    
	manager := models.NewPeriodicdataimageManager(conn)
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

func (c *PeriodicdataimageController) Insertbatch(item *[]models.Periodicdataimage) {  
    if item == nil || len(*item) == 0 {
        return
    }

    rows := len(*item)
    
    
    
	conn := c.NewConnection()
    
	manager := models.NewPeriodicdataimageManager(conn)

    for i := 0; i < rows; i++ {
	    err := manager.Insert(&((*item)[i]))
        if err != nil {
            c.Set("code", "error")    
            c.Set("error", err)
            return
        }
    }
}

func (c *PeriodicdataimageController) Update(item *models.Periodicdataimage) {
    
    
	conn := c.NewConnection()

	manager := models.NewPeriodicdataimageManager(conn)
    err := manager.Update(item)
    if err != nil {
        c.Set("code", "error")    
        c.Set("error", err)
        return
    }
}

func (c *PeriodicdataimageController) Delete(item *models.Periodicdataimage) {
    
    
    conn := c.NewConnection()

	manager := models.NewPeriodicdataimageManager(conn)

    
	err := manager.Delete(item.Id)
    if err != nil {
        c.Set("code", "error")    
        c.Set("error", err)
    }
}

func (c *PeriodicdataimageController) Deletebatch(item *[]models.Periodicdataimage) {
    
    
    conn := c.NewConnection()

	manager := models.NewPeriodicdataimageManager(conn)

    for _, v := range *item {
        
    
	    err := manager.Delete(v.Id)
        if err != nil {
            c.Set("code", "error")    
            c.Set("error", err)
            return
        }
    }
}

func (c *PeriodicdataimageController) Count() {
    
    
	conn := c.NewConnection()

	manager := models.NewPeriodicdataimageManager(conn)

    var args []interface{}
    
    _filename := c.Get("filename")
    if _filename != "" {
        args = append(args, models.Where{Column:"filename", Value:_filename, Compare:"="})
    }
    _offlinefilename := c.Get("offlinefilename")
    if _offlinefilename != "" {
        args = append(args, models.Where{Column:"offlinefilename", Value:_offlinefilename, Compare:"="})
    }
    _order := c.Geti("order")
    if _order != 0 {
        args = append(args, models.Where{Column:"order", Value:_order, Compare:"="})    
    }
    _periodicdata := c.Geti64("periodicdata")
    if _periodicdata != 0 {
        args = append(args, models.Where{Column:"periodicdata", Value:_periodicdata, Compare:"="})    
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


func (c *PeriodicdataimageController) Read(id int64) {
    
    
	conn := c.NewConnection()

	manager := models.NewPeriodicdataimageManager(conn)
	item := manager.Get(id)

    
    
    c.Set("item", item)
}

func (c *PeriodicdataimageController) Index(page int, pagesize int) {
    
    
	conn := c.NewConnection()

	manager := models.NewPeriodicdataimageManager(conn)

    var args []interface{}
    
    _filename := c.Get("filename")
    if _filename != "" {
        args = append(args, models.Where{Column:"filename", Value:_filename, Compare:"="})
    }
    _offlinefilename := c.Get("offlinefilename")
    if _offlinefilename != "" {
        args = append(args, models.Where{Column:"offlinefilename", Value:_offlinefilename, Compare:"="})
    }
    _order := c.Geti("order")
    if _order != 0 {
        args = append(args, models.Where{Column:"order", Value:_order, Compare:"="})    
    }
    _periodicdata := c.Geti64("periodicdata")
    if _periodicdata != 0 {
        args = append(args, models.Where{Column:"periodicdata", Value:_periodicdata, Compare:"="})    
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
                    str += ", pi_" + strings.Trim(v, " ")                
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





