<div class="side-body">
    <!--
    <div class="page-title">
        <span class="title">Form UI Kits</span>
        <div class="description">A ui elements use in form, input, select, etc.</div>
    </div>
    -->
    <div class="row">
        <div class="col-xs-12">
            <div class="card">
                <div class="card-header">
                    <div class="card-title">
                        <div class="title">账户</div>
                    </div>
                </div>
                <div class="card-body">
                    <form class="form-horizontal">
                        <!-- 昵称 -->
                        <div class="form-group">
                            <label for="nickname" class="col-sm-2 control-label">昵称</label>
                            <div class="col-sm-10">
                                <input type="text" class="form-control" id="nickname" name="nickname"
                                       placeholder="昵称">
                            </div>
                        </div>
                        <!-- 账户名 -->
                        <div class="form-group">
                            <label for="username" class="col-sm-2 control-label">用户名</label>
                            <div class="col-sm-10">
                                <input type="text" class="form-control" id="username" name="username"
                                       placeholder="用户名">
                            </div>
                        </div>
                        <!-- 密码 -->
                        <div class="form-group">
                            <label for="password" class="col-sm-2 control-label">密码</label>
                            <div class="col-sm-10">
                                <input type="text" class="form-control" id="password" name="password"
                                       placeholder="密码">
                            </div>
                        </div>
                        <!-- 状态 -->
                        <div class="form-group">
                            <label for="status" class="col-sm-2 control-label">状态</label>
                            <div class="col-sm-10">
                                <select id="status" name="status">
                                {{range $k,$v:=.StatusList}}
                                    <option value="{{$k}}">{{$v}}</option>
                                {{end}}
                                </select>
                            </div>
                        </div>
                        <!-- 添加账户 -->
                        <div class="form-group">
                            <div class="col-sm-offset-2 col-sm-10">
                                <button type="button" class="btn btn-success" onclick="add()">添加</button>
                            </div>
                        </div>
                    </form>
                </div>
            </div>
        </div>
    </div>
</div>
<!--  弹窗样式 -->
<link rel="stylesheet" href="static/thirdParty/layer/skin/layer.css"/>
<!--  弹窗样式 -->
<!-- 提交答案 -->
<script type="application/javascript">
    //添加用户
    function add() {
        var nickname = $('#nickname').val();
        var username = $('#username').val();
        var password = $('#password').val();
        var status = $('#status').val();
        var ajaxActionUrl = '/adminUserInsert';
        if (nickname === '') {
            layer.msg('昵称不允许为空！');
            return false;
        }
        if (username === '') {
            layer.msg('账户名不允许为空！');
            return false;
        }
        if (password === '') {
            layer.msg('密码不允许为空！');
            return false;
        }
        if (status === '') {
            layer.msg('状态不允许为空！');
            return false;
        }
        /* ajax提交 */
        $.ajax({
            //提交数据的类型 POST GET
            type: "POST",
            //提交的网址
            url: ajaxActionUrl,
            //提交的数据
            data: {nickname: nickname, username: username, status: status, password: password},
            //返回数据的格式
            datatype: "json",//"xml", "html", "script", "json", "jsonp", "text".
            //成功返回之后调用的函数
            success: function (data) {
                ////根据ajax返回参数判断添加情况
                if (data.status === 1) {
                    layer.msg(data.message);
                    window.location.reload();
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