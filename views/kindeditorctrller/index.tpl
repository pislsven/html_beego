<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>上传</title>
    <link rel="stylesheet" href="/static/editor/kindeditor/themes/default/default.css">
    <script src="/static/editor/kindeditor/kindeditor-all-min.js"></script>
    <script src="/static/editor/kindeditor/lang/zh-CN.js"></script>
    <script>
        KindEditor.ready(function (K) {
            const uploadbutton = K.uploadbutton({
                button: K('#upload_bt')[0],
                fieldName: 'imgFile',
                // url: '/upload',     //  带 / 从根目录开始  不带 从当前路径 开始
                url: 'upload',
                afterUpload: function (data) {
                    const inputInfo = document.querySelector("#imgsrc")

                    if (data.error === 0) {
                        const img = document.querySelector("#img")
                        img.setAttribute("src", data.url)

                        inputInfo.href = data.url
                        inputInfo.setAttribute("target", "_blank")
                    } else {
                        alert(data.message);
                        inputInfo.setAttribute("target", "")
                    }
                }
            })
            uploadbutton.fileBox.change(function(e){
                uploadbutton.submit();
            })
        })
    </script>
</head>
<body>
    <h1>上传</h1>
    <hr>
    <div>
        <img id="img" src="" alt="请选择上传图片" style="height: 100%" border="1">
        <a id="imgsrc" href="">预览</a>
        <button type="button" id="upload_bt" value="上传"></button>
        <!-- <input type="button" id="upload_bt" value="uploadBT"> -->
    </div>
</body>
</html>