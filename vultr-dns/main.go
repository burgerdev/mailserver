package main

import (
	"log"
	"os"

	vultr "github.com/JamesClonk/vultr/lib"
	"github.com/miekg/dns"
)

const (
	envVarName = "VULTR_API_KEY"
)

var apiKey = os.Getenv(envVarName)

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
	if apiKey == "" {
		log.Fatalf("environment variable %q must be set", envVarName)
	}

	registerDNSHandler(vultr.NewClient(apiKey, nil))

	server := &dns.Server{Addr: ":8053", Net: "udp", TsigSecret: nil, ReusePort: true}
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("stopped serving dns: %v", err)
	}
}
