// Package fountain support encoding/decoding fountain screenplay markup
//
// scrippets.go manages fetching and inlining or generating links to John August's scrippets.css
//
package fountain

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
)

var (
	scrippetsCSSUrl = "https://johnaugust.com/wp-content/plugins/wp-scrippets/scrippets.css?v2.0"
)

func getScrippetsCSS() string {
	var (
		scrippetsCSS string
	)
	// 1. Find where we've cached scrippets.css
	if _, err := os.Stat("scrippets.css"); os.IsNotExist(err) == false {
		scrippetsCSS = "scrippets.css"
	} else if _, err := os.Stat(path.Join("css", "scrippets.css")); os.IsNotExist(err) == false {
		scrippetsCSS = path.Join("css", "scrippets.css")
	}
	// otherwise download it
	if scrippetsCSS == "" {
		resp, err := http.Get(scrippetsCSSUrl)
		if err != nil {
			// handle error
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Printf("%s", err)
			return fmt.Sprintf(`<link rel="stylesheet" href=%q>`, scrippetsCSSUrl)
		}
		err = ioutil.WriteFile("scrippets.css", body, 0666)
		return fmt.Sprintf("%s", body)
	}
	src, err := ioutil.ReadFile(scrippetsCSS)
	if err != nil {
		log.Printf("%s", err)
		return fmt.Sprintf(`<link rel="stylesheet" href=%q>`, scrippetsCSSUrl)
	}
	return fmt.Sprintf("%s", src)
}

func getScrippetsCSSLink() string {
	return fmt.Sprintf(`<link rel="stylesheet" href=%q>`, scrippetsCSSUrl)
}
