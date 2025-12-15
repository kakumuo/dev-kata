package gildedrose_test

import (
	"testing"

	"github.com/emilybache/gildedrose-refactoring-kata/gildedrose"
)

func TestGeneric(t *testing.T) {
	var items = []*gildedrose.Item{
		{"Generic item", 1, 3},
	}

	gildedrose.UpdateQuality(items)

	if items[0].SellIn != 0 {
		t.Error("Expected item sellin to be zero")
	} else if items[0].Quality >= 3 {
		t.Error("Expected item quality to have decreased")
	}

	gildedrose.UpdateQuality(items)

	if items[0].SellIn >= 0 {
		t.Error("Expected item sellin to be less than zero")
	} else if items[0].Quality > 0 {
		t.Error("Expected item quality to have decreased by two")
	}

	gildedrose.UpdateQuality(items)

	if items[0].Quality != 0 {
		t.Error("Expected item quality to have decreased by two")
	}

	gildedrose.UpdateQuality(items)

	if items[0].Quality < 0 {
		t.Error("Expected item quality to remain positive")
	}

}

func TestAgedBrie(t *testing.T) {
	var items = []*gildedrose.Item{
		{"Aged Brie", 2, 48},
	}

	gildedrose.UpdateQuality(items)
	gildedrose.UpdateQuality(items)

	if items[0].Quality != 50 {
		t.Error("Expected quality to be 50")
	} else if items[0].SellIn != 0 {
		t.Error("Expected sellin date to be lowered")
	}

	gildedrose.UpdateQuality(items)
	gildedrose.UpdateQuality(items)

	if items[0].Quality != 50 {
		t.Error("Expected quality to not be past 50")
	} else if items[0].SellIn >= 0 {
		t.Error("Expected sellin date to be negative")
	}
}

func TestLegendary(t *testing.T) {
	var items = []*gildedrose.Item{
		{"Sulfuras, Hand of Ragnaros", 10, 80},
	}

	for i := 0; i < 10; i++ {
		gildedrose.UpdateQuality(items)
	}

	if items[0].Quality != 80 {
		t.Error("Expected item quality to be 80")
	} else if items[0].SellIn != 10 {
		t.Error("Expected item quality to not have changed")
	}
}

func TestPasses(t *testing.T) {
	var items = []*gildedrose.Item{
		{"Backstage passes to a TAFKAL80ETC concert", 15, 30},
	}
	item := items[0]

	for i := 0; i < 5; i++ {
		gildedrose.UpdateQuality(items)
	}

	if item.SellIn != 10 {
		t.Error("Expected quality to be 10")
	} else if item.Quality != 35 {
		t.Error("Expect quality to have increased by 5")
	}

	for i := 0; i < 5; i++ {
		gildedrose.UpdateQuality(items)
	}

	if item.SellIn != 5 {
		t.Error("Expected quality to be 5")
	} else if item.Quality != 45 {
		t.Error("Expect quality to have increased by 5")
	}

	gildedrose.UpdateQuality(items)

	if item.Quality != 48 {
		t.Error("Expect quality to have increased by 3")
	}

	gildedrose.UpdateQuality(items)

	if item.Quality != 50 {
		t.Error("Expect quality to not exceed 50")
	}

	gildedrose.UpdateQuality(items)
	gildedrose.UpdateQuality(items)
	gildedrose.UpdateQuality(items)
	gildedrose.UpdateQuality(items)

	if item.Quality != 0 {
		t.Error("Expect quality to drop to zero")
	}
}

func TestConjured(t *testing.T) {
	var items = []*gildedrose.Item{
		{"Conjured Mana Cake", 4, 20},
	}
	item := items[0]

	for i := 0; i < 4; i++ {
		gildedrose.UpdateQuality(items)
	}

	if item.Quality != 12 {
		t.Errorf("Expected quality to be 12, got %d", item.Quality)
	}

	for i := 0; i < 2; i++ {
		gildedrose.UpdateQuality(items)
	}

	if item.Quality != 4 {
		t.Errorf("Expected quality to be 4, got %d", item.Quality)
	}

	for i := 0; i < 2; i++ {
		gildedrose.UpdateQuality(items)
	}

	if item.Quality != 0 {
		t.Errorf("Expected quality to be 0, got %d", item.Quality)
	}

}
