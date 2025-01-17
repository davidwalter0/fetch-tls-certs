package main

import (
	"bytes"
	"encoding/pem"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/davidwalter0/fetchhostcerts"
)

var version = ""

func main() {
	var format string
	var template string
	var skipVerify bool
	var utc bool
	var timeout int
	var showVersion bool

	flag.StringVar(&format, "f", "simple table", "Output format. md: as markdown, json: as JSON. ")
	flag.StringVar(&format, "format", "simple table", "Output format. md: as markdown, json: as JSON. ")
	flag.StringVar(&template, "t", "", "Output format as Go template string or Go template file path.")
	flag.StringVar(&template, "template", "", "Output format as Go template string or Go template file path.")
	flag.BoolVar(&skipVerify, "k", false, "Skip verification of server's certificate chain and host name.")
	flag.BoolVar(&skipVerify, "skip-verify", false, "Skip verification of server's certificate chain and host name.")
	flag.BoolVar(&utc, "u", false, "Use UTC to represent NotBefore and NotAfter.")
	flag.BoolVar(&utc, "utc", false, "Use UTC to represent NotBefore and NotAfter.")
	flag.IntVar(&timeout, "s", 3, "Timeout seconds.")
	flag.IntVar(&timeout, "timeout", 3, "Timeout seconds.")
	flag.BoolVar(&showVersion, "v", false, "Show version.")
	flag.BoolVar(&showVersion, "version", false, "Show version.")
	flag.Parse()

	if showVersion {
		fmt.Println("cert version ", version)
		return
	}

	var certs fetchhostcerts.Certs
	var err error

	fetchhostcerts.SkipVerify = skipVerify
	fetchhostcerts.UTC = utc
	fetchhostcerts.TimeoutSeconds = timeout

	certs, err = fetchhostcerts.NewCerts(flag.Args())
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
	for _, certficate := range certs {
		var pemBytes bytes.Buffer
		for _, cert := range certficate.CertChain() {
			if err := pem.Encode(&pemBytes, &pem.Block{Type: "CERTIFICATE", Bytes: cert.Raw}); err != nil {
				if err != nil {
					log.Println("CertChainToPEM", err)
					continue
				}
			}
			f, err := ioutil.TempFile(".", certficate.DomainName+".*")
			if err != nil {
				log.Println("Tempfile", err)
				continue
			}
			if n, err := f.Write(pemBytes.Bytes()); err != nil {
				log.Printf("Wrote %d of %d Write", n, pemBytes.Len(), err)
			}
		}
	}

	if template == "" {
		switch format {
		case "md":
			fmt.Printf("%s", certs.Markdown())
		case "json":
			fmt.Printf("%s", certs.JSON())
		default:
			fmt.Printf("%s", certs)
		}
		return
	}

	if err := fetchhostcerts.SetUserTempl(template); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	fmt.Printf("%s", certs)
}
