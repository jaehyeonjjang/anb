package rest


import (
	"repair/controllers"
	"repair/models"

    "strings"
)

type PeriodicincidentalController struct {
	controllers.Controller
}



func (c *PeriodicincidentalController) GetByPeriodic(periodic int64) *models.Periodicincidental {
    
    conn := c.NewConnection()

	_manager := models.NewPeriodicincidentalManager(conn)
    
    item := _manager.GetByPeriodic(periodic)
    
    c.Set("item", item)
    
    
    
    return item
    
}


func (c *PeriodicincidentalController) Insert(item *models.Periodicincidental) {
    
    
	conn := c.NewConnection()
    
	manager := models.NewPeriodicincidentalManager(conn)
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

func (c *PeriodicincidentalController) Insertbatch(item *[]models.Periodicincidental) {  
    if item == nil || len(*item) == 0 {
        return
    }

    rows := len(*item)
    
    
    
	conn := c.NewConnection()
    
	manager := models.NewPeriodicincidentalManager(conn)

    for i := 0; i < rows; i++ {
	    err := manager.Insert(&((*item)[i]))
        if err != nil {
            c.Set("code", "error")    
            c.Set("error", err)
            return
        }
    }
}

func (c *PeriodicincidentalController) Update(item *models.Periodicincidental) {
    
    
	conn := c.NewConnection()

	manager := models.NewPeriodicincidentalManager(conn)
    err := manager.Update(item)
    if err != nil {
        c.Set("code", "error")    
        c.Set("error", err)
        return
    }
}

func (c *PeriodicincidentalController) Delete(item *models.Periodicincidental) {
    
    
    conn := c.NewConnection()

	manager := models.NewPeriodicincidentalManager(conn)

    
	err := manager.Delete(item.Id)
    if err != nil {
        c.Set("code", "error")    
        c.Set("error", err)
    }
}

func (c *PeriodicincidentalController) Deletebatch(item *[]models.Periodicincidental) {
    
    
    conn := c.NewConnection()

	manager := models.NewPeriodicincidentalManager(conn)

    for _, v := range *item {
        
    
	    err := manager.Delete(v.Id)
        if err != nil {
            c.Set("code", "error")    
            c.Set("error", err)
            return
        }
    }
}

func (c *PeriodicincidentalController) Count() {
    
    
	conn := c.NewConnection()

	manager := models.NewPeriodicincidentalManager(conn)

    var args []interface{}
    
    _result1 := c.Get("result1")
    if _result1 != "" {
        args = append(args, models.Where{Column:"result1", Value:_result1, Compare:"="})
    }
    _result2 := c.Get("result2")
    if _result2 != "" {
        args = append(args, models.Where{Column:"result2", Value:_result2, Compare:"="})
    }
    _result3 := c.Get("result3")
    if _result3 != "" {
        args = append(args, models.Where{Column:"result3", Value:_result3, Compare:"="})
    }
    _result4 := c.Get("result4")
    if _result4 != "" {
        args = append(args, models.Where{Column:"result4", Value:_result4, Compare:"="})
    }
    _result5 := c.Get("result5")
    if _result5 != "" {
        args = append(args, models.Where{Column:"result5", Value:_result5, Compare:"="})
    }
    _result6 := c.Get("result6")
    if _result6 != "" {
        args = append(args, models.Where{Column:"result6", Value:_result6, Compare:"="})
    }
    _result7 := c.Get("result7")
    if _result7 != "" {
        args = append(args, models.Where{Column:"result7", Value:_result7, Compare:"="})
    }
    _result8 := c.Get("result8")
    if _result8 != "" {
        args = append(args, models.Where{Column:"result8", Value:_result8, Compare:"="})
    }
    _result9 := c.Get("result9")
    if _result9 != "" {
        args = append(args, models.Where{Column:"result9", Value:_result9, Compare:"="})
    }
    _result10 := c.Get("result10")
    if _result10 != "" {
        args = append(args, models.Where{Column:"result10", Value:_result10, Compare:"="})
    }
    _result11 := c.Get("result11")
    if _result11 != "" {
        args = append(args, models.Where{Column:"result11", Value:_result11, Compare:"="})
    }
    _result12 := c.Get("result12")
    if _result12 != "" {
        args = append(args, models.Where{Column:"result12", Value:_result12, Compare:"="})
    }
    _result13 := c.Get("result13")
    if _result13 != "" {
        args = append(args, models.Where{Column:"result13", Value:_result13, Compare:"="})
    }
    _result14 := c.Get("result14")
    if _result14 != "" {
        args = append(args, models.Where{Column:"result14", Value:_result14, Compare:"="})
    }
    _result15 := c.Get("result15")
    if _result15 != "" {
        args = append(args, models.Where{Column:"result15", Value:_result15, Compare:"="})
    }
    _result16 := c.Get("result16")
    if _result16 != "" {
        args = append(args, models.Where{Column:"result16", Value:_result16, Compare:"="})
    }
    _result17 := c.Get("result17")
    if _result17 != "" {
        args = append(args, models.Where{Column:"result17", Value:_result17, Compare:"="})
    }
    _result18 := c.Get("result18")
    if _result18 != "" {
        args = append(args, models.Where{Column:"result18", Value:_result18, Compare:"="})
    }
    _result19 := c.Get("result19")
    if _result19 != "" {
        args = append(args, models.Where{Column:"result19", Value:_result19, Compare:"="})
    }
    _result20 := c.Get("result20")
    if _result20 != "" {
        args = append(args, models.Where{Column:"result20", Value:_result20, Compare:"="})
    }
    _result21 := c.Get("result21")
    if _result21 != "" {
        args = append(args, models.Where{Column:"result21", Value:_result21, Compare:"="})
    }
    _status1 := c.Get("status1")
    if _status1 != "" {
        args = append(args, models.Where{Column:"status1", Value:_status1, Compare:"="})
    }
    _status2 := c.Get("status2")
    if _status2 != "" {
        args = append(args, models.Where{Column:"status2", Value:_status2, Compare:"="})
    }
    _status3 := c.Get("status3")
    if _status3 != "" {
        args = append(args, models.Where{Column:"status3", Value:_status3, Compare:"="})
    }
    _status4 := c.Get("status4")
    if _status4 != "" {
        args = append(args, models.Where{Column:"status4", Value:_status4, Compare:"="})
    }
    _status5 := c.Get("status5")
    if _status5 != "" {
        args = append(args, models.Where{Column:"status5", Value:_status5, Compare:"="})
    }
    _status6 := c.Get("status6")
    if _status6 != "" {
        args = append(args, models.Where{Column:"status6", Value:_status6, Compare:"="})
    }
    _status7 := c.Get("status7")
    if _status7 != "" {
        args = append(args, models.Where{Column:"status7", Value:_status7, Compare:"="})
    }
    _status8 := c.Get("status8")
    if _status8 != "" {
        args = append(args, models.Where{Column:"status8", Value:_status8, Compare:"="})
    }
    _status9 := c.Get("status9")
    if _status9 != "" {
        args = append(args, models.Where{Column:"status9", Value:_status9, Compare:"="})
    }
    _status10 := c.Get("status10")
    if _status10 != "" {
        args = append(args, models.Where{Column:"status10", Value:_status10, Compare:"="})
    }
    _status11 := c.Get("status11")
    if _status11 != "" {
        args = append(args, models.Where{Column:"status11", Value:_status11, Compare:"="})
    }
    _status12 := c.Get("status12")
    if _status12 != "" {
        args = append(args, models.Where{Column:"status12", Value:_status12, Compare:"="})
    }
    _status13 := c.Get("status13")
    if _status13 != "" {
        args = append(args, models.Where{Column:"status13", Value:_status13, Compare:"="})
    }
    _status14 := c.Get("status14")
    if _status14 != "" {
        args = append(args, models.Where{Column:"status14", Value:_status14, Compare:"="})
    }
    _status15 := c.Get("status15")
    if _status15 != "" {
        args = append(args, models.Where{Column:"status15", Value:_status15, Compare:"="})
    }
    _status16 := c.Get("status16")
    if _status16 != "" {
        args = append(args, models.Where{Column:"status16", Value:_status16, Compare:"="})
    }
    _status17 := c.Get("status17")
    if _status17 != "" {
        args = append(args, models.Where{Column:"status17", Value:_status17, Compare:"="})
    }
    _status18 := c.Get("status18")
    if _status18 != "" {
        args = append(args, models.Where{Column:"status18", Value:_status18, Compare:"="})
    }
    _status19 := c.Get("status19")
    if _status19 != "" {
        args = append(args, models.Where{Column:"status19", Value:_status19, Compare:"="})
    }
    _status20 := c.Get("status20")
    if _status20 != "" {
        args = append(args, models.Where{Column:"status20", Value:_status20, Compare:"="})
    }
    _status21 := c.Get("status21")
    if _status21 != "" {
        args = append(args, models.Where{Column:"status21", Value:_status21, Compare:"="})
    }
    _position1 := c.Get("position1")
    if _position1 != "" {
        args = append(args, models.Where{Column:"position1", Value:_position1, Compare:"="})
    }
    _position2 := c.Get("position2")
    if _position2 != "" {
        args = append(args, models.Where{Column:"position2", Value:_position2, Compare:"="})
    }
    _position3 := c.Get("position3")
    if _position3 != "" {
        args = append(args, models.Where{Column:"position3", Value:_position3, Compare:"="})
    }
    _position4 := c.Get("position4")
    if _position4 != "" {
        args = append(args, models.Where{Column:"position4", Value:_position4, Compare:"="})
    }
    _position5 := c.Get("position5")
    if _position5 != "" {
        args = append(args, models.Where{Column:"position5", Value:_position5, Compare:"="})
    }
    _position6 := c.Get("position6")
    if _position6 != "" {
        args = append(args, models.Where{Column:"position6", Value:_position6, Compare:"="})
    }
    _position7 := c.Get("position7")
    if _position7 != "" {
        args = append(args, models.Where{Column:"position7", Value:_position7, Compare:"="})
    }
    _position8 := c.Get("position8")
    if _position8 != "" {
        args = append(args, models.Where{Column:"position8", Value:_position8, Compare:"="})
    }
    _position9 := c.Get("position9")
    if _position9 != "" {
        args = append(args, models.Where{Column:"position9", Value:_position9, Compare:"="})
    }
    _position10 := c.Get("position10")
    if _position10 != "" {
        args = append(args, models.Where{Column:"position10", Value:_position10, Compare:"="})
    }
    _position11 := c.Get("position11")
    if _position11 != "" {
        args = append(args, models.Where{Column:"position11", Value:_position11, Compare:"="})
    }
    _position12 := c.Get("position12")
    if _position12 != "" {
        args = append(args, models.Where{Column:"position12", Value:_position12, Compare:"="})
    }
    _position13 := c.Get("position13")
    if _position13 != "" {
        args = append(args, models.Where{Column:"position13", Value:_position13, Compare:"="})
    }
    _position14 := c.Get("position14")
    if _position14 != "" {
        args = append(args, models.Where{Column:"position14", Value:_position14, Compare:"="})
    }
    _position15 := c.Get("position15")
    if _position15 != "" {
        args = append(args, models.Where{Column:"position15", Value:_position15, Compare:"="})
    }
    _position16 := c.Get("position16")
    if _position16 != "" {
        args = append(args, models.Where{Column:"position16", Value:_position16, Compare:"="})
    }
    _position17 := c.Get("position17")
    if _position17 != "" {
        args = append(args, models.Where{Column:"position17", Value:_position17, Compare:"="})
    }
    _position18 := c.Get("position18")
    if _position18 != "" {
        args = append(args, models.Where{Column:"position18", Value:_position18, Compare:"="})
    }
    _position19 := c.Get("position19")
    if _position19 != "" {
        args = append(args, models.Where{Column:"position19", Value:_position19, Compare:"="})
    }
    _position20 := c.Get("position20")
    if _position20 != "" {
        args = append(args, models.Where{Column:"position20", Value:_position20, Compare:"="})
    }
    _position21 := c.Get("position21")
    if _position21 != "" {
        args = append(args, models.Where{Column:"position21", Value:_position21, Compare:"="})
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


func (c *PeriodicincidentalController) Read(id int64) {
    
    
	conn := c.NewConnection()

	manager := models.NewPeriodicincidentalManager(conn)
	item := manager.Get(id)

    
    
    c.Set("item", item)
}

func (c *PeriodicincidentalController) Index(page int, pagesize int) {
    
    
	conn := c.NewConnection()

	manager := models.NewPeriodicincidentalManager(conn)

    var args []interface{}
    
    _result1 := c.Get("result1")
    if _result1 != "" {
        args = append(args, models.Where{Column:"result1", Value:_result1, Compare:"="})
    }
    _result2 := c.Get("result2")
    if _result2 != "" {
        args = append(args, models.Where{Column:"result2", Value:_result2, Compare:"="})
    }
    _result3 := c.Get("result3")
    if _result3 != "" {
        args = append(args, models.Where{Column:"result3", Value:_result3, Compare:"="})
    }
    _result4 := c.Get("result4")
    if _result4 != "" {
        args = append(args, models.Where{Column:"result4", Value:_result4, Compare:"="})
    }
    _result5 := c.Get("result5")
    if _result5 != "" {
        args = append(args, models.Where{Column:"result5", Value:_result5, Compare:"="})
    }
    _result6 := c.Get("result6")
    if _result6 != "" {
        args = append(args, models.Where{Column:"result6", Value:_result6, Compare:"="})
    }
    _result7 := c.Get("result7")
    if _result7 != "" {
        args = append(args, models.Where{Column:"result7", Value:_result7, Compare:"="})
    }
    _result8 := c.Get("result8")
    if _result8 != "" {
        args = append(args, models.Where{Column:"result8", Value:_result8, Compare:"="})
    }
    _result9 := c.Get("result9")
    if _result9 != "" {
        args = append(args, models.Where{Column:"result9", Value:_result9, Compare:"="})
    }
    _result10 := c.Get("result10")
    if _result10 != "" {
        args = append(args, models.Where{Column:"result10", Value:_result10, Compare:"="})
    }
    _result11 := c.Get("result11")
    if _result11 != "" {
        args = append(args, models.Where{Column:"result11", Value:_result11, Compare:"="})
    }
    _result12 := c.Get("result12")
    if _result12 != "" {
        args = append(args, models.Where{Column:"result12", Value:_result12, Compare:"="})
    }
    _result13 := c.Get("result13")
    if _result13 != "" {
        args = append(args, models.Where{Column:"result13", Value:_result13, Compare:"="})
    }
    _result14 := c.Get("result14")
    if _result14 != "" {
        args = append(args, models.Where{Column:"result14", Value:_result14, Compare:"="})
    }
    _result15 := c.Get("result15")
    if _result15 != "" {
        args = append(args, models.Where{Column:"result15", Value:_result15, Compare:"="})
    }
    _result16 := c.Get("result16")
    if _result16 != "" {
        args = append(args, models.Where{Column:"result16", Value:_result16, Compare:"="})
    }
    _result17 := c.Get("result17")
    if _result17 != "" {
        args = append(args, models.Where{Column:"result17", Value:_result17, Compare:"="})
    }
    _result18 := c.Get("result18")
    if _result18 != "" {
        args = append(args, models.Where{Column:"result18", Value:_result18, Compare:"="})
    }
    _result19 := c.Get("result19")
    if _result19 != "" {
        args = append(args, models.Where{Column:"result19", Value:_result19, Compare:"="})
    }
    _result20 := c.Get("result20")
    if _result20 != "" {
        args = append(args, models.Where{Column:"result20", Value:_result20, Compare:"="})
    }
    _result21 := c.Get("result21")
    if _result21 != "" {
        args = append(args, models.Where{Column:"result21", Value:_result21, Compare:"="})
    }
    _status1 := c.Get("status1")
    if _status1 != "" {
        args = append(args, models.Where{Column:"status1", Value:_status1, Compare:"="})
    }
    _status2 := c.Get("status2")
    if _status2 != "" {
        args = append(args, models.Where{Column:"status2", Value:_status2, Compare:"="})
    }
    _status3 := c.Get("status3")
    if _status3 != "" {
        args = append(args, models.Where{Column:"status3", Value:_status3, Compare:"="})
    }
    _status4 := c.Get("status4")
    if _status4 != "" {
        args = append(args, models.Where{Column:"status4", Value:_status4, Compare:"="})
    }
    _status5 := c.Get("status5")
    if _status5 != "" {
        args = append(args, models.Where{Column:"status5", Value:_status5, Compare:"="})
    }
    _status6 := c.Get("status6")
    if _status6 != "" {
        args = append(args, models.Where{Column:"status6", Value:_status6, Compare:"="})
    }
    _status7 := c.Get("status7")
    if _status7 != "" {
        args = append(args, models.Where{Column:"status7", Value:_status7, Compare:"="})
    }
    _status8 := c.Get("status8")
    if _status8 != "" {
        args = append(args, models.Where{Column:"status8", Value:_status8, Compare:"="})
    }
    _status9 := c.Get("status9")
    if _status9 != "" {
        args = append(args, models.Where{Column:"status9", Value:_status9, Compare:"="})
    }
    _status10 := c.Get("status10")
    if _status10 != "" {
        args = append(args, models.Where{Column:"status10", Value:_status10, Compare:"="})
    }
    _status11 := c.Get("status11")
    if _status11 != "" {
        args = append(args, models.Where{Column:"status11", Value:_status11, Compare:"="})
    }
    _status12 := c.Get("status12")
    if _status12 != "" {
        args = append(args, models.Where{Column:"status12", Value:_status12, Compare:"="})
    }
    _status13 := c.Get("status13")
    if _status13 != "" {
        args = append(args, models.Where{Column:"status13", Value:_status13, Compare:"="})
    }
    _status14 := c.Get("status14")
    if _status14 != "" {
        args = append(args, models.Where{Column:"status14", Value:_status14, Compare:"="})
    }
    _status15 := c.Get("status15")
    if _status15 != "" {
        args = append(args, models.Where{Column:"status15", Value:_status15, Compare:"="})
    }
    _status16 := c.Get("status16")
    if _status16 != "" {
        args = append(args, models.Where{Column:"status16", Value:_status16, Compare:"="})
    }
    _status17 := c.Get("status17")
    if _status17 != "" {
        args = append(args, models.Where{Column:"status17", Value:_status17, Compare:"="})
    }
    _status18 := c.Get("status18")
    if _status18 != "" {
        args = append(args, models.Where{Column:"status18", Value:_status18, Compare:"="})
    }
    _status19 := c.Get("status19")
    if _status19 != "" {
        args = append(args, models.Where{Column:"status19", Value:_status19, Compare:"="})
    }
    _status20 := c.Get("status20")
    if _status20 != "" {
        args = append(args, models.Where{Column:"status20", Value:_status20, Compare:"="})
    }
    _status21 := c.Get("status21")
    if _status21 != "" {
        args = append(args, models.Where{Column:"status21", Value:_status21, Compare:"="})
    }
    _position1 := c.Get("position1")
    if _position1 != "" {
        args = append(args, models.Where{Column:"position1", Value:_position1, Compare:"="})
    }
    _position2 := c.Get("position2")
    if _position2 != "" {
        args = append(args, models.Where{Column:"position2", Value:_position2, Compare:"="})
    }
    _position3 := c.Get("position3")
    if _position3 != "" {
        args = append(args, models.Where{Column:"position3", Value:_position3, Compare:"="})
    }
    _position4 := c.Get("position4")
    if _position4 != "" {
        args = append(args, models.Where{Column:"position4", Value:_position4, Compare:"="})
    }
    _position5 := c.Get("position5")
    if _position5 != "" {
        args = append(args, models.Where{Column:"position5", Value:_position5, Compare:"="})
    }
    _position6 := c.Get("position6")
    if _position6 != "" {
        args = append(args, models.Where{Column:"position6", Value:_position6, Compare:"="})
    }
    _position7 := c.Get("position7")
    if _position7 != "" {
        args = append(args, models.Where{Column:"position7", Value:_position7, Compare:"="})
    }
    _position8 := c.Get("position8")
    if _position8 != "" {
        args = append(args, models.Where{Column:"position8", Value:_position8, Compare:"="})
    }
    _position9 := c.Get("position9")
    if _position9 != "" {
        args = append(args, models.Where{Column:"position9", Value:_position9, Compare:"="})
    }
    _position10 := c.Get("position10")
    if _position10 != "" {
        args = append(args, models.Where{Column:"position10", Value:_position10, Compare:"="})
    }
    _position11 := c.Get("position11")
    if _position11 != "" {
        args = append(args, models.Where{Column:"position11", Value:_position11, Compare:"="})
    }
    _position12 := c.Get("position12")
    if _position12 != "" {
        args = append(args, models.Where{Column:"position12", Value:_position12, Compare:"="})
    }
    _position13 := c.Get("position13")
    if _position13 != "" {
        args = append(args, models.Where{Column:"position13", Value:_position13, Compare:"="})
    }
    _position14 := c.Get("position14")
    if _position14 != "" {
        args = append(args, models.Where{Column:"position14", Value:_position14, Compare:"="})
    }
    _position15 := c.Get("position15")
    if _position15 != "" {
        args = append(args, models.Where{Column:"position15", Value:_position15, Compare:"="})
    }
    _position16 := c.Get("position16")
    if _position16 != "" {
        args = append(args, models.Where{Column:"position16", Value:_position16, Compare:"="})
    }
    _position17 := c.Get("position17")
    if _position17 != "" {
        args = append(args, models.Where{Column:"position17", Value:_position17, Compare:"="})
    }
    _position18 := c.Get("position18")
    if _position18 != "" {
        args = append(args, models.Where{Column:"position18", Value:_position18, Compare:"="})
    }
    _position19 := c.Get("position19")
    if _position19 != "" {
        args = append(args, models.Where{Column:"position19", Value:_position19, Compare:"="})
    }
    _position20 := c.Get("position20")
    if _position20 != "" {
        args = append(args, models.Where{Column:"position20", Value:_position20, Compare:"="})
    }
    _position21 := c.Get("position21")
    if _position21 != "" {
        args = append(args, models.Where{Column:"position21", Value:_position21, Compare:"="})
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
                    str += ", pi_" + strings.Trim(v, " ")                
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





