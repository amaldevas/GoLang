package main 
import("fmt"
		"net/http"
		"html/template")
var tpl *template.Template
func init(){
	tpl=template.Must(template.ParseGlob("templates/*.gohtml"))
}
func index_handler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"Whoa, Go is neat!")
}
func about_h(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"This is Go")
}
func form_handler(w http.ResponseWriter, r *http.Request){
	tpl.ExecuteTemplate(w,"form.gohtml",nil)
}
func processor(w http.ResponseWriter, r *http.Request){
	if r.Method != "POST"{

		http.Redirect(w,r,"/form/",http.StatusSeeOther)
		return
	}
	fname:=r.FormValue("fname")
	d := struct{
		First string
	}{
		First: fname,
	}
	tpl.ExecuteTemplate(w,"processor.gohtml",d)
}
func main(){
	http.HandleFunc("/",index_handler)
	http.HandleFunc("/about/",about_h)
	http.HandleFunc("/form/",form_handler)
	http.HandleFunc("/process/",processor)
	http.ListenAndServe(":8000",nil)
}