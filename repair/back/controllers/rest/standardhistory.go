package rest


import (
	"repair/controllers"
	"repair/models"

    "strings"
)

type StandardhistoryController struct {
	controllers.Controller
}



func (c *StandardhistoryController) CountByApt(apt int64) int {
    
    conn := c.NewConnection()

	_manager := models.NewStandardhistoryManager(conn)
    
    item := _manager.CountByApt(apt)
    
    
    
    c.Set("count", item)
    
    return item
    
}


func (c *StandardhistoryController) FindByApt(apt int64) []models.Standardhistory {
    
    conn := c.NewConnection()

	_manager := models.NewStandardhistoryManager(conn)
    
    item := _manager.FindByApt(apt)
    
    
    c.Set("items", item)
    
    
    return item
    
}

// @Delete()
func (c *StandardhistoryController) DeleteByApt(apt int64) {
    
    conn := c.NewConnection()

	_manager := models.NewStandardhistoryManager(conn)
    
    err := _manager.DeleteByApt(apt)
    if err != nil {
        c.Set("code", "error")    
        c.Set("error", err)
        return
    }
    
}


func (c *StandardhistoryController) CountByStandard(standard int64) int {
    
    conn := c.NewConnection()

	_manager := models.NewStandardhistoryManager(conn)
    
    item := _manager.CountByStandard(standard)
    
    
    
    c.Set("count", item)
    
    return item
    
}


func (c *StandardhistoryController) GetByStandard(standard int64) *models.Standardhistory {
    
    conn := c.NewConnection()

	_manager := models.NewStandardhistoryManager(conn)
    
    item := _manager.GetByStandard(standard)
    
    c.Set("item", item)
    
    
    
    return item
    
}


func (c *StandardhistoryController) Insert(item *models.Standardhistory) {
    
    
	conn := c.NewConnection()
    
	manager := models.NewStandardhistoryManager(conn)
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

func (c *StandardhistoryController) Insertbatch(item *[]models.Standardhistory) {  
    if item == nil || len(*item) == 0 {
        return
    }

    rows := len(*item)
    
    
    
	conn := c.NewConnection()
    
	manager := models.NewStandardhistoryManager(conn)

    for i := 0; i < rows; i++ {
	    err := manager.Insert(&((*item)[i]))
        if err != nil {
            c.Set("code", "error")    
            c.Set("error", err)
            return
        }
    }
}

func (c *StandardhistoryController) Update(item *models.Standardhistory) {
    
    
	conn := c.NewConnection()

	manager := models.NewStandardhistoryManager(conn)
    err := manager.Update(item)
    if err != nil {
        c.Set("code", "error")    
        c.Set("error", err)
        return
    }
}

func (c *StandardhistoryController) Delete(item *models.Standardhistory) {
    
    
    conn := c.NewConnection()

	manager := models.NewStandardhistoryManager(conn)

    
	err := manager.Delete(item.Id)
    if err != nil {
        c.Set("code", "error")    
        c.Set("error", err)
    }
}

func (c *StandardhistoryController) Deletebatch(item *[]models.Standardhistory) {
    
    
    conn := c.NewConnection()

	manager := models.NewStandardhistoryManager(conn)

    for _, v := range *item {
        
    
	    err := manager.Delete(v.Id)
        if err != nil {
            c.Set("code", "error")    
            c.Set("error", err)
            return
        }
    }
}

func (c *StandardhistoryController) Count() {
    
    
	conn := c.NewConnection()

	manager := models.NewStandardhistoryManager(conn)

    var args []interface{}
    
    _name := c.Get("name")
    if _name != "" {
        args = append(args, models.Where{Column:"name", Value:_name, Compare:"="})
        
    }
    _direct := c.Geti("direct")
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
    _standard := c.Geti64("standard")
    if _standard != 0 {
        args = append(args, models.Where{Column:"standard", Value:_standard, Compare:"="})    
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


func (c *StandardhistoryController) Read(id int64) {
    
    
	conn := c.NewConnection()

	manager := models.NewStandardhistoryManager(conn)
	item := manager.Get(id)

    
    
    c.Set("item", item)
}

func (c *StandardhistoryController) Index(page int, pagesize int) {
    
    
	conn := c.NewConnection()

	manager := models.NewStandardhistoryManager(conn)

    var args []interface{}
    
    _name := c.Get("name")
    if _name != "" {
        args = append(args, models.Where{Column:"name", Value:_name, Compare:"="})
        
    }
    _direct := c.Geti("direct")
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
    _standard := c.Geti64("standard")
    if _standard != 0 {
        args = append(args, models.Where{Column:"standard", Value:_standard, Compare:"="})    
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
                    str += ", sh_" + strings.Trim(v, " ")                
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





