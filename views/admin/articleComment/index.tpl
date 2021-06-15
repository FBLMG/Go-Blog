<div class="side-body">
    <!--
    <div class="page-title">
        <span class="title">Table</span>
        <div class="description">A bootstrap table for display list of data.</div>
    </div>
    -->
    <div class="row">
        <div class="col-xs-12">
            <div class="card">
                <div class="card-header">
                    <div class="card-title">
                        <form class="form-inline">
                            <!--文章查询-->
                            <div class="form-group">
                                <label for="searchArticleId">文章</label>
                                <select id="searchArticleId" name="searchArticleId" title="请选择文章">
                                {{$searchArticleIdSelect:=.getInput.SearchArticleId}}
                                {{range $k,$v:=.ArticleList}}
                                    <option value="{{$v.Id}}"
                                    {{if eq $searchArticleIdSelect $v.Id}}
                                            selected="selected"
                                    {{end}}>
                                    {{$v.Title}}
                                    </option>
                                {{end}}
                                </select>
                            </div>
                            <!--状态查询-->
                            <div class="form-group">
                                <label for="searchStatus">状态</label>
                                <select id="searchStatus" name="searchStatus" title="请选择状态">
                                {{$statusSelect:=.getInput.SearchStatus}}
                                {{range $k,$v:=.StatusList}}
                                    <option value="{{$k}}"
                                    {{if eq $statusSelect $k}}
                                            selected="selected"
                                    {{end}}>
                                    {{$v}}
                                    </option>
                                {{end}}
                                </select>
                            </div>
                            <!--是否管理员查询-->
                            <div class="form-group">
                                <label for="searchAdminStatus">管理员</label>
                                <select id="searchAdminStatus" name="searchAdminStatus" title="请选择管理员状态">
                                {{$searchAdminStatusSelect:=.getInput.SearchAdminStatus}}
                                {{range $k,$v:=.AdminStatus}}
                                    <option value="{{$k}}"
                                    {{if eq $searchAdminStatusSelect $k}}
                                            selected="selected"
                                    {{end}}>
                                    {{$v}}
                                    </option>
                                {{end}}
                                </select>
                            </div>
                            <!--操作按钮-->
                            <button type="submit" class="btn btn-info" style="border-radius: 5px;">搜索</button>
                            <a href="/articleCommentGetList" type="button" class="btn btn-danger"
                               style="border-radius: 5px;">撤回</a>
                        </form>
                    </div>
                </div>
                <div class="card-body">
                    <div class="panel panel-default">
                        <!-- Table -->
                        <table class="table">
                            <thead>
                            <tr>
                                <th>文章标题</th>
                                <th>用户昵称</th>
                                <th>用户邮箱</th>
                                <th>评论内容</th>
                                <th>状态</th>
                                <th>创建时间</th>
                                <th>操作</th>
                            </tr>
                            </thead>
                            <tbody>
                            {{range .ArticleCommentList}}
                            <tr>
                                <td>{{.ArticleTitle}}</td>
                                <td>{{.UserNickname}}</td>
                                <td>{{.UserEmail}}</td>
                                <td>{{.Content}}</td>
                                <td>
                                {{if eq .Status 1}}
                                    <label class="success">正常</label>
                                {{else if eq .Status 2}}
                                    <label class="warning">待审核</label>
                                {{else}}
                                    <label class="danger">删除</label>
                                {{end}}
                                </td>
                                <td>{{.CreateAt}}</td>
                                <td>
                                    <button type="button" class="btn btn-xs btn-info" data-toggle="modal"
                                            data-target=".bs-example-modal-sm"
                                            onclick="sendCommentDiv({{.Id}},{{.ArticleId}})">
                                        回复
                                    </button>
                                {{if eq .Status 1}}
                                    <button type="button" class="btn btn-xs btn-warning"
                                            onclick="updateStatusToClose({{.Id}})">
                                        不通过
                                    </button>
                                    <button type="button" class="btn btn-xs btn-danger"
                                            onclick="updateStatusToDelete({{.Id}})">
                                        删除
                                    </button>
                                {{else if eq .Status 2}}
                                    <button type="button" class="btn btn-xs btn-success"
                                            onclick="updateStatusToOpen({{.Id}})">
                                        通过
                                    </button>
                                    <button type="button" class="btn btn-xs btn-danger"
                                            onclick="updateStatusToDelete({{.Id}})">
                                        删除
                                    </button>
                                {{else}}
                                    <button type="button" class="btn btn-xs btn-success"
                                            onclick="updateStatusToOpen({{.Id}})">
                                        通过
                                    </button>
                                {{end}}
                                </td>
                            </tr>
                            {{end}}
                            </tbody>
                        </table>
                    </div>
                </div>
            </div>
        </div>
        <!-- 分页 -->
        <div class="layui-row" id="page" style="text-align: center;"></div>
        <!-- 分页 -->
    </div>
</div>
<!-- 模态框 -->
<div class="modal fade bs-example-modal-sm" tabindex="-1" role="dialog" aria-labelledby="mySmallModalLabel">
    <div class="modal-dialog modal-sm" role="document">
        <div class="modal-content">
            <div class="modal-header">
                <button type="button" class="close" data-dismiss="modal" aria-label="Close"><span
                        aria-hidden="true">×</span></button>
                <h4 class="modal-title" id="mySmallModalLabel">回复评论</h4>
            </div>
            <div class="modal-body">
                <form class="form-horizontal">
                    <input name="commentId" id="commentId" type="hidden" placeholder="评论ID">
                    <input name="articleId" id="articleId" type="hidden" placeholder="文章ID">
                    <div class="form-group" style="padding: 10px;">
                       <textarea class="form-control" rows="5" id="commentContent"
                                 name="commentContent"></textarea>
                    </div>
                    <div class="form-group">
                        <div class="col-sm-offset-2 col-sm-10" style="text-align: right">
                            <button type="button" class="btn btn-sm btn-success" onclick="articleCommentReply()">回复
                            </button>
                        </div>
                    </div>
                </form>
            </div>
        </div>
    </div>
</div>
<!-- 分页样式 -->
<link href="static/home/layui/css/layui.css" rel="stylesheet"/>
<!-- 分页样式 -->
<!-- 处理分页 -->
<script src="static/home/layui/layui.js?s=36"></script>
<script>
    layui.use('laypage', function () {
        var laypage = layui.laypage;
        var searchStatus = {{.getInput.SearchStatus}};
        var searchAdminStatus = {{.getInput.SearchAdminStatus}};
        var searchArticleId = {{.getInput.SearchArticleId}};
        //执行一个laypage实例
        laypage.render({
            elem: 'page' //注意，这里的 page 是 ID，不用加 # 号
            , count: {{.PageCount}}    //获取数据总条数
            , limit: {{.PageLimit}}    //每页限制多少条
            , curr:{{.PagePage}}       //当前页码
            , jump: function (obj, first) {
                //obj包含了当前分页的所有参数，比如：
                console.log(obj.curr); //得到当前页，以便向服务端请求对应页的数据。
                console.log(obj.limit); //得到每页显示的条数
                //首次不执行
                if (!first) {
                    window.location.href = "/articleCommentGetList?page=" + obj.curr + "&searchStatus=" + searchStatus
                            + "&searchAdminStatus=" + searchAdminStatus + "&searchArticleId=" + searchArticleId;
                }
            }
        });
    });
</script>
<!-- 处理分页 -->
<!-- 操作按钮 -->
<script type="application/javascript">
    //回复评论赋值
    function sendCommentDiv($commentId, $articleId) {
        var commentId = $commentId;
        var articleId = $articleId;
        $("#commentId").val(commentId);
        $("#articleId").val(articleId);
    }

    //回复评论
    function articleCommentReply() {
        //获取用户参数
        var commentId = $('#commentId').val();
        var articleId = $('#articleId').val();
        var commentContent = $('#commentContent').val();
        //判断用户参数
        if (commentId === 0) {
            layer.msg('评论Id获取失败!');
        }
        if (articleId === 0) {
            layer.msg('文章Id获取失败!');
        }
        if (commentContent === "") {
            layer.msg('请说点什么!');
        }
        //提交地址
        var ajaxActionUrl = '/articleCommentReply';
        /* ajax提交 */
        $.ajax({
            //提交数据的类型 POST GET
            type: "POST",
            //提交的网址
            url: ajaxActionUrl,
            //提交的数据
            data: {
                articleId: articleId,
                commentId: commentId,
                content: commentContent
            },
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
    }


    //审核通过
    function updateStatusToOpen(id) {
        var sureUpdateStatusToOpen = confirm("你确定审核通过该评论？");
        if (sureUpdateStatusToOpen === false) {
            alert("请确认后再重新提交");
            return false;
        }
        var ajaxActionUrl = '/articleCommentUpdateStatus';
        /* ajax提交 */
        $.ajax({
            //提交数据的类型 POST GET
            type: "POST",
            //提交的网址
            url: ajaxActionUrl,
            //提交的数据
            data: {id: id, status: 1},
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

    //审核不通过
    function updateStatusToClose(id) {
        var sureUpdateStatusToClose = confirm("你确定审核不通过该评论？");
        if (sureUpdateStatusToClose === false) {
            alert("请确认后再重新提交");
            return false;
        }
        var ajaxActionUrl = '/articleCommentUpdateStatus';
        /* ajax提交 */
        $.ajax({
            //提交数据的类型 POST GET
            type: "POST",
            //提交的网址
            url: ajaxActionUrl,
            //提交的数据
            data: {id: id, status: 2},
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

    //删除评论
    function updateStatusToDelete(id) {
        var sureDeleteData = confirm("你确定删除该评论");
        if (sureDeleteData === false) {
            alert("请确认后再重新提交");
            return false;
        }
        //
        var ajaxActionUrl = '/articleCommentUpdateStatus';
        /* ajax提交 */
        $.ajax({
            //提交数据的类型 POST GET
            type: "POST",
            //提交的网址
            url: ajaxActionUrl,
            //提交的数据
            data: {id: id, status: 3},
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
<!-- 操作按钮 -->