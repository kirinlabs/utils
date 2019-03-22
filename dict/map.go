package dict

/**
* Map
**/
func HasKey(src map[string]interface{}, key string) bool {
	if _, ok := src[key]; ok {
		return true
	}
	return false
}

func Delete(src map[string]interface{}, args ...string) {
	for _, v := range args {
		delete(src, v)
	}
}

func Keys(src map[string]interface{}) []string {
	res := make([]string, 0)
	for k, _ := range src {
		res = append(res, k)
	}
	return res
}
