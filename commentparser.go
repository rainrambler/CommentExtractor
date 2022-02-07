package main

// https://stackoverflow.com/questions/47565090/need-help-to-extract-comments-from-c-file
func formatString(content string) string {
	rs := []rune(content)
	rsnew := []rune{}

	inmulti := false
	insingle := false
	insquo := false
	indquo := false

	prev := rs[0] // first

	for _, onechar := range rs {
		switch onechar {
		case '/':
			if (prev == '/') && !(insquo || indquo) {
				insingle = true
			}
			if (prev == '*') && !(insquo || indquo) {
				inmulti = false

				rsnew = append(rsnew, '\n')
			}
			// "arch/alpha/boot/bootpz.c" in multi-line comments
			if inmulti {
				rsnew = append(rsnew, onechar)
			}
			break
		case '*':
			// "*" in multi-line comments
			if inmulti {
				rsnew = append(rsnew, onechar)
			}

			if (prev == '/') && !(insquo || indquo) {
				inmulti = true
			}

			break
		case '\n':
			insingle = false
			if insingle || inmulti {
				rsnew = append(rsnew, onechar)
			}
			break
		case '\'':
			if insingle || inmulti {

			} else {
				insquo = !insquo
			}

			break
		case '"':
			insquo = !insquo
			break
		default:
			if (insingle || inmulti) && !(insquo || indquo) {
				rsnew = append(rsnew, onechar)
			}
			break
		}
		prev = onechar
	}

	return string(rsnew)
}
