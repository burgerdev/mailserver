package main

import (
	"log"
	"os"

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

func registerDNSHandler(v *vultr.Client) {
	domains, err := v.GetDNSDomains()
	if err != nil {
		log.Fatalf("could not list DNS domains: %v", err)
	}
	for _, d := range domains {
		log.Printf("%+v", d)
	}
	dns.HandleFunc("burgerdev.de.", func(w dns.ResponseWriter, r *dns.Msg) {})
}

func main() {
	handler := &handler{vultr.NewClient(apiKey, nil)}
	dns.Handle(parentDomain, handler)

	t := map[string]string{"burgerdev-de-secret.": tsigSecret}
	server := &dns.Server{Addr: ":8053", Net: "udp", TsigSecret: t, ReusePort: true}
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("stopped serving dns: %v", err)
	}
}
