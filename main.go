package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"time"
)

func main() {

	const url = "https://ipinfo.io/ip"
	const from = ""
	const password = ""

	startIP := net.ParseIP("0.0.0.0")
	ip := getIP(url, startIP)
	timeNow := time.Now()
	hour, minute := timeNow.Hour(), timeNow.Minute()

	previousIP := ip
	previousHour, previousMinute := hour, minute

	f, err := os.OpenFile("ip_tracker.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer f.Close()
	log.SetOutput(f)

	for {

		ip = getIP(url, startIP)
		timeNow = time.Now()
		hour, minute = timeNow.Hour(), timeNow.Minute()

		if compareIP(ip, previousIP) == false {

			str_IP := string(ip[:])
			str_previousIP := string(previousIP[:])
			str_TimeNow := timeNow.String()
			str_previousTime := fmt.Sprintf("%v:%v", strconv.FormatInt(int64(previousHour), 10), strconv.FormatInt(int64(previousMinute), 10))

			log.Printf("[INFO] %v New IP address: %v", str_TimeNow, str_IP) // Logs IP change

			if (previousHour != hour) || (previousMinute != minute) {
				log.Printf("[WARNING] %v Unexpected IP address change! Previous change time %v", str_TimeNow, str_previousTime)
				mail(from, from, password, str_IP, str_previousIP, str_TimeNow, str_previousTime)
			}

			previousHour = hour
			previousMinute = minute
		}

		time.Sleep(60 * time.Second)
	}
}
