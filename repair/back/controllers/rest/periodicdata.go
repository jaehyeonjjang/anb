package rest


import (
	"repair/controllers"
	"repair/models"

    "strings"
)

type PeriodicdataController struct {
	controllers.Controller
}


// @Delete()
func (c *PeriodicdataController) DeleteByPeriodicBlueprint(periodic int64 ,blueprint int64) {
    
    conn := c.NewConnection()

	_manager := models.NewPeriodicdataManager(conn)
    
    err := _manager.DeleteByPeriodicBlueprint(periodic, blueprint)
    if err != nil {
        c.Set("code", "error")    
        c.Set("error", err)
        return
    }
    
}


func (c *PeriodicdataController) Insert(item *models.Periodicdata) {
    
    
	conn := c.NewConnection()
    
	manager := models.NewPeriodicdataManager(conn)
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

func (c *PeriodicdataController) Insertbatch(item *[]models.Periodicdata) {  
    if item == nil || len(*item) == 0 {
        return
    }

    rows := len(*item)
    
    
    
	conn := c.NewConnection()
    
	manager := models.NewPeriodicdataManager(conn)

    for i := 0; i < rows; i++ {
	    err := manager.Insert(&((*item)[i]))
        if err != nil {
            c.Set("code", "error")    
            c.Set("error", err)
            return
        }
    }
}

func (c *PeriodicdataController) Update(item *models.Periodicdata) {
    
    
	conn := c.NewConnection()

	manager := models.NewPeriodicdataManager(conn)
    err := manager.Update(item)
    if err != nil {
        c.Set("code", "error")    
        c.Set("error", err)
        return
    }
}

func (c *PeriodicdataController) Delete(item *models.Periodicdata) {
    
    
    conn := c.NewConnection()

	manager := models.NewPeriodicdataManager(conn)

    
	err := manager.Delete(item.Id)
    if err != nil {
        c.Set("code", "error")    
        c.Set("error", err)
    }
}

func (c *PeriodicdataController) Deletebatch(item *[]models.Periodicdata) {
    
    
    conn := c.NewConnection()

	manager := models.NewPeriodicdataManager(conn)

    for _, v := range *item {
        
    
	    err := manager.Delete(v.Id)
        if err != nil {
            c.Set("code", "error")    
            c.Set("error", err)
            return
        }
    }
}

func (c *PeriodicdataController) Count() {
    
    
	conn := c.NewConnection()

	manager := models.NewPeriodicdataManager(conn)

    var args []interface{}
    
    _group := c.Geti("group")
    if _group != 0 {
        args = append(args, models.Where{Column:"group", Value:_group, Compare:"="})    
    }
    _type := c.Geti("type")
    if _type != 0 {
        args = append(args, models.Where{Column:"type", Value:_type, Compare:"="})    
    }
    _part := c.Get("part")
    if _part != "" {
        args = append(args, models.Where{Column:"part", Value:_part, Compare:"="})
    }
    _member := c.Get("member")
    if _member != "" {
        args = append(args, models.Where{Column:"member", Value:_member, Compare:"="})
    }
    _shape := c.Get("shape")
    if _shape != "" {
        args = append(args, models.Where{Column:"shape", Value:_shape, Compare:"="})
    }
    _width := c.Get("width")
    if _width != "" {
        args = append(args, models.Where{Column:"width", Value:_width, Compare:"="})
    }
    _length := c.Get("length")
    if _length != "" {
        args = append(args, models.Where{Column:"length", Value:_length, Compare:"="})
    }
    _count := c.Geti("count")
    if _count != 0 {
        args = append(args, models.Where{Column:"count", Value:_count, Compare:"="})    
    }
    _progress := c.Geti("progress")
    if _progress != 0 {
        args = append(args, models.Where{Column:"progress", Value:_progress, Compare:"="})    
    }
    _remark := c.Get("remark")
    if _remark != "" {
        args = append(args, models.Where{Column:"remark", Value:_remark, Compare:"="})
    }
    _order := c.Geti("order")
    if _order != 0 {
        args = append(args, models.Where{Column:"order", Value:_order, Compare:"="})    
    }
    _content := c.Get("content")
    if _content != "" {
        args = append(args, models.Where{Column:"content", Value:_content, Compare:"="})
        
    }
    _status := c.Geti("status")
    if _status != 0 {
        args = append(args, models.Where{Column:"status", Value:_status, Compare:"="})    
    }
    _filename := c.Get("filename")
    if _filename != "" {
        args = append(args, models.Where{Column:"filename", Value:_filename, Compare:"="})
    }
    _offlinefilename := c.Get("offlinefilename")
    if _offlinefilename != "" {
        args = append(args, models.Where{Column:"offlinefilename", Value:_offlinefilename, Compare:"="})
    }
    _user := c.Geti64("user")
    if _user != 0 {
        args = append(args, models.Where{Column:"user", Value:_user, Compare:"="})    
    }
    _blueprint := c.Geti64("blueprint")
    if _blueprint != 0 {
        args = append(args, models.Where{Column:"blueprint", Value:_blueprint, Compare:"="})    
    }
    _periodic := c.Geti64("periodic")
    if _periodic != 0 {
        args = append(args, models.Where{Column:"periodic", Value:_periodic, Compare:"="})    
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


func (c *PeriodicdataController) Read(id int64) {
    
    
	conn := c.NewConnection()

	manager := models.NewPeriodicdataManager(conn)
	item := manager.Get(id)

    
    
    c.Set("item", item)
}

func (c *PeriodicdataController) Index(page int, pagesize int) {
    
    
	conn := c.NewConnection()

	manager := models.NewPeriodicdataManager(conn)

    var args []interface{}
    
    _group := c.Geti("group")
    if _group != 0 {
        args = append(args, models.Where{Column:"group", Value:_group, Compare:"="})    
    }
    _type := c.Geti("type")
    if _type != 0 {
        args = append(args, models.Where{Column:"type", Value:_type, Compare:"="})    
    }
    _part := c.Get("part")
    if _part != "" {
        args = append(args, models.Where{Column:"part", Value:_part, Compare:"="})
    }
    _member := c.Get("member")
    if _member != "" {
        args = append(args, models.Where{Column:"member", Value:_member, Compare:"="})
    }
    _shape := c.Get("shape")
    if _shape != "" {
        args = append(args, models.Where{Column:"shape", Value:_shape, Compare:"="})
    }
    _width := c.Get("width")
    if _width != "" {
        args = append(args, models.Where{Column:"width", Value:_width, Compare:"="})
    }
    _length := c.Get("length")
    if _length != "" {
        args = append(args, models.Where{Column:"length", Value:_length, Compare:"="})
    }
    _count := c.Geti("count")
    if _count != 0 {
        args = append(args, models.Where{Column:"count", Value:_count, Compare:"="})    
    }
    _progress := c.Geti("progress")
    if _progress != 0 {
        args = append(args, models.Where{Column:"progress", Value:_progress, Compare:"="})    
    }
    _remark := c.Get("remark")
    if _remark != "" {
        args = append(args, models.Where{Column:"remark", Value:_remark, Compare:"="})
    }
    _order := c.Geti("order")
    if _order != 0 {
        args = append(args, models.Where{Column:"order", Value:_order, Compare:"="})    
    }
    _content := c.Get("content")
    if _content != "" {
        args = append(args, models.Where{Column:"content", Value:_content, Compare:"="})
        
    }
    _status := c.Geti("status")
    if _status != 0 {
        args = append(args, models.Where{Column:"status", Value:_status, Compare:"="})    
    }
    _filename := c.Get("filename")
    if _filename != "" {
        args = append(args, models.Where{Column:"filename", Value:_filename, Compare:"="})
    }
    _offlinefilename := c.Get("offlinefilename")
    if _offlinefilename != "" {
        args = append(args, models.Where{Column:"offlinefilename", Value:_offlinefilename, Compare:"="})
    }
    _user := c.Geti64("user")
    if _user != 0 {
        args = append(args, models.Where{Column:"user", Value:_user, Compare:"="})    
    }
    _blueprint := c.Geti64("blueprint")
    if _blueprint != 0 {
        args = append(args, models.Where{Column:"blueprint", Value:_blueprint, Compare:"="})    
    }
    _periodic := c.Geti64("periodic")
    if _periodic != 0 {
        args = append(args, models.Where{Column:"periodic", Value:_periodic, Compare:"="})    
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





func (c *PeriodicdataController) Sum() {
    
    
	conn := c.NewConnection()

	manager := models.NewPeriodicdataManager(conn)

    var args []interface{}
    
    _group := c.Geti("group")
    if _group != 0 {
        args = append(args, models.Where{Column:"group", Value:_group, Compare:"="})    
    }
    _type := c.Geti("type")
    if _type != 0 {
        args = append(args, models.Where{Column:"type", Value:_type, Compare:"="})    
    }
    _part := c.Get("part")
    if _part != "" {
        args = append(args, models.Where{Column:"part", Value:_part, Compare:"like"})
    }
    _member := c.Get("member")
    if _member != "" {
        args = append(args, models.Where{Column:"member", Value:_member, Compare:"like"})
    }
    _shape := c.Get("shape")
    if _shape != "" {
        args = append(args, models.Where{Column:"shape", Value:_shape, Compare:"like"})
    }
    _width := c.Get("width")
    if _width != "" {
        args = append(args, models.Where{Column:"width", Value:_width, Compare:"like"})
    }
    _length := c.Get("length")
    if _length != "" {
        args = append(args, models.Where{Column:"length", Value:_length, Compare:"like"})
    }
    _count := c.Geti("count")
    if _count != 0 {
        args = append(args, models.Where{Column:"count", Value:_count, Compare:"="})    
    }
    _progress := c.Geti("progress")
    if _progress != 0 {
        args = append(args, models.Where{Column:"progress", Value:_progress, Compare:"="})    
    }
    _remark := c.Get("remark")
    if _remark != "" {
        args = append(args, models.Where{Column:"remark", Value:_remark, Compare:"like"})
    }
    _order := c.Geti("order")
    if _order != 0 {
        args = append(args, models.Where{Column:"order", Value:_order, Compare:"="})    
    }
    _content := c.Get("content")
    if _content != "" {
        args = append(args, models.Where{Column:"content", Value:_content, Compare:"="})
        
    }
    _status := c.Geti("status")
    if _status != 0 {
        args = append(args, models.Where{Column:"status", Value:_status, Compare:"="})    
    }
    _filename := c.Get("filename")
    if _filename != "" {
        args = append(args, models.Where{Column:"filename", Value:_filename, Compare:"like"})
    }
    _offlinefilename := c.Get("offlinefilename")
    if _offlinefilename != "" {
        args = append(args, models.Where{Column:"offlinefilename", Value:_offlinefilename, Compare:"like"})
    }
    _user := c.Geti64("user")
    if _user != 0 {
        args = append(args, models.Where{Column:"user", Value:_user, Compare:"="})    
    }
    _blueprint := c.Geti64("blueprint")
    if _blueprint != 0 {
        args = append(args, models.Where{Column:"blueprint", Value:_blueprint, Compare:"="})    
    }
    _periodic := c.Geti64("periodic")
    if _periodic != 0 {
        args = append(args, models.Where{Column:"periodic", Value:_periodic, Compare:"="})    
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

