package main

import (
	"fmt"
	"time"

	"github.com/bnb-chain/tss-lib/ecdsa/keygen"
	"github.com/bnb-chain/tss-lib/ecdsa/resharing"
)

func formatDuration(d time.Duration) string {
	seconds := int(d.Seconds())
	milliseconds := int((d - time.Duration(seconds)*time.Second) / time.Millisecond)
	return fmt.Sprintf("%d seconds %d milliseconds", seconds, milliseconds)
}

func main() {
	fmt.Println("Running tests...")

	// Measure elapsed time for keygen functions
	startKeygen := time.Now()

	keygen.StartRound1Paillier()
	keygen.FinishAndSaveH1H2()
	keygen.E2EConcurrentAndSaveFixtures()

	elapsedKeygen := time.Since(startKeygen)

	// Measure elapsed time for resharing function
	startResharing := time.Now()

	resharing.ResharingProtocol()

	elapsedResharing := time.Since(startResharing)

	// Calculate the total elapsed time
	elapsedTotal := elapsedKeygen + elapsedResharing

	// Output the elapsed times
	fmt.Printf("Elapsed time for keygen: %s\n", formatDuration(elapsedKeygen))
	fmt.Printf("Elapsed time for resharing: %s\n", formatDuration(elapsedResharing))
	fmt.Printf("Total elapsed time: %s\n", formatDuration(elapsedTotal))
}
