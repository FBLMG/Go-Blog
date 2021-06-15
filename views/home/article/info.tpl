<style type="text/css">
    .layui-laypage .layui-laypage-curr .layui-laypage-em {
        position: absolute;
        left: -1px;
        top: -1px;
        padding: 1px;
        width: 100%;
        height: 100%;
        background: #fff !important;
        border: 1px solid #be967f !important;
        color: #000 !important;
    }

    .layui-laypage .layui-laypage-curr em {
        position: relative;
        color: #000 !important;
    }
    .themelazer_single_content strong {
        color: #000;
        font-weight: 600;
    }
    .themelazer_single_content pre {
        display: block;
        padding: 9.5px;
        margin: 0 0 10px;
        font-size: 13px;
        line-height: 1.42857143;
        word-break: break-all;
        word-wrap: break-word;
        color: #333;
        background-color: #f5f5f5;
        border: 1px solid #ccc;
        border-radius: 4px;
    }
    .themelazer_single_content p {
        color: #343434;
        font-weight: normal;
        line-height: 24px;
        margin: 0 0 15px;
    }
</style>
<div class="col-md-8">
    <div class="single_header_wrapper">
        <div class="themelazer_post_categories">
            <a href="#">{{.ArticleInfo.ArticleType}}</a>
        </div>
        <h1>
        {{.ArticleInfo.Title}}
        </h1>
        <div class="meta-info">
            <ul>
                <li class="post-author"><a href="#" tabindex="0">
                    <img src="static/home/blog.jpg" alt="{{.ArticleInfo.AdminUser}}">{{.ArticleInfo.AdminUser}}</a>
                </li>
                <li class="post-date">{{.ArticleInfo.CreateAt}}</li>
               {{/* <li class="post-view">50K Views</li>
                <li class="post-comment">4 Comments</li>*/}}
            </ul>
        </div>
    </div>
    <div class="themelazer_single_feature">
        <img src="{{.ArticleInfo.Image}}" alt="image">
    </div>
    <div class="themelazer_single_content">
        <blockquote>
            <p style="color: #0f0f0f">{{.ArticleInfo.Describe}}</p>
        </blockquote>
        <p>{{str2html .ArticleInfo.Content}}</p>
    </div>
    <div class="themelazer_tag_share">
        <div class="blog-tags">
            <a href="#">{{.ArticleInfo.ArticleLabel}}</a>
        {{/*<a href="#">Gallery</a>*/}}
        {{/*<a href="#">Slideshow</a>*/}}
        </div>
        <div class="blog-social-list  padding-sm-tb-20">
        {{/* <a href="#" class="facebook-bg">
                <i class="fab fa-facebook-f"></i>
            </a>
            <a href="#" class="twitter-bg">
                <i class="fab fa-twitter"></i>
            </a>
            <a href="#" class="linkedin-bg">
                <i class="fab fa-linkedin-in"></i>
            </a>
            <a href="#" class="pinterest-bg">
                <i class="fab fa-pinterest-p"></i>
            </a>*/}}
        </div>
    </div>
    <!--***暂时封闭***-->
{{/*<div class="author_info">
        <div class="author_avatar"><img alt="image" src="image/blog/author3.jpg"></div>
        <div class="author_description">
            <h5 class="author_title">Jennifer<a href="#">
            </a>
            </h5>
            <div class="author_bio">
                <p>Phasellus tellus tellus, imperdiet ut imperdiet eu, iaculis a sem. Donec vehicula luctus nunc in
                    laoreet. Aliquam erat volutpat. Suspendisse vulputate porttitor imentum.</p>
            </div>
            <div class="themelazer-author-social-links">
                <div class="themelazer-social-links-items">
                    <div class="themelazer-social-links-item">
                        <a href="#" class="themelazer-social-links-link themelazer-facebook" target="_blank"> <i
                                class="fab fa-facebook-f"></i> </a>
                    </div>
                    <div class="themelazer-social-links-item">
                        <a href="#" class="themelazer-social-links-link  themelazer-twitter" target="_blank"> <i
                                class="fab fa-twitter"></i> </a>
                    </div>
                    <div class="themelazer-social-links-item">
                        <a href="#" class="themelazer-social-links-link themelazer-youtube" target="_blank"> <i
                                class="fab fa-youtube"></i> </a>
                    </div>
                </div>
            </div>
        </div>
    </div>*/}}
    <!--***相同文章***-->
    <div class="themelazer_related_post">
        <div class="container">
            <div class="themelazer_title_head">
                <h3>猜你喜欢</h3>
            </div>
        </div>
    {{range .IdenticalArticleList}}
        <div class=" blog-style-one blog-small-grid">
            <div class="single-blog-style-one">
                <div class="img-box">
                    <img src="{{.Image}}" alt="Awesome Image">
                </div>
                <div class="themelazer_post_categories">
                    <a href="/articleType?id={{.ArticleTypeId}}">{{.ArticleType}}</a>
                </div>

                <div class="text-box">
                    <h3>
                        <a href="/articleInfo?id={{.Id}}">{{.Title}}</a>
                    </h3>
                    <div class="meta-info">
                        <ul>
                            <li class="post-author"><a href="/articleInfo?id={{.Id}}" tabindex="0">
                                <img src="static/home/blog.jpg" alt="{{.AdminUser}}">{{.AdminUser}}</a>
                            </li>
                            <li class="post-date">{{.CreateAt}}</li>
                        </ul>
                    </div>
                </div>
            </div>
        </div>
    {{end}}
    </div>
    <!--***文章评论列表***-->
    <div class="themelazer_title_head">
        <h3>{{.PageCount}} Comments</h3>
    </div>
    <ul class="comments_wp_list">
    {{range .ArticleCommentList}}
        <li>
            <div class="comment_themelazer">
                <div class="comment-author">
                    <a href="#">
                    {{if eq .AdminStatus 1}}
                        <img src="static/home/admin.jpeg" alt="comments-user">
                    {{else}}
                        <img src="static/home/blog.jpeg" alt="comments-user">
                    {{end}}
                    </a>
                </div>
                <div class="comment-content">
                    <div class="comment-meta">
                        <h6>
                        {{.UserNickname}}
                        </h6>
                    </div>
                    <p>{{.Content}}</p>
                    <div class="themelazer-comment-icon">
                        <p><i class="far fa-clock"></i>
                            <span class="margin-left-10">{{.CreateAt}}</span>
                            <span class="float-right margin-left-20"><a data-toggle="modal"
                                                                        data-target=".bs-example-modal-lg"
                                                                        onclick="sendReply({{.UserNickname}},{{.Id}})">回复</a></span>
                        </p>
                    </div>
                </div>
            </div>
        </li>
    {{end}}
    </ul>
{{if  gt .PageCount 0 }}
    <div class="themelazer-pagination">
        <div class="themelazer-pagination-wrapper" id="page">
        </div>
    </div>
{{end}}
    <!--***文章评论***-->
    <div class="themelazer-comment-wrapper mb-50">
        <div class="themelazer_title_head">
            <h3>文章评论</h3>
        </div>
        <form method="post" id="contact-form">
            <div class="row" id="sendCommentDiv">
                <!---隐藏区域--->
                <input type="hidden" name="commentId" id="commentId">
                <input type="hidden" name="commentUserName" id="commentUserName">
                <!---隐藏区域--->
                <div class="col-md-6">
                    <div class="single-input-item">
                        <label>
                            <input type="text" name="userNickname" id="userNickname" placeholder="昵称 *" required="">
                        </label>
                    </div>
                </div>
                <div class="col-md-6">
                    <div class="single-input-item">
                        <label>
                            <input type="email" name="userEmail" id="userEmail" placeholder="邮箱地址 *" required="">
                        </label>
                    </div>
                </div>
                <div class="col-12">
                    <div class="single-input-item">
                        <label for="con_message" class="sr-only m-0"></label>
                        <textarea name="content" id="content" cols="30" rows="7" placeholder="评论内容 *"
                                  required="" style="resize: none;"></textarea>
                    </div>
                </div>
            </div>
            <button type="button" class="submit_themelazern_large" onclick="sendComment()">评论</button>
            <button type="button" class="submit_themelazern_large" style="margin-left: 10px;background-color:#d43f3a;"
                    onclick="cancelReply()">
                取消回复
            </button>
        </form>
    </div>
    <!--***文章评论***-->
</div>
<!--  弹窗样式 -->
<link rel="stylesheet" href="static/thirdParty/layer/skin/layer.css"/>
<!-- 评论文章 -->
<script type="application/javascript">
    //回复评论
    function sendReply($userName, $commentId) {
        //设置值
        var userName = $userName;
        var commentId = $commentId;
        //设置评论内容
        $("#content").val();
        $("#content").val("回复：" + userName + " ");
        //设置评论昵称
        $("#commentUserName").val();
        $("#commentUserName").val("回复：" + userName + " ");
        //设置评论Id
        $("#commentId").val(0);
        $("#commentId").val(commentId);
        //跳转定位
        $("html, body").animate({
            scrollTop: $("#sendCommentDiv").offset().top
        });
        return false;
    }

    //取消回复
    function cancelReply() {
        $("#content").val("");
        $("#commentId").val(0);
        $("#commentUserName").val("");
    }

    //评论文章
    function sendComment() {
        //获取用户参数
        var commentUserName = $('#commentUserName').val();
        var userNickname = $('#userNickname').val();
        var userEmail = $('#userEmail').val();
        var content = $("#content").val();
        var commentId = $("#commentId").val();
        var articleId = "{{.ArticleInfo.Id}}";
        var ajaxActionUrl = '/articleComment';
        //判断用户参数
        if (userNickname === '') {
            layer.msg('昵称不允许为空！');
            return false;
        }
        if (userEmail === '') {
            layer.msg('邮箱不允许为空！');
            return false;
        }
        if (content === '') {
            layer.msg('内容不允许为空！');
            return false;
        }
        if (articleId === '') {
            layer.msg('文章获取失败！');
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
                userNickname: userNickname,
                userEmail: userEmail,
                content: content,
                articleId: articleId,
                commentId: commentId,
                commentUserName: commentUserName
            },
            //返回数据的格式
            datatype: "json",//"xml", "html", "script", "json", "jsonp", "text".
            //成功返回之后调用的函数
            success: function (data) {
                ////根据ajax返回参数判断添加情况
                if (data.status === 1) {
                    layer.msg(data.message);
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

    //回复文章评论
</script>
<!-- 处理分页 -->
<link href="static/home/layui/css/layui.css" rel="stylesheet"/>
<script src="static/home/layui/layui.js?s=36"></script>
<script>
    layui.use('laypage', function () {
        var laypage = layui.laypage;
        var articleId = {{.ArticleInfo.Id}};
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
                    window.location.href = "/articleInfo?id=" + articleId + "&page=" + obj.curr
                }
            }
        });
    });
</script>
<!-- 处理分页 -->