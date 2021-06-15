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
                        <div class="title">编辑文章</div>
                    </div>
                </div>
                <div class="card-body">
                    <form class="form-horizontal">
                        <!-- 文章ID -->
                        <input type="hidden" name="blogId" id="blogId"
                               value="{{.Article.Id}}">
                        <!-- 文章标题 -->
                        <div class="form-group">
                            <label for="blogTitle" class="col-sm-2 control-label">标题</label>
                            <div class="col-sm-10">
                                <input type="text" class="form-control" id="blogTitle" name="blogTitle"
                                       placeholder="博客标题" value="{{.Article.Title}}">
                            </div>
                        </div>
                        <!-- 文章描述 -->
                        <div class="form-group">
                            <label for="blogDesc" class="col-sm-2 control-label">描述</label>
                            <div class="col-sm-10">
                                <input type="text" class="form-control" id="blogDesc" name="blogDesc"
                                       placeholder="博客描述" value="{{.Article.Describe}}">
                            </div>
                        </div>
                        <!-- 文章分类 -->
                        <div class="form-group">
                            <label for="typeId" class="col-sm-2 control-label">分类</label>
                            <div class="col-sm-10">
                                <select id="typeId" name="typeId">
                                {{$typeSelect:=.Article.TypeId}}
                                {{range $k,$v:=.ArticleType}}
                                    <option value="{{.Id}}"
                                    {{if eq $typeSelect $v.Id}}
                                            selected="selected"
                                    {{end}}
                                    >{{.Title}}</option>
                                {{end}}
                                </select>
                            </div>
                        </div>
                        <!-- 文章标签 -->
                        <div class="form-group">
                            <label for="labelId" class="col-sm-2 control-label">标签</label>
                            <div class="col-sm-10">
                                <select id="labelId" name="labelId">
                                {{$labelSelect:=.Article.LabelId}}
                                {{range $k,$v:=.ArticleLabel}}
                                    <option value="{{.Id}}"
                                    {{if eq $labelSelect $v.Id}}
                                            selected="selected"
                                    {{end}}
                                    >{{$v.Title}}</option>
                                {{end}}
                                </select>
                            </div>
                        </div>
                        <!-- 文章封面 -->
                        <div class="form-group">
                            <label for="blogImage" class="col-sm-2 control-label">封面</label>
                            <div class="col-sm-10">
                                <img src="{{.imageUrl}}{{.Article.Image}}" alt="点击修改图片" id="blogImage"
                                     name="blogImage" class="img-thumbnail" data-toggle="modal"
                                     data-target="#imageUpload"
                                     style="cursor:pointer; width:60px;height:60px">
                            </div>
                            <!-- 封面地址 -->
                            <input type="hidden" name="blogImageUrl" id="blogImageUrl"
                                   value="{{.Article.Image}}">
                            <!-- 封面地址 -->
                        </div>
                        <!-- 文章状态 -->
                        <div class="form-group">
                            <label for="status" class="col-sm-2 control-label">状态</label>
                            <div class="col-sm-10">
                                <select id="status" name="status">
                                {{$statusSelect:=.Article.Status}}
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
                        <!-- 文章置顶 -->
                        <div class="form-group">
                            <label for="blogSign" class="col-sm-2 control-label">置顶</label>
                            <div class="col-sm-10">
                                <select id="blogSign" name="blogSign">
                                {{if eq .Article.Sign 1}}
                                    <option value="1" selected="selected">是</option>
                                    <option value="2">否</option>
                                {{else}}
                                    <option value="1">是</option>
                                    <option value="2" selected="selected">否</option>
                                {{end}}
                                </select>
                            </div>
                        </div>
                        <!-- 文章排序 -->
                        <div class="form-group">
                            <label for="blogSort" class="col-sm-2 control-label">序号</label>
                            <div class="col-sm-10">
                                <input type="text" class="form-control" id="blogSort" name="blogSort"
                                       placeholder="博客序号" value="{{.Article.Sort}}">
                            </div>
                        </div>
                        <!-- 文章URL -->
                        <div class="form-group">
                            <label for="blogUrl" class="col-sm-2 control-label">链接</label>
                            <div class="col-sm-10">
                                <input type="text" class="form-control" id="blogUrl" name="blogUrl"
                                       placeholder="博客链接" value="{{.Article.Url}}">
                            </div>
                        </div>
                        <!-- 文章内容 -->
                        <div class="form-group">
                            <label for="container" class="col-sm-2 control-label">内容</label>
                            <div class="col-sm-10">
                               <textarea id="container" name="container">
                               {{.Article.Content}}
                               </textarea>
                            </div>
                        </div>
                        <!-- 文章编辑 -->
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
<!-- 编辑文章 -->
<script type="application/javascript">
    //编辑文章
    function edit() {
        var blogId = $('#blogId').val();
        var blogTitle = $('#blogTitle').val();
        var blogDesc = $('#blogDesc').val();
        var typeId = $("#typeId").val();
        var labelId = $("#labelId").val();
        var blogImageUrl = $('#blogImageUrl').val();
        var status = $("#status").val();
        var blogSign = $("#blogSign").val();
        var blogSort = $("#blogSort").val();
        var blogUrl = $("#blogUrl").val();
        var container = UE.getEditor('container').getContent();
        var ajaxActionUrl = '/adminArticleUpdate';
        if (blogId === '') {
            layer.msg('ID不允许为空！');
            return false;
        }
        if (blogTitle === '') {
            layer.msg('标题不允许为空！');
            return false;
        }
        if (blogDesc === '') {
            layer.msg('描述不允许为空！');
            return false;
        }
        if (typeId === '') {
            layer.msg('博客类别不允许不允许为空！');
            return false;
        }
        if (labelId === '') {
            layer.msg('博客标签不允许为空！');
            return false;
        }
        if (blogImageUrl === '') {
            layer.msg('封面不允许为空！');
            return false;
        }
        if (status === "") {
            layer.msg('状态不允许为空！');
            return false;
        }
        if (container === "") {
            layer.msg('内容不允许为空！');
            return false;
        }
        /* ajax提交 */
        $.ajax({
            //提交数据的类型 POST GET
            type: "POST",
            //提交的网址
            url: ajaxActionUrl,
            //提交的数据
            data: {
                blogId: blogId,
                blogTitle: blogTitle,
                blogDesc: blogDesc,
                typeId: typeId,
                labelId: labelId,
                blogImageUrl: blogImageUrl,
                status: status,
                blogSign: blogSign,
                blogSort: blogSort,
                blogUrl: blogUrl,
                container: container
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