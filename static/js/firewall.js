    var text = "";
    var INPUT = new Array;
    var OUTPUT = new Array;
    var FORWARD = new Array;
    var INPUT_arr = new Array;
    var OUTPUT_arr = new Array;
    var FORWARD_arr = new Array;
    var INPUT_arr2 = new Array;
    var OUTPUT_arr2 = new Array;
    var FORWARD_arr2 = new Array;
    var rule = "";

    $(window).ready(getlist());

    function getlist(){
            var list = $.ajax({
                url:"/safe/firewall",
                dataType:"JSON",
                async:true,
                type:"PUT",
                success:function(){
                    // console.log(list.responseText.split(';'));
                    text = list.responseText.replace(/\s+/g,"~");
                    var arr = text.split("Chain");
                    INPUT_arr = arr[1].split(";");
                    for(var i=2;i<INPUT_arr.length-2;i++){
                        INPUT_arr2 = INPUT_arr[i].split("~");
                        INPUT = INPUT + `<tr onclick="item('INPUT ${INPUT_arr2[0]}')"><td>${INPUT_arr2[0]}</td><td>${INPUT_arr2[1]}</td><td>${INPUT_arr2[2]}</td><td>${INPUT_arr2[3]}</td><td>${INPUT_arr2[4]}</td><td>${INPUT_arr2[5]}</td><td>${INPUT_arr2[6]}</td><td>${INPUT_arr2[7]}</td></tr>`;
                    }
                    FORWARD_arr = arr[2].split(";");
                    for(var i=2;i<FORWARD_arr.length-2;i++){
                        FORWARD_arr2 = FORWARD_arr[i].split("~");
                        FORWARD = FORWARD + `<tr onclick="item('FORWARD ${INPUT_arr2[0]}')"><td>${FORWARD_arr2[0]}</td><td>${FORWARD_arr2[1]}</td><td>${FORWARD_arr2[2]}</td><td>${FORWARD_arr2[3]}</td><td>${FORWARD_arr2[4]}</td><td>${FORWARD_arr2[5]}</td><td>${FORWARD_arr2[6]}</td><td>${FORWARD_arr2[7]}</td></tr>`;
                    }
                    OUTPUT_arr = arr[3].split(";");
                    for(var i=2;i<OUTPUT_arr.length-2;i++){
                        OUTPUT_arr2 = OUTPUT_arr[i].split("~");
                        OUTPUT = OUTPUT + `<tr onclick="item('OUTPUT ${INPUT_arr2[0]}')"><td>${OUTPUT_arr2[0]}</td><td>${OUTPUT_arr2[1]}</td><td>${OUTPUT_arr2[2]}</td><td>${OUTPUT_arr2[3]}</td><td>${OUTPUT_arr2[4]}</td><td>${OUTPUT_arr2[5]}</td><td>${OUTPUT_arr2[6]}</td><td>${OUTPUT_arr2[7]}</td></tr>`;
                    }
                    $('#INPUT').html(INPUT);
                    $('#FORWARD').html(FORWARD);
                    $('#OUTPUT').html(OUTPUT);
                },
                error:function(){
                    console.log("RuleList get info error!");
                }
            });
        }

    function item(vals){
        rule = vals;
        $('#confirmtext').html("确定删除: "+vals);
        $('#confirm').modal('show');
    }

    function tips(){
        $('#addrules').modal('hide');
        $('#tips').modal('show');
    }

    function addrules(){
        $('#addrules').modal('show');
    }

    function list(){
        getlist();
        INPUT = "";
        OUTPUT = "";
        FORWARD = "";
        $('#list').modal('show');
    }

    function exec(){
        var code = $('#code').val();
        var auto = $("input[type='checkbox']").is(':checked')
        var rule = $.ajax({
            url:"/safe/firewall",
            dataType:"JSON",
            async:true,
            type:"POST",
            data:{"code":code,"auto":auto},
            success:function(){
                $('#infotext').html(rule.responseText);
                $('#info').modal('show');
            },
            error:function(){console.log('IptablesCode get info error!');}
        });
    }

    function getinfo(){
        $('#info').modal('hide');
    }

    function ruledel(){
        var chain = rule.split(" ")[0];
        var num = rule.split(" ")[1];
        var del = $.ajax({
            url:"/safe/firewall?chain="+chain+"&num="+num,
            dataType:"JSON",
            async:true,
            type:"DELETE",
            // data:{"chain":code,"num":auto},
            success:function(){
                // $('#infotext').html(del.responseText);
                $('#infotext').html("删除成功");
                $('#info').modal('show');
            },
            error:function(){console.log('IptablesCode get info error!');}
        });
    }

    function rulesave(){
        var save = $.ajax({
            url:"/safe/firewall",
            dataType:"JSON",
            async:true,
            type:"OPTIONS",
            // data:{"chain":code,"num":auto},
            success:function(){
                // $('#infotext').html(del.responseText);
                $('#infotext').html("保存生效成功");
                $('#info').modal('show');
            },
            error:function(){console.log('IptablesCode get info error!');}
        });
    }