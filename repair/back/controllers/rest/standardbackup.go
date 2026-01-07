package rest


import (
	"repair/controllers"
	"repair/models"

    "strings"
)

type StandardbackupController struct {
	controllers.Controller
}



func (c *StandardbackupController) Insert(item *models.Standardbackup) {
    
    
	conn := c.NewConnection()
    
	manager := models.NewStandardbackupManager(conn)
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

func (c *StandardbackupController) Insertbatch(item *[]models.Standardbackup) {  
    if item == nil || len(*item) == 0 {
        return
    }

    rows := len(*item)
    
    
    
	conn := c.NewConnection()
    
	manager := models.NewStandardbackupManager(conn)

    for i := 0; i < rows; i++ {
	    err := manager.Insert(&((*item)[i]))
        if err != nil {
            c.Set("code", "error")    
            c.Set("error", err)
            return
        }
    }
}

func (c *StandardbackupController) Update(item *models.Standardbackup) {
    
    
	conn := c.NewConnection()

	manager := models.NewStandardbackupManager(conn)
    err := manager.Update(item)
    if err != nil {
        c.Set("code", "error")    
        c.Set("error", err)
        return
    }
}

func (c *StandardbackupController) Delete(item *models.Standardbackup) {
    
    
    conn := c.NewConnection()

	manager := models.NewStandardbackupManager(conn)

    
	err := manager.Delete(item.Id)
    if err != nil {
        c.Set("code", "error")    
        c.Set("error", err)
    }
}

func (c *StandardbackupController) Deletebatch(item *[]models.Standardbackup) {
    
    
    conn := c.NewConnection()

	manager := models.NewStandardbackupManager(conn)

    for _, v := range *item {
        
    
	    err := manager.Delete(v.Id)
        if err != nil {
            c.Set("code", "error")    
            c.Set("error", err)
            return
        }
    }
}

func (c *StandardbackupController) Count() {
    
    
	conn := c.NewConnection()

	manager := models.NewStandardbackupManager(conn)

    var args []interface{}
    
    _name := c.Get("name")
    if _name != "" {
        args = append(args, models.Where{Column:"name", Value:_name, Compare:"="})
        
    }
    _direct := c.Geti64("direct")
    if _direct != 0 {
        args = append(args, models.Where{Column:"direct", Value:_direct, Compare:"="})    
    }
    _labor := c.Geti("labor")
    if _labor != 0 {
        args = append(args, models.Where{Column:"labor", Value:_labor, Compare:"="})    
    }
    _cost := c.Geti("cost")
    if _cost != 0 {
        args = append(args, models.Where{Column:"cost", Value:_cost, Compare:"="})    
    }
    _unit := c.Get("unit")
    if _unit != "" {
        args = append(args, models.Where{Column:"unit", Value:_unit, Compare:"="})
    }
    _order := c.Geti("order")
    if _order != 0 {
        args = append(args, models.Where{Column:"order", Value:_order, Compare:"="})    
    }
    _original := c.Geti64("original")
    if _original != 0 {
        args = append(args, models.Where{Column:"original", Value:_original, Compare:"="})    
    }
    _category := c.Geti64("category")
    if _category != 0 {
        args = append(args, models.Where{Column:"category", Value:_category, Compare:"="})    
    }
    _apt := c.Geti64("apt")
    if _apt != 0 {
        args = append(args, models.Where{Column:"apt", Value:_apt, Compare:"="})    
    }
    _standard := c.Geti64("standard")
    if _standard != 0 {
        args = append(args, models.Where{Column:"standard", Value:_standard, Compare:"="})    
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


func (c *StandardbackupController) Read(id int64) {
    
    
	conn := c.NewConnection()

	manager := models.NewStandardbackupManager(conn)
	item := manager.Get(id)

    
    
    c.Set("item", item)
}

func (c *StandardbackupController) Index(page int, pagesize int) {
    
    
	conn := c.NewConnection()

	manager := models.NewStandardbackupManager(conn)

    var args []interface{}
    
    _name := c.Get("name")
    if _name != "" {
        args = append(args, models.Where{Column:"name", Value:_name, Compare:"="})
        
    }
    _direct := c.Geti64("direct")
    if _direct != 0 {
        args = append(args, models.Where{Column:"direct", Value:_direct, Compare:"="})    
    }
    _labor := c.Geti("labor")
    if _labor != 0 {
        args = append(args, models.Where{Column:"labor", Value:_labor, Compare:"="})    
    }
    _cost := c.Geti("cost")
    if _cost != 0 {
        args = append(args, models.Where{Column:"cost", Value:_cost, Compare:"="})    
    }
    _unit := c.Get("unit")
    if _unit != "" {
        args = append(args, models.Where{Column:"unit", Value:_unit, Compare:"="})
    }
    _order := c.Geti("order")
    if _order != 0 {
        args = append(args, models.Where{Column:"order", Value:_order, Compare:"="})    
    }
    _original := c.Geti64("original")
    if _original != 0 {
        args = append(args, models.Where{Column:"original", Value:_original, Compare:"="})    
    }
    _category := c.Geti64("category")
    if _category != 0 {
        args = append(args, models.Where{Column:"category", Value:_category, Compare:"="})    
    }
    _apt := c.Geti64("apt")
    if _apt != 0 {
        args = append(args, models.Where{Column:"apt", Value:_apt, Compare:"="})    
    }
    _standard := c.Geti64("standard")
    if _standard != 0 {
        args = append(args, models.Where{Column:"standard", Value:_standard, Compare:"="})    
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
                    str += ", sb_" + strings.Trim(v, " ")                
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





