package core

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"gohttp/mvc"
	"gohttp/storage"
	"os"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"

	"github.com/patrickmn/go-cache"
)

var instance *Application = &Application{}

type Application struct {
	// CurView fyne.CanvasObject
	CurView mvc.BaseView
	FyneApp fyne.App
	Window  fyne.Window
	cache   *cache.Cache
}

func (application *Application) loadCache() {
	filename := "cache.godb"
	_, error := os.Stat(filename)

	// check if error is "file not exists"
	if os.IsNotExist(error) {
		application.cache = cache.New(cache.NoExpiration, 10*time.Minute)
		fmt.Println("cache file not found")
	} else {
		fmt.Println("read cache file done.")
		temp := make(map[string]cache.Item, 10)
		raw, err := os.ReadFile(filename)
		if err != nil {
			panic(err)
		}
		buffer := bytes.NewBuffer(raw)
		dec := gob.NewDecoder(buffer)
		err = dec.Decode(temp)
		if err != nil {
			panic(err)
		}
		application.cache = cache.NewFrom(cache.NoExpiration, 10*time.Minute, temp)

	}
	storage.GetInstance().Cache = application.cache

}
func (application *Application) Init() {
	storage.GetInstance().InitFromFile()
	application.FyneApp = app.NewWithID("online.indigo6a.gohttp")
	application.Window = application.FyneApp.NewWindow("Container")

	// menu := fyne.NewMainMenu(fyne.NewMenu("读取变量文件"))
	// application.Window.SetMainMenu(menu)
	application.Window.Resize(fyne.NewSize(1024, 768))

}
func (application *Application) Start() {

	application.Window.ShowAndRun()

}

func GetInstance() *Application {
	return instance
}
