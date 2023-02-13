package main

import (
	"context"
	"fmt"
	_ "goservices/internal/store"
	"goservices/server"
	store "goservices/store"
	"goservices/store/factory"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World! %s", time.Now().Format("2006/01/02 15:04:05"))
	fmt.Println("hello, world")
	var book store.Book
	book.Name = "Vue3第一课时"
	book.Id = "aehyok"
	fmt.Println(book.Name)
	fmt.Println(book.Id)
}

func greets(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World! %s", time.Now().Format("2006/01/02 15:04:05"))
	fmt.Println("hello, world-get")
}
func main() {
	s, err := factory.New("mem")
	if err != nil {
		panic(err)
	}

	srv := server.NewBookStoreServer(":8080", s)

	errChan, err := srv.ListenAndServe()
	if err != nil {
		log.Println("web server start failed:", err)
		return
	}
	log.Println("web server start ok")

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)

	select {
	case err = <-errChan:
		log.Println("web server run failed:", err)
		return
	case <-c:
		log.Println("bookstore program is exiting...")
		ctx, cf := context.WithTimeout(context.Background(), time.Second)
		defer cf()
		err = srv.Shutdown(ctx)
	}

	if err != nil {
		log.Println("bookstore program exit error:", err)
		return
	}
	log.Println("bookstore program exit ok")
}
