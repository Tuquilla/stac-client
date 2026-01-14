package models

import "slices"

type Collection struct {
	Id    string `json:"id"`
	Title string `json:"title"`
	Links []Link `json:"links"`
}

type Collections struct {
	Collections []Collection `json:"collections"`
}

type Link struct {
	Href string `json:"href"`
	Rel  string `json:"rel"`
}

func (collection Collection) GetItemsLink() *Link {
	index := slices.IndexFunc(collection.Links, func(links Link) bool {
		return links.Rel == "items"
	})

	return &collection.Links[index]
}
