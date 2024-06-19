package tasks

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/vv198x/GoWB/models"
	"github.com/vv198x/GoWB/requests"
	"net/http"
	"net/url"
)

const uriSearch = "https://search.wb.ru/exactmatch/ru/common/v5/search?ab_testing=false&appType=32&curr=rub&dest=-364001&page=1&resultset=catalog&sort=popular&spp=32&suppressSpellcheck=false"

//функция с повтором

func GetPosition(ctx context.Context, query string) error {
	finalURL := fmt.Sprintf(`%s&query=%s`, uriSearch, url.QueryEscape(query))

	var wbSearch models.WbSearch

	if data, err := requests.New(http.MethodGet, finalURL, nil).DoUnauthorizedRequest(ctx); err == nil {
		if err = json.Unmarshal(data, &wbSearch); err != nil {
			return fmt.Errorf("json error %V", err)
		}
	} else {
		return fmt.Errorf("request GetPosition error %V", err)
	}

	if len(wbSearch.Data.Products) < 50 {
		return fmt.Errorf("less than 50 products")
	}

	fmt.Println("Requst:", wbSearch.Metadata.Name)
	fmt.Println("Total Products:", len(wbSearch.Data.Products))
	/*
		fmt.Println("-----------------")
		for i, product := range wbSearch.Data.Products {
			if product.Brand == "Livelyflow" {
				fmt.Println("Product:", product.Name)
				fmt.Println("SKU:", product.Id)
				fmt.Println("CPM:", product.Log.Cpm)
				fmt.Println("Position:", product.Log.Position)
				fmt.Println("Promo Position:", product.Log.PromoPosition)
				fmt.Println("Count Position:", i+1)
				fmt.Println("-----------------")
			}

		}

	*/
	// Глючное апи не всегда выдаёт log и 100 позиций
	// если позиций меньше 50 не верить
	// получить товары в компании в Апдейт
	return nil
}
