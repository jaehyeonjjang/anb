package rest


import (
	"repair/controllers"
	"repair/models"

    "strings"
)

type ReviewbasicController struct {
	controllers.Controller
}



func (c *ReviewbasicController) Insert(item *models.Reviewbasic) {
    
    
	conn := c.NewConnection()
    
	manager := models.NewReviewbasicManager(conn)
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

func (c *ReviewbasicController) Insertbatch(item *[]models.Reviewbasic) {  
    if item == nil || len(*item) == 0 {
        return
    }

    rows := len(*item)
    
    
    
	conn := c.NewConnection()
    
	manager := models.NewReviewbasicManager(conn)

    for i := 0; i < rows; i++ {
	    err := manager.Insert(&((*item)[i]))
        if err != nil {
            c.Set("code", "error")    
            c.Set("error", err)
            return
        }
    }
}

func (c *ReviewbasicController) Update(item *models.Reviewbasic) {
    
    
	conn := c.NewConnection()

	manager := models.NewReviewbasicManager(conn)
    err := manager.Update(item)
    if err != nil {
        c.Set("code", "error")    
        c.Set("error", err)
        return
    }
}

func (c *ReviewbasicController) Delete(item *models.Reviewbasic) {
    
    
    conn := c.NewConnection()

	manager := models.NewReviewbasicManager(conn)

    
	err := manager.Delete(item.Id)
    if err != nil {
        c.Set("code", "error")    
        c.Set("error", err)
    }
}

func (c *ReviewbasicController) Deletebatch(item *[]models.Reviewbasic) {
    
    
    conn := c.NewConnection()

	manager := models.NewReviewbasicManager(conn)

    for _, v := range *item {
        
    
	    err := manager.Delete(v.Id)
        if err != nil {
            c.Set("code", "error")    
            c.Set("error", err)
            return
        }
    }
}

func (c *ReviewbasicController) Count() {
    
    
	conn := c.NewConnection()

	manager := models.NewReviewbasicManager(conn)

    var args []interface{}
    
    _content := c.Get("content")
    if _content != "" {
        args = append(args, models.Where{Column:"content", Value:_content, Compare:"like"})
        
    }
    _adjust := c.Get("adjust")
    if _adjust != "" {
        args = append(args, models.Where{Column:"adjust", Value:_adjust, Compare:"="})
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
    

    
    
    total := manager.Count(args)
	c.Set("total", total)
}


func (c *ReviewbasicController) Read(id int64) {
    
    
	conn := c.NewConnection()

	manager := models.NewReviewbasicManager(conn)
	item := manager.Get(id)

    
    
    c.Set("item", item)
}

func (c *ReviewbasicController) Index(page int, pagesize int) {
    
    
	conn := c.NewConnection()

	manager := models.NewReviewbasicManager(conn)

    var args []interface{}
    
    _content := c.Get("content")
    if _content != "" {
        args = append(args, models.Where{Column:"content", Value:_content, Compare:"like"})
        
    }
    _adjust := c.Get("adjust")
    if _adjust != "" {
        args = append(args, models.Where{Column:"adjust", Value:_adjust, Compare:"="})
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
                    str += ", rv_" + strings.Trim(v, " ")                
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





