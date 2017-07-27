var arr1 = text.split("#");
// console.log(arr1);

var arr2 = arr1[0].split("\n");/*所有硬盘分区*/
var arr3 = arr1[1].split("\n");/*已挂载硬盘分区*/
// console.log(arr2);
// console.log(arr3);
var mount = "";
var umount = "";
var d = 0;
var dev = "";
var local ="";
for(var i=0;i<arr2.length-1;i++){
    for(var j=0;j<arr3.length;j++){
        if(arr3[j].indexOf(arr2[i].split("\ ")[1])>=0){
            d=1;
            break;
            // mount = mount + `<tr><td>${arr3[j].split("\ ")[0]}</td><td>${arr3[j].split("\ ")[1]}</td><td>已挂载</td></tr>`;
        }
    }
    if(d==1){
        mount = mount + `<tr class="mount"><td class="dev">${arr3[j].split("\ ")[0]}</td><td class="local">${arr3[j].split("\ ")[1]}</td><td>已挂载</td></tr>`;
        d=0;
    }else{
        umount = umount + `<tr class="umount"><td class="dev">${arr2[i].split("\ ")[1]}</td><td class="local">${arr2[i].split("\ ")[0]}</td><td>未挂载</td></tr>`;
    }
}
$('#umount').html(umount);
$('#mount').html(mount);
$('.mount').click(function(){
    $('.modal').html(`<div class="header">确定卸载?</div><div class="content"><div class="actions"><div class="ui negative button">No </div><div class="ui positive right labeled icon button" onclick="romumount()">Yes <i class="checkmark icon"></i> </div></div>`);    
    $('.small.modal').modal('show');
    dev = $(this).children('.dev').html();
    local = $(this).children('.local').html();    
});
$('.umount').click(function(){
    $('.modal').html(`<div class="header">确定挂载?</div><div class="content"><div class="ui input focus"><input id="info" type="text" placeholder="挂载点:/mnt/disk"></div></div><div class="actions"><div class="ui negative button">No </div><div class="ui positive right labeled icon button" onclick="rommount()">Yes <i class="checkmark icon"></i> </div></div>`);
    $('.small.modal').modal('show');
    dev = $(this).children('.dev').html();
});

function rommount(){
    var info = $('#info').val();
    if(info == ""){
        alert("挂载点不能为空");
    }else{
        var RomMount = $.ajax({
            url:"/device/rom",
            type:"POST",
            dataType:"JSON",
            async:true,
            data:{"dev":dev,"info":info},
            success:function(){
                var temp = RomMount.responseText;   
                var return_info = $.parseJSON(temp);
                // alert(return_info.info);     
                if(return_info.info == "success"){
                    $('.modal').html(`<div class="header">系统消息</div><div class="content"><div class="ui input focus">操作成功</div></div><div class="actions"><div class="ui positive right labeled icon button" onclick="javascript:location.href = '/device/rom';">Yes <i class="checkmark icon"></i> </div></div>`);
                    $('.small.modal').modal('show');
                }else{
                    $('.modal').html(`<div class="header">系统消息</div><div class="content"><div class="ui input focus">操作失败:${return_info.info}</div></div><div class="actions"><div class="ui positive right labeled icon button" onclick="javascript:location.href = '/device/rom';">Yes <i class="checkmark icon"></i> </div></div>`);
                    $('.small.modal').modal('show');
                }
            },
            error:function(){console.log("RomMount get info error!")}
        });
    }
}

function romumount(){
    var RomMount = $.ajax({
        url:"/device/rom",
        type:"POST",
        dataType:"JSON",
        async:true,
        data:{"local":local},
        success:function(){
            var temp = RomMount.responseText;   
            var return_info = $.parseJSON(temp);
            // alert(return_info.info);     
            if(return_info.info == "success"){
                $('.modal').html(`<div class="header">系统消息</div><div class="content"><div class="ui input focus">操作成功</div></div><div class="actions"><div class="ui positive right labeled icon button" onclick="javascript:location.href = '/device/rom';">Yes <i class="checkmark icon"></i> </div></div>`);
                $('.small.modal').modal('show');
            }else{
                $('.modal').html(`<div class="header">系统消息</div><div class="content"><div class="ui input focus">操作失败:${return_info.info}</div></div><div class="actions"><div class="ui positive right labeled icon button" onclick="javascript:location.href = '/device/rom';">Yes <i class="checkmark icon"></i> </div></div>`);
                $('.small.modal').modal('show');
            }
        },
        error:function(){console.log("RomMount get info error!")}
    });
}