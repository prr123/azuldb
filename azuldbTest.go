// azuldb
// kv db for less than 10,000 entries
// Author: prr azul software
// Date: 25. Aug 2023
// copyright 2023 prr, azul software
//

package main 

import (
	"log"
	"fmt"
	"os"

	azuldb "db/azuldb/azuldbLib"
)

func main(){

	numargs := len(os.Args)

	if numargs < 2 {
		fmt.Printf("no path to db provided\n")
		os.Exit(-1)
	}

	dirPath := os.Args[1]
	log.Printf("dirPath: %s\n", dirPath)

	db, err := azuldb.Initdb(dirPath)
	if err != nil {log.Fatalf("initdb: %v",err )}

	azuldb.PrintDb(&db)
	log.Printf("success init azuldb\n")
}
