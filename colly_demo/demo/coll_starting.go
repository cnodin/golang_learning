package main

import (
	"github.com/gocolly/colly"
	"fmt"
	"net/http"
	"net"
	"time"
	"log"
)

func main() {
	c := colly.NewCollector(
		//colly.MaxDepth(2),
		colly.Async(true),
	)
	c.WithTransport(&http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout: 30 * time.Second,
			KeepAlive: 30 * time.Second,
			DualStack: true,
		}).DialContext,
		MaxIdleConns: 10000,
		IdleConnTimeout: 90 * time.Second,
		TLSHandshakeTimeout: 10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	})
	c.OnHTML("a", func(e *colly.HTMLElement) {
		e.Request.Visit(e.Attr("href"))
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.OnResponse(func(resp *colly.Response) {
		log.Println("url:", resp.Request.URL)
		log.Println("response:", string(resp.Body))
	})

	c.Visit("http://go-colly.org")
	c.Wait()
}
