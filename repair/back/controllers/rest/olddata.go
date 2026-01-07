package rest


import (
	"repair/controllers"
	"repair/models"

    "strings"
)

type OlddataController struct {
	controllers.Controller
}



func (c *OlddataController) Insert(item *models.Olddata) {
    
    
	conn := c.NewConnection()
    
	manager := models.NewOlddataManager(conn)
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

func (c *OlddataController) Insertbatch(item *[]models.Olddata) {  
    if item == nil || len(*item) == 0 {
        return
    }

    rows := len(*item)
    
    
    
	conn := c.NewConnection()
    
	manager := models.NewOlddataManager(conn)

    for i := 0; i < rows; i++ {
	    err := manager.Insert(&((*item)[i]))
        if err != nil {
            c.Set("code", "error")    
            c.Set("error", err)
            return
        }
    }
}

func (c *OlddataController) Update(item *models.Olddata) {
    
    
	conn := c.NewConnection()

	manager := models.NewOlddataManager(conn)
    err := manager.Update(item)
    if err != nil {
        c.Set("code", "error")    
        c.Set("error", err)
        return
    }
}

func (c *OlddataController) Delete(item *models.Olddata) {
    
    
    conn := c.NewConnection()

	manager := models.NewOlddataManager(conn)

    
	err := manager.Delete(item.Id)
    if err != nil {
        c.Set("code", "error")    
        c.Set("error", err)
    }
}

func (c *OlddataController) Deletebatch(item *[]models.Olddata) {
    
    
    conn := c.NewConnection()

	manager := models.NewOlddataManager(conn)

    for _, v := range *item {
        
    
	    err := manager.Delete(v.Id)
        if err != nil {
            c.Set("code", "error")    
            c.Set("error", err)
            return
        }
    }
}

func (c *OlddataController) Count() {
    
    
	conn := c.NewConnection()

	manager := models.NewOlddataManager(conn)

    var args []interface{}
    
    _apt := c.Geti64("apt")
    if _apt != 0 {
        args = append(args, models.Where{Column:"apt", Value:_apt, Compare:"="})    
    }
    _image := c.Geti64("image")
    if _image != 0 {
        args = append(args, models.Where{Column:"image", Value:_image, Compare:"="})    
    }
    _imagetype := c.Geti("imagetype")
    if _imagetype != 0 {
        args = append(args, models.Where{Column:"imagetype", Value:_imagetype, Compare:"="})    
    }
    _user := c.Geti64("user")
    if _user != 0 {
        args = append(args, models.Where{Column:"user", Value:_user, Compare:"="})    
    }
    _type := c.Geti("type")
    if _type != 0 {
        args = append(args, models.Where{Column:"type", Value:_type, Compare:"="})    
    }
    _x := c.Geti("x")
    if _x != 0 {
        args = append(args, models.Where{Column:"x", Value:_x, Compare:"="})    
    }
    _y := c.Geti("y")
    if _y != 0 {
        args = append(args, models.Where{Column:"y", Value:_y, Compare:"="})    
    }
    _point := c.Get("point")
    if _point != "" {
        args = append(args, models.Where{Column:"point", Value:_point, Compare:"="})
    }
    _number := c.Geti("number")
    if _number != 0 {
        args = append(args, models.Where{Column:"number", Value:_number, Compare:"="})    
    }
    _group := c.Geti("group")
    if _group != 0 {
        args = append(args, models.Where{Column:"group", Value:_group, Compare:"="})    
    }
    _name := c.Get("name")
    if _name != "" {
        args = append(args, models.Where{Column:"name", Value:_name, Compare:"="})
        
    }
    _fault := c.Get("fault")
    if _fault != "" {
        args = append(args, models.Where{Column:"fault", Value:_fault, Compare:"="})
    }
    _content := c.Get("content")
    if _content != "" {
        args = append(args, models.Where{Column:"content", Value:_content, Compare:"="})
        
    }
    _width := c.Geti("width")
    if _width != 0 {
        args = append(args, models.Where{Column:"width", Value:_width, Compare:"="})    
    }
    _length := c.Geti("length")
    if _length != 0 {
        args = append(args, models.Where{Column:"length", Value:_length, Compare:"="})    
    }
    _count := c.Get("count")
    if _count != "" {
        args = append(args, models.Where{Column:"count", Value:_count, Compare:"="})
    }
    _progress := c.Get("progress")
    if _progress != "" {
        args = append(args, models.Where{Column:"progress", Value:_progress, Compare:"="})
    }
    _remark := c.Get("remark")
    if _remark != "" {
        args = append(args, models.Where{Column:"remark", Value:_remark, Compare:"="})
    }
    _imagename := c.Get("imagename")
    if _imagename != "" {
        args = append(args, models.Where{Column:"imagename", Value:_imagename, Compare:"="})
    }
    _filename := c.Get("filename")
    if _filename != "" {
        args = append(args, models.Where{Column:"filename", Value:_filename, Compare:"="})
    }
    _memo := c.Get("memo")
    if _memo != "" {
        args = append(args, models.Where{Column:"memo", Value:_memo, Compare:"="})
    }
    _report := c.Geti("report")
    if _report != 0 {
        args = append(args, models.Where{Column:"report", Value:_report, Compare:"="})    
    }
    _usermemo := c.Get("usermemo")
    if _usermemo != "" {
        args = append(args, models.Where{Column:"usermemo", Value:_usermemo, Compare:"="})
    }
    _aptmemo := c.Get("aptmemo")
    if _aptmemo != "" {
        args = append(args, models.Where{Column:"aptmemo", Value:_aptmemo, Compare:"="})
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


func (c *OlddataController) Read(id int64) {
    
    
	conn := c.NewConnection()

	manager := models.NewOlddataManager(conn)
	item := manager.Get(id)

    
    
    c.Set("item", item)
}

func (c *OlddataController) Index(page int, pagesize int) {
    
    
	conn := c.NewConnection()

	manager := models.NewOlddataManager(conn)

    var args []interface{}
    
    _apt := c.Geti64("apt")
    if _apt != 0 {
        args = append(args, models.Where{Column:"apt", Value:_apt, Compare:"="})    
    }
    _image := c.Geti64("image")
    if _image != 0 {
        args = append(args, models.Where{Column:"image", Value:_image, Compare:"="})    
    }
    _imagetype := c.Geti("imagetype")
    if _imagetype != 0 {
        args = append(args, models.Where{Column:"imagetype", Value:_imagetype, Compare:"="})    
    }
    _user := c.Geti64("user")
    if _user != 0 {
        args = append(args, models.Where{Column:"user", Value:_user, Compare:"="})    
    }
    _type := c.Geti("type")
    if _type != 0 {
        args = append(args, models.Where{Column:"type", Value:_type, Compare:"="})    
    }
    _x := c.Geti("x")
    if _x != 0 {
        args = append(args, models.Where{Column:"x", Value:_x, Compare:"="})    
    }
    _y := c.Geti("y")
    if _y != 0 {
        args = append(args, models.Where{Column:"y", Value:_y, Compare:"="})    
    }
    _point := c.Get("point")
    if _point != "" {
        args = append(args, models.Where{Column:"point", Value:_point, Compare:"="})
    }
    _number := c.Geti("number")
    if _number != 0 {
        args = append(args, models.Where{Column:"number", Value:_number, Compare:"="})    
    }
    _group := c.Geti("group")
    if _group != 0 {
        args = append(args, models.Where{Column:"group", Value:_group, Compare:"="})    
    }
    _name := c.Get("name")
    if _name != "" {
        args = append(args, models.Where{Column:"name", Value:_name, Compare:"="})
        
    }
    _fault := c.Get("fault")
    if _fault != "" {
        args = append(args, models.Where{Column:"fault", Value:_fault, Compare:"="})
    }
    _content := c.Get("content")
    if _content != "" {
        args = append(args, models.Where{Column:"content", Value:_content, Compare:"="})
        
    }
    _width := c.Geti("width")
    if _width != 0 {
        args = append(args, models.Where{Column:"width", Value:_width, Compare:"="})    
    }
    _length := c.Geti("length")
    if _length != 0 {
        args = append(args, models.Where{Column:"length", Value:_length, Compare:"="})    
    }
    _count := c.Get("count")
    if _count != "" {
        args = append(args, models.Where{Column:"count", Value:_count, Compare:"="})
    }
    _progress := c.Get("progress")
    if _progress != "" {
        args = append(args, models.Where{Column:"progress", Value:_progress, Compare:"="})
    }
    _remark := c.Get("remark")
    if _remark != "" {
        args = append(args, models.Where{Column:"remark", Value:_remark, Compare:"="})
    }
    _imagename := c.Get("imagename")
    if _imagename != "" {
        args = append(args, models.Where{Column:"imagename", Value:_imagename, Compare:"="})
    }
    _filename := c.Get("filename")
    if _filename != "" {
        args = append(args, models.Where{Column:"filename", Value:_filename, Compare:"="})
    }
    _memo := c.Get("memo")
    if _memo != "" {
        args = append(args, models.Where{Column:"memo", Value:_memo, Compare:"="})
    }
    _report := c.Geti("report")
    if _report != 0 {
        args = append(args, models.Where{Column:"report", Value:_report, Compare:"="})    
    }
    _usermemo := c.Get("usermemo")
    if _usermemo != "" {
        args = append(args, models.Where{Column:"usermemo", Value:_usermemo, Compare:"="})
    }
    _aptmemo := c.Get("aptmemo")
    if _aptmemo != "" {
        args = append(args, models.Where{Column:"aptmemo", Value:_aptmemo, Compare:"="})
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
                    str += ", d_" + strings.Trim(v, " ")                
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





func (c *OlddataController) Sum() {
    
    
	conn := c.NewConnection()

	manager := models.NewOlddataManager(conn)

    var args []interface{}
    
    _apt := c.Geti64("apt")
    if _apt != 0 {
        args = append(args, models.Where{Column:"apt", Value:_apt, Compare:"="})    
    }
    _image := c.Geti64("image")
    if _image != 0 {
        args = append(args, models.Where{Column:"image", Value:_image, Compare:"="})    
    }
    _imagetype := c.Geti("imagetype")
    if _imagetype != 0 {
        args = append(args, models.Where{Column:"imagetype", Value:_imagetype, Compare:"="})    
    }
    _user := c.Geti64("user")
    if _user != 0 {
        args = append(args, models.Where{Column:"user", Value:_user, Compare:"="})    
    }
    _type := c.Geti("type")
    if _type != 0 {
        args = append(args, models.Where{Column:"type", Value:_type, Compare:"="})    
    }
    _x := c.Geti("x")
    if _x != 0 {
        args = append(args, models.Where{Column:"x", Value:_x, Compare:"="})    
    }
    _y := c.Geti("y")
    if _y != 0 {
        args = append(args, models.Where{Column:"y", Value:_y, Compare:"="})    
    }
    _point := c.Get("point")
    if _point != "" {
        args = append(args, models.Where{Column:"point", Value:_point, Compare:"like"})
    }
    _number := c.Geti("number")
    if _number != 0 {
        args = append(args, models.Where{Column:"number", Value:_number, Compare:"="})    
    }
    _group := c.Geti("group")
    if _group != 0 {
        args = append(args, models.Where{Column:"group", Value:_group, Compare:"="})    
    }
    _name := c.Get("name")
    if _name != "" {
        args = append(args, models.Where{Column:"name", Value:_name, Compare:"="})
        
    }
    _fault := c.Get("fault")
    if _fault != "" {
        args = append(args, models.Where{Column:"fault", Value:_fault, Compare:"like"})
    }
    _content := c.Get("content")
    if _content != "" {
        args = append(args, models.Where{Column:"content", Value:_content, Compare:"="})
        
    }
    _width := c.Geti("width")
    if _width != 0 {
        args = append(args, models.Where{Column:"width", Value:_width, Compare:"="})    
    }
    _length := c.Geti("length")
    if _length != 0 {
        args = append(args, models.Where{Column:"length", Value:_length, Compare:"="})    
    }
    _count := c.Get("count")
    if _count != "" {
        args = append(args, models.Where{Column:"count", Value:_count, Compare:"like"})
    }
    _progress := c.Get("progress")
    if _progress != "" {
        args = append(args, models.Where{Column:"progress", Value:_progress, Compare:"like"})
    }
    _remark := c.Get("remark")
    if _remark != "" {
        args = append(args, models.Where{Column:"remark", Value:_remark, Compare:"like"})
    }
    _imagename := c.Get("imagename")
    if _imagename != "" {
        args = append(args, models.Where{Column:"imagename", Value:_imagename, Compare:"like"})
    }
    _filename := c.Get("filename")
    if _filename != "" {
        args = append(args, models.Where{Column:"filename", Value:_filename, Compare:"like"})
    }
    _memo := c.Get("memo")
    if _memo != "" {
        args = append(args, models.Where{Column:"memo", Value:_memo, Compare:"like"})
    }
    _report := c.Geti("report")
    if _report != 0 {
        args = append(args, models.Where{Column:"report", Value:_report, Compare:"="})    
    }
    _usermemo := c.Get("usermemo")
    if _usermemo != "" {
        args = append(args, models.Where{Column:"usermemo", Value:_usermemo, Compare:"like"})
    }
    _aptmemo := c.Get("aptmemo")
    if _aptmemo != "" {
        args = append(args, models.Where{Column:"aptmemo", Value:_aptmemo, Compare:"like"})
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

