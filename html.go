// html.go contains functions concerned with rendering web pages.
package main

// getPageForKey returns a rendered HTML page containing the contents
// of the paste identified by key. Any errors encountered while trying
// to open the paste file are returned.
func getPageForKey(pathToStore, key string) (string, error) {
	text, textErr := getTextForKey(pathToStore, key)
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
		"<input type=\"textarea\" name=\"paste\">" +
		"<input type=\"submit\" value=\"Save\">" +
		"</form>" +
		"</body>"
	return page
}
