package main

func getMainContent(feed Feed) (code string) {
	latest := feed.Rss.Channel.Item[0]
	first := feed.Rss.Channel.Item[1]
	second := feed.Rss.Channel.Item[2]

	code = `<!-- ===== MAIN ===== -->
    <main id="main" class="main">
        <!-- ===== PODCAST LIST ===== -->
        <section id="#episodes" class="section-positive">
            <div class="container">
                <!-- ===== SECTION TITLE ===== -->
                <h2 class="title-default">Other Episodes</h2>
                <div class="row">
                    <!-- ===== PODCAST CARD FULL ===== -->
                    <div class="col-sm-12 mb-40">
                        <div class="podcast-card full">
                            <figure class="podcast-image"><a href="`

	code += latest.Link
	code += `"><img src="assets/img/Logo555.png" alt="Podcast Logo" title="Podcast Logo" /></a></figure>
                            <div class="podcast-content">
                                <span class="podcast-date">`

	code += latest.PubDate[5:16] + `</span>
                                <h2 class="podcast-title"><a href="`

	code += latest.Link + `">` + latest.Title
	code += `</a></h2><ul class="podcast-meta">
                                    <li class="item"><a href="#" class="podcast-tag" rel="tag">Pop Culture</a></li>
                                    <li class="item"><i class="fa fa-clock-o"></i> ` + latest.Duration + ` mins.</li>
                                </ul>
                            </div>
                        </div>
                    </div>
                    
                    <!-- ===== PODCAST CARD BOXED ===== -->
                    <div class="col-sm-6 mb-40">
                        <div class="podcast-card boxed">
                            <figure class="podcast-image"><a href="` + first.Link + `"><img src="assets/img/Logo555.png" alt="Podcast Logo" title="Podcast Logo" /></a></figure>
                            <div class="podcast-content">
				<span class="podcast-date">` + first.PubDate[5:16] + `</span>
                                <h2 class="podcast-title"><a href="` + first.Link + `">` + ellipsis(first.Title) + `</a></h2>
                                <ul class="podcast-meta">
                                    <li class="item"><a href="#" class="podcast-tag" rel="tag">Pop Culture</a></li>
                                    <li class="item"><i class="fa fa-clock-o"></i> ` + first.Duration + ` mins.</li>
                                </ul>
                            </div>
                        </div>
                    </div>

                    <!-- ===== PODCAST CARD BOXED ===== -->
                    <div class="col-sm-6 mb-40">
                        <div class="podcast-card boxed">
                            <figure class="podcast-image"><a href="` + second.Link + `"><img src="assets/img/Logo555.png" alt="Podcast Logo" title="Podcast Logo" /></a></figure>
                            <div class="podcast-content">
				<span class="podcast-date">` + second.PubDate[5:16] + `</span>
                                <h2 class="podcast-title"><a href="` + second.Link + `">` + ellipsis(second.Title) + `</a></h2>
                               <ul class="podcast-meta">
                                    <li class="item"><a href="#" class="podcast-tag" rel="tag">Pop Culture</a></li>
                                    <li class="item"><i class="fa fa-clock-o"></i> ` + second.Duration + ` mins.</li>
                                </ul>
                            </div>
                        </div>
                    </div>
                                       
                    <!-- ===== CHECK MORE ===== -->
                    <div class="col-sm-12 mb-50">
                        <a href="intheend.html" class="btn btn-primary btn-block btn-lg">View more episodes</a>
                    </div>
                    
                </div>
                
            </div>
        </section>`

	return
}
