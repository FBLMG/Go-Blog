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
                        <div class="title">
                            <a href="/adminArticleTypeAdd" type="button" class="btn btn-success">添加分类</a>
                        </div>
                    </div>
                </div>
                <div class="card-body">
                    <div class="panel panel-default">
                        <!-- **循环数据** -->
                        <table class="table">
                            <thead>
                            <tr>
                                <th>标题</th>
                                <th>状态</th>
                                <th>创建时间</th>
                                <th>更新时间</th>
                                <th>操作</th>
                            </tr>
                            </thead>
                            <tbody>
                            {{range .Classify}}
                            <tr>
                                <td>{{.Title}}</td>
                                <td>
                                {{if eq .Status 1}}
                                    <label class="success">正常</label>
                                {{else}}
                                    <label class="danger">禁用</label>
                                {{end}}
                                </td>
                                <td>{{.CreateAt}}</td>
                                <td>{{.UpdateAt}}</td>
                                <td>
                                    <a class="btn btn-xs btn-info" href="/adminArticleTypeEdit?id={{.Id}}">
                                        编辑
                                    </a>
                                    <button type="button" class="btn btn-xs btn-danger"
                                            onclick="deleteData({{.Id}})">
                                        删除
                                    </button>
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
                                </td>
                            </tr>
                            {{end}}
                            </tbody>
                        </table>
                        <!-- **循环数据** -->
                    </div>
                </div>
            </div>
        </div>
        <!-- 分页 -->
        <div class="layui-row" id="page" style="text-align: center;"></div>
        <!-- 分页 -->
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
                    window.location.href = "/adminArticleTypeList?page=" + obj.curr
                }
            }
        });
    });
</script>
<!-- 处理分页 -->
<!-- 操作按钮 -->
<script type="application/javascript">
    //启用博客分类
    function updateStatusToOpen(id) {
        var sureUpdateStatusToOpen = confirm("你确定启用分类？");
        if (sureUpdateStatusToOpen === false) {
            alert("请确认后再重新提交");
            return false;
        }
        //
        var ajaxActionUrl = '/adminArticleTypeUpdateStatus';
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

    //禁用博客分类
    function updateStatusToClose(id) {
        var sureUpdateStatusToClose = confirm("你确定禁用分类？");
        if (sureUpdateStatusToClose === false) {
            alert("请确认后再重新提交");
            return false;
        }
        //
        var ajaxActionUrl = '/adminArticleTypeUpdateStatus';
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

    //删除博客分类
    function deleteData(id) {
        var sureDeleteData = confirm("你确定删除数据");
        if (sureDeleteData === false) {
            alert("请确认后再重新提交");
            return false;
        }
        //
        var ajaxActionUrl = '/adminArticleTypeDelete';
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
</script>
<!-- 操作按钮 -->