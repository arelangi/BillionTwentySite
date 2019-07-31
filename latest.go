package main

import (
	"strings"
)

func getLatest(feed Feed) (code string) {
	latest := feed.Rss.Channel.Item[0]

	code = `<!-- ===== LATEST PODCAST (HERO) ===== -->
    <section class="podcast-hero" style="background-image: url(assets/img/culture.png)">
        <div class="podcast-hero-inner">
            <!-- ===== PODCAST INFO ===== -->
            <div class="container">
                <div class="podcast-hero-content">
                    <span class="podcast-hero-date">`

	code += latest.PubDate[5:16]

	code += `</span>
                    <h2 class="podcast-hero-title"><a href="`

	code += latest.Link + `">` + latest.Title

	code += `<ul class="podcast-hero-meta">
                        <li class="item"><a href="#" class="podcast-hero-tag" rel="tag">Pop Culture</a></li>
                        <li class="item"><i class="fa fa-clock-o"></i> `

	code += latest.Duration + ` mins.</li>
                    </ul>
                </div>
            </div>

            <!-- ===== PODCAST PLAYER ===== -->
            <div class="podcast-hero-player-content">
                <div class="container">
                    <!-- ===== SOUNDCLOUD PLAYER ===== -->
		    <iframe width="100%" height="20" scrolling="no" frameborder="no" allow="autoplay" src="https://w.soundcloud.com/player/?url=https%3A//api.soundcloud.com/tracks/`

	code += latest.GUID.Content[strings.LastIndex(latest.GUID.Content, "/")+1:] + `&color=%23dd3d35&inverse=false&auto_play=false&show_user=true"></iframe>
                </div>
            </div>
        </div>
    </section>`

	return
}

/*
<iframe width="100%" height="20" scrolling="no" frameborder="no" allow="autoplay" src="https://w.soundcloud.com/player/?url=https%3A//api.soundcloud.com/tracks/645022230&color=%23dd3d35&inverse=false&auto_play=false&show_user=true"></iframe>

*/
