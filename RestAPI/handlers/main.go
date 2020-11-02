package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"
	"../handlers/main"
	"os/signal"
	
)
 var bindAddress = env.String("BIND_ADDRESS" , false,":9090", "Bind address for the server")
func main() {
	env.Parse()

	l := log.New(os.Stdout, "gowrks", log.LstdFlags)

	hh := handlers.NewHello(l)
	gh :=handlers.NewGoodbye(l)
	

	
	sm := http.NewServeMux()
	sm.Handle("/", hh)
	sm.Handle("/goodbye", gh)
	

	s :=http.Server{
		Addr: *binAddress,
		Handler: sm,
		ErrorLog: l,
		ReadTimeout: 5 * time.Second,
       WriteTimeout: 10 * time.Second,
	   IdleTimeout: 120 * time.Second,

	}
	go func{
		l.Println("starting server on port 9090")
	err :=s.ListenAndServe()
	if err!=nil {
		l.Printf("error starting server :%s\n",err)
		os.Exit(1)
	}
	}()
	c :=make(chan os.Signal), 1
	signal.Notify(c,os.Interrupt)
	signal.Notify(c,os.Kill)

	sig := <- c
	l.Println("recived terminate,graceful shutdown",sig)
	tc,_ :=context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(tc)

}
