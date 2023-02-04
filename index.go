package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path"
)

//for page
type HtmlContext struct{
	Title string
	Description string
}
//for html
type JsonRes struct{
	ServerName string `json:"server_name"`
}
//function as url path
func textResponse(res http.ResponseWriter,req *http.Request){
	if req.URL.Path !="/"{
		http.Error(res,"you can access some route for exp: \"/\", ",http.StatusNotFound)
		return;
	}
	if req.Method !="GET"{
		http.Error(res,"you can access from GET",http.StatusNotFound)
		return;
	}
	fmt.Fprintln(res,`
	       Help
	   -------------
	name          path
	----          ---- 
	home   ->     /home
	about  ->     /about
	json   ->     /json
	`)
	
}
func homeFunc(res http.ResponseWriter,req *http.Request){
	if req.Method !="GET"{
		http.Error(res,"you can access from GET",http.StatusNotFound)
		return;
	}
	data :=HtmlContext{Title:"Home",Description:"blank Description"}
	filepath  := path.Join("templates","render.html")
	if htmlTemplate,err:=template.ParseFiles(filepath);err==nil{
		if err:= htmlTemplate.Execute(res,data);err!=nil{
			http.Error(res, err.Error(), http.StatusInternalServerError)
		}
	}else{
		http.Error(res, err.Error(), http.StatusInternalServerError)
        return
	}
}
func aboutFunc(res http.ResponseWriter,req *http.Request){
	if req.Method !="GET"{
		http.Error(res,"you can access from GET",http.StatusNotFound)
		return;
	}
	data :=HtmlContext{Title:"about",Description:"About page...."}
	filepath  := path.Join("templates","render.html")
	if htmlTemplate,err:=template.ParseFiles(filepath);err==nil{
		if err:= htmlTemplate.Execute(res,data);err!=nil{
			http.Error(res, err.Error(), http.StatusInternalServerError)
		}
	}else{
		http.Error(res, err.Error(), http.StatusInternalServerError)
        return
	}
}
func jsonResponce(res http.ResponseWriter,req *http.Request){
	res.Header().Set("Content-Type", "application/json")
	data :=JsonRes{ServerName:"go server "} 
	byteJson,err:= json.Marshal(data)
	if err != nil {
		http.Error(res,"some thing wrong",http.StatusNotFound)
		return;
	}
	json.NewEncoder(res).Encode(string(byteJson))
}
func main(){
	 const PORT =":8000"
	http.HandleFunc("/",textResponse)
	http.HandleFunc("/home",homeFunc)
	http.HandleFunc("/about",aboutFunc)
	http.HandleFunc("/json",jsonResponce)
	fmt.Printf("server start on %v || http://localhost%v",PORT,PORT)
	if err:=http.ListenAndServe(PORT,nil);err!=nil{
		log.Fatal("server start failed")
	}
}