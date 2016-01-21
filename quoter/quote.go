package quoter

func GetRandomQuote() string {
	return getPageResult("http://bash.im/random", ".quote .text")
}
