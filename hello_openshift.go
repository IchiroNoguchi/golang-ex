package main

import (
	"fmt"
	"net/http"
	"os"

)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	s :=  `
<!DOCTYPE html>
<html>
<head>
<script src="https://ajax.googleapis.com/ajax/libs/jquery/3.5.1/jquery.min.js"></script>
<script> 
$(document).ready(function(){

    $("div").animate({left: '0px'}, 5000, 'swing');
});
</script> 
</head>
<body>



<div style="height:100px;width:100px;position:absolute;left: 2000px;">
<img width="100px" src="https://cdn.worldvectorlogo.com/logos/gopher.svg">
</div>

</body>
</html>

	` 
	response := os.Getenv("RESPONSE")
	if len(response) == 0 {
		response =  s
	}

	fmt.Fprintln(w, response)

	fmt.Println(response)
}

func listenAndServe(port string) {
	fmt.Printf("serving on %s\n", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}

func main() {
	http.HandleFunc("/", helloHandler)
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}
	go listenAndServe(port)

	port = os.Getenv("SECOND_PORT")
	if len(port) == 0 {
		port = "8888"
	}
	go listenAndServe(port)

	select {}
}
