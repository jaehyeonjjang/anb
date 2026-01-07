package rest


import (
	"repair/controllers"
	"repair/models"

    "strings"
)

type AptdongController struct {
	controllers.Controller
}



func (c *AptdongController) Insert(item *models.Aptdong) {
    
    
	conn := c.NewConnection()
    
	manager := models.NewAptdongManager(conn)
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

func (c *AptdongController) Insertbatch(item *[]models.Aptdong) {  
    if item == nil || len(*item) == 0 {
        return
    }

    rows := len(*item)
    
    
    
	conn := c.NewConnection()
    
	manager := models.NewAptdongManager(conn)

    for i := 0; i < rows; i++ {
	    err := manager.Insert(&((*item)[i]))
        if err != nil {
            c.Set("code", "error")    
            c.Set("error", err)
            return
        }
    }
}

func (c *AptdongController) Update(item *models.Aptdong) {
    
    
	conn := c.NewConnection()

	manager := models.NewAptdongManager(conn)
    err := manager.Update(item)
    if err != nil {
        c.Set("code", "error")    
        c.Set("error", err)
        return
    }
}

func (c *AptdongController) Delete(item *models.Aptdong) {
    
    
    conn := c.NewConnection()

	manager := models.NewAptdongManager(conn)

    
	err := manager.Delete(item.Id)
    if err != nil {
        c.Set("code", "error")    
        c.Set("error", err)
    }
}

func (c *AptdongController) Deletebatch(item *[]models.Aptdong) {
    
    
    conn := c.NewConnection()

	manager := models.NewAptdongManager(conn)

    for _, v := range *item {
        
    
	    err := manager.Delete(v.Id)
        if err != nil {
            c.Set("code", "error")    
            c.Set("error", err)
            return
        }
    }
}

func (c *AptdongController) Count() {
    
    
	conn := c.NewConnection()

	manager := models.NewAptdongManager(conn)

    var args []interface{}
    
    _dong := c.Get("dong")
    if _dong != "" {
        args = append(args, models.Where{Column:"dong", Value:_dong, Compare:"="})
    }
    _type := c.Get("type")
    if _type != "" {
        args = append(args, models.Where{Column:"type", Value:_type, Compare:"="})
    }
    _groundcount := c.Geti("groundcount")
    if _groundcount != 0 {
        args = append(args, models.Where{Column:"groundcount", Value:_groundcount, Compare:"="})    
    }
    _undergroundcount := c.Geti("undergroundcount")
    if _undergroundcount != 0 {
        args = append(args, models.Where{Column:"undergroundcount", Value:_undergroundcount, Compare:"="})    
    }
    _parkingcount := c.Geti("parkingcount")
    if _parkingcount != 0 {
        args = append(args, models.Where{Column:"parkingcount", Value:_parkingcount, Compare:"="})    
    }
    _topcount := c.Geti("topcount")
    if _topcount != 0 {
        args = append(args, models.Where{Column:"topcount", Value:_topcount, Compare:"="})    
    }
    _roofcount := c.Geti("roofcount")
    if _roofcount != 0 {
        args = append(args, models.Where{Column:"roofcount", Value:_roofcount, Compare:"="})    
    }
    _familycount := c.Geti("familycount")
    if _familycount != 0 {
        args = append(args, models.Where{Column:"familycount", Value:_familycount, Compare:"="})    
    }
    _area := c.Get("area")
    if _area != "" {
        args = append(args, models.Where{Column:"area", Value:_area, Compare:"="})
    }
    _remark := c.Get("remark")
    if _remark != "" {
        args = append(args, models.Where{Column:"remark", Value:_remark, Compare:"="})
    }
    _private := c.Geti("private")
    if _private != 0 {
        args = append(args, models.Where{Column:"private", Value:_private, Compare:"="})    
    }
    _order := c.Geti("order")
    if _order != 0 {
        args = append(args, models.Where{Column:"order", Value:_order, Compare:"="})    
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


func (c *AptdongController) Read(id int64) {
    
    
	conn := c.NewConnection()

	manager := models.NewAptdongManager(conn)
	item := manager.Get(id)

    
    
    c.Set("item", item)
}

func (c *AptdongController) Index(page int, pagesize int) {
    
    
	conn := c.NewConnection()

	manager := models.NewAptdongManager(conn)

    var args []interface{}
    
    _dong := c.Get("dong")
    if _dong != "" {
        args = append(args, models.Where{Column:"dong", Value:_dong, Compare:"="})
    }
    _type := c.Get("type")
    if _type != "" {
        args = append(args, models.Where{Column:"type", Value:_type, Compare:"="})
    }
    _groundcount := c.Geti("groundcount")
    if _groundcount != 0 {
        args = append(args, models.Where{Column:"groundcount", Value:_groundcount, Compare:"="})    
    }
    _undergroundcount := c.Geti("undergroundcount")
    if _undergroundcount != 0 {
        args = append(args, models.Where{Column:"undergroundcount", Value:_undergroundcount, Compare:"="})    
    }
    _parkingcount := c.Geti("parkingcount")
    if _parkingcount != 0 {
        args = append(args, models.Where{Column:"parkingcount", Value:_parkingcount, Compare:"="})    
    }
    _topcount := c.Geti("topcount")
    if _topcount != 0 {
        args = append(args, models.Where{Column:"topcount", Value:_topcount, Compare:"="})    
    }
    _roofcount := c.Geti("roofcount")
    if _roofcount != 0 {
        args = append(args, models.Where{Column:"roofcount", Value:_roofcount, Compare:"="})    
    }
    _familycount := c.Geti("familycount")
    if _familycount != 0 {
        args = append(args, models.Where{Column:"familycount", Value:_familycount, Compare:"="})    
    }
    _area := c.Get("area")
    if _area != "" {
        args = append(args, models.Where{Column:"area", Value:_area, Compare:"="})
    }
    _remark := c.Get("remark")
    if _remark != "" {
        args = append(args, models.Where{Column:"remark", Value:_remark, Compare:"="})
    }
    _private := c.Geti("private")
    if _private != 0 {
        args = append(args, models.Where{Column:"private", Value:_private, Compare:"="})    
    }
    _order := c.Geti("order")
    if _order != 0 {
        args = append(args, models.Where{Column:"order", Value:_order, Compare:"="})    
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
                    str += ", au_" + strings.Trim(v, " ")                
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





