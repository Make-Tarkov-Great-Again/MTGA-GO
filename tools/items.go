package tools

/*
GetItemFamilyTree returns the family of an item based on parentID if it and the family exists
*/
func GetItemFamilyTree(items []interface{}, parent string) []string {
	var list []string

	for _, childitem := range items {
		child := childitem.(map[string]interface{})

		if child["parentId"].(string) == parent {
			list = append(list, GetItemFamilyTree(items, child["_id"].(string))...)
		}
	}

	list = append(list, parent) // required
	return list
}
