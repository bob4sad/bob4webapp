package main

import (
    "fmt"
    "log"
    "net/http"
    "crypto/tls"
    gomail "gopkg.in/mail.v2"
    "github.com/joho/godotenv"
    "os"
)

import "flag" //DEBUG

func sendEmailData(data string) {
    m := gomail.NewMessage()
    m.SetHeader("To", os.Getenv("TO_EMAIL"))
    m.SetHeader("From", os.Getenv("FROM_EMAIL"))
    m.SetHeader("Subject", "New customer!")
    m.SetBody("text/plain", data)
    d := gomail.NewDialer("smtp.gmail.com", 587, os.Getenv("SENDER_EMAIL_USERNAME"), os.Getenv("SENDER_EMAIL_PASSWORD"))
    d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
    if err := d.DialAndSend(m); err != nil {
      fmt.Println(err)
      panic(err)
    }

    return
}


func formHandler(w http.ResponseWriter, r *http.Request) {
    if err := r.ParseForm(); err != nil {
        fmt.Fprintf(w, "ParseForm() err: %v", err)
        return
    }
    city := r.FormValue("city")
    phone := r.FormValue("phone")
    data := "City = " + city + "\n" + "Phone = " + phone
    sendEmailData(data)
    http.Redirect(w, r, "#success", 302)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/" {
        http.Error(w, "404 not found.", http.StatusNotFound)
        return
    }

    if r.Method != "GET" {
        http.Error(w, "Method is not supported.", http.StatusNotFound)
        return
    }


    fmt.Fprint(w, "Hello!")
}

func main() {

    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }



    //DEBUG
    useDegug := flag.Bool("debug", false, "use debug mode (./static_test)") 
    flag.Parse()

    var pathStatic string
    if *useDegug {
        pathStatic = "./static_test"
    } else {
        pathStatic = "./static"
    }
    //DEBUG

    
    fileServer := http.FileServer(http.Dir(pathStatic))
    http.Handle("/", fileServer)
    http.HandleFunc("/form", formHandler)

    fmt.Printf("Starting server at port 8080\n")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal(err)
    }
}