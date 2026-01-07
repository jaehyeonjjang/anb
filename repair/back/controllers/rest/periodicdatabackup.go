package rest


import (
	"repair/controllers"
	"repair/models"

    "strings"
)

type PeriodicdatabackupController struct {
	controllers.Controller
}





func (c *PeriodicdatabackupController) Count() {
    
    
	conn := c.NewConnection()

	manager := models.NewPeriodicdatabackupManager(conn)

    var args []interface{}
    
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
    _blueprint := c.Geti64("blueprint")
    if _blueprint != 0 {
        args = append(args, models.Where{Column:"blueprint", Value:_blueprint, Compare:"="})    
    }
    _count := c.Geti64("count")
    if _count != 0 {
        args = append(args, models.Where{Column:"count", Value:_count, Compare:"="})    
    }
    

    
    
    total := manager.Count(args)
	c.Set("total", total)
}


func (c *PeriodicdatabackupController) Read(id int64) {
    
    
	conn := c.NewConnection()

	manager := models.NewPeriodicdatabackupManager(conn)
	item := manager.Get(id)

    
    
    c.Set("item", item)
}

func (c *PeriodicdatabackupController) Index(page int, pagesize int) {
    
    
	conn := c.NewConnection()

	manager := models.NewPeriodicdatabackupManager(conn)

    var args []interface{}
    
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
    _blueprint := c.Geti64("blueprint")
    if _blueprint != 0 {
        args = append(args, models.Where{Column:"blueprint", Value:_blueprint, Compare:"="})    
    }
    _count := c.Geti64("count")
    if _count != 0 {
        args = append(args, models.Where{Column:"count", Value:_count, Compare:"="})    
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
                    str += ", pd_" + strings.Trim(v, " ")                
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





func (c *PeriodicdatabackupController) Sum() {
    
    
	conn := c.NewConnection()

	manager := models.NewPeriodicdatabackupManager(conn)

    var args []interface{}
    
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
    _blueprint := c.Geti64("blueprint")
    if _blueprint != 0 {
        args = append(args, models.Where{Column:"blueprint", Value:_blueprint, Compare:"="})    
    }
    _count := c.Geti64("count")
    if _count != 0 {
        args = append(args, models.Where{Column:"count", Value:_count, Compare:"="})    
    }
    

    
    
    item := manager.Sum(args)
	c.Set("item", item)
}

