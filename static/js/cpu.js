var Text = $('#cpucores').html();
for(var i=1;i<CpuCore;i++){
    Text = Text + `<div class="card"><div class="content"><div class="header">CpuCore${i}</div><div id="status${i}" class="meta"></div></div><div class="extra content"><div class="ui two buttons"><div class="ui basic green button">开启</div><div class="ui basic red button">关闭</div></div></div></div>`;
}
$('#cpucores').html(Text);

var CpuCoreOptions = $.ajax({
    url:"/device/cpu",
    type:"OPTIONS",
    dataType:"JSON",
    async:true,
    success:function(){
        var CpuCoreStatusList = CpuCoreOptions.responseText;
        var length = CpuCoreStatusList.length;
        var arr = CpuCoreStatusList.split("")
        for(var i = 1; i<length-1; i++){
            if(arr[i]=="0"){
                $('#status'+i).html("目前状态:关闭")
            }else if(arr[i]=="1"){
                $('#status'+i).html("目前状态:开启")
            }else{
                $('#status'+i).html("目前状态:异常")
            }
        }
    },
    error:function(){console.log("CpuCoreOptions get info error!")}
});