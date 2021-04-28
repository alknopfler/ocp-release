package main

import (
	"fmt"
	"log"

	"github.com/PuerkitoBio/goquery"
)

const (
	ACCEPTED = "Accepted"
	SUCCESS  = " Succeeded"
	NONE     = "None"
)

var (
	ref, status, row []string
	rows             [][]string
	ex               bool
)

func postScrape(cond string) string {

	doc, err := goquery.NewDocument("https://openshift-release.apps.ci.l2s4.p1.openshiftapps.com/#4.8.0-0.nightly")
	if err != nil {
		log.Fatal(err)
	}

	doc.Find("body > div > div.row > div > table:nth-child(10) > tbody").Each(func(index int, tablehtml *goquery.Selection) {
		tablehtml.Find("tr").Each(func(indextr int, rowhtml *goquery.Selection) {
			rowhtml.Find("td").Each(func(indexth int, tablecell *goquery.Selection) {
				row = append(row, tablecell.Text())
				if indexth == 0 {
					ref = append(ref, tablecell.Text())
				}
				if indexth == 1 {
					status = append(status, tablecell.Text())
				}

			})
			rows = append(rows, row)
			row = nil
		})
	})
	//log.Println("####### headings = ", len(ref), ref)
	//log.Println("####### status = ", len(status), status)
	//log.Println("####### rows = ", len(rows), rows)

	for i, v := range status {

		if v == ACCEPTED {
			//log.Println("selected -----> " + ref[i])

			doc2, err := goquery.NewDocument("https://openshift-release.apps.ci.l2s4.p1.openshiftapps.com/releasestream/4.8.0-0.nightly/release/" + ref[i])
			if err != nil {
				log.Fatal(err)
			}
			doc2.Find("body > div > ul:nth-child(7) > li:nth-child(2) > ul").Each(func(index int, tablehtml *goquery.Selection) {
				tablehtml.Find("li").Each(func(indextr int, rowhtml *goquery.Selection) {
					rowhtml.Find(".text-success").Each(func(indexth int, tablecell *goquery.Selection) {
						if tablecell.Text() == cond+SUCCESS {
							//log.Println("------encontracdo----> " + tablecell.Text())
							ex = true
						}
					})

				})
			})
			if ex {
				return ref[i]
			}
		}
	}
	return NONE
}

func main() {
	fmt.Println(postScrape("metal-assisted"))
}
