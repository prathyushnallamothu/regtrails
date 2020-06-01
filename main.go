package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/prathyushnallamothu/cleverdbconnection"
)
var Email,Password string
var tpl *template.Template
type data struct{
	Emailid string
	Username string
	Password string
	Country string
	Occupation string
}
func init(){
	tpl=template.Must(template.ParseGlob("templates/*.html"))
}
func main(){
	m:=mux.NewRouter().StrictSlash(true)
	m.HandleFunc("/",homehandler)
	m.HandleFunc("/register",registerhandler)
	m.HandleFunc("/registersucess",registersucesshandler)
	m.HandleFunc("/login",loginhandler)
	m.HandleFunc("/loginprocess",loginprocesshandler)
	m.HandleFunc("/dashboard",dashboardhandler)
	http.ListenAndServe(":8080",m)
}
func homehandler(w http.ResponseWriter,r *http.Request){
	tpl.ExecuteTemplate(w,"register.html",nil)
}
func registerhandler(w http.ResponseWriter,r *http.Request){
	x:=data{
		Emailid: r.FormValue("emailid"),
		Username: r.FormValue("username"),
		Password: r.FormValue("password"),
		Country: r.FormValue("country"),
		Occupation: r.FormValue("occupation"),
	}
	db:=dbconnection.Connect()
	defer db.Close()
	result,err:=db.Query("select (emailid) from registration where emailid=?",x.Emailid)
	if err!=nil{
		log.Fatal(err)
	}
	for result.Next(){
		err=result.Scan(&Email)
		if err!=nil{
			log.Fatal(err)
		}
	}
	if Email==x.Emailid{
		fmt.Fprintf(w,"already registered please use another email")
	}
	if Email!=x.Emailid{
		result1,err1:=db.Query("insert into registration(emailid,username,password,country,occupation) values(?,?,?,?,?)",x.Emailid,x.Username,x.Password,x.Country,x.Occupation)
		if err1!=nil{
			log.Fatal(err1)
		}
		fmt.Println(result1)
		http.Redirect(w,r,"/registersucess",307)
	}
}
func registersucesshandler(w http.ResponseWriter,r *http.Request){
	fmt.Fprintf(w,"registered sucessfully,will update login page shortly kindly please wait")
	http.Redirect(w,r,"/login",307)
}
func loginhandler(w http.ResponseWriter,r *http.Request){
	tmp.ExecuteTemplate(w,"login.html",nil)
}
func loginprocesshandler(w http.ResponseWriter,r *http.Request){
	x:=data{
		Emailid: r.FormValue("emailid"),
		Password: r.FormValue("password"),
	}
	db:=dbconnection.Connect()
	defer db.Close()
	result,err:=db.Query("select (emailid,password) from registration where emailid=?",x.Emailid)
	if err!=nil{
		log.Fatal(err)
	}
	for result.Next(){
		err=result.Scan(&Email,&Password)
		if err!=nil{
			log.Fatal(err)
		}
	}
	if x.Emailid==Email{
		if x.Password==Password{
			http.Redirect(w,r,"/dashboard",307)
		}else{
			fmt.Fprintf(w,"email or password incorrect")
		}
	}
	if x.Emailid!=Email{
		fmt.Fprintf(w,"please register to pran")
	}
}
func dashboardhandler(w http.ResponseWriter,r *http.Request){
	fmt.Fprintf(w,"WELCOME TO PRAN,THIS SITE IS STILL UNDER CONSTRUCTION")
}
