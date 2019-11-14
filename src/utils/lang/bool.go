package lang

//BoolIncludes 布尔包含判断
func BoolIncludes(arr []bool, item bool) bool {
	for _, a := range arr {
		if a == item {
			return true
		}
	}
	return false
}
