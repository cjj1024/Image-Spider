package main

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/wonderivan/logger"
)

func init() {
	logger.SetLogger("log.json")
}

func GetImageNameAndUrl(url string) []Image {
	doc, err := goquery.NewDocument(url)
	if err != nil {
		logger.Error("Get Document %s error!", url)
		logger.Error(err)
	}

	var images []Image

	doc.Find(".fbi").Each(func(i int, selection *goquery.Selection) {
		tmp := selection.Find("a")
		url_tmp, _ := tmp.Attr("href")
		url = getImageUrl(url_tmp)

		name, _ := tmp.Find("img").Attr("src")
		name_slice := strings.Split(name, "/")
		name = name_slice[len(name_slice)-1]
		name = strings.Split(name, ".")[0]

		url = url + name + ".jpg"

		images = append(images, Image{name, url})

		logger.Info("Add a image url, name: %s, url: %s", name, url)
	})

	return images
}

// Cookie
//
// cmd	JekxvVbCarH6ZH16F2C8dQnWu4eJaQgn
// Hm_lpvt_633fe378147652d6cb58809821524bec	1552975739
// Hm_lvt_633fe378147652d6cb58809821524bec	1552552570,1552556853,1552891683,1552975723
// uid	255971
// upw	a0663602d3f2ce2f0194a855af395d19

// cmd="Y8ZLdpKLvWmN3F3lf1I0tPge2Uc6WraXQ",
// Hm_lpvt_633fe378147652d6cb58809821524bec = "1552975739",
// Hm_lvt_633fe378147652d6cb58809821524bec="1552552570,1552556853,1552891683,1552975723",
// uid="255971",
// upw="a0663602d3f2ce2f0194a855af395d19"

var cookie string = `cmd=JekxvVbCarH6ZH16F2C8dQnWu4eJaQgn;Hm_lpvt_633fe378147652d6cb58809821524bec=1552975739;Hm_lvt_633fe378147652d6cb58809821524bec=1552552570,1552556853,1552891683,1552975723;uid=255971;upw=a0663602d3f2ce2f0194a855af395d19`

func getImageUrl(url string) string {
	// url = "https://www.roame.net/" + url
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		logger.Error("Set Request %s Error!", url)
		logger.Error(err)
	}
	req.Header.Set("Cookie", cookie)
	req.Header.Add("Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/64.0.3282.140 Safari/537.36 Edge/18.17763")
	resp, err := client.Do(req)
	if err != nil {
		logger.Error("Get Response Error!")
		logger.Error(err)
	}
	html_byte, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	html := string(html_byte)
	// fmt.Println(html)

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		logger.Error("Get Document %s Error!", url)
		logger.Error(err)
		return ""
	}

	img_url, _ := doc.Find("#darlnks").Find("a").Eq(1).Attr("href")

	return img_url
}

func GetImage(url string) []byte {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		logger.Error("Set Request Error!")
		logger.Error(err)
	}
	req.Header.Set("Cookie", cookie)
	req.Header.Add("Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/64.0.3282.140 Safari/537.36 Edge/18.17763")
	resp, err := client.Do(req)
	if err != nil {
		logger.Error("Downlaod Image Error!")
		logger.Error(err)
	}
	defer resp.Body.Close()
	pix, err := ioutil.ReadAll(resp.Body)
	logger.Info("Download Image %s Succussfully!", url)

	return pix
}

func SaveImage(image []byte, name string) {
	name = name + ".jpg"
	out, err := os.Create(name)
	if err != nil {
		logger.Error("Create Image File Error!")
		logger.Error(err)
	}
	defer out.Close()
	io.Copy(out, bytes.NewReader(image))
	logger.Info("Save Image %s Suffessfully!", name)
}
