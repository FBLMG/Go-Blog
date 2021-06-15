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
                        <!-- 账户ID -->
                        <input type="hidden" class="form-control" id="id" name="id"
                               placeholder="账户ID" value="{{.adminUser.Id}}">
                        <!-- 昵称 -->
                        <div class="form-group">
                            <label for="nickname" class="col-sm-2 control-label">昵称</label>
                            <div class="col-sm-10">
                                <input type="text" class="form-control" id="nickname" name="nickname"
                                       placeholder="昵称" value="{{.adminUser.Nickname}}">
                            </div>
                        </div>
                        <!-- 账户名 -->
                        <div class="form-group">
                            <label for="username" class="col-sm-2 control-label">用户名</label>
                            <div class="col-sm-10">
                                <input type="text" class="form-control" id="username" name="username"
                                       placeholder="用户名" value="{{.adminUser.Username}}">
                            </div>
                        </div>
                        <!-- 密码 -->
                        <div class="form-group">
                            <label for="password" class="col-sm-2 control-label">密码</label>
                            <div class="col-sm-10">
                                <input type="password" class="form-control" id="password" name="password"
                                       placeholder="密码" value="{{.adminUser.Password}}">
                            </div>
                        </div>
                        <!-- 状态 -->
                        <div class="form-group">
                            <label for="status" class="col-sm-2 control-label">状态</label>
                            <div class="col-sm-10">
                                <select id="status" name="status">
                                {{$statusSelect:=.adminUser.Status}}
                                {{range $k,$v:=.StatusList}}
                                    <option value="{{$k}}"
                                    {{if eq $statusSelect $k}}
                                            selected="selected"
                                    {{end}}
                                    >{{$v}}</option>
                                {{end}}
                                </select>
                            </div>
                        </div>
                        <!-- 编辑账户 -->
                        <div class="form-group">
                            <div class="col-sm-offset-2 col-sm-10">
                                <button type="button" class="btn btn-success" onclick="edit()">编辑</button>
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
    //编辑用户
    function edit() {
        var id = $('#id').val();
        var nickname = $('#nickname').val();
        var username = $('#username').val();
        var password = $('#password').val();
        var status = $('#status').val();
        var ajaxActionUrl = '/adminUserUpdate';
        if (id === '') {
            layer.msg('账户ID不允许为空！');
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
            data: {nickname: nickname, id: id, username: username, status: status, password: password},
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