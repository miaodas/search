package hiendy

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"testing"
)

func TestStart(t *testing.T) {
	respon := Start()
	if respon > 0 {
		t.Log("failed !")
	} else {
		t.Log("success !")
	}
}

func TestHotTu(t *testing.T) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://huotu42.com/category/short_video/", strings.NewReader(""))
	if err != nil {
		log.Printf("NewRequest err: %+v", err)
	}

	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Linux, U, Android 3.0, en-us, Xoom Build/HRI39) AppleWebKit/534.13 (KHTML, like Gecko) Version/4.0 Safari/534.13")
	// req.Header.Set("Host", "archiver.hiendy.com")
	response, err := client.Do(req)
	if err != nil {
		log.Printf("NewRequest err: %+v", err)
	}
	s, err := ioutil.ReadAll(response.Body)
	err = ioutil.WriteFile("page0001.html", s, 0644)
	if err != nil {
		log.Printf("WriteFile err: %+v", err)
	}
	log.Printf("body: %+v", string(s))

}
