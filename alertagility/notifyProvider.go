package main

import (
        "bytes"
        "fmt"
        "net/http"
        "encoding/json"
//      "net/smtp"
        "github.com/mailgun/mailgun-go"
        log "github.com/sirupsen/logrus"
)


func main() {
        http.HandleFunc("/", handleRequest)
        http.ListenAndServe(":9192", nil)
}


type Alert struct {
        AlertId string `json:"alert_id"`
        Description string `json:"description"`
        Details string `json:"details"`
        Host string `json:"host"`
        Name string `json:"name"`
        Service string `json:"service"`
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
        fmt.Println("Notify provider email got called.")
        decoder := json.NewDecoder(r.Body)
        var alert Alert
        err := decoder.Decode(&alert)
        if err != nil {
                fmt.Println("Error decoding alert ... ", err)
        }

        fmt.Println(alert)

        subject := alert.Details +  " " + alert.Description

        if contains(subject, "API") {
                sendEmail("bibek.shing@prabhubank.com", "Subisu link down", "Subject: "+subject)
        } else if contains(subject, "worldlink") {
                sendEmail("ramesh.ghale@prabhubank.com", "Worldlink link down", "Subject: "+subject)
        }

        fmt.Fprintln(w, "Email sent!")
}


func contains(str, substr string) bool {
        return bytes.Contains([]byte(str), []byte(substr))
}


func sendEmail(to, subject, body string) error {
        mg := mailgun.NewMailgun("mg.priceblick.com", "")
        //mg.SetAPIBase("https://api.eu.mailgun.net/v3")
        from := "info@priceblick.com"
        if to == "" {
            to = "sulochan@gmail.com"
        }

        msg := []byte("To: " + to + "\r\n" +
                "Subject: " + subject + "\r\n" +
                "\r\n" +
                body + "\r\n")

        message := mg.NewMessage(from, subject, string(msg), to)


        log.Info("@@@@@ SENDING EMAIL with mg.inithive.com @@@@@")
        if msg, id, err := mg.Send(message); err != nil {
                log.Error(msg, id, err)
                return err
        }
        return nil
}
