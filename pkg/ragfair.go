package pkg

import (
	"MT-GO/data"
	"fmt"
)

func GetFlea(handbookId string) (data.Flea, error) {
	flea := data.GetFlea().Market

	if handbookId == "" {
		return flea, nil
	}

	if main, err := data.HasGetMainHandbookCategory(handbookId); err == nil {
		for _, mainValue := range main {
			if sub, err := data.HasGetHandbookSubCategory(mainValue); err == nil {
				for _, subValue := range sub {
					catalog, err := data.GetFleaCatalog(subValue)
					if err != nil {
						continue
					}
					flea.Offers = append(flea.Offers, catalog...)
				}
			}
		}
		return flea, nil
	}

	if sub, err := data.HasGetHandbookSubCategory(handbookId); err == nil {
		for _, value := range sub {
			catalog, err := data.GetFleaCatalog(value)
			if err != nil {
				continue
			}
			flea.Offers = append(flea.Offers, catalog...)
		}
		return flea, nil
	}

	catalog, err := data.GetFleaCatalog(handbookId)
	if err != nil {
		return flea, fmt.Errorf("handbookID %s invalid, flea not populated", handbookId)
	}
	flea.Offers = append(flea.Offers, catalog...)
	return flea, nil

}
