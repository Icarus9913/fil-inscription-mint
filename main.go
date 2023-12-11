package main

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/filecoin-project/go-address"
	jsonrpc "github.com/filecoin-project/go-jsonrpc"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/builtin"
	lotusapi "github.com/filecoin-project/lotus/api"
	"github.com/filecoin-project/lotus/chain/types"
	"github.com/filecoin-project/lotus/chain/wallet/key"
	"github.com/filecoin-project/lotus/lib/sigs"
	_ "github.com/filecoin-project/lotus/lib/sigs/secp"
	"github.com/ipfs/go-cid"
)

const (
	fil20 = `data:,{"p":"fil-20","op":"mint","tick":"fils","amt":"1000"}`

	MainNet = "https://api.node.glif.io/"
	TestNet = "https://api.calibration.node.glif.io/"

	walletAddr       = ""
	walletPrivateKey = ""
)

type Admin struct {
	Api    *lotusapi.FullNodeStruct
	Closer *jsonrpc.ClientCloser

	Wallet  address.Address
	KeyInfo *types.KeyInfo
}

func main() {
	ctx := context.TODO()
	admin := initAdmin(ctx)
	defer (*admin.Closer)()

	// Now you can call any API you're interested in.
	tipset, err := admin.Api.ChainHead(context.Background())
	if err != nil {
		log.Fatalf("calling chain head: %s", err)
	}
	log.Printf("Current chain head height is: %s", tipset.Height())

	//////////////////////
	msg, err := admin.generateMsg(ctx)
	if nil != err {
		log.Fatalf("failed to generate message, error: %v", err)
	}

	log.Printf("gas limit: %v", msg.GasLimit)
	log.Printf("gas fee cap: %v", msg.GasFeeCap)
	log.Printf("gas premium: %v", msg.GasPremium)

	err = admin.mint(ctx, msg)
	if nil != err {
		log.Fatalf("failed to mint message, error: %v", err)
	}
}

func initAdmin(ctx context.Context) *Admin {
	wallet, err := address.NewFromString(walletAddr)
	if nil != err {
		log.Fatal(err.Error())
	}

	decodeString, err := hex.DecodeString(walletPrivateKey)
	if nil != err {
		log.Fatalf(err.Error())
	}
	var ki types.KeyInfo
	err = json.Unmarshal(decodeString, &ki)
	if nil != err {
		log.Fatal(err.Error())
	}

	var api lotusapi.FullNodeStruct
	closer, err := jsonrpc.NewMergeClient(
		ctx,
		MainNet, "Filecoin",
		[]interface{}{&api.Internal, &api.CommonStruct.Internal},
		http.Header{})
	if err != nil {
		log.Fatalf("connecting with lotus failed: %s", err)
	}

	admin := &Admin{
		Api:    &api,
		Closer: &closer,

		Wallet:  wallet,
		KeyInfo: &ki,
	}

	return admin
}

func (a *Admin) generateMsg(ctx context.Context) (*types.Message, error) {
	var err error
	msg := &types.Message{
		To:     a.Wallet,
		From:   a.Wallet,
		Value:  abi.NewTokenAmount(0), // 1 fil = 1,000,000,000,000,000,000 = 1 * 10^18
		Method: builtin.MethodsEVM.InvokeContract,
		Params: []byte(fil20),
	}
	msg, err = a.Api.GasEstimateMessageGas(ctx, msg, nil, types.EmptyTSK)
	if nil != err {
		return nil, fmt.Errorf("failed to estimate message gas, error: %v", err)
	}

	return msg, nil
}

func (a *Admin) mint(ctx context.Context, msg *types.Message) error {
	signature, err := sigs.Sign(key.ActSigType(a.KeyInfo.Type), a.KeyInfo.PrivateKey, msg.Cid().Bytes())
	if nil != err {
		return fmt.Errorf("failed to sign, error: %v", err)
	}

	signedMessage := &types.SignedMessage{
		Message:   *msg,
		Signature: *signature,
	}
	log.Printf("message signed")

	msgCid, err := a.Api.MpoolPush(ctx, signedMessage)
	if nil != err {
		return fmt.Errorf("failed to push message, error: %v", err)
	}

	log.Printf("message pushed, cid: %s", msgCid)
	return nil
}

func (a *Admin) getMsg(ctx context.Context, cidStr string) {
	decodeCID, err := cid.Decode(cidStr)
	if nil != err {
		log.Fatalf(err.Error())
	}

	message, err := a.Api.ChainGetMessage(ctx, decodeCID)
	if nil != err {
		log.Fatalf(err.Error())
	}

	marshalJSON, err := message.MarshalJSON()
	if nil != err {
		log.Fatalf(err.Error())
	}
	log.Println(string(marshalJSON))
	log.Println(string(message.Params))
}
