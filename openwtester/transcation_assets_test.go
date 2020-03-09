/*
 * Copyright 2018 The openwallet Authors
 * This file is part of the openwallet library.
 *
 * The openwallet library is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Lesser General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * The openwallet library is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
 * GNU Lesser General Public License for more details.
 */

package openwtester

import (
	"github.com/blocktree/openwallet/openw"
	"testing"
	"time"

	"github.com/blocktree/openwallet/log"
	"github.com/blocktree/openwallet/openwallet"
)

func testGetAssetsAccountBalance(tm *openw.WalletManager, walletID, accountID string) {
	balance, err := tm.GetAssetsAccountBalance(testApp, walletID, accountID)
	if err != nil {
		log.Error("GetAssetsAccountBalance failed, unexpected error:", err)
		return
	}
	log.Info("balance:", balance)
}

func testGetAssetsAccountTokenBalance(tm *openw.WalletManager, walletID, accountID string, contract openwallet.SmartContract) {
	balance, err := tm.GetAssetsAccountTokenBalance(testApp, walletID, accountID, contract)
	if err != nil {
		log.Error("GetAssetsAccountTokenBalance failed, unexpected error:", err)
		return
	}
	log.Info("token balance:", balance.Balance)
}

func testCreateTransactionStep(tm *openw.WalletManager, walletID, accountID, to, amount, feeRate string, contract *openwallet.SmartContract) (*openwallet.RawTransaction, error) {

	//err := tm.RefreshAssetsAccountBalance(testApp, accountID)
	//if err != nil {
	//	log.Error("RefreshAssetsAccountBalance failed, unexpected error:", err)
	//	return nil, err
	//}

	rawTx, err := tm.CreateTransaction(testApp, walletID, accountID, amount, to, feeRate, "", contract)

	if err != nil {
		log.Error("CreateTransaction failed, unexpected error:", err)
		return nil, err
	}

	return rawTx, nil
}

func testCreateSummaryTransactionStep(
	tm *openw.WalletManager,
	walletID, accountID, summaryAddress, minTransfer, retainedBalance, feeRate string,
	start, limit int,
	contract *openwallet.SmartContract) ([]*openwallet.RawTransaction, error) {

	rawTxArray, err := tm.CreateSummaryTransaction(testApp, walletID, accountID, summaryAddress, minTransfer,
		retainedBalance, feeRate, start, limit, contract)

	if err != nil {
		log.Error("CreateSummaryTransaction failed, unexpected error:", err)
		return nil, err
	}

	return rawTxArray, nil
}

func testSignTransactionStep(tm *openw.WalletManager, rawTx *openwallet.RawTransaction) (*openwallet.RawTransaction, error) {

	_, err := tm.SignTransaction(testApp, rawTx.Account.WalletID, rawTx.Account.AccountID, "12345678", rawTx)
	if err != nil {
		log.Error("SignTransaction failed, unexpected error:", err)
		return nil, err
	}

	log.Infof("rawTx: %+v", rawTx)
	return rawTx, nil
}

func testVerifyTransactionStep(tm *openw.WalletManager, rawTx *openwallet.RawTransaction) (*openwallet.RawTransaction, error) {

	//log.Info("rawTx.Signatures:", rawTx.Signatures)

	_, err := tm.VerifyTransaction(testApp, rawTx.Account.WalletID, rawTx.Account.AccountID, rawTx)
	if err != nil {
		log.Error("VerifyTransaction failed, unexpected error:", err)
		return nil, err
	}

	log.Infof("rawTx: %+v", rawTx)
	return rawTx, nil
}

func testSubmitTransactionStep(tm *openw.WalletManager, rawTx *openwallet.RawTransaction) (*openwallet.RawTransaction, error) {

	tx, err := tm.SubmitTransaction(testApp, rawTx.Account.WalletID, rawTx.Account.AccountID, rawTx)
	if err != nil {
		log.Error("SubmitTransaction failed, unexpected error:", err)
		return nil, err
	}

	log.Std.Info("tx: %+v", tx)
	log.Info("wxID:", tx.WxID)
	log.Info("txID:", rawTx.TxID)

	return rawTx, nil
}

func TestTransfer_ERC20(t *testing.T) {

	addrs := []string{
		//"0x065881528680c7794594c136e145e94b8eae6908",
		//"0x51cfa12d6390dba967ad8f0bd9c040b5564e8ece",
		//"0x71010e21e2f7aaf9b6ef5dc651aeb380ba0d8e97",
		//"0x850c16cd1d6e1e440d8147c8190d6b250e50244b",
		//"0xa5b2006965e33993f4ca0082d16c6521a6d0daf7",
		//"0xaf3eafcd23d1110174118053b9d10f51b60483f5",

		"0xfb5cd467bc2d88f7308aa6838bc89bc3e53adb70",
		"0x0c18d28ad4a0be2f6fa81317ac03081996d09317",
	}

	tm := testInitWalletManager()
	walletID := "W8uQM5k5NLZEF3ge4Vx2wgxHPiHFQkdsZs"
	accountID := "AvE1MqXm8gqm7hKR79S1DcZiCSHeCvyLX4x4SxAi154n" //0xc97ac4202b860e381659851c8f3e272554aa9d6e

	contract := openwallet.SmartContract{
		Address:  "0x46ecadaa38f0562f2108bde63fc605ccaaad649e",
		Symbol:   "FAC",
		Name:     "FUQI-TEST",
		Token:    "TFUQ",
		Decimals: 18,
	}

	for _, to := range addrs {
		testGetAssetsAccountTokenBalance(tm, walletID, accountID, contract)

		rawTx, err := testCreateTransactionStep(tm, walletID, accountID, to, "1000", "", &contract)
		if err != nil {
			return
		}

		_, err = testSignTransactionStep(tm, rawTx)
		if err != nil {
			return
		}

		_, err = testVerifyTransactionStep(tm, rawTx)
		if err != nil {
			return
		}

		_, err = testSubmitTransactionStep(tm, rawTx)
		if err != nil {
			return
		}
	}
}

func TestSummary_ERC20(t *testing.T) {
	tm := testInitWalletManager()
	walletID := "WMTUzB3LWaSKNKEQw9Sn73FjkEoYGHEp4B"
	accountID := "59t47qyjHUMZ6PGAdjkJopE9ffAPUkdUhSinJqcWRYZ1"
	summaryAddress := "0xd35f9Ea14D063af9B3567064FAB567275b09f03D"

	contract := openwallet.SmartContract{
		Address:  "4092678e4E78230F46A1534C0fbc8fA39780892B",
		Symbol:   "ETH",
		Name:     "OCoin",
		Token:    "OCN",
		Decimals: 18,
	}

	testGetAssetsAccountBalance(tm, walletID, accountID)

	testGetAssetsAccountTokenBalance(tm, walletID, accountID, contract)

	rawTxArray, err := testCreateSummaryTransactionStep(tm, walletID, accountID,
		summaryAddress, "", "", "",
		0, 100, &contract)
	if err != nil {
		log.Errorf("CreateSummaryTransaction failed, unexpected error: %v", err)
		return
	}

	//执行汇总交易
	for _, rawTx := range rawTxArray {
		_, err = testSignTransactionStep(tm, rawTx)
		if err != nil {
			return
		}

		_, err = testVerifyTransactionStep(tm, rawTx)
		if err != nil {
			return
		}

		_, err = testSubmitTransactionStep(tm, rawTx)
		if err != nil {
			return
		}
	}

}

func TestTransfer_ERC20_Timer(t *testing.T) {

	addrs := []string{
		"0x065881528680c7794594c136e145e94b8eae6908",
	}

	tm := testInitWalletManager()
	walletID := "W8uQM5k5NLZEF3ge4Vx2wgxHPiHFQkdsZs"
	accountID := "AvE1MqXm8gqm7hKR79S1DcZiCSHeCvyLX4x4SxAi154n" //0xc97ac4202b860e381659851c8f3e272554aa9d6e

	contract := openwallet.SmartContract{
		Address:  "0x46ecadaa38f0562f2108bde63fc605ccaaad649e",
		Symbol:   "FAC",
		Name:     "FUQI-TEST",
		Token:    "TFUQ",
		Decimals: 18,
	}

	for {

		for _, to := range addrs {
			testGetAssetsAccountTokenBalance(tm, walletID, accountID, contract)

			rawTx, err := testCreateTransactionStep(tm, walletID, accountID, to, "1", "", &contract)
			if err != nil {
				return
			}

			_, err = testSignTransactionStep(tm, rawTx)
			if err != nil {
				return
			}

			_, err = testVerifyTransactionStep(tm, rawTx)
			if err != nil {
				return
			}

			_, err = testSubmitTransactionStep(tm, rawTx)
			if err != nil {
				return
			}
		}

		time.Sleep(8 * time.Second)
	}
}
