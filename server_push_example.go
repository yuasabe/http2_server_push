package main

import (
//"os"
"fmt"
"io/ioutil"
"net/http"
)

var image []byte
var image2 []byte

func init() {
	var err error
	image, err = ioutil.ReadFile("./image.jpg")
	image2, err = ioutil.ReadFile("./image2.jpg")
	if err != nil {
		panic(err)
	}
}

func handlerHtml(w http.ResponseWriter, r *http.Request) {

	pusher, ok := w.(http.Pusher)
	if ok {
		pusher.Push("image", nil)
		pusher.Push("./image2.jpg", nil)
		pusher.Push("/css/style.css", nil)
		pusher.Push("/css/a_new.css", nil)
		pusher.Push("/js/script.js", nil)
		pusher.Push("/js/another_script.js", nil)
	}

	w.Header().Add("Content-Type", "text/html")
	fmt.Fprintf(w, indexHTML)
}

func handlerImage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/jpeg")
	// w.Write(image)
	w.Write(image2)
}

func main() {
//	os.Setenv("GODEBUG", "http2server=0") // Access with HTTP/1.1
	http.HandleFunc("/", handlerHtml)
	http.HandleFunc("/image", handlerImage)
	fmt.Println("start http listening :18443")
	err := http.ListenAndServeTLS(":443", "cert.pem", "privkey.pem", nil)
	fmt.Println(err)
}

const indexHTML = `
<html>
<head>
	<link rel="stylesheet" href="/css/style.css">
	<script src="/js/script.js"></script>
	<link rel="stylesheet" href="/css/a_new.css">
	<script src="/js/another_script.js"></script>
</head>
<body>
	<img src="/image"/>
	<img src="image2.jpg">
</body>
</html>
`
