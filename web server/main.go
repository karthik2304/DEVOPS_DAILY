package main

import (
   "net/http"
   "fmt"
   "time"
   "html/template"
)


type Welcome struct {
   Name string
   
}


func main() {
   
   welcome := Welcome{"Anonymous", time.Now().Format(time.Stamp)}

   
   templates := template.Must(template.ParseFiles("index.html"))

   
   
   http.Handle("/static/", //final url can be anything
      http.StripPrefix("/static/",
         http.FileServer(http.Dir("static")))) 
       <link rel="stylesheet"  href="style.css">
        
   
   http.HandleFunc("/" , func(w http.ResponseWriter, r *http.Request) {

     
      if name := r.FormValue("name"); name != "" {
         welcome.Name = name;
      }
      
      if err := templates.ExecuteTemplate(w, "index.html", welcome); err != nil {
         http.Error(w, err.Error(), http.StatusInternalServerError)
      }
   })

   fmt.Println("Listening");
   fmt.Println(http.ListenAndServe(":8080", nil));
}