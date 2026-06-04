package main

func MergeBanners(base map[rune][]string, priority map[rune][]string) map[rune][]string {
	m1 := base
	m2 := priority
	m3 := map[rune][]string{}

	for k, v := range m1 {
		m3[k] = v
	}

	for k, v := range m2 {
		m3[k] = v
	}
	return m3
}
