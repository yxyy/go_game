function ban(role_id, account, role_name, ip, server) {
    $("#role_id").val(role_id);
    $("#account").val(account);
    $("#role_name").val(role_name);
    $("#ip").val(ip);
    $("#server").val(server);
    $("#banModal").modal("show");

}
function ban_submit() {
    var formData = new FormData($('#banForm')[0]);
    $.ajax({
        url: '/admin/ban/ban',
        data: formData,
        type: "POST",
        dataType: "json",
        contentType: false,
        processData: false,
        success: function (data) {
            if (data.status == 1) {
                alert(data.msg);
                $('#banModal').modal('hide');
                window.location.replace(data.url);
            } else {
                alert(data.msg);
            }
        }
    });
}