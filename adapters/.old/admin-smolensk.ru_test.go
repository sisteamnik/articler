package adapters

import (
	"log"
	"testing"
)

func TestAdminGetDate(t *testing.T) {
	rp := NewAdminParser()
	/*var dates = map[string]time.Time{
		"21 ИЮНЬ 2015 09:46": time.Time{},
	}*/
	date, err := rp.getDate(adminTestArticle)
	if err != nil {
		t.Error(err)
	}
	if DEBUG {
		log.Println(date)
	}

}

func TestAdminIsArticle(t *testing.T) {
	rp := NewAdminParser()
	if !rp.IsArticle("/news/news_11724.html") {
		t.Error("Is article failed")
	}

	if rp.IsArticle("/neasdfs") {
		t.Error("Is article failed")
	}
}

func TestAdminGetDateUpdated(t *testing.T) {
	rp := NewAdminParser()
	var html = `<p style="color: #666666;font-size: 0.917em;text-align: right;">Дата последнего изменения 22.06.2015 16:01</p>`
	date, err := rp.getDate(html)
	if err != nil {
		t.Error(err)
	}
	if DEBUG {
		log.Println(date)
	}
}

//it use metworking
func TestAdminLastArticles(t *testing.T) {
	/*rp := NewAdminParser()
	urls, err := rp.LastArticles()
	if err != nil {
		t.Error(err)
	}
	if DEBUG {
		log.Println("Admin articles",urls)
	}*/
}

func TestAdminParse(t *testing.T) {
	rp := NewAdminParser()
	art, err := rp.Parse([]byte(adminTestArticle))
	if err != nil {
		t.Error(err)
	}
	if DEBUG {
		log.Println(art.Title)
		log.Println(art.Published)
	}
}

var adminTestArticle = `
<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.1//EN" "http://www.w3.org/TR/xhtml11/DTD/xhtml11.dtd">
<html xmlns="http://www.w3.org/1999/xhtml" xml:lang="ru" lang="ru">
    <head>
        <meta http-equiv="content-type" content="text/html; charset=utf-8" />
        <meta name="description" content="Пресс-служба Администрации Смоленской области - НОВОСТИ">
        <meta name="keywords" content="новости, Смоленск, новости Смоленска">
        <title>Смоляне зажгли «Свечи Памяти»   </title>
        <link rel="stylesheet" type="text/css" href="/css/cached_css.css" media="all" />
        <link rel="stylesheet" type="text/css" href="/css/datepicker.css" media="all" />
        <link rel="stylesheet" type="text/css" href="/css/jquery.fancybox.css" media="all" />
        <link rel="stylesheet" type="text/css" href="/css/skins/tango/skin.css" media="all" />
        <link href="/themes/canape1/style.css" type="text/css" rel="stylesheet" media="screen,projection" charset="utf-8" />
        <link rel="shortcut icon" type="image/x-icon" href="/img/3_favicon_(8).ico" />
        <!--[if lte IE 7]>
            <link href="/themes/canape1/main.ie.css" type="text/css" rel="stylesheet" media="screen,projection" charset="utf-8" />
        <![endif]-->
        
        
        
        <script src="/js/jquery.js" type="text/javascript" charset="utf-8"></script>
        <script type="text/javascript" src="/js/ajax.js"></script>
        <script type="text/javascript" src="/js/date.js"></script>
        <script type="text/javascript" src="/js/datepicker.js"></script>
        <script type="text/javascript" src="/js/form_sender.js"></script>
        <script type="text/javascript" src="/js/jquery.fancybox-1.2.1.js"></script>
        <script type="text/javascript" src="/js/main.js"></script>
        <script type="text/javascript" src="/js/jquery.jcarousel.js"></script>
        <script type="text/javascript" src="/js/voting.js"></script>
        <script type="text/javascript" src="/js/ajax.js"></script>
        <script type="text/javascript" src="/js/coockie.js"></script>        

    </head>
    <body>
        <div class="l-container">
            <div class="container__limiter">
                <div class="container__wrapper">
                    <div class="container__page">
                        
                        <div class="l-content" style="background-color: #ffffff;">
                            <div class="content__wrapper" style="position: relative;">
                                <div class="content__center" style="margin: 0;">
                                    <!--Версия для слепых-->
                                    
                                    <div class="content__indent">
                                          
                                        <div class="b-sevice">
                                            <ul>                                      
                                            
                                                      
                                            
                                            <li ><a href="http://admin-smolensk.ru/"><ins></ins>Главная</a></li>
                                            
                                            
                                                      
                                            
                                            <li ><a href="http://admin-smolensk.ru/meropriyatiya/"><ins></ins>Мероприятия</a></li>
                                            
                                            
                                                      
                                            
                                            <li ><a href="http://admin-smolensk.ru/calendar/"><ins></ins>Календарь</a></li>
                                            
                                            
                                                      
                                            
                                            
                                            <li class="on" ><a href="http://admin-smolensk.ru/news/"><ins></ins>Новости</a></li>
                                            
                                                      
                                            
                                            <li ><a href="http://admin-smolensk.ru/mass_media/"><ins></ins>СМИ</a></li>
                                            
                                            
                                                      
                                            
                                            <li ><a href="/pda/"><ins></ins>КПК версия</a></li>
                                            
                                            
                                                      
                                            
                                            <li style="border:none;"><a href="http://admin-smolensk.ru/rss.php"><ins></ins>RSS</a></li>
                                            
                                            
                                            
                                                        
                                            </ul>
                                        </div>                                      
                                                                                 
                                    </div>
                                </div>
                                <div style="position: absolute; top: 15px; right: 25px;">
                                    <a href="http://admin-smolensk.ru"><img src="/img/russian.gif"></a>
                                    <a href="/en/"><img src="/img/english.gif"></a>
                                </div>
                            </div>
                            <div class="content__left">
                                <div class="content__indent">
                                    
                                    &#160;
                                    
                                </div>
                            </div>                            
                            <div class="content__right">
                                <div class="content__indent">
                                    
                                    <div class="b-phone"><!--dhs--><!--dhs--></div>
                                    
                                </div>
                            </div>                            
                        </div>
                        
                        <div class="l-content l-content_lc">
                            <div class="content__wrapper">
                                <div class="content__center">
                                    <div class="content__indent">
                                    

<div class="b-picture" style="background: url(/img/529_banner.gif) 100% 0 no-repeat; height:172px;">
  <div class="picture__wrapper">
    <div class="picture__motto">&nbsp;</div>
  </div>
</div>


                                    <div class="main_info">
                                        <p>
  <span style="font-size: 24px;">Администрация</span></p>
<p>
  <span style="font-size: 24px;"><strong>СМОЛЕНСКОЙ ОБЛАСТИ<br />
  </strong></span></p>
<p>
  &nbsp;</p>
<p>
  <span style="font-size: 18px;">официальный портал органов власти</span></p>

                                    </div>         
                                    <div class="main_info2">Версия для слабовидящих</div>
                                    </div>
                                </div>
                            </div>
                            <div class="content__left">
                                <div class="content__indent">
                                    <div class="b-logo"><a href="http://admin-smolensk.ru"><img src="/images/empty.jpg" /></a></div>
                                </div>
                            </div>                            
                        </div>

                        <div class="l-content">
                            <div class="content__wrapper">
                                <div class="content__center ">
                                    <div class="content__indent " style="padding: 9px 2px;">
                                        <!--
                                            b-path — ""
                                            b-path_vbar — "|"
                                            b-path_middot — "·"
                                            b-path_arrow — "→"
                                        -->
                                        <div class="b-path b-path_arrow">
                                            <ul>
                                            
                                            </ul>
                                        </div>
                                        
                                        <h1>Новости</h1>
                                        <div class="b-editor">
   
<div class="b-news b-news_page">
    <span style="color:#676767; border-right:#676767 solid 1px;">22.06.2015&nbsp;&nbsp;&nbsp;</span>&nbsp;&nbsp;&nbsp;
      <span class="news-title">Смоляне зажгли «Свечи Памяти»   </span>
<div style="margin-top: 10px;">
<p style="text-align: justify;">
  <strong>Ровно 74 года назад, в 4 часа утра 22 июня, началась Великая Отечественная война &ndash; самый кровопролитный военный конфликт в истории человечества, который унес жизни 27 миллионов советских граждан. В память об этой дате жители всей России, а также стран СНГ сегодня зажгли свечи. Смоленщина традиционно присоединилась к международной акции, которая проходит в День памяти и скорби. </strong></p>
<p style="text-align: center;">
  <strong><a href="http://admin-smolensk.ru/img/image/news/2015/6-15/22_6_15svechaa_2.jpg" class="thickbox_resized"><img alt="" src="/photos/resized/22_6_15svechaa_2_400_266.jpg" style="width: 400px; height: 266px;" /></a></strong></p>
<p style="text-align: center;">
  <a href="http://admin-smolensk.ru/img/image/news/2015/6-15/22_6_15svecha20.jpg" class="thickbox_resized"><img alt="" src="/photos/resized/22_6_15svecha20_400_267.jpg" style="width: 400px; height: 267px;" /></a></p>
<p style="text-align: justify;">
  Ранним утром тысячи смолян вышли на улицы, чтобы вспомнить погибших в боях за Родину. Ветераны, представители органов власти региона и муниципальных образований, федеральных структур, священнослужители, активисты молодежных и других общественных объединений собрались на главных площадях городов и сел, воинских мемориалах.</p>
<p style="text-align: center;">
  <a href="http://admin-smolensk.ru/img/image/news/2015/6-15/22_6_15svecha18.jpg" class="thickbox_resized"><img alt="" src="/photos/resized/22_6_15svecha18_400_300.jpg" style="width: 400px; height: 300px;" /></a></p>
<p style="text-align: center;">
  <a href="http://admin-smolensk.ru/img/image/news/2015/6-15/22_6_15svecha19.jpg" class="thickbox_resized"><img alt="" src="/photos/resized/22_6_15svecha19_400_267.jpg" style="width: 400px; height: 267px;" /></a></p>
<p style="text-align: center;">
  <a href="http://admin-smolensk.ru/img/image/news/2015/6-15/22_6_15svechaa_3.jpg" class="thickbox_resized"><img alt="" src="/photos/resized/22_6_15svechaa_3_400_266.jpg" style="width: 400px; height: 266px;" /></a></p>
<p style="text-align: justify;">
  В областном центре акция &laquo;Свеча Памяти&raquo; была организована в сквере Памяти Героев, на воинском кладбище на улице Фрунзе и возле Братской могилы на проезде Маршала Конева. Во всех муниципальных образованиях региона также состоялись заупокойные литии, смоляне приносили к могилам воинов зажжённые свечи и цветы. В память о погибших защитниках Отечества в День памяти и скорби на Смоленщине прошла Минута молчания.</p>
<p style="text-align: center;">
  <a href="http://admin-smolensk.ru/img/image/news/2015/6-15/22_6_15svecha23.jpg" class="thickbox_resized"><img alt="" src="/photos/resized/22_6_15svecha23_400_267.jpg" style="width: 400px; height: 267px;" /></a></p>
<p style="text-align: center;">
  <a href="http://admin-smolensk.ru/img/image/news/2015/6-15/22_6_15svecha11.jpg" class="thickbox_resized"><img alt="" src="/photos/resized/22_6_15svecha11_400_266.jpg" style="width: 400px; height: 266px;" /></a></p>
<p style="text-align: center;">
  <a href="http://admin-smolensk.ru/img/image/news/2015/6-15/22_6_15svecha12.jpg" class="thickbox_resized"><img alt="" src="/photos/resized/22_6_15svecha12_400_266.jpg" style="width: 400px; height: 266px;" /></a></p>
<p style="text-align: center;">
  <a href="http://admin-smolensk.ru/img/image/news/2015/6-15/22_6_15svecha10.jpg" class="thickbox_resized"><img alt="" src="/photos/resized/22_6_15svecha10_400_266.jpg" style="width: 400px; height: 266px;" /></a></p>
<p style="text-align: center;">
  <a href="http://admin-smolensk.ru/img/image/news/2015/6-15/22_6_15svecha16.jpg" class="thickbox_resized"><img alt="" src="/photos/resized/22_6_15svecha16_400_306.jpg" style="width: 400px; height: 306px;" /></a></p>
<p style="text-align: center;">
  <a href="http://admin-smolensk.ru/img/image/news/2015/6-15/22_6_15svecha6.jpg" class="thickbox_resized"><img alt="" src="/photos/resized/22_6_15svecha6_400_266.jpg" style="width: 400px; height: 266px;" /></a></p>
<p style="text-align: center;">
  <a href="http://admin-smolensk.ru/img/image/news/2015/6-15/22_6_15svecha5.jpg" class="thickbox_resized"><img alt="" src="/photos/resized/22_6_15svecha5_400_266.jpg" style="width: 400px; height: 266px;" /></a></p>

</div>
</div>

<br>

<p style="color: #666666;font-size: 0.917em;text-align: right;">Дата последнего изменения 22.06.2015 16:01</p>

<div class="news_detail_back">
<img src="/img/back_pic.gif">&nbsp;<a href="http://admin-smolensk.ru/news/page_0.html">Вернуться к списку</a>
</div>



    
</div>
                                     
                                     </div>   
                                </div>
                            </div>
                            <div class="content__left">
                                <div class="content__indent">
                                  <div class="b-searchsp">
                                      <form id="search_form" method="GET" action="/search/" class="b-search">
                                                                            
                                          <table>
                                              <tr>
                                                  <td class="input">
                                                      <div class="b-input"><input name="search_text" value="Поиск"  onfocus="this.value=''" /></div>
                                                  </td>
                                                  <td class="button">
                                                      <button><img src="/images/searchsp.gif" alt="Поиск" /></button>
                                                  </td>
                                              </tr>
                                          </table>
                                      </form>
                                  </div>
                                  <div class="add_menu">
                                  <a href="http://admin-smolensk.ru/news/news_11724.html?cmd_version=set_sp_version"><img src="/images/sppanel/head-icon1.png"></a>                                  
                                  <a href="/sitemap/"><span title="Карта сайта" style="background:url('/img/sitemap.gif') no-repeat 50% 50%;">&nbsp;&nbsp;&nbsp;</span></a>
                                  <a href="http://admin-smolensk.ru/applications/"><span title="Написать письмо" style="background:url('/img/contacts.gif') no-repeat 50% 50%;">&nbsp;&nbsp;&nbsp;</span></a>
                                  <a href="http://admin-smolensk.ru/subscribe/"><span title="Подписаться на рассылку" style="background:url('/images/subscribe.gif') no-repeat 50% 50%;">&nbsp;&nbsp;&nbsp;</span></a>
                                  <a href="javascript: void();" onclick="window.location=window.location+'?version=print';"><span title="Версия для печати" style="border: 0; background: url(/images/printv.gif) 50% 50% no-repeat;">&nbsp;&nbsp;&nbsp;</span></a>                                                                    
                                  </div>
                                  
                                  <div class="left-top-banners">
                                      
  


                                  </div>
                                    
                                    <div class="b-menu">
                                    <ul class="level-1">                                    
                                    
                                              
                                    
                                    <li class="item-1 first"><span><a href="http://admin-smolensk.ru/authorities/"><ins></ins>Органы власти</a></span></li>
                                    
                                    
                                              
                                    
                                    <li class="item-1"><span><a href="http://admin-smolensk.ru/zakonodatelstvo/"><ins></ins>Законодательство </a></span></li>
                                    
                                    
                                              
                                    
                                    <li class="item-1"><span><a href="http://admin-smolensk.ru/our_region/"><ins></ins>О регионе</a></span></li>
                                    
                                    
                                              
                                    
                                    <li class="item-1"><span><a href="http://ums.admin-smolensk.ru/"><ins></ins>Международное сотрудничество</a></span></li>
                                    
                                    
                                              
                                    
                                    <li class="item-1"><span><a href="http://admin-smolensk.ru/sov_gos_upr/"><ins></ins>Совершенствование государственного управления</a></span></li>
                                    
                                    
                                              
                                    
                                    <li class="item-1"><span><a href="http://its.admin-smolensk.ru/admin_reforma/"><ins></ins>Административная реформа</a></span></li>
                                    
                                    
                                              
                                    
                                    <li class="item-1"><span><a href="http://67.gosuslugi.ru "><ins></ins>Портал государственных услуг </a></span></li>
                                    
                                    
                                              
                                    
                                    <li class="item-1"><span><a href="http://admin-smolensk.ru/goszakaz/"><ins></ins>Госзаказ</a></span></li>
                                    
                                    
                                              
                                    
                                    <li class="item-1"><span><a href="http://zakupki.gov.ru/epz/main/public/home.html"><ins></ins>Реестр госконтрактов</a></span></li>
                                    
                                    
                                              
                                    
                                    <li class="item-1"><span><a href="http://gis.admin-smolensk.ru/ "><ins></ins>ГИС Смоленской области </a></span></li>
                                    
                                    
                                              
                                    
                                    <li class="item-1"><span><a href="http://admin-smolensk.ru/information_systems_list/"><ins></ins>Информационные системы</a></span></li>
                                    
                                    
                                              
                                    
                                    <li class="item-1"><span><a href="http://opendata.admin-smolensk.ru/"><ins></ins>Открытые данные</a></span></li>
                                    
                                    
                                              
                                    
                                    <li class="item-1"><span><a href="http://admin-smolensk.ru/obrascheniya_grazhdan/"><ins></ins>Обращения граждан</a></span></li>
                                    
                                    
                                              
                                    
                                    <li class="item-1"><span><a href="http://admin-smolensk.ru/jurist/"><ins></ins>Бесплатная юридическая помощь</a></span></li>
                                    
                                    
                                              
                                    
                                    <li class="item-1"><span><a href="http://strateg-prf.admin-smolensk.ru/"><ins></ins>Реализация стратегических инициатив Президента РФ в Смоленской области</a></span></li>
                                    
                                    
                                              
                                    
                                    <li class="item-1"><span><a href="http://admin-smolensk.ru/staffing/"><ins></ins>Кадровое обеспечение</a></span></li>
                                    
                                    
                                              
                                    
                                    <li class="item-1"><span><a href="/links/organi_vlasti/"><ins></ins>Ссылки</a></span></li>
                                    
                                    
                                              
                                    
                                    <li class="item-1"><span><a href="http://admin-smolensk.ru/directories/"><ins></ins>Справочники</a></span></li>
                                    
                                    
                                    
                                    </li>
                                                
                                    </ul>
                                    </div>                                    
                                                                         
                                    <div class="b-searchbox">
                                        <form id="search_form" method="GET" action="/search/" class="b-search">
                                            <!--
                                            <h2>Поиcк</h2>
                                            -->                                        
                                            <table>
                                                <tr>
                                                    <td class="input">
                                                        <div class="b-input"><input name="search_text" value="Поиск"  onfocus="this.value=''" /></div>
                                                    </td>
                                                    <td class="button">
                                                        <button><img src="http://admin-smolensk.ru/images/shpic/find_button.gif" alt="Поиск" /></button>
                                                    </td>
                                                </tr>
                                            </table>
                                        </form>
                                    </div>

                                    <!--*left*-->

                                    <div class="b-news b-special">
                                        
  
    <div><p style="">
  <a href="http://www.smoloblduma.ru/"><img alt="" src="/img/image/smolduma.jpg" style="width: 198px; height: 79px; " /></a></p>
<p style="">
  <a href="http://www.finsmol.ru/open" target="_blank"><img alt="Открытый бюджет для граждан" src="/img/image/grafics/open_budget.jpg" style="width: 200px; height: 83px;" /></a></p>
<p style="">
  <a href="http://fkremont.admin-smolensk.ru/" target="_blank"><img alt="НО Региональный фонд капитального ремонта многоквартирных домов Смоленской области" src="/img/image/grafics/fkremont.jpg" style="width: 200px; height: 83px;" /></a></p>
<p style="">
  <a href="http://jkh.admin-smolensk.ru/news.php?rubrica=130&amp;D=11&amp;M=10&amp;Y=2013&amp;id=598" target="_blank"><img alt="Приемная ЖКХ" src="/img/image/grafics/virtual.gif" /></a></p>
<p style="">
  <a href="/socz_proektu/" target="_blank"><img alt="Cоциально значимые проекты религиозных организаций Смоленской области" src="/img/image/grafics/socz_proektu.jpg" style="width: 200px; height: 80px;" /></a></p>
<p style="">
  <a href="http://torgi.gov.ru" target="_blank"><img alt="torgi.gov.ru" src="/img/image/grafics/torgi.jpg" style="width: 198px; height: 58px;" /></a></p>
<p>
  <a href="http://gosuslugi.ru/" target="_blank"><img alt="Единый портал государственных и муниципальных услуг (функций)" src="/img/image/grafics/gosyslugi_fed.jpg" style="width: 198px; height: 78px;" /></a></p>
<p>
  <a href="http://pravo.gov.ru/" target="_blank"><img alt="Официальный интернет-портал правовой информации" src="/img/image/grafics/pravo_gov.jpg" style="width: 198px; height: 44px;" /></a></p>
<p>
  <a href="http://www.smol.ranepa.ru/" target="_blank"><img alt="Смоленский филиал РАНХиГС" src="/img/image/grafics/ranepa.jpg" style="width: 198px; height: 72px;" /></a></p>
</div>
  


                                    </div>
                                           
                                    <div class="b-news">
                                        
                                        <div id="voitingContent"></div>
                                    </div>
                                    <div class="b-news b-special" style="text-align: center;">
                                        <a target="_blank" href="http://www.rbc.ru"><img height="88" border="0" width="120" src="http://pics.rbc.ru/img/grinf/elections3.gif"></a>
                                    </div>
                                                
                                </div>
                            </div>
                            <div class="content__right" >
                                <div class="content__indent" style="padding: 35px 20px 0 0;">
                                    
                                    <div class="b-special">
                                        
  
    <div><p>
  <noindex></noindex></p>
<p style="text-align: center;">
  <a href="http://orphus.ru/" target="_blank"><img alt="" src="/img/image/1_(12).jpg" style="width: 210px; height: 80px;" /></a></p>
<p style="text-align: center;">
  <a href="/70let/"><img alt="70-летие Победы в Великой Отечественной войне" src="/img/image/grafics/70let.gif" style="width: 200px; height: 109px;" /></a></p>
</div>
  


                                        
                                         
                                        
                                            <div class="calendar"> 
                                                 <center>
                                        <div id="date-picker"></div></center>
                                        <script>
                                        $('#date-picker').datePicker({
                                            inline:true, 
                                            startDate: '01/01/1970',
                                            endDate: (new Date()).asString(),
                                            renderCallback:function($td, thisDate, month, year)
                                          {
                                              var d = thisDate.getDate();
                                              $td.bind(
                                                'click',
                                                function()
                                                {
                                                  window.location = '/news/?date='+d+'_'+(Number(month)+1)+'_'+year;
                                                }
                                              );
                                            
                                          }
                                         });
                                        </script>  
                                            </div>
                                        
                                        <div class="references"> 
                                            
  
    <div><p style="text-align: right">
  <a href="http://strateg-prf.admin-smolensk.ru/" target="_blank"><img alt="Реализация стратегических инициатив Президента РФ в Смоленской области" src="/img/image/grafics/trateg_in.jpg" style="width: 200px; height: 83px;" /></a></p>
<p style="text-align: right">
  <a href="http://smolinvest.com/" target="_blank"><img alt="Инвестиционный Портал Смоленской области" src="/img/image/grafics/invest_portal2.jpg" style="width: 200px; height: 83px;" /></a></p>
<p style="text-align: right">
  <a href="http://мфц67.рф/" target="_blank"><img alt="Смоленский многофункциональный центр" src="/img/image/grafics/smol_mfc.jpg" style="width: 200px; height: 83px" /></a></p>
<p style="text-align: right">
  <a href="http://equalsmol.admin-smolensk.ru/" target="_blank"><img alt="Смоленщина без границ" src="/img/image/grafics/equalsmol2.jpg" style="width: 200px; height: 83px" /></a></p>
<p style="text-align: right">
  <a href="/vserossiyskie_i_oblastnie_konkursi/" target="_blank"><img alt="Всероссийские и областные конкурсы" src="/img/image/competitions.gif" style="width: 200px; height: 83px" /></a></p>
<p style="text-align: right">
  <a href="http://smolfond.ru/" target="_blank"><img alt="Объявления по продаже имущества, находящегося в государственной собственности Смоленской области" src="/img/image/grafics/kon_im.jpg" style="width: 200px; height: 83px" /></a></p>
<p style="text-align: right">
  <a href="/vserossiyskie_i_oblastnie_konkursi/sub/" target="_blank"><img alt="Конкурсы на предоставление субсидий" src="/img/image/grafics/kon_sub.jpg" style="width: 200px; height: 83px" /></a></p>
<p style="text-align: right">
  <a href="http://prirod.admin-smolensk.ru/deyatelnost/konkursy-i-aukciony/" target="_blank"><img alt="Конкурсы и аукционы - Недропользование и Водопользование" src="/img/image/grafics/kon_prirod.jpg" style="width: 200px; height: 83px" /></a></p>
<p style="text-align: right">
  <a href="http://uec.admin-smolensk.ru/" target="_blank"><img alt=" Универсальная Электронная Карта  Смоленской области" src="/img/image/grafics/uec.jpg" style="width: 200px; height: 83px" /></a></p>
<p style="text-align: right">
  <a href="http://www.admin-smolensk.ru/esia/" target="_blank"><img alt="Единая система идентификации и аутентификации" src="/img/image/esia.jpg" style="width: 200px; height: 83px" /></a></p>
<p style="text-align: right">
  <a href="http://corruption.admin-smolensk.ru/" target="_blank"><img alt="Противодействие коррупции" src="/img/image/grafics/anticorruption.gif" style="width: 200px; height: 83px" /></a></p>
<p style="text-align: right">
  <a href="http://antiterror.admin-smolensk.ru/" target="_blank"><img alt="Комиссия по Антитеррору" src="/img/image/grafics/antiterror.jpg" style="width: 200px; height: 83px" /></a></p>
<p style="text-align: right">
  <a href="http://www.67.mchs.gov.ru/" target="_blank"><img alt="МЧС" src="/img/image/grafics/m4s.jpg" style="width: 200px; height: 83px" /></a></p>
<p style="text-align: right">
  <a href="http://korso.admin-smolensk.ru" target="_blank"><img alt="Координационное совещание по обеспечению правопорядка в Смоленской области" src="/img/image/grafics/koord_sov.jpg" style="width: 200px; height: 83px" /></a></p>
<p style="text-align: right">
  <a href="http://antinark.admin-smolensk.ru/" target="_blank"><img alt="Антинаркотическая комиссия в Смоленской области " src="/img/image/grafics/antinark2.jpg" style="width: 200px; height: 83px" /></a></p>
<p style="text-align: right">
  <a href="http://admin-smolensk.ru/opros_gl"><img alt="Итоги опроса населения об эффективности деятельности руководителей" src="/img/image/grafics/opros_effect.jpg" style="width: 200px; height: 83px" /></a></p>
<p style="text-align: right">
  <a href="http://admin-smolensk.ru/web-kameri_goroda_smolenska/"><img alt="Wеb-камеры Смоленска" src="/img/image/grafics/web2.jpg" style="width: 200px; height: 83px" /></a></p>
<p style="text-align: right">
  <a href="http://antinark.admin-smolensk.ru/anketa/" target="_blank"><img alt="Опрос Антинаркотической комиссии" src="/img/image/grafics/opros.jpg" style="width: 200px; height: 83px" /></a></p>
<p style="text-align: right">
  <a href="http://admin-smolensk.ru/"><img alt="Погода" src="/img/image/grafics/gis_meteo.jpg" style="width: 200px; height: 83px" /></a></p>
<p style="text-align: right">
  <a href="http://admin-smolensk.ru/valut/"><img alt="ПогодаКурсы валют ЦБ РФ" src="/img/image/grafics/valut.jpg" style="width: 200px; height: 83px" /></a></p>
<p style="text-align: right">
  <a href="http://smolspas.ru/" target="_blank"><img alt="Пожарно-
спасательный
центр" src="/img/image/grafics/radiacija.jpg" style="width: 200px; height: 83px" /></a></p>
</div>
  


                                        </div>
                                        <style type='text/css'>
.gmtbl{background-color: #ffffff; border: 1px solid #ffffff;}
.gmtdttl{font-size: 100%; font-weight: bold; color: #0a58ab; text-align:center;  background-color: #ffffff}
.gmtdtext{font-size: 100%; font-weight: normal; color: #000000;}text-align:center;}
.tddiv { text-align: left !important;}
div#cntdiv{ position:absolute; display:block;}
div#infselectlist{ background-color: #ffffff; border: 1px solid #7c7c7c; width:100%; height:100%; position:relative; top:1; left:0; right:0; visibility:hidden; cursor:pointer;}
div#hiddentl{ position:relative; top:1; left:0; right:0; visibility:hidden;}
</style>

<script language='JavaScript' type='text/javascript' src='http://informer.gismeteo.ru/html/js/showtlist_new.js'></script>
<script language='JavaScript' type='text/javascript' src='http://informer.gismeteo.ru/html/js/ldata_new.js'></script>
<table style="float: right;" border=0 width=190><tr><td>
 <div id='informer1'></div>
<div id='infscript'></div>
<script language='JavaScript' type='text/javascript' src='http://informer.gismeteo.ru/html/2.php?tnumber=1&city0=4239%D0%A1%D0%BC%D0%BE%D0%BB%D0%B5%D0%BD%D1%81%D0%BA&codepg=utf-8&par=4&inflang=rus&domain=ru&vieinf=1&p=1&w=1&tblstl=gmtbl&tdttlstl=gmtdttl&tdtext=gmtdtext&new_scheme=1'></script>
</td></tr></table>
                                        <a href="http://orphus.ru" id="orphus" target="_blank"><img alt="Система Orphus" src="/images/orphus_1px.gif" border="0" width="2" height="80" /></a>
                                    </div>
                                    
                                </div>
                            </div>
                        </div>
                        
                    </div>
                    
                </div>
            </div>
            <div class="b-field b-field_left">
                <div class="field__1"></div>
                <div class="field__2"></div>
            </div>
            <div class="b-field b-field_right">
                <div class="field__1"></div>
                <div class="field__2"></div>
            </div>
        </div> 
<div class="p-container">
            <div class="container__limiter">
                <div class="container__wrapper">
                    <div class="l-grid l-grid_3 l-grid_footer">
                        <div class="grid__wrapper">
                        <div class="footer_layer">
                            <script type="text/javascript">
    $(document).ready(function() {
        jQuery('#mycarousel').jcarousel({
            scroll: 4,
            visible: 4,
            auto: 5,
            animation: 'slow',
            wrap: 'circular'
        });
    });
</script>


<div class="carousel">
    <div class="jcarousel-skin-tango">
        <div class="jcarousel-container">
            <div class="jcarousel-clip" >
                <ul id="mycarousel" class="jcarousel-skin-tango" class="jcarousel-list jcarousel-clip-horizontal">

                    
                    <li><a href="http://incotech.admin-smolensk.ru/udostoveryaschij-centr/" target="_blank"><img src="/img/image/grafics/official_document.jpg" alt="Удостоверяющий центр органов исполнительной власти"></a></li>
                    
                    <li><a href="/free_soft/" target="_blank"><img src="/img/image/grafics/legal_soft.jpg" alt="Рекомендованное к использованию бесплатное ПО"></a></li>
                    
                    <li><a href="/gosudarstvennie_uslugi/" target="_blank"><img src="/img/image/grafics/state_services.jpg" alt="Государственные услуги для граждан и организаций"></a></li>
                    
                    <li><a href="http://www.sudrf.ru/index.php?id=300&amp;amp;court_type=RS&amp;amp;court_subj=67&amp;amp;act=go_search" target="_blank"><img src="/img/image/grafics/femidas_hand.jpg" alt="ГАС РФ Правосудие"></a></li>
                    
                    <li><a href="http://econ.admin-smolensk.ru/~ekon/deyatelnost/investicii/investicionnie_ploschadki/" target="_blank"><img src="/img/image/grafics/invest.jpg" alt="Инвестиционные площадки"></a></li>
                    
                    <li><a href="http://www.nasledie-smolensk.ru/" target="_blank"><img src="/img/image/grafics/cultural_past.jpg" alt="Культурное наследие Земли Смоленской"></a></li>
                    
                    <li><a href="http://visit-smolensk.ru/" target="blank"><img src="/img/terem2.jpg" alt="Смоленский областной туристско-информационный центр «Смоленский терем»"></a></li>
                    
                    <li><a href="http://websprav.admin-smolensk.ru/tour_sp/" target="_blank"><img src="/img/image/grafics/tourist_help.jpg" alt="Туристический справочник Вас приглашает Смоленщина"></a></li>
                    
                    <li><a href="http://www.smolensk-travel.ru/" target="_blank"><img src="/img/image/grafics/smol_tour_port.jpg" alt="Смоленский туристический портал"></a></li>
                    
                    <li><a href="http://вашсмоленск.рф/" target="blank"><img src="/img/3445_smolenskie_klad.jpg" alt="&quot;Ваш Смоленск&quot; Въездной туризм"></a></li>
                    
                    <li><a href="http://websprav.admin-smolensk.ru/nagrad/index.html" target="_blank"><img src="/img/image/grafics/medals.jpg" alt="Система награждения Смоленской области"></a></li>
                    
                    <li><a href="/poslanie_prezidenta_federalnomu_sobraniyu_rossiyskoy_federacii/" target="_blank"><img src="/img/image/grafics/presidents_notes.jpg" alt="Послание Президента Федеральному собранию"></a></li>
                    
                    <li><a href="http://opsmol.admin-smolensk.ru/" target="_blank"><img src="/img/image/grafics/common_cabinet.jpg" alt="Общественная палата Смоленской области"></a></li>
                    
                    <li><a href="http://ombudsmanbiz67.ru/" target="blank"><img src="/img/predpr.jpg" alt="Уполномоченный по защите прав предпринимателей в Смоленской области"></a></li>
                    
                    <li><a href="http://cfo.gov.ru/" target="_blank"><img src="/img/3_cfo.jpg" alt="Центральный федеральный округ "></a></li>
                    
                    <li><a href="http://smolenskcci.ru/" target="_blank"><img src="/img/image/grafics/merchants_headquaters.jpg" alt="Смоленская торгово-промышленная палата"></a></li>
                    
                    <li><a href="http://www.smolensk-notarius.ru/" target="_blank"><img src="/img/image/grafics/lawers_headquaters.jpg" alt="Смоленская областная нотариальная палата"></a></li>
                    
                    <li><a href="http://www.r67.nalog.ru/" target="_blank"><img src="/img/image/grafics/yfns.jpg" alt="УФНС России по Смоленской области"></a></li>
                    
                    <li><a href="http://admin.smolensk.ru/~ufrs/adm_ref.html" target="_blank"><img src="/img/image/grafics/rosreestr.jpg" alt="Создание единого учётно-регистрационного органа Управление Россреестра по Смоленской области"></a></li>
                    
                    <li><a href="http://sror-nps.ru/" target="_blank"><img src="/img/image/grafics/sror_nps.jpg" alt="СРОР НПС"></a></li>
                    
                    <li><a href="http://gossluzhba.gov.ru/" target="_blank"><img src="/img/portal_gsl.jpg" alt="Портал Государственной службы"></a></li>
                    
                    <li><a href="http://apparat.admin-smolensk.ru/index.php?option=com_content&amp;view=article&amp;id=213" target="_blank"><img src="/img/image/grafics/managers_forge.jpg" alt="Президентская программа подготовки управленческих кадров"></a></li>
                    
                    <li><a href="/pochetnie_zvaniya_i_nagradi_gorodov_smolenskoy_oblasti/" target="_blank"><img src="/img/image/grafics/revered_titles.jpg" alt="Почетные звания и награды городов Смоленской области"></a></li>
                    
                    <li><a href="http://www.obd-memorial.ru/" target="_blank"><img src="/img/image/grafics/memorial_db.jpg" alt="Обобщенный банк данных Мемориал"></a></li>
                    
                    <li><a href="http://websprav.admin-smolensk.ru/kniga_pam/" target="_blank"><img src="/img/image/grafics/digital_memory_book.jpg" alt="Электронная книга памяти"></a></li>
                    
                    <li><a href="http://websprav.admin-smolensk.ru/pobeda/" target="blank"><img src="/img/pobeda.jpg" alt="51 том Книги Памяти – результат работы области по увековечиванию смолян"></a></li>
                    
                    <li><a href="http://websprav.admin-smolensk.ru/wow/" target="_blank"><img src="/img/image/grafics/soldatu_pobedu.jpg" alt="Солдаты Победы"></a></li>
                    
                    <li><a href="http://equalsmol.admin-smolensk.ru/" target="_blank"><img src="/img/image/grafics/equalsmol.jpg" alt="Смоленщина без границ"></a></li>
                    
                    <li><a href="http://www.zabota-smolensk.info/" target="_blank"><img src="/img/image/grafics/cearness.jpg" alt="ЗАБОТА Смоленск - Социальный портрет региона"></a></li>
                    
                    <li><a href="/ipoteka/" target="_blank"><img src="/img/image/grafics/home_for_credit.jpg" alt="Ипотечное жилищное кредитование в Смоленской области"></a></li>
                    
                    <li><a href="http://www.km67.ru/" target="_blank"><img src="/img/3_4539_smolizba.jpg" alt="Смоленская Изба - художественная  мастерская гончарного искусства"></a></li>
                    
                    <li><a href="http://www.smolgazeta.ru/" target="_blank"><img src="/img/image/grafics/smol_gazeta.jpg" alt="Смоленская газета"></a></li>
                    
                    <li><a href="http://www.molsm.ru/" target="_blank"><img src="/img/image/grafics/mol_sm.jpg" alt="Молодежный Смоленск"></a></li>
                    
                    <li><a href="http://www.molodezh67.ru/" target="_blank"><img src="/img/image/grafics/smol_adults.jpg" alt="Молодежь Смоленщины"></a></li>
                    
                    <li><a href="http://www.sofpmp.ru/" target="_blank"><img src="/img/image/grafics/fond_predprinim.jpg" alt="Смоленский областной фонд поддержки предпринимательства"></a></li>
                    
                    <li><a href="http://www.r67.fssprus.ru/iss/ip/" target="_blank"><img src="/img/3_isppro.jpg" alt="Банк данных исполнительных производств "></a></li>
                    
                    <li><a href="http://www.smolcity.ru/" target="_blank"><img src="/img/image/grafics/smolcity.jpg" alt="smolcity.ru"></a></li>
                    
                    <li><a href="http://www.rabochy-put.ru/" target="_blank"><img src="/img/3_rabpyt.jpg" alt="Областная общественно-политическая газета &quot;Рабочий путь&quot;"></a></li>
                    
                    <li><a href="http://www.smolgrad.ru/" target="_blank"><img src="/img/image/grafics/our_smolensk.jpg" alt="НАШ ГОРОД Информационный портал"></a></li>
                    
                    <li><a href="http://all-smolensk.ru/" target="_blank"><img src="/img/image/grafics/whole_smolensk.jpg" alt="Справочник Весь Смоленск"></a></li>
                    
                    <li><a href="http://www.vsn-smol.info/" target=""><img src="/img/image/grafics/vsn_smol_info.jpg" alt="ВСН-СМОЛ инфо Информационный портал"></a></li>
                    
                    <li><a href="http://fcdnepr.ru/" target="_blank"><img src="/img/image/grafics/fk_dnepr.jpg" alt="Футбольный клуб ДНЕПР"></a></li>
                    
                    <li><a href="/sluzhba_avtospaseniya/" target="_blank"><img src="/img/image/grafics/god_save_as_everyone.jpg" alt="Служба автоспасения"></a></li>
                    
                    <li><a href="/arbexpo/" target=""><img src="/img/1325_arbexpo.jpg" alt="Выставочный интернет-комплекс Ассоциации российских банков АРБЭКСПО"></a></li>
                    
                    <li><a href="http://admin-smolensk.ru/pages/fingram.html" target=""><img src="/img/1325_fingram.jpg" alt="Марафон финансовой грамотности «Ипотечное кредитование в России. Из региона в регион НОН СТОП!»"></a></li>
                    
                    <li><a href="http://helpinver.ru" target="blank"><img src="/img/helpinver.jpg" alt="Всероссийский проект «Хелпинвер - Открой новую Россию»"></a></li>
                    

                </ul>
            </div>
            <div disabled="disabled" class="jcarousel-prev jcarousel-prev-disabled"></div>
            <div class="jcarousel-next"></div>
        </div>
    </div>
</div>
<div>
    <a href="http://admin-smolensk.ru/vse_resursi/">Все ресурсы</a>
</div>


                        </div>
                            <div class="grid__1">
                                <div class="grid__indent">
                                <p>© Администрация Смоленской области 2010-2015  – официальный сайт </p><br />
<p>Вопросы, предложения и сведения о неполадках работы сайта следует сообщать ведущему специалисту отдела телекоммуникаций, связи и цифрового телевидения Департамента Смоленской области по информационным технологиям Володькину Сергею Александровичу по телефону (4812) 29-20-67 (внутр. 22067) или по электронной почте <a href="mailto:volodkin_sa@admin-smolensk.ru">volodkin_sa@admin-smolensk.ru</a>.</p><br />
<p>«Твинс» — <a href="http://www.web-canape.ru/authorities/">сайты для органов государственной власти</a></p>

                                </div>
                            </div>
                            <div class="grid__2">
                                <div class="grid__indent">

                                    <div class="b-counter">
                                       <div align="center"><a href="http://its.admin-smolensk.ru/">i</a></div>

<noindex>
<table border="0" align="center">
  <tr align="center">
    <td rowspan="3" align="center">
<!--LiveInternet counter-->
<script language="JavaScript">
<!--
document.write('<a href="http://www.liveinternet.ru/click" '+
'target=liveinternet><img src="http://counter.yadro.ru/hit?t27.11;r'+
escape(document.referrer)+((typeof(screen)=='undefined')?'':
';s'+screen.width+'*'+screen.height+'*'+(screen.colorDepth?
screen.colorDepth:screen.pixelDepth))+';'+Math.random()+
'" alt="liveinternet.ru: показано количество просмотров и посетителей" '+
'border=0 width=88 height=120></a>')//
-->
</script>
<!--/LiveInternet--> 
</td>
    <td align="center">
 <a href="http://www.yandex.ru/cy?base=0&host=admin-smolensk.ru"><img
src="http://www.yandex.ru/cycounter?admin-smolensk.ru" width="88" height="31"
alt="Яндекс цитирования" border="0"> </a>
</td>
    <td rowspan="3" align="center"> <!-- HotLog --> <script language="javascript">
hotlog_js="1.0";
hotlog_r=""+Math.random()+"&s=157377&im=201&r="+escape(document.referrer)+"&pg="+
escape(window.location.href);
document.cookie="hotlog=1; path=/"; hotlog_r+="&c="+(document.cookie?"Y":"N");
</script><script
language="javascript1.1">
hotlog_js="1.1";hotlog_r+="&j="+(navigator.javaEnabled()?"Y":"N")</script> <script language="javascript1.2">
hotlog_js="1.2";
hotlog_r+="&wh="+screen.width+'x'+screen.height+"&px="+
(((navigator.appName.substring(0,3)=="Mic"))?
screen.colorDepth:screen.pixelDepth)</script> <script
language="javascript1.3">hotlog_js="1.3"</script> <script language="javascript">hotlog_r+="&js="+hotlog_js;
document.write("<a href='http://click.hotlog.ru/?157377' target='_top'><img "+
" src='http://hit6.hotlog.ru/cgi-bin/hotlog/count?"+
hotlog_r+"&' border=0 width=88 height=100 alt=HotLog></a>")</script> <noscript><a
href="http://click.hotlog.ru/?157377" target="_top"><img
src="http://hit6.hotlog.ru/cgi-bin/hotlog/count?s=157377&im=201" border="0" width="88"
height="100" alt="HotLog"></a></noscript> <!-- /HotLog --><br>
</noindex> </td>
  </tr>
  <tr align="center">
    <td align="center"><!-- begin of Top100 code -->

<script id="top100Counter" type="text/javascript" src="http://counter.rambler.ru/top100.jcn?160503"></script>
<noscript>
<a href="http://top100.rambler.ru/navi/160503/">
<img src="http://counter.rambler.ru/top100.cnt?160503" alt="Rambler's Top100" border="0" />
</a>
</noscript>
<!-- end of Top100 code -->
 </td>
  </tr>
  <tr align="center">
    <td align="center"><noindex> <!-- SpyLOG f:0211 --> <script
language="javascript"><!--
Mu="u2326.85.spylog.com";Md=document;Mnv=navigator;Mp=0;
Md.cookie="b=b";Mc=0;if(Md.cookie)Mc=1;Mrn=Math.random();
Mn=(Mnv.appName.substring(0,2)=="Mi")?0:1;Mt=(new Date()).getTimezoneOffset();
Mz="p="+Mp+"&rn="+Mrn+"&c="+Mc+"&t="+Mt;
if(self!=top){Mfr=1;}else{Mfr=0;}Msl="1.0";
//--></script><script language="javascript1.1"><!--
Mpl="";Msl="1.1";Mj = (Mnv.javaEnabled()?"Y":"N");Mz+='&j='+Mj;
//--></script><script
language="javascript1.2"><!-- 
Msl="1.2";Ms=screen;Mpx=(Mn==0)?Ms.colorDepth:Ms.pixelDepth;
Mz+="&wh="+Ms.width+'x'+Ms.height+"&px="+Mpx;
//--></script><script language="javascript1.3"><!--
Msl="1.3";//--></script><script
language="javascript"><!--
My="";My+="<a href='http://"+Mu+"/cnt?cid=232685&f=3&p="+Mp+"&rn="+Mrn+"' target='_blank'>";
My+="<img src='http://"+Mu+"/cnt?cid=232685&"+Mz+"&sl="+Msl+"&r="+escape(Md.referrer)+"&fr="+Mfr+"&pg="+escape(window.location.href);
My+="' border=0 width=88 height=31 alt='SpyLOG'>";
My+="</a>";Md.write(My);//--></script><noscript> <a
href="http://u2326.85.spylog.com/cnt?cid=232685&f=3&p=0" target="_blank"><img
src="http://u2326.85.spylog.com/cnt?cid=232685&p=0" alt="SpyLOG" border="0" width="88"
height="31"> </a></noscript> <!-- SpyLOG --></td>

  </tr>
</table>
</noindex>
                                    </div>

                                </div>
                            </div>
                            <!--div class="grid__3">
                                <div class="grid__indent">

                                </div>
                            </div-->
                            <div class="grid__3">
                                <div class="grid__indent" style="text-align:right;">
                                    <p>
                                    <p>Адрес: 214008, г. Смоленск, площадь им. Ленина, 1<br />
Телефон/факс справочной службы: (4812) 38-61-65<br />
E-mail: <a href="mailto:region@admin.smolensk.ru">region@admin.smolensk.ru</a></p>
<br>
<noindex>
<a href="http://orphus.ru/" target=_blank">
<img alt="Заметили ошибку?" src="/img/image/orphus.gif" style="width: 240px; height: 80px;" />
</a>

<!-- widget start -->
<script type="text/javascript" src="http://gosmonitor.ru/widget.js?id=267"></script>
<!-- / widget end -->
</noindex>
                                    </p>

                                </div>
                            </div>
                        </div>
                    </div>

                </div>
            </div>
            <div class="x-field x-field_left">
                <div class="field__1"></div>
                <div class="field__2"></div>
            </div>
            <div class="x-field x-field_right">
                <div class="field__1"></div>
                <div class="field__2"></div>
            </div>
        </div>
    <script type="text/javascript" src="/js/orphus.js"></script>
    </body>
</html>
`
