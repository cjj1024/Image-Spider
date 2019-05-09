package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"

	"github.com/wonderivan/logger"
)

func init() {
	logger.SetLogger("log.json")
}

type Image struct {
	Name string
	Url  string
}

var existedImage []string

func main() {
	var name string
	var pageNum int
	// fmt.Print("Input Name: ")
	// fmt.Scanf("%s", &name)
	logger.Info("Start Download %s", name)
	name = "id=233580&name=Overlord+Wallpapers"
	pageNum = 10

	GetExistedImage()

	url := "https://wall.alphacoders.com/by_sub_category.php?" + name
	logger.Info("##################Start Download Page 1##################")
	DownloadImage(url)

	for i := 2; i < pageNum; i++ {
		url = "https://wall.alphacoders.com/by_sub_category.php?" + name + "&page=" + strconv.Itoa(i)
		logger.Info("##################Start Download Page %d##################", i)
		DownloadImage(url)
	}
}

func DownloadImage(url string) {
	logger.Info("Start Download Page %s", url)
	images := GetImageUrl(url)

	for i := 0; i < len(images); i++ {
		if IsExisted(images[i].Name) {
			logger.Warn("Image %s Has Existed!", images[i].Name)
			continue
		}
		image := GetImage(images[i].Url)
		SaveImage(image, "D:\\Image Spider\\alphacoders\\"+images[i].Name)
	}
}

func GetImage(url string) []byte {
	resp, err := http.Get(url)
	if err != nil {
		logger.Error("Get Image %s Error!", url)
		logger.Error(err)
		return nil
	}
	if resp.StatusCode != 200 {
		logger.Error("Get Image %s Error!", url)
		logger.Error(err)
		return nil
	}

	image, _ := ioutil.ReadAll(resp.Body)
	logger.Info("Get Image %s Successfully!", url)

	return image
}

func SaveImage(image []byte, name string) {
	out, err := os.Create(name)
	if err != nil {
		logger.Error("Save Image %s Error!", name)
		logger.Error(err)
		return
	}
	defer out.Close()

	io.Copy(out, bytes.NewReader(image))
	logger.Info("Save Image %s Successfully", name)
}

func GetImageUrl(url string) []Image {
	resp, err := http.Get(url)
	if err != nil {
		logger.Error("Get Page %s Error!", url)
		logger.Error(err)
		return nil
	}
	if resp.StatusCode != 200 {
		logger.Error("Get Page %s Error!", url)
		logger.Error(err)
		return nil
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		logger.Error("Get Goquery Document Error!")
		logger.Error(err)
		return nil
	}
	var images []Image
	doc.Find(".boxgrid").Each(func(i int, selection *goquery.Selection) {
		tmp, _ := selection.Find("a").Find("img").Attr("data-src")

		fmt.Println(tmp)
		// https://images2.alphacoders.com/694/694207.png

		// https://images3.alphacoders.com/845/thumb-350-845999.png
		// https://images3.alphacoders.com/845/845999.png
		tmp2 := strings.Split(tmp, "/")
		tmp3 := strings.Split(tmp2[len(tmp2)-1], "-")
		tmp2[len(tmp2)-1] = tmp3[len(tmp3)-1]

		var image Image
		image.Name = tmp3[len(tmp3)-1]
		image.Url = strings.Join(tmp2, "/")
		images = append(images, image)
		logger.Info("Add Image Id %s, Url %s", image.Name, image.Url)
	})

	return images
}

func IsExisted(name string) bool {
	for _, existedName := range existedImage {
		if existedName == name {
			return true
		}
	}

	return false
}

func GetExistedImage() {
	list, err := ioutil.ReadDir("D:\\Image Spider\\alphacoders\\")
	if err != nil {
		logger.Error("Cant Not Read Alphacoders Directory")
		logger.Error(err)
		return
	}

	for _, v := range list {
		existedImage = append(existedImage, v.Name())
	}
}
