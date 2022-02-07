package main

// https://stackoverflow.com/questions/47565090/need-help-to-extract-comments-from-c-file
func formatString(content string) string {
	rs := []rune(content)
	rsnew := []rune{}

	inmulti := false  /* in multi-line comment flag */
	insingle := false /* in single-line comment flag */
	insquo := false   /* within single-quotes */
	indquo := false   /* within double-quotes */

	prev := rs[0] // first

	for pos := 1; pos < len(rs); pos++ {
		onechar := rs[pos]

		switch onechar {
		case '/':
			if insquo || indquo {

			} else {
				if insingle {
					// "// /" <== the last
					rsnew = append(rsnew, onechar)
				} else {
					if prev == '/' {
						insingle = true
					} else {
						// not comment
					}
				}
				if inmulti {
					if prev == '*' {
						inmulti = false
					} else {
						// "arch/alpha/boot/bootpz.c" in multi-line comments
						rsnew = append(rsnew, onechar)
					}
				} else {
					// not comment
				}
			}

			break
		case '*':
			if insquo || indquo {
				// within quotes, no oper
			} else {
				if insingle {
					// "// *" <== the last
					rsnew = append(rsnew, onechar)
				} else if inmulti {
					// "/* * **/"
					var nextChar rune = 0
					if pos < len(rs)-1 {
						nextChar = rs[pos+1]
					}

					if nextChar == '/' {

					} else {
						rsnew = append(rsnew, onechar)
					}
				} else {
					if prev == '/' {
						inmulti = true
					} else {
						// not comment
					}
				}
			}

			if !(insquo || indquo) {
				if prev == '/' {
					inmulti = true
				} else {

				}
			} else {
				// within quotes, no oper
			}

			break
		case '\n':
			if insingle {
				insingle = false
				rsnew = append(rsnew, onechar)
			}
			if inmulti {
				rsnew = append(rsnew, onechar)
			}
			break
		case '\'':
			if insingle || inmulti {
				// within comment
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
