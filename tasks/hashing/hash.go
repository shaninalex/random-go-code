package hashing

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"log"
)

func ExecuteHash() {
	tx := "CrYBCqABCiMvY29zbW9zLnN0YWtpbmcudjFiZXRhMS5Nc2dEZWxlZ2F0ZRJ5Ci9jZWxlc3RpYTFkZWFhbXVqdzdrZ2dwNHo0ZzN4aHRzeDIyMmw4c2VnemRkdDhrNRI2Y2VsZXN0aWF2YWxvcGVyMWhtMmQ4ZTZuZDVuZ3RsdGUzaGxkZWQwM3ZnajNyZXI5NHZtZG1sGg4KBHV0aWESBjIwNjM1NhIRRGVsZWdhdGUocmV3YXJkcykSZwpQCkYKHy9jb3Ntb3MuY3J5cHRvLnNlY3AyNTZrMS5QdWJLZXkSIwohAuz4Ttqc44OX6LuTTskoZjEqWNl/tg49/2lUIukWm0kFEgQKAgh/GAMSEwoNCgR1dGlhEgU1MzU0MBCwqxAaQN2IsiNwYlLqTvYYZLx3sVssaddbG2m8Uwb261C81/+EOi1hYibD+nUkftn0G98vjFf4aIWGSbvmwzV0yRF2vhY="

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
