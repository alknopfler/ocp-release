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

func getBestReleaseCandidate(version string, cond []string) {

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
	num := 0
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
						for v := range cond {
							if tablecell.Text() == cond[v]+SUCCESS {
								//log.Println("------encontracdo----> " + tablecell.Text())
								//fmt.Println(OUTPUT + ref[i])
								ex = true
								num++
							}
						}
					})

				})
			})
			doc2.Find("body > div > ul:nth-child(7) > li:nth-child(1) > ul").Each(func(index int, tablehtml *goquery.Selection) {
				tablehtml.Find("li").Each(func(indextr int, rowhtml *goquery.Selection) {
					rowhtml.Find(".text-success").Each(func(indexth int, tablecell *goquery.Selection) {
						for v := range cond {
							if tablecell.Text() == cond[v]+SUCCESS {
								//log.Println("------encontracdo----> " + tablecell.Text())
								//fmt.Println(OUTPUT + ref[i])
								ex = true
								num++
							}
						}
					})

				})
			})
			if ex && len(cond) == num {
				fmt.Println(OUTPUT + ref[i])
				return
			}
		}
	}
}

func main() {

	v := flag.String("v", "", "Version to be used: -v [ci|nightly] ")
	flag.Bool("c", false, "condition array to be present. Ex: '-c metal-assisted aws metal-ipi'")
	h := flag.Bool("h", false, "Help usage example: ./ocp-release -v nightly -c metal-ipi aws gcp")
	flag.Parse()
	tail := flag.Args()

	if *h {
		flag.PrintDefaults()
		os.Exit(1)
	}
	if len(tail) < 1 {
		flag.PrintDefaults()
		fmt.Println("sssss")
		os.Exit(2)
	}

	getBestReleaseCandidate(*v, tail)
}
