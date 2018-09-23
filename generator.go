package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

// Google I/O 2012 - Go Concurrency Patterns

// <-chan - canal somente-leitura
func titulo(urls ...string) <-chan string {
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

func main() {
	t1 := titulo("https://www.cod3r.com.br", "https://www.google.com")
	t2 := titulo("https://www.amazon.com", "https://www.youtube.com")
	fmt.Println("Primeiros:", <-t1, "|", <-t2)
	fmt.Println("Segundos:", <-t1, "|", <-t2)
}
