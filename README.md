Multi-Party Threshold ECDSA and EdDSA Library

This repository contains an implementation of multi-party {t,n}-threshold ECDSA (Elliptic Curve Digital Signature Algorithm) based on Gennaro and Goldfeder CCS 2018 1 and EdDSA (Edwards-curve Digital Signature Algorithm) following a similar approach. The library is permissively MIT Licensed.

Overview
The library includes three protocols:

Key Generation for creating secret shares with no trusted dealer ("keygen").
Signing for using the secret shares to generate a signature ("signing").
Dynamic Groups to change the group of participants while keeping the secret ("resharing").

The benchmarks were created by rewriting tests and following the instructions from the README.md included in the repository and documentation available for bnb-chain/tss-lib (https://github.com/bnb-chain/tss-lib) by Binance.

How to run it?

Directory tss-lib-benchmark-master\ecdsa\ecdsa-benchmark contains the benchmark which can be run by using either default benchmarking Go commands or simply by running the functions in benchmark-test.go

Number of parties and threshold needs to be edited manually in tss-lib-benchmark-master\test\config.go but the text fixture files in test/_fixtures/ must be deleted first.
