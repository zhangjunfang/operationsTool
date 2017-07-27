    function reboot(){
        var RebootInfo = $.ajax({
        url:"/resuce/reboot",
        type:"PUT",
        dataType:"JSON",
        async:true,
        success:function(){
            $('#infotext').html("正在重启所有功能将会暂停运作...")
        },
        error:function(){console.log("Reboot get info error!")}
    });
    $('#info').modal('show');
}
    function getinfo(){
        $('#info').modal('hide');
        location.href = "/";
    }