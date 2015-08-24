package adapters

import (
	"github.com/sisteamnik/articler"
	"github.com/stretchr/testify/assert"
	"github.com/ungerik/go-dry"
	"io/ioutil"
	"net/url"
	"path/filepath"
	"testing"
	"time"
)

var TEST_DATA_PATH = "./testdata"

func TestSuperExtractTime(t *testing.T) {
	ac, err := articler.ParseAdapterConfig([]byte(aconf))
	if err != nil {
		t.Error(err)
	}
	sp, err := NewParser(ac)
	if err != nil {
		t.Error(err)
	}
	date, err := extractTime(sp.cnf.DateSelector, "", sp.cnf.DateRegex,
		sp.cnf.DateExtractFunc, []byte(data))
	if err != nil {
		t.Error(err)
	}
	t.Log(date)
}

func TestTimeParse(t *testing.T) {
	/*str := "2015-06-24T07:00:53+00:00"
	date, err := extractTime("", time.RFC3339, "", nil, []byte(str))
	if err != nil {
		t.Error(err)
	}
	//if DEBUG {
	log.Println("TestTimeParse", date)
	//}

	date, err = time.Parse(time.RFC3339, str)
	if err != nil {
		t.Error(err)
	}
	//if DEBUG {
	log.Println("TestTimeParse", date)
	//}*/
}

func TestSuperParseArticle(t *testing.T) {
	/*ac, err := articler.ParseAdapterConfig([]byte(aconf))
	if err != nil {
		t.Error(err)
	}
	sp, err := NewParser(ac)
	if err != nil {
		t.Error(err)
	}*/

	// _, err := sp.LastArticles()
	// if err != nil {
	// 	t.Error(err)
	// }
	//log.Println(links)

	/*art, err := sp.Parse([]byte(data))
	if err != nil {
		t.Error(err)
	}
	log.Println(string(art.Body))*/
}

func TestPublishingDate(t *testing.T) {
	type exp struct {
		*url.URL
		time string
	}
	usualyTimeFormat := "02.01.2006 15:04"
	files, err := dry.ListDirFiles(TEST_DATA_PATH)
	if err != nil {
		t.Error(err)
	}
	n := time.Now()
	var urls = map[string]exp{
		"smoldaily.ru":      exp{time: "27.06.2015 00:03", URL: &url.URL{Scheme: "http", Host: "smoldaily.ru", Path: "/smolyan-na-vyxodnyx-zhdut-livni-i-grozy"}},
		"admin-smolensk.ru": exp{time: "26.06.2015 23:18", URL: &url.URL{Scheme: "http", Host: "admin-smolensk.ru", Path: "/news/news_11768.html"}},
		"smolensk-i.ru":     exp{time: "26.06.2015 23:53", URL: &url.URL{Scheme: "http", Host: "smolensk-i.ru", Path: "/authority/v-smolenske-i-oblasti-perestanut-prodavat-energeticheskie-alkogolnyie-napitki_115157"}},
		"smol.kp.ru":        exp{time: n.Add(-24*time.Hour).Format("02.01.2006") + " 17:36", URL: &url.URL{Scheme: "http", Host: "www.smol.kp.ru", Path: "/online/news/2093465"}},
		"rabochy-put.ru":    exp{time: "26.06.2015 16:25", URL: &url.URL{Scheme: "http", Host: "rabochy-put.ru", Path: "/society/64854-11-zolotykh-let-i-6-brilliantov-anastasiya-nikiforova-luchshaya-vypusknitsa-smolenska-2015.html"}},
		"readovka.ru":       exp{time: "24.06.2015 14:12", URL: &url.URL{Scheme: "http", Host: "readovka.ru", Path: "/culture/6338-angelina_interview"}},
	}
	for _, file := range files {
		if u, ok := urls[file]; ok {
			bts, err := ioutil.ReadFile(filepath.Join(TEST_DATA_PATH, file))
			if err != nil {
				t.Error(err)
			}
			date, err := getPublishingDate(u.URL, bts)
			if err != nil {
				t.Error("getPubl not working", err)
			} else {
				assert.Equal(t, u.time, date.Format(usualyTimeFormat), u.Host)
				t.Log("Found date", date)
			}
		}
	}
}

func TestRelativeTime(t *testing.T) {
	str := "вчера, 14:51"
	if !IsRelativeTime(str) {
		t.Error("Error define rel time")
	}
	str = FixRelativeTime(str)
	t.Log(str)
	date, err := ParseTime(str)
	if err != nil {
		t.Error(err)
	}
	t.Log(date)
}

var aconf = `
host: smolensk-i.ru
name: smolenski

feedtype: html
articleuriregex: "^/[a-z-_0-9]*/[a-z-_0-9]*$"
feedselector: ".contentColumn h1 a"

bodyselector: .entry-content

dateselector: time
`

var data = `

<!DOCTYPE html>
<!--[if lt IE 7 ]> <html lang="en" class="no-js ie6"> <![endif]-->
<!--[if IE 7 ]>    <html lang="en" class="no-js ie7"> <![endif]-->
<!--[if IE 8 ]>    <html lang="en" class="no-js ie8"> <![endif]-->
<!--[if IE 9 ]>    <html lang="en" class="no-js ie9"> <![endif]-->
<!--[if (gt IE 9)|!(IE)]><!-->
<html lang="ru" class="no-js" xmlns="http://www.w3.org/1999/xhtml"
      xmlns:og="http://ogp.me/ns#"
      xmlns:fb="https://www.facebook.com/2008/fbml">
<!--<![endif]-->
<head>
<meta charset="utf-8">
<meta name="viewport" content="width=device-width" />
<!--[if IE]> <meta http-equiv="X-UA-Compatible" content="IE=edge,chrome=1"><![endif]-->
<title>
В Смоленске состоится очередная «Фотосушка» - Общество | О чём говорит Смоленск каждый день</title>
<link rel="stylesheet" type="text/css" media="all" href="http://smolensk-i.ru/wp-content/themes/smolenski/css/application.css?v20140318" />
<!--[if IE]><link rel="stylesheet" type="text/css" media="all" href="http://smolensk-i.ru/wp-content/themes/smolenski/css/ie.css" /> <![endif]-->
<link rel="pingback" href="http://smolensk-i.ru/xmlrpc.php" />
<link rel="shortcut icon" href="/favicon.ico">
<link rel="apple-touch-icon" href="/apple-touch-icon.png">

<script src="http://smolensk-i.ru/wp-content/themes/smolenski/js/libs/modernizr-1.6.min.js"></script>
	<link rel="canonical" href="http://smolensk-i.ru/society/v-smolenske-sostoitsya-ocherednaya-fotosushka_114844" />
	<meta name="description" content="Этим летом сушим свои фотоснимки будем около «Кургозора»" />
	<meta prefix="og: http://ogp.me/ns#" property="og:type" content="article">
	<meta prefix="og: http://ogp.me/ns#" property="og:title" content="В Смоленске состоится очередная «Фотосушка»">
	<meta prefix="og: http://ogp.me/ns#" property="og:description" content="Этим летом сушим свои фотоснимки будем около «Кургозора»">
	<meta prefix="og: http://ogp.me/ns#" property="og:url" content="http://smolensk-i.ru/society/v-smolenske-sostoitsya-ocherednaya-fotosushka_114844">
	<meta prefix="og: http://ogp.me/ns#" property="og:image" content="http://smolensk-i.ru/wp-content/uploads/2015/06/jg4aLNugwvU.jpg">
	<meta property="article:published_time" content="2015-06-24">
	<meta property="article:modified_time" content="2015-06-24">
	<meta property="article:section" content="Общество">
	<meta property="article:tag" content="акция">
	<meta property="article:tag" content="Кругозор">
	<meta property="article:tag" content="Смоленск">
	<meta property="article:tag" content="Фотосушка">
	<meta prefix="og: http://ogp.me/ns#" property="og:site_name" content="О чём говорит Смоленск каждый день">
	<meta prefix="fb: http://ogp.me/ns/fb#" property="fb:app_id" content="431149116925061">
	<meta property="twitter:card" content="summary">
<link rel="alternate" type="application/rss+xml" title="О чём говорит Смоленск каждый день &raquo; Лента" href="http://smolensk-i.ru/feed" />
<link rel="alternate" type="application/rss+xml" title="О чём говорит Смоленск каждый день &raquo; Лента комментариев" href="http://smolensk-i.ru/comments/feed" />
<link rel='stylesheet' id='fotorama.css-css'  href='http://smolensk-i.ru/wp-content/plugins/fotorama/fotorama.css?ver=3.6.1' type='text/css' media='all' />
<link rel='stylesheet' id='fotorama-wp.css-css'  href='http://smolensk-i.ru/wp-content/plugins/fotorama/fotorama-wp.css?ver=3.6.1' type='text/css' media='all' />
<link rel='stylesheet' id='jquery.fancybox-css'  href='http://smolensk-i.ru/wp-content/plugins/fancybox2/jquery.fancybox.css?ver=2.0.3' type='text/css' media='all' />
<link rel='stylesheet' id='jquery.fancybox-thumbs-css'  href='http://smolensk-i.ru/wp-content/plugins/fancybox2/helpers/jquery.fancybox-thumbs.css?ver=2.0.3' type='text/css' media='all' />
<script type='text/javascript' src='http://smolensk-i.ru/js/jquery-1.9.1.min.js?ver=1.9.1'></script>
<script type='text/javascript' src='http://smolensk-i.ru/wp-content/plugins/fotorama/fotorama.js?ver=3.6.1'></script>
<script type='text/javascript' src='http://smolensk-i.ru/wp-content/plugins/fotoramaDefaults.js?ver=3.6.1'></script>
<script type='text/javascript' src='http://smolensk-i.ru/wp-content/plugins/fotorama/fotorama-wp.js?ver=3.6.1'></script>
<script type='text/javascript' src='http://smolensk-i.ru/wp-includes/js/jquery/jquery-migrate.min.js?ver=1.2.1'></script>
<script type='text/javascript' src='http://smolensk-i.ru/wp-content/themes/smolenski/js/jquery.tagcloud.min.js?ver=2.0.3'></script>
<script type='text/javascript' src='http://smolensk-i.ru/wp-content/themes/smolenski/js/jquery.scrollTo.min.js?ver=1.4.6'></script>
<script type='text/javascript' src='http://smolensk-i.ru/wp-content/themes/smolenski/js/jquery.ui.totop.js?ver=3.6.1'></script>
<script type='text/javascript' src='http://smolensk-i.ru/wp-includes/js/comment-reply.min.js?ver=3.6.1'></script>
<script type='text/javascript' src='http://smolensk-i.ru/wp-content/plugins/fancybox2/jquery.fancybox.js?ver=2.0.3'></script>
<script type='text/javascript' src='http://smolensk-i.ru/wp-content/plugins/fancybox2/helpers/jquery.fancybox-thumbs.js?ver=2.0.3'></script>
<script type='text/javascript' src='http://smolensk-i.ru/wp-content/plugins/fancybox2/jquery.easing.js?ver=1.6'></script>

<!-- This site is using AdRotate v3.11.1 to display their advertisements - https://ajdg.solutions/products/adrotate-for-wordpress/ -->
<!-- AdRotate CSS -->
<style type="text/css" media="screen">
	.g { margin:0px; padding:0px; overflow:hidden; line-height:1; zoom:1; }
	.g-col { position:relative; float:left; }
	.g-col:first-child { margin-left: 0; }
	.g-col:last-child { margin-right: 0; }
	.g-4 { margin:1px; }
	@media only screen and (max-width: 480px) {
		.g-col, .g-dyn, .g-single { width:100%; margin-left:0; margin-right:0; }
	}
</style>
<!-- /AdRotate CSS -->

<script type="text/javascript">
	jQuery(document).ready(function($){
		var select = $('a[href$=".bmp"],a[href$=".gif"],a[href$=".jpg"],a[href$=".jpeg"],a[href$=".png"],a[href$=".BMP"],a[href$=".GIF"],a[href$=".JPG"],a[href$=".JPEG"],a[href$=".PNG"]', 'div.entry-content:not(#fotorama)');
		select.attr('rel', 'fancybox-thumb');
		select.attr('class', 'fancybox-thumb');
		select.fancybox({
			openEffect	: 'none',
			closeEffect	: 'none',
			prevEffect	: 'elastic',
			nextEffect	: 'elastic',
			nextClick : true,
			closeBtn: true,
			padding:0,
			margin: 40,
			loop: true,
			helpers	: {
				thumbs	: {
					width	: 48,
					height	: 48
				},
				title	: {
					type: 'outside'
				}
			}		
		});
	});
</script>
<!-- ## NXS/OG ## -->
<!-- ## NXS/OG ## -->
<script src="http://smolensk-i.ru/wp-content/themes/smolenski/js/onload.js"></script>
<!--[if lte IE 7 ]><script src="http://smolensk-i.ru/wp-content/themes/smolenski/js/jquery.old-ie.js"></script><![endif]-->
</head>
<body class="single single-post postid-114844 single-format-standard singular two-column right-sidebar">
<div id="fb-root"></div>
<script>(function(d, s, id) {
  var js, fjs = d.getElementsByTagName(s)[0];
  if (d.getElementById(id)) return;
  js = d.createElement(s); js.id = id;
  js.src = "//connect.facebook.net/ru_RU/all.js#xfbml=1&appId=740584415961242";
  fjs.parentNode.insertBefore(js, fjs);
}(document, 'script', 'facebook-jssdk'));</script>
  <div class="banner top">
  
  </div>
<div id="container">
<header id="banner">

  <div class="cover">
    <span class="today"><a href="#" title="все новости за сегодня">24.06.2015</a></span>
    <div class="clearfix logo">
    <img src="/logo_iru_360.jpg" style="position:absolute;top:-9999px;left:-9999px">
    <img src="/wp-content/themes/smolenski/images/logo_i-ru.png" class="logo-iRu" alt="i-Ru: Группа ГС">
    <a class="label-Smolensk" href="/">
    О чём говорит <span>Смоленск</span>
    <img src="/wp-content/themes/smolenski/images/logo_daily.png" alt="Каждый день" title="... да, каждый день" style="">
    </a>
    </div>
    <!-- <a href="http://smolensk-i.ru/culture/parad-pobedyi-v-smolenske-fotoreportazh_71916" title="С Великой Победой!"><img src="/wp-content/themes/smolenski/images/lenta(21).png" alt="С Великой Победой!" title="... да, каждый день" style="position: absolute;top: 18px;right: 160px;height: 70px;"></a>-->
            <form method="get" id="searchform" class="search" action="http://smolensk-i.ru/">
        <input type="text" class="text" name="s" id="s" placeholder="поиск" value=""/>
        <input type="submit" class="submit" name="submit" id="searchsubmit" title="искать на сайте" value="Поиск" />
    </form>
      <ul class="social">
        <li><span class="icon age-limit" title="предназначено для читателей старше 16 лет"></span></li>
          <li><a href="http://www.facebook.com/O4GSDAILY" title="официальная группа в Facebook" target="_blank" rel="external nofollow"><span class="icon fb">Facebook</span></a> </li>
          <li><a class="tw" href="http://twitter.com/O4GS_DAILY" title="официальный Twitter канал" target="_blank" rel="external nofollow"><span class="icon tw">Twitter</span></a> </li>
          <li><a class="vk" href="https://vk.com/club36053190" title="официальная группа в Сети ВКонтакте" target="_blank" rel="external follow"><span class="icon vk">ВКонтакте</span></a> </li>
          <li><a class="ok" href="http://www.odnoklassniki.ru/group/55899698430001" title="официальная группа в Одноклассниках" target="_blank" rel="external nofollow"><span class="icon ok">Одноклассники</span></a></li>
      </ul>    
  </div>

  <div class="topLine">
    <table class="menu">
    <tr>
      <td id="menu-item-166" class=" menu-item menu-item-type-taxonomy menu-item-object-category"><a href="http://smolensk-i.ru/news/politics" title="новости в рубрике «Политика»">Политика</a></td>
<td id="menu-item-163" class=" menu-item menu-item-type-taxonomy menu-item-object-category"><a href="http://smolensk-i.ru/news/authority" title="новости в рубрике «Власть»">Власть</a></td>
<td id="menu-item-165" class=" menu-item menu-item-type-taxonomy menu-item-object-category current-post-ancestor current-menu-parent current-post-parent"><a href="http://smolensk-i.ru/news/society" title="новости в рубрике «Общество»">Общество</a></td>
<td id="menu-item-167" class=" menu-item menu-item-type-taxonomy menu-item-object-category"><a href="http://smolensk-i.ru/news/accidents" title="новости в рубрике «Происшествия»">Происшествия</a></td>
<td id="menu-item-95728" class=" menu-item menu-item-type-taxonomy menu-item-object-category"><a href="http://smolensk-i.ru/news/week-events" title="новости в рубрике «Рейтинг событий»">Рейтинг событий</a></td>
<td id="menu-item-168" class=" menu-item menu-item-type-taxonomy menu-item-object-category"><a href="http://smolensk-i.ru/news/sport" title="новости в рубрике «Спорт»">Спорт</a></td>
<td id="menu-item-164" class=" menu-item menu-item-type-taxonomy menu-item-object-category"><a href="http://smolensk-i.ru/news/culture" title="новости в рубрике «Культура»">Культура</a></td>
<td id="menu-item-170" class=" menu-item menu-item-type-taxonomy menu-item-object-category"><a href="http://smolensk-i.ru/news/auto" title="новости в рубрике «Авто»">Авто</a></td>
<td id="menu-item-162" class=" menu-item menu-item-type-taxonomy menu-item-object-category"><a href="http://smolensk-i.ru/news/business" title="новости в рубрике «Бизнес»">Бизнес</a></td>
    </tr>
    </table>
  </div>

</header>
<div id="content">

<div class="contentColumn wide">
    <article id="post-114844" class="post-114844 post type-post status-publish format-standard hentry category-society tag-aktsiya tag-krugozor tag-smolensk tag-fotosushka">
  <time datetime="2015-06-24T07:00:53+00:00"><span><a href="/2015/06/24/" title="все новости за день">24</a> <a href="/2015/06/" title="все новости за месяц">Июня</a> <a href="/2015/" title="все новости за год">2015 года</a></span> в 07:00, <i>5 часов назад</i></time>
  <h1 class="entry-title">В Смоленске состоится очередная «Фотосушка»</h1>
   
  <div class="entry-content">
    <h2>Этим летом сушим свои фотоснимки будем около «Кургозора»</h2>
<p><a href="http://smolensk-i.ru/wp-content/uploads/2015/06/AAnVY88kZrs.jpg"><img class="aligncenter size-full wp-image-114846" alt="AAnVY88kZrs" src="http://smolensk-i.ru/wp-content/uploads/2015/06/AAnVY88kZrs.jpg" width="800" height="533" /></a></p>
<p>Акция «Фотосушка», приуроченная ко дню молодёжи состоится 26 июня в 16.00 в сквере около магазина «Кургозор» в Смоленске.</p>
<p>Принять участие в акции очень просто. Для этого нужно выбрать свои фотографии согласно правилам, установленным организаторами. Затем распечатать снимки в формате не более А3 и прийти на «Фотосушку», чтобы с помощью прищепок повесить свои фотографии.</p>
<p>В рамках акции можно совершенно бесплатно взять те фотографии, которые понравились больше всего, познакомиться с интересными людьми и пообщаться на тему фотографии.</p>
<p>Предыдущая «Фотосушка» состоялась около кинотеатра «Октябрь» <a href="http://smolensk-i.ru/society/prosushi-svoi-snimki-v-smolenske-proydet-fotosushka_106888" target="_blank">весной этого года</a>.</p>
<p>«Сушка» — это международная акция по обмену фотографиями, придуманная и впервые реализованная в Санкт–Петербурге в 2010 году. С тех пор «Сушка» прошла в более чем 100 городах России и мира и собрала более 45 000 участников.<br />
<em></em></p>
<p>&nbsp;</p>
<p>&nbsp;</p>
<p><em>текст: Анна Романова</em><br />
<em>фото: Евгений Гаврилов</em></p>
  </div>
  
  <footer class="entry-meta">
        <div class="entry-tags">
      <ins class="postIcon tags" title="теги"> </ins><a href="http://smolensk-i.ru/tag/aktsiya" rel="tag">акция</a> <a href="http://smolensk-i.ru/tag/krugozor" rel="tag">Кругозор</a> <a href="http://smolensk-i.ru/tag/smolensk" rel="tag">Смоленск</a> <a href="http://smolensk-i.ru/tag/fotosushka" rel="tag">Фотосушка</a>    </div>
                    <span class="entry-views"><ins class="postIcon views" title="просмотры"> </ins>166 просмотров </span>
                <span class="entry-comments"><ins class="postIcon comments" title="комментарии"> </ins><a href="http://smolensk-i.ru/society/v-smolenske-sostoitsya-ocherednaya-fotosushka_114844#comments">Прокомментируйте это первым!</a></span>
    <script type="text/javascript" src="//yandex.st/share/cnt.share.js" charset="utf-8"></script>
        <br>
        <div class="social-share">
        
        <span class="label-tell">Расскажите о новости в соцсетях:</span> 
        <div style="display:inline-block" class="yashare-auto-init" data-yashareL10n="ru" data-yashareType="small" data-yashareQuickServices="vkontakte,facebook,twitter,odnoklassniki,moimir,gplus,yaru" data-yashareTheme="counter">
        </div>
        </div>       
              </footer>
  <div class="cc"></div>
  <script src="http://smolensk-i.ru/wp-content/themes/smolenski/js/jquery.sticky.js"></script>
  <script>
  jQuery('.social-share').sticky({bottomSpacing:0, wrapperClassName:'footer-wrapper'});
  </script>
</article>
        <div id="comments">
    
    
                    								<div id="respond" class="comment-respond">
				<h3 id="reply-title" class="comment-reply-title">Комментировать <em>или оценить комментарии</em> <small><a rel="nofollow" id="cancel-comment-reply-link" href="/society/v-smolenske-sostoitsya-ocherednaya-fotosushka_114844#respond" style="display:none;">Отменить ответ</a></small></h3>
									<form action="http://smolensk-i.ru/wp-comments-post.php" method="post" id="commentform" class="comment-form">
						<script src="//ulogin.ru/js/ulogin.js" type="text/javascript"></script><div class="ulogin_block"><div class="ulogin_label">Войти с помощью:&nbsp;</div><div id=uLogin0 class="ulogin_panel" data-ulogin="display=small;providers=facebook,twitter,vkontakte,odnoklassniki,google,yandex,mailru;hidden=other;fields=first_name,photo;optional=phone,last_name,email,;redirect_uri=http%3A%2F%2Fsmolensk-i.ru%2F%3Fulogin%3Dtoken%26backurl%3Dhttp%253A%252F%252Fsmolensk-i.ru%252Fsociety%252Fv-smolenske-sostoitsya-ocherednaya-fotosushka_114844%2523commentform;"></div><div style="clear:both"></div></div><script>uLogin.customInit('uLogin0')</script>																				
												<textarea disabled="disabled" class="disabled"></textarea>												<p class="form-submit">
							<input name="submit" type="submit" id="submit-disabled" value="Отправить комментарий" />
							<input type='hidden' name='comment_post_ID' value='114844' id='comment_post_ID' />
<input type='hidden' name='comment_parent' id='comment_parent' value='0' />
						</p>
											</form>
							</div><!-- #respond -->
						    
</div><!-- #comments -->
      <div class="entry-related">
    <h3>Новости по теме</h3>
    <ul>
     <li><span class="date">2 дня назад</span><a href="http://smolensk-i.ru/culture/smolenskie-muzyikantyi-swanky-tunes-snyali-klip-v-lokatsii-karantin_114642" rel="bookmark" title="Смоленские музыканты Swanky Tunes сняли клип в локации «Карантин»"><img width="130" height="130" src="http://smolensk-i.ru/wp-content/uploads/2015/06/wpdDx-VtzdE-130x130.jpg" class="attachment-thumbnail wp-post-image" alt="http://smolensk-i.ru/culture/smolenskie-muzyikantyi-swanky-tunes-snyali-klip-v-lokatsii-karantin_114642"  /><strong>Смоленские музыканты Swanky Tunes сняли клип в локации «Карантин»</strong></a><small>Новое видео можно будет посмотреть в сентябре, а пока доступны только фотографии


Группа диджеев Swanky Tunes сняла клип в action–room «Карантин» в Смоленске.

Всего в съёмках участвовало более 20 человек. Само действо заняло...</small></li><li><span class="date">3 недели назад</span><a href="http://smolensk-i.ru/auto/i-vnov-v-smolenske-pogonyayut-na-krutyih-mashinah_112922" rel="bookmark" title="И вновь в Смоленске погоняют на крутых машинах"><img width="130" height="130" src="http://smolensk-i.ru/wp-content/uploads/2015/06/Bieber-Drag-Racing-Miami_041907792459-130x130.jpg" class="attachment-thumbnail wp-post-image" alt="http://smolensk-i.ru/auto/i-vnov-v-smolenske-pogonyayut-na-krutyih-mashinah_112922"  /><strong>И вновь в Смоленске погоняют на крутых машинах</strong></a><small>Началось лето, а значит, открывается сезон Drag Racing


Соревнования по Drag Racing от 67 Region Club пройдут в Смоленске 6 июня.

Сбор участников состоится на парковке Промышленной администрации в 19.00.

Гостей вечера ожидают танцы...</small></li><li><span class="date">3 недели назад</span><a href="http://smolensk-i.ru/culture/etim-letom-v-smolensk-priplyivyot-titanik_112396" rel="bookmark" title="Этим летом в Смоленск приплывёт «Титаник»"><img width="130" height="130" src="http://smolensk-i.ru/wp-content/uploads/2015/05/the_titanic__gordon_johnson_10-130x130.jpg" class="attachment-thumbnail wp-post-image" alt="http://smolensk-i.ru/culture/etim-letom-v-smolensk-priplyivyot-titanik_112396"  /><strong>Этим летом в Смоленск приплывёт «Титаник»</strong></a><small>История уникального судна предстанет в фотографиях


Открытие выставки «Титаник: 100 лет истории» состоится в Культурно–выставочном центре имени Тенишевых в Смоленске 11 июня. — КВЦ

В экспозиции будут представлены 80 фотографий, которые расскажут об...</small></li><li><span class="date">28.03.2015</span><a href="http://smolensk-i.ru/society/prosushi-svoi-snimki-v-smolenske-proydet-fotosushka_106888" rel="bookmark" title="Просуши свои снимки. В Смоленске пройдет «Фотосушка»"><img width="130" height="130" src="http://smolensk-i.ru/wp-content/uploads/2015/03/0gd5JOrirvM-130x130.jpg" class="attachment-thumbnail wp-post-image" alt="http://smolensk-i.ru/society/prosushi-svoi-snimki-v-smolenske-proydet-fotosushka_106888"  /><strong>Просуши свои снимки. В Смоленске пройдет «Фотосушка»</strong></a><small>В Смоленске состоится акция «Фотосушка», в течение которой жители города смогут обменяться друг с другом различными фотографиями.


Акция пройдет с 14.00 до 18.00 5 апреля в Лопатинском саду.

Цель «Фотосушки» — общение и...</small></li><li><span class="date">14.09.2012</span><a href="http://smolensk-i.ru/society/v-smolenske-otkryivaetsya-sinagoga_2020" rel="bookmark" title="В Смоленске открывается синагога"><img width="130" height="130" src="http://smolensk-i.ru/wp-content/uploads/2012/09/0_86e18_ff44c54a_XL-200x200.jpeg" class="attachment-thumbnail wp-post-image" alt="http://smolensk-i.ru/society/v-smolenske-otkryivaetsya-sinagoga_2020"  /><strong>В Смоленске открывается синагога</strong></a><small>23 сентября в Смоленске откроется синагога.
На торжественном открытии ленточку перережут члены еврейской общины и гости праздника. Затем состоится церемония освящения синагоги, после которой запланирован праздничный концерт.
Напомним, что синагогу в Смоленске начали...</small></li>    </ul>
  </div>
    
  </div>
<aside class="rightColumn widget-area">
              <div id="text-4" class="widget widget_text">			<div class="textwidget"><a target="_BLANK" href="https://www.vtb24.ru/personal/loans/personal/cash/cash4/Pages/default.aspx?geo=smolensk&utm_source=credcash_smolenskiru&utm_medium=media&utm_campaign=advregcredcash&NoMobileRedirect=true#/0401/01/00/01/04"><img src="http://smolensk-i.ru/wp-content/banners/220x275_КН.GIF" width="220"></a>
<br /><br />
<a target="_BLANK" href="http://www.vesnasm.ru"><img src="http://smolensk-i.ru/wp-content/banners/vesnaSmFM102.7.jpg" width="220"></a>
</div>
		</div><div id="recentcomments" class="widget widget_recentcomments"><h3 class="widget-title">Последние комментарии</h3><ul><li class="rc-navi rc-clearfix"><span class="rc-loading">Загрузка...</span></li><li id="rc-comment-temp" class="rc-item rc-comment rc-clearfix"><div class="rc-info"></div><div class="rc-excerpt"></div></li><li id="rc-ping-temp" class="rc-item rc-ping rc-clearfix"><span class="rc-label"></span></li></ul></div>
</aside>
<br class="cc" />
<div class="bottomLine">
</div>

<div class="cc"></div>

<div id="nav-wrapper">
    <span class="nav-previous">
  <a href="http://smolensk-i.ru/culture/zhiteli-smolenska-vspomnyat-vremena-kluba-orfey_114834" rel="prev"><ins></ins></a>    </span>
     <span class="nav-next">
  <a href="http://smolensk-i.ru/society/s-iyulya-zhiteli-smolenska-i-oblasti-budut-bolshe-platit-za-gaz_114849" rel="next"><ins></ins></a>  
    </span>
    
</div>
</div>

<footer id="colophon" role="contentinfo">
  
  
       <table class="copyrights">
            <tr>
            <td>Новостной сайт журнала <a href="http://journal.smolensk-i.ru" target="_blank" title="Перейти к журналу"><em>&laquo;О чём говорит Смоленск&raquo;</em></a></td>
            <td>&copy; ООО <em>&laquo;Группа ГС&raquo;</em>. Все права защищены.</td>
            <td class="last">При перепечатке материалов обязательна ссылка</td>
            </tr>
            <tr>
            <td colspan="3" class="hr">
            <img src="/wp-content/themes/smolenski/images/footer_hr.png">
            </td>
            </tr>
            </table>
 
  <div class="left">
    <ul class="contacts">
		  <li>Главный редактор <em>Светлана Савенок</em></li>    
    	  <li>Шеф-редактор <em>Евгений Ванифатов</em></li>
		  
		  <li>E-mail отдела новостей <em><a href="mailto:press@smolensk-i.ru">press@smolensk-i.ru</a></em></li>
                  <li>E-mail отдела рекламы <em><a href="mailto:smolredaktor@yandex.ru">smolredaktor@yandex.ru</a> [тел. 56-58-23]</em></li>
                  
    </ul>
  </div>  
  <div class="left">
    <ul class="contacts dev">
	  <li>&nbsp;</li>
	  <li>&nbsp;</li>
	  <li>&nbsp;</li>
	  <li>&nbsp;</li>
	  <!-- <li></li>-->
    </ul>
  </div>
  <div class="right">
  <div class="social">
    <ul>
    <li><span class="icon age-limit" title="предназначено для читателей старше 16 лет"></span></li>
      <li><a href="http://www.facebook.com/groups/smolenski" title="группа читателей печатной версии в Facebook" target="_blank" rel="external nofollow"><span class="icon fb">Facebook</span></a> </li>
      <li><a class="tw" href="http://twitter.com/O4GS_DAILY" title="официальный Twitter канал" target="_blank" rel="external nofollow"><span class="icon tw">Twitter</span></a> </li>
      <li><a class="vk" href="http://vk.com/club36053190" title="официальный паблик во Вконтакте" target="_blank" rel="external follow"><span class="icon vk">Twitter</span></a> </li>
      <li><a class="ok" href="http://www.odnoklassniki.ru/group/55899698430001" title="официальная группа в Одноклассниках" target="_blank" rel="external nofollow"><span class="icon ok">Одноклассники</span></a></li>
    </ul>
    
<!--LiveInternet counter--><script type="text/javascript">document.write("<a href='http://www.liveinternet.ru/click' target=_blank><img src='//counter.yadro.ru/hit?t14.1;r" + escape(document.referrer) + ((typeof(screen)=="undefined")?"":";s"+screen.width+"*"+screen.height+"*"+(screen.colorDepth?screen.colorDepth:screen.pixelDepth)) + ";u" + escape(document.URL) +";h"+escape(document.title.substring(0,80)) +  ";" + Math.random() + "' border=0 width=88 height=31 alt='' title='LiveInternet: показано число просмотров за 24 часа, посетителей за 24 часа и за сегодня' style='position:absolute; top:112px; right:20px;'><\/a>")</script><!--/LiveInternet-->
  </div>

</footer>
</div>
<script>
/* <![CDATA[ */
var rcGlobal = {
	serverUrl		:'http://smolensk-i.ru',
	infoTemp		:'%REVIEWER% к новости %POST%',
	loadingText		:'Загрузка',
	noCommentsText	:'No comments',
	newestText		:'К началу',
	newerText		:'Позднее',
	olderText		:'Раньше',
	showContent		:'1',
	external		:'1',
	avatarSize		:'40',
	avatarPosition	:'left',
	anonymous		:'Аноним'
};
/* ]]> */
</script>
<div id="su-footer-links" style="text-align: center;"></div><link rel='stylesheet' id='ulogin-style-css'  href='http://smolensk-i.ru/wp-content/plugins/ulogin/css/ulogin.css?ver=3.6.1' type='text/css' media='all' />
<script type='text/javascript' src='http://smolensk-i.ru/wp-content/plugins/comment-rating/js/comment-rating.js?ver=3.6.1'></script>
<script type='text/javascript' src='http://smolensk-i.ru/wp-content/plugins/wp-recentcomments/js/wp-recentcomments-jquery.js?ver=2.2.7'></script>
<script type='text/javascript'>
/* <![CDATA[ */
var impression_object = {"ajax_url":"http:\/\/smolensk-i.ru\/wp-admin\/admin-ajax.php"};
/* ]]> */
</script>
<script type='text/javascript' src='http://smolensk-i.ru/wp-content/plugins/adrotate/library/jquery.adrotate.dyngroup.js?ver=0.7'></script>
<script type='text/javascript'>
/* <![CDATA[ */
var click_object = {"ajax_url":"http:\/\/smolensk-i.ru\/wp-admin\/admin-ajax.php"};
/* ]]> */
</script>
<script type='text/javascript' src='http://smolensk-i.ru/wp-content/plugins/adrotate/library/jquery.adrotate.clicktracker.js?ver=0.7'></script>
<script type='text/javascript'>
/* <![CDATA[ */
var viewsCacheL10n = {"admin_ajax_url":"http:\/\/smolensk-i.ru\/wp-admin\/admin-ajax.php","post_id":"114844"};
/* ]]> */
</script>
<script type='text/javascript' src='http://smolensk-i.ru/wp-content/plugins/wp-postviews/postviews-cache.js?ver=1.68'></script>
<!-- AdRotate JS -->
<script type="text/javascript">
</script>
<!-- /AdRotate JS -->

<script type="text/javascript">(function (d, w, c) { (w[c] = w[c] || []).push(function() { try { w.yaCounter16941661 = new Ya.Metrika({id:16941661, enableAll: true, trackHash:true, webvisor:true}); } catch(e) { } }); var n = d.getElementsByTagName("script")[0], s = d.createElement("script"), f = function () { n.parentNode.insertBefore(s, n); }; s.type = "text/javascript"; s.async = true; s.src = (d.location.protocol == "https:" ? "https:" : "http:") + "//mc.yandex.ru/metrika/watch.js"; if (w.opera == "[object Opera]") { d.addEventListener("DOMContentLoaded", f); } else { f(); } })(document, window, "yandex_metrika_callbacks");</script><noscript><div><img src="//mc.yandex.ru/watch/16941661" style="position:absolute; left:-9999px;" alt="" /></div></noscript><!-- /Yandex.Metrika counter -->
<script type="text/javascript">
var _gaq = _gaq || []; _gaq.push(['_setAccount', 'UA-34598388-1']); _gaq.push(['_setDomainName', 'smolensk-i.ru']); _gaq.push(['_trackPageview']);
(function() {var ga = document.createElement('script'); ga.type = 'text/javascript'; ga.async = true; ga.src = ('https:' == document.location.protocol ? 'https://ssl' : 'http://www') + '.google-analytics.com/ga.js';var s = document.getElementsByTagName('script')[0]; s.parentNode.insertBefore(ga, s);})();
</script>
</body>
</html> 
 
<!-- Dynamic page generated in 5.295 seconds. -->
<!-- Cached page generated by WP-Super-Cache on 2015-06-24 12:05:51 -->

<!-- super cache -->
`
