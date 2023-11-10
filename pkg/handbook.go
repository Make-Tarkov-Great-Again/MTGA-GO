package pkg

var currencyName = map[string]string{
	"RUB": "5449016a4bdc2d6f028b456f",
	"EUR": "569668774bdc2da2298b4568",
	"USD": "5696686a4bdc2da3298b456a",
}

var currencyByID = map[string]*struct{}{
	"5449016a4bdc2d6f028b456f": nil, //RUB
	"569668774bdc2da2298b4568": nil, //EUR
	"5696686a4bdc2da3298b456a": nil, //USD
}

func IsCurrencyByID(UID string) bool {
	_, ok := currencyByID[UID]
	return ok
}

func GetCurrencyByName(name string) *string {
	currency, ok := currencyName[name]
	if ok {
		return &currency
	}
	return nil
}
