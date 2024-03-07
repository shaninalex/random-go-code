package hashing

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"log"
)

func ExecuteHash() {
	tx := "CqQBCqEBCiMvY29zbW9zLnN0YWtpbmcudjFiZXRhMS5Nc2dEZWxlZ2F0ZRJ6Ci9jZWxlc3RpYTF5OXkydDNyamVtdzlwNnJzd2NjZDM4NmEzZ2p1cTRmamY3cXN0eRI2Y2VsZXN0aWF2YWxvcGVyMXpxanBmeHR2M3lwNmtkbGdyYTRoYzl6ZWh4Z3ZwYXc4Mmh4cjV3Gg8KBHV0aWESBzE0MDAwMDASZgpQCkYKHy9jb3Ntb3MuY3J5cHRvLnNlY3AyNTZrMS5QdWJLZXkSIwohA0fkeXrT3IajKrpFSebnEOmAuc9VZe5dlpg8F7AjjJibEgQKAgh/GAUSEgoMCgR1dGlhEgQ0NDc1ENXTDRpAmQ5tQtnHISo2bGhLLwyBsdxhz43/DURSlYeMUTX14t5lujn3lKmdZHcJ69EJZs0SJcs0c7k5EPpqEzeb6Da1cw=="

	encodedTx, err := base64.StdEncoding.DecodeString(tx)

	if err != nil {
		panic(err)
	}

	hash := sha256.Sum256(encodedTx)
	txHash := hex.EncodeToString(hash[:])
	log.Println("encodedTx: ", string(encodedTx))
	log.Println("txHash: ", txHash)

	// tx: "CpgBCpUBChwvY29zbW9zLmJhbmsudjFiZXRhMS5Nc2dTZW5kEnUKL2NlbGVzdGlhMXhmZWNmYWRuZjhtbWp3dGE4bGdtejR4ODNkZHpyOXlxcHByZ2o5Ei9jZWxlc3RpYTE2NjAyc2V6M2RrY2t5cjVhN3dqMng1bmxqdHN0NWNlcmU5dmpocxoRCgR1dGlhEgkxNjI5ODAwMDASaQpOCkYKHy9jb3Ntb3MuY3J5cHRvLnNlY3AyNTZrMS5QdWJLZXkSIwohAtPs2XMyzVfai0Z4u7JgyBTLybM6tz+M4d9E1KLMdFcsEgQKAggBEhcKEQoEdXRpYRIJMTYyOTgwMDAwENDoDBpAtT4C4w+uq7vpfpp8U+u6tYaDmXPzttw2vEO1e58fJzscmw2nwkVx6ZMIvz+8iZpmU04TAkiSqcBb8ZKl1qDeTA=="
	// txHash: fcda965b374dc35340c367e13d7944bab33f7014be46fc30c506956f5c492081
}
