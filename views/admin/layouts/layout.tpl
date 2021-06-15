<!DOCTYPE html>
<html>
<head>
    <title>{{.WebTitle}}</title>
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <!-- CSS Libs -->
    <link rel="stylesheet" type="text/css" href="static/admin/lib/css/bootstrap.min.css">
    <link rel="stylesheet" type="text/css" href="static/admin/lib/css/font-awesome.min.css">
    <link rel="stylesheet" type="text/css" href="static/admin/lib/css/animate.min.css">
    <link rel="stylesheet" type="text/css" href="static/admin/lib/css/bootstrap-switch.min.css">
    <link rel="stylesheet" type="text/css" href="static/admin/lib/css/checkbox3.min.css">
    <link rel="stylesheet" type="text/css" href="static/admin/lib/css/jquery.dataTables.min.css">
    <link rel="stylesheet" type="text/css" href="static/admin/lib/css/dataTables.bootstrap.css">
    <link rel="stylesheet" type="text/css" href="static/admin/lib/css/select2.min.css">
    <!-- CSS App -->
    <link rel="stylesheet" type="text/css" href="static/admin/css/style.css">
    <link rel="stylesheet" type="text/css" href="static/admin/css/themes/flat-blue.css">
    <!-- 引入样式方法 -->
{{.ThirdPartyCss}}
    <!-- 引入样式方法 -->
    <!-- 导航图标 -->
    <link rel="shortcut icon" sizes="200x200" href="static/thirdParty/icon/icon.jpg">
    <!-- 导航图标 -->
</head>

<body class="flat-blue">
<div class="app-container">
    <div class="row content-container">
        <nav class="navbar navbar-default navbar-fixed-top navbar-top">
            <div class="container-fluid">
                <div class="navbar-header">
                    <button type="button" class="navbar-expand-toggle">
                        <i class="fa fa-bars icon"></i>
                    </button>
                    <ol class="breadcrumb navbar-breadcrumb">
                        <li class="active">{{.WebTitle}}</li>
                    </ol>
                    <button type="button" class="navbar-right-expand-toggle pull-right visible-xs">
                        <i class="fa fa-th icon"></i>
                    </button>
                </div>
                <ul class="nav navbar-nav navbar-right">
                    <button type="button" class="navbar-right-expand-toggle pull-right visible-xs">
                        <i class="fa fa-times icon"></i>
                    </button>

                    <li class="dropdown">
                        <a href="#" class="dropdown-toggle" data-toggle="dropdown" role="button"
                           aria-expanded="false"><i class="glyphicon glyphicon-home"></i></a>
                        <ul class="dropdown-menu animated fadeInDown">
                            <li class="title">
                                操作
                            </li>
                            <li class="message">
                                <a href="/" style="pointer-events: painted" target="_blank">前往前台</a>
                            </li>
                        </ul>
                    </li>
                    <li class="dropdown profile">
                        <a href="#" class="dropdown-toggle" data-toggle="dropdown" role="button"
                           aria-expanded="false">{{.LeftActive.adminNickname}}<span class="caret"></span></a>
                        <ul class="dropdown-menu animated fadeInDown">
                            <li class="profile-img">
                                <img src="static/admin/img/profile/picjumbo.com_HNCK4153_resize.jpg"
                                     class="profile-img">
                            </li>
                            <li>
                                <div class="profile-info">
                                    <h4 class="username">{{.LeftActive.adminNickname}}</h4>
                                    <div class="btn-group margin-bottom-2x" role="group">
                                        <a type="button" class="btn btn-default"
                                           href="/adminLoginOut"><i
                                                class="fa fa-sign-out"></i>
                                            退出
                                        </a>
                                    </div>
                                </div>
                            </li>
                        </ul>
                    </li>
                </ul>
            </div>
        </nav>
        <div class="side-menu sidebar-inverse">
            <nav class="navbar navbar-default" role="navigation">
                <div class="side-menu-container">
                    <div class="navbar-header">
                        <a class="navbar-brand" href="/admin">
                            <div class="icon fa fa-paper-plane"></div>
                            <div class="title">{{.LeftActive.adminNickname}}</div>
                        </a>
                        <button type="button" class="navbar-expand-toggle pull-right visible-xs">
                            <i class="fa fa-times icon"></i>
                        </button>
                    </div>
                    <!---***  左侧导航栏 ***--->
                    <ul class="nav navbar-nav {{.LeftActive.BlogActive}}">
                        <li class="{{.LeftActive.HomeActive}}">
                            <a href="/admin">
                                <span class="icon fa fa-tachometer"></span><span class="title">首页</span>
                            </a>
                        </li>
                        <li class="{{.LeftActive.BlogActive}}">
                            <a href="/adminArticleList">
                                <span class="icon glyphicon glyphicon-book"></span><span class="title">文章</span>
                            </a>
                            <!-- 文章下级菜单 -->
                            <div id="dropdown-table" class="panel-collapse collapse">
                                <div class="panel-body">
                                    <ul class="nav navbar-nav">
                                        <li><a href="/AdminAddBlog">发表文章</a>
                                        </li>
                                        <li><a href="/AdminBlogList">管理文章</a>
                                        </li>
                                    </ul>
                                </div>
                            </div>
                        </li>
                        <!-- 博客分类 -->
                        <li class="{{.LeftActive.ClassifyActive}}">
                            <a href="/adminArticleTypeList">
                                <span class="icon glyphicon glyphicon-list"></span><span class="title">博客分类</span>
                            </a>
                        </li>
                        <!-- 标签-->
                        <li class="{{.LeftActive.LabelActive}}">
                            <a href="/adminArticleLabelList">
                                <span class="icon glyphicon glyphicon-tag"></span><span class="title">博客标签</span>
                            </a>
                        </li>
                        <!-- 博客评论-->
                        <li class="{{.LeftActive.ArticleCommentActive}}">
                            <a href="/articleCommentGetList">
                                <span class="icon glyphicon glyphicon-eye-open"></span><span class="title">博客评论</span>
                            </a>
                        </li>
                        <!-- 友情链接-->
                        <li class="{{.LeftActive.FriendActive}}">
                            <a href="/adminFriendshipList">
                                <span class="icon glyphicon glyphicon-random"></span><span class="title">友情链接</span>
                            </a>
                        </li>
                        <!-- Dropdown-->
                        <li class="panel panel-default dropdown {{.LeftActive.SettingActive}}">
                            <a data-toggle="collapse" href="#dropdown-setting">
                                <span class="icon glyphicon glyphicon-cog"></span><span class="title">系统配置</span>
                            </a>
                            <!-- Dropdown level 1 -->
                            <div id="dropdown-setting" class="panel-collapse collapse">
                                <div class="panel-body">
                                    <ul class="nav navbar-nav">
                                        <li><a href="/adminSettingGetBottomConfig">底部配置</a>
                                        </li>
                                        <li><a href="/adminSettingGetUserConfig">自我介绍</a>
                                        </li>
                                        <li><a href="/adminSettingGetAdConfig">广告配置</a>
                                        </li>
                                        <li><a href="/adminSettingGetWebTitleConfig">网站标题</a>
                                        </li>
                                    </ul>
                                </div>
                            </div>
                        </li>
                        <!-- 账号管理-->
                        <li class="{{.LeftActive.AdminActive}}">
                            <a href="/adminUserList">
                                <span class="icon glyphicon glyphicon-user"></span><span class="title">账号管理</span>
                            </a>
                        </li>
                        <!-- 统计记录-->
                        <li class="panel panel-default dropdown {{.LeftActive.StatisticActive}}">
                            <a data-toggle="collapse" href="#dropdown-statistic">
                                <span class="icon glyphicon glyphicon-paperclip"></span><span class="title">统计记录</span>
                            </a>
                            <!-- Dropdown level 1 -->
                            <div id="dropdown-statistic" class="panel-collapse collapse">
                                <div class="panel-body">
                                    <ul class="nav navbar-nav">
                                        <li><a href="/adminStatisticUserLogin">用户登陆</a>
                                        </li>
                                        <li><a href="/adminStatisticArticle">文章访问</a>
                                        </li>
                                    </ul>
                                </div>
                            </div>
                        </li>
                        <!-- 退出 -->
                        <li>
                            <a href="/adminLoginOut">
                                <span class="icon glyphicon glyphicon-road"></span><span class="title">退出</span>
                            </a>
                        </li>
                    </ul>
                    <!---***  左侧导航栏 ***--->
                </div>
                <!-- /.navbar-collapse -->
            </nav>
        </div>
        <!-- Main Content -->
        <div class="container-fluid">
            <!---***  中间内容 ***--->
        {{.LayoutContent}}
            <!---***  中间内容 ***--->
        </div>
    </div>
    <footer class="app-footer">
        <div class="wrapper">
            <span class="pull-right">2.1 <a href="#"><i class="fa fa-long-arrow-up"></i></a></span> © 2020 Copyright.
        {{.LeftActive.bottomConfig.webTitle}}
            - Collect from <a href="{{.LeftActive.bottomConfig.recordUrl}}"
                              target="_blank">{{.LeftActive.bottomConfig.recordTitle}}</a>
        </div>
    </footer>
    <!-- Javascript Libs -->
    <script type="text/javascript" src="static/admin/lib/js/jquery.min.js"></script>
    <script type="text/javascript" src="static/admin/lib/js/bootstrap.min.js"></script>
    <!-- 弹窗 -->
    <script src="static/thirdParty/layer/layer.js"></script>
    <!-- 弹窗 -->
    <script type="text/javascript" src="static/admin/lib/js/Chart.min.js"></script>
    <script type="text/javascript" src="static/admin/lib/js/bootstrap-switch.min.js"></script>
    <script type="text/javascript" src="static/admin/lib/js/jquery.matchHeight-min.js"></script>
    <script type="text/javascript" src="static/admin/lib/js/jquery.dataTables.min.js"></script>
    <script type="text/javascript" src="static/admin/lib/js/dataTables.bootstrap.min.js"></script>
    <script type="text/javascript" src="static/admin/lib/js/select2.full.min.js"></script>
    <script type="text/javascript" src="static/admin/lib/js/ace/ace.js"></script>
    <script type="text/javascript" src="static/admin/lib/js/ace/mode-html.js"></script>
    <script type="text/javascript" src="static/admin/lib/js/ace/theme-github.js"></script>
    <!-- Javascript -->
    <script type="text/javascript" src="static/admin/js/app.js"></script>
{{/* <script type="text/javascript" src="static/admin/js/index.js"></script>*/}}
    <!-- 引入图片上传JS方法 -->
{{.ThirdPartyImageJs}}
    <!-- 引入图片上传JS方法 -->
</div>
</body>

</html>
<!-- 引入JS方法 -->
{{.ThirdPartyJs}}
<!-- 引入JS方法 -->