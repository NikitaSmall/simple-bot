package quoter

var bibler = Quoter{
	url:          "http://www.sandersweb.net/bible/verse.php",
	query:        "div.esv-text p",
	fromEncoding: "utf-8",
}

func GetBibleQuote() (string, error) {
	return bibler.getPageResult()
}
