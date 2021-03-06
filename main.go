package main

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/asalih/guardian_ns/models"
	"github.com/miekg/dns"
)

var _dnsHandler *DNSHandler

func main() {
	models.InitConfig()

	_dnsHandler = NewDNSHandler()

	ticker := time.NewTicker(2 * time.Minute)
	defer ticker.Stop()

	go func() {
		for ; true; <-ticker.C {
			fmt.Println("Tick at", time.Now())
			_dnsHandler.LoadTargets()
			fmt.Println("Targets loaded: " + strconv.Itoa(len(_dnsHandler.Targets)))
		}
	}()

	srv := &dns.Server{Addr: ":" + strconv.Itoa(53), Net: "udp"}
	srv.Handler = _dnsHandler
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("Failed to set udp listener %s\n", err.Error())
	}
}
