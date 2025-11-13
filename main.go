package main

import (
	"fmt" //Imprime texto formateado
	"net/http" //Maneja el servidor HTTP y peticiones
)

func handler(w http.ResponseWriter, r *http.Request) { //handler escribe la respuesta HTTP cuando se accede a "/"
	fmt.Fprintln(w, "<h1>Soy alumno/a de la UOC</h1><img src='/static/imagen.jpg'>") 
	//Fprintln escribe la cadena en el cuerpo de la respuesta (incluye salto de línea)
	// w es el ResponseWriter usado para enviar la respuesta al cliente
	// r contiene la información de la solicitud HTTP entrante
}

func main() { 
	http.HandleFunc("/", handler) //Asocia la ruta raíz "/" con la función handler
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	// Sirve archivos estáticos desde el directorio local "./static"
	// Las peticiones a "/static/..." se redirigen al FileServer quitando el prefijo "/static/"
	http.ListenAndServe(":8080", nil)
	// Inicia el servidor HTTP en el puerto 8080; el segundo parámetro nil usa el DefaultServeMux
	// Si ListenAndServe devuelve error (por ejemplo puerto ocupado), lo ignoramos aquí — en producción conviene manejarlo.
}
