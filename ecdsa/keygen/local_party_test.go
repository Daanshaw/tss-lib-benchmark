package keygen

//const (
//	testParticipants = TestParticipants
//	testThreshold    = TestThreshold
//)

//func setUp(level string) {
//	if err := log.SetLogLevel("tss-lib", level); err != nil {
//		panic(err)
//	}
//}
//
//func TestStartRound1Paillier(t *testing.T) {
//	setUp("debug")
//
//	pIDs := tss.GenerateTestPartyIDs(1)
//	p2pCtx := tss.NewPeerContext(pIDs)
//	threshold := 1
//	params := tss.NewParameters(tss.EC(), p2pCtx, pIDs[0], len(pIDs), threshold)
//
//	fixtures, pIDs, err := LoadKeygenTestFixtures(testParticipants)
//	if err != nil {
//		common.Logger.Info("No test fixtures were found, so the safe primes will be generated from scratch. This may take a while...")
//		pIDs = tss.GenerateTestPartyIDs(testParticipants)
//	}
//
//	var lp *LocalParty
//	out := make(chan tss.Message, len(pIDs))
//	if 0 < len(fixtures) {
//		lp = NewLocalParty(params, out, nil, fixtures[0].LocalPreParams).(*LocalParty)
//	} else {
//		lp = NewLocalParty(params, out, nil).(*LocalParty)
//	}
//	if err := lp.Start(); err != nil {
//		common.Logger.Error("Failed to start: ", err)
//		return
//	}
//	<-out
//
//	// Paillier modulus 2048 (two 1024-bit primes)
//	// round up to 256, it was used to be flaky, sometimes comes back with 1 byte less
//	len1 := len(lp.data.PaillierSK.LambdaN.Bytes())
//	len2 := len(lp.data.PaillierSK.PublicKey.N.Bytes())
//	if len1%2 != 0 {
//		len1 = len1 + (256 - (len1 % 256))
//	}
//	if len2%2 != 0 {
//		len2 = len2 + (256 - (len2 % 256))
//	}
//	if len1 != 2048/8 || len2 != 2048/8 {
//		common.Logger.Error("Paillier modulus lengths do not match expected values")
//	}
//}
//
//func TestFinishAndSaveH1H2(t *testing.T) {
//	setUp("debug")
//
//	pIDs := tss.GenerateTestPartyIDs(1)
//	p2pCtx := tss.NewPeerContext(pIDs)
//	threshold := 1
//	params := tss.NewParameters(tss.EC(), p2pCtx, pIDs[0], len(pIDs), threshold)
//
//	fixtures, pIDs, err := LoadKeygenTestFixtures(testParticipants)
//	if err != nil {
//		common.Logger.Info("No test fixtures were found, so the safe primes will be generated from scratch. This may take a while...")
//		pIDs = tss.GenerateTestPartyIDs(testParticipants)
//	}
//
//	var lp *LocalParty
//	out := make(chan tss.Message, len(pIDs))
//	if 0 < len(fixtures) {
//		lp = NewLocalParty(params, out, nil, fixtures[0].LocalPreParams).(*LocalParty)
//	} else {
//		lp = NewLocalParty(params, out, nil).(*LocalParty)
//	}
//	if err := lp.Start(); err != nil {
//		common.Logger.Error("Failed to start: ", err)
//		return
//	}
//
//	// RSA modulus 2048 (two 1024-bit primes)
//	// round up to 256
//	len1 := len(lp.data.H1j[0].Bytes())
//	len2 := len(lp.data.H2j[0].Bytes())
//	len3 := len(lp.data.NTildej[0].Bytes())
//	if len1%2 != 0 {
//		len1 = len1 + (256 - (len1 % 256))
//	}
//	if len2%2 != 0 {
//		len2 = len2 + (256 - (len2 % 256))
//	}
//	if len3%2 != 0 {
//		len3 = len3 + (256 - (len3 % 256))
//	}
//
//	if len1 != 256 || len2 != 256 || len3 != 256 {
//		common.Logger.Error("h1, h2, and n-tilde lengths do not match expected values")
//	}
//	if lp.data.H1i.Cmp(big.NewInt(0)) == 0 || lp.data.H2i.Cmp(big.NewInt(0)) == 0 || lp.data.NTildei.Cmp(big.NewInt(0)) == 0 {
//		common.Logger.Error("h1, h2, and n-tilde should be non-zero")
//	}
//}
//
//func TestBadMessageCulprits(t *testing.T) {
//	setUp("debug")
//
//	pIDs := tss.GenerateTestPartyIDs(2)
//	p2pCtx := tss.NewPeerContext(pIDs)
//	params := tss.NewParameters(tss.S256(), p2pCtx, pIDs[0], len(pIDs), 1)
//
//	fixtures, pIDs, err := LoadKeygenTestFixtures(testParticipants)
//	if err != nil {
//		common.Logger.Info("No test fixtures were found, so the safe primes will be generated from scratch. This may take a while...")
//		pIDs = tss.GenerateTestPartyIDs(testParticipants)
//	}
//
//	var lp *LocalParty
//	out := make(chan tss.Message, len(pIDs))
//	if 0 < len(fixtures) {
//		lp = NewLocalParty(params, out, nil, fixtures[0].LocalPreParams).(*LocalParty)
//	} else {
//		lp = NewLocalParty(params, out, nil).(*LocalParty)
//	}
//	if err := lp.Start(); err != nil {
//		common.Logger.Error("Failed to start: ", err)
//		return
//	}
//
//	badMsg, _ := NewKGRound1Message(pIDs[1], zero, &paillier.PublicKey{N: zero}, zero, zero, zero, new(dlnproof.Proof), new(dlnproof.Proof))
//	ok, err2 := lp.Update(badMsg)
//	t.Log(err2)
//	if !ok {
//		common.Logger.Info("Message update failed")
//	}
//	if err2 != nil {
//		if len(err2.Culprits()) != 1 || !reflect.DeepEqual(pIDs[1], err2.Culprits()[0]) {
//			common.Logger.Error("Culprit count or identity mismatch")
//		}
//		expectedError := "task ecdsa-keygen, party {0,P[1]}, round 1, culprits [{1,2}]: message failed ValidateBasic: Type: binance.tsslib.ecdsa.keygen.KGRound1Message, From: {1,2}, To: all"
//		if err2.Error() != expectedError {
//			common.Logger.Error("Error message mismatch")
//		}
//	}
//}
//
//func TestE2EConcurrentAndSaveFixtures(t *testing.T) {
//	setUp("info")
//
//	threshold := testThreshold
//	fixtures, pIDs, err := LoadKeygenTestFixtures(testParticipants)
//	if err != nil {
//		common.Logger.Info("No test fixtures were found, so the safe primes will be generated from scratch. This may take a while...")
//		pIDs = tss.GenerateTestPartyIDs(testParticipants)
//	}
//
//	p2pCtx := tss.NewPeerContext(pIDs)
//	parties := make([]*LocalParty, 0, len(pIDs))
//
//	errCh := make(chan *tss.Error, len(pIDs))
//	outCh := make(chan tss.Message, len(pIDs))
//	endCh := make(chan LocalPartySaveData, len(pIDs))
//
//	updater := test.SharedPartyUpdater
//
//	//startGR := runtime.NumGoroutine()
//
//	// init the parties
//	for i := 0; i < len(pIDs); i++ {
//		var P *LocalParty
//		params := tss.NewParameters(tss.S256(), p2pCtx, pIDs[i], len(pIDs), threshold)
//		if i < len(fixtures) {
//			P = NewLocalParty(params, outCh, endCh, fixtures[i].LocalPreParams).(*LocalParty)
//		} else {
//			P = NewLocalParty(params, outCh, endCh).(*LocalParty)
//		}
//		parties = append(parties, P)
//		go func(P *LocalParty) {
//			if err := P.Start(); err != nil {
//				errCh <- err
//			}
//		}(P)
//	}
//
//	// PHASE: keygen
//	var ended int32
//keygen:
//	for {
//		//fmt.Printf("ACTIVE GOROUTINES: %d\n", runtime.NumGoroutine())
//		select {
//		case err := <-errCh:
//			common.Logger.Errorf("Error: %s", err)
//			break keygen
//
//		case msg := <-outCh:
//			dest := msg.GetTo()
//			if dest == nil { // broadcast!
//				for _, P := range parties {
//					if P.PartyID().Index == msg.GetFrom().Index {
//						continue
//					}
//					go updater(P, msg, errCh)
//				}
//			} else { // point-to-point!
//				if dest[0].Index == msg.GetFrom().Index {
//					common.Logger.Errorf("party %d tried to send a message to itself (%d)", dest[0].Index, msg.GetFrom().Index)
//					return
//				}
//				go updater(parties[dest[0].Index], msg, errCh)
//			}
//
//		case save := <-endCh:
//			// SAVE a test fixture file for this P (if it doesn't already exist)
//			// .. here comes a workaround to recover this party's index (it was removed from save data)
//			index, err := save.OriginalIndex()
//			if err != nil {
//				common.Logger.Error("Error getting a party's index from save data")
//			} else {
//				tryWriteTestFixtureFile(index, save)
//			}
//
//			atomic.AddInt32(&ended, 1)
//			if atomic.LoadInt32(&ended) == int32(len(pIDs)) {
//				common.Logger.Infof("Done. Received save data from %d participants", ended)
//
//				// combine shares for each Pj to get u
//				u := new(big.Int)
//				for j, Pj := range parties {
//					pShares := make(vss.Shares, 0)
//					for _, P := range parties {
//						vssMsgs := P.temp.kgRound2Message1s
//						share := vssMsgs[j].Content().(*KGRound2Message1).Share
//						shareStruct := &vss.Share{
//							Threshold: threshold,
//							ID:        P.PartyID().KeyInt(),
//							Share:     new(big.Int).SetBytes(share),
//						}
//						pShares = append(pShares, shareStruct)
//					}
//					uj, err := pShares[:threshold+1].ReConstruct(tss.S256())
//					if err != nil {
//						common.Logger.Infof("Error while reconstructing shares: %v", err)
//						// Handle the error here as appropriate for your application
//					}
//
//					// uG test: u*G[j] == V[0]
//					if uj.Cmp(Pj.temp.ui) == 0 {
//						// The reconstructed value matches the expected value
//						// Do whatever you need to do here
//						common.Logger.Infof("Reconstructed value matches expected value: %v", uj)
//					}
//					//uG := crypto.ScalarBaseMult(tss.EC(), uj)
//					//assert.True(t, uG.Equals(Pj.temp.vs[0]), "ensure u*G[j] == V_0")
//
//					// xj tests: BigXj == xj*G
//					//xj := Pj.data.Xi
//					//gXj := crypto.ScalarBaseMult(tss.EC(), xj)
//					//BigXj := Pj.data.BigXj[j]
//					//assert.True(t, BigXj.Equals(gXj), "ensure BigX_j == g^x_j")
//
//					// fails if threshold cannot be satisfied (bad share)
//
//					// fails if threshold cannot be satisfied (bad share)
//					{
//						badShares := pShares[:threshold]
//						badShares[len(badShares)-1].Share.Set(big.NewInt(0))
//						uj, err := pShares[:threshold].ReConstruct(tss.S256())
//						//
//						//common.Logger.Infof("Error while reconstructing shares: %v", err)
//						if err != nil {
//							common.Logger.Error("An error occurred:", err)
//							// You can also handle the error as needed
//						}
//						// Handle the error here as appropriate for your application
//						if parties[j].temp.ui != (uj) {
//							common.Logger.Infof("Not equal")
//						}
//						// assert.NotEqual(t, parties[j].temp.ui, uj)
//						if parties[j].temp.ui != (uj) {
//							common.Logger.Infof("Not equal")
//						}
//						BigXjX, BigXjY := tss.EC().ScalarBaseMult(uj.Bytes())
//						// assert.NotEqual(t, BigXjX, Pj.temp.vs[0].X())
//						// assert.NotEqual(t, BigXjY, Pj.temp.vs[0].Y())
//						if BigXjX != (Pj.temp.vs[0].X()) {
//							common.Logger.Infof("Not equal")
//						}
//						if BigXjY != (Pj.temp.vs[0].Y()) {
//							common.Logger.Infof("Not equal")
//						}
//					}
//					u = new(big.Int).Add(u, uj)
//				}
//
//				// build ecdsa key pair
//				pkX, pkY := save.ECDSAPub.X(), save.ECDSAPub.Y()
//				pk := ecdsa.PublicKey{
//					Curve: tss.EC(),
//					X:     pkX,
//					Y:     pkY,
//				}
//				sk := ecdsa.PrivateKey{
//					PublicKey: pk,
//					D:         u,
//				}
//				// test pub key, should be on curve and match pkX, pkY
//				if !sk.IsOnCurve(pkX, pkY) {
//					common.Logger.Error("public key must be on curve")
//				}
//
//				// public key tests
//				if u.Cmp(big.NewInt(0)) == 0 {
//					common.Logger.Error("u should not be zero")
//				}
//				ourPkX, ourPkY := tss.EC().ScalarBaseMult(u.Bytes())
//				if pkX.Cmp(ourPkX) != 0 {
//					common.Logger.Error("pkX should match expected pk derived from u")
//				}
//				if pkY.Cmp(ourPkY) != 0 {
//					common.Logger.Error("pkY should match expected pk derived from u")
//				}
//				common.Logger.Info("Public key tests done.")
//
//				// make sure everyone has the same ECDSA public key
//				for _, Pj := range parties {
//					if pkX.Cmp(Pj.data.ECDSAPub.X()) != 0 || pkY.Cmp(Pj.data.ECDSAPub.Y()) != 0 {
//						common.Logger.Error("ECDSA public keys should match for all parties")
//					}
//				}
//				common.Logger.Info("Public key distribution test done.")
//
//				// test sign/verify
//				data := make([]byte, 32)
//				for i := range data {
//					data[i] = byte(i)
//				}
//				r, s, err := ecdsa.Sign(rand.Reader, &sk, data)
//				if err != nil {
//					common.Logger.Error("sign should not throw an error")
//				}
//				ok := ecdsa.Verify(&pk, data, r, s)
//				if !ok {
//					common.Logger.Error("signature should be ok")
//				}
//				common.Logger.Info("ECDSA signing testing done")
//
//				break keygen
//			}
//		}
//	}
//}
//
//func tryWriteTestFixtureFile(index int, data LocalPartySaveData) {
//	fixtureFileName := makeTestFixtureFilePath(index)
//
//	// fixture file does not already exist?
//	// if it does, we won't re-create it here
//	fi, err := os.Stat(fixtureFileName)
//	if !(err == nil && fi != nil && !fi.IsDir()) {
//		fd, err := os.OpenFile(fixtureFileName, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
//		if err != nil {
//			common.Logger.Errorf("unable to open fixture file %s for writing", fixtureFileName)
//		}
//		bz, err := json.Marshal(&data)
//		if err != nil {
//			common.Logger.Fatal("unable to marshal save data for fixture file %s", fixtureFileName)
//		}
//		_, err = fd.Write(bz)
//		if err != nil {
//			common.Logger.Fatal("unable to write to fixture file %s", fixtureFileName)
//		}
//		common.Logger.Infof("Saved a test fixture file for party %d: %s", index, fixtureFileName)
//	} else {
//		common.Logger.Infof("Fixture file already exists for party %d; not re-creating: %s", index, fixtureFileName)
//	}
//	//
//}
