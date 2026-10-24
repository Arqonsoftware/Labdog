package main

import (
	"flag"
	"log"
	"os/exec"
	"time"
)

func reachable(addr string) bool {
	return exec.Command("ping", "-c1", "-W2", addr).Run() == nil
}

// runWatchdog shuts the host down when a peer that shares its power supply
// (e.g. a Raspberry Pi with no UPS) stops answering pings.
func runWatchdog(args []string) int {
	fs := flag.NewFlagSet("watchdog", flag.ExitOnError)
	peer := fs.String("peer", "", "address of the power-loss sensor (required)")
	failAfter := fs.Duration("fail-after", 60*time.Second, "how long the peer must be unreachable before shutting down")
	interval := fs.Duration("interval", 10*time.Second, "ping interval")
	dryRun := fs.Bool("dry-run", false, "log instead of shutting down")
	fs.Parse(args)

	if *peer == "" {
		log.Print("watchdog: -peer is required")
		return 2
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
					return 0
				}
				if err := exec.Command("shutdown", "-h", "now", "peer unreachable: power outage").Run(); err != nil {
					log.Printf("shutdown failed: %v", err)
					return 1
				}
				return 0
			}
		}
		time.Sleep(*interval)
	}
}
