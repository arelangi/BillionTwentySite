package main

func getEpisodes(feed Feed) (code string) {

	episodes := feed.Rss.Channel.Item[1:]

	code = `<!-- ===== MAIN ===== -->
	<main id="main" class="main">
		<!-- ===== PODCAST LIST ===== -->
		<section id="#episodes" class="section-positive">
			<div class="container">
				
				<!-- ===== SECTION TITLE ===== -->
				<h2 class="title-default">Episodes</h2>
				
			<div class="row">`

	for _, v := range episodes {

		code += `<!-- ===== PODCAST CARD FULL ===== -->
					<div class="col-sm-12 mb-40">
						<div class="podcast-card full">
							<figure class="podcast-image"><a href="` + v.Link + `"><img src="assets/img/Logo555.png" alt="Podcast Logo" title="Int The End Episode" /></a></figure>
							<div class="podcast-content">
								<span class="podcast-date">` + v.PubDate[5:16] + `</span>
								<h2 class="podcast-title"><a href="` + v.Link + `">` + v.Title + `</a></h2>
								<p class="podcast-excerpt"><a href="` + v.Link + `">` + v.Description + `</a></p>
								<ul class="podcast-meta">
					                <li class="item"><a href="#" class="podcast-tag" rel="tag">Pop Culture</a></li>
					                <li class="item"><i class="fa fa-clock-o"></i> ` + v.Duration + ` mins.</li>
					            </ul>
							</div>
						</div>
					</div>`

	}

	code += `</div>
			</div>
		</section>
	</main>`

	return
}
