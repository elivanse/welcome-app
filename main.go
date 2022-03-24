package main

// 1. “net/http” to access the core go http functionality
// 2. “fmt” for formatting our text
// 3. “html/template” a library that allows us to interact with our html file.
// 4. "time" - a library for working with date and time.
import (
	"fmt"
	"html/template"
	"net/http"
	"time"
)

// creamos la estructura que   contiene la info que quweremos mostrar en el html
type Welcome struct {
	Name string
	Time string
}

//punto de entrada de la apliccion go
func main() {
	//instanciamos un objeto Welcome struct y le pasamos random info
	//deberemos obtener el nombre del usuario como un parametro desde el url
	welcome := Welcome{"Anonymous", time.Now().Format(time.Stamp)}

	//le decimos a go donde encontrar el html. le pedimos parsear el html en path relativo
	//lo enpaquetamos in la llamada template.Must() que chequea los errores y salta
	templates := template.Must(template.ParseFiles("template/welcome-template.html"))

	//Our HTML comes with CSS that go needs to provide when we run the app. Here we tell go to create
	// a handle that looks in the static directory, go then uses the "/static/" as a url that our
	//html can refer to when looking for our css and other files.

	http.Handle("/static/", //final url can be anything
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("static")))) //Go looks in the relative "static" directory first using http.FileServer(), then matches it to a
	//url of our choice as shown in http.Handle("/static/"). This url is what we need when referencing our css files
	//once the server begins. Our html code would therefore be <link rel="stylesheet"  href="/static/stylesheet/...">
	//It is important to note the url in http.Handle can be whatever we like, so long as we are consistent.

	//This method takes in the URL path "/" and a function that takes in a response writer, and a http request.
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		//Takes the name from the URL query e.g ?name=Martin, will set welcome.Name = Martin.
		if name := r.FormValue("name"); name != "" {
			welcome.Name = name
		}
		//If errors show an internal server error message
		//I also pass the welcome struct to the welcome-template.html file.
		if err := templates.ExecuteTemplate(w, "welcome-template.html", welcome); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	//Start the web server, set the port to listen to 8080. Without a path it assumes localhost
	//Print any errors from starting the webserver using fmt
	fmt.Println("Listening")
	fmt.Println(http.ListenAndServe(":8220", nil))

}
