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
</style>
<div class="col-md-8 themelazer_content">
    <!--循环文章-->
{{range .Article}}
    <div class="themelazer_blog_style_one">
        <div class="sigle_blog_style_one">
            <a href="/articleInfo?id={{.Id}}" class="themelazer_post_categories">{{.articleType}}</a>
            <h3>
                <a href="/articleInfo?id={{.Id}}">{{.Title}}</a>
            </h3>
            <div class="meta-info">
                <ul>
                    <li class="post-author"><a href="/articleInfo?id={{.Id}}" tabindex="0">
                        <img src="static/home/blog.jpg" alt="Amelia">{{.Author}}</a>
                    </li>
                    <li class="post-date">{{.CreateAt}}</li>
                </ul>
            </div>
            <!--封面-->
            <div class="themelazer_grid_image_wrapper">
                <div class="themelazer_xl_grid_image_front">
                    <a href="/articleInfo?id={{.Id}}"><span class="themelazer_xl_grid_image_header_absolute"
                                                            style="background-image: url({{.Image}});"> </span>
                    </a>
                </div>
            </div>
            <!--封面-->
            <!--描述-->
            <div class="text_box">
                <div class="themelazer_post_entry">
                    <p>
                    {{.Describe}}
                    </p>
                </div>
            </div>
            <!--描述-->
            <!--阅读按钮-->
            <div class="themelazer_post_meta" style="text-align: center;">
                <div class="themelazer_grid_image_wrapper">
                    <div class="themelazer_read_more_style2_wrapper" style="float:none !important;margin-right: 0 !important;">
                        <a href="/articleInfo?id={{.Id}}" class="themelazer_read_more">阅读</a>
                    </div>
                </div>
            </div>
            <!--阅读按钮-->
        </div>
    </div>
{{end}}
    <!--循环文章-->
    <!--分页-->
{{if gt .PageCount 0 }}
    <div class="themelazer-pagination">
        <div class="themelazer-pagination-wrapper" id="page">
        </div>
    </div>
{{end}}
    <!--分页-->

</div>
<!-- 处理分页 -->
<link href="static/home/layui/css/layui.css" rel="stylesheet"/>
<script src="static/home/layui/layui.js?s=36"></script>
<script>
    layui.use('laypage', function () {
        var laypage = layui.laypage;
        var articleTypeId ={{.articleTypeId}};
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
                    window.location.href = "/articleIndex?page=" + obj.curr + "&articleTypeId=" + articleTypeId
                }
            }
        });
    });
</script>
<!-- 处理分页 -->