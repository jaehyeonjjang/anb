package rest


import (
	"repair/controllers"
	"repair/models"

    "strings"
)

type BreakdownController struct {
	controllers.Controller
}



func (c *BreakdownController) CountByApt(apt int64) int {
    
    conn := c.NewConnection()

	_manager := models.NewBreakdownManager(conn)
    
    item := _manager.CountByApt(apt)
    
    
    
    c.Set("count", item)
    
    return item
    
}


func (c *BreakdownController) FindByApt(apt int64) []models.Breakdown {
    
    conn := c.NewConnection()

	_manager := models.NewBreakdownManager(conn)
    
    item := _manager.FindByApt(apt)
    
    
    c.Set("items", item)
    
    
    return item
    
}

// @Delete()
func (c *BreakdownController) DeleteByApt(apt int64) {
    
    conn := c.NewConnection()

	_manager := models.NewBreakdownManager(conn)
    
    err := _manager.DeleteByApt(apt)
    if err != nil {
        c.Set("code", "error")    
        c.Set("error", err)
        return
    }
    
}


func (c *BreakdownController) CountByAptDong(apt int64 ,dong int64) int {
    
    conn := c.NewConnection()

	_manager := models.NewBreakdownManager(conn)
    
    item := _manager.CountByAptDong(apt, dong)
    
    
    
    c.Set("count", item)
    
    return item
    
}


func (c *BreakdownController) CountByAptStandard(apt int64 ,standard int64) int {
    
    conn := c.NewConnection()

	_manager := models.NewBreakdownManager(conn)
    
    item := _manager.CountByAptStandard(apt, standard)
    
    
    
    c.Set("count", item)
    
    return item
    
}

// @Put()
func (c *BreakdownController) UpdateDuedateById(duedate int ,id int64) {
    
    conn := c.NewConnection()

	_manager := models.NewBreakdownManager(conn)
    
    err := _manager.UpdateDuedateById(duedate, id)
    if err != nil {
        c.Set("code", "error")    
        c.Set("error", err)
        return
    }
    
}

// @Put()
func (c *BreakdownController) UpdateLastdateById(lastdate int ,id int64) {
    
    conn := c.NewConnection()

	_manager := models.NewBreakdownManager(conn)
    
    err := _manager.UpdateLastdateById(lastdate, id)
    if err != nil {
        c.Set("code", "error")    
        c.Set("error", err)
        return
    }
    
}


func (c *BreakdownController) FindByCategory(category int64) []models.Breakdown {
    
    conn := c.NewConnection()

	_manager := models.NewBreakdownManager(conn)
    
    item := _manager.FindByCategory(category)
    
    
    c.Set("items", item)
    
    
    return item
    
}

// @Delete()
func (c *BreakdownController) DeleteByCategory(category int64) {
    
    conn := c.NewConnection()

	_manager := models.NewBreakdownManager(conn)
    
    err := _manager.DeleteByCategory(category)
    if err != nil {
        c.Set("code", "error")    
        c.Set("error", err)
        return
    }
    
}


func (c *BreakdownController) FindByMethod(method int64) []models.Breakdown {
    
    conn := c.NewConnection()

	_manager := models.NewBreakdownManager(conn)
    
    item := _manager.FindByMethod(method)
    
    
    c.Set("items", item)
    
    
    return item
    
}

// @Put()
func (c *BreakdownController) UpdateSubcategoryByCategory(subcategory int64 ,category int64) {
    
    conn := c.NewConnection()

	_manager := models.NewBreakdownManager(conn)
    
    err := _manager.UpdateSubcategoryByCategory(subcategory, category)
    if err != nil {
        c.Set("code", "error")    
        c.Set("error", err)
        return
    }
    
}


func (c *BreakdownController) Insert(item *models.Breakdown) {
    
    
	conn := c.NewConnection()
    
	manager := models.NewBreakdownManager(conn)
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

func (c *BreakdownController) Insertbatch(item *[]models.Breakdown) {  
    if item == nil || len(*item) == 0 {
        return
    }

    rows := len(*item)
    
    
    
	conn := c.NewConnection()
    
	manager := models.NewBreakdownManager(conn)

    for i := 0; i < rows; i++ {
	    err := manager.Insert(&((*item)[i]))
        if err != nil {
            c.Set("code", "error")    
            c.Set("error", err)
            return
        }
    }
}

func (c *BreakdownController) Update(item *models.Breakdown) {
    
    
	conn := c.NewConnection()

	manager := models.NewBreakdownManager(conn)
    err := manager.Update(item)
    if err != nil {
        c.Set("code", "error")    
        c.Set("error", err)
        return
    }
}

func (c *BreakdownController) Delete(item *models.Breakdown) {
    
    
    conn := c.NewConnection()

	manager := models.NewBreakdownManager(conn)

    
	err := manager.Delete(item.Id)
    if err != nil {
        c.Set("code", "error")    
        c.Set("error", err)
    }
}

func (c *BreakdownController) Deletebatch(item *[]models.Breakdown) {
    
    
    conn := c.NewConnection()

	manager := models.NewBreakdownManager(conn)

    for _, v := range *item {
        
    
	    err := manager.Delete(v.Id)
        if err != nil {
            c.Set("code", "error")    
            c.Set("error", err)
            return
        }
    }
}

func (c *BreakdownController) Count() {
    
    
	conn := c.NewConnection()

	manager := models.NewBreakdownManager(conn)

    var args []interface{}
    
    _topcategory := c.Geti64("topcategory")
    if _topcategory != 0 {
        args = append(args, models.Where{Column:"topcategory", Value:_topcategory, Compare:"="})    
    }
    _subcategory := c.Geti64("subcategory")
    if _subcategory != 0 {
        args = append(args, models.Where{Column:"subcategory", Value:_subcategory, Compare:"="})    
    }
    _category := c.Geti64("category")
    if _category != 0 {
        args = append(args, models.Where{Column:"category", Value:_category, Compare:"="})    
    }
    _method := c.Geti64("method")
    if _method != 0 {
        args = append(args, models.Where{Column:"method", Value:_method, Compare:"="})    
    }
    _count := c.Geti("count")
    if _count != 0 {
        args = append(args, models.Where{Column:"count", Value:_count, Compare:"="})    
    }
    _lastdate := c.Geti("lastdate")
    if _lastdate != 0 {
        args = append(args, models.Where{Column:"lastdate", Value:_lastdate, Compare:"="})    
    }
    _duedate := c.Geti("duedate")
    if _duedate != 0 {
        args = append(args, models.Where{Column:"duedate", Value:_duedate, Compare:"="})    
    }
    _remark := c.Get("remark")
    if _remark != "" {
        args = append(args, models.Where{Column:"remark", Value:_remark, Compare:"="})
    }
    _elevator := c.Geti("elevator")
    if _elevator != 0 {
        args = append(args, models.Where{Column:"elevator", Value:_elevator, Compare:"="})    
    }
    _percent := c.Geti("percent")
    if _percent != 0 {
        args = append(args, models.Where{Column:"percent", Value:_percent, Compare:"="})    
    }
    _rate := c.Geti("rate")
    if _rate != 0 {
        args = append(args, models.Where{Column:"rate", Value:_rate, Compare:"="})    
    }
    _type := c.Geti("type")
    if _type != 0 {
        args = append(args, models.Where{Column:"type", Value:_type, Compare:"="})    
    }
    _dong := c.Geti64("dong")
    if _dong != 0 {
        args = append(args, models.Where{Column:"dong", Value:_dong, Compare:"="})    
    }
    _standard := c.Geti64("standard")
    if _standard != 0 {
        args = append(args, models.Where{Column:"standard", Value:_standard, Compare:"="})    
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


func (c *BreakdownController) Read(id int64) {
    
    
	conn := c.NewConnection()

	manager := models.NewBreakdownManager(conn)
	item := manager.Get(id)

    
    
    c.Set("item", item)
}

func (c *BreakdownController) Index(page int, pagesize int) {
    
    
	conn := c.NewConnection()

	manager := models.NewBreakdownManager(conn)

    var args []interface{}
    
    _topcategory := c.Geti64("topcategory")
    if _topcategory != 0 {
        args = append(args, models.Where{Column:"topcategory", Value:_topcategory, Compare:"="})    
    }
    _subcategory := c.Geti64("subcategory")
    if _subcategory != 0 {
        args = append(args, models.Where{Column:"subcategory", Value:_subcategory, Compare:"="})    
    }
    _category := c.Geti64("category")
    if _category != 0 {
        args = append(args, models.Where{Column:"category", Value:_category, Compare:"="})    
    }
    _method := c.Geti64("method")
    if _method != 0 {
        args = append(args, models.Where{Column:"method", Value:_method, Compare:"="})    
    }
    _count := c.Geti("count")
    if _count != 0 {
        args = append(args, models.Where{Column:"count", Value:_count, Compare:"="})    
    }
    _lastdate := c.Geti("lastdate")
    if _lastdate != 0 {
        args = append(args, models.Where{Column:"lastdate", Value:_lastdate, Compare:"="})    
    }
    _duedate := c.Geti("duedate")
    if _duedate != 0 {
        args = append(args, models.Where{Column:"duedate", Value:_duedate, Compare:"="})    
    }
    _remark := c.Get("remark")
    if _remark != "" {
        args = append(args, models.Where{Column:"remark", Value:_remark, Compare:"="})
    }
    _elevator := c.Geti("elevator")
    if _elevator != 0 {
        args = append(args, models.Where{Column:"elevator", Value:_elevator, Compare:"="})    
    }
    _percent := c.Geti("percent")
    if _percent != 0 {
        args = append(args, models.Where{Column:"percent", Value:_percent, Compare:"="})    
    }
    _rate := c.Geti("rate")
    if _rate != 0 {
        args = append(args, models.Where{Column:"rate", Value:_rate, Compare:"="})    
    }
    _type := c.Geti("type")
    if _type != 0 {
        args = append(args, models.Where{Column:"type", Value:_type, Compare:"="})    
    }
    _dong := c.Geti64("dong")
    if _dong != 0 {
        args = append(args, models.Where{Column:"dong", Value:_dong, Compare:"="})    
    }
    _standard := c.Geti64("standard")
    if _standard != 0 {
        args = append(args, models.Where{Column:"standard", Value:_standard, Compare:"="})    
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
                    str += ", b_" + strings.Trim(v, " ")                
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





func (c *BreakdownController) Sum() {
    
    
	conn := c.NewConnection()

	manager := models.NewBreakdownManager(conn)

    var args []interface{}
    
    _topcategory := c.Geti64("topcategory")
    if _topcategory != 0 {
        args = append(args, models.Where{Column:"topcategory", Value:_topcategory, Compare:"="})    
    }
    _subcategory := c.Geti64("subcategory")
    if _subcategory != 0 {
        args = append(args, models.Where{Column:"subcategory", Value:_subcategory, Compare:"="})    
    }
    _category := c.Geti64("category")
    if _category != 0 {
        args = append(args, models.Where{Column:"category", Value:_category, Compare:"="})    
    }
    _method := c.Geti64("method")
    if _method != 0 {
        args = append(args, models.Where{Column:"method", Value:_method, Compare:"="})    
    }
    _count := c.Geti("count")
    if _count != 0 {
        args = append(args, models.Where{Column:"count", Value:_count, Compare:"="})    
    }
    _lastdate := c.Geti("lastdate")
    if _lastdate != 0 {
        args = append(args, models.Where{Column:"lastdate", Value:_lastdate, Compare:"="})    
    }
    _duedate := c.Geti("duedate")
    if _duedate != 0 {
        args = append(args, models.Where{Column:"duedate", Value:_duedate, Compare:"="})    
    }
    _remark := c.Get("remark")
    if _remark != "" {
        args = append(args, models.Where{Column:"remark", Value:_remark, Compare:"like"})
    }
    _elevator := c.Geti("elevator")
    if _elevator != 0 {
        args = append(args, models.Where{Column:"elevator", Value:_elevator, Compare:"="})    
    }
    _percent := c.Geti("percent")
    if _percent != 0 {
        args = append(args, models.Where{Column:"percent", Value:_percent, Compare:"="})    
    }
    _rate := c.Geti("rate")
    if _rate != 0 {
        args = append(args, models.Where{Column:"rate", Value:_rate, Compare:"="})    
    }
    _type := c.Geti("type")
    if _type != 0 {
        args = append(args, models.Where{Column:"type", Value:_type, Compare:"="})    
    }
    _dong := c.Geti64("dong")
    if _dong != 0 {
        args = append(args, models.Where{Column:"dong", Value:_dong, Compare:"="})    
    }
    _standard := c.Geti64("standard")
    if _standard != 0 {
        args = append(args, models.Where{Column:"standard", Value:_standard, Compare:"="})    
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

