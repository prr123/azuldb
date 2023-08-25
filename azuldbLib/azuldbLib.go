// azuldb
// kv db for less than 10,000 entries
// Author: prr azul software
// Date: 25. Aug 2023
// copyright 2023 prr, azul software
//


package azuldb

import (
	"os"
	"fmt"
//	"log"

)

type tabdb struct {
	Dbg bool
	Ndb bool
	DirPath string
	TabFilList [16]string
	TabList [16]*tab
}

type tab struct {
	grp uint32
	hash []uint32
	key []string
	val []string
}

func Initdb(dirpath string) (azulDB tabdb, err error){

	// find dir
 	_, err = os.Open(dirpath)
	if err != nil {
		if os.IsNotExist(err) {
			//create directory
			err1 := os.Mkdir(dirpath, 0755)
			if err1 != nil {return azulDB, fmt.Errorf("could not create dir: %v", err1)}

			azulDB.DirPath = dirpath
			azulDB.Ndb = true

			//create files

			return azulDB, nil
		} else {
			return azulDB, fmt.Errorf("could not open dir: %v", err)
		}
	}

	azulDB.DirPath = dirpath
	azulDB.Ndb = false

	files, err := os.ReadDir(dirpath)
	if err != nil {return azulDB, fmt.Errorf("could not read dir: %v", err)}
	for _, file := range files {
        fmt.Println(file.Name(), file.IsDir())
    }

	return azulDB, nil
}


func PrintDb(azuldb *tabdb) {

	fmt.Printf("******* AzulDB: %s *******\n", azuldb.DirPath)
	fmt.Printf("Dir:    %s\n",(*azuldb).DirPath)
	fmt.Printf("New DB: %t\n",(*azuldb).Ndb)
	fmt.Printf("  table  entries\n")
	for i:=0; i< len((*azuldb).TabList); i++ {
		fmt.Printf("  [%2d]: ", i+1)

		if (*azuldb).TabList[i] == nil {
			fmt.Printf("   - \n")
		} else {
			table := (*azuldb).TabList[i]
			num := len((*table).hash)
			fmt.Printf(" %3d\n", num)
		}
	}
	fmt.Printf("********* End AzulDB: *******\n")
	return
}
