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

func ExecDistribute(min, max *int) {

	Min := 3
	Max := 10

	if *min != 0 {
		Min = *min
	}

	if *max != 0 {
		Max = *max
	}

	for i := range 15 {
		distribute(i, Min, Max)
	}
}

func distribute(iter, min, max int) {
	set := generateSetOfHashes(rand.IntN(max-min) + min)

	// TODO: check node availability before assign

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

		// if part of the set has length - assign tasks (transactions) to node
		if len(set[startIndex:endIndex]) != 0 {
			distributedSets[node] = set[startIndex:endIndex]
		}

		startIndex = endIndex
	}

	// just add total_task key to validate that all tasks distributed between nodes
	distributedSets["total_tasks"] = []string{fmt.Sprintf("%d", len(set))}

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
