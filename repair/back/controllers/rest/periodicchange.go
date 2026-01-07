package rest


import (
	"repair/controllers"
	"repair/models"

    "strings"
)

type PeriodicchangeController struct {
	controllers.Controller
}



func (c *PeriodicchangeController) Insert(item *models.Periodicchange) {
    
    
	conn := c.NewConnection()
    
	manager := models.NewPeriodicchangeManager(conn)
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

func (c *PeriodicchangeController) Insertbatch(item *[]models.Periodicchange) {  
    if item == nil || len(*item) == 0 {
        return
    }

    rows := len(*item)
    
    
    
	conn := c.NewConnection()
    
	manager := models.NewPeriodicchangeManager(conn)

    for i := 0; i < rows; i++ {
	    err := manager.Insert(&((*item)[i]))
        if err != nil {
            c.Set("code", "error")    
            c.Set("error", err)
            return
        }
    }
}

func (c *PeriodicchangeController) Update(item *models.Periodicchange) {
    
    
	conn := c.NewConnection()

	manager := models.NewPeriodicchangeManager(conn)
    err := manager.Update(item)
    if err != nil {
        c.Set("code", "error")    
        c.Set("error", err)
        return
    }
}

func (c *PeriodicchangeController) Delete(item *models.Periodicchange) {
    
    
    conn := c.NewConnection()

	manager := models.NewPeriodicchangeManager(conn)

    
	err := manager.Delete(item.Id)
    if err != nil {
        c.Set("code", "error")    
        c.Set("error", err)
    }
}

func (c *PeriodicchangeController) Deletebatch(item *[]models.Periodicchange) {
    
    
    conn := c.NewConnection()

	manager := models.NewPeriodicchangeManager(conn)

    for _, v := range *item {
        
    
	    err := manager.Delete(v.Id)
        if err != nil {
            c.Set("code", "error")    
            c.Set("error", err)
            return
        }
    }
}

func (c *PeriodicchangeController) Count() {
    
    
	conn := c.NewConnection()

	manager := models.NewPeriodicchangeManager(conn)

    var args []interface{}
    
    _content1 := c.Get("content1")
    if _content1 != "" {
        args = append(args, models.Where{Column:"content1", Value:_content1, Compare:"="})
    }
    _content2 := c.Get("content2")
    if _content2 != "" {
        args = append(args, models.Where{Column:"content2", Value:_content2, Compare:"="})
    }
    _content3 := c.Get("content3")
    if _content3 != "" {
        args = append(args, models.Where{Column:"content3", Value:_content3, Compare:"="})
    }
    _content4 := c.Get("content4")
    if _content4 != "" {
        args = append(args, models.Where{Column:"content4", Value:_content4, Compare:"="})
    }
    _content5 := c.Get("content5")
    if _content5 != "" {
        args = append(args, models.Where{Column:"content5", Value:_content5, Compare:"="})
    }
    _content6 := c.Get("content6")
    if _content6 != "" {
        args = append(args, models.Where{Column:"content6", Value:_content6, Compare:"="})
    }
    _content7 := c.Get("content7")
    if _content7 != "" {
        args = append(args, models.Where{Column:"content7", Value:_content7, Compare:"="})
    }
    _type := c.Geti("type")
    if _type != 0 {
        args = append(args, models.Where{Column:"type", Value:_type, Compare:"="})    
    }
    _order := c.Geti("order")
    if _order != 0 {
        args = append(args, models.Where{Column:"order", Value:_order, Compare:"="})    
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


func (c *PeriodicchangeController) Read(id int64) {
    
    
	conn := c.NewConnection()

	manager := models.NewPeriodicchangeManager(conn)
	item := manager.Get(id)

    
    
    c.Set("item", item)
}

func (c *PeriodicchangeController) Index(page int, pagesize int) {
    
    
	conn := c.NewConnection()

	manager := models.NewPeriodicchangeManager(conn)

    var args []interface{}
    
    _content1 := c.Get("content1")
    if _content1 != "" {
        args = append(args, models.Where{Column:"content1", Value:_content1, Compare:"="})
    }
    _content2 := c.Get("content2")
    if _content2 != "" {
        args = append(args, models.Where{Column:"content2", Value:_content2, Compare:"="})
    }
    _content3 := c.Get("content3")
    if _content3 != "" {
        args = append(args, models.Where{Column:"content3", Value:_content3, Compare:"="})
    }
    _content4 := c.Get("content4")
    if _content4 != "" {
        args = append(args, models.Where{Column:"content4", Value:_content4, Compare:"="})
    }
    _content5 := c.Get("content5")
    if _content5 != "" {
        args = append(args, models.Where{Column:"content5", Value:_content5, Compare:"="})
    }
    _content6 := c.Get("content6")
    if _content6 != "" {
        args = append(args, models.Where{Column:"content6", Value:_content6, Compare:"="})
    }
    _content7 := c.Get("content7")
    if _content7 != "" {
        args = append(args, models.Where{Column:"content7", Value:_content7, Compare:"="})
    }
    _type := c.Geti("type")
    if _type != 0 {
        args = append(args, models.Where{Column:"type", Value:_type, Compare:"="})    
    }
    _order := c.Geti("order")
    if _order != 0 {
        args = append(args, models.Where{Column:"order", Value:_order, Compare:"="})    
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
                    str += ", pc_" + strings.Trim(v, " ")                
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





