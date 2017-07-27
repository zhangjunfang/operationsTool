
    $(window).ready(
        function(){
            // setInterval(function(){},1000)
            var cpu = $.ajax({
                url:"/device/cpu",
                async:true,
                dataType:"JSON",
                type:"PUT",
                success:function(){
                    var cpu_obj = $.parseJSON(cpu.responseText);
                    $('#cpu_name').html("型号:"+cpu_obj.cpu_name);
                    $('#cpu_processor').html("线程:"+cpu_obj.cpu_processor);
                    $('#cpu_MHz').html("频率:"+cpu_obj.cpu_MHz);
                    $('#cpu_cores').html("核心数:"+cpu_obj.cpu_cores);
                    $('#cpu_pre').html("占用比:"+cpu_obj.cpu_pre);
                },
                error:function(){console.log("Cpu get info error!")}
            });
            var system = $.ajax({
                url:"/soft/system",
                async:true,
                dataType:"JSON",
                type:"PUT",
                success:function(){
                    var system_obj = $.parseJSON(system.responseText);
                    $('#system_info').html(system_obj.system_info);
                },
                error:function(){console.log("Cpu get info error!")}
            });
            var ram = $.ajax({
                url:"/device/ram",
                async:true,
                dataType:"JSON",
                type:"PUT",
                success:function(){
                    var ram_obj = $.parseJSON(ram.responseText);
                    $('#ram_total').html("总内存:"+ram_obj.ram_total);
                    $('#ram_free').html("剩余内存:"+ram_obj.ram_free);
                    $('#ram_pre').html("占用比:"+ram_obj.ram_pre);
                },
                error:function(){console.log("Ram get info error!")}
            });
            var rom = $.ajax({
                url:"/device/rom",
                async:true,
                dataType:"JSON",
                type:"PUT",
                success:function(){
                    var rom_obj = $.parseJSON(rom.responseText);
                    var rom_content = "";
                    for(var i in rom_obj){
                        var temparr = rom_obj[i].split(",");
                        rom_content = rom_content + "<tr><td>"+temparr[5]+"</td><td>"+temparr[1]+"</td><td>"+temparr[2]+"</td><td>"+temparr[3]+"</td><td>"+temparr[4]+"</td></tr>";
                    }
                    $('#rom_content').html(
                        rom_content
                    );
                },
                error:function(){console.log("Rom get info error!")}
            });
        }
    );

