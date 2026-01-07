package rest


import (
	"repair/controllers"
	"repair/models"

    "strings"
)

type OldimageController struct {
	controllers.Controller
}



func (c *OldimageController) Insert(item *models.Oldimage) {
    
    
	conn := c.NewConnection()
    
	manager := models.NewOldimageManager(conn)
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

func (c *OldimageController) Insertbatch(item *[]models.Oldimage) {  
    if item == nil || len(*item) == 0 {
        return
    }

    rows := len(*item)
    
    
    
	conn := c.NewConnection()
    
	manager := models.NewOldimageManager(conn)

    for i := 0; i < rows; i++ {
	    err := manager.Insert(&((*item)[i]))
        if err != nil {
            c.Set("code", "error")    
            c.Set("error", err)
            return
        }
    }
}

func (c *OldimageController) Update(item *models.Oldimage) {
    
    
	conn := c.NewConnection()

	manager := models.NewOldimageManager(conn)
    err := manager.Update(item)
    if err != nil {
        c.Set("code", "error")    
        c.Set("error", err)
        return
    }
}

func (c *OldimageController) Delete(item *models.Oldimage) {
    
    
    conn := c.NewConnection()

	manager := models.NewOldimageManager(conn)

    
	err := manager.Delete(item.Id)
    if err != nil {
        c.Set("code", "error")    
        c.Set("error", err)
    }
}

func (c *OldimageController) Deletebatch(item *[]models.Oldimage) {
    
    
    conn := c.NewConnection()

	manager := models.NewOldimageManager(conn)

    for _, v := range *item {
        
    
	    err := manager.Delete(v.Id)
        if err != nil {
            c.Set("code", "error")    
            c.Set("error", err)
            return
        }
    }
}

func (c *OldimageController) Count() {
    
    
	conn := c.NewConnection()

	manager := models.NewOldimageManager(conn)

    var args []interface{}
    
    _apt := c.Geti64("apt")
    if _apt != 0 {
        args = append(args, models.Where{Column:"apt", Value:_apt, Compare:"="})    
    }
    _name := c.Get("name")
    if _name != "" {
        args = append(args, models.Where{Column:"name", Value:_name, Compare:"="})
        
    }
    _level := c.Geti("level")
    if _level != 0 {
        args = append(args, models.Where{Column:"level", Value:_level, Compare:"="})    
    }
    _parent := c.Geti64("parent")
    if _parent != 0 {
        args = append(args, models.Where{Column:"parent", Value:_parent, Compare:"="})    
    }
    _last := c.Geti("last")
    if _last != 0 {
        args = append(args, models.Where{Column:"last", Value:_last, Compare:"="})    
    }
    _title := c.Get("title")
    if _title != "" {
        args = append(args, models.Where{Column:"title", Value:_title, Compare:"="})
        
    }
    _type := c.Geti("type")
    if _type != 0 {
        args = append(args, models.Where{Column:"type", Value:_type, Compare:"="})    
    }
    _floortype := c.Geti("floortype")
    if _floortype != 0 {
        args = append(args, models.Where{Column:"floortype", Value:_floortype, Compare:"="})    
    }
    _filename := c.Get("filename")
    if _filename != "" {
        args = append(args, models.Where{Column:"filename", Value:_filename, Compare:"="})
    }
    _order := c.Geti("order")
    if _order != 0 {
        args = append(args, models.Where{Column:"order", Value:_order, Compare:"="})    
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
    _standard := c.Geti("standard")
    if _standard != 0 {
        args = append(args, models.Where{Column:"standard", Value:_standard, Compare:"="})    
    }
    

    
    
    total := manager.Count(args)
	c.Set("total", total)
}


func (c *OldimageController) Read(id int64) {
    
    
	conn := c.NewConnection()

	manager := models.NewOldimageManager(conn)
	item := manager.Get(id)

    
    
    c.Set("item", item)
}

func (c *OldimageController) Index(page int, pagesize int) {
    
    
	conn := c.NewConnection()

	manager := models.NewOldimageManager(conn)

    var args []interface{}
    
    _apt := c.Geti64("apt")
    if _apt != 0 {
        args = append(args, models.Where{Column:"apt", Value:_apt, Compare:"="})    
    }
    _name := c.Get("name")
    if _name != "" {
        args = append(args, models.Where{Column:"name", Value:_name, Compare:"="})
        
    }
    _level := c.Geti("level")
    if _level != 0 {
        args = append(args, models.Where{Column:"level", Value:_level, Compare:"="})    
    }
    _parent := c.Geti64("parent")
    if _parent != 0 {
        args = append(args, models.Where{Column:"parent", Value:_parent, Compare:"="})    
    }
    _last := c.Geti("last")
    if _last != 0 {
        args = append(args, models.Where{Column:"last", Value:_last, Compare:"="})    
    }
    _title := c.Get("title")
    if _title != "" {
        args = append(args, models.Where{Column:"title", Value:_title, Compare:"="})
        
    }
    _type := c.Geti("type")
    if _type != 0 {
        args = append(args, models.Where{Column:"type", Value:_type, Compare:"="})    
    }
    _floortype := c.Geti("floortype")
    if _floortype != 0 {
        args = append(args, models.Where{Column:"floortype", Value:_floortype, Compare:"="})    
    }
    _filename := c.Get("filename")
    if _filename != "" {
        args = append(args, models.Where{Column:"filename", Value:_filename, Compare:"="})
    }
    _order := c.Geti("order")
    if _order != 0 {
        args = append(args, models.Where{Column:"order", Value:_order, Compare:"="})    
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
    _standard := c.Geti("standard")
    if _standard != 0 {
        args = append(args, models.Where{Column:"standard", Value:_standard, Compare:"="})    
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
                    str += ", i_" + strings.Trim(v, " ")                
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





