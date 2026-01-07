package rest


import (
	"repair/controllers"
	"repair/models"

    "strings"
)

type CategoryController struct {
	controllers.Controller
}



func (c *CategoryController) GetByLevelName(level int ,name string) *models.Category {
    
    conn := c.NewConnection()

	_manager := models.NewCategoryManager(conn)
    
    item := _manager.GetByLevelName(level, name)
    
    c.Set("item", item)
    
    
    
    return item
    
}


func (c *CategoryController) GetByLevelParentName(level int ,parent int64 ,name string) *models.Category {
    
    conn := c.NewConnection()

	_manager := models.NewCategoryManager(conn)
    
    item := _manager.GetByLevelParentName(level, parent, name)
    
    c.Set("item", item)
    
    
    
    return item
    
}


func (c *CategoryController) CountByApt(apt int64) int {
    
    conn := c.NewConnection()

	_manager := models.NewCategoryManager(conn)
    
    item := _manager.CountByApt(apt)
    
    
    
    c.Set("count", item)
    
    return item
    
}


func (c *CategoryController) FindByApt(apt int64) []models.Category {
    
    conn := c.NewConnection()

	_manager := models.NewCategoryManager(conn)
    
    item := _manager.FindByApt(apt)
    
    
    c.Set("items", item)
    
    
    return item
    
}


func (c *CategoryController) CountByAptLevel(apt int64 ,level int) int {
    
    conn := c.NewConnection()

	_manager := models.NewCategoryManager(conn)
    
    item := _manager.CountByAptLevel(apt, level)
    
    
    
    c.Set("count", item)
    
    return item
    
}


func (c *CategoryController) FindByAptLevel(apt int64 ,level int) []models.Category {
    
    conn := c.NewConnection()

	_manager := models.NewCategoryManager(conn)
    
    item := _manager.FindByAptLevel(apt, level)
    
    
    c.Set("items", item)
    
    
    return item
    
}


func (c *CategoryController) GetByAptLevelParentName(apt int64 ,level int ,parent int64 ,name string) *models.Category {
    
    conn := c.NewConnection()

	_manager := models.NewCategoryManager(conn)
    
    item := _manager.GetByAptLevelParentName(apt, level, parent, name)
    
    c.Set("item", item)
    
    
    
    return item
    
}


func (c *CategoryController) GetByAptName(apt int64 ,name string) *models.Category {
    
    conn := c.NewConnection()

	_manager := models.NewCategoryManager(conn)
    
    item := _manager.GetByAptName(apt, name)
    
    c.Set("item", item)
    
    
    
    return item
    
}


func (c *CategoryController) CountByAptParent(apt int64 ,parent int64) int {
    
    conn := c.NewConnection()

	_manager := models.NewCategoryManager(conn)
    
    item := _manager.CountByAptParent(apt, parent)
    
    
    
    c.Set("count", item)
    
    return item
    
}


func (c *CategoryController) FindByAptParent(apt int64 ,parent int64) []models.Category {
    
    conn := c.NewConnection()

	_manager := models.NewCategoryManager(conn)
    
    item := _manager.FindByAptParent(apt, parent)
    
    
    c.Set("items", item)
    
    
    return item
    
}


func (c *CategoryController) FindByAptOrder(apt int64 ,order int) []models.Category {
    
    conn := c.NewConnection()

	_manager := models.NewCategoryManager(conn)
    
    item := _manager.FindByAptOrder(apt, order)
    
    
    c.Set("items", item)
    
    
    return item
    
}


func (c *CategoryController) Insert(item *models.Category) {
    
    
	conn := c.NewConnection()
    
	manager := models.NewCategoryManager(conn)
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

func (c *CategoryController) Insertbatch(item *[]models.Category) {  
    if item == nil || len(*item) == 0 {
        return
    }

    rows := len(*item)
    
    
    
	conn := c.NewConnection()
    
	manager := models.NewCategoryManager(conn)

    for i := 0; i < rows; i++ {
	    err := manager.Insert(&((*item)[i]))
        if err != nil {
            c.Set("code", "error")    
            c.Set("error", err)
            return
        }
    }
}

func (c *CategoryController) Update(item *models.Category) {
    
    
	conn := c.NewConnection()

	manager := models.NewCategoryManager(conn)
    err := manager.Update(item)
    if err != nil {
        c.Set("code", "error")    
        c.Set("error", err)
        return
    }
}

func (c *CategoryController) Delete(item *models.Category) {
    
    
    conn := c.NewConnection()

	manager := models.NewCategoryManager(conn)

    
	err := manager.Delete(item.Id)
    if err != nil {
        c.Set("code", "error")    
        c.Set("error", err)
    }
}

func (c *CategoryController) Deletebatch(item *[]models.Category) {
    
    
    conn := c.NewConnection()

	manager := models.NewCategoryManager(conn)

    for _, v := range *item {
        
    
	    err := manager.Delete(v.Id)
        if err != nil {
            c.Set("code", "error")    
            c.Set("error", err)
            return
        }
    }
}

func (c *CategoryController) Count() {
    
    
	conn := c.NewConnection()

	manager := models.NewCategoryManager(conn)

    var args []interface{}
    
    _name := c.Get("name")
    if _name != "" {
        args = append(args, models.Where{Column:"name", Value:_name, Compare:"like"})
        
    }
    _level := c.Geti("level")
    if _level != 0 {
        args = append(args, models.Where{Column:"level", Value:_level, Compare:"="})    
    }
    _parent := c.Geti64("parent")
    if _parent != 0 {
        args = append(args, models.Where{Column:"parent", Value:_parent, Compare:"="})    
    }
    _cycle := c.Geti("cycle")
    if _cycle != 0 {
        args = append(args, models.Where{Column:"cycle", Value:_cycle, Compare:"="})    
    }
    _percent := c.Geti("percent")
    if _percent != 0 {
        args = append(args, models.Where{Column:"percent", Value:_percent, Compare:"="})    
    }
    _unit := c.Get("unit")
    if _unit != "" {
        args = append(args, models.Where{Column:"unit", Value:_unit, Compare:"="})
    }
    _elevator := c.Geti("elevator")
    if _elevator != 0 {
        args = append(args, models.Where{Column:"elevator", Value:_elevator, Compare:"="})    
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


func (c *CategoryController) Read(id int64) {
    
    
	conn := c.NewConnection()

	manager := models.NewCategoryManager(conn)
	item := manager.Get(id)

    
    
    c.Set("item", item)
}

func (c *CategoryController) Index(page int, pagesize int) {
    
    
	conn := c.NewConnection()

	manager := models.NewCategoryManager(conn)

    var args []interface{}
    
    _name := c.Get("name")
    if _name != "" {
        args = append(args, models.Where{Column:"name", Value:_name, Compare:"like"})
        
    }
    _level := c.Geti("level")
    if _level != 0 {
        args = append(args, models.Where{Column:"level", Value:_level, Compare:"="})    
    }
    _parent := c.Geti64("parent")
    if _parent != 0 {
        args = append(args, models.Where{Column:"parent", Value:_parent, Compare:"="})    
    }
    _cycle := c.Geti("cycle")
    if _cycle != 0 {
        args = append(args, models.Where{Column:"cycle", Value:_cycle, Compare:"="})    
    }
    _percent := c.Geti("percent")
    if _percent != 0 {
        args = append(args, models.Where{Column:"percent", Value:_percent, Compare:"="})    
    }
    _unit := c.Get("unit")
    if _unit != "" {
        args = append(args, models.Where{Column:"unit", Value:_unit, Compare:"="})
    }
    _elevator := c.Geti("elevator")
    if _elevator != 0 {
        args = append(args, models.Where{Column:"elevator", Value:_elevator, Compare:"="})    
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
                    str += ", c_" + strings.Trim(v, " ")                
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





