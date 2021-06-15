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
                        <div class="title">网站名称配置</div>
                    </div>
                </div>
                <div class="card-body">
                    <form class="form-horizontal">
                        <div class="form-group">
                            <label for="webTopTitle" class="col-sm-2 control-label">网站名称</label>
                            <div class="col-sm-10">
                                <input type="text" class="form-control" id="webTopTitle" name="webTopTitle"
                                       placeholder="网站名称" value="{{.webTopTitle}}">
                            </div>
                            <div class="form-group">
                                <div class="col-sm-offset-2 col-sm-10">
                                    <button type="button" class="btn btn-success" onclick="set()">设置</button>
                                </div>
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
<!-- 配置 -->
<script type="application/javascript">
    //配置
    function set() {
        //获取用户参数
        var webTopTitle = $('#webTopTitle').val();
        var ajaxActionUrl = '/adminSettingSetWebTitleConfig';
        /* ajax提交 */
        $.ajax({
            //提交数据的类型 POST GET
            type: "POST",
            //提交的网址
            url: ajaxActionUrl,
            //提交的数据
            data: {
                webTopTitle: webTopTitle,
            },
            //返回数据的格式
            datatype: "json",//"xml", "html", "script", "json", "jsonp", "text".
            //成功返回之后调用的函数
            success: function (data) {
                ////根据ajax返回参数判断验证码发送情况
                if (data.status === 1) {
                    layer.msg(data.message);
                    window.location.reload();
                } else if (data.status === -1) {
                    layer.msg(data.message);
                } else {
                    layer.msg('未知错误!');
                }
                ////根据ajax返回参数判断验证码发送情况
            }
        });
        /* ajax提交 */
    }
</script>