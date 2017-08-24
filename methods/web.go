package main

import (
	"fmt"
	"image"
	"net/http"
	"time"
)

type timeHandler struct {
	zone *time.Location
}

func (th *timeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	tm := time.Now().In(th.zone).Format(time.RFC1123)
	w.Write([]byte("The time is: " + tm))
}

func newTimeHandler(name string) *timeHandler {
	return &timeHandler{zone: time.FixedZone(name, 0)}
}

func main() {
	myHandler := newTimeHandler("EST")
	//Custom http server
	s := &http.Server{
		Addr:           ":4000",
		Handler:        myHandler,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	err := s.ListenAndServe()
	if err != nil {
		fmt.Printf("Server failed: ", err.Error())
	}
	m := image.NewRGBA(image.Rect(0, 0, 100, 100))
	fmt.Println(m.Bounds())
	fmt.Println(m.At(0, 0).RGBA())
}
