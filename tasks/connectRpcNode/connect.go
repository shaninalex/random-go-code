package connectrpcnode

import (
	"context"
	"fmt"
	"log"

	tmclient "github.com/tendermint/tendermint/rpc/client/http"
	rpctypes "github.com/tendermint/tendermint/rpc/jsonrpc/types"
)

func ExecuteConnect() {
	// client, err := tmclient.New("http://localhost:26657")
	client, err := tmclient.New("http://consensus.lunaroasis.net:26657")
	if err != nil {
		log.Fatalf("Error creating RPC client: %v", err)
	}

	blockHeight := int64(942647)

	// Fetch block information
	result, err := client.Block(context.Background(), &blockHeight)
	if err != nil {
		if rpcErr, ok := err.(rpctypes.RPCError); ok && rpcErr.Code == 1 {
			log.Fatalf("Error: block %d not found", blockHeight)
		} else {
			log.Fatalf("Error fetching block: %v", err)
		}
	}

	txHashes := make([]string, len(result.Block.Txs))
	for i, tx := range result.Block.Txs {
		txHashes[i] = fmt.Sprintf("%X", tx.Hash())
	}

	fmt.Println("Transaction hashes in block", blockHeight)
	for i := range txHashes {
		log.Println(i, txHashes[i])
	}
}
