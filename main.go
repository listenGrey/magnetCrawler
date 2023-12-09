package main

import (
	"javCrawler/model"
	"javCrawler/persist/serve"
)

func main() {
	profile := model.Profile{
		Url:         "test",
		Id:          "test",
		Code:        "test",
		Title:       "test",
		Cover:       "test",
		ReleaseData: "test",
		Duration:    "test",
		Director:    "test",
		Company: model.Tag{
			Content: "test",
			Url:     "test",
		},
		Series: model.Tag{
			Content: "test",
			Url:     "test",
		},
		Stars:       "test",
		StarsPerson: "test",
		Type: []model.Tag{
			model.Tag{
				Content: "test1",
				Url:     "test1",
			}, model.Tag{
				Content: "test2",
				Url:     "test2",
			},
		},
		Actor: []model.Tag{
			model.Tag{
				Content: "test1",
				Url:     "test1",
			}, model.Tag{
				Content: "test2",
				Url:     "test2",
			},
		},
		FetchTime: "test",
		IntroPics: []string{
			"test1",
			"test2",
		},
		Magnets: []model.Magnet{
			model.Magnet{
				Content: "test1",
				Url:     "test1",
				Data:    "test1",
			}, model.Magnet{
				Content: "test2",
				Url:     "test2",
				Data:    "test2",
			},
		},
	}
	add := serve.ListenPort("192.168.43.82:8964")
	serve.SaveItem(profile, add)
}
