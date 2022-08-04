// Package fountain is a Golang package supporting Fountain screenplay markup.
//
// css.go manages setting up CSS for either inline style elements or links.
//
package fountain

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

var (
	// SourceCSS is a default CSS to use when rendering
	SourceCSS = `
/**
 * fountain.css - CSS for displaying foutain2html generated HTML.
 * It was inspired by scrippet.css found on the Fountain
 * website at https://fountain.io/_css/scrippets.css which is attributed
 * to John August, updated in 2012.
 *
 * 2019-06-18, RSD
 */

.fountain {
    margin: 0;
    padding: 0;
    display: block;
}

.fountain {
    max-width: 400px;
    background: #fffffc;
    color: #000000;
    padding: 5px 14px 15px 14px !important;
    clear: both;
    margin-bottom: 2.5em;
    margin-top: 2.5em;
}

section.title-page,
section.script {
    width: 36em;
    padding-left: 1em;
    padding-bottom: 2em;
    margin-bottom: 2em;
    border: 0.12em solid #d2d2d2;
    border-radius: 3px;
}

.title-page,
.script {
    min-height: 20em;
}

/* Simulate a page feed */
.script .page-feed {
    overflow: visible;
    border: 0;
    padding: 0;
    margin: 0;
    margin-top: 4em;
    margin-left: -1em;
    /* For IE */ 
    height: 30px;
    border-style: solid;
    border-color: black;
    border-width: 1px 0 0 0;
    border-radius: 8px;
} 
.script .page-feed:before { 
    display: block;
    content: "";
    height: 30px;
    border: 0;
    padding: 0;
    margin: 0;
    margin-top: -31px;
    border-style: solid;
    border-color: black;
    border-width: 0 0 1px 0;
    border-radius: 8px;
}


.title {
    text-align: center;
    padding-left: 33%;
    padding-right: 33%;
    text-transform: uppercase;
    text-decoration: underline;
    margin-top: 1em;
    margin-bottom: 1em;
}

.author {
    text-align: center;
    padding-left: 33%;
    padding-right: 33%;
    margin-top: 0em;
    margin-bottom: 0em;
}

.draft-date,
.date {
    text-align: center;
    padding-left: 33%;
    padding-right: 33%;
    margin-top: 0;
    margin-bottom: 6em;
}

.copyright {
    display: block;
    padding: 0;
    margin: 0;
    text-align: left;
    text-transform: none;
    text-decoration: none;
}

.contact {
    display: block;
    padding: 0;
    margin: 0;
    text-align: left;
    text-transform: none;
    text-decoration: none;
}

.script {
    padding-top: 2em;
    padding-left: 0;
    padding-right: 0;
    padding-bottom: 2em;
}

.scene-heading,
.action,
.character,
.parenthetical,
.transition,
.dialogue {
    font: 12px/14px Courier, "Courier New", monospace;
    text-align: left !important;
    letter-spacing: 0 !important;
    margin-top: 0px !important;
    margin-bottom: 0px !important;
    display: block;
}

.scene-heading,
.action,
.character {
    padding-top: 1.5ex !important;
    display: block;
}

.action {
    padding-right: 5% !important;
    font-size: 12px !important;
    line-height: 14px !important;
}

.character {
    padding-left: 40% !important;
}

.dialogue {
    padding-left: 20% !important;
    padding-right: 20% !important;
}

.parenthetical {
    display: block;
    padding-left: 32% !important;
    padding-right: 30% !important;
}

.dialogue+.parenthetical {
    padding-bottom: 0 !important;
}

.left-align {
    float: left;
    padding-left: 2em;
    text-align: left;
}

.centered {
    padding-left: 33%;
    padding-right: 33%;
    text-align: center;
}

.right-align {
    float: right;
    padding-right: 2em;
    text-align: right;
}

.empty {
    display: none;
    height: 0;
}
`
)

// getCSS() checks to see if there is any custom CSS filenamed fountain.css
// in the current work directory or in the CSS folder and gets that
// otherwise it'll fall back the value in SourceCSS.
func getCSS() (string, error) {
	var (
		src []byte
		err error
	)
	// NOTE: If CSS value not set then look for fountain.css in
	// current work directory and in css/fountain.css before falling
	// back to the default SourceCSS value.
	if CSS == "" {
		// 1. Find where we've put any custom CSS
		if _, err = os.Stat("fountain.css"); os.IsNotExist(err) == false {
			CSS = "fountain.css"
		} else if _, err = os.Stat(path.Join("css", "fountain.css")); os.IsNotExist(err) == false {
			CSS = path.Join("css", "fountain.css")
		}
	}
	src, err = ioutil.ReadFile(CSS)
	if err != nil {
		return "", err
	} else {
		// Replace default CSS with requested CSS
		SourceCSS = fmt.Sprintf("%s", src)
	}
	// 2. Otherwise provide default
	return createElement("style", []string{}, SourceCSS), nil
}

func getCSSLink() (string, error) {
	var err error
	if strings.Contains(CSS, "://") == false {
		_, err = os.Stat(CSS)
	}
	return fmt.Sprintf("<link rel=%q href=%q>\n", "stylesheet", CSS), err
}
