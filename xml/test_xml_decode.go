package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

func main() {

	configFile := "config.xml"
	cf, err := os.Open(configFile)
	defer cf.Close()
	if err != nil {
		fmt.Printf("open config file failed, %s", err)
	}
	/*
		data, err := ioutil.ReadAll(cf)
		if err != nil {
			fmt.Printf("read config file data failed, %s", err)
		}
	*/

	decoder := xml.NewDecoder(cf)
	t, _ := decoder.Token()
	if t == nil {
		fmt.Printf("error when token xml file")
	}
	// fmt.Println()
	switch se := t.(type) {
	case xml.StartElement:
		fmt.Println("start element:")
		fmt.Println(se.Name)
	default:
		fmt.Println(se)
	}
}
