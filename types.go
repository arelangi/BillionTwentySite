package main

type Feed struct {
	Rss struct {
		Channel struct {
			Title     string   `json:"title"`
			Copyright []string `json:"copyright"`
			Image     []struct {
				Href  string `json:"-href,omitempty"`
				URL   string `json:"url,omitempty"`
				Title string `json:"title,omitempty"`
				Link  string `json:"link,omitempty"`
			} `json:"image"`
			Item      []Episode     `json:"item"`
			Link      []interface{} `json:"link"`
			PubDate   string        `json:"pubDate"`
			Language  string        `json:"language"`
			WebMaster string        `json:"webMaster"`
			Info      struct {
				URI string `json:"-uri"`
			} `json:"info"`
			Keywords      []string      `json:"keywords"`
			LastBuildDate string        `json:"lastBuildDate"`
			TTL           string        `json:"ttl"`
			Description   []interface{} `json:"description"`
			Author        string        `json:"author"`
			Category      []struct {
				Content  string `json:"#content,omitempty"`
				Scheme   string `json:"-scheme,omitempty"`
				Text     string `json:"-text,omitempty"`
				Category struct {
					Text string `json:"-text"`
				} `json:"category,omitempty"`
			} `json:"category"`
			Owner struct {
				Email string `json:"email"`
				Name  string `json:"name"`
			} `json:"owner"`
			Rating    string `json:"rating"`
			Subtitle  string `json:"subtitle"`
			Explicit  string `json:"explicit"`
			Thumbnail struct {
				URL string `json:"-url"`
			} `json:"thumbnail"`
			Summary   string `json:"summary"`
			FeedFlare []struct {
				Content string `json:"#content"`
				Href    string `json:"-href"`
				Src     string `json:"-src"`
			} `json:"feedFlare"`
			Credit struct {
				Content string `json:"#content"`
				Role    string `json:"-role"`
			} `json:"credit"`
		} `json:"channel"`
		Itunes     string `json:"-itunes"`
		Atom       string `json:"-atom"`
		Media      string `json:"-media"`
		Feedburner string `json:"-feedburner"`
		Version    string `json:"-version"`
	} `json:"rss"`
}

type Episode struct {
	Summary     string   `json:"summary"`
	Description string   `json:"description"`
	Keywords    string   `json:"keywords"`
	Title       string   `json:"title"`
	Explicit    string   `json:"explicit"`
	OrigLink    string   `json:"origLink"`
	PubDate     string   `json:"pubDate"`
	Author      []string `json:"author"`
	Image       struct {
		Href string `json:"-href"`
	} `json:"image"`
	Duration  string `json:"duration"`
	Enclosure struct {
		Length string `json:"-length"`
		Type   string `json:"-type"`
		URL    string `json:"-url"`
	} `json:"enclosure"`
	Subtitle string `json:"subtitle"`
	Content  struct {
		URL      string `json:"-url"`
		FileSize string `json:"-fileSize"`
		Type     string `json:"-type"`
	} `json:"content,omitempty"`
	GUID struct {
		Content     string `json:"#content"`
		IsPermaLink string `json:"-isPermaLink"`
	} `json:"guid"`
	Link string `json:"link"`
}
