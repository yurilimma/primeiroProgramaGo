package html

import (
	"io/ioutil"
	"net/http"
	"regexp"
)

//Titulo obtem o titulo de uma página html
func Titulo(urls ...string) <-chan string {
	c := make(chan string)
	for _, url := range urls {
		//função anonima
		go func(url string) {
			//pega a pagina html via url
			resp, _ := http.Get(url)
			//pega o body da pagina html
			html, _ := ioutil.ReadAll(resp.Body)
			//pega o titulo da página html
			r, _ := regexp.Compile("<title>(.*?)<\\/title>")
			//faz um cast pra string do titulo da pagina html
			c <- r.FindStringSubmatch(string(html))[1]
		}(url)
	}
	return c
}
