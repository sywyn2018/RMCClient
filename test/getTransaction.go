package main

import (
	"context"
	"log"
	// "math/big"
	"encoding/hex"
	"git.weilaicaijing.com/RMCEth/ethereumRMC/common"
	"git.weilaicaijing.com/RMCEth/ethereumRMC/core/types"
	"git.weilaicaijing.com/RMCEth/ethereumRMC/ethclient"
)
func main() {
    client, err := ethclient.Dial("http://chain-node.galaxynetwork.vip")
    if err != nil {
        log.Fatal(err)
	}
	var transaction *types.Transaction
	
	TransactionHash := common.HexToHash("0xa7ab5a9d119afb4e3e2a97c627e6bd31b1232db6055a77e0f28dab97bfec7af1")
	transaction,isPending,err:=client.TransactionByHash(context.Background(),TransactionHash)
	if err != nil {
        log.Fatal(err)
	}

	if isPending==false{
    chainID, err := client.NetworkID(context.Background())
    if err != nil {
        log.Fatal(err)
	}
//get "From" must change msg type
	msg, err := transaction.AsMessage(types.NewEIP155Signer(chainID))
	if err != nil {
		log.Fatal("getFromErr",err)
	}

	log.Println("交易Hash:",transaction.Hash().Hex()) 
	log.Println("From:",msg.From().Hex())   
	log.Println("To:",transaction.To().Hex())      
	log.Println("Value:",transaction.Value().String())    
	log.Println("Gas:",transaction.Gas())               
	log.Println("GasPrice:",transaction.GasPrice().Uint64()) 
	log.Println("Nonce:",transaction.Nonce())             
	log.Println("Data:",hex.EncodeToString(transaction.Data()))
}
}