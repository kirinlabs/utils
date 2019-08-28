package dict

func Delete(src map[string]interface{}, args ...string) {
	for _, v := range args {
		delete(src, v)
	}
}
