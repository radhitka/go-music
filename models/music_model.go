package models

type Music struct {
	ID          int
	Title       string
	Artist      string
	IsPublished bool
}

var Musics = []Music{
	{
		ID:     1,
		Title:  "Sudah",
		Artist: "Ardhito Pramodu",
	},
	{
		ID:     2,
		Title:  "Muak",
		Artist: "Aruma",
	},
	{
		ID:     3,
		Title:  "Tak Segampang Itu",
		Artist: "Anggi Marito",
	},
}
