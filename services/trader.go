package services

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
