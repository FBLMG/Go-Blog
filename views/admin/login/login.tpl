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
    <link rel="shortcut icon" sizes="200x200" href="static/thirdParty/icon/icon.jpg">
</head>

<body class="flat-blue login-page" style="background:url('static/admin/img/app-header-bg.jpg')">
<div class="container">
    <div class="login-box">
        <div>
            <div class="login-form row">
                <div class="col-sm-12 text-center login-header">
                    <i class="login-logo fa fa-connectdevelop fa-5x"></i>
                    <h4 class="login-title"></h4>
                </div>
                <div class="col-sm-12">
                    <div class="login-body">
                        <div class="progress hidden" id="login-progress">
                            <div class="progress-bar progress-bar-success progress-bar-striped active"
                                 role="progressbar" aria-valuenow="100" aria-valuemin="0" aria-valuemax="100"
                                 style="width: 100%">
                                登陆
                            </div>
                        </div>
                        <form>
                            <div class="control">
                                <input type="text" class="form-control" id="username" name="username"
                                       placeholder="用户账号"/>
                            </div>
                            <div class="control">
                                <input type="password" class="form-control" id="password" name="password"
                                       placeholder="用户密码"/>
                            </div>
                            <div class="login-button text-center">
                                <input type="button" class="btn btn-primary" value="登陆" onclick="doLogin()">
                            </div>
                        </form>
                    </div>
                    <div class="login-footer">
                        <!--  <span class="text-right"><a href="#" class="color-white">忘记  密码?</a></span> -->
                    </div>
                </div>
            </div>
        </div>
    </div>
</div>
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
</body>
<!-- 提交答案 -->
<script type="application/javascript">
    //处理用户登陆
    function doLogin() {
        var username = $('#username').val();
        var password = $('#password').val();
        var ajaxActionUrl = '/adminDoLogin';
        if (username === '') {
            layer.msg('用户名不允许为空！');
            return false;
        }
        if (password === "") {
            layer.msg('密码不允许为空！');
            return false;
        }
        /* ajax提交 */
        $.ajax({
            //提交数据的类型 POST GET
            type: "POST",
            //提交的网址
            url: ajaxActionUrl,
            //提交的数据
            data: {username: username, password: password},
            //返回数据的格式
            datatype: "json",//"xml", "html", "script", "json", "jsonp", "text".
            //成功返回之后调用的函数
            success: function (data) {
                ////根据ajax返回参数判断添加情况
                if (data.status === 1) {
                    layer.msg(data.message);
                    //跳转至首页
                    window.location.href = "/admin";
                } else if (data.status === -1) {
                    layer.msg(data.message);
                } else {
                    layer.msg('未知错误!');
                }
                ////根据ajax返回参数判断添加情况
            }
        });
        /* ajax提交 */
    }
</script>
<!-- 用户登陆 -->
</html>
