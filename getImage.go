package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gocolly/colly"
)

type Registry struct {
	URL string
}

const (
	urlBase    = "https://unsplash.com/t/"
	classPhoto = "._1hjZT._1jjdS._1CBrG._1WPby.xLon9.Onk5k._17avz._1B083._3d86A._22Rl1"
	classTopic = ".qvEaq._1CBrG"
)

func main() {
	dirBase := "./downloads/"
	_, err := os.Stat(dirBase)
	if err != nil && !os.IsExist(err) {
		if err := os.Mkdir(dirBase, 0744); err != nil && err != os.ErrExist {
			log.Fatal(err)
		}
	}

	search := classPhoto
	topic := flag.String("topic", "", "topic for search")
	flag.Parse()
	if *topic == "" {
		search = classTopic
	}

	registries := []Registry{}
	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.OnHTML(search, func(e *colly.HTMLElement) {
		reg := Registry{}
		reg.URL = e.Attr("href")
		registries = append(registries, reg)
	})

	c.OnScraped(func(r *colly.Response) {
		if *topic == "" {
			fmt.Println("Topics found")
			for _, r := range registries {
				if len(r.URL) > 3 {
					fmt.Println("   ", r.URL[3:])
				}
			}
			os.Exit(0)
		}

		for _, r := range registries {
			fmt.Println(r.URL)
		}
	})

	c.Visit(urlBase + *topic)
}

func download(path, file string) error {
	res, err := http.Get(path)
	if err != nil {
		return fmt.Errorf("download: %w", err)
	}
	defer res.Body.Close()
	f, err := os.Create(file)
	if err != nil {
		return fmt.Errorf("download: %w", err)
	}
	defer f.Close()
	if _, err := io.Copy(f, res.Body); err != nil {
		return fmt.Errorf("download: %w", err)
	}
	return nil
}
