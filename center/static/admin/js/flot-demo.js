var oilprices = [];
function search_online(){
    var stime = $("#stime").val();

    var server_arr = new Array();
    $('input[name="s_server"]:checked').each(function () {
        server_arr.push($(this).val());//向数组中添加元素
    });
    var server_ids = server_arr.join(',');//将数组元素连接起来以构建一个字符串

    var s_channel_arr = new Array();
    $('input[name="s_channel"]:checked').each(function () {
        s_channel_arr.push($(this).val());//向数组中添加元素
    });
    var s_channel = s_channel_arr.join("','");//将数组元素连接起来以构建一个字符串

    var s_packager_arr = new Array();
    $('input[name="s_package"]:checked').each(function () {
        s_packager_arr.push($(this).val());//向数组中添加元素
    });
    var s_package = s_packager_arr.join(',');//将数组元素连接起来以构建一个字符串

    $.ajax({
        url:"online5Minute",
        data:{"stime":stime,"server":server_ids,"channel":s_channel,"package":s_package},
        type:'POST',
        success:function(data){
            // console.log(1111111);
            var table_dom = $("#online_hour_table");
            if(data.code == 1){
                var myDate = new Date();
                var hourOnline = [];
                $.each(data.data,function(key,val){
                    oilprices.push([Number(val.create_time),Number(val.online_num)]);
                    if(val.create_time%1000 == 0){
                        myDate.setTime(val.create_time);
                        hourOnline.push([myDate.toLocaleTimeString(),val.online_num]);
                    }
                });
                table_dom.html("");
                table_dom.append('<tr><th>时间</th><th>人数</th></tr>');
                $.each(hourOnline,function(k,v){
                    table_dom.append('<tr><th>'+v[0]+'</th><th>'+v[1]+'</th></tr>');
                });
            }else{
                table_dom.html("");
            }

            // console.log(oilprices);

            doPlot("right");
            oilprices = [];
        }
    });
}

search_online();

function euroFormatter(v, axis) {
    return "&yen;"+v.toFixed(axis.tickDecimals);
}

function doPlot(position) {
    $.plot($("#flot-line-chart-multi"), [{
        data: oilprices,
        label: "实时在线 (人数)",
        lines: { show: true, fill: true }
    }], {
        xaxes: [{
            mode: 'time'
        }],
        yaxes: [{
            min: 0
        }, {
            // align if we are to the right
            alignTicksWithAxis: position == "right" ? 1 : null,
            position: position,
            tickFormatter: euroFormatter
        }],
        legend: {
            position: 'sw'
        },
        colors: ["#1ab394"],
        grid: {
            color: "#999999",
            hoverable: true,
            clickable: true,
            tickColor: "#D4D4D4",
            borderWidth:0,
        },
        tooltip: true,
        tooltipOpts: {
            content: "%s %x 为 %y",
            xDateFormat: "%H:%M",
        }

    });
}




