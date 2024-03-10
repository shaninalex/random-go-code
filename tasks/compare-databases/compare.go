package comparedatabases

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	celeniumDSN string = "host=localhost user=user password=password dbname=celenium_data port=5432 sslmode=disable"
	qubeBaseDSN string = "host=localhost user=user password=password dbname=qube_data port=5434 sslmode=disable"
)

type TxItem struct {
	Hash   string `gorm:"column:hash"`
	Fee    int64  `gorm:"column:fee"`
	Height int64  `gorm:"column:height"`
	IsPfb  bool   `gorm:"is_pfb"`
}

func createDb(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

func Execute() {

	// result file
	file, err := os.Create("compare.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	// initialize csv writer
	writer := csv.NewWriter(file)
	defer writer.Flush()

	errfile, err := os.Create("errors.csv")
	if err != nil {
		log.Fatal(errfile)
	}
	defer errfile.Close()
	// initialize csv writer
	errWriter := csv.NewWriter(errfile)
	defer errWriter.Flush()
	errWriter.Write([]string{"block", "celBlockItems", "qubeBlockItems"})

	headers := []string{
		"cel block number",
		"cel tx hash",
		"cel fee",
		"cel total block fee",
		"qube block number",
		"qube tx hash",
		"qube fee",
		"qube total block fee",
		"is_pfb",
		"fee difference",
	}
	writer.Write(headers)

	var mainData [][]string

	// celenium base
	celeniumDB, err := createDb(celeniumDSN)
	if err != nil {
		panic(err)
	}

	// qube base
	qubeDB, err := createDb(qubeBaseDSN)
	if err != nil {
		panic(err)
	}

	// get all celenium items
	var celItems []TxItem
	celeniumDB.Raw("SELECT t.hash, t.fee, t.height FROM tx_items t ORDER BY t.id DESC").Scan(&celItems)
	if celeniumDB.Error != nil {
		log.Println(celeniumDB.Error)
	}

	// get all qube items
	var qubeItems []TxItem
	qubeDB.Raw("SELECT t.hash, t.fee, t.block_number AS height, t.is_pfb FROM txes t ORDER BY t.id DESC").Scan(&qubeItems)
	if qubeDB.Error != nil {
		log.Println(qubeDB.Error)
	}

	for i := 100000; i > 0; i-- {
		b := int64(300000 - i)
		// log.Println(b)
		celBlockItems := findItemsByHeight(celItems, b)
		qubeBlockItems := findItemsByHeight(qubeItems, b)

		if len(celBlockItems) != len(qubeBlockItems) {
			log.Println("block transactions are not equeal from 2 dbs")
			errWriter.Write([]string{
				fmt.Sprintf("%d", b),
				fmt.Sprintf("%d", len(celBlockItems)),
				fmt.Sprintf("%d", len(qubeBlockItems)),
			})
			continue
		}

		for i := 0; i < len(celBlockItems); i++ {
			if qubeBlockItems[i].IsPfb {
				log.Println(qubeBlockItems[i].Height, strconv.FormatBool(qubeBlockItems[i].IsPfb))
			}

			data := []string{
				// celenium data
				fmt.Sprintf("%d", celBlockItems[i].Height),
				celBlockItems[i].Hash,
				fmt.Sprintf("%d", celBlockItems[i].Fee),
				fmt.Sprintf("%d", totalBlockFee(celBlockItems)),

				// qube data
				fmt.Sprintf("%d", qubeBlockItems[i].Height),
				strings.ToLower(qubeBlockItems[i].Hash),
				fmt.Sprintf("%d", qubeBlockItems[i].Fee),
				fmt.Sprintf("%d", totalBlockFee(qubeBlockItems)),
				strconv.FormatBool(qubeBlockItems[i].IsPfb),
				fmt.Sprintf("%d", totalBlockFee(celBlockItems)-totalBlockFee(qubeBlockItems)),
			}

			mainData = append(mainData, data)
		}
	}

	writer.WriteAll(mainData)
}

func findItemsByHeight(items []TxItem, height int64) []TxItem {
	var matchingItems []TxItem
	for _, item := range items {
		if item.Height == height {
			matchingItems = append(matchingItems, item)
		}
	}
	return matchingItems
}

func totalBlockFee(items []TxItem) int64 {
	var total int64 = 0
	for i := range items {
		total += items[i].Fee
	}
	return total
}
