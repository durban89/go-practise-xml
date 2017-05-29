package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

type Servers struct {
	XMLName     xml.Name `xml:"servers"`
	Version     string   `xml:"version,attr"`
	Server      []Server `xml:"server"`
	Description string   `xml:",innerxml"`
}

type Server struct {
	XMLName    xml.Name `xml:"server"`
	ServerName string   `xml:"serverName"`
	ServerIP   string   `xml:"serverIP"`
}

type XMLServers struct {
	XMLName xml.Name    `xml:"servers"`
	Version string      `xml:"version,attr"`
	Server  []XMLServer `xml:"server"`
}

type XMLServer struct {
	ServerName string `xml:"serverName"`
	ServerIP   string `xml:"serverIP"`
}

func BuildXml() {
	v := &XMLServers{Version: "1"}
	v.Server = append(v.Server, XMLServer{"Shanghai_VPN", "127.0.0.1"})
	v.Server = append(v.Server, XMLServer{"Beijing_VPN", "127.0.0.2"})

	output, err := xml.MarshalIndent(v, " ", "  ")
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}

	os.Stdout.Write([]byte(xml.Header))
	os.Stdout.Write(output)
}

func main() {
	// For read access.
	file, err := os.Open("servers.xml")
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}

	defer file.Close()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}

	v := Servers{}
	err = xml.Unmarshal(data, &v)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}

	fmt.Println(v)

	BuildXml()
}
