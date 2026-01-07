package rest


import (
	"repair/controllers"
	"repair/models"

    "strings"
)

type YearreportController struct {
	controllers.Controller
}



func (c *YearreportController) CountByApt(apt int64) int {
    
    conn := c.NewConnection()

	_manager := models.NewYearreportManager(conn)
    
    item := _manager.CountByApt(apt)
    
    
    
    c.Set("count", item)
    
    return item
    
}


func (c *YearreportController) FindByApt(apt int64) []models.Yearreport {
    
    conn := c.NewConnection()

	_manager := models.NewYearreportManager(conn)
    
    item := _manager.FindByApt(apt)
    
    
    c.Set("items", item)
    
    
    return item
    
}




func (c *YearreportController) Count() {
    
    
	conn := c.NewConnection()

	manager := models.NewYearreportManager(conn)

    var args []interface{}
    
    _apt := c.Geti64("apt")
    if _apt != 0 {
        args = append(args, models.Where{Column:"apt", Value:_apt, Compare:"="})    
    }
    _topcategory := c.Geti64("topcategory")
    if _topcategory != 0 {
        args = append(args, models.Where{Column:"topcategory", Value:_topcategory, Compare:"="})    
    }
    _subcategory := c.Geti64("subcategory")
    if _subcategory != 0 {
        args = append(args, models.Where{Column:"subcategory", Value:_subcategory, Compare:"="})    
    }
    _category := c.Geti64("category")
    if _category != 0 {
        args = append(args, models.Where{Column:"category", Value:_category, Compare:"="})    
    }
    _standard := c.Geti64("standard")
    if _standard != 0 {
        args = append(args, models.Where{Column:"standard", Value:_standard, Compare:"="})    
    }
    _method := c.Geti64("method")
    if _method != 0 {
        args = append(args, models.Where{Column:"method", Value:_method, Compare:"="})    
    }
    _rate := c.Geti("rate")
    if _rate != 0 {
        args = append(args, models.Where{Column:"rate", Value:_rate, Compare:"="})    
    }
    _duedate := c.Geti("duedate")
    if _duedate != 0 {
        args = append(args, models.Where{Column:"duedate", Value:_duedate, Compare:"="})    
    }
    _count := c.Geti("count")
    if _count != 0 {
        args = append(args, models.Where{Column:"count", Value:_count, Compare:"="})    
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


func (c *YearreportController) Read(id int64) {
    
    
	conn := c.NewConnection()

	manager := models.NewYearreportManager(conn)
	item := manager.Get(id)

    
    
    c.Set("item", item)
}

func (c *YearreportController) Index(page int, pagesize int) {
    
    
	conn := c.NewConnection()

	manager := models.NewYearreportManager(conn)

    var args []interface{}
    
    _apt := c.Geti64("apt")
    if _apt != 0 {
        args = append(args, models.Where{Column:"apt", Value:_apt, Compare:"="})    
    }
    _topcategory := c.Geti64("topcategory")
    if _topcategory != 0 {
        args = append(args, models.Where{Column:"topcategory", Value:_topcategory, Compare:"="})    
    }
    _subcategory := c.Geti64("subcategory")
    if _subcategory != 0 {
        args = append(args, models.Where{Column:"subcategory", Value:_subcategory, Compare:"="})    
    }
    _category := c.Geti64("category")
    if _category != 0 {
        args = append(args, models.Where{Column:"category", Value:_category, Compare:"="})    
    }
    _standard := c.Geti64("standard")
    if _standard != 0 {
        args = append(args, models.Where{Column:"standard", Value:_standard, Compare:"="})    
    }
    _method := c.Geti64("method")
    if _method != 0 {
        args = append(args, models.Where{Column:"method", Value:_method, Compare:"="})    
    }
    _rate := c.Geti("rate")
    if _rate != 0 {
        args = append(args, models.Where{Column:"rate", Value:_rate, Compare:"="})    
    }
    _duedate := c.Geti("duedate")
    if _duedate != 0 {
        args = append(args, models.Where{Column:"duedate", Value:_duedate, Compare:"="})    
    }
    _count := c.Geti("count")
    if _count != 0 {
        args = append(args, models.Where{Column:"count", Value:_count, Compare:"="})    
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
                    str += ", b_" + strings.Trim(v, " ")                
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





func (c *YearreportController) Sum() {
    
    
	conn := c.NewConnection()

	manager := models.NewYearreportManager(conn)

    var args []interface{}
    
    _apt := c.Geti64("apt")
    if _apt != 0 {
        args = append(args, models.Where{Column:"apt", Value:_apt, Compare:"="})    
    }
    _topcategory := c.Geti64("topcategory")
    if _topcategory != 0 {
        args = append(args, models.Where{Column:"topcategory", Value:_topcategory, Compare:"="})    
    }
    _subcategory := c.Geti64("subcategory")
    if _subcategory != 0 {
        args = append(args, models.Where{Column:"subcategory", Value:_subcategory, Compare:"="})    
    }
    _category := c.Geti64("category")
    if _category != 0 {
        args = append(args, models.Where{Column:"category", Value:_category, Compare:"="})    
    }
    _standard := c.Geti64("standard")
    if _standard != 0 {
        args = append(args, models.Where{Column:"standard", Value:_standard, Compare:"="})    
    }
    _method := c.Geti64("method")
    if _method != 0 {
        args = append(args, models.Where{Column:"method", Value:_method, Compare:"="})    
    }
    _rate := c.Geti("rate")
    if _rate != 0 {
        args = append(args, models.Where{Column:"rate", Value:_rate, Compare:"="})    
    }
    _duedate := c.Geti("duedate")
    if _duedate != 0 {
        args = append(args, models.Where{Column:"duedate", Value:_duedate, Compare:"="})    
    }
    _count := c.Geti("count")
    if _count != 0 {
        args = append(args, models.Where{Column:"count", Value:_count, Compare:"="})    
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
    

    
    
    item := manager.Sum(args)
	c.Set("item", item)
}

