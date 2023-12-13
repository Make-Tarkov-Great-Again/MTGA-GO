package pkg

import (
	"MT-GO/data"
	"fmt"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"sort"
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
	output := make([]*data.TraderBase, 0, traders.Len())

	traders.ForEach(func(_ string, trader *data.Trader) bool {
		output = append(output, trader.Base)
		return true
	})

	sort.SliceStable(output, func(i, j int) bool {
		return output[i].ID < output[j].ID
	})
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

	return nil, fmt.Errorf("trader %s suits does not exist", id)
}

func GetTraderAssort(r *http.Request) (*data.Assort, error) {
	character, err := data.GetCharacterByID(GetSessionID(r))
	if err != nil {
		log.Fatalln(err)
	}
	trader, err := data.GetTraderByUID(chi.URLParam(r, "id"))
	if err != nil {
		return nil, err
	}
	assort, err := trader.GetStrippedAssort(character)
	if err != nil {
		return nil, err
	}

	return assort, nil
}
