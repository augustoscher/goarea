package html

import (
	"io/ioutil"
	"net/http"
	"regexp"
)

// Title busca o titulo de uma página html com base nas URLs passadas e retorna através de um canal
func Title(urls ...string) <-chan string {
	c := make(chan string)
	for _, url := range urls {
		go func(url string) {
			result, _ := http.Get(url)
			html, _ := ioutil.ReadAll(result.Body)

			r, _ := regexp.Compile("<title>(.*?)<\\/title>")
			c <- r.FindStringSubmatch(string(html))[1]
		}(url)
	}

	// O retorno é executado antes das requests terminarem
	// Quem ler o canal irá receber os retornos a medida que as request terminarem
	return c
}
