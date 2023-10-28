package tools

/*
GetItemFamilyTree returns the family of an item based on parentID if it and the family exists
*/
func GetItemFamilyTree(items []any, parent string) []string {
	var list []string

	for _, childItem := range items {
		child := childItem.(map[string]any)

		if child["parentId"].(string) == parent {
			list = append(list, GetItemFamilyTree(items, child["_id"].(string))...)
		}
	}

	list = append(list, parent) // required
	return list
}
