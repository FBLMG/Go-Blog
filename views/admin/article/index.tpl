<div class="side-body">
    <!--
    <div class="page-title">
        <span class="title">Datatable</span>
        <div class="description">with jquery Datatable for display data with most usage functional. such as search, ajax
            loading, pagination, etc.
        </div>
    </div>
    -->
    <div class="row">
        <div class="col-xs-12">
            <div class="card">
                <div class="card-header">
                    <div class="card-title">
                        <form class="form-inline">
                            <div class="form-group">
                                <label for="searchTitle">标题</label>
                                <input type="text" class="form-control" id="searchTitle" placeholder="标题"
                                       name="searchTitle" style="border-radius: 5px;" value="{{.getInput.searchTitle}}">
                            </div>
                            <div class="form-group">
                                <label for="searchTypeId">分类</label>
                                <select id="searchTypeId" name="searchTypeId" title="请选择分类">
                                    <option value="0">全部</option>
                                {{$typeIdSelect:=.getInput.searchTypeId}}
                                {{range $k,$v:=.ArticleType}}
                                    <option value="{{.Id}}"
                                    {{if eq $typeIdSelect $v.Id}}
                                            selected="selected"
                                    {{end}}>
                                    {{.Title}}
                                    </option>
                                {{end}}
                                </select>
                            </div>
                            <div class="form-group">
                                <label for="searchLabelId">标签</label>
                                <select id="searchLabelId" name="searchLabelId" title="请选择标签">
                                    <option value="0">全部</option>
                                {{$labelSelect:=.getInput.searchLabelId}}
                                {{range $k,$v:=.ArticleLabel}}
                                    <option value="{{.Id}}"
                                    {{if eq $labelSelect $v.Id}}
                                            selected="selected"
                                    {{end}}>
                                    {{$v.Title}}
                                    </option>
                                {{end}}
                                </select>
                            </div>
                            <div class="form-group">
                                <label for="searchSign">置顶</label>
                                <select id="searchSign" name="searchSign" title="请选择置顶">
                                {{$SignSelect:=.getInput.searchSign}}
                                {{range $k,$v:=.SignList}}
                                    <option value="{{$k}}"
                                    {{if eq $SignSelect $k}}
                                            selected="selected"
                                    {{end}}
                                    >{{$v}}</option>
                                {{end}}
                                </select>
                            </div>
                            <button type="submit" class="btn btn-info" style="border-radius: 5px;">搜索</button>
                            <a href="/adminArticleAdd" type="button" class="btn btn-success"
                               style="border-radius: 5px;">添加文章</a>
                            <a href="/adminArticleList" type="button" class="btn btn-danger"
                               style="border-radius: 5px;">撤回</a>
                        </form>

                    </div>
                </div>
                <div class="card-body">
                    <table class="table " cellspacing="0" width="100%">
                        <thead>
                        <tr>
                            <th>标题</th>
                            <th>封面</th>
                            <th>状态</th>
                            <th>点赞</th>
                            <th>嘘声</th>
                            <th>发表时间</th>
                            <th>操作</th>
                        </tr>
                        </thead>
                        <tbody>
                        {{range .Article}}
                        <tr>
                            <td>{{.Title}}</td>
                            <td>
                                <img src="{{.Image}}" class="profile-img pull-left" style="width: 152px;height: 93px">
                            </td>
                            <td>
                            {{if eq .Status 1}}
                                <label class="success">正常</label>
                            {{else}}
                                <label class="danger">禁用</label>
                            {{end}}
                            </td>
                            <td>{{.Thumbs}}</td>
                            <td>{{.Hiss}}</td>
                            <td>{{.CreateAt}}</td>
                            <td>
                                <button type="button" class="btn btn-xs btn-success" data-toggle="modal"
                                        data-target=".bs-example-modal-sm"
                                        onclick="sendCommentDiv({{.Id}})">
                                    评论
                                </button>
                                <a type="button" class="btn btn-xs btn-info" href="/adminArticleEdit?id={{.Id}}">
                                    编辑
                                </a>
                                <button type="button" class="btn btn-xs btn-danger"
                                        onclick="deleteData({{.Id}})">
                                    删除
                                </button>
                                <a type="button" class="btn btn-xs btn-success"
                                   href="/articleCommentGetList?searchArticleId={{.Id}}">
                                    评论列表
                                </a>
                                <!--编辑状态-->
                            {{if eq .Status 1}}
                                <button type="button" class="btn btn-xs btn-warning"
                                        onclick="updateStatusToClose({{.Id}})">
                                    禁用
                                </button>
                            {{else}}
                                <button type="button" class="btn btn-xs btn-success"
                                        onclick="updateStatusToOpen({{.Id}})">
                                    恢复
                                </button>
                            {{end}}
                                <!--编辑置顶-->
                            {{if eq .Sign 1}}
                                <button type="button" class="btn btn-xs btn-primary"
                                        onclick="updateSignToClose({{.Id}})">
                                    不置顶
                                </button>
                            {{else}}
                                <button type="button" class="btn btn-xs btnOrchid"
                                        onclick="updateSignToOpen({{.Id}})">
                                    置顶
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
                <h4 class="modal-title" id="mySmallModalLabel">评论</h4>
            </div>
            <div class="modal-body">
                <form class="form-horizontal">
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
        var searchTitle = {{.getInput.searchTitle}};
        var searchTypeId = {{.getInput.searchTypeId}};
        var searchLabelId = {{.getInput.searchLabelId}};
        var searchSign = {{.getInput.searchSign}};
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
                    window.location.href = "/adminArticleList?page=" + obj.curr + "&searchTitle=" + searchTitle
                            + "&searchTypeId=" + searchTypeId + "&searchLabelId=" + searchLabelId + "&searchSign=" + searchSign
                }
            }
        });
    });
</script>
<!-- 处理分页 -->

<!-- 操作按钮 -->
<script type="application/javascript">
    //回复评论赋值
    function sendCommentDiv($articleId) {
        var articleId = $articleId;
        $("#articleId").val(articleId);
    }

    //回复评论
    function articleCommentReply() {
        //获取用户参数
        var articleId = $('#articleId').val();
        var commentContent = $('#commentContent').val();
        //判断用户参数
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
                commentId: 0,
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


    //启用博客
    function updateStatusToOpen(id) {
        var sureUpdateStatusToOpen = confirm("你确定启用文章？");
        if (sureUpdateStatusToOpen === false) {
            alert("请确认后再重新提交");
            return false;
        }
        var ajaxActionUrl = '/adminArticleUpdateStatus';
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

    //禁用博客
    function updateStatusToClose(id) {
        var sureUpdateStatusToClose = confirm("你确定禁用文章？");
        if (sureUpdateStatusToClose === false) {
            alert("请确认后再重新提交");
            return false;
        }
        var ajaxActionUrl = '/adminArticleUpdateStatus';
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

    //删除博客
    function deleteData(id) {
        var sureDeleteData = confirm("你确定删除数据");
        if (sureDeleteData === false) {
            alert("请确认后再重新提交");
            return false;
        }
        var ajaxActionUrl = '/adminArticleDelete';
        /* ajax提交 */
        $.ajax({
            //提交数据的类型 POST GET
            type: "POST",
            //提交的网址
            url: ajaxActionUrl,
            //提交的数据
            data: {id: id},
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

    //置顶博客
    function updateSignToOpen(id) {
        var sureUpdateStatusToOpen = confirm("你确定置顶文章？");
        if (sureUpdateStatusToOpen === false) {
            alert("请确认后再重新提交");
            return false;
        }
        var ajaxActionUrl = '/adminArticleUpdateSign';
        /* ajax提交 */
        $.ajax({
            //提交数据的类型 POST GET
            type: "POST",
            //提交的网址
            url: ajaxActionUrl,
            //提交的数据
            data: {id: id, sign: 1},
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

    //取消置顶博客
    function updateSignToClose(id) {
        var sureUpdateStatusToClose = confirm("你确定取消置顶文章？");
        if (sureUpdateStatusToClose === false) {
            alert("请确认后再重新提交");
            return false;
        }
        var ajaxActionUrl = '/adminArticleUpdateSign';
        /* ajax提交 */
        $.ajax({
            //提交数据的类型 POST GET
            type: "POST",
            //提交的网址
            url: ajaxActionUrl,
            //提交的数据
            data: {id: id, sign: 2},
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