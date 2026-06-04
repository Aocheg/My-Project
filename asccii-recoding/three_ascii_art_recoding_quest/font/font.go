package main

import "unicode"

func GenerateFont() map[rune][]string {
	f := map[rune][]string{}

	for ch := rune(32); ch <= 126; ch++ {
		g := make([]string, 8)

		for y := 0; y < 8; y++ {
			r := make([]rune, 8)

			for x := 0; x < 8; x++ {
				c := ' '

				if ch != ' ' {
					if x == 0 || x == 7 || y == 0 || y == 7 {
						c = '.'
					}

					if unicode.IsLetter(ch) &&
						(ch == 'A' || ch == 'E' || ch == 'I' || ch == 'O' || ch == 'U' ||
							ch == 'a' || ch == 'e' || ch == 'i' || ch == 'o' || ch == 'u') {
						if (x+y+int(ch))%2 == 0 {
							c = '*'
						}
					} else if (x*y+int(ch))%5 == 0 {
						c = '*'
					}

					if (ch == 'g' || ch == 'j' || ch == 'p' || ch == 'q' || ch == 'y') && y > 5 {
						c = '*'
					}
				}

				r[x] = c
			}

			g[y] = string(r)
		}

		f[ch] = g
	}

	return f
}
