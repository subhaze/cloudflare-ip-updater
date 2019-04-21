package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"

	cloudflare "github.com/cloudflare/cloudflare-go"
)

func notFound(res http.ResponseWriter, err error) {
	fmt.Println(err)
	res.WriteHeader(404)
	res.Write([]byte(err.Error()))
}

func missingQueryParam(param string, value string, res http.ResponseWriter) bool {
	if value == "" {
		notFound(res, errors.New("Query param '"+param+"' was not found."))
		return true
	}
	return false
}

func handleIPUpdate(res http.ResponseWriter, req *http.Request) {
	site := req.URL.Query().Get("site")
	domain := req.URL.Query().Get("domain")
	ip := req.URL.Query().Get("ip")
	if missingQueryParam("site", site, res) || missingQueryParam("domain", domain, res) || missingQueryParam("ip", ip, res) {
		return
	}
	updateIP(site, domain, ip, res)
}

func updateIP(site string, domain string, ip string, res http.ResponseWriter) {
	fmt.Println(os.Getenv("CF_API_EMAIL"))
	api, err := cloudflare.New(os.Getenv("CF_API_KEY"), os.Getenv("CF_API_EMAIL"))
	if err != nil {
		notFound(res, err)
		return
	}

	id, err := api.ZoneIDByName(site) // Assuming example.com exists in your Cloudflare account already
	if err != nil {
		notFound(res, err)
		return
	}
	zone, err := api.ZoneDetails(id)
	if err != nil {
		notFound(res, err)
		return
	}
	dnsRecords, err := api.DNSRecords(zone.ID, cloudflare.DNSRecord{Name: domain})
	if err != nil || len(dnsRecords) < 1 {
		notFound(res, err)
		return
	}
	dnsRecord := dnsRecords[0]
	fmt.Printf("[info] Attempting to update %s IP %s\n", domain, dnsRecord.Content)
	dnsRecord.Content = ip
	err = api.UpdateDNSRecord(zone.ID, dnsRecord.ID, dnsRecord)
	if err != nil {
		fmt.Printf("[error] %s", err.Error())
		notFound(res, err)
		return
	}
	fmt.Printf("[info] Success! Updated %s IP to %s\n", domain, dnsRecord.Content)
	res.WriteHeader(200)
	res.Write([]byte("success!"))
}

func main() {
	http.HandleFunc("/cf-ip-update", handleIPUpdate)
	http.ListenAndServe(":8080", nil)
}
