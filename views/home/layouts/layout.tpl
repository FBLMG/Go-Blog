<!DOCTYPE html>
<html lang="en" dir="ltr">
<head>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <meta name="author" content="">
    <meta name="description" content="">
    <meta name="keywords" content="">
    <!-- Title-->
    <title>{{.BaseData.webTitle}}</title>
    <!-- Stylesheets-->
    <link rel="stylesheet" href="static/home/css/bootstrap.min.css">
    <link rel="stylesheet" href="static/home/css/all.min.css">
    <link rel="stylesheet" href="static/home/css/style.css">
    <link rel="shortcut icon" sizes="200x200" href="static/thirdParty/icon/icon.jpg">
</head>
<body>
<!--***顶部导航栏***-->
<header class="themelazer_main_header">
    <div class="themelazer_middle_header white_bg">
        <div class="container clearfix">
            <div class="row">
                <div class="col-md-12">
                    <div class="themelazer_promomenu_wrapper">
                        <div class="themelazer_header_social_icons">
                        </div>
                        <div class="themelazer_mobile_logo ">
                            <a href="/"><img src="static/home/image/relovan.png" alt="" title=""></a>
                        </div>
                        <div class="themelazer-nav clearfix">
                            <!-- 文章分类 -->
                            <div class="themelazer-navigation">
                                <ul class="menu black_color">
                                {{range .BaseData.articleType}}
                                    <li><a href="/articleType?articleTypeId={{.Id}}">{{.Title}}</a></li>
                                {{end}}
                                </ul>
                            </div>
                            <!-- 文章分类 -->
                        </div>
                        <ul class="header-s-m black_color">
                        {{/* <li class="nav-search">
                                <a href="#header-search" title="Search"><i class="ti-search"></i></a>
                            </li>*/}}
                            <li class="themelazer_mb_themelazern sidemenuoption-open is-active"><i class="ti-menu"></i>
                            </li>
                        </ul>
                    </div>
                </div>
            </div>
        </div>
        <div id="header-search" class="header-search">
            <button type="button" class="close">
                <i class="ti-close"></i>
            </button>
            <form class="header-search-form">
                <input type="search" value="" placeholder="What are you looking for?">
                <button type="submit" class="search-btn">
                    <i class="ti-search"></i>
                </button>
            </form>
        </div>
    </div>
    <!--Header Top-->
    <div id="themelazer_top_bar">
        <div class="container">
            <div class="row">
                <div class="col-md-12">
                    <div class="themelazer_logo_header2">
                        <a href="/"><img src="static/home/image/relovan.png" alt="" title=""></a>
                    </div>
                </div>
            </div>
        </div>
    </div>
    <!-- End Header Top -->
</header>
<!--***顶部导航栏***-->
<!--***中间内容***-->
<div class="themelazer-blog-body {{.ContentCss}}">
    <div class="container">
        <div class="row">
        {{.LayoutContent}}
            <div class="col-md-4 themelazer_sidebar">
                <div class="sidebar">
                    <div class="themelazer-widget-author">
                        <div class="author-container">
                            <h5 class="themelazer-author-title"><a href="#"
                                                                   rel="author">{{.BaseData.userConfig.userTitle}}</a>
                            </h5>
                            <div class="themelazer-author-avatar">
                                <a href="#" rel="author"><img alt="" src="{{.BaseData.userConfig.userImage}}"
                                                              class="avatar avatar-80 photo"> </a>
                            </div>
                            <div class="themelazer-author-data">
                                <div class="author-description">{{.BaseData.userConfig.userDesc}}
                                </div>
                            {{/* <div class="themelazer-autograph-about">
                                    <img src="image/sign_relovan.png" alt="">
                                </div>
                                <div class="themelazer-author-social-links">
                                    <div class="themelazer-social-links-items">
                                        <div class="themelazer-social-links-item">
                                            <a href="#" class="themelazer-social-links-link themelazer-facebook"
                                               target="_blank"> <i class="fab fa-facebook-f"></i> </a>
                                        </div>
                                        <div class="themelazer-social-links-item">
                                            <a href="#" class="themelazer-social-links-link  themelazer-twitter"
                                               target="_blank"> <i class="fab fa-twitter"></i> </a>
                                        </div>
                                        <div class="themelazer-social-links-item">
                                            <a href="#" class="themelazer-social-links-link themelazer-youtube"
                                               target="_blank"> <i class="fab fa-youtube"></i> </a>
                                        </div>
                                    </div>
                                </div>*/}}
                            </div>
                        </div>
                    </div>
                    <div class="single-sidebar recent-post-widget">
                        <!--置顶文章-->
                        <div class="title">
                            <h3>置顶文章</h3>
                        </div>
                        <div class="recent-post-wrapper">
                        {{range .BaseData.signArticle}}
                            <div class="themelazer_article_list">
                                <div class="post-outer">
                                    <div class="post-inner">
                                        <div class="post-thumbnail sidebar">
                                            <img src="{{.Image}}" alt="image">
                                            <span class="themelazer_site_count">{{.k}}</span>
                                            <a href="/articleInfo?id={{.Id}}"></a>
                                        </div>
                                    </div>
                                    <div class="post-inner">
                                        <div class="entry-header">
                                            <div class="themelazer_post_categories">
                                                <a href="/articleInfo?id={{.Id}}">{{.ArticleTypeTitle}}</a>
                                            </div>
                                            <h2 class="entry-title"><a href="/articleInfo?id={{.Id}}">{{.Title}}</a>
                                            </h2>
                                            <div class="meta-info">
                                                <ul>
                                                    <li class="post-date">{{.CreateAt}}</li>
                                                {{/*<li class="post-view">89K Views</li>*/}}
                                                </ul>
                                            </div>
                                        </div>
                                    </div>
                                </div>
                            </div>
                        {{end}}
                        </div>
                        <!--置顶文章-->
                        <!--标签-->
                        <div class="title">
                            <h3>标签</h3>
                        </div>
                        <div class="themelazer_widget_categories">
                            <ul>
                            {{range .BaseData.articleLabel}}
                                <li><a href="/articleLabel?labelId={{.Id}}">{{.Title}}<span>hot</span></a></li>
                            {{end}}
                            </ul>
                        </div>
                        <!--标签-->
                        <!--广告-->
                    {{if  eq .BaseData.adConfig.adStatus 1 }}
                        <div class="title">
                            <h3>{{.BaseData.adConfig.adTitle}}</h3>
                        </div>
                        <div class="themelazer_banner_spot">
                            <div class="themelazer_content_banner">
                                <div class="themelazer_bg_image_banner">
                                    <a href="{{.BaseData.adConfig.adUrl}}" target="_blank">
                                        <img src="{{.BaseData.adConfig.adImage}}" alt="">
                                    </a>
                                </div>
                            </div>
                        </div>
                    {{end}}
                        <!--友情链接-->
                        <div class="title">
                            <h3>友链</h3>
                        </div>
                        <div class="themelazer_widget_content">
                            <div class="themelazer_tagcloud">
                            {{range .BaseData.friendship}}
                                <a target="_blank" href="{{.Url}}" class="tag-cloud-link">{{.Title}}</a>
                            {{end}}
                            </div>
                        </div>
                        <!--友情链接-->
                    </div>
                </div>
            </div>
        </div>
    </div>
</div>
<!--***中间内容***-->
<!--***底部***-->
<footer class="footer-area">
    <div class="copyright-area">
        <div class="container">
            <div class="row">
                <div class="col-lg-12">
                    <div class="copyright-area-inner">
                        © 2020 Copyright. {{.BaseData.bottomConfig.webTitle}} - <a
                            href="{{.BaseData.bottomConfig.recordUrl}}">{{.BaseData.bottomConfig.recordTitle}}</a> All
                        rights reserved.
                    </div>
                </div>
            </div>
        </div>
    </div>
</footer>
<aside class="sidemenuoption">
    <div class="sidemenuoption-inner">
        <span class="menuoption-close"><i class="ti-close"></i></span>
        <div class="site-name-logo">
            <div class="site-name">
                <a href="/"><img src="static/home/image/relovan.png" alt="" title=""></a>
            </div>
        </div>
        <div class="themelazer_mobile_menu"></div>
    </div>
</aside>
<div class="body-overlay"></div>
<div class="scroll-totop-wrapper">
    <button class="scroll-totop "><i class="ti-angle-up"></i></button>
</div>
<!--***底部***-->
<!--***JS样式***-->
<script src="static/home/js/jquery.js"></script>
<script src="static/thirdParty/layer/layer.js"></script>
<script src="static/home/js/jquery.sticky.js"></script>
<script src="static/home/js/theia-sticky-sidebar.js"></script>
<script src="static/home/js/fluidvids.js"></script>
<script src="static/home/js/justified.js"></script>
<script src="static/home/js/slick.js"></script>
<script src="static/home/js/masonry.pkgd.min.js"></script>
<script src="static/home/js/main.js"></script>

<!--***JS样式***-->
</body>
</html>