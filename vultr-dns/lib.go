package main

import (
	"log"
	"time"

	vultr "github.com/JamesClonk/vultr/lib"
	"github.com/miekg/dns"
)

type handler struct {
	client *vultr.Client
}

func (h *handler) doSomething() {
	dd, err := h.client.GetDNSDomains()
	if err != nil {
		log.Printf("could not obtain domains: %v", err)
		return
	}
	for _, d := range dd {
		log.Printf("DNSDomain%+v", d)
		records, err := h.client.GetDNSRecords(d.Domain)
		if err != nil {
			log.Printf("could not obtain records for domain %q: %v", d.Domain, err)
			continue
		}
		for _, r := range records {
			log.Printf("DNSRecord%+v", r)
		}
	}
}

func (h *handler) ServeDNS(w dns.ResponseWriter, r *dns.Msg) {
  log.Printf("new request from %s: dns.Msg%+v", w.RemoteAddr(), *r)

	tsig := r.IsTsig()
	if tsig == nil {
		log.Printf("request lacks TSIG")
		// unauth(w, r)
		// return
	} else {
		log.Printf("TSIG%+v", *tsig)
	}
	status := w.TsigStatus()
	if status != nil {
		log.Printf("ResponseWriter unexpectedly has a TSIG status: %v", status)
		fail(w, r)
		return
	}
	// m := new(dns.Msg)
	// m.SetReply(r)
	// m.SetTsig(tsig.Hdr.Name, tsig.Algorithm, 300, time.Now().Unix())
	// w.WriteMsg(m)

	doSomething()
	fail(w, r)
}

func fail(w dns.ResponseWriter, r *dns.Msg) {
	m := new(dns.Msg)
	m.SetReply(r)
	m.Rcode = dns.RcodeServerFailure
	w.WriteMsg(m)
}

func unauth(w dns.ResponseWriter, r *dns.Msg) {
	m := new(dns.Msg)
	m.SetReply(r)
	m.Rcode = dns.RcodeNotAuth
	w.WriteMsg(m)
}
