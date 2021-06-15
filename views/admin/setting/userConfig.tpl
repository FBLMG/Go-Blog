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
                        <div class="title">自我介绍</div>
                    </div>
                </div>
                <div class="card-body">
                    <form class="form-horizontal">
                        <div class="form-group">
                            <label for="userTitle" class="col-sm-2 control-label">昵称</label>
                            <div class="col-sm-10">
                                <input type="text" class="form-control" id="userTitle" name="userTitle"
                                       placeholder="昵称" value="{{.userTitle}}">
                            </div>
                        </div>
                        <div class="form-group">
                            <label for="inputPassword3" class="col-sm-2 control-label">头像</label>
                            <div class="col-sm-10">
                                <img src="{{.userImage}}" alt="点击修改图片" id="settingImage"
                                     name="settingImage" class="img-thumbnail" data-toggle="modal"
                                     data-target="#imageUpload"
                                     style="cursor:pointer; width:100px;height:100px">
                                <input type="hidden" name="settingImageUrl" id="settingImageUrl"
                                       value="{{.userImage}}">
                            </div>
                        </div>
                        <div class="form-group">
                            <label for="userDesc" class="col-sm-2 control-label">自我介绍</label>
                            <div class="col-sm-10">
                                <textarea class="form-control" rows="3" id="userDesc"
                                          name="userDesc">{{.userDesc}}</textarea>
                            </div>
                        </div>
                        <div class="form-group">
                            <div class="col-sm-offset-2 col-sm-10">
                                <button type="button" class="btn btn-success" onclick="set()">设置</button>
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
        var userTitle = $('#userTitle').val();
        var userDesc = $('#userDesc').val();
        var settingImageUrl = $('#settingImageUrl').val();
        var ajaxActionUrl = '/adminSettingSetUserConfig';
        /* ajax提交 */
        $.ajax({
            //提交数据的类型 POST GET
            type: "POST",
            //提交的网址
            url: ajaxActionUrl,
            //提交的数据
            data: {
                userTitle: userTitle,
                userDesc: userDesc,
                userImage: settingImageUrl,
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