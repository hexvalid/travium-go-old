package mari

import "regexp"

const hostUrl = "https://www.travian.com"

var regexWindowData = regexp.MustCompile("window.__data='(.*?)'")
