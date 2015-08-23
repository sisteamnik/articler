package adapters

import (
	"log"
	"testing"
)

func TestFasGetDate(t *testing.T) {
	rp := NewFasParser()
	/*var dates = map[string]time.Time{
		"21 ИЮНЬ 2015 09:46": time.Time{},
	}*/
	date, err := rp.getDate("22 июня 2015, 15:03")
	if err != nil {
		t.Error(err)
	}
	if DEBUG {
		log.Println(date)
	}

}

func TestFasIsArticle(t *testing.T) {
	rp := NewFasParser()
	if !rp.IsArticle("/news/11897") {
		t.Error("Is article failed")
	}

	if rp.IsArticle("/2133") {
		t.Error("Is article failed")
	}
}

func TestFasGetDateUpdated(t *testing.T) {
	rp := NewFasParser()
	var html = `<div class="submitted left">22 июня 2015, 15:03</div>`
	date, err := rp.getDate(html)
	if err != nil {
		t.Error(err)
	}
	if DEBUG {
		log.Println(date)
	}
}

//it use metworking
func TestFasLastArticles(t *testing.T) {
	/*rp := NewFasParser()
	urls, err := rp.LastArticles()
	if err != nil {
		t.Error(err)
	}
	if DEBUG {
		log.Println(urls)
	}*/
}

func TestFasParse(t *testing.T) {
	rp := NewFasParser()
	art, err := rp.Parse([]byte(fasTestArticle))
	if err != nil {
		t.Error(err)
	}
	if DEBUG {
		log.Println(art.Title)
		log.Println(art.Published)
	}
}

var fasTestArticle = `
<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML+RDFa 1.0//EN"
  "http://www.w3.org/MarkUp/DTD/xhtml-rdfa-1.dtd">
<html xmlns="http://www.w3.org/1999/xhtml" xml:lang="ru" version="XHTML+RDFa 1.0" dir="ltr"
  xmlns:content="http://purl.org/rss/1.0/modules/content/"
  xmlns:dc="http://purl.org/dc/terms/"
  xmlns:foaf="http://xmlns.com/foaf/0.1/"
  xmlns:og="http://ogp.me/ns#"
  xmlns:rdfs="http://www.w3.org/2000/01/rdf-schema#"
  xmlns:sioc="http://rdfs.org/sioc/ns#"
  xmlns:sioct="http://rdfs.org/sioc/types#"
  xmlns:skos="http://www.w3.org/2004/02/skos/core#"
  xmlns:xsd="http://www.w3.org/2001/XMLSchema#">

<head profile="http://www.w3.org/1999/xhtml/vocab">
  <meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
<meta about="/news/11890" property="sioc:num_replies" content="0" datatype="xsd:integer" />
<link rel="shortcut icon" href="http://smolensk.fas.gov.ru/sites/smolensk.f.isfb.ru/files/favicon-fas.ico" type="image/vnd.microsoft.icon" />
<meta content="Данные мониторинга средневзвешенных оптово-отпускных цен по состоянию на 11.06.2015" about="/news/11890" property="dc:title" />
<link rel="shortlink" href="/node/11890" />
<meta name="Generator" content="Drupal 7 (http://drupal.org)" />
<link rel="canonical" href="/news/11890" />
  <title>Данные мониторинга средневзвешенных оптово-отпускных цен по состоянию на 11.06.2015 | Смоленское УФАС</title>
  <link type="text/css" rel="stylesheet" href="http://smolensk.fas.gov.ru/sites/smolensk.f.isfb.ru/files/css/css_pbm0lsQQJ7A7WCCIMgxLho6mI_kBNgznNUWmTWcnfoE.css" media="all" />
<link type="text/css" rel="stylesheet" href="http://smolensk.fas.gov.ru/sites/smolensk.f.isfb.ru/files/css/css_-TNq6F6EH1K3WcBMUMQP90OkyCq0Lyv1YnyoEj3kxiU.css" media="screen" />
<link type="text/css" rel="stylesheet" href="http://smolensk.fas.gov.ru/sites/smolensk.f.isfb.ru/files/css/css_GosY8GqvTU6t9wT7re_sxx7RRpK9yvvjHJHd3WIMet8.css" media="all" />
<link type="text/css" rel="stylesheet" href="http://smolensk.fas.gov.ru/sites/smolensk.f.isfb.ru/files/css/css_2083LY1J1ajaPycSoN7vcU4xNtcaB7xKyJ2HicTOKKE.css" media="all" />
<link type="text/css" rel="stylesheet" href="http://smolensk.fas.gov.ru/sites/smolensk.f.isfb.ru/files/css/css_Y8HdaQKgD2Q75TocMqoQMY0pjRpZHceX6U7TudJ_XvA.css" media="screen" />
<link type="text/css" rel="stylesheet" href="http://smolensk.fas.gov.ru/sites/smolensk.f.isfb.ru/files/css/css_Iy2FKCge4uVrxP6RFknn8ZOBkOGXBtgONJGX9nfEQpw.css" media="print" />

<!--[if lte IE 8]>
<link type="text/css" rel="stylesheet" href="http://smolensk.fas.gov.ru/sites/all/themes/fasregion/css/ie.css?noloqo" media="all" />
<![endif]-->

<!--[if IE 6]>
<link type="text/css" rel="stylesheet" href="http://smolensk.fas.gov.ru/sites/all/themes/fasregion/css/ie6.css?noloqo" media="all" />
<![endif]-->

<!--[if IE 7]>
<link type="text/css" rel="stylesheet" href="http://smolensk.fas.gov.ru/sites/all/themes/fasregion/css/ie7.css?noloqo" media="all" />
<![endif]-->
  <script type="text/javascript" src="http://smolensk.fas.gov.ru/misc/jquery.js?v=1.4.4"></script>
<script type="text/javascript" src="http://smolensk.fas.gov.ru/misc/jquery.once.js?v=1.2"></script>
<script type="text/javascript" src="http://smolensk.fas.gov.ru/misc/drupal.js?noloqo"></script>
<script type="text/javascript" src="http://smolensk.fas.gov.ru/sites/all/libraries/shadowbox/shadowbox.js?v=3.0.3"></script>
<script type="text/javascript" src="http://smolensk.fas.gov.ru/sites/all/modules/mollom/mollom.js?noloqo"></script>
<script type="text/javascript" src="http://smolensk.fas.gov.ru/sites/smolensk.f.isfb.ru/files/languages/ru_EIOed_Gf2G-6AmJ_J5F38aTaK-QxCpbCUQDjwtWfc0I.js?noloqo"></script>
<script type="text/javascript" src="http://smolensk.fas.gov.ru/sites/all/libraries/chosen/chosen/chosen.jquery.min.js?noloqo"></script>
<script type="text/javascript" src="http://smolensk.fas.gov.ru/sites/all/modules/chosen/chosen.js?noloqo"></script>
<script type="text/javascript" src="http://smolensk.fas.gov.ru/sites/all/libraries/colorbox/colorbox/jquery.colorbox-min.js?noloqo"></script>
<script type="text/javascript" src="http://smolensk.fas.gov.ru/sites/all/modules/colorbox/js/colorbox.js?noloqo"></script>
<script type="text/javascript" src="http://smolensk.fas.gov.ru/sites/all/modules/colorbox/styles/default/colorbox_default_style.js?noloqo"></script>
<script type="text/javascript" src="http://smolensk.fas.gov.ru/sites/all/modules/nice_menus/superfish/js/superfish.js?noloqo"></script>
<script type="text/javascript" src="http://smolensk.fas.gov.ru/sites/all/modules/nice_menus/superfish/js/jquery.bgiframe.min.js?noloqo"></script>
<script type="text/javascript" src="http://smolensk.fas.gov.ru/sites/all/modules/nice_menus/superfish/js/jquery.hoverIntent.minified.js?noloqo"></script>
<script type="text/javascript" src="http://smolensk.fas.gov.ru/sites/all/modules/nice_menus/nice_menus.js?noloqo"></script>
<script type="text/javascript" src="http://smolensk.fas.gov.ru/sites/all/modules/panels/js/panels.js?noloqo"></script>
<script type="text/javascript" src="http://smolensk.fas.gov.ru/sites/all/modules/video/js/video.js?noloqo"></script>
<script type="text/javascript" src="http://smolensk.fas.gov.ru/sites/all/themes/fasregion/js/jquery.bgiframe.min.js?noloqo"></script>
<script type="text/javascript" src="http://smolensk.fas.gov.ru/sites/all/themes/fasregion/js/jquery.hoverIntent.minified.js?noloqo"></script>
<script type="text/javascript" src="http://smolensk.fas.gov.ru/sites/all/themes/fasregion/js/jquery.livequery.min.js?noloqo"></script>
<script type="text/javascript" src="http://smolensk.fas.gov.ru/sites/all/themes/fasregion/js/superfish.js?noloqo"></script>
<script type="text/javascript" src="http://smolensk.fas.gov.ru/sites/all/themes/fasregion/js/jquery-pass2text.js?noloqo"></script>
<script type="text/javascript" src="http://smolensk.fas.gov.ru/sites/all/themes/fasregion/js/jquery-pngfix.js?noloqo"></script>
<script type="text/javascript" src="http://smolensk.fas.gov.ru/sites/all/themes/fasregion/js/layout.js?noloqo"></script>
<script type="text/javascript" src="http://smolensk.fas.gov.ru/sites/all/themes/fasregion/js/social.js?noloqo"></script>
<script type="text/javascript">
<!--//--><![CDATA[//><!--
jQuery.extend(Drupal.settings, {"basePath":"\/","pathPrefix":"","ajaxPageState":{"theme":"fasregion","theme_token":"yidbiN1kR-TlXGR9lJBQWcbugsQ-WLTT_byGnHm7pyU","js":{"0":1,"misc\/jquery.js":1,"misc\/jquery.once.js":1,"misc\/drupal.js":1,"sites\/all\/libraries\/shadowbox\/shadowbox.js":1,"sites\/all\/modules\/mollom\/mollom.js":1,"public:\/\/languages\/ru_EIOed_Gf2G-6AmJ_J5F38aTaK-QxCpbCUQDjwtWfc0I.js":1,"sites\/all\/libraries\/chosen\/chosen\/chosen.jquery.min.js":1,"sites\/all\/modules\/chosen\/chosen.js":1,"sites\/all\/libraries\/colorbox\/colorbox\/jquery.colorbox-min.js":1,"sites\/all\/modules\/colorbox\/js\/colorbox.js":1,"sites\/all\/modules\/colorbox\/styles\/default\/colorbox_default_style.js":1,"sites\/all\/modules\/nice_menus\/superfish\/js\/superfish.js":1,"sites\/all\/modules\/nice_menus\/superfish\/js\/jquery.bgiframe.min.js":1,"sites\/all\/modules\/nice_menus\/superfish\/js\/jquery.hoverIntent.minified.js":1,"sites\/all\/modules\/nice_menus\/nice_menus.js":1,"sites\/all\/modules\/panels\/js\/panels.js":1,"sites\/all\/modules\/video\/js\/video.js":1,"sites\/all\/themes\/fasregion\/js\/jquery.bgiframe.min.js":1,"sites\/all\/themes\/fasregion\/js\/jquery.hoverIntent.minified.js":1,"sites\/all\/themes\/fasregion\/js\/jquery.livequery.min.js":1,"sites\/all\/themes\/fasregion\/js\/superfish.js":1,"sites\/all\/themes\/fasregion\/js\/jquery-pass2text.js":1,"sites\/all\/themes\/fasregion\/js\/jquery-pngfix.js":1,"sites\/all\/themes\/fasregion\/js\/layout.js":1,"sites\/all\/themes\/fasregion\/js\/social.js":1},"css":{"modules\/system\/system.base.css":1,"modules\/system\/system.menus.css":1,"modules\/system\/system.messages.css":1,"modules\/system\/system.theme.css":1,"sites\/all\/libraries\/shadowbox\/shadowbox.css":1,"modules\/comment\/comment.css":1,"sites\/all\/modules\/date\/date_api\/date.css":1,"sites\/all\/modules\/date\/date_popup\/themes\/datepicker.1.7.css":1,"sites\/all\/modules\/ISFB\/fasto_search\/search.css":1,"modules\/field\/theme\/field.css":1,"sites\/all\/modules\/mollom\/mollom.css":1,"modules\/node\/node.css":1,"modules\/poll\/poll.css":1,"modules\/user\/user.css":1,"sites\/all\/modules\/views\/css\/views.css":1,"sites\/all\/libraries\/chosen\/chosen\/chosen.css":1,"sites\/all\/modules\/colorbox\/styles\/default\/colorbox_default_style.css":1,"sites\/all\/modules\/ctools\/css\/ctools.css":1,"sites\/all\/modules\/nice_menus\/nice_menus.css":1,"sites\/all\/modules\/nice_menus\/nice_menus_default.css":1,"sites\/all\/modules\/panels\/css\/panels.css":1,"sites\/all\/modules\/video\/css\/video.css":1,"sites\/all\/themes\/fasregion\/css\/reset.css":1,"sites\/all\/themes\/fasregion\/css\/fasregion.css":1,"sites\/all\/themes\/fasregion\/css\/jquery.ui.theme.css":1,"sites\/all\/themes\/fasregion\/css\/region.css":1,"sites\/all\/themes\/fasregion\/css\/print.css":1,"sites\/all\/themes\/fasregion\/css\/ie.css":1,"sites\/all\/themes\/fasregion\/css\/ie6.css":1,"sites\/all\/themes\/fasregion\/css\/ie7.css":1}},"chosen":{"selector":".chzn-select","minimum":"0"},"colorbox":{"opacity":"0.85","current":"{current} \u0438\u0437 {total}","previous":"\u00ab \u041f\u0440\u0435\u0434\u044b\u0434\u0443\u0449\u0438\u0439","next":"\u0421\u043b\u0435\u0434\u0443\u044e\u0449\u0438\u0439 \u00bb","close":"\u0417\u0430\u043a\u0440\u044b\u0442\u044c","maxWidth":"100%","maxHeight":"100%","fixed":true,"__drupal_alter_by_ref":["default"]},"nice_menus_options":{"delay":"800","speed":"slow"},"shadowbox":{"animate":true,"animateFade":true,"animSequence":"wh","auto_enable_all_images":0,"auto_gallery":0,"autoplayMovies":true,"continuous":false,"counterLimit":10,"counterType":"default","displayCounter":true,"displayNav":true,"enableKeys":true,"fadeDuration":0.35,"handleOversize":"resize","handleUnsupported":"link","initialHeight":160,"initialWidth":320,"language":"ru","modal":false,"overlayColor":"#000","overlayOpacity":0.8,"resizeDuration":0.55,"showMovieControls":true,"slideshowDelay":0,"viewportPadding":20,"useSizzle":false}});
//--><!]]>
</script>
</head>
<body class="html not-front not-logged-in no-sidebars page-node page-node- page-node-11890 node-type-news" >
  <div id="skip-link">
    <a href="#main-content" class="element-invisible element-focusable">Перейти к основному содержанию</a>
  </div>
    


<div id="fb-root"></div>
<div id="page">
<div id="header" class="">
    <div class="logos">
              <a href="/" class="logo"><img src="http://smolensk.fas.gov.ru/sites/smolensk.f.isfb.ru/files/coat_of_arms_of_smolensk_smolensk_oblast_2001.png"
                                      class="left"/><span class="left">Управление Федеральной антимонопольной службы<br />по Смоленской области</span></a>
          </div>
          <div id="navlist">
        <h2 class="element-invisible">Основные ссылки</h2>
    <ul class="links inline main-menu">
          <li class="menu-217 first"><a title="" href="/taxonomy/term/3">Антимонопольное регулирование</a>
                  <ul>        <li><a href="/taxonomy/term/47">Банковские и страховые услуги</a></li>
                    <li><a href="/taxonomy/term/36">ЖКХ</a></li>
                    <li><a href="/taxonomy/term/50">Контроль проведения торгов</a></li>
                    <li><a href="/taxonomy/term/48">Регулирование торговой деятельности</a></li>
                    <li><a href="/taxonomy/term/30">Розничная торговля</a></li>
                    <li><a href="/taxonomy/term/46">Транспорт и связь</a></li>
                    <li><a href="/taxonomy/term/1">Электроэнергетика</a></li>
                    <li><a href="/taxonomy/term/31">Контроль органов власти</a></li>
      
                                                </ul>        </li>
          <li class="menu-497"><a title="" href="/taxonomy/term/4">Контроль госзакупок</a>
            
                                                
        </li>
          <li class="menu-498 last"><a title="" href="/taxonomy/term/5">Контроль рекламы и недобросовестной конкуренции</a>
            
                                                        </li>
        </ul>
    </div>
    <div class="icons-top left">
        <a href="/" class="home"></a>
        <a href="/node/3786" class="phone"></a>
        <a href="http://smolensk.fas.gov.ru/webform/5970" class="mail"></a>
        <a href="/sitemap" class="map"></a>
    </div>
    <div class="region-select left">
        <ul class="select-region sf-js-enabled">
            <li class=""><a href="">Выбор региона</a><br>
                <ul class="regions" style="visibility: hidden; display: none;">
                    <li class="first_column column">
                      <ul class="regions_main">
                        <li class="subheader"><span>Центральный федеральный округ:</span></li>
                        <li><a href="http://belgorod.fas.gov.ru">Белгородское УФАС России</a></li>
                        <li><a href="http://bryansk.fas.gov.ru">Брянское УФАС России</a></li>
                        <li><a href="http://vladimir.fas.gov.ru">Владимирское УФАС России</a></li>
                        <li><a href="http://voronezh.fas.gov.ru">Воронежское УФАС России</a></li>
                        <li><a href="http://ivanovo.fas.gov.ru">Ивановское УФАС России</a></li>
                        <li><a href="http://kaluga.fas.gov.ru">Калужское УФАС России</a></li>
                        <li><a href="http://kostroma.fas.gov.ru">Костромское УФАС России</a></li>
                        <li><a href="http://kursk.fas.gov.ru">Курское УФАС России</a></li>
                        <li><a href="http://lipetsk.fas.gov.ru">Липецкое УФАС России</a></li>
                        <li><a href="http://mo.fas.gov.ru">Московское областное УФАС России</a></li>
                        <li><a href="http://orel.fas.gov.ru">Орловское УФАС России</a></li>
                        <li><a href="http://ryazan.fas.gov.ru">Рязанское УФАС России</a></li>
                        <li><a href="http://smolensk.fas.gov.ru">Смоленское УФАС России</a></li>
                        <li><a href="http://tambov.fas.gov.ru">Тамбовское УФАС России</a></li>
                        <li><a href="http://tver.fas.gov.ru">Тверское УФАС России</a></li>
                        <li><a href="http://tula.fas.gov.ru">Тульское УФАС России</a></li>
                        <li><a href="http://yaroslavl.fas.gov.ru">Ярославское УФАС России</a></li>
                        <li><a href="http://moscow.fas.gov.ru">Московское УФАС России</a></li>
                        <li class="subheader"><span>Южный федеральный округ:</span></li>
                        <li><a href="http://adygea.fas.gov.ru">Адыгейское УФАС России</a></li>
                        <li><a href="http://kalmykia.fas.gov.ru">Калмыцкое УФАС России</a></li>
                        <li><a href="http://krasnodar.fas.gov.ru">Краснодарское УФАС России</a></li>
                        <li><a href="http://astrahan.fas.gov.ru">Астраханское УФАС России</a></li>
                        <li><a href="http://volgograd.fas.gov.ru">Волгоградское УФАС России</a></li>
                        <li><a href="http://rostov.fas.gov.ru">Ростовское УФАС России</a></li>
                      </ul>
                    </li>
                    <li class="second_column column">
                      <ul>
                        <li class="subheader"><span>Северо-Западный федеральный округ:</span></li>
                        <li><a href="http://karelia.fas.gov.ru">Карельское УФАС России</a></li>
                        <li><a href="http://komi.fas.gov.ru">Коми УФАС России</a></li>
                        <li><a href="http://arhangelsk.fas.gov.ru">Архангельское УФАС России</a></li>
                        <li><a href="http://vologda.fas.gov.ru">Вологодское УФАС России</a></li>
                        <li><a href="http://kaliningrad.fas.gov.ru">Калининградское УФАС России</a></li>
                        <li><a href="http://lenobl.fas.gov.ru">Ленинградское областное УФАС России</a></li>
                        <li><a href="http://murmansk.fas.gov.ru">Мурманское УФАС России</a></li>
                        <li><a href="http://novgorod.fas.gov.ru">Новгородское УФАС России</a></li>
                        <li><a href="http://pskov.fas.gov.ru">Псковское УФАС России</a></li>
                        <li><a href="http://spb.fas.gov.ru">Санкт-Петербургское УФАС России</a></li>
                        <li><a href="http://nao.fas.gov.ru">Ненецкое УФАС России</a></li>
                        <li class="subheader"><span>Дальневосточный федеральный округ:</span></li>
                        <li><a href="http://sakha.fas.gov.ru">Якутское УФАС России</a></li>
                        <li><a href="http://kamchatka.fas.gov.ru">Камчатское УФАС России</a></li>
                        <li><a href="http://primorie.fas.gov.ru">Приморское УФАС России</a></li>
                        <li><a href="http://habarovsk.fas.gov.ru">Хабаровское УФАС России</a></li>
                        <li><a href="http://amur.fas.gov.ru">Амурское УФАС России</a></li>
                        <li><a href="http://magadan.fas.gov.ru">Магаданское УФАС России</a></li>
                        <li><a href="http://sahalin.fas.gov.ru">Сахалинское УФАС России</a></li>
                        <li><a href="http://eao.fas.gov.ru">Еврейское УФАС России</a></li>
                        <li><a href="http://chukotka.fas.gov.ru">Чукотское УФАС России</a></li>
                      </ul>
                    </li>
                    <li class="third_column column">
                      <ul>
                        <li class="subheader"><span>Сибирский федеральный округ:</span></li>
                        <li><a href="http://altr.fas.gov.ru">Алтайское республиканское УФАС России</a></li>
                        <li><a href="http://buryatia.fas.gov.ru">Бурятское УФАС России</a></li>
                        <li><a href="http://tuva.fas.gov.ru">Тывинское УФАС России</a></li>
                        <li><a href="http://hakasia.fas.gov.ru">Хакасское УФАС России</a></li>
                        <li><a href="http://altk.fas.gov.ru">Алтайское краевое УФАС России</a></li>
                        <li><a href="http://zab.fas.gov.ru">Забайкальское УФАС России</a></li>
                        <li><a href="http://krsk.fas.gov.ru">Красноярское УФАС России</a></li>
                        <li><a href="http://irkutsk.fas.gov.ru">Иркутское УФАС России</a></li>
                        <li><a href="http://kemerovo.fas.gov.ru">Кемеровское УФАС России</a></li>
                        <li><a href="http://novosibirsk.fas.gov.ru">Новосибирское УФАС России</a></li>
                        <li><a href="http://omsk.fas.gov.ru">Омское УФАС России</a></li>
                        <li><a href="http://tomsk.fas.gov.ru">Томское УФАС России</a></li>
                        <li class="subheader"><span>Уральский федеральный округ:</span></li>
                        <li><a href="http://kurgan.fas.gov.ru">Курганское УФАС России</a></li>
                        <li><a href="http://sverdlovsk.fas.gov.ru">Свердловское УФАС России</a></li>
                        <li><a href="http://tyumen.fas.gov.ru">Тюменское УФАС России</a></li>
                        <li><a href="http://chel.fas.gov.ru">Челябинское УФАС России</a></li>
                        <li><a href="http://hmao.fas.gov.ru">Ханты-Мансийское УФАС России</a></li>
                        <li><a href="http://yamal.fas.gov.ru">Ямало-Ненецкое УФАС России</a></li>
                      </ul>
                    </li>
                    <li class="fourth_column column last">
                      <ul>
                        <li class="subheader"><span>Приволжский федеральный округ:</span></li>
                        <li><a href="http://bash.fas.gov.ru">Башкортостанское УФАС России</a></li>
                        <li><a href="http://mari-el.fas.gov.ru">Марийское УФАС России</a></li>
                        <li><a href="http://mordovia.fas.gov.ru">Мордовское УФАС России</a></li>
                        <li><a href="http://fasrt.ru">Татарстанское УФАС России</a></li>
                        <li><a href="http://udmurtia.fas.gov.ru">Удмуртское УФАС России</a></li>
                        <li><a href="http://chuvashia.fas.gov.ru">Чувашское УФАС России</a></li>
                        <li><a href="http://kirov.fas.gov.ru">Кировское УФАС России</a></li>
                        <li><a href="http://n-novgorod.fas.gov.ru">Нижегородское УФАС России</a></li>
                        <li><a href="http://orenburg.fas.gov.ru">Оренбургское УФАС России</a></li>
                        <li><a href="http://penza.fas.gov.ru">Пензенское УФАС России</a></li>
                        <li><a href="http://ulyanovsk.fas.gov.ru">Ульяновское УФАС России</a></li>
                        <li><a href="http://samara.fas.gov.ru">Самарское УФАС России</a></li>
                        <li><a href="http://saratov.fas.gov.ru">Саратовское УФАС России</a></li>
                        <li><a href="http://perm.fas.gov.ru">Пермское УФАС России</a></li>
                        <li class="subheader"><span>Северо - Кавказский федеральный округ:</span></li>
                        <li><a href="http://dagestan.fas.gov.ru">Дагестанское УФАС России</a></li>
                        <li><a href="http://che-in.fas.gov.ru">Ингушское УФАС России</a></li>
                        <li><a href="http://kbr.fas.gov.ru">Кабардино-Балкарское УФАС России</a></li>
                        <li><a href="http://kchr.fas.gov.ru">Карачаево-Черкесское УФАС России</a></li>
                        <li><a href="http://so-alania.fas.gov.ru">Северо-осетинское УФАС России</a></li>
                        <li><a href="http://chechnya.fas.gov.ru/">Чеченское УФАС России</a></li>
                        <li><a href="http://stavropol.fas.gov.ru">Ставропольское УФАС России</a></li>
                        <li class="subheader"><span>Крымский федеральный округ:</span></li>
                        <li><a href="http://krym.fas.gov.ru/">Крымское УФАС России</a></li>
                      </ul>
                    </li>
                </ul>
            </li>
        </ul>
    </div>
      <div class="authoriz left"><a href="/user">Подписка</a></div>
  
    <div class="region region-header">
        <div id="block-search-form" class="block block-search">
                    <form action="/news/11890" method="post" id="fasto-search-box" accept-charset="UTF-8"><div><div class="form-item form-type-textfield form-item-search-block-form">
  <label class="element-invisible" for="edit-search-block-form">Поиск </label>
 <input title="Введите ключевые слова для поиска." type="text" id="edit-search-block-form" name="search_block_form" value="" size="15" maxlength="128" class="form-text" />
</div>
<input type="hidden" name="form_build_id" value="form-zKpKQvtSvr927_yDUqapzmbBeDP1fF5zzT6VdO4QAb8" />
<input type="hidden" name="form_id" value="fasto_search_box" />
<div class="form-actions form-wrapper" id="edit-actions"><input type="submit" id="edit-submit" name="op" value="Поиск" class="form-submit" /></div></div></form>            <div class="advanced_search"><a href="/search">Расширенный поиск</a></div>
                  </div>
    </div>



  </div>

<div class="main-content">
  <div id="left">
    <div class="inner">
      <div class="region region-left">
                  <div id="block-menu-menu-materials" class="block block-menu">

                      
  <div class="content">
    <ul class="menu"><li class="first leaf"><a href="/news">Новости</a></li>
<li class="leaf"><a href="/media" title="">УФАС в СМИ</a></li>
<li class="leaf"><a href="/fas_sso/redirect_to_solutions_search" title="">Решения</a></li>
<li class="last leaf"><a href="/analytic" title="">Аналитические материалы</a></li>
</ul>  </div>
</div>
        
                  <div id="block-menu-menu-activiti-information" class="block block-menu">

                      
  <div class="content">
    <ul class="menu"><li class="first expanded"><a href="/calendarpage" title="">Об управлении</a><ul class="menu"><li class="first leaf"><a href="/calendarpage" title="">График рассмотрения дел</a></li>
<li class="leaf"><a href="/structure" title="">Структура УФАС</a></li>
<li class="leaf"><a href="/reports" title="">Отчетность</a></li>
<li class="leaf"><a href="/node/3791">Закупки</a></li>
<li class="leaf"><a href="/node/3792">Соглашения о взаимодействии</a></li>
<li class="leaf"><a href="/page/9234">История</a></li>
<li class="leaf"><a href="/page/9284">Миссия, цели, ценности</a></li>
<li class="leaf"><a href="/page/9280">О руководстве Смоленского УФАС России</a></li>
<li class="leaf"><a href="/page/9289">О сайте</a></li>
<li class="leaf"><a href="/page/9481" title="Руководство для поступающих на государственную службу в территориальный орган ФАС">Первые шаги в Смоленском УФАС России</a></li>
<li class="leaf"><a href="/page/9260">Проверки</a></li>
<li class="last leaf"><a href="/page/9262">Руководитель ФАС России</a></li>
</ul></li>
<li class="leaf"><a href="/documents_base" title="">Нормативно-правовая база</a></li>
<li class="leaf"><a href="/advice/9261" title="">Общественно-консультативный совет при Смоленском УФАС России</a></li>
<li class="expanded"><a href="/page/9252" title="">Госслужба</a><ul class="menu"><li class="first leaf"><a href="/page/9252">Поступление на госслужбу</a></li>
<li class="leaf"><a href="http://smolensk.fas.gov.ru/page/10319" title="">Вакансии</a></li>
<li class="last leaf"><a href="/page/9287">Проводимые конкурсы</a></li>
</ul></li>
<li class="expanded"><a href="/corruption">Противодействие коррупции</a><ul class="menu"><li class="first leaf"><a href="http://fas.gov.ru/corruption/legislative-acts/">Нормативные правовые и иные акты в сфере противодействия коррупции</a></li>
<li class="leaf"><a href="http://fas.gov.ru/corruption/expertise/">Антикоррупционная экспертиза</a></li>
<li class="leaf"><a href="http://fas.gov.ru/corruption/materials/">Методические материалы</a></li>
<li class="leaf"><a href="http://fas.gov.ru/corruption/blanks/">Формы документов, связанных с противодействием коррупции, для заполнения</a></li>
<li class="leaf"><a href="/corruption/gains">Сведения о доходах, расходах, об имуществе и обязательствах имущественного характера</a></li>
<li class="leaf"><a href="/corruption/commission">Комиссия по соблюдению требований к служебному поведению и урегулированию конфликта интересов</a></li>
<li class="leaf"><a href="http://www.fas.gov.ru/anticorruption/feedback/">Обратная связь для сообщений о фактах коррупции</a></li>
<li class="leaf"><a href="/corruption/control_acts">Нормативные правовые и иные акты территориального управления в сфере противодействия коррупции</a></li>
<li class="last leaf"><a href="/corruption/reports">Доклады, отчеты, обзоры</a></li>
</ul></li>
<li class="leaf"><a href="/photogallery">Фотогалерея</a></li>
<li class="expanded"><a href="/node/3786" title="">Обратная связь</a><ul class="menu"><li class="first leaf"><a href="/node/3786">Контактная информация</a></li>
<li class="leaf"><a href="/page/9265">Порядок обращения</a></li>
<li class="leaf"><a href="/page/9282">График приема граждан</a></li>
<li class="leaf"><a href="/page/9268">Порядок обжалования</a></li>
<li class="leaf"><a href="/page/9269">Примеры обращений</a></li>
<li class="leaf"><a href="/page/9266">Реквизиты для оплаты штрафов и пошлин</a></li>
<li class="last leaf"><a href="/webform/5970" title="">Написать в УФАС</a></li>
</ul></li>
<li class="last leaf"><a href="/page/9397">Реестр хозяйствующих субъектов, имеющих долю на рынке определенного товара в размере более чем 35%</a></li>
</ul>  </div>
</div>
        
                  <div class="corruption_interview">
            <a href="/interview/corruption" title="Опрос">
              <img src="/sites/all/themes/fasregion/img/banner/anticorr_005.png" alt="Опрос по противодействию коррупции" />
            </a>
          </div>
        
        <ul class="smo">
          <li><a href="https://twitter.com/rus_fas"><img src="/sites/all/themes/fasregion/img/banner/tw_50x50.png" alt="twitter"></a></li>
          <li><a href="https://www.facebook.com/rus.fas?v=wall"><img src="/sites/all/themes/fasregion/img/banner/fb_50x50.png" alt="facebook"></a></li>
          <li><a href="http://fasovka.livejournal.com/"><img src="/sites/all/themes/fasregion/img/banner/lj_50x50.png" alt="livejournal"></a></li>
          <li><a href="http://vk.com/fas_rus"><img src="/sites/all/themes/fasregion/img/banner/vk_50x50.png" alt="vkontakte"></a></li>
          <li><a href="http://www.youtube.com/user/FASvideotube?feature=guide"><img src="/sites/all/themes/fasregion/img/banner/yout_50x50.png" alt="youtube"></a></li>
        </ul>

        <ul class="banner-ul" style="text-align: center">
          <li>
            <a href="http://fas.gov.ru/overpricing/"><img
                src="/sites/all/themes/fasregion/img/banner/overpricing.png" width="189"></a>
          </li>
          <li>
            <a href="http://www.fas.gov.ru/"><img
                src="/sites/all/themes/fasregion/img/banner/000.png" width="189"></a>
          </li>
          <li>
            <a href="http://www.fas.gov.ru/about/structure/cartel"><img
                src="/sites/all/themes/fasregion/img/banner/001.png"></a>
          </li>
          <li>
            <a href="http://fas.gov.ru/about/list-of-reports/?ct[3]=1&text_search="><img
                src="/sites/all/themes/fasregion/img/banner/010.png"></a>
          </li>
          <li>
            <a href="http://fas.gov.ru/eljournal/"><img
                src="/sites/all/themes/fasregion/img/banner/002.png"></a>
          </li>
          <li>
            <a href="http://emc.fas.gov.ru/"><img
                src="/sites/all/themes/fasregion/img/banner/004.png"></a>
          </li>
          <li>
            <a href="http://anticartel.ru/"><img
                src="/sites/all/themes/fasregion/img/banner/003.png"></a>
          </li>
          <li>
            <a href="http://www.fas.gov.ru/tags/tags_54.html?num=71"><img
                src="/sites/all/themes/fasregion/img/banner/006.png"></a>
          </li>
          <li>
            <a href="http://zakupki.gov.ru/pgz/public/action/search/complaint/run"><img
                src="/sites/all/themes/fasregion/img/banner/007.png"></a>
          </li>
          <li>
            <a href="http://reestr.fas.gov.ru/"><img
                src="/sites/all/themes/fasregion/img/banner/008.png"></a>
          </li>
          <li>
            <a href="http://rnp.fas.gov.ru/"><img
                src="/sites/all/themes/fasregion/img/banner/009.gif"></a>
          </li>
          <li>
            <a href="http://fas.gov.ru/international-partnership/partnership-cis-participants/"><img
                src="/sites/all/themes/fasregion/img/banner/013.png"></a>
          </li>
        </ul>

                  <div id="block-block-3" class="block block-block">

    
  <div class="content">
    
<ul class="banner-ul">
			<li id="banner-11735">
							                  <a href="/banner/click/11735"><img typeof="foaf:Image" src="http://smolensk.fas.gov.ru/sites/smolensk.f.isfb.ru/files/styles/banner/public/banner/2015/04/20/banner_001.png?itok=5Fg1SIxR" alt="ФАС помнит" /></a>
					</li>
			<li id="banner-11677">
							                  <a href="/banner/click/11677"><img typeof="foaf:Image" src="http://smolensk.fas.gov.ru/sites/smolensk.f.isfb.ru/files/styles/banner/public/banner/2015/03/24/banner_journalist_comprtition_001.png?itok=t6G2wx5R" alt="Конкурс ФАС России для журналистов" /></a>
					</li>
	</ul>
  </div>
</div>
        

        <center><!--Openstat-->
<span id="openstat2221695"></span>
<script type="text/javascript">
var openstat = { counter: 2221695, image: 5081, color: "007085", next: openstat };
(function(d, t, p) {
var j = d.createElement(t); j.async = true; j.type = "text/javascript";
j.src = ("https:" == p ? "https:" : "http:") + "//openstat.net/cnt.js";
var s = d.getElementsByTagName(t)[0]; s.parentNode.insertBefore(j, s);
})(document, "script", document.location.protocol);
</script>
<!--/Openstat--></center>
      </div>

    </div>
  </div>
    <div id="content">
            
                    <div class="content-wrapper">
            <div class="element-invisible"><a id="main-content"></a></div>
                    
          
                    
            <div class="region region-content">
    <div id="block-system-main" class="block block-system">

                      
  <div class="content">
    <div id="node-11890"
     class="node node-news node-promoted node-news-full full-node clearfix" about="/news/11890" typeof="sioc:Item foaf:Document">
    <h1  property="dc:title" datatype="">Данные мониторинга средневзвешенных оптово-отпускных цен по состоянию на 11.06.2015</h1>
        <div class="node-tags"><span>Теги:</span> <a href="/%D1%82%D0%B5%D0%B3%D0%B8/%D0%BC%D0%BE%D0%BD%D0%B8%D1%82%D0%BE%D1%80%D0%B8%D0%BD%D0%B3-%D1%86%D0%B5%D0%BD">мониторинг цен</a></div>
  
        <div class="node-tags"><span>Сфера деятельности:</span>  <a href="/taxonomy/term/5">Контроль рекламы и недобросовестной конкуренции</a></div>
      <div class="subtitle oh">
        <div class="submitted left">19 июня 2015, 09:28</div>
        <div class="node-actions right">
            <a href="" class="ins-blog">Вставить в блог</a>
            <a href="#" class="print" onclick="$('.content.cl').print();return false;">Напечатать</a>
            <a href="/rtf/11890" class="rtf">RTF версия</a>
        </div>
    </div>


    <div class="content cl">
              <div class="text-version"><p>В целях недопущения необоснованного роста цен на продовольственные товары, Смоленское УФАС России продолжает проводить&nbsp; еженедельный мониторинг средневзвешенных оптово-отпускных цен&nbsp; одиннадцати производителей и оптовых поставщиков.</p><p>По результатам мониторинга за отчетный период с 05.06.2015 по 11.06.2015 повышения оптово-отпускных цен более чем на 5 % не выявлено ни у одного хозяйствующего субъекта.</p><p>В отчетном периоде отмечено снижение оптово-отпускных цен на кур замороженных (не разделенных на части) &ndash; в среднем 16,3%, на рыбу хек мороженую неразделанную &ndash; 8,6% и на капусту свежую белокочанную &ndash; 8,5%.</p><p>За истекший период повысились оптово-отпускные цены на картофель свежий - 2,8% и на яблоки свежие &ndash; 1,9%.</p><p>По состоянию на 11.06.2015 признаков нарушения антимонопольного законодательства на агропродовольственных рынках Смоленской области не выявлено.</p><p>&nbsp;</p><p>&nbsp;</p><p align="right"><em>Пресс-служба Смоленского УФАС</em></p><p><u>Для справки</u>: перечень продовольственных товаров, подлежащих мониторингу:</p><p>Говядина (в живом весе), свинина свежая или охлажденная (кроме бескостного мяса), свинина замороженная (кроме бескостного мяса), куры свежие или охлажденные (не разделенные на части), куры замороженные (не разделенные на части), рыба хек мороженая неразделанная, рыба минтай мороженая неразделанная, масло сливочное с массовой долей жира&nbsp; от 50% до 85%, молоко пастеризованное жирностью 3,2%, картофель свежий, капуста свежая белокочанная, лук репчатый свежий, морковь свежая, яблоки свежие и крупа гречневая.</p></div>
        <div class="blog-version" style="display:none">
                      <textarea>&lt;div style=&quot;font-size: 11px; font-weight: bold; margin-bottom: 5px; border-bottom: 1px solid #000&quot;&gt;Смоленское УФАС России&lt;/div&gt;&lt;p&gt;В целях недопущения необоснованного роста цен на продовольственные товары, Смоленское УФАС России продолжает проводить&amp;nbsp; еженедельный мониторинг средневзвешенных оптово-отпускных цен&amp;nbsp; одиннадцати производителей и оптовых поставщиков.&lt;/p&gt;&lt;p&gt;По результатам мониторинга за отчетный период с 05.06.2015 по 11.06.2015 повышения оптово-отпускных цен более чем на 5 % не выявлено ни у одного хозяйствующего субъекта.&lt;/p&gt;&lt;p&gt;В отчетном периоде отмечено снижение оптово-отпускных цен на кур замороженных (не разделенных на части) &amp;ndash; в среднем 16,3%, на рыбу хек мороженую неразделанную &amp;ndash; 8,6% и на капусту свежую белокочанную &amp;ndash; 8,5%.&lt;/p&gt;&lt;p&gt;За истекший период повысились оптово-отпускные цены на картофель свежий - 2,8% и на яблоки свежие &amp;ndash; 1,9%.&lt;/p&gt;&lt;p&gt;По состоянию на 11.06.2015 признаков нарушения антимонопольного законодательства на агропродовольственных рынках Смоленской области не выявлено.&lt;/p&gt;&lt;p&gt;&amp;nbsp;&lt;/p&gt;&lt;p&gt;&amp;nbsp;&lt;/p&gt;&lt;p align=&quot;right&quot;&gt;&lt;em&gt;Пресс-служба Смоленского УФАС&lt;/em&gt;&lt;/p&gt;&lt;p&gt;&lt;u&gt;Для справки&lt;/u&gt;: перечень продовольственных товаров, подлежащих мониторингу:&lt;/p&gt;&lt;p&gt;Говядина (в живом весе), свинина свежая или охлажденная (кроме бескостного мяса), свинина замороженная (кроме бескостного мяса), куры свежие или охлажденные (не разделенные на части), куры замороженные (не разделенные на части), рыба хек мороженая неразделанная, рыба минтай мороженая неразделанная, масло сливочное с массовой долей жира&amp;nbsp; от 50% до 85%, молоко пастеризованное жирностью 3,2%, картофель свежий, капуста свежая белокочанная, лук репчатый свежий, морковь свежая, яблоки свежие и крупа гречневая.&lt;/p&gt;&lt;div style=&quot;font-size: 11px; font-weight: bold; margin-top: 5px; border-top: 1px solid #000&quot;&gt;&lt;a href=&quot;http://smolensk.fas.gov.ru/node/11890&quot;&gt;полная версия статьи&lt;/a&gt;&lt;/div&gt;</textarea>
        </div>
          </div>
          <div class="likes">
        <div class="fb_like bottom_like">
            <div class="fb-like" data-href="http://smolensk.fas.gov.ru/news/11890"
                 data-send="false" data-layout="button_count" data-width="210" data-show-faces="false"></div>
        </div>
        <div class="vk_like bottom_like">
            <div id="vk_like"></div>
        </div>
        <div class="tw_like bottom_like">
            <a href="https://twitter.com/share" class="twitter-share-button" data-url="http://smolensk.fas.gov.ru/news/11890" data-lang="ru">Твитнуть</a>
        </div>
    </div>
    <div class="back"><a href="/news">Назад к списку новостей</a></div>
</div>



  </div>
</div>
  </div>
                  </div>
    </div>
  </div>
<div class="bufer"></div>
<div id="footer">
    <div class="border-t-footer"></div>
    <div class="footer-content">
        <div class="left">
          <div class="copy">&copy;

            <a href="http://fas.gov.ru">УФАС России</a> 2015.
            Все права защищены.</div>

            <div class="pad-t-5">При полном или частичном использовании<br/>
ссылка на Смоленское УФАС России обязательна.</div>
          </div>
        <div class="right t-right">
                      <div>
                <span>Адрес:</span> 214000, г. Смоленск, ул. Октябрьской революции, д. 14-а            </div>
            <div class="pad-t-5">
                <span>Телефон/факс:</span> (4812) 38-62-22<br/>
                <span>E-mail:</span> <a
                    href="mailto:to67@fas.gov.ru">to67@fas.gov.ru</a>
            </div>
        </div>
    </div>
</div>
</div><!-- //#page -->

  <script type="text/javascript">
<!--//--><![CDATA[//><!--

          Shadowbox.path = "/sites/all/libraries/shadowbox/";
	      Shadowbox.init(Drupal.settings.shadowbox);
	    
//--><!]]>
</script>
</body>
</html>
`
