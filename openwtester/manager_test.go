package openwtester

import (
	"encoding/hex"
	"fmt"
	"github.com/blocktree/openwallet/hdkeystore"
	"github.com/blocktree/openwallet/log"
	"github.com/blocktree/openwallet/openw"
	"github.com/blocktree/openwallet/openwallet"
	"path/filepath"
	"testing"
)

var (
	testApp        = "fukubukuro-adapter"
	configFilePath = filepath.Join("conf")
)

func testInitWalletManager() *openw.WalletManager {
	log.SetLogFuncCall(true)
	tc := openw.NewConfig()

	tc.ConfigDir = configFilePath
	tc.EnableBlockScan = false
	tc.SupportAssets = []string{
		"FAC",
	}
	return openw.NewWalletManager(tc)
	//tm.Init()
}

func TestWalletManager_CreateWallet(t *testing.T) {
	tm := testInitWalletManager()
	w := &openwallet.Wallet{Alias: "HELLO FAC", IsTrust: true, Password: "12345678"}
	nw, key, err := tm.CreateWallet(testApp, w)
	if err != nil {
		log.Error(err)
		return
	}

	log.Info("wallet:", nw)
	log.Info("key:", key)

}

func TestWalletManager_GetWalletInfo(t *testing.T) {

	tm := testInitWalletManager()

	wallet, err := tm.GetWalletInfo(testApp, "W8uQM5k5NLZEF3ge4Vx2wgxHPiHFQkdsZs")
	if err != nil {
		log.Error("unexpected error:", err)
		return
	}
	log.Info("wallet:", wallet)
}

func TestWalletManager_GetWalletList(t *testing.T) {

	tm := testInitWalletManager()

	list, err := tm.GetWalletList(testApp, 0, 10000000)
	if err != nil {
		log.Error("unexpected error:", err)
		return
	}
	for i, w := range list {
		log.Info("wallet[", i, "] :", w)
	}
	log.Info("wallet count:", len(list))

	tm.CloseDB(testApp)
}

func TestWalletManager_CreateAssetsAccount(t *testing.T) {

	tm := testInitWalletManager()

	walletID := "W8uQM5k5NLZEF3ge4Vx2wgxHPiHFQkdsZs"
	account := &openwallet.AssetsAccount{Alias: "mainnetFAC", WalletID: walletID, Required: 1, Symbol: "FAC", IsTrust: true}
	account, address, err := tm.CreateAssetsAccount(testApp, walletID, "12345678", account, nil)
	if err != nil {
		log.Error(err)
		return
	}

	log.Info("account:", account)
	log.Info("address:", address)

	tm.CloseDB(testApp)
}

func TestWalletManager_GetAssetsAccountList(t *testing.T) {

	tm := testInitWalletManager()

	walletID := "W8uQM5k5NLZEF3ge4Vx2wgxHPiHFQkdsZs"
	list, err := tm.GetAssetsAccountList(testApp, walletID, 0, 10000000)
	if err != nil {
		log.Error("unexpected error:", err)
		return
	}
	for i, w := range list {
		log.Info("account[", i, "] :", w)
	}
	log.Info("account count:", len(list))

	tm.CloseDB(testApp)

}

func TestWalletManager_CreateAddress(t *testing.T) {

	tm := testInitWalletManager()

	walletID := "W8uQM5k5NLZEF3ge4Vx2wgxHPiHFQkdsZs"
	accountID := "cm3dSCFP47yhgPfhYLab8ChprH2wHfwrWtEuSEkNW29"
	address, err := tm.CreateAddress(testApp, walletID, accountID, 5)
	if err != nil {
		log.Error(err)
		return
	}

	log.Info("address:", address)

	tm.CloseDB(testApp)
}

func TestWalletManager_GetAddressList(t *testing.T) {

	tm := testInitWalletManager()

	walletID := "W8uQM5k5NLZEF3ge4Vx2wgxHPiHFQkdsZs"
	//accountID := "AvE1MqXm8gqm7hKR79S1DcZiCSHeCvyLX4x4SxAi154n" //0xc97ac4202b860e381659851c8f3e272554aa9d6e
	accountID := "cm3dSCFP47yhgPfhYLab8ChprH2wHfwrWtEuSEkNW29" //0xaf3eafcd23d1110174118053b9d10f51b60483f5
	list, err := tm.GetAddressList(testApp, walletID, accountID, 0, -1, false)
	if err != nil {
		log.Error("unexpected error:", err)
		return
	}
	for _, w := range list {
		fmt.Printf("%s \n", w.Address)
	}
	log.Info("address count:", len(list))

	tm.CloseDB(testApp)
}

func TestGenerateSeed(t *testing.T) {
	seed, _ := hdkeystore.GenerateSeed(32)
	log.Infof("seed: %s", hex.EncodeToString(seed))
}