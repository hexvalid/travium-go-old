package mari

import "regexp"

var regexWindowData = regexp.MustCompile("window.__data='(.*?)'")
