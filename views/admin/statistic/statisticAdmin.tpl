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
                        <!--插入表单查询-->
                        <form class="form-inline">
                            <div class="form-group">
                                <label for="searchStartDateFilter">开始日期</label>
                                <input type="text" class="form-control" id="searchStartDateFilter" placeholder="开始日期"
                                       style="border-radius: 5px;"
                                       value="{{.getInput.searchStartDate}}">
                                <input type="hidden" id="searchStartDate" name="searchStartDate"
                                       value="{{.getInput.searchStartDate}}">
                            </div>
                            <div class="form-group">
                                <label for="searchEndDateFilter">结束日期</label>
                                <input type="text" class="form-control" id="searchEndDateFilter" placeholder="结束日期"
                                       style="border-radius: 5px;" value="{{.getInput.searchEndDate}}">
                                <input type="hidden" id="searchEndDate" name="searchEndDate"
                                       value="{{.getInput.searchEndDate}}">
                            </div>
                            <div class="form-group">
                                <label for="searchAdminId">管理员</label>
                                <select id="searchAdminId" name="searchAdminId" title="请选择管理员">
                                    <option value="0">全部</option>
                                {{$adminIdSelect:=.getInput.searchAdminId}}
                                {{range $k,$v:=.AdminUser}}
                                    <option value="{{$v.Id}}"
                                    {{if eq $adminIdSelect $v.Id}}
                                            selected="selected"
                                    {{end}}>
                                    {{$v.Nickname}}
                                    </option>
                                {{end}}
                                </select>
                            </div>
                            <button type="submit" class="btn btn-info" style="border-radius: 5px;">搜索</button>
                            <a href="/adminStatisticUserLogin" type="button" class="btn btn-danger"
                               style="border-radius: 5px;">撤回</a>
                            <button type="button" class="btn btn-warning" style="border-radius: 5px;"
                                    onclick="deleteData()">删除
                            </button>
                            <a href="/adminStatisticUserLoginExcel?searchStartDate={{.getInput.searchStartDate}}&searchEndDate
={{.getInput.searchEndDate}}&searchAdminId={{.getInput.searchAdminId}}" type="button" class="btn btn-success"
                               style="border-radius: 5px;">导出数据</a>
                        </form>
                        <!--插入表单查询-->
                    </div>
                </div>
                <div class="card-body">
                    <div class="panel panel-default">
                        <!-- Table -->
                        <table class="table">
                            <thead>
                            <tr>
                                <th>管理员</th>
                                <th>IP</th>
                                <th>地址</th>
                                <th>访问时间</th>
                            </tr>
                            </thead>
                            <tbody>
                            {{range .StatisticList}}
                            <tr>
                                <td>{{.AdminUserName}}</td>
                                <td>{{.Ip}}</td>
                                <td>{{.Address}}</td>
                                <td>{{.CreateAt}}</td>
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
<!-- 分页样式 -->
<link href="static/home/layui/css/layui.css" rel="stylesheet"/>
<!-- 分页样式 -->
<!-- 处理分页 -->
<script type="text/javascript" src="static/thirdParty/layui/layui.all.js"></script>
<script>
    layui.use('laypage', function () {
        var laypage = layui.laypage;
        var searchStartDate = {{.getInput.searchStartDate}};
        var searchEndDate = {{.getInput.searchEndDate}};
        var searchAdminId = {{.getInput.searchAdminId}};
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
                    window.location.href = "/adminStatisticUserLogin?page=" + obj.curr + "&searchStartDate=" + searchStartDate
                            + "&searchEndDate=" + searchEndDate + "&searchAdminId=" + searchAdminId
                }
            }
        });
    });
</script>
<!-- 处理分页 -->
<!-- 初始化日期 -->
<script>
    //设置日期
    layui.use('laydate', function () {
        var laydate = layui.laydate;
        laydate.render({
            elem: '#searchStartDateFilter',
            type: 'date',
            done: function (value) {
                document.getElementById('searchStartDate').value = value;
            }
        });
        laydate.render({
            elem: '#searchEndDateFilter'
            , type: 'date'
            , done: function (value) {
                //赋值
                document.getElementById('searchEndDate').value = value;
            }
        });
    });

    //删除数据
    function deleteData() {
        //获取用户参数
        var searchStartDate = $('#searchStartDate').val();
        var searchEndDate = $('#searchEndDate').val();
        var searchAdminId = $('#searchAdminId').val();
        var ajaxActionUrl = '/adminStatisticUserLoginDelete';
        /* ajax提交 */
        $.ajax({
            //提交数据的类型 POST GET
            type: "POST",
            //提交的网址
            url: ajaxActionUrl,
            //提交的数据
            data: {
                searchStartDate: searchStartDate,
                searchEndDate: searchEndDate,
                searchAdminId: searchAdminId,
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
        /* ajax提交 */

    }
</script>
<!-- 初始化日期 -->