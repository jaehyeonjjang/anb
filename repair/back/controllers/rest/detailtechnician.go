package rest


import (
	"repair/controllers"
	"repair/models"

    "strings"
)

type DetailtechnicianController struct {
	controllers.Controller
}



func (c *DetailtechnicianController) CountByDetail(detail int64) int {
    
    conn := c.NewConnection()

	_manager := models.NewDetailtechnicianManager(conn)
    
    item := _manager.CountByDetail(detail)
    
    
    
    c.Set("count", item)
    
    return item
    
}


func (c *DetailtechnicianController) FindByDetail(detail int64) []models.Detailtechnician {
    
    conn := c.NewConnection()

	_manager := models.NewDetailtechnicianManager(conn)
    
    item := _manager.FindByDetail(detail)
    
    
    c.Set("items", item)
    
    
    return item
    
}

// @Delete()
func (c *DetailtechnicianController) DeleteByDetail(detail int64) {
    
    conn := c.NewConnection()

	_manager := models.NewDetailtechnicianManager(conn)
    
    err := _manager.DeleteByDetail(detail)
    if err != nil {
        c.Set("code", "error")    
        c.Set("error", err)
        return
    }
    
}


func (c *DetailtechnicianController) Insert(item *models.Detailtechnician) {
    
    
	conn := c.NewConnection()
    
	manager := models.NewDetailtechnicianManager(conn)
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

func (c *DetailtechnicianController) Insertbatch(item *[]models.Detailtechnician) {  
    if item == nil || len(*item) == 0 {
        return
    }

    rows := len(*item)
    
    
    
	conn := c.NewConnection()
    
	manager := models.NewDetailtechnicianManager(conn)

    for i := 0; i < rows; i++ {
	    err := manager.Insert(&((*item)[i]))
        if err != nil {
            c.Set("code", "error")    
            c.Set("error", err)
            return
        }
    }
}

func (c *DetailtechnicianController) Update(item *models.Detailtechnician) {
    
    
	conn := c.NewConnection()

	manager := models.NewDetailtechnicianManager(conn)
    err := manager.Update(item)
    if err != nil {
        c.Set("code", "error")    
        c.Set("error", err)
        return
    }
}

func (c *DetailtechnicianController) Delete(item *models.Detailtechnician) {
    
    
    conn := c.NewConnection()

	manager := models.NewDetailtechnicianManager(conn)

    
	err := manager.Delete(item.Id)
    if err != nil {
        c.Set("code", "error")    
        c.Set("error", err)
    }
}

func (c *DetailtechnicianController) Deletebatch(item *[]models.Detailtechnician) {
    
    
    conn := c.NewConnection()

	manager := models.NewDetailtechnicianManager(conn)

    for _, v := range *item {
        
    
	    err := manager.Delete(v.Id)
        if err != nil {
            c.Set("code", "error")    
            c.Set("error", err)
            return
        }
    }
}

func (c *DetailtechnicianController) Count() {
    
    
	conn := c.NewConnection()

	manager := models.NewDetailtechnicianManager(conn)

    var args []interface{}
    
    _type := c.Geti("type")
    if _type != 0 {
        args = append(args, models.Where{Column:"type", Value:_type, Compare:"="})    
    }
    _part := c.Get("part")
    if _part != "" {
        args = append(args, models.Where{Column:"part", Value:_part, Compare:"="})
    }
    _startsignupstartdate := c.Get("startsignupstartdate")
    _endsignupstartdate := c.Get("endsignupstartdate")

    if _startsignupstartdate != "" && _endsignupstartdate != "" {        
        var v [2]string
        v[0] = _startsignupstartdate
        v[1] = _endsignupstartdate  
        args = append(args, models.Where{Column:"signupstartdate", Value:v, Compare:"between"})    
    } else if  _startsignupstartdate != "" {          
        args = append(args, models.Where{Column:"signupstartdate", Value:_startsignupstartdate, Compare:">="})
    } else if  _endsignupstartdate != "" {          
        args = append(args, models.Where{Column:"signupstartdate", Value:_endsignupstartdate, Compare:"<="})            
    }
    _startsignupenddate := c.Get("startsignupenddate")
    _endsignupenddate := c.Get("endsignupenddate")

    if _startsignupenddate != "" && _endsignupenddate != "" {        
        var v [2]string
        v[0] = _startsignupenddate
        v[1] = _endsignupenddate  
        args = append(args, models.Where{Column:"signupenddate", Value:v, Compare:"between"})    
    } else if  _startsignupenddate != "" {          
        args = append(args, models.Where{Column:"signupenddate", Value:_startsignupenddate, Compare:">="})
    } else if  _endsignupenddate != "" {          
        args = append(args, models.Where{Column:"signupenddate", Value:_endsignupenddate, Compare:"<="})            
    }
    _remark := c.Get("remark")
    if _remark != "" {
        args = append(args, models.Where{Column:"remark", Value:_remark, Compare:"="})
    }
    _order := c.Geti("order")
    if _order != 0 {
        args = append(args, models.Where{Column:"order", Value:_order, Compare:"="})    
    }
    _technician := c.Geti64("technician")
    if _technician != 0 {
        args = append(args, models.Where{Column:"technician", Value:_technician, Compare:"="})    
    }
    _detail := c.Geti64("detail")
    if _detail != 0 {
        args = append(args, models.Where{Column:"detail", Value:_detail, Compare:"="})    
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


func (c *DetailtechnicianController) Read(id int64) {
    
    
	conn := c.NewConnection()

	manager := models.NewDetailtechnicianManager(conn)
	item := manager.Get(id)

    
    
    c.Set("item", item)
}

func (c *DetailtechnicianController) Index(page int, pagesize int) {
    
    
	conn := c.NewConnection()

	manager := models.NewDetailtechnicianManager(conn)

    var args []interface{}
    
    _type := c.Geti("type")
    if _type != 0 {
        args = append(args, models.Where{Column:"type", Value:_type, Compare:"="})    
    }
    _part := c.Get("part")
    if _part != "" {
        args = append(args, models.Where{Column:"part", Value:_part, Compare:"="})
    }
    _startsignupstartdate := c.Get("startsignupstartdate")
    _endsignupstartdate := c.Get("endsignupstartdate")
    if _startsignupstartdate != "" && _endsignupstartdate != "" {        
        var v [2]string
        v[0] = _startsignupstartdate
        v[1] = _endsignupstartdate  
        args = append(args, models.Where{Column:"signupstartdate", Value:v, Compare:"between"})    
    } else if  _startsignupstartdate != "" {          
        args = append(args, models.Where{Column:"signupstartdate", Value:_startsignupstartdate, Compare:">="})
    } else if  _endsignupstartdate != "" {          
        args = append(args, models.Where{Column:"signupstartdate", Value:_endsignupstartdate, Compare:"<="})            
    }
    _startsignupenddate := c.Get("startsignupenddate")
    _endsignupenddate := c.Get("endsignupenddate")
    if _startsignupenddate != "" && _endsignupenddate != "" {        
        var v [2]string
        v[0] = _startsignupenddate
        v[1] = _endsignupenddate  
        args = append(args, models.Where{Column:"signupenddate", Value:v, Compare:"between"})    
    } else if  _startsignupenddate != "" {          
        args = append(args, models.Where{Column:"signupenddate", Value:_startsignupenddate, Compare:">="})
    } else if  _endsignupenddate != "" {          
        args = append(args, models.Where{Column:"signupenddate", Value:_endsignupenddate, Compare:"<="})            
    }
    _remark := c.Get("remark")
    if _remark != "" {
        args = append(args, models.Where{Column:"remark", Value:_remark, Compare:"="})
    }
    _order := c.Geti("order")
    if _order != 0 {
        args = append(args, models.Where{Column:"order", Value:_order, Compare:"="})    
    }
    _technician := c.Geti64("technician")
    if _technician != 0 {
        args = append(args, models.Where{Column:"technician", Value:_technician, Compare:"="})    
    }
    _detail := c.Geti64("detail")
    if _detail != 0 {
        args = append(args, models.Where{Column:"detail", Value:_detail, Compare:"="})    
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
                    str += ", dt_" + strings.Trim(v, " ")                
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





