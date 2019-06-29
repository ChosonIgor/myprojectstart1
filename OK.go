package main

import (
	"encoding/json"
	"log"
	"net/http"
)
func main() {
	http.HandleFunc("/", myController)
	log.Fatal(http.ListenAndServe(":3000", nil))
}
func myController(w http.ResponseWriter, r *http.Request) {
	a := Test{Text: "OK"}
	jsonString := r.Body
	mess,_ := json.Marshal(r.Body)
	//json.NewDecoder(r.Body).Decode(mess)
	//value, ok := mess["text"]
	/*if ok {
			if value == "OK"{
				w.Write([]byte(mess["text"]))
			}
	*/.
	w.Write([]byte("OK"))
}

type Test struct{
	Test1 string `json: "test1"`
	
}