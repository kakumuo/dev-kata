package gildedrose

import "strings"

type Item struct {
	Name            string
	SellIn, Quality int
}

type ItemType int

const (
	IT_Generic ItemType = iota
	IT_Aged
	IT_Legendary
	IT_Pass
	IT_Conjured
)

func GetItemType(item *Item) ItemType {
	name := strings.ToLower(item.Name)

	if strings.Contains(name, "aged brie") {
		return IT_Aged
	} else if strings.Contains(name, "sulfuras") {
		return IT_Legendary
	} else if strings.Contains(name, "backstage pass") {
		return IT_Pass
	} else if strings.Contains(name, "conjured") {
		return IT_Conjured
	}

	return IT_Generic
}

const MAX_QUALITY, MIN_QUALITY = 50, 0

func UpdateQuality(items []*Item) {
	for _, item := range items {
		itemType := GetItemType(item)

		if itemType != IT_Legendary {
			item.SellIn -= 1
		}

		switch itemType {
		case IT_Legendary:
		case IT_Aged:
			item.Quality += 1
		case IT_Conjured:
			item.Quality -= 2
			if item.SellIn < 0 {
				item.Quality -= 2
			}
		case IT_Pass:
			if item.SellIn < 0 {
				item.Quality = 0
			} else if item.SellIn < 5 {
				item.Quality += 3
			} else if item.SellIn < 10 {
				item.Quality += 2
			} else {
				item.Quality += 1
			}
		case IT_Generic:
			item.Quality -= 1
			if item.SellIn < 0 {
				item.Quality -= 1
			}
		}

		if itemType != IT_Legendary {
			item.Quality = max(min(item.Quality, MAX_QUALITY), MIN_QUALITY)
		}
	}

}
