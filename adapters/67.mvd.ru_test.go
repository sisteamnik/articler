package adapters

import (
	"log"
	"testing"
)

func TestMvdGetDate(t *testing.T) {
	rp := NewMvdParser()
	/*var dates = map[string]time.Time{
		"21 ИЮНЬ 2015 09:46": time.Time{},
	}*/
	date, err := rp.getDate("Сегодня 17:10")
	if err != nil {
		t.Error(err)
	}
	if DEBUG {
		log.Println(date)
	}

	date, err = rp.getDate("16 Июня 17:17")
	if err != nil {
		t.Error(err)
	}
	if DEBUG {
		log.Println(date)
	}
	//todo implemet year
	date, err = rp.getDate("18 Июня 2014 17:41")
	if err == nil {
		t.Error(err)
	}

}

func TestMvdIsArticle(t *testing.T) {
	rp := NewMvdParser()
	if !rp.IsArticle("/news/item/3599309/") {
		t.Error("Is article failed")
	}

	if rp.IsArticle("/news/dsfd/") {
		t.Error("Is article failed")
	}
}

//it use metworking
func TestMvdLastArticles(t *testing.T) {
	/*rp := NewMvdParser()
	urls, err := rp.LastArticles()
	if err != nil {
		t.Error(err)
	}
	if DEBUG {
		log.Println("Mvd articles", urls)
	}*/
}

func TestMvdParse(t *testing.T) {
	rp := NewMvdParser()
	art, err := rp.Parse([]byte(MvdTestArticle))
	if err != nil {
		t.Error(err)
	}
	if len(art.Body) < 10 {
		t.Error("to short")
	}
	if DEBUG {
		log.Println("Mvd parsing", art.Title)
		log.Println(art.Published)
	}
}

var MvdTestArticle = `
<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
<html xmlns="//www.w3.org/1999/xhtml" xml:lang="ru">
<head><title>24 июня в УМВД России по Смоленской области состоится «прямая линия» с гражданами по вопросам антикоррупционного просвещения, отнесенным к сфере деятельности органов внутренних дел</title>
<meta http-equiv="Content-Type" content="text/html; charset=UTF-8" />
<meta name = "viewport" content = "maximum-scale = 1.0, width = device-width">
<meta name="keywords" content=""/>
<meta name="description" content=""/>
<meta http-equiv="X-UA-Compatible" content="IE=edge,chrome=1"/>


<link rel="alternate" type="application/rss+xml" title="Новости" href="/news/rss/" />
<link rel="apple-touch-icon" href="/favico.ico">
<link rel="icon" href="/favico.ico" type="image/x-icon">
<link rel="shortcut icon" href="/favico.ico" type="image/x-icon">

<script>
    !function(d,s,id){var js,fjs=d.getElementsByTagName(s)[0];if(!d.getElementById(id)){js=d.createElement(s);js.id=id;js.src="//platform.twitter.com/widgets.js";fjs.parentNode.insertBefore(js,fjs);}}(document,"script","twitter-wjs");
</script>

<script type="text/javascript" src="/media/mvd-2014/js/jquery-1.10.2.min.js"></script><script type="text/javascript" src="/media/mvd-2014/js/f.js?20150313"></script><script type="text/javascript" src="/media/mvd-2014/js/jquery.mcustomscrollbar.min.js"></script><script type="text/javascript" src="/media/mvd-2014/js/jquery.slider2.js"></script><script type="text/javascript" src="/media/mvd-2014/js/jquery.mediaelement-and-player.min.js"></script><script type="text/javascript" src="/media/mvd-2014/js/overlay.js"></script><script type="text/javascript" src="/media/mvd-2014/js/jquery.columnizer.js"></script><script type="text/javascript" src="/media/mvd-2014/js/jquery.mousewheel.min.js"></script><link type="text/css" href="/media/mvd-2014/css/old/add.css?2015032415" rel="stylesheet" /><link type="text/css" href="/media/mvd-2014/css/old/colorbox.css" rel="stylesheet" /><link type="text/css" href="/media/mvd-2014/css/old/default.css?20150313" rel="stylesheet" /><link type="text/css" href="/media/mvd-2014/css/old/general.css?20150218" rel="stylesheet" /><link type="text/css" href="/media/mvd-2014/css/old/jquery-ui-1.7.2.custom.css" rel="stylesheet" /><link type="text/css" href="/media/mvd-2014/css/old/media-tabs.css" rel="stylesheet" /><script type="text/javascript" src="/media/mvd-2014/js/old/add.js?20140826"></script><script type="text/javascript" src="/media/mvd-2014/js/old/interface.js?20140821"></script><script type="text/javascript" src="/media/mvd-2014/js/old/jquery.colorbox-min.js"></script><script type="text/javascript" src="/media/mvd-2014/js/old/jquery.slider.js"></script><script type="text/javascript" src="https://api-maps.yandex.ru/2.0/?load=package.standard&amp;lang=ru-RU"></script><link type="text/css" href="/media/mvd-2014/css/holster.css?20150526" rel="stylesheet" /><link type="text/css" href="/media/mvd-2014/css/blocks.css?20150601" rel="stylesheet" /><link type="text/css" href="/media/mvd-2014/css/mcustomscrollbar.css" rel="stylesheet" /><link type="text/css" href="/media/mvd-2014/css/mediaelement-and-player.css" rel="stylesheet" /><link type="text/css" href="/media/mvd-2014/css/timeTo.css?20140529" rel="stylesheet" /><script src="//code.jquery.com/ui/1.10.4/jquery-ui.js"></script>
<!--[if IE 8]>
<link type="text/css" href="/media/mvd-2014/css/style-ie8.css" rel="stylesheet" /><![endif]-->
<!--[if IE 6]>
<link type="text/css" href="/media/mvd/css/ie6.css" rel="stylesheet" /><![endif]-->


<script type="text/javascript">
    var site_lang  = "ru";
    var text_print = "Печать";
    var text_back = "Вернуться";
    var text_address = "Адрес данной страницы в интернете";
    var text_mvd = "Официальный сайт Министерства внутренних дел Российской Федерации";
    var text_copy = '© 2011, МВД России  <a href="#copyright">Об использовании информации сайта</a>';
    var text_rights = "Все права охраняются законодательством Российской Федерации";
</script>
<link type="text/css" href="/media/regions/css/add.css" rel="stylesheet" />


</head>
<body>

<div class="ln-page">
    <div class="ln-wrapper">

        <div class="ln-header wrapper">
    <div class="bn-logo" style="top: 25px;"><a href="/"><img src="/upload/site68/4eqxWWT3rf.png" alt="" style="max-width: 140px; max-height: 110px;" /></a></div>
    <div class="bn-logo-name">
        <span>Управление МВД России <br>по Смоленской области</span>
    </div>
    <div class="bn-logo-text">Служим России, служим закону!</div>
    <ul class="bn-links no-style">
            <li>
                                            <a href="//67.mvd.ru/banners/redirect?bid=7417" target="_blank">
                    МВД России                 </a>
                    </li>
            <li>
                                            <a href="//67.mvd.ru/banners/redirect?bid=6870" target="_blank">
                    Наши проекты                </a>
                    </li>
            <li>
                                            <a href="//67.mvd.ru/banners/redirect?bid=7220" target="_blank">
                    Наши герои                </a>
                    </li>
            <li>
                                            <a href="//67.mvd.ru/banners/redirect?bid=7222" target="_blank">
                    Госуслуги                </a>
                    </li>
    </ul>    <ul class="bn-special no-style">
        <li><a class="ico-but b_hb5" target="_blank" href="/radio.html" onclick="popupWin = window.open(this.href, 'win', 'status=no,toolbar=no,scrollbars=no,titlebar=no,menubar=no,resizable=yes,width=287,height=181,directories=no,location=no');popupWin.focus(); return false;"></a></li>
                                            <li><a class="ico-but b_tw2" target="_blank" href="https://twitter.com/mvd67">Tw</a></li>
                                <li><a class="ico-but b_hb3" target="_blank" href="/news/rss">RSS</a></li>
    </ul>

    <img style="position: absolute; right: 200px; top: 62px;" src="/media/mvd-2014/img/102ru.png" alt="">

    <img src="/upload/site68/OeohMJW3TR.jpg" style="position: absolute; right: 0; max-width: 170px; max-height: 110px; bottom: 15px;"></div>        <div class="bn-top-menu no-style" >
    <div class="wrapper">
        <div class="bm-links f-right">
            <div id="search-1-holder" class="bn-search">
                <form method="get" action="/search/">
                    <input id="search-1" type="text" value="" name="q" /><button id="search-1-but" class="ico-but b_sr">Найти</button>
                    <div class="bs-ext">
                        <a class="link red-text" href="/search/">Расширенный поиск</a>
                    </div>
                </form>
            </div>
        </div>
        <ul id="menu-1">
                            <li><a href="/gumvd/rukovodstvo/">УМВД</a></li>
                            <li><a href="/slujba/priem/">Деятельность</a></li>
                            <li><a href="/appeals">Для граждан</a></li>
                            <li><a href="/Kontakti">Контакты</a></li>
                            <li><a href="/folder/2917316">Пресс-служба </a></li>
                    </ul>
    </div>
</div>
<div class="bn-top-submenu no-style">
    <div id="menu-1-sub" class="bm-row">
                    <div class="bm-cell">
                <ul>
                                            <li><a href="/gumvd/rukovodstvo">Руководство</a></li>
                                            <li><a href="/gumvd/structure">Структура</a></li>
                                            <li><a href="/gumvd/Koordinacionnie_i_soveshhatelnie_organi">Координационные и совещательные органы</a></li>
                                            <li><a href="/gumvd/Obshhestvennij_sovet_pri_UMVD_po_Smolens">Общественный совет при УМВД по Смоленской области</a></li>
                                            <li><a href="/folder/923374">Противодействие коррупции</a></li>
                                            <li><a href="/gumvd/Nashi_geroi">Наши герои</a></li>
                                    </ul>
            </div>
                    <div class="bm-cell">
                <ul>
                                            <li><a href="http://mvd.ru/plan_mvd_2018">План 2013-2018</a></li>
                                            <li><a href="/slujba/Sluzhba._Vakansii_FGGS">Служба. Вакансии ФГГС. Учебные заведения</a></li>
                                            <li><a href="/slujba/Dejatelnost">Отчеты должностных лиц</a></li>
                                            <li><a href="/slujba/Plani_i_rezultati_proverok">Планы и результаты проверок</a></li>
                                            <li><a href="/slujba/Vzaimodejstvie_s_organami_ispolnitelnoj">Взаимодействие с органами исполнительной власти субъектов и МСУ</a></li>
                                    </ul>
            </div>
                    <div class="bm-cell">
                <ul>
                                            <li><a href="/request_main">Прием обращений</a></li>
                                            <li><a href="/citizens/graph">Графики приема граждан</a></li>
                                            <li><a href="/citizens/Obzori_obrashhenij_grazhdan">Работа с обращениями граждан</a></li>
                                            <li><a href="/citizens/recommendation">Полиция рекомендует</a></li>
                                            <li><a href="/citizens/GOSUSLUGI">Государственные услуги</a></li>
                                            <li><a href="/citizens/PRAVINF">Правовое информирование</a></li>
                                            <li><a href="/citizens/Vnimanie_rozisk">Внимание, розыск!</a></li>
                                            <li><a href="/citizens/Im_blagodarna_policija">Им благодарна полиция</a></li>
                                            <li><a href="/citizens/Informacija_dlja_lic_s_ogranichennimi_vo">Информация для лиц с ограниченными возможностями</a></li>
                                    </ul>
            </div>
                    <div class="bm-cell">
                <ul>
                                            <li><a href="/Kontakti/contact">Приемная</a></li>
                                            <li><a href="/Kontakti/02">Экстренный вызов</a></li>
                                            <li><a href="http://mvd.ru/help/district">Ваш участковый/отдел полиции</a></li>
                                            <li><a href="/Kontakti/units">Территориальные подразделения </a></li>
                                            <li><a href="/Kontakti/Organizacii_i_uchrezhdenija">Организации и учреждения</a></li>
                                    </ul>
            </div>
                    <div class="bm-cell">
                <ul>
                                            <li><a href="/Press_sluzhba/press">Пресс-служба</a></li>
                                            <li><a href="/Press_sluzhba/smi">Контакты</a></li>
                                            <li><a href="/folder/1305857">Наши проекты </a></li>
                                            <li><a href="/Press_sluzhba/MVD_v_socialnih_setjah">МВД России в социальных медиа</a></li>
                                            <li><a href="/Press_sluzhba/Vedomstvennie_SMI">Ведомственные СМИ</a></li>
                                            <li><a href="/Press_sluzhba/Fotoarhiv">Фотогалерея</a></li>
                                            <li><a href="/Press_sluzhba/video">Видеосюжеты</a></li>
                                            <li><a href="/Press_sluzhba/Pravovoj_likbez">Правовой ликбез</a></li>
                                    </ul>
            </div>
            </div>
</div>        <div class="ln-content wrapper old_content">
            <div class="ln-content-center">
                                    <div class="ln-content-holder">
                                    
<div class="bn-breadcrumb margin1 links">
                            <a href="/">Главная</a>
            <span class="bb-sep">&nbsp;&rarr;&nbsp;</span>
                                    Новости                    </div>                    <h1>24 июня в УМВД России по Смоленской области состоится «прямая линия» с гражданами по вопросам антикоррупционного просвещения, отнесенным к сфере деятельности органов внутренних дел</h1>
<div class="article-date-item">Сегодня 17:10</div>
<div class="article">
    
    <p style="outline: none; margin: 10px 0px; line-height: 23.7999992370605px; text-align: justify; color: #000000; font-family: 'PT Sans', Arial, Helvetica, sans-serif; font-size: 14px; font-style: normal; font-variant: normal; font-weight: normal; letter-spacing: normal; orphans: auto; text-indent: 0px; text-transform: none; white-space: normal; widows: 1; word-spacing: 0px; -webkit-text-stroke-width: 0px; background-color: #ffffff;">24 июня 2015 года, с 15.00 до 16.00 по телефону 8-910-767-77-31 в УМВД России по Смоленской области состоится &laquo;прямая линия&raquo; с гражданами по вопросам антикоррупционного просвещения, отнесенным к сфере деятельности органов внутренних дел.</p>
<p style="outline: none; margin: 10px 0px; line-height: 23.7999992370605px; text-align: justify; color: #000000; font-family: 'PT Sans', Arial, Helvetica, sans-serif; font-size: 14px; font-style: normal; font-variant: normal; font-weight: normal; letter-spacing: normal; orphans: auto; text-indent: 0px; text-transform: none; white-space: normal; widows: 1; word-spacing: 0px; -webkit-text-stroke-width: 0px; background-color: #ffffff;">Отвечать на вопросы граждан будут представители&nbsp;управления экономической безопасности и противодействия коррупции, оперативно-разыскной части службы собственной безопасности,&nbsp;правовой службы и &nbsp;управления по работе с личным составом УМВД России по Смоленской области.&nbsp;&nbsp;</p>
<p style="outline: none; margin: 10px 0px; line-height: 23.7999992370605px; text-align: justify; color: #000000; font-family: 'PT Sans', Arial, Helvetica, sans-serif; font-size: 14px; font-style: normal; font-variant: normal; font-weight: normal; letter-spacing: normal; orphans: auto; text-indent: 0px; text-transform: none; white-space: normal; widows: 1; word-spacing: 0px; -webkit-text-stroke-width: 0px; background-color: #ffffff;">Позвонив по телефону 8-910-767-77-31, граждане смогут задать интересующие их вопросы на тему противодействия коррупции и получить информацию, в том числе, о правах и обязанностях полицейских, уточнить законность действий сотрудников полиции в тех или иных жизненных ситуациях, а также сообщить о совершении правонарушений и противоправных действий сотрудниками органов внутренних дел.</p>
<p style="outline: none; margin: 10px 0px; line-height: 23.7999992370605px; text-align: left; color: #000000; font-family: 'PT Sans', Arial, Helvetica, sans-serif; font-size: 14px; font-style: normal; font-variant: normal; font-weight: normal; letter-spacing: normal; orphans: auto; text-indent: 0px; text-transform: none; white-space: normal; widows: 1; word-spacing: 0px; -webkit-text-stroke-width: 0px; background-color: #ffffff;"><strong style="outline: none; line-height: 1.7em;">Пресс-служба УМВД России по Смоленской области</strong>&nbsp;&nbsp;</p>
    
<div class="border more_links social">
    <ul class="links">
        <li class="ico ico_2"><a href="/news/item/3599309/?print=1" target="_blank" id="print">Версия для печати</a></li>
        <li class="ico ico_4"><a href="javascript:void(0);" id="link_page_a">Ссылка на страницу</a></li>
                <li class="ico ico_3"><a href="javascript:void(0);" onclick="showFormSendLinkToEmail()">Отправить по почте</a></li>

    </ul>

    <div id="link_page" style="left: 701px; top: 754px; display: none; ">
        <div class="padding">
            <h3>Ссылка на страницу</h3>
            <a href="javascript:void(0);" class="close">закрыть</a>
            <form action="subscribe" method="post">
                <table>
                    <tbody><tr>
                        <td><input value="" type="text" class="input" onclick="this.select()"></td>

                    </tr>
                    </tbody></table>
            </form>
        </div>
    </div>

    <div id="subscribe_b_send_email">
        <div class="padding">
            <h3>Отправить ссылку на E-mail</h3>
            <a href="javascript:void(0);" class="close" onclick="closeFormSendLinkToEmail()">закрыть</a>
            <div class="ajax-loader-send-link" style="text-align: center; display: none;">
                <img src="/media/default/img/ajax-loader.gif">
            </div>
            <div class="bsl-send-success">Ссылка успешно отправлена!</div>
            <div class="bsl-form">
                <div class="rgf-line">
                    <label class="rgf-label">Ссылка</label>
                    <div class="rgf-inpit ff-inp"><input type="text" name="link" disabled value="https://67.mvd.ru/news/item/3599309" id="send_link_item" /></div>
                </div>
                <div class="rgf-line">
                    <label class="rgf-label">E-mail <span class="rgf-required" style="color: #E68D00; display: inline-block;">*</span></label>
                    <div class="rgf-inpit ff-inp"><input type="text" name="email" value="" id="send_link_email"/></div>
                </div>
                <div class="rgf-line">
                    <label class="rgf-label">Комментарии</label>
                    <div class="rgf-inpit ff-inp"><textarea name="comments" id="send_link_comments"></textarea></div>
                </div>
                <div class="rgf-line" >
                    <label class="rgf-label">&nbsp;</label>
                    <input type="submit" class="quiz-submit rgf-submit send-link-submit" value="Отправить" onclick="sendLinkToEmail()"/>
                </div>
            </div>

        </div>
    </div>

    <div class="clear"></div>
</div>
<div class="clearfix"></div>
<table width="100%" cellspacing="0" cellpadding="0" class="no_border table-links-layout">
    <tbody><tr valign="center" align="center" style="border:0px;background:none;">
        <td width="150" valign="center" class="odd links-title">
            <span style="color:#676767;">Поместить ссылку в</span>
        </td>

        
        <td width="34" valign="center" align="center"  class="social_icon odd">
            <noindex><a onClick="popup('https://www.facebook.com/sharer.php?u='+'https://67.mvd.ru/news/item/3599309&t=24 июня в УМВД России по Смоленской области состоится «прямая линия» с гражданами по вопросам антикоррупционного просвещения, отнесенным к сфере деятельности органов внутренних дел')" target="_blank" title="Facebook"><img src="/media/default/img/social_fb.png" title="Facebook"></a></noindex>
        </td>

        <td width="33" valign="center" align="center" class="social_icon odd">
            <noindex><a class="ico-but b_stw" style="background-position: -20px -160px;" onclick="popup('https://twitter.com/share?url=https://67.mvd.ru/news/item/3599309&text=24 июня в УМВД России по Смоленской области состоится «прямая линия» с гражданами по вопросам антикоррупционного просвещения, отнесенным к сфере деятельности органов внутренних дел')" target="_blank" title="Twitter"></a></noindex>
        </td>

        <td width="33" valign="center" align="center" class="social_icon odd">
            <noindex><a onclick="popup('https://vkontakte.ru/share.php?url=https://67.mvd.ru/news/item/3599309&title=24 июня в УМВД России по Смоленской области состоится «прямая линия» с гражданами по вопросам антикоррупционного просвещения, отнесенным к сфере деятельности органов внутренних дел')" target="_blank" title="ВКонтакте"><img src="/media/default/img/social_vk.png" title="ВКонтакте"></a></noindex>
        </td>

        <td width="33" valign="center" align="center" class="social_icon odd">
            <script src="/media/mvd-2014/js/soc/odkl_share.js" type="text/javascript" ></script>
            <noindex><a href="https://67.mvd.ru/news/item/3599309" onclick="ODKL.Share(this);return false;" ><img width="20px" height="20px" src="/media/mvd/img/test1.png" title="Однокласники+"></a></noindex>
        </td>

        <td width="33" valign="center" align="center" class="social_icon odd">
            <noindex><a href="https://www.livejournal.com/update.bml?event=https://67.mvd.ru/news/item/3599309&;subject=24 июня в УМВД России по Смоленской области состоится «прямая линия» с гражданами по вопросам антикоррупционного просвещения, отнесенным к сфере деятельности органов внутренних дел" target="_blank" title="LiveJournal"><img src="/media/mvd/images/social/social_lj.png" title="LiveJournal"></a></noindex>
        </td>
        <td width="33" valign="center" align="center" class="social_icon odd">
            <noindex><a href="https://plus.google.com/share?url=https://67.mvd.ru/news/item/3599309" target="_blank" title="Google+"><img width="20px" height="20px" src="/media/mvd/images/social/social_gp.png" title="Google+"></a></noindex>
        </td>
        <td width="33" valign="center" align="center" class="social_icon odd">
            <noindex><a href="https://www.myspace.com/Modules/PostTo/Pages/?u=https://67.mvd.ru/news/item/3599309&;t=24 июня в УМВД России по Смоленской области состоится «прямая линия» с гражданами по вопросам антикоррупционного просвещения, отнесенным к сфере деятельности органов внутренних дел" target="_blank" title="MySpace"><img src="/media/mvd/images/social/social_mysp.png" title="MySpace"></a></noindex>
        </td>
        <td width="33" valign="center" align="center" class="social_icon odd">
            <noindex><a href="https://www.linkedin.com/shareArticle?mini=true&;url=https://67.mvd.ru/news/item/3599309&;title=24 июня в УМВД России по Смоленской области состоится «прямая линия» с гражданами по вопросам антикоррупционного просвещения, отнесенным к сфере деятельности органов внутренних дел" target="_blank" title="Linked In"><img src="/media/mvd/images/social/social_lin.png" title="Linked In"></a></noindex>
        </td>
        <td width="33" valign="center" align="center" class="social_icon odd">
            <noindex><a href="https://bookmarks.yahoo.com/toolbar/savebm?u=https://67.mvd.ru/news/item/3599309&;t=24 июня в УМВД России по Смоленской области состоится «прямая линия» с гражданами по вопросам антикоррупционного просвещения, отнесенным к сфере деятельности органов внутренних дел" target="_blank" title="Yahoo! Bookmarks"><img src="/media/mvd/images/social/social_yb.png" title="Yahoo! Bookmarks"></a></noindex>
        </td>

        <!--<td width="33" valign="center" align="center" class="social_icon odd">
            <noindex><a href="http://www.google.com/reader/link?url=&;title=" target="_blank" title="Google Reader"><img src="/media/mvd/images/social/social_gr.png" title="Google Reader"></a></noindex>
        </td>-->

        <td width="33" valign="center" align="center" class="social_icon odd">
            <noindex><a href="https://connect.mail.ru/share?share_url=https://67.mvd.ru/news/item/3599309" target="_blank" title="Мой мир на Mail.ru"><img src="/media/mvd/images/social/social_mail.png" title="Мой мир на Mail.ru"></a></noindex>
        </td>

        <!--<td width="33" valign="center" align="center" class="social_icon odd">
            <noindex><a href="http://zakladki.yandex.ru/userarea/links/addfromfav.asp?bAddLink_x=1&;lurl=&;lname=" target="_blank" title="Яндекс.Закладки"><img src="/media/mvd/images/social/social_ya.png" title="Яндекс.Закладки" style="padding-right: 4px;"></a></noindex>
        </td>-->

    </tr>
    </tbody></table>

<div style="padding: 10px 0;">
        <a href="https://twitter.com/mvd_official" class="twitter-follow-button" data-show-count="false">Читать @mvd_official</a>
    </div>
</div>                </div>
            </div>
                            <div class="ln-content-right">
                    
<div class="bn-filter type-5 margin2" id="calendar-block">
    <div class="bf-holder">
        <div class="bf-item">
            Год:
            <div class="n-select type-3">
                <span class="select-input select-input-holder" id="year">2014</span>
                <div class="select-list calendar">
                                            <label ><input type="radio" value="2006" name="year" />2006</label>
                                            <label ><input type="radio" value="2011" name="year" />2011</label>
                                            <label ><input type="radio" value="2012" name="year" />2012</label>
                                            <label ><input type="radio" value="2013" name="year" />2013</label>
                                            <label  class="active" ><input type="radio" value="2014" name="year" />2014</label>
                                            <label ><input type="radio" value="2015" name="year" />2015</label>
                                    </div>
            </div>
        </div>
        <div class="bf-item">
            Месяц:
            <div class="n-select type-3">
                <span class="select-input select-input-holder">Июнь</span>
                <div class="select-list calendar" id="month">
                                            <label ><input type="radio" value="1" name="month" />Январь</label>
                                            <label ><input type="radio" value="2" name="month" />Февраль</label>
                                            <label ><input type="radio" value="3" name="month" />Март</label>
                                            <label ><input type="radio" value="4" name="month" />Апрель</label>
                                            <label ><input type="radio" value="5" name="month" />Май</label>
                                            <label  class="active" ><input type="radio" value="6" name="month" />Июнь</label>
                                            <label ><input type="radio" value="7" name="month" />Июль</label>
                                            <label ><input type="radio" value="8" name="month" />Август</label>
                                            <label ><input type="radio" value="9" name="month" />Сентябрь</label>
                                            <label ><input type="radio" value="10" name="month" />Октябрь</label>
                                            <label ><input type="radio" value="11" name="month" />Ноябрь</label>
                                            <label ><input type="radio" value="12" name="month" />Декабрь</label>
                                    </div>
            </div>
        </div>
    </div>

    <div class="bf-holder2">
        <div class="bn-calendar">
            
<table class="bc-day">
    <thead>
    <tr>
        <th scope="col" abbr="Пн" title="Пн" class="weekday">Пн</th>
        <th scope="col" abbr="Вт" title="Вт" class="weekday">Вт</th>
        <th scope="col" abbr="Ср" title="Ср" class="weekday">Ср</th>
        <th scope="col" abbr="Чт" title="Чт" class="weekday">Чт</th>
        <th scope="col" abbr="Пт" title="Пт" class="weekday">Пт</th>
        <th scope="col" abbr="Сб" title="Сб" class="weekend">Сб</th>
        <th scope="col" abbr="Вс" title="Вс" class="weekend">Вс</th>
    </tr>
    </thead>
    <tbody>
            <tr>
                                                <td></td>
                                                                    <td></td>
                                                                    <td></td>
                                                                    <td></td>
                                                                    <td></td>
                                                                    <td></td>
                                                                                        <td class=""><a
                            href="/news?year=2014&month=6&day=1">1</a>
                    </td>
                                        </tr>
            <tr>
                                                                    <td class=""><a
                            href="/news?year=2014&month=6&day=2">2</a>
                    </td>
                                                                                        <td class=""><a
                            href="/news?year=2014&month=6&day=3">3</a>
                    </td>
                                                                                        <td class=""><a
                            href="/news?year=2014&month=6&day=4">4</a>
                    </td>
                                                                                        <td class=""><a
                            href="/news?year=2014&month=6&day=5">5</a>
                    </td>
                                                                                        <td class=""><a
                            href="/news?year=2014&month=6&day=6">6</a>
                    </td>
                                                                                        <td class=""><a
                            href="/news?year=2014&month=6&day=7">7</a>
                    </td>
                                                                                        <td class=""><a
                            href="/news?year=2014&month=6&day=8">8</a>
                    </td>
                                        </tr>
            <tr>
                                                                    <td class=""><a
                            href="/news?year=2014&month=6&day=9">9</a>
                    </td>
                                                                                        <td class=""><a
                            href="/news?year=2014&month=6&day=10">10</a>
                    </td>
                                                                                        <td class=""><a
                            href="/news?year=2014&month=6&day=11">11</a>
                    </td>
                                                                                        <td class=""><a
                            href="/news?year=2014&month=6&day=12">12</a>
                    </td>
                                                                                        <td class=""><a
                            href="/news?year=2014&month=6&day=13">13</a>
                    </td>
                                                                                        <td class=""><a
                            href="/news?year=2014&month=6&day=14">14</a>
                    </td>
                                                                                        <td class=""><a
                            href="/news?year=2014&month=6&day=15">15</a>
                    </td>
                                        </tr>
            <tr>
                                                                    <td class=""><a
                            href="/news?year=2014&month=6&day=16">16</a>
                    </td>
                                                                                        <td class=""><a
                            href="/news?year=2014&month=6&day=17">17</a>
                    </td>
                                                                                        <td class=""><a
                            href="/news?year=2014&month=6&day=18">18</a>
                    </td>
                                                                                        <td class=""><a
                            href="/news?year=2014&month=6&day=19">19</a>
                    </td>
                                                                                        <td class=""><a
                            href="/news?year=2014&month=6&day=20">20</a>
                    </td>
                                                                                        <td class=""><a
                            href="/news?year=2014&month=6&day=21">21</a>
                    </td>
                                                                                        <td class=""><a
                            href="/news?year=2014&month=6&day=22">22</a>
                    </td>
                                        </tr>
            <tr>
                                                                    <td class=""><a
                            href="/news?year=2014&month=6&day=23">23</a>
                    </td>
                                                                                        <td class=""><a
                            href="/news?year=2014&month=6&day=24">24</a>
                    </td>
                                                                                        <td class=""><a
                            href="/news?year=2014&month=6&day=25">25</a>
                    </td>
                                                                                        <td class=""><a
                            href="/news?year=2014&month=6&day=26">26</a>
                    </td>
                                                                                        <td class=""><a
                            href="/news?year=2014&month=6&day=27">27</a>
                    </td>
                                                                                        <td class=""><a
                            href="/news?year=2014&month=6&day=28">28</a>
                    </td>
                                                                                        <td class=""><a
                            href="/news?year=2014&month=6&day=29">29</a>
                    </td>
                                        </tr>
            <tr>
                                                                    <td class=""><a
                            href="/news?year=2014&month=6&day=30">30</a>
                    </td>
                                                                    <td></td>
                                                                    <td></td>
                                                                    <td></td>
                                                                    <td></td>
                                                                    <td></td>
                                                                    <td></td>
                                        </tr>
        </tbody>
</table>        </div>
    </div>
</div>

<script type="text/javascript">
    $(function () {
        $('.calendar').find('input').click(function () {
            var year = $('#year').text();
            var month = $('#month').find('label[class=active]').children().val();

            var url = '/news/get_days/?year=' + year + '&month=' + month;
            var calendar = $('#calendar-block');
            calendar.fadeTo('fast', 0.5);

            $.get(url, function(html){
                calendar.find('.bn-calendar').html(html);
                calendar.fadeTo('fast', 1);
            }, 'html');
        });
    });

</script>
    <div class="bn-block type-6 margin2 rightmenu">
        <ul class="bb-holder no-style">
                                        <li >
                    <a href="/news/rubric/986/">В центре внимания</a>

                                                        </li>
                            <li >
                    <a href="/news/rubric/1636/">Навстречу 70-летию Великой Победы</a>

                                                        </li>
                            <li >
                    <a href="/news/rubric/1574/">Имиджевые мероприятия УМВД </a>

                                                        </li>
                            <li >
                    <a href="/news/rubric/1816/">"Ангелы в погонах"</a>

                                                        </li>
                            <li >
                    <a href="/news/rubric/987/">Календарь мужества</a>

                                                        </li>
                            <li >
                    <a href="/news/rubric/988/">Информация для СМИ</a>

                                                        </li>
                            <li >
                    <a href="/news/rubric/989/">Безопасность на дороге</a>

                                                        </li>
                            <li >
                    <a href="/news/rubric/991/">Полицейская жизнь</a>

                                                        </li>
                            <li >
                    <a href="/news/rubric/1033/">Из кабинета следователя</a>

                                                        </li>
                            <li >
                    <a href="/news/rubric/1453/">Правоохранительно-патриотическая смена "Патриот"</a>

                                                        </li>
                        </li>
        </ul>
    </div>





    <div class="bn-block type-2 margin2">
        <div class="bb-title"><h3 class="h3">ДОРОГИ БЕЗ ПРОБЛЕМ</h3></div>
        <div class="bb-holder clearfix">
                            <div class="bb-item-image" style="float: none;">
                    <a  href="http://www.dorogibezproblem.ru/interactive-map/">
                        <img src="/upload/site68/Hiwoa30pvM-250x0.jpg" alt="" />
                    </a>
                </div>
                        <div class="bb-item-holder" style="margin: 10px 0 0 0;">
                <div class="bb-item-title"><a  href="http://www.dorogibezproblem.ru/interactive-map/"></a>
                                                        </div>
            </div>
        </div>
    </div>


    <div class="bn-block type-2 margin2">
        <div class="bb-title"><h3 class="h3">Телефон доверия </h3></div>
        <div class="bb-holder clearfix">
                        <div class="bb-item-holder" style="margin: 10px 0 0 0;">
                <div class="bb-item-title"><a target="_blank" href="http://67.mvd.ru/Kontakti/contact"><p>Телефон доверия УМВД России по Смоленской области</p>
<p>8 4812 38 05 35</p>
<p>Ваш звонок очень важен для нас</p>
<p></p></a>
                                                        </div>
            </div>
        </div>
    </div>


    <div class="bn-block type-2 margin2">
        <div class="bb-title"><h3 class="h3">Фотоконкурс Открытый взгляд</h3></div>
        <div class="bb-holder clearfix">
                            <div class="bb-item-image" style="float: none;">
                    <a target="_blank" href="http://foto-mvd.ru/">
                        <img src="/upload/site68/MFqc0KjZGk-250x0.png" alt="" />
                    </a>
                </div>
                        <div class="bb-item-holder" style="margin: 10px 0 0 0;">
                <div class="bb-item-title"><a target="_blank" href="http://foto-mvd.ru/"></a>
                                                        </div>
            </div>
        </div>
    </div>


    <div class="bn-block type-2 margin2">
        <div class="bb-title"><h3 class="h3">Доброе слово</h3></div>
        <div class="bb-holder clearfix">
                            <div class="bb-item-image" style="float: none;">
                    <a target="_blank" href="https://67.mvd.ru/Press_sluzhba/Nashi_proekti/Dobrie_slova_ot_nashih_grazhdan">
                        <img src="/upload/site68/a0NS6DIyFm-250x0.jpg" alt="" />
                    </a>
                </div>
                        <div class="bb-item-holder" style="margin: 10px 0 0 0;">
                <div class="bb-item-title"><a target="_blank" href="https://67.mvd.ru/Press_sluzhba/Nashi_proekti/Dobrie_slova_ot_nashih_grazhdan"></a>
                                                        </div>
            </div>
        </div>
    </div>


    <div class="bn-block type-2 margin2">
        <div class="bb-title"><h3 class="h3">Мемориал "Солдатам правопорядка, погибшим при исполнении служебного долга"</h3></div>
        <div class="bb-holder clearfix">
                            <div class="bb-item-image" style="float: none;">
                    <a target="_blank" href="https://67.mvd.ru/Press_sluzhba/Nashi_proekti/Memorial_Soldatam_pravoporjadka_pogibshi/Otkritie_memoriala_Soldatam_pravoporjadk">
                        <img src="/upload/site68/K2CTIjMIq8-250x0.jpg" alt="" />
                    </a>
                </div>
                        <div class="bb-item-holder" style="margin: 10px 0 0 0;">
                <div class="bb-item-title"><a target="_blank" href="https://67.mvd.ru/Press_sluzhba/Nashi_proekti/Memorial_Soldatam_pravoporjadka_pogibshi/Otkritie_memoriala_Soldatam_pravoporjadk"></a>
                                                        </div>
            </div>
        </div>
    </div>


    <div class="bn-block type-2 margin2">
        <div class="bb-title"><h3 class="h3">СПАЙСАМ.НЕТ</h3></div>
        <div class="bb-holder clearfix">
                            <div class="bb-item-image" style="float: none;">
                    <a target="_blank" href="https://67.mvd.ru/Press_sluzhba/Nashi_proekti/SPAJSAM.NET">
                        <img src="/upload/site68/3OcqD0mli1-250x0.jpg" alt="" />
                    </a>
                </div>
                        <div class="bb-item-holder" style="margin: 10px 0 0 0;">
                <div class="bb-item-title"><a target="_blank" href="https://67.mvd.ru/Press_sluzhba/Nashi_proekti/SPAJSAM.NET"></a>
                                                        </div>
            </div>
        </div>
    </div>


    <div class="bn-block type-2 margin2">
        <div class="bb-title"><h3 class="h3">Противодействие коррупции</h3></div>
        <div class="bb-holder clearfix">
                            <div class="bb-item-image" style="float: none;">
                    <a target="_blank" href="https://67.mvd.ru/folder/923374">
                        <img src="/upload/site68/vvwHHHsAvk-250x0.jpg" alt="" />
                    </a>
                </div>
                        <div class="bb-item-holder" style="margin: 10px 0 0 0;">
                <div class="bb-item-title"><a target="_blank" href="https://67.mvd.ru/folder/923374"></a>
                                                        </div>
            </div>
        </div>
    </div>


    <div class="bn-block type-2 margin2">
        <div class="bb-title"><h3 class="h3">Мы вместе за безопасность дорожного движения</h3></div>
        <div class="bb-holder clearfix">
                            <div class="bb-item-image" style="float: none;">
                    <a target="_blank" href="https://67.mvd.ru/folder/3195657">
                        <img src="/upload/site68/dvsdVe5OgY-250x0.jpg" alt="" />
                    </a>
                </div>
                        <div class="bb-item-holder" style="margin: 10px 0 0 0;">
                <div class="bb-item-title"><a target="_blank" href="https://67.mvd.ru/folder/3195657"></a>
                                                        </div>
            </div>
        </div>
    </div>



<div class="bn-left-menu">
                    <a target="_blank" class="first" href="http://mvd.ru/apps"><span>Мобильное приложение МВД России</span></a>
                            <a target="_blank" class="" href="http://67.mvd.ru/gumvd/rukovodstvo"><span>Руководство</span></a>
                            <a target="_blank" class="" href="http://67.mvd.ru/gumvd/structure"><span>Структура</span></a>
                            <a target="_blank" class="" href="http://67.mvd.ru/Kontakti/units"><span>Территориальные подразделения</span></a>
                            <a target="_blank" class="last" href="http://67.mvd.ru/citizens/graph"><span>Графики приема граждан</span></a>
            </div>
<div class="bn-radio margin2">
    <table>
        <tbody>
        <tr>
            <td>
                <img src="/media/mvd-2014/img/img7.jpg" alt=""/>
            </td>
            <td>
                <a class="ico i_rd" onclick="popupWin = window.open(this.href, 'win', 'status=no,toolbar=no,scrollbars=no,titlebar=no,menubar=no,resizable=yes,width=287,height=181,directories=no,location=no');popupWin.focus(); return false;" target="_blank" href="https://67.mvd.ru/radio.html">Слушать</a>
            </td>
        </tr>
        </tbody>
    </table>
</div>
<!--TODO new radio-->
<!--<div id="block-1" class="bn-modern_radio margin2">-->
<!--    <div id="radio_pleyer-poster" class="bnr-poster"></div>-->
<!--    <div class="bnr-stat">-->
<!--        <audio id="radio_pleyer" src="http://stream06.media.rambler.ru/mv128.mp3?32" type="audio/mp3" controls="none"></audio>-->
<!--    </div>-->
<!--</div>-->
<!--<script type="text/javascript">-->
<!--    var poster = $('#radio_pleyer-poster');-->
<!--    $('#radio_pleyer').mediaelementplayer({-->
<!--        audioWidth: '100%',-->
<!--        type: 'audio/mpeg',-->
<!--        loop: true,-->
<!--        features: ['playpause','current','volume'],-->
<!--        enablePluginDebug: true,-->
<!--        alwaysShowControls: true,-->
<!--        alwaysShowHours: false-->
<!--    }).bind('play', function () {-->
<!--        poster.hide();-->
<!--    }).bind('pause', function () {-->
<!--        poster.show();-->
<!--    });-->
<!--</script>-->

    
<div id="poll_place"></div>
<script type="text/javascript">
   $.ajax({
        data: '',
        type:'get',
        url:'/poll',
        success:function (data) {
            $('#poll_place').replaceWith(data);
        }
   });
</script>                </div>
                        <div class="clean"></div>
        </div>
        <div class="bn-federal-site wrapper">
        <div class="bs-title">Ссылки на сайты органов государственной власти:</div>
    <div class="bs-holder2">
            <ul class="bs-holder no-style links">
                                                <li class="ico i_nw">
                        <a href="https://67.mvd.ru/banners/redirect?bid=8059" target="_blank">Сайт Президента России</a>
                    </li>
                                                                <li class="ico i_nw">
                        <a href="https://67.mvd.ru/banners/redirect?bid=7226" target="_blank">Сервер органов государственной власти</a>
                    </li>
                                                                <li class="ico i_nw">
                        <a href="https://67.mvd.ru/banners/redirect?bid=7269" target="_blank">Общественный совет при УМВД России по Смоленской области</a>
                    </li>
                                    </ul>
        </div>
            <div class="m-b2">
        <div id="button-1" class="bs-title margin1"><a class="grey-text half-link ico i_arr" href="/">Полезные ресурсы</a></div>
        <div class="bs-holder2 advanced" style="display: none;">
                            <ul class="bs-holder no-style links">
                                                                        <li class="ico i_nw">
                                <a href="https://67.mvd.ru/banners/redirect?bid=6873"
                                   target="_blank">Закон о полиции</a>
                            </li>
                                                                                                <li class="ico i_nw">
                                <a href="https://67.mvd.ru/banners/redirect?bid=6874"
                                   target="_blank">Госзакупки</a>
                            </li>
                                                                                                <li class="ico i_nw">
                                <a href="https://67.mvd.ru/banners/redirect?bid=7200"
                                   target="_blank">Управление вневедомственной охраны УМВД России по Смоленской области</a>
                            </li>
                                                                                                <li class="ico i_nw">
                                <a href="https://67.mvd.ru/banners/redirect?bid=7201"
                                   target="_blank">УФМС России по Смоленской области</a>
                            </li>
                                                                                                <li class="ico i_nw">
                                <a href="https://67.mvd.ru/banners/redirect?bid=7227"
                                   target="_blank">УГИБДД УМВД России по Смоленской области</a>
                            </li>
                                                            </ul>
                            <ul class="bs-holder no-style links">
                                                                        <li class="ico i_nw">
                                <a href="https://67.mvd.ru/banners/redirect?bid=7228"
                                   target="_blank">Поисковый отряд Сальвар</a>
                            </li>
                                                                                                <li class="ico i_nw">
                                <a href="https://67.mvd.ru/banners/redirect?bid=7242"
                                   target="_blank">Учебные заведения системы МВД России </a>
                            </li>
                                                                                                <li class="ico i_nw">
                                <a href="https://67.mvd.ru/banners/redirect?bid=7243"
                                   target="_blank">История МВД России</a>
                            </li>
                                                                                                <li class="ico i_nw">
                                <a href="https://67.mvd.ru/banners/redirect?bid=7270"
                                   target="_blank">Портал правовой информации</a>
                            </li>
                                                                                                <li class="ico i_nw">
                                <a href="https://67.mvd.ru/banners/redirect?bid=7418"
                                   target="_blank">Полиция рекомендует</a>
                            </li>
                                                            </ul>
                            <ul class="bs-holder no-style links">
                                                                        <li class="ico i_nw">
                                <a href="https://67.mvd.ru/banners/redirect?bid=7993"
                                   target="_blank">Санаторно-курортное обеспечение</a>
                            </li>
                                                                                                <li class="ico i_nw">
                                <a href="https://67.mvd.ru/banners/redirect?bid=7994"
                                   target="_blank">Информация для пенсионеров ОВД</a>
                            </li>
                                                                                                <li class="ico i_nw">
                                <a href="https://67.mvd.ru/banners/redirect?bid=8033"
                                   target="_blank">Лицензионно-разрешительная работа</a>
                            </li>
                                                                                                <li class="ico i_nw">
                                <a href="https://67.mvd.ru/banners/redirect?bid=8034"
                                   target="_blank">Информационный центр УМВД</a>
                            </li>
                                                                                                <li class="ico i_nw">
                                <a href="https://67.mvd.ru/banners/redirect?bid=8046"
                                   target="_blank">Правовое информирование</a>
                            </li>
                                                            </ul>
                            <ul class="bs-holder no-style links">
                                                                        <li class="ico i_nw">
                                <a href="https://67.mvd.ru/banners/redirect?bid=8066"
                                   target="_blank">Электронное правительство Госуслуги</a>
                            </li>
                                                                                                <li class="ico i_nw">
                                <a href="https://67.mvd.ru/banners/redirect?bid=5420"
                                   target="_blank">Закон о полиции</a>
                            </li>
                                                                                                <li class="ico i_nw">
                                <a href="https://67.mvd.ru/banners/redirect?bid=8048"
                                   target="_blank">Спас-экстрим. Портал детской безопасности.</a>
                            </li>
                                                                                                <li class="ico i_nw">
                                <a href="https://67.mvd.ru/banners/redirect?bid=8186"
                                   target="_blank">Народные дружины</a>
                            </li>
                                                                                                <li class="ico i_nw">
                                <a href="https://67.mvd.ru/banners/redirect?bid=8287"
                                   target="_blank">Совет ветеранов МВД</a>
                            </li>
                                                            </ul>
                            <ul class="bs-holder no-style links">
                                                                        <li class="ico i_nw">
                                <a href="https://67.mvd.ru/banners/redirect?bid=8920"
                                   target="_blank">Фонд поддержки детей</a>
                            </li>
                                                            </ul>
                    </div>
    </div>
    </div>
<script type="text/javascript">
    $(function () {
        $('#button-1').on('click', function () {
            var n = $(this).next();
            if (n.is(':visible')) {
                n.slideUp(300);
            } else {
                n.slideDown(300);
            }
            return false;
        });
    });
</script>    </div>
    <div class="ln-footer">
    <div class="bn-sitecopy margin2">
        <div class="counters no_border">
            <table><tbody><tr>
                    <td><!-- Yandex.Metrika informer -->
<a href="http://metrika.yandex.ru/stat/?id=13929682&amp;from=informer"
target="_blank" rel="nofollow"><img src="//bs.yandex.ru/informer/13929682/3_1_FFFFFFFF_EFEFEFFF_0_pageviews"
style="width:88px; height:31px; border:0;" alt="Яндекс.Метрика" title="Яндекс.Метрика: данные за сегодня (просмотры, визиты и уникальные посетители)" onclick="try{Ya.Metrika.informer({i:this,id:13929682,type:0,lang:'ru'});return false}catch(e){}"/></a>
<!-- /Yandex.Metrika informer -->

<!-- Yandex.Metrika counter -->
<script type="text/javascript">
(function (d, w, c) {
    (w[c] = w[c] || []).push(function() {
        try {
            w.yaCounter13929682 = new Ya.Metrika({id:13929682,
                    clickmap:true,
                    trackLinks:true,
                    accurateTrackBounce:true});
        } catch(e) { }
    });

    var n = d.getElementsByTagName("script")[0],
        s = d.createElement("script"),
        f = function () { n.parentNode.insertBefore(s, n); };
    s.type = "text/javascript";
    s.async = true;
    s.src = (d.location.protocol == "https:" ? "https:" : "http:") + "//mc.yandex.ru/metrika/watch.js";

    if (w.opera == "[object Opera]") {
        d.addEventListener("DOMContentLoaded", f, false);
    } else { f(); }
})(document, window, "yandex_metrika_callbacks");
</script>
<noscript><div><img src="//mc.yandex.ru/watch/13929682" style="position:absolute; left:-9999px;" alt="" /></div></noscript>
<!-- /Yandex.Metrika counter --></td>
                </tr></tbody></table>
        </div>
        <br />© 2015, Управление МВД России по Смоленской области    </div>
    <div class="bn-footer-links margin2">
        <div class="f-left">
            <a class="half-link red-text" href="/sitemap/">Карта сайта</a>
            <a class="link copyright" href="#copyright">Об использовании информации сайта</a>
                            <a class="grey-text" href="mailto:pressuvdsml@yandex.ru?subject=Ошибка%20на%20сайте">Нашли ошибку на сайте?</a>
                    </div>
    </div>
</div>

<div class="mvd_copyright">
    <div class="padding">
        <h3>Об использовании информации сайта</h3>
        <p>Все материалы сайта Министерства внутренних дел Российской Федерации могут быть воспроизведены в любых
            средствах массовой информации, на серверах сети Интернет или на любых иных носителях без каких-либо
            ограничений по объему и срокам публикации.</p>
        <p>Это разрешение в равной степени распространяется на газеты, журналы, радиостанции, телеканалы, сайты и
            страницы сети Интернет. Единственным условием перепечатки и ретрансляции является ссылка на
            первоисточник.</p>
        <p>Никакого предварительного согласия на перепечатку со стороны Министерства внутренних дел Российской Федерации            не требуется.</p>
    </div>
</div>

<div id="window-overlay" class="bn-popup2 shadow">
    <a class="bp-close-but" href="#">close</a>
    <div class="content">

    </div>
</div></div>
</body>
</html>`
