package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type document struct {
	Types map[string]struct {
		ID                    string        `json:"id"`
		Title                 string        `json:"title"`
		Pid                   string        `json:"pid"`
		DefaultModificationID string        `json:"default_modification_id"`
		Videos                []interface{} `json:"videos"`
		Modifications         map[string]struct {
			ID            string      `json:"id"`
			Pid           string      `json:"pid"`
			Type          string      `json:"type"`
			CartSliceType string      `json:"cart_slice_type"`
			CategoryID    string      `json:"category_id"`
			ProductID     string      `json:"product_id"`
			ProductClass  interface{} `json:"product_class"`
			ProductType   string      `json:"product_type"`
			Prices        struct {
				NormalPrice string `json:"normal_price"`
				SellPrice   string `json:"sell_price"`
			} `json:"prices"`
			Discounts struct {
				PromoDiscount     string `json:"promo_discount"`
				OurDiscount       string `json:"our_discount"`
				PromoSellDiscount string `json:"promo_sell_discount"`
				IsNew             string `json:"is_new"`
				SellPrice         string `json:"sell_price"`
				VendorID          string `json:"vendor_id"`
				Status            string `json:"status"`
				PageType          string `json:"pageType"`
			} `json:"discounts"`
		} `json:"modifications"`
		Defaults struct {
			Vendor string `json:"vendor"`
		} `json:"defaults"`
	} `json:"types"`
	Defaults struct {
		ColorID string `json:"colorId"`
	} `json:"defaults"`
	IsComparable   bool   `json:"isComparable"`
	AddedToCompare bool   `json:"addedToCompare"`
	Status         string `json:"status"`
}

func main() {
	if err := process(); err != nil {
		log.Fatal(err)
	}
}

func process() error {
	var doc document
	if err := json.Unmarshal(jsonStr, &doc); err != nil {
		return err
	}
	fmt.Printf("%+v\n", doc)

	for tKey, t := range doc.Types {
		fmt.Println("Type key", tKey)
		fmt.Println("Type ID", t.ID)

		fmt.Println("Type Modifications", t.Modifications)
		for mKey, m := range t.Modifications {
			fmt.Println("Modification key", mKey)
			fmt.Println("Modification ID", m.ID)
		}
	}

	return nil
}

var jsonStr = []byte(`{
    "types": {
        "15590": {
            "id": "15590",
            "title": "",
            "pid": "1",
            "default_modification_id": "232385",
            "videos": [],
            "modifications": {
                "232385": {
                    "id": "232385",
                    "pid": "1",
                    "type": "default",
                    "cart_slice_type": "0",
                    "category_id": "150",
                    "product_id": "89876",
                    "product_class": null,
                    "product_type": "product",
                    "prices": {
                        "normal_price": "176.99",
                        "sell_price": "160.00"
                    },
                    "discounts": {
                        "promo_discount": "9",
                        "our_discount": "6",
                        "promo_sell_discount": "16",
                        "is_new": "0",
                        "sell_price": "160.00",
                        "vendor_id": "1",
                        "status": "active",
                        "pageType": "product"
                    }
                }
            },
            "defaults": {
                "vendor": "1"
            }
        }
    },
    "defaults": {
        "colorId": "555"
    },
    "isComparable": true,
    "addedToCompare": false,
    "status": "active"
}`)
