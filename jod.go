package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

const url = "https://api.jokes.one/jod"

type jsonResponse struct {
	Contents contents `json:"contents"`
}

func (rj jsonResponse) String() string {
	return rj.Contents.String()
}

type contents struct {
	Jokes []jokes `json:"jokes"`
}

func (c contents) String() string {
	return c.Jokes[0].String()
}

type jokes struct {
	Joke joke `json:"joke"`
}

func (js jokes) String() string {
	return js.Joke.String()
}

type joke struct {
	Text string `json:"text"`
}

func (j joke) String() string {
	return j.Text
}

func jod() (msg string, err error) {
	r, err := http.Get(url)
	if err != nil {
		return "", err
	}

	rd, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return "", err
	}

	j := &jsonResponse{}
	err = json.Unmarshal(rd, j)

	if err != nil {
		return "", err
	}

	return j.String(), nil
}
