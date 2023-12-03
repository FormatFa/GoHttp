package storage

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/patrickmn/go-cache"
)

var instance *Storage = &Storage{
	FileName: "cache.godb",
}

type Storage struct {
	FileName string
	Cache    *cache.Cache
}

func (storage *Storage) InitFromFile() {
	_, error := os.Stat(storage.FileName)

	// check if error is "file not exists"
	if os.IsNotExist(error) {
		storage.Cache = cache.New(cache.NoExpiration, 10*time.Minute)
		fmt.Println("cache file not found")
	} else {
		fmt.Println("read cache file done.")
		temp := make(map[string]cache.Item, 10)
		raw, err := os.ReadFile(storage.FileName)
		if err != nil {
			panic(err)
		}
		buffer := bytes.NewBuffer(raw)
		dec := gob.NewDecoder(buffer)
		err = dec.Decode(&temp)
		if err != nil {
			panic(err)
		}
		storage.Cache = cache.NewFrom(cache.NoExpiration, 10*time.Minute, temp)

	}
}
func (storage *Storage) SaveToFile() {
	buffer := new(bytes.Buffer)
	encoder := gob.NewEncoder(buffer)
	err := encoder.Encode(storage.Cache.Items())
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile(storage.FileName, buffer.Bytes(), 0600)
	if err != nil {
		panic(err)
	}
}
func GetInstance() *Storage {
	return instance
}
