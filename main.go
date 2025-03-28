package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("domain, hasMX, hasSPF, sprRecord, hasDMARC, dmarcRecord\n")

	for scanner.Scan() {
		checkDomain(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error could not read from input: %v\n", err)
	}
}

func checkDomain(domain string) {
	var hasMX, hasDMARC, hasSPF bool
	var dmarcRecord, spfRecord string

	mxRecord, err := net.LookupMX(domain)
	if err != nil {
		log.Fatalf("Could not lookup MX record: %v", err)
	}

	if len(mxRecord) > 0 {
		hasMX = true
	}

	txtRecord, err := net.LookupTXT(domain)
	if err != nil {
		log.Fatalf("Could not lookup TXT record: %v", err)
	}

	for _, record := range txtRecord {
		if strings.HasPrefix(record, "v=spf1") {
			hasSPF = true
			spfRecord = record
			break
		}
	}

	dRecord, err := net.LookupTXT("_dmarc." + domain)
	if err != nil {
		log.Fatalf("Could not lookup Dmarc record: %v", err)
	}

	for _, record := range dRecord {
		if strings.HasPrefix(record, "v=DMARC1") {
			hasDMARC = true
			dmarcRecord = record
			break
		}
	}
	fmt.Printf("%v, %v, %v, %v %v, %v", domain, hasMX, hasSPF, spfRecord, hasDMARC, dmarcRecord)

}
