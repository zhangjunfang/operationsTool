
    function getinfo(){
        $('#info').modal('hide');
    }

    function allow(){
        var con = $.ajax({
        url:"/safe/stopping",
        type:"PUT",
        dataType:"JSON",
        async:true,
        success:function(){
            // alert(con.responseText);
            if(con.responseText.indexOf("0")>=0){
                $('#infotext').html("Ping已允许");
            }else{
                $('#infotext').html("启动失败");                
            }
        },
        error:function(){console.log("Allowping get info error!")}
    });
        $('#info').modal('show');
    }

    function stop(){
        var con = $.ajax({
        url:"/safe/stopping",
        type:"DELETE",
        dataType:"JSON",
        async:true,
        success:function(){
            // alert(con.responseText);
            if(con.responseText.indexOf("1")>=0){
                $('#infotext').html("Ping已禁止");
            }else{
                $('#infotext').html("禁止失败");                
            }
        },
        error:function(){console.log("Stopping get info error!")}
    });
        $('#info').modal('show');
    }
