package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	//"github.com/prathyushnallamothu/mycookie"
	"github.com/gorilla/mux"
	"github.com/prathyushnallamothu/cleverdbconnection"
)
var Email,Password,Username,Cname string
var tpl *template.Template
type Developerdata struct{
	Emailid string
	Username string
	Password string
	Skills   string
}
type Companydata struct{
	Emailid string
	Companyname string
	Password string
	Url   string
}
type Mainstruct struct{
	Value string
	Developer []Developerdata
	Company  []Companydata
}
func init(){
	tpl=template.Must(template.ParseGlob("templates/*.html"))
}
func main(){
	m:=mux.NewRouter().StrictSlash(true)
	m.PathPrefix("/images/").Handler(http.StripPrefix("/images/",http.FileServer(http.Dir("./images"))))
	m.PathPrefix("/css/").Handler(http.StripPrefix("/css/",http.FileServer(http.Dir("./css"))))
	m.PathPrefix("/fonts/").Handler(http.StripPrefix("/fonts/",http.FileServer(http.Dir("./fonts"))))
	m.PathPrefix("/js/").Handler(http.StripPrefix("/js/",http.FileServer(http.Dir("./js"))))
	m.PathPrefix("/sass/").Handler(http.StripPrefix("/sass/",http.FileServer(http.Dir("./sass"))))
	m.HandleFunc("/",homehandler)
	m.HandleFunc("/companyregister",companyregisterhandler)
	m.HandleFunc("/register",registerhandler)
	m.HandleFunc("/registersucessful",registersucessfulhandler)
	m.HandleFunc("/registersucess",registersucesshandler)
	m.HandleFunc("/login",loginhandler)
	m.HandleFunc("/loginprocess",loginprocesshandler)
	m.HandleFunc("/dashboard",dashboardhandler)
	m.HandleFunc("/logout",logouthandler)
	http.ListenAndServe(":8080",m)
}
func homehandler(w http.ResponseWriter,r *http.Request){
	tpl.ExecuteTemplate(w,"index.html",nil)
}
func companyregisterhandler(w http.ResponseWriter,r *http.Request){
http.Redirect(w,r,"/register?q=company",307)
}
func registerhandler(w http.ResponseWriter,r *http.Request){
	if r.FormValue("q")=="developers"{
		tpl.ExecuteTemplate(w,"developerregister.html",nil)
	}
	if r.FormValue("q")=="company"{
		tpl.ExecuteTemplate(w,"companyregister.html",nil)
	}
}
func registersucessfulhandler(w http.ResponseWriter,r *http.Request){
	if r.FormValue("q")=="developers"{
	x:=Developerdata{
		Emailid: r.FormValue("emailid"),
		Username: r.FormValue("username"),
		Password: r.FormValue("password"),
		Skills: r.FormValue("skills"),
	}
	db:=dbconnection.Connect()
	defer db.Close()
	result,err:=db.Query("select emailid from developer where emailid=?",x.Emailid)
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
		result1,err1:=db.Query("insert into developer(emailid,username,password,skills) values(?,?,?,?)",x.Emailid,x.Username,x.Password,x.Skills)
		if err1!=nil{
			log.Fatal(err1)
		}
		fmt.Println(result1)
		http.Redirect(w,r,"/registersucess",307)
	}
	}
	if r.FormValue("q")=="company"{
		y:=Companydata{
			Emailid: r.FormValue("emailid"),
			Companyname: r.FormValue("cname"),
			Password: r.FormValue("password"),
			Url: r.FormValue("url"),
		}
		db:=dbconnection.Connect()
		defer db.Close()
		result,err:=db.Query("select (emailid) from company where emailid=?",y.Emailid)
		if err!=nil{
			log.Fatal(err)
		}
		for result.Next(){
			err=result.Scan(&Email)
			if err!=nil{
				log.Fatal(err)
			}
		}
		if Email==y.Emailid{
			fmt.Fprintf(w,"already registered please use another email")
		}
		if Email!=y.Emailid{
			result1,err1:=db.Query("insert into company(emailid,cname,password,url) values(?,?,?,?)",y.Emailid,y.Companyname,y.Password,y.Url)
			if err1!=nil{
				log.Fatal(err1)
			}
			fmt.Println(result1)
			http.Redirect(w,r,"/registersucess",307)
		}
	}
}
func registersucesshandler(w http.ResponseWriter,r *http.Request){
	fmt.Fprintf(w,"sucessfully registered")
}
func loginhandler(w http.ResponseWriter,r *http.Request){
	if r.FormValue("q")=="developers"{
	tpl.ExecuteTemplate(w,"devlogin.html",nil)
}
	if r.FormValue("q")=="company"{
		tpl.ExecuteTemplate(w,"companylogin.html",nil)
	}
}
func loginprocesshandler(w http.ResponseWriter,r *http.Request){
	if r.FormValue("q")=="developers"{
	x:=Developerdata{
		Emailid: r.FormValue("emailid"),
		Password: r.FormValue("password"),
	}
	db:=dbconnection.Connect()
	defer db.Close()
	result,err:=db.Query("select emailid,password from developer where emailid=?",x.Emailid)
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
			
			http.Redirect(w,r,"/dashboard?q=developers&email="+Email,307)
		}else{
			fmt.Fprintf(w,"email or password incorrect")
		}
	}
	if x.Emailid!=Email{
		fmt.Fprintf(w,"please register to pran")
	}
}
if r.FormValue("q")=="company"{
	x:=Companydata{
		Emailid: r.FormValue("emailid"),
		Password: r.FormValue("password"),
	}
	db:=dbconnection.Connect()
	defer db.Close()
	result,err:=db.Query("select emailid,password from company where emailid=?",x.Emailid)
	if err!=nil{
		log.Fatal(err)
	}
	for result.Next(){
		err=result.Scan(&Email,&Password)
		if err!=nil{
			log.Fatal(err)
		}
	}
	if Email=="anukruthi.pulimi02@gmail.com"{
		if x.Password==Password{
			http.Redirect(w,r,"/dashboard?q=anu",307)
		}
	}
	if x.Emailid==Email{
		if x.Password==Password{
		
			http.Redirect(w,r,"/dashboard?q=company&email="+Email,307)
		}else{
			fmt.Fprintf(w,"email or password incorrect")
		}
	}
	if x.Emailid!=Email{
		fmt.Fprintf(w,"please register to pran")
	}
}
}
func dashboardhandler(w http.ResponseWriter,r *http.Request){
	
	if r.FormValue("q")=="developers"{
		x:=r.FormValue("email")
		db:=dbconnection.Connect()
		defer db.Close()
		result,err:=db.Query("select username from developer where emailid=?",x)
		if err!=nil{
			log.Fatal(err)
		}
		var y Developerdata
		var z Mainstruct
		for result.Next(){
			err=result.Scan(&Username)
			if err!=nil{
				log.Fatal(err)
			}
			y.Username=Username
			z.Developer=append(z.Developer,y)
		}
		z.Value="developer"
		if x=="anukruthi.pulimi02@gmail.com"{
			tpl.ExecuteTemplate(w,"anudashboard.html",z)
		}
		if x!="anukruthi.pulimi02@gmail.com"{
			tpl.ExecuteTemplate(w,"dashboard.html",z)
		}
				
	}
	if r.FormValue("q")=="company"{
		x:=r.FormValue("email")
		db:=dbconnection.Connect()
		defer db.Close()
		result1,err1:=db.Query("select cname from company where emailid=?",y)
		if err1!=nil{
			log.Fatal(err1)
		}
		var y Companydata
		var z Mainstruct
		for result1.Next(){
			err1=result1.Scan(&Cname)
			if err1!=nil{
				log.Fatal(err1)
			}
			y.Companyname=Cname
			z.Company=append(z.Company,y)
		}
		z.Value="company"
		tpl.ExecuteTemplate(w,"dashboard.html",z)		
	}
}
func logouthandler(w http.ResponseWriter,r *http.Request){
	if r.FormValue("q")=="developers"{
		http.Redirect(w,r,"/",307)
	}
	if r.FormValue("q")=="company"{
		http.Redirect(w,r,"/",307)
	}
}
