    var name;
    $(window).ready(function(){
        getUserList();
        getGroupList();
    });

    function getUserList(){
        var userlists = $.ajax({
            url:"/safe/user",
            dataType:"JSON",
            type:"GET",
            async:true,
            success:function(){
                var s = userlists.responseText.replace('\"',"");
                var arr = s.split("\\n");
                var text = "";
                for (var i=0;i<arr.length-1;i++){
                    text = text + `<div class="column ui button user-item" style="width:30%" onclick="userdel('${arr[i]}')">${arr[i]}</div>`;
                }
                $('#usertext').html(`<div class="ui grid" style="text-align:center"><div class="four column">${text}</div></div>`);
            },
            error:function(){
                console.log("UserList get info error!");
            }
        });
    }

    function getGroupList(){
        var grouplists = $.ajax({
            url:"/safe/group",
            dataType:"JSON",
            type:"GET",
            async:true,
            success:function(){
                var s = grouplists.responseText.replace('\"',"");
                var arr = s.split("\\n");
                var text = "";
                for (var i=0;i<arr.length-1;i++){
                    text = text + `<div class="column ui button group-item" style="width:30%" onclick="groupdel('${arr[i]}')">${arr[i]}</div>`;
                }
                $('#grouptext').html(`<div class="ui grid" style="text-align:center"><div class="four column">${text}</div></div>`);
            },
            error:function(){
                console.log("GroupList get info error!");
            }
        });
    }

    function userlist(){
        getUserList();
        $('#userlist').modal('show');
    }

    function useradd(){
        $('#useradd').modal('show');
    }

    function usermod(){
        $('#usermod').modal('show');
    }

    function grouplist(){
        getGroupList();
        $('#grouplist').modal('show');
    }

    function groupadd(){
        $('#groupadd').modal('show');
    }

    function userdel(temp){
        name = temp;
        $('#userdeltext').html("确定删除用户: " + name)
        $('#userdel').modal('show');
    }

    function groupdel(temp){
        name = temp;
        $('#groupdeltext').html("确定删除组: " + name)        
        $('#groupdel').modal('show');
    }
    
    function userdel_do(){
        var del = $.ajax({
            url:"/safe/user?auth="+name,
            dataType:"JSON",
            type:"DELETE",
            async:true,
            success:function(){
                if (del.responseText.indexOf("del")<0){
                    $('#infotext').html("删除成功");
                }else{
                    $('#infotext').html(del.responseText);
                }
                $('#info').modal('show');
            },
            error:function(){
                console.log("UserDelete get info error!");
            }
        });
    }

    function useradd_do(){
        var auth = $('#auth').val();
        $('#auth').val("");
        var passwd = $('#passwd').val();
        $('#passwd').val("");
        var add = $.ajax({
            url:"/safe/user",
            dataType:"JSON",
            type:"POST",
            data:{"auth":auth,"passwd":passwd},
            async:true,
            success:function(){
                if (add.responseText.indexOf("add")<0){
                    $('#infotext').html("添加成功");
                }else{
                    $('#infotext').html(add.responseText);
                }
                $('#info').modal('show');
            },
            error:function(){
                console.log("Useradd get info error!");
            }
        });
    }

    function usermod_do(){
        var auth = $('#authmod').val();
        $('#authmod').val("");
        var group = $('#groupmod').val();
        $('#groupmod').val("");
        var mod = $.ajax({
            url:"/safe/user?auth="+auth+"&group="+group,
            dataType:"JSON",
            type:"OPTIONS",
            async:true,
            success:function(){
                if (mod.responseText.indexOf("mod")<0){
                    $('#infotext').html("转移成功");
                }else{
                    $('#infotext').html(mod.responseText);
                }
                $('#info').modal('show');
            },
            error:function(){
                console.log("Usermod get info error!");
            }
        });
    }

    function groupadd_do(){
        var group = $('#group').val();
        $('#group').val("");
        var add = $.ajax({
            url:"/safe/group",
            dataType:"JSON",
            type:"POST",
            data:{"group":group},
            async:true,
            success:function(){
                if (add.responseText.indexOf("add")<0){
                    $('#infotext').html("添加成功");
                }else{
                    $('#infotext').html(add.responseText);
                }
                $('#info').modal('show');
            },
            error:function(){
                console.log("Groupadd get info error!");
            }
        });
    }

    function groupdel_do(){
        var del = $.ajax({
            url:"/safe/group?group="+name,
            dataType:"JSON",
            type:"DELETE",
            async:true,
            success:function(){
                if (del.responseText.indexOf("del")<0){
                    $('#infotext').html("删除成功");
                }else{
                    $('#infotext').html(del.responseText);
                }
                $('#info').modal('show');
            },
            error:function(){
                console.log("GroupDelete get info error!");
            }
        });
    }

    function getinfo(){
        $('#info').modal('hide');
    }
