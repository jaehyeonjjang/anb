package rest


import (
	"repair/controllers"
	"repair/models"

    "strings"
)

type DetailController struct {
	controllers.Controller
}



func (c *DetailController) Insert(item *models.Detail) {
    
    
	conn := c.NewConnection()
    
	manager := models.NewDetailManager(conn)
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

func (c *DetailController) Insertbatch(item *[]models.Detail) {  
    if item == nil || len(*item) == 0 {
        return
    }

    rows := len(*item)
    
    
    
	conn := c.NewConnection()
    
	manager := models.NewDetailManager(conn)

    for i := 0; i < rows; i++ {
	    err := manager.Insert(&((*item)[i]))
        if err != nil {
            c.Set("code", "error")    
            c.Set("error", err)
            return
        }
    }
}

func (c *DetailController) Update(item *models.Detail) {
    
    
	conn := c.NewConnection()

	manager := models.NewDetailManager(conn)
    err := manager.Update(item)
    if err != nil {
        c.Set("code", "error")    
        c.Set("error", err)
        return
    }
}

func (c *DetailController) Delete(item *models.Detail) {
    
    
    conn := c.NewConnection()

	manager := models.NewDetailManager(conn)

    
	err := manager.Delete(item.Id)
    if err != nil {
        c.Set("code", "error")    
        c.Set("error", err)
    }
}

func (c *DetailController) Deletebatch(item *[]models.Detail) {
    
    
    conn := c.NewConnection()

	manager := models.NewDetailManager(conn)

    for _, v := range *item {
        
    
	    err := manager.Delete(v.Id)
        if err != nil {
            c.Set("code", "error")    
            c.Set("error", err)
            return
        }
    }
}

func (c *DetailController) Count() {
    
    
	conn := c.NewConnection()

	manager := models.NewDetailManager(conn)

    var args []interface{}
    
    _name := c.Get("name")
    if _name != "" {
        args = append(args, models.Where{Column:"name", Value:_name, Compare:"like"})
        
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
    _price := c.Geti("price")
    if _price != 0 {
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


func (c *DetailController) Read(id int64) {
    
    
	conn := c.NewConnection()

	manager := models.NewDetailManager(conn)
	item := manager.Get(id)

    
    
    c.Set("item", item)
}

func (c *DetailController) Index(page int, pagesize int) {
    
    
	conn := c.NewConnection()

	manager := models.NewDetailManager(conn)

    var args []interface{}
    
    _name := c.Get("name")
    if _name != "" {
        args = append(args, models.Where{Column:"name", Value:_name, Compare:"like"})
        
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
    _price := c.Geti("price")
    if _price != 0 {
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





func (c *DetailController) Sum() {
    
    
	conn := c.NewConnection()

	manager := models.NewDetailManager(conn)

    var args []interface{}
    
    _name := c.Get("name")
    if _name != "" {
        args = append(args, models.Where{Column:"name", Value:_name, Compare:"like"})
        
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
    _price := c.Geti("price")
    if _price != 0 {
        args = append(args, models.Where{Column:"price", Value:_price, Compare:"="})    
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

