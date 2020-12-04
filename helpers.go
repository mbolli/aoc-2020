package main

func indexOf(element string, data []string) int {
	for k, v := range data {
		if element == v {
			return k
		}
	}
	return -1 //not found.
}

func contains(element string, data []string) bool {
	return indexOf(element, data) != -1
}
