package main

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"net/smtp"
)

// Requires link used to fetch IP and IP address that will be returned in case of an error
func getIP(link string, previousIP net.IP) net.IP {

	res, err := http.Get(link)

	if err != nil {
		return previousIP
	}

	resBody, err := io.ReadAll(res.Body)

	if err != nil {
		return previousIP
	}

	return resBody
}

// Returns true if IP address has changed
func compareIP(new_ip net.IP, last_ip net.IP) bool {
	return new_ip.Equal(last_ip)
}

func mail(from string, to string, password string, newIP string, previousIP string, timeStamp string, previousTime string) bool {
	const smtpServer = "smtp.gmail.com:587"
	smtpAuth := smtp.PlainAuth("", from, password, "smtp.gmail.com")

	msg := "From: " + from + "\r\n" +
		"To: " + to + "\r\n" +
		"Subject:❗ IP address change time has changed!\r\n\r\n" +
		"Previous change time was:  " + previousTime + "\r\n" +
		"Previous IP address: " + string(previousIP[:]) + "\r\n" +
		"New IP address: " + string(newIP[:]) + "\r\n" +
		"Time stamp: " + timeStamp

	err := smtp.SendMail(smtpServer, smtpAuth, from, []string{to}, []byte(msg))

	if err != nil {
		fmt.Println(err)
		return false
	}

	return true
}
