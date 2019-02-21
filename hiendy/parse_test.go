package hiendy

import (
	"bytes"
	"io/ioutil"
	"log"
	"testing"

	"github.com/PuerkitoBio/goquery"
)

func TestParseHtml(t *testing.T) {
	// err := ParseHtml()
	// if err != nil {
	// 	t.Log("failed !")
	// } else {
	// 	t.Log("success !")
	// }

}

func TestParseTest(t *testing.T) {
	body, err := ioutil.ReadFile("page0001.html")
	if err != nil {
		log.Printf("read file err: %+v", err)
	}
	r := bytes.NewReader(body)

	dom, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		log.Printf("NewDocumentFromReader err: %+v", err)
	}

	dom.Find("div.entry-content source").Each(func(i int, selection *goquery.Selection) {
		if len(selection.Nodes) > 1 {
			for index := 0; index < len(selection.Nodes); index++ {
				log.Printf("nodes: %+v", selection.Nodes[index])
			}
		} else if len(selection.Nodes) == 1 {
			log.Printf("nodes: %+v", selection.Nodes[0])
		}
	})
}
