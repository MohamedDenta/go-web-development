package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/cookiejar"
)

// Use Go as a client

var client http.Client

func init() {
	jar, err := cookiejar.New(nil)
	if err != nil {
		fmt.Println("Error : ", err)
	}
	client = http.Client{
		Jar: jar,
	}
}
func makeExtReq() {
	cookie := &http.Cookie{
		Name:   "token",
		Value:  "some_token",
		MaxAge: 300,
	}
	cookie2 := &http.Cookie{
		Name:   "clicked",
		Value:  "true",
		MaxAge: 300,
	}
	// new req
	req, err := http.NewRequest("GET", "https://google.com", nil)
	if err != nil {
		log.Fatalf("Got error %s", err.Error())
	}
	// add cookies
	req.AddCookie(cookie)
	req.AddCookie(cookie2)

	for _, c := range req.Cookies() {
		fmt.Println(c)
	}

	// send req
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error occured. Error is: %s", err.Error())
	}
	defer resp.Body.Close()
	fmt.Printf("StatusCode: %d\n", resp.StatusCode)
}
func main() {
	makeExtReq()
	http.HandleFunc("/doc", docHandler)
	http.ListenAndServe(":8080", nil)
}

func docHandler(w http.ResponseWriter, r *http.Request) {
	cs := r.Cookies()
	fmt.Println(cs)
	// c, err := r.Cookie("id")
	// if err != nil {
	// w.Write([]byte(err.Error() + "\n"))
	if len(cs) < 1 {
		cookie := &http.Cookie{
			Name:   "id",
			Value:  "asd",
			MaxAge: 300,
		}
		http.SetCookie(w, cookie)
		w.WriteHeader(200)
		w.Write([]byte("Doc Get Successful"))
	} else {
		w.WriteHeader(200)
		w.Write([]byte("Cookie exists"))
	}

	// } else {
	// 	w.Write([]byte(c.Value))
	// }

}
