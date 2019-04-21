package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

// IJs ...
type IJs struct {
	WG [][]float64 `json:"wg"`
}

// IJss ...
type IJss struct {
	BG [][]float64 `json:"bg"`
}

// IJsss ...
type IJsss struct {
	IVA [][][]float64 `json:"iva"`
}

func setRand() [][]float64 {
	ijss := IJs{}
	db := "./wg.json"
	content, err := ioutil.ReadFile(db)
	if err != nil {
		log.Fatalln(err)
	}
	// fmt.Println("content:", content)
	if err = json.Unmarshal(content, &ijss); err != nil {
		log.Fatalln(err)
	}

	return ijss.WG
}

func setRand2() [][]float64 {
	ijss := IJss{}
	db := "./bg.json"
	content, err := ioutil.ReadFile(db)
	if err != nil {
		log.Fatalln(err)
	}
	// fmt.Println("content:", content)
	if err = json.Unmarshal(content, &ijss); err != nil {
		log.Fatalln(err)
	}

	return ijss.BG
}

func setRand3() [][][]float64 {
	ijss := IJsss{}
	db := "./iva.json"
	content, err := ioutil.ReadFile(db)
	if err != nil {
		log.Fatalln(err)
	}
	// fmt.Println("content:", content)
	if err = json.Unmarshal(content, &ijss); err != nil {
		log.Fatalln(err)
	}

	return ijss.IVA
}
