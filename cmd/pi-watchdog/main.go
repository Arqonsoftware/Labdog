// pi-watchdog shuts the host down gracefully when a peer that shares its
// power supply (e.g. a Raspberry Pi with no UPS) stops answering pings.
package main

import (
	"flag"
	"log"
	"os/exec"
	"time"
)

var version = "dev"

func reachable(addr string) bool {
	return exec.Command("ping", "-c1", "-W2", addr).Run() == nil
}

func main() {
	peer := flag.String("peer", "", "address of the power-loss sensor (required)")
	failAfter := flag.Duration("fail-after", 60*time.Second, "how long the peer must be unreachable before shutting down")
	interval := flag.Duration("interval", 10*time.Second, "ping interval")
	dryRun := flag.Bool("dry-run", false, "log instead of shutting down")
	showVersion := flag.Bool("version", false, "print version and exit")
	flag.Parse()

	if *showVersion {
		log.Println(version)
		return
	}
	if *peer == "" {
		log.Fatal("-peer is required")
	}

	// Armed only after the peer has been seen up, so a peer that is down
	// at boot cannot cause a shutdown loop.
	armed := false
	var downSince time.Time
	for {
		switch {
		case reachable(*peer):
			if !armed {
				log.Printf("peer %s seen up; armed", *peer)
			}
			armed, downSince = true, time.Time{}
		case armed:
			if downSince.IsZero() {
				downSince = time.Now()
				log.Printf("peer %s unreachable", *peer)
			}
			if time.Since(downSince) >= *failAfter {
				log.Printf("peer down for %s: assuming power outage, shutting down", *failAfter)
				if *dryRun {
					return
				}
				if err := exec.Command("shutdown", "-h", "now", "peer unreachable: power outage").Run(); err != nil {
					log.Fatalf("shutdown failed: %v", err)
				}
				return
			}
		}
		time.Sleep(*interval)
	}
}
