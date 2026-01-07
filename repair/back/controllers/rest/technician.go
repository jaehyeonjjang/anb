package rest


import (
	"repair/controllers"
	"repair/models"

    "strings"
)

type TechnicianController struct {
	controllers.Controller
}



func (c *TechnicianController) Insert(item *models.Technician) {
    
    
	conn := c.NewConnection()
    
	manager := models.NewTechnicianManager(conn)
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

func (c *TechnicianController) Insertbatch(item *[]models.Technician) {  
    if item == nil || len(*item) == 0 {
        return
    }

    rows := len(*item)
    
    
    
	conn := c.NewConnection()
    
	manager := models.NewTechnicianManager(conn)

    for i := 0; i < rows; i++ {
	    err := manager.Insert(&((*item)[i]))
        if err != nil {
            c.Set("code", "error")    
            c.Set("error", err)
            return
        }
    }
}

func (c *TechnicianController) Update(item *models.Technician) {
    
    
	conn := c.NewConnection()

	manager := models.NewTechnicianManager(conn)
    err := manager.Update(item)
    if err != nil {
        c.Set("code", "error")    
        c.Set("error", err)
        return
    }
}

func (c *TechnicianController) Delete(item *models.Technician) {
    
    
    conn := c.NewConnection()

	manager := models.NewTechnicianManager(conn)

    
	err := manager.Delete(item.Id)
    if err != nil {
        c.Set("code", "error")    
        c.Set("error", err)
    }
}

func (c *TechnicianController) Deletebatch(item *[]models.Technician) {
    
    
    conn := c.NewConnection()

	manager := models.NewTechnicianManager(conn)

    for _, v := range *item {
        
    
	    err := manager.Delete(v.Id)
        if err != nil {
            c.Set("code", "error")    
            c.Set("error", err)
            return
        }
    }
}

func (c *TechnicianController) Count() {
    
    
	conn := c.NewConnection()

	manager := models.NewTechnicianManager(conn)

    var args []interface{}
    
    _name := c.Get("name")
    if _name != "" {
        args = append(args, models.Where{Column:"name", Value:_name, Compare:"like"})
        
    }
    _grade := c.Geti("grade")
    if _grade != 0 {
        args = append(args, models.Where{Column:"grade", Value:_grade, Compare:"="})    
    }
    _stamp := c.Get("stamp")
    if _stamp != "" {
        args = append(args, models.Where{Column:"stamp", Value:_stamp, Compare:"="})
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


func (c *TechnicianController) Read(id int64) {
    
    
	conn := c.NewConnection()

	manager := models.NewTechnicianManager(conn)
	item := manager.Get(id)

    
    
    c.Set("item", item)
}

func (c *TechnicianController) Index(page int, pagesize int) {
    
    
	conn := c.NewConnection()

	manager := models.NewTechnicianManager(conn)

    var args []interface{}
    
    _name := c.Get("name")
    if _name != "" {
        args = append(args, models.Where{Column:"name", Value:_name, Compare:"like"})
        
    }
    _grade := c.Geti("grade")
    if _grade != 0 {
        args = append(args, models.Where{Column:"grade", Value:_grade, Compare:"="})    
    }
    _stamp := c.Get("stamp")
    if _stamp != "" {
        args = append(args, models.Where{Column:"stamp", Value:_stamp, Compare:"="})
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
                    str += ", te_" + strings.Trim(v, " ")                
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





