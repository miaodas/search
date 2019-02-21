package hiendy

import (
	"errors"
	"io"
	// "io/ioutil"
	"log"
	// "strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/jinzhu/gorm"
	"golang.org/x/net/html"
)

func ParseHtml(r io.Reader, c chan int) {
	// body, err := ioutil.ReadFile("page0001.html")
	// if err != nil {
	// 	log.Printf("read file err: %+v", err)
	// 	return err
	// }

	dom, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		log.Printf("NewDocumentFromReader err: %+v", err)
	}

	conn, err := GetConnection()

	dom.Find("div.entry-content source").Each(func(i int, selection *goquery.Selection) {
		if len(selection.Nodes) > 1 {
			for index := 0; index < len(selection.Nodes); index++ {
				for _, value := range selection.Nodes[index].Attr {
					// if value.Key == "file" {
					// 	GetImageUrl(&value, conn)
					// }
					if value.Key == "src" {
						GetUrl(&value, conn, "video")
					}
				}
			}
		} else if len(selection.Nodes) == 1 {
			for _, value := range selection.Nodes[0].Attr {
				// if value.Key == "file" {
				// 	GetImageUrl(&value, conn)
				// }
				if value.Key == "src" {
					GetUrl(&value, conn, "video")
				}
			}
		}
	})

	// dom.Find("div.entry-content img").Each(func(i int, selection *goquery.Selection) {
	// 	if len(selection.Nodes) > 1 {
	// 		for index := 0; index < len(selection.Nodes); index++ {
	// 			for _, value := range selection.Nodes[index].Attr {
	// 				// if value.Key == "file" {
	// 				// 	GetImageUrl(&value, conn)
	// 				// }
	// 				if value.Key == "src" {
	// 					GetUrl(&value, conn, "image")
	// 				}
	// 			}
	// 		}
	// 	} else if len(selection.Nodes) == 1 {
	// 		for _, value := range selection.Nodes[0].Attr {
	// 			// if value.Key == "file" {
	// 			// 	GetImageUrl(&value, conn)
	// 			// }
	// 			if value.Key == "src" {
	// 				GetUrl(&value, conn, "image")
	// 			}
	// 		}
	// 	}
	// })
	// log.Printf("read file success: %+v", bodyStr)
	log.Printf("ParseHtml before chan!")
	c <- 1
	log.Printf("ParseHtml finish!")
	// return nil
}

// save to mysql
func GetUrl(attr *html.Attribute, conn *gorm.DB, types string) {
	res := &Resource{
		Url:  attr.Val,
		From: "https://huotu42.com/",
		Type: types,
	}
	err := SaveResource(conn, res)
	if err != nil {
		log.Printf("save err: %s", err)
	}
}

// SaveResource
func SaveResource(conn *gorm.DB, resource *Resource) (err error) {
	conn.NewRecord(resource)
	conn.Create(resource)
	if conn.NewRecord(resource) {
		// fail
		err = errors.New("save failed")
	}
	return err
}
