package main

import (
	"fmt"
	"io/ioutil"
	"path"
	"strconv"
	"strings"

	"github.com/wonderivan/logger"
)

// logger.SetLogger(`{"Console": {"level": "DEBG"}`)

func init() {
	logger.SetLogger("log.json")
}

type Image struct {
	Name string
	Url  string
}

var img_list []string

func main() {
	var name string
	fmt.Print("Input Name: ")
	fmt.Scanf("%s", &name)

	GetRoameImageName()

    logger.Info("############################### Page 1 Start ###############################")
	url := "https://www.roame.net/" + name + "/index.html"
	logger.Info("Start download page %s", url)
	DownloadImage(url)
    logger.Info("############################### Page 1 End ###############################")

	for i := 2; i < 19; i++ {
        logger.Info("############################### Page %d Start ###############################", i)
		url = "https://www.roame.net/" + name + "/index_" + strconv.Itoa(i) + ".html"
		logger.Info("Start Downlaod Page %s", url)
		DownloadImage(url)
        logger.Info("############################### Page %d Start ###############################", i)
	}
}

func DownloadImage(url string) {
	images := GetImageNameAndUrl(url)

	for i := 0; i < len(images); i++ {
		if HasImage(images[i].Name) {
			logger.Warn("Image %s Has Downloaded!", images[i].Name)
			continue
		}
		image := GetImage(images[i].Url)
		SaveImage(image, "D:\\roame\\Image Spider\\"+images[i].Name)
	}
}

func GetRoameImageName() {
	list, err := ioutil.ReadDir("D:\\roame\\Image Spider\\")
	if err != nil {
		logger.Error("Can Not Read Roame Directory!")
		logger.Error(err)
		return
	}

	for _, v := range list {
		img_list = append(img_list, strings.TrimSuffix(v.Name(), path.Ext(v.Name())))
	}
}

func HasImage(name string) bool {
	for _, exname := range img_list {
		if name == exname {
			return true
		}
	}

	return false
}
