var arr = text.split("\n");
var table = "";
var list = "";
var lists = new Array();
for(var i=1; i<arr.length-1; i++){
    var temparr = arr[i].split(/\s+/g);
    for(var j=0; j<8; j++){
        if(j==1){
            list = list + "<td class='pid'>" + temparr[j] + "</td>"; 
        }else{
            list = list + "<td>" + temparr[j] + "</td>";             
        }
    }
    lists[i] = `<tr class='list'>"+${list}+"</tr>`;
    list = "";
}

$("#process").html(lists);
var pid = "";
$(".list").click(function(){
    pid = $(this).children(".pid").html();
    $('#tips').modal('show');
    $('#tipsinfo').html("将关闭进程pid："+$(this).children(".pid").html());
});

function killprocess(){
    var kill = $.ajax({
        url:"/device/ram",
        type:"POST",
        dataType:"JSON",
        async:true,
        data:{"pid":pid},
        success:function(){
            $('#infotext').html('关闭进程成功');
        },
        error:function(){console.log("Killprocess get info error!")}
    });
    $('#info').modal('show');
}

function getinfo(){
    $('#info').modal('hide');
    location.href = "/device/ram";    
}
