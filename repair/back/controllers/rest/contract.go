package rest


import (
	"repair/controllers"
	"repair/models"

    "strings"
)

type ContractController struct {
	controllers.Controller
}



func (c *ContractController) GetByEstimate(estimate int64) *models.Contract {
    
    conn := c.NewConnection()

	_manager := models.NewContractManager(conn)
    
    item := _manager.GetByEstimate(estimate)
    
    c.Set("item", item)
    
    
    
    return item
    
}


func (c *ContractController) Insert(item *models.Contract) {
    
    
	conn := c.NewConnection()
    
	manager := models.NewContractManager(conn)
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

func (c *ContractController) Insertbatch(item *[]models.Contract) {  
    if item == nil || len(*item) == 0 {
        return
    }

    rows := len(*item)
    
    
    
	conn := c.NewConnection()
    
	manager := models.NewContractManager(conn)

    for i := 0; i < rows; i++ {
	    err := manager.Insert(&((*item)[i]))
        if err != nil {
            c.Set("code", "error")    
            c.Set("error", err)
            return
        }
    }
}

func (c *ContractController) Update(item *models.Contract) {
    
    
	conn := c.NewConnection()

	manager := models.NewContractManager(conn)
    err := manager.Update(item)
    if err != nil {
        c.Set("code", "error")    
        c.Set("error", err)
        return
    }
}

func (c *ContractController) Delete(item *models.Contract) {
    
    
    conn := c.NewConnection()

	manager := models.NewContractManager(conn)

    
	err := manager.Delete(item.Id)
    if err != nil {
        c.Set("code", "error")    
        c.Set("error", err)
    }
}

func (c *ContractController) Deletebatch(item *[]models.Contract) {
    
    
    conn := c.NewConnection()

	manager := models.NewContractManager(conn)

    for _, v := range *item {
        
    
	    err := manager.Delete(v.Id)
        if err != nil {
            c.Set("code", "error")    
            c.Set("error", err)
            return
        }
    }
}

func (c *ContractController) Count() {
    
    
	conn := c.NewConnection()

	manager := models.NewContractManager(conn)

    var args []interface{}
    
    _type := c.Geti("type")
    if _type != 0 {
        args = append(args, models.Where{Column:"type", Value:_type, Compare:"="})    
    }
    _startcontractdate := c.Get("startcontractdate")
    _endcontractdate := c.Get("endcontractdate")

    if _startcontractdate != "" && _endcontractdate != "" {        
        var v [2]string
        v[0] = _startcontractdate
        v[1] = _endcontractdate  
        args = append(args, models.Where{Column:"contractdate", Value:v, Compare:"between"})    
    } else if  _startcontractdate != "" {          
        args = append(args, models.Where{Column:"contractdate", Value:_startcontractdate, Compare:">="})
    } else if  _endcontractdate != "" {          
        args = append(args, models.Where{Column:"contractdate", Value:_endcontractdate, Compare:"<="})            
    }
    _startcontractstartdate := c.Get("startcontractstartdate")
    _endcontractstartdate := c.Get("endcontractstartdate")

    if _startcontractstartdate != "" && _endcontractstartdate != "" {        
        var v [2]string
        v[0] = _startcontractstartdate
        v[1] = _endcontractstartdate  
        args = append(args, models.Where{Column:"contractstartdate", Value:v, Compare:"between"})    
    } else if  _startcontractstartdate != "" {          
        args = append(args, models.Where{Column:"contractstartdate", Value:_startcontractstartdate, Compare:">="})
    } else if  _endcontractstartdate != "" {          
        args = append(args, models.Where{Column:"contractstartdate", Value:_endcontractstartdate, Compare:"<="})            
    }
    _startcontractenddate := c.Get("startcontractenddate")
    _endcontractenddate := c.Get("endcontractenddate")

    if _startcontractenddate != "" && _endcontractenddate != "" {        
        var v [2]string
        v[0] = _startcontractenddate
        v[1] = _endcontractenddate  
        args = append(args, models.Where{Column:"contractenddate", Value:v, Compare:"between"})    
    } else if  _startcontractenddate != "" {          
        args = append(args, models.Where{Column:"contractenddate", Value:_startcontractenddate, Compare:">="})
    } else if  _endcontractenddate != "" {          
        args = append(args, models.Where{Column:"contractenddate", Value:_endcontractenddate, Compare:"<="})            
    }
    _price := c.Geti("price")
    if _price != 0 {
        args = append(args, models.Where{Column:"price", Value:_price, Compare:"="})    
    }
    _vat := c.Geti("vat")
    if _vat != 0 {
        args = append(args, models.Where{Column:"vat", Value:_vat, Compare:"="})    
    }
    _startinvoice := c.Get("startinvoice")
    _endinvoice := c.Get("endinvoice")

    if _startinvoice != "" && _endinvoice != "" {        
        var v [2]string
        v[0] = _startinvoice
        v[1] = _endinvoice  
        args = append(args, models.Where{Column:"invoice", Value:v, Compare:"between"})    
    } else if  _startinvoice != "" {          
        args = append(args, models.Where{Column:"invoice", Value:_startinvoice, Compare:">="})
    } else if  _endinvoice != "" {          
        args = append(args, models.Where{Column:"invoice", Value:_endinvoice, Compare:"<="})            
    }
    _startdepositdate := c.Get("startdepositdate")
    _enddepositdate := c.Get("enddepositdate")

    if _startdepositdate != "" && _enddepositdate != "" {        
        var v [2]string
        v[0] = _startdepositdate
        v[1] = _enddepositdate  
        args = append(args, models.Where{Column:"depositdate", Value:v, Compare:"between"})    
    } else if  _startdepositdate != "" {          
        args = append(args, models.Where{Column:"depositdate", Value:_startdepositdate, Compare:">="})
    } else if  _enddepositdate != "" {          
        args = append(args, models.Where{Column:"depositdate", Value:_enddepositdate, Compare:"<="})            
    }
    _remark := c.Get("remark")
    if _remark != "" {
        args = append(args, models.Where{Column:"remark", Value:_remark, Compare:"="})
    }
    _user := c.Geti64("user")
    if _user != 0 {
        args = append(args, models.Where{Column:"user", Value:_user, Compare:"="})    
    }
    _estimate := c.Geti64("estimate")
    if _estimate != 0 {
        args = append(args, models.Where{Column:"estimate", Value:_estimate, Compare:"="})    
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


func (c *ContractController) Read(id int64) {
    
    
	conn := c.NewConnection()

	manager := models.NewContractManager(conn)
	item := manager.Get(id)

    
    
    c.Set("item", item)
}

func (c *ContractController) Index(page int, pagesize int) {
    
    
	conn := c.NewConnection()

	manager := models.NewContractManager(conn)

    var args []interface{}
    
    _type := c.Geti("type")
    if _type != 0 {
        args = append(args, models.Where{Column:"type", Value:_type, Compare:"="})    
    }
    _startcontractdate := c.Get("startcontractdate")
    _endcontractdate := c.Get("endcontractdate")
    if _startcontractdate != "" && _endcontractdate != "" {        
        var v [2]string
        v[0] = _startcontractdate
        v[1] = _endcontractdate  
        args = append(args, models.Where{Column:"contractdate", Value:v, Compare:"between"})    
    } else if  _startcontractdate != "" {          
        args = append(args, models.Where{Column:"contractdate", Value:_startcontractdate, Compare:">="})
    } else if  _endcontractdate != "" {          
        args = append(args, models.Where{Column:"contractdate", Value:_endcontractdate, Compare:"<="})            
    }
    _startcontractstartdate := c.Get("startcontractstartdate")
    _endcontractstartdate := c.Get("endcontractstartdate")
    if _startcontractstartdate != "" && _endcontractstartdate != "" {        
        var v [2]string
        v[0] = _startcontractstartdate
        v[1] = _endcontractstartdate  
        args = append(args, models.Where{Column:"contractstartdate", Value:v, Compare:"between"})    
    } else if  _startcontractstartdate != "" {          
        args = append(args, models.Where{Column:"contractstartdate", Value:_startcontractstartdate, Compare:">="})
    } else if  _endcontractstartdate != "" {          
        args = append(args, models.Where{Column:"contractstartdate", Value:_endcontractstartdate, Compare:"<="})            
    }
    _startcontractenddate := c.Get("startcontractenddate")
    _endcontractenddate := c.Get("endcontractenddate")
    if _startcontractenddate != "" && _endcontractenddate != "" {        
        var v [2]string
        v[0] = _startcontractenddate
        v[1] = _endcontractenddate  
        args = append(args, models.Where{Column:"contractenddate", Value:v, Compare:"between"})    
    } else if  _startcontractenddate != "" {          
        args = append(args, models.Where{Column:"contractenddate", Value:_startcontractenddate, Compare:">="})
    } else if  _endcontractenddate != "" {          
        args = append(args, models.Where{Column:"contractenddate", Value:_endcontractenddate, Compare:"<="})            
    }
    _price := c.Geti("price")
    if _price != 0 {
        args = append(args, models.Where{Column:"price", Value:_price, Compare:"="})    
    }
    _vat := c.Geti("vat")
    if _vat != 0 {
        args = append(args, models.Where{Column:"vat", Value:_vat, Compare:"="})    
    }
    _startinvoice := c.Get("startinvoice")
    _endinvoice := c.Get("endinvoice")
    if _startinvoice != "" && _endinvoice != "" {        
        var v [2]string
        v[0] = _startinvoice
        v[1] = _endinvoice  
        args = append(args, models.Where{Column:"invoice", Value:v, Compare:"between"})    
    } else if  _startinvoice != "" {          
        args = append(args, models.Where{Column:"invoice", Value:_startinvoice, Compare:">="})
    } else if  _endinvoice != "" {          
        args = append(args, models.Where{Column:"invoice", Value:_endinvoice, Compare:"<="})            
    }
    _startdepositdate := c.Get("startdepositdate")
    _enddepositdate := c.Get("enddepositdate")
    if _startdepositdate != "" && _enddepositdate != "" {        
        var v [2]string
        v[0] = _startdepositdate
        v[1] = _enddepositdate  
        args = append(args, models.Where{Column:"depositdate", Value:v, Compare:"between"})    
    } else if  _startdepositdate != "" {          
        args = append(args, models.Where{Column:"depositdate", Value:_startdepositdate, Compare:">="})
    } else if  _enddepositdate != "" {          
        args = append(args, models.Where{Column:"depositdate", Value:_enddepositdate, Compare:"<="})            
    }
    _remark := c.Get("remark")
    if _remark != "" {
        args = append(args, models.Where{Column:"remark", Value:_remark, Compare:"="})
    }
    _user := c.Geti64("user")
    if _user != 0 {
        args = append(args, models.Where{Column:"user", Value:_user, Compare:"="})    
    }
    _estimate := c.Geti64("estimate")
    if _estimate != 0 {
        args = append(args, models.Where{Column:"estimate", Value:_estimate, Compare:"="})    
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
                    str += ", co_" + strings.Trim(v, " ")                
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





func (c *ContractController) Sum() {
    
    
	conn := c.NewConnection()

	manager := models.NewContractManager(conn)

    var args []interface{}
    
    _type := c.Geti("type")
    if _type != 0 {
        args = append(args, models.Where{Column:"type", Value:_type, Compare:"="})    
    }
    _startcontractdate := c.Get("startcontractdate")
    _endcontractdate := c.Get("endcontractdate")
    if _startcontractdate != "" && _endcontractdate != "" {        
        var v [2]string
        v[0] = _startcontractdate
        v[1] = _endcontractdate  
        args = append(args, models.Where{Column:"contractdate", Value:v, Compare:"between"})    
    } else if  _startcontractdate != "" {          
        args = append(args, models.Where{Column:"contractdate", Value:_startcontractdate, Compare:">="})
    } else if  _endcontractdate != "" {          
        args = append(args, models.Where{Column:"contractdate", Value:_endcontractdate, Compare:"<="})            
    }
    _startcontractstartdate := c.Get("startcontractstartdate")
    _endcontractstartdate := c.Get("endcontractstartdate")
    if _startcontractstartdate != "" && _endcontractstartdate != "" {        
        var v [2]string
        v[0] = _startcontractstartdate
        v[1] = _endcontractstartdate  
        args = append(args, models.Where{Column:"contractstartdate", Value:v, Compare:"between"})    
    } else if  _startcontractstartdate != "" {          
        args = append(args, models.Where{Column:"contractstartdate", Value:_startcontractstartdate, Compare:">="})
    } else if  _endcontractstartdate != "" {          
        args = append(args, models.Where{Column:"contractstartdate", Value:_endcontractstartdate, Compare:"<="})            
    }
    _startcontractenddate := c.Get("startcontractenddate")
    _endcontractenddate := c.Get("endcontractenddate")
    if _startcontractenddate != "" && _endcontractenddate != "" {        
        var v [2]string
        v[0] = _startcontractenddate
        v[1] = _endcontractenddate  
        args = append(args, models.Where{Column:"contractenddate", Value:v, Compare:"between"})    
    } else if  _startcontractenddate != "" {          
        args = append(args, models.Where{Column:"contractenddate", Value:_startcontractenddate, Compare:">="})
    } else if  _endcontractenddate != "" {          
        args = append(args, models.Where{Column:"contractenddate", Value:_endcontractenddate, Compare:"<="})            
    }
    _price := c.Geti("price")
    if _price != 0 {
        args = append(args, models.Where{Column:"price", Value:_price, Compare:"="})    
    }
    _vat := c.Geti("vat")
    if _vat != 0 {
        args = append(args, models.Where{Column:"vat", Value:_vat, Compare:"="})    
    }
    _startinvoice := c.Get("startinvoice")
    _endinvoice := c.Get("endinvoice")
    if _startinvoice != "" && _endinvoice != "" {        
        var v [2]string
        v[0] = _startinvoice
        v[1] = _endinvoice  
        args = append(args, models.Where{Column:"invoice", Value:v, Compare:"between"})    
    } else if  _startinvoice != "" {          
        args = append(args, models.Where{Column:"invoice", Value:_startinvoice, Compare:">="})
    } else if  _endinvoice != "" {          
        args = append(args, models.Where{Column:"invoice", Value:_endinvoice, Compare:"<="})            
    }
    _startdepositdate := c.Get("startdepositdate")
    _enddepositdate := c.Get("enddepositdate")
    if _startdepositdate != "" && _enddepositdate != "" {        
        var v [2]string
        v[0] = _startdepositdate
        v[1] = _enddepositdate  
        args = append(args, models.Where{Column:"depositdate", Value:v, Compare:"between"})    
    } else if  _startdepositdate != "" {          
        args = append(args, models.Where{Column:"depositdate", Value:_startdepositdate, Compare:">="})
    } else if  _enddepositdate != "" {          
        args = append(args, models.Where{Column:"depositdate", Value:_enddepositdate, Compare:"<="})            
    }
    _remark := c.Get("remark")
    if _remark != "" {
        args = append(args, models.Where{Column:"remark", Value:_remark, Compare:"like"})
    }
    _user := c.Geti64("user")
    if _user != 0 {
        args = append(args, models.Where{Column:"user", Value:_user, Compare:"="})    
    }
    _estimate := c.Geti64("estimate")
    if _estimate != 0 {
        args = append(args, models.Where{Column:"estimate", Value:_estimate, Compare:"="})    
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
    

    
    
    item := manager.Sum(args)
	c.Set("item", item)
}

