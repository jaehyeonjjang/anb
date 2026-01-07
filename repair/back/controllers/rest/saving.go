package rest


import (
	"repair/controllers"
	"repair/models"

    "strings"
)

type SavingController struct {
	controllers.Controller
}



func (c *SavingController) CountByApt(apt int64) int {
    
    conn := c.NewConnection()

	_manager := models.NewSavingManager(conn)
    
    item := _manager.CountByApt(apt)
    
    
    
    c.Set("count", item)
    
    return item
    
}


func (c *SavingController) FindByApt(apt int64) []models.Saving {
    
    conn := c.NewConnection()

	_manager := models.NewSavingManager(conn)
    
    item := _manager.FindByApt(apt)
    
    
    c.Set("items", item)
    
    
    return item
    
}

// @Delete()
func (c *SavingController) DeleteByApt(apt int64) {
    
    conn := c.NewConnection()

	_manager := models.NewSavingManager(conn)
    
    err := _manager.DeleteByApt(apt)
    if err != nil {
        c.Set("code", "error")    
        c.Set("error", err)
        return
    }
    
}


func (c *SavingController) Insert(item *models.Saving) {
    
    
	conn := c.NewConnection()
    
	manager := models.NewSavingManager(conn)
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

func (c *SavingController) Insertbatch(item *[]models.Saving) {  
    if item == nil || len(*item) == 0 {
        return
    }

    rows := len(*item)
    
    
    
	conn := c.NewConnection()
    
	manager := models.NewSavingManager(conn)

    for i := 0; i < rows; i++ {
	    err := manager.Insert(&((*item)[i]))
        if err != nil {
            c.Set("code", "error")    
            c.Set("error", err)
            return
        }
    }
}

func (c *SavingController) Update(item *models.Saving) {
    
    
	conn := c.NewConnection()

	manager := models.NewSavingManager(conn)
    err := manager.Update(item)
    if err != nil {
        c.Set("code", "error")    
        c.Set("error", err)
        return
    }
}

func (c *SavingController) Delete(item *models.Saving) {
    
    
    conn := c.NewConnection()

	manager := models.NewSavingManager(conn)

    
	err := manager.Delete(item.Id)
    if err != nil {
        c.Set("code", "error")    
        c.Set("error", err)
    }
}

func (c *SavingController) Deletebatch(item *[]models.Saving) {
    
    
    conn := c.NewConnection()

	manager := models.NewSavingManager(conn)

    for _, v := range *item {
        
    
	    err := manager.Delete(v.Id)
        if err != nil {
            c.Set("code", "error")    
            c.Set("error", err)
            return
        }
    }
}

func (c *SavingController) Count() {
    
    
	conn := c.NewConnection()

	manager := models.NewSavingManager(conn)

    var args []interface{}
    
    _year := c.Geti("year")
    if _year != 0 {
        args = append(args, models.Where{Column:"year", Value:_year, Compare:"="})    
    }
    _forward := c.Geti64("forward")
    if _forward != 0 {
        args = append(args, models.Where{Column:"forward", Value:_forward, Compare:"="})    
    }
    _interest := c.Geti64("interest")
    if _interest != 0 {
        args = append(args, models.Where{Column:"interest", Value:_interest, Compare:"="})    
    }
    _surplus := c.Geti64("surplus")
    if _surplus != 0 {
        args = append(args, models.Where{Column:"surplus", Value:_surplus, Compare:"="})    
    }
    _saving := c.Geti64("saving")
    if _saving != 0 {
        args = append(args, models.Where{Column:"saving", Value:_saving, Compare:"="})    
    }
    _etc := c.Geti64("etc")
    if _etc != 0 {
        args = append(args, models.Where{Column:"etc", Value:_etc, Compare:"="})    
    }
    _use := c.Geti64("use")
    if _use != 0 {
        args = append(args, models.Where{Column:"use", Value:_use, Compare:"="})    
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


func (c *SavingController) Read(id int64) {
    
    
	conn := c.NewConnection()

	manager := models.NewSavingManager(conn)
	item := manager.Get(id)

    
    
    c.Set("item", item)
}

func (c *SavingController) Index(page int, pagesize int) {
    
    
	conn := c.NewConnection()

	manager := models.NewSavingManager(conn)

    var args []interface{}
    
    _year := c.Geti("year")
    if _year != 0 {
        args = append(args, models.Where{Column:"year", Value:_year, Compare:"="})    
    }
    _forward := c.Geti64("forward")
    if _forward != 0 {
        args = append(args, models.Where{Column:"forward", Value:_forward, Compare:"="})    
    }
    _interest := c.Geti64("interest")
    if _interest != 0 {
        args = append(args, models.Where{Column:"interest", Value:_interest, Compare:"="})    
    }
    _surplus := c.Geti64("surplus")
    if _surplus != 0 {
        args = append(args, models.Where{Column:"surplus", Value:_surplus, Compare:"="})    
    }
    _saving := c.Geti64("saving")
    if _saving != 0 {
        args = append(args, models.Where{Column:"saving", Value:_saving, Compare:"="})    
    }
    _etc := c.Geti64("etc")
    if _etc != 0 {
        args = append(args, models.Where{Column:"etc", Value:_etc, Compare:"="})    
    }
    _use := c.Geti64("use")
    if _use != 0 {
        args = append(args, models.Where{Column:"use", Value:_use, Compare:"="})    
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
                    str += ", sa_" + strings.Trim(v, " ")                
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





