package main

import "strings"

type ArtBuilder struct {
	t, s string
}

func NewArtBuilder() *ArtBuilder {
	return &ArtBuilder{"", "normal"}
}

func (a *ArtBuilder) AddText(t string) *ArtBuilder {
	a.t += t
	return a
}

func (a *ArtBuilder) SetStyle(s string) *ArtBuilder {
	if s != "normal" && s != "bold" && s != "italic" && s != "outline" {
		panic("invalid style")
	}
	a.s = s
	return a
}

func (a *ArtBuilder) Build() string {
	r := make([]string, 8)

	for i := 0; i < 8; i++ {
		x := a.t

		if a.s == "bold" {
			x = strings.ReplaceAll(x, "", " ")
		}
		if a.s == "italic" {
			x = strings.Repeat(" ", i+1) + "/" + x
		}
		if a.s == "outline" {
			x = "[" + x + "]"
		}

		r[i] = x
	}

	return strings.Join(r, "\n")
}
