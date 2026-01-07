package rest


import (
	"repair/controllers"
	"repair/models"

    "strings"
)

type PeriodicresultController struct {
	controllers.Controller
}



func (c *PeriodicresultController) Insert(item *models.Periodicresult) {
    
    
	conn := c.NewConnection()
    
	manager := models.NewPeriodicresultManager(conn)
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

func (c *PeriodicresultController) Insertbatch(item *[]models.Periodicresult) {  
    if item == nil || len(*item) == 0 {
        return
    }

    rows := len(*item)
    
    
    
	conn := c.NewConnection()
    
	manager := models.NewPeriodicresultManager(conn)

    for i := 0; i < rows; i++ {
	    err := manager.Insert(&((*item)[i]))
        if err != nil {
            c.Set("code", "error")    
            c.Set("error", err)
            return
        }
    }
}

func (c *PeriodicresultController) Update(item *models.Periodicresult) {
    
    
	conn := c.NewConnection()

	manager := models.NewPeriodicresultManager(conn)
    err := manager.Update(item)
    if err != nil {
        c.Set("code", "error")    
        c.Set("error", err)
        return
    }
}

func (c *PeriodicresultController) Delete(item *models.Periodicresult) {
    
    
    conn := c.NewConnection()

	manager := models.NewPeriodicresultManager(conn)

    
	err := manager.Delete(item.Id)
    if err != nil {
        c.Set("code", "error")    
        c.Set("error", err)
    }
}

func (c *PeriodicresultController) Deletebatch(item *[]models.Periodicresult) {
    
    
    conn := c.NewConnection()

	manager := models.NewPeriodicresultManager(conn)

    for _, v := range *item {
        
    
	    err := manager.Delete(v.Id)
        if err != nil {
            c.Set("code", "error")    
            c.Set("error", err)
            return
        }
    }
}

func (c *PeriodicresultController) Count() {
    
    
	conn := c.NewConnection()

	manager := models.NewPeriodicresultManager(conn)

    var args []interface{}
    
    _defect := c.Get("defect")
    if _defect != "" {
        args = append(args, models.Where{Column:"defect", Value:_defect, Compare:"="})
    }
    _reinforcement := c.Get("reinforcement")
    if _reinforcement != "" {
        args = append(args, models.Where{Column:"reinforcement", Value:_reinforcement, Compare:"="})
    }
    _remark := c.Get("remark")
    if _remark != "" {
        args = append(args, models.Where{Column:"remark", Value:_remark, Compare:"="})
    }
    _periodic := c.Geti64("periodic")
    if _periodic != 0 {
        args = append(args, models.Where{Column:"periodic", Value:_periodic, Compare:"="})    
    }
    _aptdong := c.Geti64("aptdong")
    if _aptdong != 0 {
        args = append(args, models.Where{Column:"aptdong", Value:_aptdong, Compare:"="})    
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


func (c *PeriodicresultController) Read(id int64) {
    
    
	conn := c.NewConnection()

	manager := models.NewPeriodicresultManager(conn)
	item := manager.Get(id)

    
    
    c.Set("item", item)
}

func (c *PeriodicresultController) Index(page int, pagesize int) {
    
    
	conn := c.NewConnection()

	manager := models.NewPeriodicresultManager(conn)

    var args []interface{}
    
    _defect := c.Get("defect")
    if _defect != "" {
        args = append(args, models.Where{Column:"defect", Value:_defect, Compare:"="})
    }
    _reinforcement := c.Get("reinforcement")
    if _reinforcement != "" {
        args = append(args, models.Where{Column:"reinforcement", Value:_reinforcement, Compare:"="})
    }
    _remark := c.Get("remark")
    if _remark != "" {
        args = append(args, models.Where{Column:"remark", Value:_remark, Compare:"="})
    }
    _periodic := c.Geti64("periodic")
    if _periodic != 0 {
        args = append(args, models.Where{Column:"periodic", Value:_periodic, Compare:"="})    
    }
    _aptdong := c.Geti64("aptdong")
    if _aptdong != 0 {
        args = append(args, models.Where{Column:"aptdong", Value:_aptdong, Compare:"="})    
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
                    str += ", pr_" + strings.Trim(v, " ")                
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





