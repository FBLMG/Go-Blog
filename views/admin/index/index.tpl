<div class="side-body padding-top">
    <div class="row">
        <div class="col-lg-3 col-md-6 col-sm-6 col-xs-12">
            <a href="/adminArticleList">
                <div class="card red summary-inline">
                    <div class="card-body">
                        <i class="icon glyphicon glyphicon-book fa-4x"></i>
                        <div class="content">
                            <div class="title">{{.ArticleCount}}</div>
                            <div class="sub-title">日志</div>
                        </div>
                        <div class="clear-both"></div>
                    </div>
                </div>
            </a>
        </div>
        <div class="col-lg-3 col-md-6 col-sm-6 col-xs-12">
            <a href="/adminArticleTypeList">
                <div class="card yellow summary-inline">
                    <div class="card-body">
                        <i class="icon glyphicon glyphicon-list fa-4x"></i>
                        <div class="content">
                            <div class="title">{{.ClassifyCount}}</div>
                            <div class="sub-title">分类</div>
                        </div>
                        <div class="clear-both"></div>
                    </div>
                </div>
            </a>
        </div>
        <div class="col-lg-3 col-md-6 col-sm-6 col-xs-12">
            <a href="/adminArticleLabelList">
                <div class="card green summary-inline">
                    <div class="card-body">
                        <i class="icon fa fa-tags fa-4x"></i>
                        <div class="content">
                            <div class="title">{{.LabelCount}}</div>
                            <div class="sub-title">标签</div>
                        </div>
                        <div class="clear-both"></div>
                    </div>
                </div>
            </a>
        </div>
        <div class="col-lg-3 col-md-6 col-sm-6 col-xs-12">
            <a href="/adminFriendshipList">
                <div class="card blue summary-inline">
                    <div class="card-body">
                        <i class="icon fa fa-share-alt fa-4x"></i>
                        <div class="content">
                            <div class="title">{{.FriendCount}}</div>
                            <div class="sub-title">友情链接</div>
                        </div>
                        <div class="clear-both"></div>
                    </div>
                </div>
            </a>
        </div>
    </div>
    <div class="row  no-margin-bottom">
        <!--置顶2篇文章-->
        <div class="col-sm-6 col-xs-12">
            <div class="row">
            {{range .NewArticleTwo}}
                <div class="col-md-6 col-sm-12">
                    <div class="thumbnail no-margin-bottom">
                        <img src="{{.Image}}" class="img-responsive">
                        <div class="caption">
                            <h3 id="thumbnail-label">{{.Title}}
                                <a class="anchorjs-link" href="#thumbnail-label">
                                    <span class="anchorjs-icon"></span>
                                </a>
                            </h3>
                            <p>{{.Describe}}</p>
                            <p>
                                <button onclick="updateSignToClose({{.Id}})" class="btn btn-danger" role="button">取消
                                </button>
                                <a href="/adminArticleEdit?id={{.Id}}" class="btn btn-info" role="button">编辑</a>
                            </p>
                        </div>
                    </div>
                </div>
            {{end}}
            </div>
        </div>
        <!--置顶2篇文章-->
        <!--待审核评论-->
        <div class="col-sm-6 col-xs-12">
            <div class="card card-success">
                <div class="card-header">
                    <div class="card-title">
                        <div class="title"><i class="fa fa-comments-o"></i> 待审核评论</div>
                    </div>
                    <div class="clear-both"></div>
                </div>
                <div class="card-body no-padding">
                    <ul class="message-list">
                    {{range .NewArticleFour}}
                        <a href="/articleCommentGetList?searchCommentId={{.Id}}">
                            <li>
                                <img src="static/home/blog.jpeg" class="profile-img pull-left">
                                <div class="message-block">
                                    <div><span class="username">{{.UserNickname}}</span> <span
                                            class="message-datetime"> {{.CreateAt}}</span>
                                    </div>
                                    <div class="message">{{.Content}}
                                    </div>
                                </div>
                            </li>
                        </a>
                    {{end}}
                        <a href="/articleCommentGetList" id="message-load-more">
                            <li class="text-center load-more">
                                <i class="fa fa-refresh"></i> 查看更多..
                            </li>
                        </a>
                    </ul>
                </div>
            </div>
        </div>
        <!--待审核评论-->
    </div>
</div>
<!-- 操作按钮 -->
<script type="application/javascript">
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
