package main

import (
	"testing"

	"github.com/bnb-chain/tss-lib/ecdsa/keygen"
	"github.com/bnb-chain/tss-lib/ecdsa/resharing"
)

func runKeygen() {
	keygen.StartRound1Paillier()
	keygen.FinishAndSaveH1H2()
	keygen.E2EConcurrentAndSaveFixtures()
}

func runResharing() {
	resharing.ResharingProtocol()
}

func BenchmarkKeygen(b *testing.B) {
	for i := 0; i < b.N; i++ {
		runKeygen()
	}
}

func BenchmarkResharing(b *testing.B) {
	for i := 0; i < b.N; i++ {
		runResharing()
	}
}

func BenchmarkTotal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		runKeygen()
		runResharing()
	}
}

func BenchmarkCombined(b *testing.B) {
	for i := 0; i < b.N; i++ {
		runKeygen()
		runResharing()
	}
}
