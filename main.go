package main

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/asalih/guardian/models"
	"github.com/miekg/dns"
)

var _dnsHandler *DNSHandler

func main() {
	models.InitConfig()

	_dnsHandler = NewDNSHandler()

	ticker := time.NewTicker(3 * time.Minute)
	defer ticker.Stop()

	go func() {
		for ; true; <-ticker.C {
			fmt.Println("Tick at", time.Now())

		}
	}()

	srv := &dns.Server{Addr: ":" + strconv.Itoa(53), Net: "udp"}
	srv.Handler = _dnsHandler
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("Failed to set udp listener %s\n", err.Error())
	}
}
