/*
Copyright © 2025 Abinand P <abinand0911@gmail.com>
*/
package redisclient

type RedisSearchResult struct {
	TotalResults int64           `json:"total_results"`
	Results      []RedisDocument `json:"results"`
}

type RedisDocument struct {
	ID              string          `json:"id"`
	ExtraAttributes ExtraAttributes `json:"extra_attributes"`
}

type ExtraAttributes struct {
	Command      string `json:"command"`
	Os           string `json:"os"`
	TextChunk    string `json:"text_chunk"`
	VectorAnswer string `json:"vectoranswer"`
}

type QuerySearchResult struct {
	Command        string  `json:"command"`
	Os             string  `json:"os"`
	TextChunk      string  `json:"text_chunk"`
	VectorDistance float64 `json:"vector_distance"`
}
