// html.go contains functions concerned with rendering web pages.
package main

// getPageForKey returns a rendered HTML page containing the contents
// of the paste identified by key. Any errors encountered while trying
// to open the paste file are returned.
func getPageForKey(key string) (string, error) {
	text, textErr := getTextForKey(key)
	if textErr != nil {
		return "", textErr
	}

	escapedTextString := escape(text)

	page := "<!DOCTYPE html>\n" +
		"<head>\n\t" +
		"<meta charset=\"UTF-8\">\n\t" +
		"<title>GoPaste</title>\n" +
		"</head>\n\n" +
		"<body>\n\t" +
		"<pre><tt>" +
		escapedTextString +
		"\t</tt></pre>\n" +
		"</body>"

	return page, nil
}

func getNewPastePage() string {
	page := "<!DOCTYPE html>\n" +
		"<head>\n\t" +
		"<meta charset=\"UTF-8\">\n\t" +
		"<title>GoPaste</title>\n" +
		"</head>\n\n" +
		"<body>\n\t" +
		"<form action=\"/\" method=\"post\">" +
		"<textarea autofocus required rows=\"32\" cols=\"80\" name=\"paste\"></textarea>" +
		"<br>" +
		"<input type=\"submit\" value=\"Save\">" +
		"</form>" +
		"</body>"
	return page
}
