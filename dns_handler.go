package main

import (
	"fmt"
	"net"
	"sync"

	"github.com/asalih/guardian_ns/data"
	"github.com/miekg/dns"
)

//DNSHandler the dns handler
type DNSHandler struct {
	Targets  map[string]string
	DBHelper *data.DNSDBHelper

	mutex sync.Mutex
}

//NewDNSHandler Init dns handler
func NewDNSHandler() *DNSHandler {
	handler := &DNSHandler{nil, &data.DNSDBHelper{}, sync.Mutex{}}

	handler.LoadTargets()

	return handler
}

//ServeDNS ...
func (h *DNSHandler) ServeDNS(w dns.ResponseWriter, r *dns.Msg) {
	msg := dns.Msg{}
	msg.SetReply(r)

	switch r.Question[0].Qtype {
	case dns.TypeA:
		msg.Authoritative = true
		domain := msg.Question[0].Name
		address, ok := h.Targets[domain]

		fmt.Println("Requested: " + domain)
		if ok {
			msg.Answer = append(msg.Answer, &dns.A{
				Hdr: dns.RR_Header{Name: domain, Rrtype: dns.TypeA, Class: dns.ClassINET, Ttl: 60},
				A:   net.ParseIP(address),
			})
		}
	}
	w.WriteMsg(&msg)
}

//LoadTargets ...
func (h *DNSHandler) LoadTargets() {
	h.mutex.Lock()
	defer h.mutex.Unlock()

	h.Targets = h.DBHelper.GetTargetsList()
	h.Targets["ntp.ubuntu.com."] = "91.189.91.157"
}
