package main

import (
    "encoding/xml"
    "fmt"
)

type Config struct {
    XMLName xml.Name `xml:"config"`
    Version string   `xml:"version,attr"`
    Servers []Server `xml:"servers>server"`
}

type Server struct {
    IP      string `xml:"ip"`
    Port    int    `xml:"port"`
    Enabled bool   `xml:"enabled"`
}

// Сериализация в XML
func toXML(v any) ([]byte, error) {
    return xml.MarshalIndent(v, "", "  ")
}

// Десериализация из XML
func fromXML(data []byte, v any) error {
    return xml.Unmarshal(data, v)
}

// XML демо
func xmlDemo() {
    // Создаём конфиг
    config := Config{
        Version: "1.0",
        Servers: []Server{
            {"192.168.0.1", 8080, true},
            {"10.0.0.1", 80, false},
        },
    }
    
    // Сериализуем
    xmlData, _ := toXML(config)
    writeFile("config.xml", xmlData)
    
    // Читаем и десериализуем
    data, _ := readFile("config.xml")
    var newConfig Config
    fromXML(data, &newConfig)
    
    fmt.Println("\nXML Demo:")
    fmt.Println("Исходный:", config)
    fmt.Println("Из файла:", newConfig)
    fmt.Println("Сырой XML:", string(data))
}