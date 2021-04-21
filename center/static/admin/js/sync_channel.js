function channelSync() {

    layer.confirm('确定要同步吗？', {
        btn: ['确定', '取消'], //按钮
        shade: false //不显示遮罩
    }, function () {
        $.ajax({
            type: "POST",
            url: '/admin/server_channel/sync',
            async: false,
            error: function (request) {
                layer.msg('网络错误', {icon: 1});
            },
            success: function (data) {
                if (data.status == 1) {
                    layer.msg(data.msg, {time: 2000, icon: 1});
                } else {
                    layer.msg(data.msg, {icon: 2});
                }
            }
        });
    }, function () {
        layer.msg('取消', {time: 2000, icon: 2});
    });
}

function channelClear() {

    layer.confirm('确定要清除同步记录吗吗？', {
        btn: ['确定', '取消'], //按钮
        shade: false //不显示遮罩
    }, function () {
        $.ajax({
            type: "POST",
            url: '/admin/server_channel/clearLog',
            async: false,
            error: function (request) {
                layer.msg('网络错误', {icon: 1});
            },
            success: function (data) {
                if (data.status == 1) {
                    layer.msg(data.msg, {time: 2000, icon: 1});
                } else {
                    layer.msg(data.msg, {icon: 2});
                }
            }
        });
    }, function () {
        layer.msg('取消', {time: 2000, icon: 2});
    });
}