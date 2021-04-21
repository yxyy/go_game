/**
 * Created by Administrator on 2017/2/22.
 */
function isAllPackage(obj){
    var channel = $(obj).data("channel");
    var this_val = $(obj).data("is_all");
    if(this_val == 0){
        $("."+channel).prop("checked",true);
        $(obj).data("is_all",1);
    }else{
        $("."+channel).prop("checked",false);
        $(obj).data("is_all",0);
    }
}

function isSelect(obj){
    if($(obj).attr("checked") != true){
        var channel_checkbox = $("#"+$(obj).data('channel'));
        channel_checkbox.prop("checked",false);
        channel_checkbox.data("is_all",0);
    }
}

//修改权限组信息
function modifyUserGroup(){
    var formData = new FormData($("#user_group_info")[0]);
    $.ajax({
        url:'/backend/user_group/modifyUserGroup',
        data:formData,
        type:"post",
        dataType:"json",
        contentType:false,
        processData:false,
        success:function(data){
            alert(data.msg);
        },
        error:function(msg){
            console.log(msg);
        }
    });
}

function isSelectAll(type,obj,Superfluous){
    if($(obj).data('val') == 0){
        $("input[type=checkbox][name='"+type+"\[\]']").prop("checked",true);
        $("input[type=checkbox][name='"+Superfluous+"\[\]']").prop("checked",true);
        $(obj).data('val',1);
        $(obj).html("取消全选");
    }else if($(obj).data('val') == 1){
        $("input[type=checkbox][name='"+type+"\[\]']").prop("checked",false);
        $("input[type=checkbox][name='"+Superfluous+"\[\]']").prop("checked",false);
        $(obj).data('val',0);
        $(obj).html("全选");
    }
}