package rest


import (
	"repair/controllers"
	"repair/models"

    "strings"
)

type AptlistController struct {
	controllers.Controller
}





func (c *AptlistController) Count() {
    
    
	conn := c.NewConnection()

	manager := models.NewAptlistManager(conn)

    var args []interface{}
    
    _name := c.Get("name")
    if _name != "" {
        args = append(args, models.Where{Column:"name", Value:_name, Compare:"="})
        
    }
    _tel := c.Get("tel")
    if _tel != "" {
        args = append(args, models.Where{Column:"tel", Value:_tel, Compare:"="})
    }
    _fax := c.Get("fax")
    if _fax != "" {
        args = append(args, models.Where{Column:"fax", Value:_fax, Compare:"="})
    }
    _testdate := c.Get("testdate")
    if _testdate != "" {
        args = append(args, models.Where{Column:"testdate", Value:_testdate, Compare:"="})
    }
    _email := c.Get("email")
    if _email != "" {
        args = append(args, models.Where{Column:"email", Value:_email, Compare:"="})
    }
    _personalemail := c.Get("personalemail")
    if _personalemail != "" {
        args = append(args, models.Where{Column:"personalemail", Value:_personalemail, Compare:"="})
    }
    _zip := c.Get("zip")
    if _zip != "" {
        args = append(args, models.Where{Column:"zip", Value:_zip, Compare:"="})
    }
    _address := c.Get("address")
    if _address != "" {
        args = append(args, models.Where{Column:"address", Value:_address, Compare:"="})
    }
    _address2 := c.Get("address2")
    if _address2 != "" {
        args = append(args, models.Where{Column:"address2", Value:_address2, Compare:"="})
    }
    _completeyear := c.Get("completeyear")
    if _completeyear != "" {
        args = append(args, models.Where{Column:"completeyear", Value:_completeyear, Compare:"="})
    }
    _type := c.Get("type")
    if _type != "" {
        args = append(args, models.Where{Column:"type", Value:_type, Compare:"="})
    }
    _flatcount := c.Get("flatcount")
    if _flatcount != "" {
        args = append(args, models.Where{Column:"flatcount", Value:_flatcount, Compare:"="})
    }
    _familycount := c.Get("familycount")
    if _familycount != "" {
        args = append(args, models.Where{Column:"familycount", Value:_familycount, Compare:"="})
    }
    _floor := c.Get("floor")
    if _floor != "" {
        args = append(args, models.Where{Column:"floor", Value:_floor, Compare:"="})
    }
    _fmsloginid := c.Get("fmsloginid")
    if _fmsloginid != "" {
        args = append(args, models.Where{Column:"fmsloginid", Value:_fmsloginid, Compare:"="})
    }
    _fmspasswd := c.Get("fmspasswd")
    if _fmspasswd != "" {
        args = append(args, models.Where{Column:"fmspasswd", Value:_fmspasswd, Compare:"="})
    }
    _repairdate := c.Get("repairdate")
    if _repairdate != "" {
        args = append(args, models.Where{Column:"repairdate", Value:_repairdate, Compare:"="})
    }
    _periodicdate := c.Get("periodicdate")
    if _periodicdate != "" {
        args = append(args, models.Where{Column:"periodicdate", Value:_periodicdate, Compare:"="})
    }
    

    
    
    total := manager.Count(args)
	c.Set("total", total)
}


func (c *AptlistController) Read(id int64) {
    
    
	conn := c.NewConnection()

	manager := models.NewAptlistManager(conn)
	item := manager.Get(id)

    
    
    c.Set("item", item)
}

func (c *AptlistController) Index(page int, pagesize int) {
    
    
	conn := c.NewConnection()

	manager := models.NewAptlistManager(conn)

    var args []interface{}
    
    _name := c.Get("name")
    if _name != "" {
        args = append(args, models.Where{Column:"name", Value:_name, Compare:"="})
        
    }
    _tel := c.Get("tel")
    if _tel != "" {
        args = append(args, models.Where{Column:"tel", Value:_tel, Compare:"="})
    }
    _fax := c.Get("fax")
    if _fax != "" {
        args = append(args, models.Where{Column:"fax", Value:_fax, Compare:"="})
    }
    _testdate := c.Get("testdate")
    if _testdate != "" {
        args = append(args, models.Where{Column:"testdate", Value:_testdate, Compare:"="})
    }
    _email := c.Get("email")
    if _email != "" {
        args = append(args, models.Where{Column:"email", Value:_email, Compare:"="})
    }
    _personalemail := c.Get("personalemail")
    if _personalemail != "" {
        args = append(args, models.Where{Column:"personalemail", Value:_personalemail, Compare:"="})
    }
    _zip := c.Get("zip")
    if _zip != "" {
        args = append(args, models.Where{Column:"zip", Value:_zip, Compare:"="})
    }
    _address := c.Get("address")
    if _address != "" {
        args = append(args, models.Where{Column:"address", Value:_address, Compare:"="})
    }
    _address2 := c.Get("address2")
    if _address2 != "" {
        args = append(args, models.Where{Column:"address2", Value:_address2, Compare:"="})
    }
    _completeyear := c.Get("completeyear")
    if _completeyear != "" {
        args = append(args, models.Where{Column:"completeyear", Value:_completeyear, Compare:"="})
    }
    _type := c.Get("type")
    if _type != "" {
        args = append(args, models.Where{Column:"type", Value:_type, Compare:"="})
    }
    _flatcount := c.Get("flatcount")
    if _flatcount != "" {
        args = append(args, models.Where{Column:"flatcount", Value:_flatcount, Compare:"="})
    }
    _familycount := c.Get("familycount")
    if _familycount != "" {
        args = append(args, models.Where{Column:"familycount", Value:_familycount, Compare:"="})
    }
    _floor := c.Get("floor")
    if _floor != "" {
        args = append(args, models.Where{Column:"floor", Value:_floor, Compare:"="})
    }
    _fmsloginid := c.Get("fmsloginid")
    if _fmsloginid != "" {
        args = append(args, models.Where{Column:"fmsloginid", Value:_fmsloginid, Compare:"="})
    }
    _fmspasswd := c.Get("fmspasswd")
    if _fmspasswd != "" {
        args = append(args, models.Where{Column:"fmspasswd", Value:_fmspasswd, Compare:"="})
    }
    _repairdate := c.Get("repairdate")
    if _repairdate != "" {
        args = append(args, models.Where{Column:"repairdate", Value:_repairdate, Compare:"="})
    }
    _periodicdate := c.Get("periodicdate")
    if _periodicdate != "" {
        args = append(args, models.Where{Column:"periodicdate", Value:_periodicdate, Compare:"="})
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
                    str += ", a_" + strings.Trim(v, " ")                
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





