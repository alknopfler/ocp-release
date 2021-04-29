package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/PuerkitoBio/goquery"
)

const (
	ACCEPTED  = "Accepted"
	SUCCESS   = " Succeeded"
	NONE      = "None"
	URL_BASE  = "https://openshift-release.apps.ci.l2s4.p1.openshiftapps.com/"
	URL_MID   = URL_BASE + "releasestream/"
	URL_END   = "/release/"
	ERROR     = "Error"
	ERROR_DOC = "Error reading URL"
	VERSION   = "4.8.0-0"
	OUTPUT    = "registry.ci.openshift.org/ocp/release:"
)

var (
	ref, status, row []string
	rows             [][]string
	ex               bool
)

func getTagFromVersion(version string) string {
	switch version {
	case "ci", "CI", "Ci":
		return VERSION + ".ci"
	default:
		return VERSION + ".nightly"
	}
}

func getBestReleaseCandidate(version, cond string) {

	doc, err := goquery.NewDocument(URL_BASE + "#" + getTagFromVersion(version))
	if err != nil {
		log.Fatal(err)
		return
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

	for i, v := range status {

		if v == ACCEPTED {
			doc2, err := goquery.NewDocument(URL_MID + getTagFromVersion(version) + URL_END + ref[i])
			if err != nil {
				log.Fatal(err)
				return
			}
			doc2.Find("body > div > ul:nth-child(7) > li:nth-child(2) > ul").Each(func(index int, tablehtml *goquery.Selection) {
				tablehtml.Find("li").Each(func(indextr int, rowhtml *goquery.Selection) {
					rowhtml.Find(".text-success").Each(func(indexth int, tablecell *goquery.Selection) {
						if tablecell.Text() == cond+SUCCESS {
							//log.Println("------encontracdo----> " + tablecell.Text())
							fmt.Println(OUTPUT + ref[i])
							ex = true
						}
					})

				})
			})
			if ex {
				return
			}
		}
	}
}

func main() {
	v := flag.String("v", "nightly", "Version to be used: [ci|nightly] ")
	c := flag.String("c", "metal-assisted", "condition to be present. Ex: 'metal-assisted'")
	h := flag.Bool("h", false, "Help")
	flag.Parse()

	if *h {
		flag.PrintDefaults()
		os.Exit(1)
	}
	if len(os.Args) < 2 {
		flag.PrintDefaults()
		os.Exit(1)
	}

	getBestReleaseCandidate(*v, *c)
}
