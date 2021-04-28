package main

import (
	"fmt"
	"log"

	"github.com/PuerkitoBio/goquery"
)

const (
	ACCEPTED = "Accepted"
)
var(
	ref, status, row []string
	rows [][]string
)

func postScrape() {

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
		fmt.Println("####### headings = ", len(ref), ref)
		fmt.Println("####### status = ", len(status), status)
		fmt.Println("####### rows = ", len(rows), rows)

		for i, v := range status {

			if v == ACCEPTED {
				fmt.Println(ref[i])
			}
		}
}




func main() {
	postScrape()
}