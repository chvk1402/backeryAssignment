package config

var Code = map[string]string{
	"Vegemite Scroll":  "VS5",
	"Blueberry Muffin": "MB11",
	"Croissant":        "CF",
}

var PriceMatrix = map[string]map[int]float32{
	"VS5": {
		3: 6.99,
		5: 8.99,
	},
	"MB11": {
		2: 9.95,
		5: 16.95,
		8: 24.95,
	},
	"CF": {
		3: 5.95,
		5: 9.95,
		9: 16.99,
	},
}
