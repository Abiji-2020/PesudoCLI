/*
Copyright © 2025 Abinand P <abinand0911@gmail.com>
*/

package gemini

import (
	"strconv"
	"strings"

	"github.com/Abiji-2020/PesudoCLI/internal/redisclient"
)

func BuildContextPrompt(queryChunks []redisclient.QuerySearchResult) string {
	var builder strings.Builder

	builder.WriteString("You are a helpful assistant for command line tasks. You have a vast knowledge on the command line and terminal commands of linux, macOs, OSX, windows, andriod and ios. You are able to answer questions about these operating systems and their command line interfaces.\n\n")
	builder.WriteString("Here are some relevant commands of the question, and they are targeted with the relavant operating system and the distance measure where the distance mesasure less means it is more accurate or relavant to the question. \n\n")
	for _, chunk := range queryChunks {
		builder.WriteString("Command: " + chunk.Command + "\n")
		builder.WriteString("OS: " + chunk.Os + "\n")
		builder.WriteString("Text Chunk: " + chunk.TextChunk + "\n")
		builder.WriteString("Vector Distance: " + strconv.FormatFloat(chunk.VectorDistance, 'f', 6, 64) + "\n\n")
	}
	builder.WriteString("You should provide a detailed answer in the following described json format, which has command, os, explanation and answer.\n")
	builder.WriteString("The explanation should be very small, more like a one liner")
	builder.WriteString("Answer should be a detailed one and it should have examples and relavant details to it. and only contains plain text formatted by tags.. No markdown or other kind of formats for any values within it.\n")
	builder.WriteString("The answer should be in the following json format:\n")
	builder.WriteString("{\n")
	builder.WriteString("  \"command\": \"\",\n")
	builder.WriteString("  \"os\": \"\",\n")
	builder.WriteString("  \"explanation\": \"\",\n")
	builder.WriteString("  \"answer\": \"\"\n")
	builder.WriteString("}\n")
	builder.WriteString("Make sure to provide a detailed answer and not just a simple one. The answer should be in the context of the question and the commands provided.\n")
	builder.WriteString("If you don't know the answer, just say that you don't know the answer.\n")
	builder.WriteString("If the question is not related to command line or terminal commands, just say that you don't know the answer in both explanaiton and answer and leave other two as empty values .\n")
	builder.WriteString("If the question is not related to the operating systems mentioned, just say that you don't know the answer in both explanation and answer and leave other two as a empty string.\n")
	return builder.String()
}
