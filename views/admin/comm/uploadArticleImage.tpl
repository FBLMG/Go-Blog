<!-----图片上传模态框------>
<link rel="stylesheet" type="text/css" href="static/thirdParty/layui/css/layui.css">
<!---上传封面模态框---->
<div class="modal fade" id="imageUpload" tabindex="-1" role="dialog" aria-labelledby="myModalLabel">
    <div class="modal-dialog" role="document" style="text-align: center;height: 120px;width: 300px;">
        <div class="modal-content">
            <!-- -->
            <div class="layui-upload" style="text-align: center;height: 120px;width: 300px;line-height: 120px;">
                <button type="button" class="layui-btn" id="test1">
                    上传图片
                </button>
                <input class="layui-upload-file" type="file" name="file">
            </div>
            <!-- -->
        </div>
    </div>
</div>
<!-----图片上传模态框------>
<!--- 图片上传方法 --->
<script type="application/javascript">
    layui.use('upload', function () {
        var upload = layui.upload;
        //执行实例
        var uploadInst = upload.render({
            elem: '#test1'               //绑定元素
            , url: "/adminUploadArticle" //上传接口
            , done: function (res) {
                if (res.status === 1) {
                    document.getElementById("blogImage").src = res.message;         //重新设置博客封面地址
                    document.getElementById("blogImageUrl").value = res.message;    //赋值图片地址
                    $('#imageUpload').modal('hide');                                //隐藏模态框
                    return layer.msg('图片上传成功！');
                } else if (res.status === -1) {
                    return layer.msg(res.message);
                }
            }
            , error: function () {
                //请求异常回调
                return layer.msg('图片上传异常');
            }
        });
    });
</script>
<!--- 图片上传方法 --->
<!-- 富文本编辑框 -->
<script src="static/thirdParty/ueditor/ueditor.config.js"></script>
<script src="static/thirdParty/ueditor/ueditor.all.js"></script>
<script type="text/javascript">
    var ue = UE.getEditor('container', {
        "initialFrameWidth": null
    });
</script>
<!-- 富文本编辑框 -->