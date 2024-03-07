package distr

import (
	"encoding/json"
	"fmt"
	"math/rand/v2"
	"os"

	"github.com/google/uuid"
)

var (
	nodes []string = []string{
		"http://celestia-rpc.easy2stake.com",
		"http://celestia-rpc.brightlystake.com",
		"http://rpc.celestia.pops.one:26657",
		"http://celestia-rpc.spidey.services",
		"http://rpc-celestia-full.avril14th.org",
	}
)

func ExecDistribute() {
	for i := range 15 {
		distribute(i)
	}
}

func distribute(iter int) {
	set := generateSetOfHashes(rand.IntN(45-3) + 3)

	itemsPerNode := len(set) / len(nodes)
	remainder := len(set) % len(nodes)

	distributedSets := make(map[string][]string)
	startIndex := 0
	for _, node := range nodes {
		endIndex := startIndex + itemsPerNode
		if remainder > 0 {
			endIndex++
			remainder--
		}
		distributedSets[node] = set[startIndex:endIndex]
		startIndex = endIndex
	}

	f, _ := os.Create(fmt.Sprintf("tasks/distr/debug-%d.json", iter))
	defer f.Close()
	as_json, _ := json.MarshalIndent(distributedSets, "", "\t")
	f.Write(as_json)
}

func generateSetOfHashes(length int) []string {
	set := []string{}
	for range length {
		set = append(set, uuid.NewString())
	}
	return set
}
