package rest


import (
	"repair/controllers"
	"repair/models"

    "strings"
)

type FileController struct {
	controllers.Controller
}



func (c *FileController) CountByApt(apt int64) int {
    
    conn := c.NewConnection()

	_manager := models.NewFileManager(conn)
    
    item := _manager.CountByApt(apt)
    
    
    
    c.Set("count", item)
    
    return item
    
}


func (c *FileController) FindByApt(apt int64) []models.File {
    
    conn := c.NewConnection()

	_manager := models.NewFileManager(conn)
    
    item := _manager.FindByApt(apt)
    
    
    c.Set("items", item)
    
    
    return item
    
}

// @Delete()
func (c *FileController) DeleteByApt(apt int64) {
    
    conn := c.NewConnection()

	_manager := models.NewFileManager(conn)
    
    err := _manager.DeleteByApt(apt)
    if err != nil {
        c.Set("code", "error")    
        c.Set("error", err)
        return
    }
    
}


func (c *FileController) Insert(item *models.File) {
    
    
	conn := c.NewConnection()
    
	manager := models.NewFileManager(conn)
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

func (c *FileController) Insertbatch(item *[]models.File) {  
    if item == nil || len(*item) == 0 {
        return
    }

    rows := len(*item)
    
    
    
	conn := c.NewConnection()
    
	manager := models.NewFileManager(conn)

    for i := 0; i < rows; i++ {
	    err := manager.Insert(&((*item)[i]))
        if err != nil {
            c.Set("code", "error")    
            c.Set("error", err)
            return
        }
    }
}

func (c *FileController) Update(item *models.File) {
    
    
	conn := c.NewConnection()

	manager := models.NewFileManager(conn)
    err := manager.Update(item)
    if err != nil {
        c.Set("code", "error")    
        c.Set("error", err)
        return
    }
}

func (c *FileController) Delete(item *models.File) {
    
    
    conn := c.NewConnection()

	manager := models.NewFileManager(conn)

    
	err := manager.Delete(item.Id)
    if err != nil {
        c.Set("code", "error")    
        c.Set("error", err)
    }
}

func (c *FileController) Deletebatch(item *[]models.File) {
    
    
    conn := c.NewConnection()

	manager := models.NewFileManager(conn)

    for _, v := range *item {
        
    
	    err := manager.Delete(v.Id)
        if err != nil {
            c.Set("code", "error")    
            c.Set("error", err)
            return
        }
    }
}

func (c *FileController) Count() {
    
    
	conn := c.NewConnection()

	manager := models.NewFileManager(conn)

    var args []interface{}
    
    _title := c.Get("title")
    if _title != "" {
        args = append(args, models.Where{Column:"title", Value:_title, Compare:"="})
        
    }
    _filename := c.Get("filename")
    if _filename != "" {
        args = append(args, models.Where{Column:"filename", Value:_filename, Compare:"="})
    }
    _originalfilename := c.Get("originalfilename")
    if _originalfilename != "" {
        args = append(args, models.Where{Column:"originalfilename", Value:_originalfilename, Compare:"="})
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


func (c *FileController) Read(id int64) {
    
    
	conn := c.NewConnection()

	manager := models.NewFileManager(conn)
	item := manager.Get(id)

    
    
    c.Set("item", item)
}

func (c *FileController) Index(page int, pagesize int) {
    
    
	conn := c.NewConnection()

	manager := models.NewFileManager(conn)

    var args []interface{}
    
    _title := c.Get("title")
    if _title != "" {
        args = append(args, models.Where{Column:"title", Value:_title, Compare:"="})
        
    }
    _filename := c.Get("filename")
    if _filename != "" {
        args = append(args, models.Where{Column:"filename", Value:_filename, Compare:"="})
    }
    _originalfilename := c.Get("originalfilename")
    if _originalfilename != "" {
        args = append(args, models.Where{Column:"originalfilename", Value:_originalfilename, Compare:"="})
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
                    str += ", f_" + strings.Trim(v, " ")                
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





