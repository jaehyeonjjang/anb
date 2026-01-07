package rest


import (
	"repair/controllers"
	"repair/models"

    "strings"
)

type AreaController struct {
	controllers.Controller
}



func (c *AreaController) CountByApt(apt int64) int {
    
    conn := c.NewConnection()

	_manager := models.NewAreaManager(conn)
    
    item := _manager.CountByApt(apt)
    
    
    
    c.Set("count", item)
    
    return item
    
}


func (c *AreaController) FindByApt(apt int64) []models.Area {
    
    conn := c.NewConnection()

	_manager := models.NewAreaManager(conn)
    
    item := _manager.FindByApt(apt)
    
    
    c.Set("items", item)
    
    
    return item
    
}

// @Delete()
func (c *AreaController) DeleteByApt(apt int64) {
    
    conn := c.NewConnection()

	_manager := models.NewAreaManager(conn)
    
    err := _manager.DeleteByApt(apt)
    if err != nil {
        c.Set("code", "error")    
        c.Set("error", err)
        return
    }
    
}


func (c *AreaController) Insert(item *models.Area) {
    
    
	conn := c.NewConnection()
    
	manager := models.NewAreaManager(conn)
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

func (c *AreaController) Insertbatch(item *[]models.Area) {  
    if item == nil || len(*item) == 0 {
        return
    }

    rows := len(*item)
    
    
    
	conn := c.NewConnection()
    
	manager := models.NewAreaManager(conn)

    for i := 0; i < rows; i++ {
	    err := manager.Insert(&((*item)[i]))
        if err != nil {
            c.Set("code", "error")    
            c.Set("error", err)
            return
        }
    }
}

func (c *AreaController) Update(item *models.Area) {
    
    
	conn := c.NewConnection()

	manager := models.NewAreaManager(conn)
    err := manager.Update(item)
    if err != nil {
        c.Set("code", "error")    
        c.Set("error", err)
        return
    }
}

func (c *AreaController) Delete(item *models.Area) {
    
    
    conn := c.NewConnection()

	manager := models.NewAreaManager(conn)

    
	err := manager.Delete(item.Id)
    if err != nil {
        c.Set("code", "error")    
        c.Set("error", err)
    }
}

func (c *AreaController) Deletebatch(item *[]models.Area) {
    
    
    conn := c.NewConnection()

	manager := models.NewAreaManager(conn)

    for _, v := range *item {
        
    
	    err := manager.Delete(v.Id)
        if err != nil {
            c.Set("code", "error")    
            c.Set("error", err)
            return
        }
    }
}

func (c *AreaController) Count() {
    
    
	conn := c.NewConnection()

	manager := models.NewAreaManager(conn)

    var args []interface{}
    
    _name := c.Get("name")
    if _name != "" {
        args = append(args, models.Where{Column:"name", Value:_name, Compare:"like"})
        
    }
    _familycount := c.Geti("familycount")
    if _familycount != 0 {
        args = append(args, models.Where{Column:"familycount", Value:_familycount, Compare:"="})    
    }
    _size := c.Geti("size")
    if _size != 0 {
        args = append(args, models.Where{Column:"size", Value:_size, Compare:"="})    
    }
    _order := c.Geti("order")
    if _order != 0 {
        args = append(args, models.Where{Column:"order", Value:_order, Compare:"="})    
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


func (c *AreaController) Read(id int64) {
    
    
	conn := c.NewConnection()

	manager := models.NewAreaManager(conn)
	item := manager.Get(id)

    
    
    c.Set("item", item)
}

func (c *AreaController) Index(page int, pagesize int) {
    
    
	conn := c.NewConnection()

	manager := models.NewAreaManager(conn)

    var args []interface{}
    
    _name := c.Get("name")
    if _name != "" {
        args = append(args, models.Where{Column:"name", Value:_name, Compare:"like"})
        
    }
    _familycount := c.Geti("familycount")
    if _familycount != 0 {
        args = append(args, models.Where{Column:"familycount", Value:_familycount, Compare:"="})    
    }
    _size := c.Geti("size")
    if _size != 0 {
        args = append(args, models.Where{Column:"size", Value:_size, Compare:"="})    
    }
    _order := c.Geti("order")
    if _order != 0 {
        args = append(args, models.Where{Column:"order", Value:_order, Compare:"="})    
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
                    str += ", ar_" + strings.Trim(v, " ")                
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





