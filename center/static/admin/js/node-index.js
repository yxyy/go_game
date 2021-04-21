/**
 * Created by Administrator on 2017/6/22.
 */
//模态框关闭时，清空form的下几项数据
$('#myModal').on('hide.bs.modal', function () {
    $("#group_name").val("")
    $("#permit_manager").val("")
    $("#controller_action").val("")
})

$(document).ready(function () {
    var updateOutput = function (e) {
        var list = e.length ? e : $(e.target),
            output = list.data('output');
        if (window.JSON) {
            output.val(window.JSON.stringify(list.nestable('serialize'))); //, null, 2));
        } else {
            output.val('浏览器不支持');
        }
    };
});

//权限列表是否收起/展开
function is_show_permit(obj){
    var is_show = $(obj).html();
    if(is_show == '-'){
        $(obj).next().next().hide();
        $(obj).html('+');
        $(obj).attr("data-beforeContent","+");
    } else if(is_show == '+'){
        $(obj).next().next().show();
        $(obj).html('-')
        $(obj).attr("data-beforeContent","-");
    }
}

//修改submit_method的值为add,并显示模态框
function submit_add_method(value){
    $("#submit_method").val(value);
    $('#myModal').modal('show');
}

//修改权限填上模态框的信息
function submit_edit_method(value,obj){
    $("#submit_method").val(value);
    var pid = $(obj).data('id');
    $("#pid").val(pid);
    var formData = new FormData();
    formData.append("pid",pid);
    $.ajax({
        url:"getPermitInfoById",
        data:formData,
        type:"post",
        dataType:"json",
        contentType:false,
        processData:false,
        success:function(data){
            if(data.code == 0){
                console.log(data.code);
                alert(data.msg);
                return;
            }else if(data.code == 1){
                $("#group_name").val(data.data.group)
                $("#permit_manager").val(data.data.name)
                $("#controller_action").val(data.data.modules)
                $('#myModal').modal('show');
            }
        },
        error:function(msg){
            console.log(msg)
        }
    });

}

//删除权限
function delPermitById(obj){
    var pid = $(obj).data('id');
    $("#pid").val(pid);
    var formData = new FormData();
    formData.append("pid",pid);
    $.ajax({
        url:"delPermitById",
        data:formData,
        type:"post",
        dataType:"json",
        contentType:false,
        processData:false,
        success:function(data){
            alert(data.msg);
            if(data.code == 1){
                self.location.href = "";
            }
        },
        error:function(msg){
            console.log(msg)
        }
    });
}

//表单的提交
function submit_form(){
    var form = $('#permit_info')[0];
    var formData = new FormData(form);

    $.ajax({
        url:"addOrEditPermit",
        data:formData,
        type:"post",
        dataType:'json',
        contentType:false,
        processData:false,
        success:function(returnData){
            $('#myModal').modal('hide');
            alert(returnData.msg);
            if(returnData.code == 1){
                self.location.href = "";
            }
        },
        error:function(msg){
            console.log(msg);
        }
    });
}
