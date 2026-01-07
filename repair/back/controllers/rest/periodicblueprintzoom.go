package rest


import (
	"repair/controllers"
	"repair/models"

    "strings"
)

type PeriodicblueprintzoomController struct {
	controllers.Controller
}



func (c *PeriodicblueprintzoomController) GetByPeriodicBlueprint(periodic int64 ,blueprint int64) *models.Periodicblueprintzoom {
    
    conn := c.NewConnection()

	_manager := models.NewPeriodicblueprintzoomManager(conn)
    
    item := _manager.GetByPeriodicBlueprint(periodic, blueprint)
    
    c.Set("item", item)
    
    
    
    return item
    
}

// @Put()
func (c *PeriodicblueprintzoomController) UpdateStatusByPeriodicBlueprint(status int ,periodic int64 ,blueprint int64) {
    
    conn := c.NewConnection()

	_manager := models.NewPeriodicblueprintzoomManager(conn)
    
    err := _manager.UpdateStatusByPeriodicBlueprint(status, periodic, blueprint)
    if err != nil {
        c.Set("code", "error")    
        c.Set("error", err)
        return
    }
    
}

// @Delete()
func (c *PeriodicblueprintzoomController) DeleteByPeriodicBlueprint(periodic int64 ,blueprint int64) {
    
    conn := c.NewConnection()

	_manager := models.NewPeriodicblueprintzoomManager(conn)
    
    err := _manager.DeleteByPeriodicBlueprint(periodic, blueprint)
    if err != nil {
        c.Set("code", "error")    
        c.Set("error", err)
        return
    }
    
}


func (c *PeriodicblueprintzoomController) Insert(item *models.Periodicblueprintzoom) {
    
    
	conn := c.NewConnection()
    
	manager := models.NewPeriodicblueprintzoomManager(conn)
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

func (c *PeriodicblueprintzoomController) Insertbatch(item *[]models.Periodicblueprintzoom) {  
    if item == nil || len(*item) == 0 {
        return
    }

    rows := len(*item)
    
    
    
	conn := c.NewConnection()
    
	manager := models.NewPeriodicblueprintzoomManager(conn)

    for i := 0; i < rows; i++ {
	    err := manager.Insert(&((*item)[i]))
        if err != nil {
            c.Set("code", "error")    
            c.Set("error", err)
            return
        }
    }
}

func (c *PeriodicblueprintzoomController) Update(item *models.Periodicblueprintzoom) {
    
    
	conn := c.NewConnection()

	manager := models.NewPeriodicblueprintzoomManager(conn)
    err := manager.Update(item)
    if err != nil {
        c.Set("code", "error")    
        c.Set("error", err)
        return
    }
}

func (c *PeriodicblueprintzoomController) Delete(item *models.Periodicblueprintzoom) {
    
    
    conn := c.NewConnection()

	manager := models.NewPeriodicblueprintzoomManager(conn)

    
	err := manager.Delete(item.Id)
    if err != nil {
        c.Set("code", "error")    
        c.Set("error", err)
    }
}

func (c *PeriodicblueprintzoomController) Deletebatch(item *[]models.Periodicblueprintzoom) {
    
    
    conn := c.NewConnection()

	manager := models.NewPeriodicblueprintzoomManager(conn)

    for _, v := range *item {
        
    
	    err := manager.Delete(v.Id)
        if err != nil {
            c.Set("code", "error")    
            c.Set("error", err)
            return
        }
    }
}

func (c *PeriodicblueprintzoomController) Count() {
    
    
	conn := c.NewConnection()

	manager := models.NewPeriodicblueprintzoomManager(conn)

    var args []interface{}
    
    _iconzoom := c.Geti("iconzoom")
    if _iconzoom != 0 {
        args = append(args, models.Where{Column:"iconzoom", Value:_iconzoom, Compare:"="})    
    }
    _numberzoom := c.Geti("numberzoom")
    if _numberzoom != 0 {
        args = append(args, models.Where{Column:"numberzoom", Value:_numberzoom, Compare:"="})    
    }
    _crackzoom := c.Geti("crackzoom")
    if _crackzoom != 0 {
        args = append(args, models.Where{Column:"crackzoom", Value:_crackzoom, Compare:"="})    
    }
    _zoom := c.Geti("zoom")
    if _zoom != 0 {
        args = append(args, models.Where{Column:"zoom", Value:_zoom, Compare:"="})    
    }
    _status := c.Geti("status")
    if _status != 0 {
        args = append(args, models.Where{Column:"status", Value:_status, Compare:"="})    
    }
    _blueprint := c.Geti64("blueprint")
    if _blueprint != 0 {
        args = append(args, models.Where{Column:"blueprint", Value:_blueprint, Compare:"="})    
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


func (c *PeriodicblueprintzoomController) Read(id int64) {
    
    
	conn := c.NewConnection()

	manager := models.NewPeriodicblueprintzoomManager(conn)
	item := manager.Get(id)

    
    
    c.Set("item", item)
}

func (c *PeriodicblueprintzoomController) Index(page int, pagesize int) {
    
    
	conn := c.NewConnection()

	manager := models.NewPeriodicblueprintzoomManager(conn)

    var args []interface{}
    
    _iconzoom := c.Geti("iconzoom")
    if _iconzoom != 0 {
        args = append(args, models.Where{Column:"iconzoom", Value:_iconzoom, Compare:"="})    
    }
    _numberzoom := c.Geti("numberzoom")
    if _numberzoom != 0 {
        args = append(args, models.Where{Column:"numberzoom", Value:_numberzoom, Compare:"="})    
    }
    _crackzoom := c.Geti("crackzoom")
    if _crackzoom != 0 {
        args = append(args, models.Where{Column:"crackzoom", Value:_crackzoom, Compare:"="})    
    }
    _zoom := c.Geti("zoom")
    if _zoom != 0 {
        args = append(args, models.Where{Column:"zoom", Value:_zoom, Compare:"="})    
    }
    _status := c.Geti("status")
    if _status != 0 {
        args = append(args, models.Where{Column:"status", Value:_status, Compare:"="})    
    }
    _blueprint := c.Geti64("blueprint")
    if _blueprint != 0 {
        args = append(args, models.Where{Column:"blueprint", Value:_blueprint, Compare:"="})    
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
                    str += ", pb_" + strings.Trim(v, " ")                
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





