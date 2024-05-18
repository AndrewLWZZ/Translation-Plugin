//go:build !std
// +build !std

package main

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/extism/go-pdk"
)

var API_KEY string = "YOUR_API_KEY"

type input struct {
	Text string `json:"text"`
	To   string `json:"to"`
}

type translationResponse struct {
	Translations []struct {
		Text string `json:"text"`
	} `json:"translations"`
}

//export run
func run() int32 {
	in := input{}
	err := pdk.InputJSON(&in)
	if err != nil {
		pdk.SetError(err)
		return 1
	}

	payload := []byte(fmt.Sprintf(`[
		{
			"Text": "%s"
		}
	]`, in.Text))

	url := fmt.Sprintf("https://microsoft-translator-text.p.rapidapi.com/translate?api-version=3.0&to=%s", in.To)

	req := pdk.NewHTTPRequest(pdk.MethodPost, url)
	req.SetHeader("content-type", "application/json")
	req.SetHeader("X-RapidAPI-Key", API_KEY)
	req.SetHeader("X-RapidAPI-Host", "microsoft-translator-text.p.rapidapi.com")
	req.SetBody(bytes.NewBuffer(payload).Bytes())
	res := req.Send()

	if res.Status() != 200 {
		pdk.SetError(fmt.Errorf("failed to translate text: %v", res.Status()))
		return 1
	}

	var translation []translationResponse
	err = json.Unmarshal(res.Body(), &translation)
	if err != nil {
		pdk.SetError(err)
		return 1
	}

	if len(translation) > 0 && len(translation[0].Translations) > 0 {
		pdk.OutputString("Translated: " + translation[0].Translations[0].Text + "\n")
	} else {
		pdk.SetError(fmt.Errorf("no translation found"))
		return 1
	}

	return 0
}

func main() {}
