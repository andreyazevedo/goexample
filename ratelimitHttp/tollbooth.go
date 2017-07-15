/*
* Example tollbooth
*
* @package     main
* @author      @jeffotoni
* @size        15/07/2017
*
 */

package main

import (
	"fmt"
	"github.com/didip/tollbooth"
	"net/http"
	"time"
)

func Hello(w http.ResponseWriter, req *http.Request) {

	w.Write([]byte("Hello, Jefferson !"))
}

func Login(w http.ResponseWriter, req *http.Request) {

	w.Write([]byte(`{"valid":"login"}`))
}

func main() {

	port := "12345"

	// Create a request limiter per handler.
	http.Handle("/", tollbooth.LimitFuncHandler(tollbooth.NewLimiter(900, time.Millisecond), Hello))

	// Create a request limiter per handler.
	http.Handle("/login", tollbooth.LimitFuncHandler(tollbooth.NewLimiter(2, time.Second), Login))

	fmt.Println("Start port:", port)
	fmt.Println("Endpoints:")
	fmt.Println("http://localhost:" + port + "/")
	fmt.Println("http://localhost:" + port + "/login")

	s := &http.Server{

		Addr: ":" + port,
		// Handler:        myHandler,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()

}