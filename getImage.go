package main

import (
	"flag"
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/alexsetta/scrapper"
)

const (
	urlBase = "https://unsplash.com/t/"
)

var (
	rePhoto = regexp.MustCompile(`<a class=\"_1hjZT _1jjdS _1CBrG _1WPby xLon9 Onk5k _17avz _1B083 _3d86A _22Rl1\" title=\"Download photo\" href=\"(.*?)\"`)
	reTopic = regexp.MustCompile(`<span class="_1WMnM xLon9">(.*?)<`)
)

func main() {
	search := reTopic
	dirBase := "./downloads/"
	scrapper.MakeDir(dirBase)
	topic := flag.String("topic", "", "topic for search")
	flag.Parse()
	if *topic != "" {
		search = rePhoto
		scrapper.MakeDir(dirBase + *topic)
	}

	fmt.Println("Acessando site de imagens")
	links, err := scrapper.List(search, urlBase+*topic)
	if err != nil {
		log.Fatal(err)
	}

	for _, l := range links {
		if *topic != "" {
			fileName := strings.Split(l.Value, "/")[4]
			fmt.Println("baixando ", dirBase+*topic+"/"+fileName+".jpg")
			scrapper.Download(l.Value, dirBase+*topic+"/"+fileName+".jpg")
		} else {
			fmt.Println(l.Value)
		}
	}
}
