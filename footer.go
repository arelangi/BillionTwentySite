package main

func getFooter(feed Feed) (code string) {
	latest := feed.Rss.Channel.Item[0]
	first := feed.Rss.Channel.Item[1]

	code = `    <!-- ===== FOOTER ===== -->
    <footer class="footer">
        <!-- ===== FOOTER CONTENT INFORMATION ===== -->
        <section class="section-positive">
            <div class="container">

                <div class="row mt-70 mb-30">

                    <!-- ===== SITEMAP ===== -->
                    <div class="col-sm-4 mb-40">
                        <h2 class="title-separator white">Sitemap</h2>
                        <ul class="footer-list">
                            <li><a href="index.html">Home</a></li>
                            <li><a href="podcasts.html">Shows</a></li>
                            <li><a href="about-us.html">About Us</a></li>
                            <li><a href="contact.html">Contact</a></li>
                        </ul>
                    </div>

                    <!-- ===== PODCASTS EPISODES ===== -->
                    <div class="col-sm-4 mb-40">
                        <h2 class="title-separator white">Latest Episodes</h2>
                        <ul class="footer-complement">
                            <li>
                                <a href="` + latest.Link + `">` + ellipsis(latest.Title) + `</a>
                                <span>` + latest.PubDate[5:16] + `</span>
                            </li>
                            <li>
                                <a href="` + first.Link + `">` + ellipsis(first.Title) + `</a>
				<span>` + first.PubDate[5:16] + `</span>
                            </li>
                        </ul>
                    </div>

                    <!-- ===== SOCIAL CONNECTION ===== -->
                    <div class="col-sm-4 mb-40">
                        <h2 class="title-separator white">We are social</h2>
                        <ul class="social-list">
                            <li class="social-item"><a href="https://facebook.com" target="_blank"><i class="fa fa-facebook"></i></a></li>
                            <li class="social-item"><a href="https://twitter.com/billiontwenty" target="_blank"><i class="fa fa-twitter"></i></a></li>
                            <li class="social-item"><a href="https://instagram.com/billiontwenty"" target="_blank"><i class="fa fa-instagram"></i></a></li>
                            <li class="social-item"><a href="#" target="_blank"><i class="fa fa-google"></i></a></li>
                            <li class="social-item"><a href="https://soundcloud.com/billiontwenty" target="_blank"><i class="fa fa-soundcloud"></i></a></li>
                            <li class="social-item"><a href="#" target="_blank"><i class="fa fa-youtube"></i></a></li>
                            <li class="social-item"><a href="https://open.spotify.com/show/76TTxV7gAUgvVDGHTTovcB?si=kWyuwdDZQcyektE_BNs7rw" target="_blank"><i class="fa fa-spotify"></i></a></li>
                            <li class="social-item"><a href="https://podcasts.apple.com/us/podcast/in-the-end-telugu/id1473265557" target="_blank"><i class="fa fa-podcast"></i></a></li>
                        </ul>
                    </div>

                </div>

            </div>
        </section>

        <!-- ===== FOOTER INFORMATION ===== -->
        <section class="footer-credits">
            <div class="container">

                <div class="row">

                    <!-- ===== CREDIT LOGO ===== -->
                    <div class="col-sm-6 footer-logo">
                        <h2>
                            <a href="index.html"><img id="logo-footer" src="assets/img/LL3.png" alt="Billion Twenty Media" title="Billion Twenty Media" /></a>
                        </h2>
                    </div>

                    <!-- ===== CREDIT LOGO ===== -->
                    <div class="col-sm-6 text-right">
                        Billion Twenty Media - 2019. All rights reserved.
                    </div>

                </div>

            </div>
        </section>
    </footer>

    <!-- =================================== -->
    <!-- 			  SCRIPTS 				 -->
    <!-- =================================== -->

    <!-- JQUERY -->
    <script src="assets/js/jquery-1.11.min.js"></script>

    <!-- BOOTSTRAP JS -->
    <script src="assets/js/bootstrap.min.js"></script>

    <!-- MEDIA ELEMENT -->
    <script src="assets/js/mediaelement-and-player.min.js"></script>

    <!-- MAGNIFIC POPUP -->
    <script src="assets/js/magnific-popup.min.js"></script>

    <!-- FORM VALIDATE -->
    <script src="assets/js/validate.min.js"></script>

    <!-- PLACEHOLDER FOR IE -->
    <script src="assets/js/placeholder.min.js"></script>

    <!-- THEME JS -->
    <script src="assets/js/main.js"></script>

    <!-- Global site tag (gtag.js) - Google Analytics -->
    <script async src="https://www.googletagmanager.com/gtag/js?id=UA-146878075-2">
    </script>
    <script>
      window.dataLayer = window.dataLayer || [];
      function gtag(){dataLayer.push(arguments);}
      gtag('js', new Date());

      gtag('config', 'UA-146878075-2');
    </script>

</body>

</html>`

	return
}
