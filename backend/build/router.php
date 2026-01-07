#!/opt/homebrew/bin/php
<?php

$argc = count($argv) - 1;

if ($argc == 0) {
    $cwd = getcwd();
} else {
    $cwd = $argv[1];
}

$dir = $cwd . '/../controllers';

$config_file = $cwd . '/' . 'router.config.inc';
if (!is_file($config_file)) {
    echo "config file not found\n";
    exit;
}

require_once $config_file;


function getModel($name) {
    global $models;

    if (array_key_exists(strtolower($name), $models)) {
    return ucfirst($models[strtolower($name)]);
    }

    return $name;
}
// 핸들 획득
$handle  = opendir($dir);

$files = array();

while (false !== ($pathname = readdir($handle))) {
    if($pathname == "." || $pathname == ".."){
        continue;
    }

    if(is_dir($dir . "/" . $pathname)) {
    $inner_handle  = opendir($dir . "/" . $pathname);

    $files[$pathname] = array();

    while (false !== ($filename2 = readdir($inner_handle))) {
        if($filename2 == "." || $filename2 == "..") {
            continue;
        }

        $full_filename = $dir . "/" . $pathname . "/" . $filename2;
        if(is_file($full_filename)) {
            $files[$pathname][] = $filename2;
        }
    }

    closedir($inner_handle);
    }
}

closedir($handle);

$str = '';

$str .= 'package router

import (
    "net/http"
    _ "strconv"
';
foreach ($files as $key => $value) {
    if ($key == 'mobile')
        continue;

    $str .= '    "anb/controllers/' . $key . '"'."\n";
}

$str .= '
    _ "anb/models"
    "github.com/gin-gonic/gin"
)

func SetRouter(r *gin.Engine) {';



function process($dir, $key, $item, $newname, $newkey, $category) {
    $str = '';

    $data = file($dir . '/' . $key . '/' . $item);

    $pathname = $key;
    $name = str_replace('.go', '', $item);
    $name2 = ucfirst($name);

    $funcs = array();
    for ($j = 0; $j < count($data); $j++) {
        if (substr($data[$j], 0, 6) == 'func (') {
            preg_match('/([a-zA-z0-9]+)\(/', $data[$j], $matches);
            $funcs[] = $matches[1];
        }
    }

    if ($newname != '') {
        $name = $newname;
        $name2 = ucfirst($name);
        $key = $newkey;
    }

    for ($j = 0; $j < count($funcs); $j++) {
        $func = $funcs[$j];
        $func2 = strtolower($func);

        if ($func == 'Index') {
                $str .= "
        ${key}Group.GET(\"/${name}\", func(c *gin.Context) {
            var controller ${key}.${name2}Controller
";
                if ($category > 0) {
                    $str .= "            controller.Category = " . $category . "\n";
                }
                $str .= "
            controller.Init(c)
            controller.Set(\"current\", \"/${key}/${name}\")
            controller.Current = \"/${key}/${name}\"
            controller.Index()
";
                if ($category > 0) {
                    $str .= "            controller.Set(\"category\", " . $category . ")\n";
                }
                $str .= "

            controller.Set(\"self\", \"${name}\")
            controller.Display(\"/${key}/${name}.html\")
            controller.Close()
        })\n";
            } else if ($func == 'View') {
                $str .= "
        ${key}Group.GET(\"/${name}/view/:id\", func(c *gin.Context) {
            id, _ := strconv.ParseInt(c.Param(\"id\"), 10, 64)

            var controller ${key}.${name2}Controller
";
                if ($category > 0) {
                    $str .= "            controller.Category = " . $category;
                }
                $str .= "
            controller.Init(c)
            conn := controller.NewConnection()

            manager := models.New".@getModel(${name2})."Manager(conn)
            item := manager.Get(id)
            controller.Set(\"current\", \"/${key}/${name}\")
            controller.Current = \"/${key}/${name}\"
            item = controller.View(item)
";
                if ($category > 0) {
                    $str .= "            controller.Set(\"category\", " . $category . ")\n";
                }
                $str .= "
            controller.Set(\"item\", item)

            controller.Set(\"self\", \"${name}\")
            controller.Display(\"/${key}/${name}_view.html\")
            controller.Close()
        })\n";
            } else if ($func == 'Insert') {
                $str .= "
        ${key}Group.GET(\"/${name}/insert\", func(c *gin.Context) {
            var controller ${key}.${name2}Controller
";
                if ($category > 0) {
                    $str .= "            controller.Category = " . $category;
                }
                $str .= "
            controller.Init(c)
            controller.Set(\"current\", \"/${key}/${name}\")
            controller.Current = \"/${key}/${name}\"
            var item models.".@getModel(${name2})."
            controller.Insert(&item)
";
                if ($category > 0) {
                    $str .= "            controller.Set(\"category\", " . $category . ")\n";
                }
                $str .= "
            controller.Set(\"item\", item)
            controller.Set(\"mode\", \"insert\")
            controller.Set(\"action\", \"/${key}/${name}/insert_process\")

            controller.Set(\"self\", \"${name}\")
            controller.Display(\"/${key}/${name}_insert.html\")
            controller.Close()
        })\n";
            } else if ($func == 'InsertProcess') {
                $str .= "
        ${key}Group.POST(\"/${name}/insert_process\", func(c *gin.Context) {
            var controller ${key}.${name2}Controller
";
                if ($category > 0) {
                    $str .= "            controller.Category = " . $category;
                }
                $str .= "
            controller.Init(c)

            var item models.".@getModel(${name2})."
            c.ShouldBind(&item)
            controller.InsertProcess(&item)

            controller.Close()
            c.Redirect(http.StatusMovedPermanently, \"/${key}/${name}\")
        })\n";
            } else if ($func == 'Update') {
                $str .= "
        ${key}Group.GET(\"/${name}/update/:id\", func(c *gin.Context) {
            id, _ := strconv.ParseInt(c.Param(\"id\"), 10, 64)

            var controller ${key}.${name2}Controller
";
                if ($category > 0) {
                    $str .= "            controller.Category = " . $category;
                }
                $str .= "
            controller.Init(c)
            controller.Set(\"current\", \"/${key}/${name}\")
            controller.Current = \"/${key}/${name}\"

            conn := controller.NewConnection()

            manager := models.New".@getModel(${name2})."Manager(conn)
            item := manager.Get(id)
            controller.Update(item)
";
                if ($category > 0) {
                    $str .= "            controller.Set(\"category\", " . $category . ")\n";
                }
                $str .= "
            controller.Set(\"item\", item)
            controller.Set(\"mode\", \"update\")
            controller.Set(\"action\", \"/${key}/${name}/update_process\")

            controller.Set(\"self\", \"${name}\")
            controller.Display(\"/${key}/${name}_insert.html\")

            controller.Close()
        })\n";
            } else if ($func == 'UpdateProcess') {
                $str .= "
        ${key}Group.POST(\"/${name}/update_process\", func(c *gin.Context) {
            id, _ := strconv.ParseInt(c.PostForm(\"id\"), 10, 64)

            var controller ${key}.${name2}Controller
";
                if ($category > 0) {
                    $str .= "            controller.Category = " . $category;
                }
                $str .= "
            controller.Init(c)
            conn := controller.NewConnection()

            manager := models.New".@getModel(${name2})."Manager(conn)
            item := manager.Get(id)

            c.ShouldBind(item)
            controller.UpdateProcess(item)

            controller.Close()
            c.Redirect(http.StatusMovedPermanently, \"/${key}/${name}\")
        })\n";
            } else if ($func == 'Delete') {
                $str .= "
        ${key}Group.GET(\"/${name}/delete/:id\", func(c *gin.Context) {
            id, _ := strconv.ParseInt(c.Param(\"id\"), 10, 64)

            var controller ${key}.${name2}Controller
";
                if ($category > 0) {
                    $str .= "            controller.Category = " . $category;
                }
                $str .= "
            controller.Init(c)

            controller.Delete(id)

            controller.Close()
            c.Redirect(http.StatusMovedPermanently, \"/${key}/${name}\")
        })\n";
            } else if (substr($func, 0, 4) == 'Ajax') {
                $func3 = substr($func2, 4);

                $full = '/';
                $full .= $name;
                if ($key != 'api') {
                    $full .= '/ajax';
                }

                if ($func != 'AjaxIndex') {
                    $full .= '/'.$func3;
                }
                $str .= "
        ${key}Group.Any(\"${full}\", func(c *gin.Context) {
            var controller ${key}.${name2}Controller
";
                if ($category > 0) {
                    $str .= "            controller.Category = " . $category;
                }
                $str .= "
            controller.Init(c)

            controller.${func}()
            controller.Close()
            c.JSON(http.StatusOK, controller.Result)
        })\n";
            } else if (substr($func, 0, 8) == 'Download') {
                $func3 = substr($func2, 8);

                $full = '/' . $name . '/download/' . $func3;
                $str .= "
        ${key}Group.GET(\"${full}\", func(c *gin.Context) {
            var controller ${key}.${name2}Controller
            controller.Init(c)

            controller.${func}()
";
                if ($category > 0) {
                    $str .= "            controller.Set(\"category\", " . $category . ")\n";
                }
                $str .= "
            controller.Set(\"self\", \"${name}\")
            controller.Close()
        })\n";
            } else if ($func == 'PageView') {
                $func3 = substr($func2, 4);
                $str .= "
        ${key}Group.GET(\"/${name}/page/${func3}/:id\", func(c *gin.Context) {
            id, _ := strconv.Atoi(c.Param(\"id\"))

            var controller ${key}.${name2}Controller
";
                if ($category > 0) {
                    $str .= "            controller.Category = " . $category;
                }
                $str .= "
            controller.Init(c)

            controller.${func}(id)
            controller.Display(\"/${key}/${name}_page_${func3}.html\")
            controller.Close()
        })\n";
            } else if (substr($func, 0, 4) == 'Page') {
                $func3 = substr($func2, 4);
                $str .= "
        ${key}Group.GET(\"/${name}/page/${func3}\", func(c *gin.Context) {
            var controller ${key}.${name2}Controller
";
                if ($category > 0) {
                    $str .= "            controller.Category = " . $category;
                }
                $str .= "
            controller.Init(c)

            controller.${func}()
            controller.Display(\"/${key}/${name}_page_${func3}.html\")
            controller.Close()
        })\n";
            } else if (substr($func, 0, 5) == 'Popup') {
                $func3 = substr($func2, 5);
                $str .= "
        ${key}Group.GET(\"/${name}/popup/${func3}\", func(c *gin.Context) {
            var controller ${key}.${name2}Controller
";
                if ($category > 0) {
                    $str .= "            controller.Category = " . $category;
                }
                $str .= "
            controller.Init(c)

            controller.${func}()
";
                if ($category > 0) {
                    $str .= "            controller.Set(\"category\", " . $category . ")\n";
                }
                $str .= "
            controller.Set(\"self\", \"${name}\")
            controller.Display(\"/${key}/${name}_popup_${func3}.html\")
            controller.Close()
        })\n";
            } else {
                if (strpos($func, "Process") == false) {
                    $m = 'GET';
                } else {
                    $m = 'POST';
                }

                $str .= "
        ${key}Group.${m}(\"/${name}/${func2}\", func(c *gin.Context) {
            var controller ${key}.${name2}Controller
";
                if ($category > 0) {
                    $str .= "            controller.Category = " . $category;
                }
                $str .= "
            controller.Init(c)
            controller.Set(\"current\", \"/${key}/${name}\")
            controller.${func}()
";
                if ($category > 0) {
                    $str .= "            controller.Set(\"category\", " . $category . ")\n";
                }
                $str .= "
            controller.Set(\"self\", \"${name}\")";

                if (strpos($func, "Process") == false) {
                    $str .= "
            controller.Display(\"/${key}/${name}_${func2}.html\")";
                }
                $str .= "
            controller.Close()
        })\n";
            }
        }

        return $str;
}



foreach ($files as $key => $value) {
    $rows = count($value);

    $str .= "
    ${key}Group := r.Group(\"/$key\")";


    if ($key != 'login') {
        if ($key == 'api') {
            $str .= "
    //${key}Group.Use(TokenRequired())
";
        } else if ($key =='admin') {
            $str .= "
    ${key}Group.Use(AuthAdminRequired(false))
";
        } else if ($key == 'mypage') {
            $str .= "
    ${key}Group.Use(AuthRequired(false))
";
        }
    }

    $str .="
    {";
    for ($i = 0; $i < $rows; $i++) {
        $item = $value[$i];
        $name = str_replace('.go', '', $item);
        $keyname = $key.'/'.$name;
        $ret = process($dir, $key, $item, '', '', 0);

        $str .= $ret;

        if (array_key_exists($keyname, $subclass)) {
            $str .= process($dir, $subclass[$keyname][0], $subclass[$keyname][1].".go", $name, $key, $subclass[$keyname][2]);
        }
    }

    $str .= "    }
";

}

$str .= '}
';

file_put_contents($cwd . '/../router/router.go', $str);

?>
