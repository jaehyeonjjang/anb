package models

import (
	"anb/config"
	"anb/global"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	logrus "github.com/sirupsen/logrus"
)

func DatabaseSync() {
	lists, err := global.GetPage(config.ServiceUrl + "/api/json/list")

	if err != nil {
		log.Println("url read error")
		return
	}

	if global.WriteFile("webdata/json/list.json", lists) != nil {
		log.Println("write file error")
		return
	}

	datas, err := global.GetPage(config.ServiceUrl + "/api/json/data")

	if err != nil {
		log.Println("url read error")
		return
	}

	SaveDatas(datas)

	images := GetImages()
	if images != nil {
		for _, item := range images {
			if item.Filename != "" {
				logrus.Printf("file download : %v\n", config.ServiceUrl+"/"+item.Filename)
				global.DownloadFile(item.Filename, config.ServiceUrl+"/"+item.Filename)
			}
		}
	}

	statuss := GetStatuss()
	if statuss != nil {
		for _, item := range statuss {
			if item.Type == 8 {
				if item.Content != "" {
					logrus.Printf("file download : %v\n", config.ServiceUrl+"/"+item.Content)
					global.DownloadFile(item.Content, config.ServiceUrl+"/"+item.Content)
				}
			}
		}
	}

	items := GetDatas()
	if items != nil {
		for _, item := range items {
			if item.Filename != "" {
				if item.Filename != "" {
					logrus.Printf("file download : %v\n", config.ServiceUrl+"/"+item.Filename)
					global.DownloadFile(item.Content, config.ServiceUrl+"/"+item.Filename)
				}
			}
		}
	}
}

func DatabaseUploadSync() {
	files, err := ioutil.ReadDir(config.ImagePath)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fmt.Println(file.Name())
	}

	syncs := GetSyncs()
	if syncs == nil || len(syncs) == 0 {
		return
	}

	datas := GetDatas()

	syncArray := make(map[int64]Sync)

	for _, sync := range syncs {
		syncArray[sync.Image] = sync
	}

	for _, sync := range syncArray {
		params := make(map[string]string, 0)

		imageItem := GetImage(sync.Image)

		if imageItem == nil {
			continue
		}

		params["apt"] = fmt.Sprintf("%v", imageItem.Apt)
		params["image"] = fmt.Sprintf("%v", imageItem.Id)
		fn := fmt.Sprintf("webdata/%v-%v.png", imageItem.Apt, imageItem.Id)

		_, err := os.Stat(fn)
		if err == nil {
			request, err := global.NewfileUploadRequest(config.ServiceUrl+"/api/image/upload2", params, "upload", fn)
			if err != nil {
				continue
			}
			client := &http.Client{}
			resp, err := client.Do(request)
			if err != nil {
				log.Println(err)
			} else {
				body, _ := ioutil.ReadAll(resp.Body)
				var ff global.Upload
				json.Unmarshal(body, &ff)
				log.Printf("filename = %v", ff.Filename)

			}
		}
	}

	for i, item := range datas {
		flag := false
		for _, sync := range syncs {
			if item.Image == sync.Image {
				flag = true
				break
			}
		}

		if flag == false {
			continue
		}

		if item.Imagename != "" {
			pos := global.Atoi(item.Imagename)

			if pos > 0 {
				filename := files[pos]
				path := config.ImagePath + "/" + filename.Name()
				log.Printf("path = %v\n", path)
				request, err := global.NewfileUploadRequest(config.ServiceUrl+"/api/image/upload", nil, "upload", path)
				if err != nil {
					continue
				}
				client := &http.Client{}
				resp, err := client.Do(request)
				if err != nil {
					log.Println(err)
				} else {
					body, _ := ioutil.ReadAll(resp.Body)
					var ff global.Upload
					json.Unmarshal(body, &ff)

					log.Printf("filename = %v", ff.Filename)
					datas[i].Filename = string(ff.Filename)
				}
			}
		} else if item.Filename != "" {
			path := item.Filename
			request, err := global.NewfileUploadRequest(config.ServiceUrl+"/api/image/upload", nil, "upload", path)
			if err != nil {
				continue
			}
			client := &http.Client{}
			resp, err := client.Do(request)
			if err != nil {
				log.Println(err)
			} else {
				body, _ := ioutil.ReadAll(resp.Body)
				var ff global.Upload
				json.Unmarshal(body, &ff)

				log.Printf("filename = %v", ff.Filename)
				datas[i].Filename = string(ff.Filename)

			}
		}
	}


	var lists Datas
	lists.Datas = datas
	lists.Syncs = syncs
	content, _ := json.Marshal(lists)

	buff := bytes.NewBuffer(content)
	resp, err := http.Post(config.ServiceUrl+"/api/sync/upload", "application/json", buff)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err == nil {
		str := string(respBody)
		println(str)
	}

	TruncateSync()
}
