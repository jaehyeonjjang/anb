package rest


import (
	"repair/controllers"
	"repair/models"

    "strings"
)

type OutlineController struct {
	controllers.Controller
}



func (c *OutlineController) Insert(item *models.Outline) {
    
    
	conn := c.NewConnection()
    
	manager := models.NewOutlineManager(conn)
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

func (c *OutlineController) Insertbatch(item *[]models.Outline) {  
    if item == nil || len(*item) == 0 {
        return
    }

    rows := len(*item)
    
    
    
	conn := c.NewConnection()
    
	manager := models.NewOutlineManager(conn)

    for i := 0; i < rows; i++ {
	    err := manager.Insert(&((*item)[i]))
        if err != nil {
            c.Set("code", "error")    
            c.Set("error", err)
            return
        }
    }
}

func (c *OutlineController) Update(item *models.Outline) {
    
    
	conn := c.NewConnection()

	manager := models.NewOutlineManager(conn)
    err := manager.Update(item)
    if err != nil {
        c.Set("code", "error")    
        c.Set("error", err)
        return
    }
}

func (c *OutlineController) Delete(item *models.Outline) {
    
    
    conn := c.NewConnection()

	manager := models.NewOutlineManager(conn)

    
	err := manager.Delete(item.Id)
    if err != nil {
        c.Set("code", "error")    
        c.Set("error", err)
    }
}

func (c *OutlineController) Deletebatch(item *[]models.Outline) {
    
    
    conn := c.NewConnection()

	manager := models.NewOutlineManager(conn)

    for _, v := range *item {
        
    
	    err := manager.Delete(v.Id)
        if err != nil {
            c.Set("code", "error")    
            c.Set("error", err)
            return
        }
    }
}

func (c *OutlineController) Count() {
    
    
	conn := c.NewConnection()

	manager := models.NewOutlineManager(conn)

    var args []interface{}
    
    _startyear := c.Geti("startyear")
    if _startyear != 0 {
        args = append(args, models.Where{Column:"startyear", Value:_startyear, Compare:"="})    
    }
    _endyear := c.Geti("endyear")
    if _endyear != 0 {
        args = append(args, models.Where{Column:"endyear", Value:_endyear, Compare:"="})    
    }
    _startmonth := c.Geti("startmonth")
    if _startmonth != 0 {
        args = append(args, models.Where{Column:"startmonth", Value:_startmonth, Compare:"="})    
    }
    _endmonth := c.Geti("endmonth")
    if _endmonth != 0 {
        args = append(args, models.Where{Column:"endmonth", Value:_endmonth, Compare:"="})    
    }
    _rate := c.Geti("rate")
    if _rate != 0 {
        args = append(args, models.Where{Column:"rate", Value:_rate, Compare:"="})    
    }
    _price := c.Geti("price")
    if _price != 0 {
        args = append(args, models.Where{Column:"price", Value:_price, Compare:"="})    
    }
    _remark := c.Get("remark")
    if _remark != "" {
        args = append(args, models.Where{Column:"remark", Value:_remark, Compare:"="})
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


func (c *OutlineController) Read(id int64) {
    
    
	conn := c.NewConnection()

	manager := models.NewOutlineManager(conn)
	item := manager.Get(id)

    
    
    c.Set("item", item)
}

func (c *OutlineController) Index(page int, pagesize int) {
    
    
	conn := c.NewConnection()

	manager := models.NewOutlineManager(conn)

    var args []interface{}
    
    _startyear := c.Geti("startyear")
    if _startyear != 0 {
        args = append(args, models.Where{Column:"startyear", Value:_startyear, Compare:"="})    
    }
    _endyear := c.Geti("endyear")
    if _endyear != 0 {
        args = append(args, models.Where{Column:"endyear", Value:_endyear, Compare:"="})    
    }
    _startmonth := c.Geti("startmonth")
    if _startmonth != 0 {
        args = append(args, models.Where{Column:"startmonth", Value:_startmonth, Compare:"="})    
    }
    _endmonth := c.Geti("endmonth")
    if _endmonth != 0 {
        args = append(args, models.Where{Column:"endmonth", Value:_endmonth, Compare:"="})    
    }
    _rate := c.Geti("rate")
    if _rate != 0 {
        args = append(args, models.Where{Column:"rate", Value:_rate, Compare:"="})    
    }
    _price := c.Geti("price")
    if _price != 0 {
        args = append(args, models.Where{Column:"price", Value:_price, Compare:"="})    
    }
    _remark := c.Get("remark")
    if _remark != "" {
        args = append(args, models.Where{Column:"remark", Value:_remark, Compare:"="})
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
                    str += ", o_" + strings.Trim(v, " ")                
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





func (c *OutlineController) Sum() {
    
    
	conn := c.NewConnection()

	manager := models.NewOutlineManager(conn)

    var args []interface{}
    
    _startyear := c.Geti("startyear")
    if _startyear != 0 {
        args = append(args, models.Where{Column:"startyear", Value:_startyear, Compare:"="})    
    }
    _endyear := c.Geti("endyear")
    if _endyear != 0 {
        args = append(args, models.Where{Column:"endyear", Value:_endyear, Compare:"="})    
    }
    _startmonth := c.Geti("startmonth")
    if _startmonth != 0 {
        args = append(args, models.Where{Column:"startmonth", Value:_startmonth, Compare:"="})    
    }
    _endmonth := c.Geti("endmonth")
    if _endmonth != 0 {
        args = append(args, models.Where{Column:"endmonth", Value:_endmonth, Compare:"="})    
    }
    _rate := c.Geti("rate")
    if _rate != 0 {
        args = append(args, models.Where{Column:"rate", Value:_rate, Compare:"="})    
    }
    _price := c.Geti("price")
    if _price != 0 {
        args = append(args, models.Where{Column:"price", Value:_price, Compare:"="})    
    }
    _remark := c.Get("remark")
    if _remark != "" {
        args = append(args, models.Where{Column:"remark", Value:_remark, Compare:"like"})
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
    

    
    
    item := manager.Sum(args)
	c.Set("item", item)
}

