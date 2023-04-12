Multi-Party Threshold ECDSA and EdDSA Library
This repository contains an implementation of multi-party {t,n}-threshold ECDSA (Elliptic Curve Digital Signature Algorithm) based on Gennaro and Goldfeder CCS 2018 1 and EdDSA (Edwards-curve Digital Signature Algorithm) following a similar approach. The library is permissively MIT Licensed.

Overview
The library includes three protocols:

Key Generation for creating secret shares with no trusted dealer ("keygen").
Signing for using the secret shares to generate a signature ("signing").
Dynamic Groups to change the group of participants while keeping the secret ("resharing").
The implementation was created by rewriting tests and following the instructions from the README.md included in the repository and documentation available for bnb-chain/tss-lib by Binance.
