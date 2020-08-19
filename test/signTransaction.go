package main

import (
    "context"
    "crypto/ecdsa"
    "fmt"
    "log"
    "math/big"
    "git.weilaicaijing.com/RMCEth/ethereumRMC/common"
    "git.weilaicaijing.com/RMCEth/ethereumRMC/core/types"
    "git.weilaicaijing.com/RMCEth/ethereumRMC/crypto"
    "git.weilaicaijing.com/RMCEth/ethereumRMC/ethclient"
)
func main() {
    client, err := ethclient.Dial("http://chain-node.galaxynetwork.vip")
    if err != nil {
        log.Fatal(err)
    }
    var data []byte
    data=[]byte("")

//0xf5403E4F120901407eF221E2419583D1F3556953
    privateKey, err := crypto.HexToECDSA("cfcfa295cab51ccae9110ed6932c2c68dc0b94dba300baf5b8890906b248b50b")
    if err != nil {
        log.Fatal(err)
    }
    publicKey := privateKey.Public()
    publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
    if !ok {
        log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
    }
    fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
    if err != nil {
        log.Fatal(err)
    }

    value,_:= new(big.Int).SetString("1",10)

    gasPrice, err := client.SuggestGasPrice(context.Background())
    if err != nil {
        log.Fatal(err)
    }

    toAddress := common.HexToAddress("RMC6cBe9DF6DF54281D363e7a5e1790dc66212438C7")

        
    nonce, err := client.PendingNonceAt(context.Background(), fromAddress)

    tx := types.NewTransaction(nonce, toAddress, value,3000000, big.NewInt(gasPrice.Int64()), data)

    chainID, err := client.NetworkID(context.Background())
    if err != nil {
        log.Fatal(err)
    }

    var signedTx *types.Transaction
    signedTx, err = types.SignTx(tx, types.NewEIP155Signer(big.NewInt(chainID.Int64())), privateKey)
    if err != nil {
        log.Fatal(err)    
    }
    
    err = client.SendTransaction(context.Background(), signedTx)
    if err != nil {
        log.Fatal(err)
    }

    log.Printf("tx Hash: %v\n", signedTx.Hash().Hex())
    log.Println("Waiting for the transaction, about 4 minutes...")
	for {
    tx, isPending, err := client.TransactionByHash(context.Background(), signedTx.Hash())
    if err != nil {
        log.Fatal(err)
    }
    if isPending==false{
         fmt.Println("transaction is successful!!")
		 receipt, err := client.TransactionReceipt(context.Background(), tx.Hash())
		if err != nil {
			log.Fatal(err)
        }
        if receipt.Status==0{
            log.Fatal( "Error: Transaction has been reverted by the EVM")
        }
		fmt.Printf("receipt.Status:%v\n",receipt.Status)
		return 
    }
   }
}