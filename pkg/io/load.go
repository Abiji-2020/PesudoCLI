/*
Copyright © 2025 Abinand P <abinand0911@gmail.com>
*/

package io

import (
	"embed"
	"encoding/csv"
	"fmt"
	"io"
	"strings"
)

//go:embed command.csv
var commandCSV embed.FS

type CommandDoc struct {
	Command   string    `json:"command"`
	Os        string    `json:"os"`
	TextChunk string    `json:"text_chunk"`
	Embedding []float32 `json:"embedding"`
}

func LoadComandDocs() ([]CommandDoc, error) {
	fileBytes, err := commandCSV.ReadFile("command.csv")
	if err != nil {
		return nil, err
	}

	reader := csv.NewReader(strings.NewReader(string(fileBytes)))
	reader.TrimLeadingSpace = true

	_, err = reader.Read() // Skip header
	if err != nil {
		return nil, err
	}
	var docs []CommandDoc
	for {
		row, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			continue
		}
		docs = append(docs, CommandDoc{
			Command:   row[0],
			Os:        row[1],
			TextChunk: row[2],
		})
	}
	fmt.Print("Loaded ", len(docs), " command docs\n")
	return docs, nil
}
