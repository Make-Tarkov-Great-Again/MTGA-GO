package pkg

import (
	"MT-GO/data"
	"fmt"
	"net/http"
)

// GetCorrectAmountOfItemsPurchased returns a new slice of which each index
// represents a new item and the value of that index is the StackObjectsCount
// of that item
func GetCorrectAmountOfItemsPurchased(amountPurchased int32, itemStackSize int32) []int32 {
	howManyItems := amountPurchased / itemStackSize
	remainder := amountPurchased % itemStackSize
	var stackSlice []int32
	if remainder != 0 {
		stackSlice = make([]int32, 0, howManyItems+1)
		for i := int32(0); i < howManyItems; i++ {
			stackSlice = append(stackSlice, itemStackSize)
		}
		stackSlice = append(stackSlice, remainder)
	} else {
		stackSlice = make([]int32, 0, howManyItems)
		for i := int32(0); i < howManyItems; i++ {
			stackSlice = append(stackSlice, itemStackSize)
		}
	}

	return stackSlice
}

func GetSuitesStorage(sessionID string) (map[string]any, error) {
	storage, err := data.GetStorageByID(sessionID)
	if err != nil {
		return nil, err
	}

	return map[string]any{
		"_id":    sessionID,
		"suites": storage.Suites,
	}, nil
}

func GetTraderSettings() []*data.TraderBase {
	traders := data.GetTraders()
	output := make([]*data.TraderBase, 0, len(traders))

	for _, trader := range traders {
		output = append(output, trader.Base)
	}

	return output
}

func GetTraderSuits(id string) ([]data.TraderSuits, error) {
	trader, err := data.GetTraderByUID(id)
	if err != nil {
		return nil, err
	}

	if trader.Suits != nil {
		return trader.Suits, nil
	}

	return nil, fmt.Errorf("Trader %s suits does not exist", id)
}

func GetTraderAssort(r *http.Request) (*data.Assort, error) {
	trader, err := data.GetTraderByUID(r.URL.Path[36:])
	if err != nil {
		return nil, err
	}

	character := data.GetCharacterByID(GetSessionID(r))
	assort, err := trader.GetStrippedAssort(character)
	if err != nil {
		return nil, err
	}

	return assort, nil
}
