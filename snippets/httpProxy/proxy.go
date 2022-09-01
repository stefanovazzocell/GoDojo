package main

import (
	"io"
	"log"
	"net/http"
	"time"
)

type Proxy struct {
	client      http.Client
	forwardAddr string
}

func (p Proxy) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	frwReq, err := http.NewRequest(req.Method, p.forwardAddr+req.URL.RawPath+"?"+req.URL.RawQuery, req.Body)
	if err != nil {
		log.Println(err)
		res.WriteHeader(500)
		res.Write([]byte("500 Internal Server Error"))
		return
	}
	frwRes, err := p.client.Do(frwReq)
	if err != nil {
		log.Println(err)
		res.WriteHeader(500)
		res.Write([]byte("500 Internal Server Error"))
		return
	}
	res.WriteHeader(frwRes.StatusCode)
	io.Copy(res, frwRes.Body)
	frwRes.Body.Close()
}

func NewProxy(forwardAddr string) *Proxy {
	return &Proxy{
		client: http.Client{
			Timeout: 5 * time.Second,
		},
		forwardAddr: forwardAddr,
	}
}
