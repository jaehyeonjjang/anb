package rest


import (
	"repair/controllers"
	"repair/models"

    "strings"
)

type AptusagefloorController struct {
	controllers.Controller
}



func (c *AptusagefloorController) CountByApt(apt int64) int {
    
    conn := c.NewConnection()

	_manager := models.NewAptusagefloorManager(conn)
    
    item := _manager.CountByApt(apt)
    
    
    
    c.Set("count", item)
    
    return item
    
}


func (c *AptusagefloorController) FindByApt(apt int64) []models.Aptusagefloor {
    
    conn := c.NewConnection()

	_manager := models.NewAptusagefloorManager(conn)
    
    item := _manager.FindByApt(apt)
    
    
    c.Set("items", item)
    
    
    return item
    
}

// @Delete()
func (c *AptusagefloorController) DeleteByApt(apt int64) {
    
    conn := c.NewConnection()

	_manager := models.NewAptusagefloorManager(conn)
    
    err := _manager.DeleteByApt(apt)
    if err != nil {
        c.Set("code", "error")    
        c.Set("error", err)
        return
    }
    
}


func (c *AptusagefloorController) Insert(item *models.Aptusagefloor) {
    
    
	conn := c.NewConnection()
    
	manager := models.NewAptusagefloorManager(conn)
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

func (c *AptusagefloorController) Insertbatch(item *[]models.Aptusagefloor) {  
    if item == nil || len(*item) == 0 {
        return
    }

    rows := len(*item)
    
    
    
	conn := c.NewConnection()
    
	manager := models.NewAptusagefloorManager(conn)

    for i := 0; i < rows; i++ {
	    err := manager.Insert(&((*item)[i]))
        if err != nil {
            c.Set("code", "error")    
            c.Set("error", err)
            return
        }
    }
}

func (c *AptusagefloorController) Update(item *models.Aptusagefloor) {
    
    
	conn := c.NewConnection()

	manager := models.NewAptusagefloorManager(conn)
    err := manager.Update(item)
    if err != nil {
        c.Set("code", "error")    
        c.Set("error", err)
        return
    }
}

func (c *AptusagefloorController) Delete(item *models.Aptusagefloor) {
    
    
    conn := c.NewConnection()

	manager := models.NewAptusagefloorManager(conn)

    
	err := manager.Delete(item.Id)
    if err != nil {
        c.Set("code", "error")    
        c.Set("error", err)
    }
}

func (c *AptusagefloorController) Deletebatch(item *[]models.Aptusagefloor) {
    
    
    conn := c.NewConnection()

	manager := models.NewAptusagefloorManager(conn)

    for _, v := range *item {
        
    
	    err := manager.Delete(v.Id)
        if err != nil {
            c.Set("code", "error")    
            c.Set("error", err)
            return
        }
    }
}

func (c *AptusagefloorController) Count() {
    
    
	conn := c.NewConnection()

	manager := models.NewAptusagefloorManager(conn)

    var args []interface{}
    
    _floor := c.Get("floor")
    if _floor != "" {
        args = append(args, models.Where{Column:"floor", Value:_floor, Compare:"="})
    }
    _purpose := c.Get("purpose")
    if _purpose != "" {
        args = append(args, models.Where{Column:"purpose", Value:_purpose, Compare:"="})
    }
    _area := c.Get("area")
    if _area != "" {
        args = append(args, models.Where{Column:"area", Value:_area, Compare:"="})
    }
    _remark := c.Get("remark")
    if _remark != "" {
        args = append(args, models.Where{Column:"remark", Value:_remark, Compare:"="})
    }
    _order := c.Geti("order")
    if _order != 0 {
        args = append(args, models.Where{Column:"order", Value:_order, Compare:"="})    
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


func (c *AptusagefloorController) Read(id int64) {
    
    
	conn := c.NewConnection()

	manager := models.NewAptusagefloorManager(conn)
	item := manager.Get(id)

    
    
    c.Set("item", item)
}

func (c *AptusagefloorController) Index(page int, pagesize int) {
    
    
	conn := c.NewConnection()

	manager := models.NewAptusagefloorManager(conn)

    var args []interface{}
    
    _floor := c.Get("floor")
    if _floor != "" {
        args = append(args, models.Where{Column:"floor", Value:_floor, Compare:"="})
    }
    _purpose := c.Get("purpose")
    if _purpose != "" {
        args = append(args, models.Where{Column:"purpose", Value:_purpose, Compare:"="})
    }
    _area := c.Get("area")
    if _area != "" {
        args = append(args, models.Where{Column:"area", Value:_area, Compare:"="})
    }
    _remark := c.Get("remark")
    if _remark != "" {
        args = append(args, models.Where{Column:"remark", Value:_remark, Compare:"="})
    }
    _order := c.Geti("order")
    if _order != 0 {
        args = append(args, models.Where{Column:"order", Value:_order, Compare:"="})    
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
                    str += ", af_" + strings.Trim(v, " ")                
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





