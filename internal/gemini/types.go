/*
Copyright © 2025 Abinand P <abinand0911@gmail.com>
*/

package gemini

type GeminiResponse struct {
	Answer      string `json:"answer"`
	Command     string `json:"command"`
	Os          string `json:"os"`
	Explanation string `json:"explanation"`
}
