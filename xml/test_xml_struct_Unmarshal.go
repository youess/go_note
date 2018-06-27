package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

// 解析的结构体Field第一个字母必须大写
type SysConf struct {
	XMLName     xml.Name `xml:"sysconfig"` // 是XML块中的数据能够引用出来
	Port        int      `xml:"port,attr"`
	ProjectName string   `xml:"project_name,attr"`
	extra       string
}

// 多重只需要对应其中一个结构体就好
type KConf struct {
	XMLName xml.Name `xml:"kcfg"`
	Topic   string   `xml:"topic,attr"`
	Group   string   `xml:"group,attr"`
	From    string   `xml:"from,attr"`
	Addrs   []string `xml:"addr"`
	extra   string
}

type MyXMLConfig struct {
	XMLName   xml.Name `xml:"config"`
	SysConfig *SysConf `xml:"sysconfig"`
	Kcfg      []*KConf `xml:"kcfg"`
	extra     string
}

func main() {
	configFile := "config.xml"
	cf, err := os.Open(configFile)
	defer cf.Close()
	if err != nil {
		fmt.Printf("open config file failed, %s", err)
	}
	data, err := ioutil.ReadAll(cf)
	if err != nil {
		fmt.Printf("read config file data failed, %s", err)
	}
	// fmt.Println(data)
	cfg := &MyXMLConfig{}
	err = xml.Unmarshal(data, cfg)
	if err != nil {
		fmt.Printf("xml data parse failed, %s", err)
	}
	// fmt.Println(*cfg)

	fmt.Printf("config name: %s \n", cfg.XMLName)
	fmt.Printf("sys config app_port: %d project_name: %s\n",
		cfg.SysConfig.Port, cfg.SysConfig.ProjectName)
	/*
		fmt.Printf("kcfg topic: %s, group: %s, from: %s, addrs: %s",
			cfg.Kcfg.Topic, cfg.Kcfg.Group, cfg.Kcfg.From, cfg.Kcfg.Addrs)
	*/
	for _, kc := range cfg.Kcfg {
		fmt.Println(" ")
		fmt.Println("Topic: ", kc.Topic)
		fmt.Println("Group: ", kc.Group)
		fmt.Println("From: ", kc.From)
		fmt.Println("Addrs: ", kc.Addrs)
	}
	// fmt.Println(cfg.Kcfg)
	fmt.Printf("\nRefer to Here to Checkout more XML Export Problem:\nhttps://github.com/astaxie/build-web-application-with-golang/blob/master/zh/07.1.md")
}
