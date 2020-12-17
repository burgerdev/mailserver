package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	vultr "github.com/JamesClonk/vultr/lib"
	"github.com/miekg/dns"
)

func mandatoryEnvVar(name string) string {
	value := os.Getenv(name)
	if value == "" {
		log.Fatalf("environment variable %q must be set", name)
	}
	return value
}

var apiKey = mandatoryEnvVar("VULTR_API_KEY")
var parentDomain = mandatoryEnvVar("VULTR_PARENT_DOMAIN")
var tsigSecret = mandatoryEnvVar("TSIG_SECRET")
var tsigSecretName = mandatoryEnvVar("TSIG_SECRET_NAME")

func serve(net string) {
	t := map[string]string{tsigSecretName: tsigSecret}
	server := &dns.Server{Addr: ":8053", Net: net, TsigSecret: t, ReusePort: true}
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("stopped serving dns: %v", err)
	}
}

func main() {
	handler := &handler{vultr.NewClient(apiKey, nil)}
	dns.Handle(parentDomain, handler)
	dns.HandleFunc(".", func(w dns.ResponseWriter, r *dns.Msg) {
		log.Printf("Refusing to handle:\n\tdns.Msg%+v", *r)
		m := new(dns.Msg)
		m.SetReply(r)
		m.Rcode = dns.RcodeRefused
		w.WriteMsg(m)
	})

	go serve("tcp")
	go serve("udp")
	sig := make(chan os.Signal)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	s := <-sig
	log.Printf("Signal (%s) received, stopping\n", s)
}
