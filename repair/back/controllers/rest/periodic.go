package rest


import (
	"repair/controllers"
	"repair/models"

    "strings"
)

type PeriodicController struct {
	controllers.Controller
}



func (c *PeriodicController) Insert(item *models.Periodic) {
    
    if c.Session == nil {
        item = nil
        return
    }

    if c.Session.Level < 3 {
    
        if c.Session.Apt == 0 {
            item = nil
            return
        } else {
            item.Apt = c.Session.Apt
        }    
    
    } else {
    
        if item.Apt == 0 {
            item.Apt = c.Session.Apt
        }    
    
    }
    
    
	conn := c.NewConnection()
    
	manager := models.NewPeriodicManager(conn)
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

func (c *PeriodicController) Insertbatch(item *[]models.Periodic) {  
    if item == nil || len(*item) == 0 {
        return
    }

    rows := len(*item)
    
    
    if c.Session == nil {
        return
    }

    if c.Session.Level < 3 {
    
        if c.Session.Apt == 0 {
            return
        } else {
            for i := 0; i < rows; i++ {
                (*item)[i].Apt = c.Session.Apt
            }
        }    
    
    }
    
    
	conn := c.NewConnection()
    
	manager := models.NewPeriodicManager(conn)

    for i := 0; i < rows; i++ {
	    err := manager.Insert(&((*item)[i]))
        if err != nil {
            c.Set("code", "error")    
            c.Set("error", err)
            return
        }
    }
}

func (c *PeriodicController) Update(item *models.Periodic) {
    
    if c.Session == nil {
        item = nil
        return
    }
    
    if c.Session.Level < 3 {
    
        if c.Session.Apt == 0 {
            item = nil
            return
        } else {
            item.Apt = c.Session.Apt
        }    
    
    }
    
    
	conn := c.NewConnection()

	manager := models.NewPeriodicManager(conn)
    err := manager.Update(item)
    if err != nil {
        c.Set("code", "error")    
        c.Set("error", err)
        return
    }
}

func (c *PeriodicController) Delete(item *models.Periodic) {
    
    if c.Session == nil {
        item = nil
        return
    }
    
    
    conn := c.NewConnection()

	manager := models.NewPeriodicManager(conn)

    
    n := manager.Get(item.Id)

    if c.Session.Level < 3 {
    
        if c.Session.Apt == 0 {
            item = nil
            return
        } else {
            if n.Apt != c.Session.Apt {
                item = nil
                return
            }
        }    
    
    }
    
	err := manager.Delete(item.Id)
    if err != nil {
        c.Set("code", "error")    
        c.Set("error", err)
    }
}

func (c *PeriodicController) Deletebatch(item *[]models.Periodic) {
    
    if c.Session == nil {
        return
    }
    
    
    conn := c.NewConnection()

	manager := models.NewPeriodicManager(conn)

    for _, v := range *item {
        
        n := manager.Get(v.Id)

        if c.Session.Level < 3 {
        
            if c.Session.Apt == 0 {
                return
            } else {
                if n.Apt != c.Session.Apt {
                    return
                }
            }    
        
        }
        
    
	    err := manager.Delete(v.Id)
        if err != nil {
            c.Set("code", "error")    
            c.Set("error", err)
            return
        }
    }
}

func (c *PeriodicController) Count() {
    
    if c.Session == nil {
        c.Result["code"] = "auth error"
        return
    }

    if c.Session.Level < 3 {
    
        if c.Session.Apt == 0 {
            c.Result["code"] = "auth error"
            return
        }
    
    }
    
    
	conn := c.NewConnection()

	manager := models.NewPeriodicManager(conn)

    var args []interface{}
    
    _name := c.Get("name")
    if _name != "" {
        args = append(args, models.Where{Column:"name", Value:_name, Compare:"like"})
        
    }
    _aptname := c.Get("aptname")
    if _aptname != "" {
        args = append(args, models.Where{Column:"aptname", Value:_aptname, Compare:"="})
    }
    _taskrange := c.Get("taskrange")
    if _taskrange != "" {
        args = append(args, models.Where{Column:"taskrange", Value:_taskrange, Compare:"="})
    }
    _startreportdate := c.Get("startreportdate")
    _endreportdate := c.Get("endreportdate")

    if _startreportdate != "" && _endreportdate != "" {        
        var v [2]string
        v[0] = _startreportdate
        v[1] = _endreportdate  
        args = append(args, models.Where{Column:"reportdate", Value:v, Compare:"between"})    
    } else if  _startreportdate != "" {          
        args = append(args, models.Where{Column:"reportdate", Value:_startreportdate, Compare:">="})
    } else if  _endreportdate != "" {          
        args = append(args, models.Where{Column:"reportdate", Value:_endreportdate, Compare:"<="})            
    }
    _startstartdate := c.Get("startstartdate")
    _endstartdate := c.Get("endstartdate")

    if _startstartdate != "" && _endstartdate != "" {        
        var v [2]string
        v[0] = _startstartdate
        v[1] = _endstartdate  
        args = append(args, models.Where{Column:"startdate", Value:v, Compare:"between"})    
    } else if  _startstartdate != "" {          
        args = append(args, models.Where{Column:"startdate", Value:_startstartdate, Compare:">="})
    } else if  _endstartdate != "" {          
        args = append(args, models.Where{Column:"startdate", Value:_endstartdate, Compare:"<="})            
    }
    _startenddate := c.Get("startenddate")
    _endenddate := c.Get("endenddate")

    if _startenddate != "" && _endenddate != "" {        
        var v [2]string
        v[0] = _startenddate
        v[1] = _endenddate  
        args = append(args, models.Where{Column:"enddate", Value:v, Compare:"between"})    
    } else if  _startenddate != "" {          
        args = append(args, models.Where{Column:"enddate", Value:_startenddate, Compare:">="})
    } else if  _endenddate != "" {          
        args = append(args, models.Where{Column:"enddate", Value:_endenddate, Compare:"<="})            
    }
    _supply := c.Get("supply")
    if _supply != "" {
        args = append(args, models.Where{Column:"supply", Value:_supply, Compare:"="})
    }
    _contract := c.Get("contract")
    if _contract != "" {
        args = append(args, models.Where{Column:"contract", Value:_contract, Compare:"="})
    }
    _price := c.Get("price")
    if _price != "" {
        args = append(args, models.Where{Column:"price", Value:_price, Compare:"="})
    }
    _safetygrade := c.Get("safetygrade")
    if _safetygrade != "" {
        args = append(args, models.Where{Column:"safetygrade", Value:_safetygrade, Compare:"="})
    }
    _status := c.Geti("status")
    if _status != 0 {
        args = append(args, models.Where{Column:"status", Value:_status, Compare:"="})    
    }
    _startprestartdate := c.Get("startprestartdate")
    _endprestartdate := c.Get("endprestartdate")

    if _startprestartdate != "" && _endprestartdate != "" {        
        var v [2]string
        v[0] = _startprestartdate
        v[1] = _endprestartdate  
        args = append(args, models.Where{Column:"prestartdate", Value:v, Compare:"between"})    
    } else if  _startprestartdate != "" {          
        args = append(args, models.Where{Column:"prestartdate", Value:_startprestartdate, Compare:">="})
    } else if  _endprestartdate != "" {          
        args = append(args, models.Where{Column:"prestartdate", Value:_endprestartdate, Compare:"<="})            
    }
    _startpreenddate := c.Get("startpreenddate")
    _endpreenddate := c.Get("endpreenddate")

    if _startpreenddate != "" && _endpreenddate != "" {        
        var v [2]string
        v[0] = _startpreenddate
        v[1] = _endpreenddate  
        args = append(args, models.Where{Column:"preenddate", Value:v, Compare:"between"})    
    } else if  _startpreenddate != "" {          
        args = append(args, models.Where{Column:"preenddate", Value:_startpreenddate, Compare:">="})
    } else if  _endpreenddate != "" {          
        args = append(args, models.Where{Column:"preenddate", Value:_endpreenddate, Compare:"<="})            
    }
    _startresearchstartdate := c.Get("startresearchstartdate")
    _endresearchstartdate := c.Get("endresearchstartdate")

    if _startresearchstartdate != "" && _endresearchstartdate != "" {        
        var v [2]string
        v[0] = _startresearchstartdate
        v[1] = _endresearchstartdate  
        args = append(args, models.Where{Column:"researchstartdate", Value:v, Compare:"between"})    
    } else if  _startresearchstartdate != "" {          
        args = append(args, models.Where{Column:"researchstartdate", Value:_startresearchstartdate, Compare:">="})
    } else if  _endresearchstartdate != "" {          
        args = append(args, models.Where{Column:"researchstartdate", Value:_endresearchstartdate, Compare:"<="})            
    }
    _startresearchenddate := c.Get("startresearchenddate")
    _endresearchenddate := c.Get("endresearchenddate")

    if _startresearchenddate != "" && _endresearchenddate != "" {        
        var v [2]string
        v[0] = _startresearchenddate
        v[1] = _endresearchenddate  
        args = append(args, models.Where{Column:"researchenddate", Value:v, Compare:"between"})    
    } else if  _startresearchenddate != "" {          
        args = append(args, models.Where{Column:"researchenddate", Value:_startresearchenddate, Compare:">="})
    } else if  _endresearchenddate != "" {          
        args = append(args, models.Where{Column:"researchenddate", Value:_endresearchenddate, Compare:"<="})            
    }
    _startanalyzestartdate := c.Get("startanalyzestartdate")
    _endanalyzestartdate := c.Get("endanalyzestartdate")

    if _startanalyzestartdate != "" && _endanalyzestartdate != "" {        
        var v [2]string
        v[0] = _startanalyzestartdate
        v[1] = _endanalyzestartdate  
        args = append(args, models.Where{Column:"analyzestartdate", Value:v, Compare:"between"})    
    } else if  _startanalyzestartdate != "" {          
        args = append(args, models.Where{Column:"analyzestartdate", Value:_startanalyzestartdate, Compare:">="})
    } else if  _endanalyzestartdate != "" {          
        args = append(args, models.Where{Column:"analyzestartdate", Value:_endanalyzestartdate, Compare:"<="})            
    }
    _startanalyzeenddate := c.Get("startanalyzeenddate")
    _endanalyzeenddate := c.Get("endanalyzeenddate")

    if _startanalyzeenddate != "" && _endanalyzeenddate != "" {        
        var v [2]string
        v[0] = _startanalyzeenddate
        v[1] = _endanalyzeenddate  
        args = append(args, models.Where{Column:"analyzeenddate", Value:v, Compare:"between"})    
    } else if  _startanalyzeenddate != "" {          
        args = append(args, models.Where{Column:"analyzeenddate", Value:_startanalyzeenddate, Compare:">="})
    } else if  _endanalyzeenddate != "" {          
        args = append(args, models.Where{Column:"analyzeenddate", Value:_endanalyzeenddate, Compare:"<="})            
    }
    _startratingstartdate := c.Get("startratingstartdate")
    _endratingstartdate := c.Get("endratingstartdate")

    if _startratingstartdate != "" && _endratingstartdate != "" {        
        var v [2]string
        v[0] = _startratingstartdate
        v[1] = _endratingstartdate  
        args = append(args, models.Where{Column:"ratingstartdate", Value:v, Compare:"between"})    
    } else if  _startratingstartdate != "" {          
        args = append(args, models.Where{Column:"ratingstartdate", Value:_startratingstartdate, Compare:">="})
    } else if  _endratingstartdate != "" {          
        args = append(args, models.Where{Column:"ratingstartdate", Value:_endratingstartdate, Compare:"<="})            
    }
    _startratingenddate := c.Get("startratingenddate")
    _endratingenddate := c.Get("endratingenddate")

    if _startratingenddate != "" && _endratingenddate != "" {        
        var v [2]string
        v[0] = _startratingenddate
        v[1] = _endratingenddate  
        args = append(args, models.Where{Column:"ratingenddate", Value:v, Compare:"between"})    
    } else if  _startratingenddate != "" {          
        args = append(args, models.Where{Column:"ratingenddate", Value:_startratingenddate, Compare:">="})
    } else if  _endratingenddate != "" {          
        args = append(args, models.Where{Column:"ratingenddate", Value:_endratingenddate, Compare:"<="})            
    }
    _startwritestartdate := c.Get("startwritestartdate")
    _endwritestartdate := c.Get("endwritestartdate")

    if _startwritestartdate != "" && _endwritestartdate != "" {        
        var v [2]string
        v[0] = _startwritestartdate
        v[1] = _endwritestartdate  
        args = append(args, models.Where{Column:"writestartdate", Value:v, Compare:"between"})    
    } else if  _startwritestartdate != "" {          
        args = append(args, models.Where{Column:"writestartdate", Value:_startwritestartdate, Compare:">="})
    } else if  _endwritestartdate != "" {          
        args = append(args, models.Where{Column:"writestartdate", Value:_endwritestartdate, Compare:"<="})            
    }
    _startwriteenddate := c.Get("startwriteenddate")
    _endwriteenddate := c.Get("endwriteenddate")

    if _startwriteenddate != "" && _endwriteenddate != "" {        
        var v [2]string
        v[0] = _startwriteenddate
        v[1] = _endwriteenddate  
        args = append(args, models.Where{Column:"writeenddate", Value:v, Compare:"between"})    
    } else if  _startwriteenddate != "" {          
        args = append(args, models.Where{Column:"writeenddate", Value:_startwriteenddate, Compare:">="})
    } else if  _endwriteenddate != "" {          
        args = append(args, models.Where{Column:"writeenddate", Value:_endwriteenddate, Compare:"<="})            
    }
    _startprintstartdate := c.Get("startprintstartdate")
    _endprintstartdate := c.Get("endprintstartdate")

    if _startprintstartdate != "" && _endprintstartdate != "" {        
        var v [2]string
        v[0] = _startprintstartdate
        v[1] = _endprintstartdate  
        args = append(args, models.Where{Column:"printstartdate", Value:v, Compare:"between"})    
    } else if  _startprintstartdate != "" {          
        args = append(args, models.Where{Column:"printstartdate", Value:_startprintstartdate, Compare:">="})
    } else if  _endprintstartdate != "" {          
        args = append(args, models.Where{Column:"printstartdate", Value:_endprintstartdate, Compare:"<="})            
    }
    _startprintenddate := c.Get("startprintenddate")
    _endprintenddate := c.Get("endprintenddate")

    if _startprintenddate != "" && _endprintenddate != "" {        
        var v [2]string
        v[0] = _startprintenddate
        v[1] = _endprintenddate  
        args = append(args, models.Where{Column:"printenddate", Value:v, Compare:"between"})    
    } else if  _startprintenddate != "" {          
        args = append(args, models.Where{Column:"printenddate", Value:_startprintenddate, Compare:">="})
    } else if  _endprintenddate != "" {          
        args = append(args, models.Where{Column:"printenddate", Value:_endprintenddate, Compare:"<="})            
    }
    _blueprint1 := c.Geti("blueprint1")
    if _blueprint1 != 0 {
        args = append(args, models.Where{Column:"blueprint1", Value:_blueprint1, Compare:"="})    
    }
    _blueprint2 := c.Geti("blueprint2")
    if _blueprint2 != 0 {
        args = append(args, models.Where{Column:"blueprint2", Value:_blueprint2, Compare:"="})    
    }
    _blueprint3 := c.Geti("blueprint3")
    if _blueprint3 != 0 {
        args = append(args, models.Where{Column:"blueprint3", Value:_blueprint3, Compare:"="})    
    }
    _blueprint4 := c.Geti("blueprint4")
    if _blueprint4 != 0 {
        args = append(args, models.Where{Column:"blueprint4", Value:_blueprint4, Compare:"="})    
    }
    _blueprint5 := c.Geti("blueprint5")
    if _blueprint5 != 0 {
        args = append(args, models.Where{Column:"blueprint5", Value:_blueprint5, Compare:"="})    
    }
    _blueprint6 := c.Geti("blueprint6")
    if _blueprint6 != 0 {
        args = append(args, models.Where{Column:"blueprint6", Value:_blueprint6, Compare:"="})    
    }
    _blueprint7 := c.Geti("blueprint7")
    if _blueprint7 != 0 {
        args = append(args, models.Where{Column:"blueprint7", Value:_blueprint7, Compare:"="})    
    }
    _blueprint8 := c.Geti("blueprint8")
    if _blueprint8 != 0 {
        args = append(args, models.Where{Column:"blueprint8", Value:_blueprint8, Compare:"="})    
    }
    _blueprint9 := c.Geti("blueprint9")
    if _blueprint9 != 0 {
        args = append(args, models.Where{Column:"blueprint9", Value:_blueprint9, Compare:"="})    
    }
    _blueprint10 := c.Get("blueprint10")
    if _blueprint10 != "" {
        args = append(args, models.Where{Column:"blueprint10", Value:_blueprint10, Compare:"="})
    }
    _blueprint11 := c.Geti("blueprint11")
    if _blueprint11 != 0 {
        args = append(args, models.Where{Column:"blueprint11", Value:_blueprint11, Compare:"="})    
    }
    _blueprint1save := c.Geti("blueprint1save")
    if _blueprint1save != 0 {
        args = append(args, models.Where{Column:"blueprint1save", Value:_blueprint1save, Compare:"="})    
    }
    _owner := c.Get("owner")
    if _owner != "" {
        args = append(args, models.Where{Column:"owner", Value:_owner, Compare:"="})
    }
    _manager := c.Get("manager")
    if _manager != "" {
        args = append(args, models.Where{Column:"manager", Value:_manager, Compare:"="})
    }
    _agent := c.Get("agent")
    if _agent != "" {
        args = append(args, models.Where{Column:"agent", Value:_agent, Compare:"="})
    }
    _result1 := c.Geti("result1")
    if _result1 != 0 {
        args = append(args, models.Where{Column:"result1", Value:_result1, Compare:"="})    
    }
    _result2 := c.Geti("result2")
    if _result2 != 0 {
        args = append(args, models.Where{Column:"result2", Value:_result2, Compare:"="})    
    }
    _result3 := c.Geti("result3")
    if _result3 != 0 {
        args = append(args, models.Where{Column:"result3", Value:_result3, Compare:"="})    
    }
    _result4 := c.Geti("result4")
    if _result4 != 0 {
        args = append(args, models.Where{Column:"result4", Value:_result4, Compare:"="})    
    }
    _result5 := c.Geti("result5")
    if _result5 != 0 {
        args = append(args, models.Where{Column:"result5", Value:_result5, Compare:"="})    
    }
    _resulttext1 := c.Get("resulttext1")
    if _resulttext1 != "" {
        args = append(args, models.Where{Column:"resulttext1", Value:_resulttext1, Compare:"="})
    }
    _resulttext2 := c.Get("resulttext2")
    if _resulttext2 != "" {
        args = append(args, models.Where{Column:"resulttext2", Value:_resulttext2, Compare:"="})
    }
    _resulttext3 := c.Get("resulttext3")
    if _resulttext3 != "" {
        args = append(args, models.Where{Column:"resulttext3", Value:_resulttext3, Compare:"="})
    }
    _resulttext4 := c.Get("resulttext4")
    if _resulttext4 != "" {
        args = append(args, models.Where{Column:"resulttext4", Value:_resulttext4, Compare:"="})
    }
    _resulttext5 := c.Get("resulttext5")
    if _resulttext5 != "" {
        args = append(args, models.Where{Column:"resulttext5", Value:_resulttext5, Compare:"="})
    }
    _past := c.Get("past")
    if _past != "" {
        args = append(args, models.Where{Column:"past", Value:_past, Compare:"="})
    }
    _category := c.Geti("category")
    if _category != 0 {
        args = append(args, models.Where{Column:"category", Value:_category, Compare:"="})    
    }
    _user := c.Geti64("user")
    if _user != 0 {
        args = append(args, models.Where{Column:"user", Value:_user, Compare:"="})    
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
    

    
    if c.Session.Level < 3 {
    
        args = append(args, models.Where{Column:"apt", Value:c.Session.Apt, Compare:"="})    
    
    }
    
    
    total := manager.Count(args)
	c.Set("total", total)
}


func (c *PeriodicController) Read(id int64) {
    
    if c.Session == nil {
        c.Result["code"] = "auth error"
        return
    }
    if c.Session.Level < 3 { 
    
        if c.Session.Apt == 0 {
            c.Result["code"] = "auth error"
            return
        }    
    
    }
    
    
	conn := c.NewConnection()

	manager := models.NewPeriodicManager(conn)
	item := manager.Get(id)

    
    if c.Session.Level < 3 {
    
        if item.Apt != c.Session.Apt {
            c.Result["code"] = "auth error"
            return
        }
    
    }
    
    
    c.Set("item", item)
}

func (c *PeriodicController) Index(page int, pagesize int) {
    
    if c.Session == nil {
        c.Result["code"] = "auth error"
        return
    }

    if c.Session.Level < 3 {
    
        if c.Session.Apt == 0 {
            c.Result["code"] = "auth error"
            return
        }
    
    }
    
    
	conn := c.NewConnection()

	manager := models.NewPeriodicManager(conn)

    var args []interface{}
    
    _name := c.Get("name")
    if _name != "" {
        args = append(args, models.Where{Column:"name", Value:_name, Compare:"like"})
        
    }
    _aptname := c.Get("aptname")
    if _aptname != "" {
        args = append(args, models.Where{Column:"aptname", Value:_aptname, Compare:"="})
    }
    _taskrange := c.Get("taskrange")
    if _taskrange != "" {
        args = append(args, models.Where{Column:"taskrange", Value:_taskrange, Compare:"="})
    }
    _startreportdate := c.Get("startreportdate")
    _endreportdate := c.Get("endreportdate")
    if _startreportdate != "" && _endreportdate != "" {        
        var v [2]string
        v[0] = _startreportdate
        v[1] = _endreportdate  
        args = append(args, models.Where{Column:"reportdate", Value:v, Compare:"between"})    
    } else if  _startreportdate != "" {          
        args = append(args, models.Where{Column:"reportdate", Value:_startreportdate, Compare:">="})
    } else if  _endreportdate != "" {          
        args = append(args, models.Where{Column:"reportdate", Value:_endreportdate, Compare:"<="})            
    }
    _startstartdate := c.Get("startstartdate")
    _endstartdate := c.Get("endstartdate")
    if _startstartdate != "" && _endstartdate != "" {        
        var v [2]string
        v[0] = _startstartdate
        v[1] = _endstartdate  
        args = append(args, models.Where{Column:"startdate", Value:v, Compare:"between"})    
    } else if  _startstartdate != "" {          
        args = append(args, models.Where{Column:"startdate", Value:_startstartdate, Compare:">="})
    } else if  _endstartdate != "" {          
        args = append(args, models.Where{Column:"startdate", Value:_endstartdate, Compare:"<="})            
    }
    _startenddate := c.Get("startenddate")
    _endenddate := c.Get("endenddate")
    if _startenddate != "" && _endenddate != "" {        
        var v [2]string
        v[0] = _startenddate
        v[1] = _endenddate  
        args = append(args, models.Where{Column:"enddate", Value:v, Compare:"between"})    
    } else if  _startenddate != "" {          
        args = append(args, models.Where{Column:"enddate", Value:_startenddate, Compare:">="})
    } else if  _endenddate != "" {          
        args = append(args, models.Where{Column:"enddate", Value:_endenddate, Compare:"<="})            
    }
    _supply := c.Get("supply")
    if _supply != "" {
        args = append(args, models.Where{Column:"supply", Value:_supply, Compare:"="})
    }
    _contract := c.Get("contract")
    if _contract != "" {
        args = append(args, models.Where{Column:"contract", Value:_contract, Compare:"="})
    }
    _price := c.Get("price")
    if _price != "" {
        args = append(args, models.Where{Column:"price", Value:_price, Compare:"="})
    }
    _safetygrade := c.Get("safetygrade")
    if _safetygrade != "" {
        args = append(args, models.Where{Column:"safetygrade", Value:_safetygrade, Compare:"="})
    }
    _status := c.Geti("status")
    if _status != 0 {
        args = append(args, models.Where{Column:"status", Value:_status, Compare:"="})    
    }
    _startprestartdate := c.Get("startprestartdate")
    _endprestartdate := c.Get("endprestartdate")
    if _startprestartdate != "" && _endprestartdate != "" {        
        var v [2]string
        v[0] = _startprestartdate
        v[1] = _endprestartdate  
        args = append(args, models.Where{Column:"prestartdate", Value:v, Compare:"between"})    
    } else if  _startprestartdate != "" {          
        args = append(args, models.Where{Column:"prestartdate", Value:_startprestartdate, Compare:">="})
    } else if  _endprestartdate != "" {          
        args = append(args, models.Where{Column:"prestartdate", Value:_endprestartdate, Compare:"<="})            
    }
    _startpreenddate := c.Get("startpreenddate")
    _endpreenddate := c.Get("endpreenddate")
    if _startpreenddate != "" && _endpreenddate != "" {        
        var v [2]string
        v[0] = _startpreenddate
        v[1] = _endpreenddate  
        args = append(args, models.Where{Column:"preenddate", Value:v, Compare:"between"})    
    } else if  _startpreenddate != "" {          
        args = append(args, models.Where{Column:"preenddate", Value:_startpreenddate, Compare:">="})
    } else if  _endpreenddate != "" {          
        args = append(args, models.Where{Column:"preenddate", Value:_endpreenddate, Compare:"<="})            
    }
    _startresearchstartdate := c.Get("startresearchstartdate")
    _endresearchstartdate := c.Get("endresearchstartdate")
    if _startresearchstartdate != "" && _endresearchstartdate != "" {        
        var v [2]string
        v[0] = _startresearchstartdate
        v[1] = _endresearchstartdate  
        args = append(args, models.Where{Column:"researchstartdate", Value:v, Compare:"between"})    
    } else if  _startresearchstartdate != "" {          
        args = append(args, models.Where{Column:"researchstartdate", Value:_startresearchstartdate, Compare:">="})
    } else if  _endresearchstartdate != "" {          
        args = append(args, models.Where{Column:"researchstartdate", Value:_endresearchstartdate, Compare:"<="})            
    }
    _startresearchenddate := c.Get("startresearchenddate")
    _endresearchenddate := c.Get("endresearchenddate")
    if _startresearchenddate != "" && _endresearchenddate != "" {        
        var v [2]string
        v[0] = _startresearchenddate
        v[1] = _endresearchenddate  
        args = append(args, models.Where{Column:"researchenddate", Value:v, Compare:"between"})    
    } else if  _startresearchenddate != "" {          
        args = append(args, models.Where{Column:"researchenddate", Value:_startresearchenddate, Compare:">="})
    } else if  _endresearchenddate != "" {          
        args = append(args, models.Where{Column:"researchenddate", Value:_endresearchenddate, Compare:"<="})            
    }
    _startanalyzestartdate := c.Get("startanalyzestartdate")
    _endanalyzestartdate := c.Get("endanalyzestartdate")
    if _startanalyzestartdate != "" && _endanalyzestartdate != "" {        
        var v [2]string
        v[0] = _startanalyzestartdate
        v[1] = _endanalyzestartdate  
        args = append(args, models.Where{Column:"analyzestartdate", Value:v, Compare:"between"})    
    } else if  _startanalyzestartdate != "" {          
        args = append(args, models.Where{Column:"analyzestartdate", Value:_startanalyzestartdate, Compare:">="})
    } else if  _endanalyzestartdate != "" {          
        args = append(args, models.Where{Column:"analyzestartdate", Value:_endanalyzestartdate, Compare:"<="})            
    }
    _startanalyzeenddate := c.Get("startanalyzeenddate")
    _endanalyzeenddate := c.Get("endanalyzeenddate")
    if _startanalyzeenddate != "" && _endanalyzeenddate != "" {        
        var v [2]string
        v[0] = _startanalyzeenddate
        v[1] = _endanalyzeenddate  
        args = append(args, models.Where{Column:"analyzeenddate", Value:v, Compare:"between"})    
    } else if  _startanalyzeenddate != "" {          
        args = append(args, models.Where{Column:"analyzeenddate", Value:_startanalyzeenddate, Compare:">="})
    } else if  _endanalyzeenddate != "" {          
        args = append(args, models.Where{Column:"analyzeenddate", Value:_endanalyzeenddate, Compare:"<="})            
    }
    _startratingstartdate := c.Get("startratingstartdate")
    _endratingstartdate := c.Get("endratingstartdate")
    if _startratingstartdate != "" && _endratingstartdate != "" {        
        var v [2]string
        v[0] = _startratingstartdate
        v[1] = _endratingstartdate  
        args = append(args, models.Where{Column:"ratingstartdate", Value:v, Compare:"between"})    
    } else if  _startratingstartdate != "" {          
        args = append(args, models.Where{Column:"ratingstartdate", Value:_startratingstartdate, Compare:">="})
    } else if  _endratingstartdate != "" {          
        args = append(args, models.Where{Column:"ratingstartdate", Value:_endratingstartdate, Compare:"<="})            
    }
    _startratingenddate := c.Get("startratingenddate")
    _endratingenddate := c.Get("endratingenddate")
    if _startratingenddate != "" && _endratingenddate != "" {        
        var v [2]string
        v[0] = _startratingenddate
        v[1] = _endratingenddate  
        args = append(args, models.Where{Column:"ratingenddate", Value:v, Compare:"between"})    
    } else if  _startratingenddate != "" {          
        args = append(args, models.Where{Column:"ratingenddate", Value:_startratingenddate, Compare:">="})
    } else if  _endratingenddate != "" {          
        args = append(args, models.Where{Column:"ratingenddate", Value:_endratingenddate, Compare:"<="})            
    }
    _startwritestartdate := c.Get("startwritestartdate")
    _endwritestartdate := c.Get("endwritestartdate")
    if _startwritestartdate != "" && _endwritestartdate != "" {        
        var v [2]string
        v[0] = _startwritestartdate
        v[1] = _endwritestartdate  
        args = append(args, models.Where{Column:"writestartdate", Value:v, Compare:"between"})    
    } else if  _startwritestartdate != "" {          
        args = append(args, models.Where{Column:"writestartdate", Value:_startwritestartdate, Compare:">="})
    } else if  _endwritestartdate != "" {          
        args = append(args, models.Where{Column:"writestartdate", Value:_endwritestartdate, Compare:"<="})            
    }
    _startwriteenddate := c.Get("startwriteenddate")
    _endwriteenddate := c.Get("endwriteenddate")
    if _startwriteenddate != "" && _endwriteenddate != "" {        
        var v [2]string
        v[0] = _startwriteenddate
        v[1] = _endwriteenddate  
        args = append(args, models.Where{Column:"writeenddate", Value:v, Compare:"between"})    
    } else if  _startwriteenddate != "" {          
        args = append(args, models.Where{Column:"writeenddate", Value:_startwriteenddate, Compare:">="})
    } else if  _endwriteenddate != "" {          
        args = append(args, models.Where{Column:"writeenddate", Value:_endwriteenddate, Compare:"<="})            
    }
    _startprintstartdate := c.Get("startprintstartdate")
    _endprintstartdate := c.Get("endprintstartdate")
    if _startprintstartdate != "" && _endprintstartdate != "" {        
        var v [2]string
        v[0] = _startprintstartdate
        v[1] = _endprintstartdate  
        args = append(args, models.Where{Column:"printstartdate", Value:v, Compare:"between"})    
    } else if  _startprintstartdate != "" {          
        args = append(args, models.Where{Column:"printstartdate", Value:_startprintstartdate, Compare:">="})
    } else if  _endprintstartdate != "" {          
        args = append(args, models.Where{Column:"printstartdate", Value:_endprintstartdate, Compare:"<="})            
    }
    _startprintenddate := c.Get("startprintenddate")
    _endprintenddate := c.Get("endprintenddate")
    if _startprintenddate != "" && _endprintenddate != "" {        
        var v [2]string
        v[0] = _startprintenddate
        v[1] = _endprintenddate  
        args = append(args, models.Where{Column:"printenddate", Value:v, Compare:"between"})    
    } else if  _startprintenddate != "" {          
        args = append(args, models.Where{Column:"printenddate", Value:_startprintenddate, Compare:">="})
    } else if  _endprintenddate != "" {          
        args = append(args, models.Where{Column:"printenddate", Value:_endprintenddate, Compare:"<="})            
    }
    _blueprint1 := c.Geti("blueprint1")
    if _blueprint1 != 0 {
        args = append(args, models.Where{Column:"blueprint1", Value:_blueprint1, Compare:"="})    
    }
    _blueprint2 := c.Geti("blueprint2")
    if _blueprint2 != 0 {
        args = append(args, models.Where{Column:"blueprint2", Value:_blueprint2, Compare:"="})    
    }
    _blueprint3 := c.Geti("blueprint3")
    if _blueprint3 != 0 {
        args = append(args, models.Where{Column:"blueprint3", Value:_blueprint3, Compare:"="})    
    }
    _blueprint4 := c.Geti("blueprint4")
    if _blueprint4 != 0 {
        args = append(args, models.Where{Column:"blueprint4", Value:_blueprint4, Compare:"="})    
    }
    _blueprint5 := c.Geti("blueprint5")
    if _blueprint5 != 0 {
        args = append(args, models.Where{Column:"blueprint5", Value:_blueprint5, Compare:"="})    
    }
    _blueprint6 := c.Geti("blueprint6")
    if _blueprint6 != 0 {
        args = append(args, models.Where{Column:"blueprint6", Value:_blueprint6, Compare:"="})    
    }
    _blueprint7 := c.Geti("blueprint7")
    if _blueprint7 != 0 {
        args = append(args, models.Where{Column:"blueprint7", Value:_blueprint7, Compare:"="})    
    }
    _blueprint8 := c.Geti("blueprint8")
    if _blueprint8 != 0 {
        args = append(args, models.Where{Column:"blueprint8", Value:_blueprint8, Compare:"="})    
    }
    _blueprint9 := c.Geti("blueprint9")
    if _blueprint9 != 0 {
        args = append(args, models.Where{Column:"blueprint9", Value:_blueprint9, Compare:"="})    
    }
    _blueprint10 := c.Get("blueprint10")
    if _blueprint10 != "" {
        args = append(args, models.Where{Column:"blueprint10", Value:_blueprint10, Compare:"="})
    }
    _blueprint11 := c.Geti("blueprint11")
    if _blueprint11 != 0 {
        args = append(args, models.Where{Column:"blueprint11", Value:_blueprint11, Compare:"="})    
    }
    _blueprint1save := c.Geti("blueprint1save")
    if _blueprint1save != 0 {
        args = append(args, models.Where{Column:"blueprint1save", Value:_blueprint1save, Compare:"="})    
    }
    _owner := c.Get("owner")
    if _owner != "" {
        args = append(args, models.Where{Column:"owner", Value:_owner, Compare:"="})
    }
    _manager := c.Get("manager")
    if _manager != "" {
        args = append(args, models.Where{Column:"manager", Value:_manager, Compare:"="})
    }
    _agent := c.Get("agent")
    if _agent != "" {
        args = append(args, models.Where{Column:"agent", Value:_agent, Compare:"="})
    }
    _result1 := c.Geti("result1")
    if _result1 != 0 {
        args = append(args, models.Where{Column:"result1", Value:_result1, Compare:"="})    
    }
    _result2 := c.Geti("result2")
    if _result2 != 0 {
        args = append(args, models.Where{Column:"result2", Value:_result2, Compare:"="})    
    }
    _result3 := c.Geti("result3")
    if _result3 != 0 {
        args = append(args, models.Where{Column:"result3", Value:_result3, Compare:"="})    
    }
    _result4 := c.Geti("result4")
    if _result4 != 0 {
        args = append(args, models.Where{Column:"result4", Value:_result4, Compare:"="})    
    }
    _result5 := c.Geti("result5")
    if _result5 != 0 {
        args = append(args, models.Where{Column:"result5", Value:_result5, Compare:"="})    
    }
    _resulttext1 := c.Get("resulttext1")
    if _resulttext1 != "" {
        args = append(args, models.Where{Column:"resulttext1", Value:_resulttext1, Compare:"="})
    }
    _resulttext2 := c.Get("resulttext2")
    if _resulttext2 != "" {
        args = append(args, models.Where{Column:"resulttext2", Value:_resulttext2, Compare:"="})
    }
    _resulttext3 := c.Get("resulttext3")
    if _resulttext3 != "" {
        args = append(args, models.Where{Column:"resulttext3", Value:_resulttext3, Compare:"="})
    }
    _resulttext4 := c.Get("resulttext4")
    if _resulttext4 != "" {
        args = append(args, models.Where{Column:"resulttext4", Value:_resulttext4, Compare:"="})
    }
    _resulttext5 := c.Get("resulttext5")
    if _resulttext5 != "" {
        args = append(args, models.Where{Column:"resulttext5", Value:_resulttext5, Compare:"="})
    }
    _past := c.Get("past")
    if _past != "" {
        args = append(args, models.Where{Column:"past", Value:_past, Compare:"="})
    }
    _category := c.Geti("category")
    if _category != 0 {
        args = append(args, models.Where{Column:"category", Value:_category, Compare:"="})    
    }
    _user := c.Geti64("user")
    if _user != 0 {
        args = append(args, models.Where{Column:"user", Value:_user, Compare:"="})    
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
    

    
    if c.Session.Level < 3 {
    
        args = append(args, models.Where{Column:"apt", Value:c.Session.Apt, Compare:"="})    
    
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





func (c *PeriodicController) Sum() {
    
    if c.Session == nil {
        c.Result["code"] = "auth error"
        return
    }

    if c.Session.Level < 3 {
    
    if c.Session.Apt == 0 {
        c.Result["code"] = "auth error"
        return
    }
    
    }
    
    
	conn := c.NewConnection()

	manager := models.NewPeriodicManager(conn)

    var args []interface{}
    
    _name := c.Get("name")
    if _name != "" {
        args = append(args, models.Where{Column:"name", Value:_name, Compare:"like"})
        
    }
    _aptname := c.Get("aptname")
    if _aptname != "" {
        args = append(args, models.Where{Column:"aptname", Value:_aptname, Compare:"like"})
    }
    _taskrange := c.Get("taskrange")
    if _taskrange != "" {
        args = append(args, models.Where{Column:"taskrange", Value:_taskrange, Compare:"like"})
    }
    _startreportdate := c.Get("startreportdate")
    _endreportdate := c.Get("endreportdate")
    if _startreportdate != "" && _endreportdate != "" {        
        var v [2]string
        v[0] = _startreportdate
        v[1] = _endreportdate  
        args = append(args, models.Where{Column:"reportdate", Value:v, Compare:"between"})    
    } else if  _startreportdate != "" {          
        args = append(args, models.Where{Column:"reportdate", Value:_startreportdate, Compare:">="})
    } else if  _endreportdate != "" {          
        args = append(args, models.Where{Column:"reportdate", Value:_endreportdate, Compare:"<="})            
    }
    _startstartdate := c.Get("startstartdate")
    _endstartdate := c.Get("endstartdate")
    if _startstartdate != "" && _endstartdate != "" {        
        var v [2]string
        v[0] = _startstartdate
        v[1] = _endstartdate  
        args = append(args, models.Where{Column:"startdate", Value:v, Compare:"between"})    
    } else if  _startstartdate != "" {          
        args = append(args, models.Where{Column:"startdate", Value:_startstartdate, Compare:">="})
    } else if  _endstartdate != "" {          
        args = append(args, models.Where{Column:"startdate", Value:_endstartdate, Compare:"<="})            
    }
    _startenddate := c.Get("startenddate")
    _endenddate := c.Get("endenddate")
    if _startenddate != "" && _endenddate != "" {        
        var v [2]string
        v[0] = _startenddate
        v[1] = _endenddate  
        args = append(args, models.Where{Column:"enddate", Value:v, Compare:"between"})    
    } else if  _startenddate != "" {          
        args = append(args, models.Where{Column:"enddate", Value:_startenddate, Compare:">="})
    } else if  _endenddate != "" {          
        args = append(args, models.Where{Column:"enddate", Value:_endenddate, Compare:"<="})            
    }
    _supply := c.Get("supply")
    if _supply != "" {
        args = append(args, models.Where{Column:"supply", Value:_supply, Compare:"like"})
    }
    _contract := c.Get("contract")
    if _contract != "" {
        args = append(args, models.Where{Column:"contract", Value:_contract, Compare:"like"})
    }
    _price := c.Get("price")
    if _price != "" {
        args = append(args, models.Where{Column:"price", Value:_price, Compare:"like"})
    }
    _safetygrade := c.Get("safetygrade")
    if _safetygrade != "" {
        args = append(args, models.Where{Column:"safetygrade", Value:_safetygrade, Compare:"like"})
    }
    _status := c.Geti("status")
    if _status != 0 {
        args = append(args, models.Where{Column:"status", Value:_status, Compare:"="})    
    }
    _startprestartdate := c.Get("startprestartdate")
    _endprestartdate := c.Get("endprestartdate")
    if _startprestartdate != "" && _endprestartdate != "" {        
        var v [2]string
        v[0] = _startprestartdate
        v[1] = _endprestartdate  
        args = append(args, models.Where{Column:"prestartdate", Value:v, Compare:"between"})    
    } else if  _startprestartdate != "" {          
        args = append(args, models.Where{Column:"prestartdate", Value:_startprestartdate, Compare:">="})
    } else if  _endprestartdate != "" {          
        args = append(args, models.Where{Column:"prestartdate", Value:_endprestartdate, Compare:"<="})            
    }
    _startpreenddate := c.Get("startpreenddate")
    _endpreenddate := c.Get("endpreenddate")
    if _startpreenddate != "" && _endpreenddate != "" {        
        var v [2]string
        v[0] = _startpreenddate
        v[1] = _endpreenddate  
        args = append(args, models.Where{Column:"preenddate", Value:v, Compare:"between"})    
    } else if  _startpreenddate != "" {          
        args = append(args, models.Where{Column:"preenddate", Value:_startpreenddate, Compare:">="})
    } else if  _endpreenddate != "" {          
        args = append(args, models.Where{Column:"preenddate", Value:_endpreenddate, Compare:"<="})            
    }
    _startresearchstartdate := c.Get("startresearchstartdate")
    _endresearchstartdate := c.Get("endresearchstartdate")
    if _startresearchstartdate != "" && _endresearchstartdate != "" {        
        var v [2]string
        v[0] = _startresearchstartdate
        v[1] = _endresearchstartdate  
        args = append(args, models.Where{Column:"researchstartdate", Value:v, Compare:"between"})    
    } else if  _startresearchstartdate != "" {          
        args = append(args, models.Where{Column:"researchstartdate", Value:_startresearchstartdate, Compare:">="})
    } else if  _endresearchstartdate != "" {          
        args = append(args, models.Where{Column:"researchstartdate", Value:_endresearchstartdate, Compare:"<="})            
    }
    _startresearchenddate := c.Get("startresearchenddate")
    _endresearchenddate := c.Get("endresearchenddate")
    if _startresearchenddate != "" && _endresearchenddate != "" {        
        var v [2]string
        v[0] = _startresearchenddate
        v[1] = _endresearchenddate  
        args = append(args, models.Where{Column:"researchenddate", Value:v, Compare:"between"})    
    } else if  _startresearchenddate != "" {          
        args = append(args, models.Where{Column:"researchenddate", Value:_startresearchenddate, Compare:">="})
    } else if  _endresearchenddate != "" {          
        args = append(args, models.Where{Column:"researchenddate", Value:_endresearchenddate, Compare:"<="})            
    }
    _startanalyzestartdate := c.Get("startanalyzestartdate")
    _endanalyzestartdate := c.Get("endanalyzestartdate")
    if _startanalyzestartdate != "" && _endanalyzestartdate != "" {        
        var v [2]string
        v[0] = _startanalyzestartdate
        v[1] = _endanalyzestartdate  
        args = append(args, models.Where{Column:"analyzestartdate", Value:v, Compare:"between"})    
    } else if  _startanalyzestartdate != "" {          
        args = append(args, models.Where{Column:"analyzestartdate", Value:_startanalyzestartdate, Compare:">="})
    } else if  _endanalyzestartdate != "" {          
        args = append(args, models.Where{Column:"analyzestartdate", Value:_endanalyzestartdate, Compare:"<="})            
    }
    _startanalyzeenddate := c.Get("startanalyzeenddate")
    _endanalyzeenddate := c.Get("endanalyzeenddate")
    if _startanalyzeenddate != "" && _endanalyzeenddate != "" {        
        var v [2]string
        v[0] = _startanalyzeenddate
        v[1] = _endanalyzeenddate  
        args = append(args, models.Where{Column:"analyzeenddate", Value:v, Compare:"between"})    
    } else if  _startanalyzeenddate != "" {          
        args = append(args, models.Where{Column:"analyzeenddate", Value:_startanalyzeenddate, Compare:">="})
    } else if  _endanalyzeenddate != "" {          
        args = append(args, models.Where{Column:"analyzeenddate", Value:_endanalyzeenddate, Compare:"<="})            
    }
    _startratingstartdate := c.Get("startratingstartdate")
    _endratingstartdate := c.Get("endratingstartdate")
    if _startratingstartdate != "" && _endratingstartdate != "" {        
        var v [2]string
        v[0] = _startratingstartdate
        v[1] = _endratingstartdate  
        args = append(args, models.Where{Column:"ratingstartdate", Value:v, Compare:"between"})    
    } else if  _startratingstartdate != "" {          
        args = append(args, models.Where{Column:"ratingstartdate", Value:_startratingstartdate, Compare:">="})
    } else if  _endratingstartdate != "" {          
        args = append(args, models.Where{Column:"ratingstartdate", Value:_endratingstartdate, Compare:"<="})            
    }
    _startratingenddate := c.Get("startratingenddate")
    _endratingenddate := c.Get("endratingenddate")
    if _startratingenddate != "" && _endratingenddate != "" {        
        var v [2]string
        v[0] = _startratingenddate
        v[1] = _endratingenddate  
        args = append(args, models.Where{Column:"ratingenddate", Value:v, Compare:"between"})    
    } else if  _startratingenddate != "" {          
        args = append(args, models.Where{Column:"ratingenddate", Value:_startratingenddate, Compare:">="})
    } else if  _endratingenddate != "" {          
        args = append(args, models.Where{Column:"ratingenddate", Value:_endratingenddate, Compare:"<="})            
    }
    _startwritestartdate := c.Get("startwritestartdate")
    _endwritestartdate := c.Get("endwritestartdate")
    if _startwritestartdate != "" && _endwritestartdate != "" {        
        var v [2]string
        v[0] = _startwritestartdate
        v[1] = _endwritestartdate  
        args = append(args, models.Where{Column:"writestartdate", Value:v, Compare:"between"})    
    } else if  _startwritestartdate != "" {          
        args = append(args, models.Where{Column:"writestartdate", Value:_startwritestartdate, Compare:">="})
    } else if  _endwritestartdate != "" {          
        args = append(args, models.Where{Column:"writestartdate", Value:_endwritestartdate, Compare:"<="})            
    }
    _startwriteenddate := c.Get("startwriteenddate")
    _endwriteenddate := c.Get("endwriteenddate")
    if _startwriteenddate != "" && _endwriteenddate != "" {        
        var v [2]string
        v[0] = _startwriteenddate
        v[1] = _endwriteenddate  
        args = append(args, models.Where{Column:"writeenddate", Value:v, Compare:"between"})    
    } else if  _startwriteenddate != "" {          
        args = append(args, models.Where{Column:"writeenddate", Value:_startwriteenddate, Compare:">="})
    } else if  _endwriteenddate != "" {          
        args = append(args, models.Where{Column:"writeenddate", Value:_endwriteenddate, Compare:"<="})            
    }
    _startprintstartdate := c.Get("startprintstartdate")
    _endprintstartdate := c.Get("endprintstartdate")
    if _startprintstartdate != "" && _endprintstartdate != "" {        
        var v [2]string
        v[0] = _startprintstartdate
        v[1] = _endprintstartdate  
        args = append(args, models.Where{Column:"printstartdate", Value:v, Compare:"between"})    
    } else if  _startprintstartdate != "" {          
        args = append(args, models.Where{Column:"printstartdate", Value:_startprintstartdate, Compare:">="})
    } else if  _endprintstartdate != "" {          
        args = append(args, models.Where{Column:"printstartdate", Value:_endprintstartdate, Compare:"<="})            
    }
    _startprintenddate := c.Get("startprintenddate")
    _endprintenddate := c.Get("endprintenddate")
    if _startprintenddate != "" && _endprintenddate != "" {        
        var v [2]string
        v[0] = _startprintenddate
        v[1] = _endprintenddate  
        args = append(args, models.Where{Column:"printenddate", Value:v, Compare:"between"})    
    } else if  _startprintenddate != "" {          
        args = append(args, models.Where{Column:"printenddate", Value:_startprintenddate, Compare:">="})
    } else if  _endprintenddate != "" {          
        args = append(args, models.Where{Column:"printenddate", Value:_endprintenddate, Compare:"<="})            
    }
    _blueprint1 := c.Geti("blueprint1")
    if _blueprint1 != 0 {
        args = append(args, models.Where{Column:"blueprint1", Value:_blueprint1, Compare:"="})    
    }
    _blueprint2 := c.Geti("blueprint2")
    if _blueprint2 != 0 {
        args = append(args, models.Where{Column:"blueprint2", Value:_blueprint2, Compare:"="})    
    }
    _blueprint3 := c.Geti("blueprint3")
    if _blueprint3 != 0 {
        args = append(args, models.Where{Column:"blueprint3", Value:_blueprint3, Compare:"="})    
    }
    _blueprint4 := c.Geti("blueprint4")
    if _blueprint4 != 0 {
        args = append(args, models.Where{Column:"blueprint4", Value:_blueprint4, Compare:"="})    
    }
    _blueprint5 := c.Geti("blueprint5")
    if _blueprint5 != 0 {
        args = append(args, models.Where{Column:"blueprint5", Value:_blueprint5, Compare:"="})    
    }
    _blueprint6 := c.Geti("blueprint6")
    if _blueprint6 != 0 {
        args = append(args, models.Where{Column:"blueprint6", Value:_blueprint6, Compare:"="})    
    }
    _blueprint7 := c.Geti("blueprint7")
    if _blueprint7 != 0 {
        args = append(args, models.Where{Column:"blueprint7", Value:_blueprint7, Compare:"="})    
    }
    _blueprint8 := c.Geti("blueprint8")
    if _blueprint8 != 0 {
        args = append(args, models.Where{Column:"blueprint8", Value:_blueprint8, Compare:"="})    
    }
    _blueprint9 := c.Geti("blueprint9")
    if _blueprint9 != 0 {
        args = append(args, models.Where{Column:"blueprint9", Value:_blueprint9, Compare:"="})    
    }
    _blueprint10 := c.Get("blueprint10")
    if _blueprint10 != "" {
        args = append(args, models.Where{Column:"blueprint10", Value:_blueprint10, Compare:"like"})
    }
    _blueprint11 := c.Geti("blueprint11")
    if _blueprint11 != 0 {
        args = append(args, models.Where{Column:"blueprint11", Value:_blueprint11, Compare:"="})    
    }
    _blueprint1save := c.Geti("blueprint1save")
    if _blueprint1save != 0 {
        args = append(args, models.Where{Column:"blueprint1save", Value:_blueprint1save, Compare:"="})    
    }
    _owner := c.Get("owner")
    if _owner != "" {
        args = append(args, models.Where{Column:"owner", Value:_owner, Compare:"like"})
    }
    _manager := c.Get("manager")
    if _manager != "" {
        args = append(args, models.Where{Column:"manager", Value:_manager, Compare:"like"})
    }
    _agent := c.Get("agent")
    if _agent != "" {
        args = append(args, models.Where{Column:"agent", Value:_agent, Compare:"like"})
    }
    _result1 := c.Geti("result1")
    if _result1 != 0 {
        args = append(args, models.Where{Column:"result1", Value:_result1, Compare:"="})    
    }
    _result2 := c.Geti("result2")
    if _result2 != 0 {
        args = append(args, models.Where{Column:"result2", Value:_result2, Compare:"="})    
    }
    _result3 := c.Geti("result3")
    if _result3 != 0 {
        args = append(args, models.Where{Column:"result3", Value:_result3, Compare:"="})    
    }
    _result4 := c.Geti("result4")
    if _result4 != 0 {
        args = append(args, models.Where{Column:"result4", Value:_result4, Compare:"="})    
    }
    _result5 := c.Geti("result5")
    if _result5 != 0 {
        args = append(args, models.Where{Column:"result5", Value:_result5, Compare:"="})    
    }
    _resulttext1 := c.Get("resulttext1")
    if _resulttext1 != "" {
        args = append(args, models.Where{Column:"resulttext1", Value:_resulttext1, Compare:"like"})
    }
    _resulttext2 := c.Get("resulttext2")
    if _resulttext2 != "" {
        args = append(args, models.Where{Column:"resulttext2", Value:_resulttext2, Compare:"like"})
    }
    _resulttext3 := c.Get("resulttext3")
    if _resulttext3 != "" {
        args = append(args, models.Where{Column:"resulttext3", Value:_resulttext3, Compare:"like"})
    }
    _resulttext4 := c.Get("resulttext4")
    if _resulttext4 != "" {
        args = append(args, models.Where{Column:"resulttext4", Value:_resulttext4, Compare:"like"})
    }
    _resulttext5 := c.Get("resulttext5")
    if _resulttext5 != "" {
        args = append(args, models.Where{Column:"resulttext5", Value:_resulttext5, Compare:"like"})
    }
    _past := c.Get("past")
    if _past != "" {
        args = append(args, models.Where{Column:"past", Value:_past, Compare:"like"})
    }
    _category := c.Geti("category")
    if _category != 0 {
        args = append(args, models.Where{Column:"category", Value:_category, Compare:"="})    
    }
    _user := c.Geti64("user")
    if _user != 0 {
        args = append(args, models.Where{Column:"user", Value:_user, Compare:"="})    
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
    

    
    if c.Session.Level < 3 {
    
    args = append(args, models.Where{Column:"apt", Value:c.Session.Apt, Compare:"="})    
    
    }
    
    
    item := manager.Sum(args)
	c.Set("item", item)
}

