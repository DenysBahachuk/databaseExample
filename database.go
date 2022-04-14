package database

import (
	"encoding/gob"
	"fmt"
	"log"
	"os"
)

type inMemoryDB struct {
	path     string
	accounts map[string]string
}

func NewDataBase() *inMemoryDB {
	return &inMemoryDB{
		path:     "C:\\Users\\Appartment_103\\Desktop\\Golang\\database\\data.gob",
		accounts: make(map[string]string),
	}
}

func (db *inMemoryDB) Open() error {
	log.Println("Loading accounts:", db.path)
	file, err := os.Open(db.path)
	if err != nil {
		fmt.Println("The file doesn't exist.")
		return err
	}
	defer file.Close()
	decoder := gob.NewDecoder(file)
	err1 := decoder.Decode(&db.accounts)
	if err1 != nil {
		log.Println("Error decoding the file")
		return err1
	}
	return nil
}

func (db *inMemoryDB) Save() error {
	log.Println("Saving accounts: ", db.path)

	err := os.Remove(db.path)
	if err != nil {
		log.Println("Error removing file", err)
	}

	file, err1 := os.Create(db.path)
	if err1 != nil {
		fmt.Println("Error creating file")
		return err1
	}
	defer file.Close()

	encoder := gob.NewEncoder(file)
	err2 := encoder.Encode(&db.accounts)
	if err2 != nil {
		fmt.Println("Error encoding file")
		return err2
	}
	return nil
}

func (db *inMemoryDB) GetAccountsInfo() map[string]string {
	return db.accounts
}

func (db *inMemoryDB) AddAccount(name string, password string) {
	db.accounts[name] = password
}

func (db *inMemoryDB) ChangeName(oldName string, newName string, password string) {
	delete(db.accounts, oldName)
	db.AddAccount(newName, password)
}
