package main

func ellipsis(title string) (text string) {
	text = title
	if len(title) > 55 {
		text = title[:55] + "..."
	}
	return
}
