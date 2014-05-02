


function activate(id){
    document.getElementById('show_'+id).className = "inactive";
    document.getElementById(id).className = "active";
    document.getElementById(id).focus()
}

function deactivate(elem){
    elem.className = "inactive";
}

function printList(){
    var output = "";
    var savebuttonclass = "inactive";
    for (var i = 0; i < todos.length; i++) {
        
        output+= '<div class="status ';
        if(i>=archiv.length || todos[i]['text']!=archiv[i]['text'] || todos[i]['date']!=archiv[i]['date'] || todos[i]['type']!=archiv[i]['type'] ){
            output+= "modified";
            savebuttonclass = "active";
        }
        
        output+= '">&nbsp;</div>';
        output+= printLine(i,todos[i]['text'],todos[i]['date'],todos[i]['type']);
    }
    if(todos.length==0){
        output+= '<a href="javascript:addFirstLine()">start List</a> ';
    }
    if(archiv.length>todos.length){
        savebuttonclass = "active";
    }
    document.getElementById("data").value=JSON.stringify(todos).replace("&nbsp;","");
    document.getElementById("container").innerHTML =output;
    document.getElementById("save").className = savebuttonclass;
}

function printLine(nr,text,date,type){
    var output = "";
    
    output+= '<div class="buttons">';
    
    output+= '<a href="javascript:moveUp('+nr+')">';
    output+= '<img class="arrow up" src="img/arrow_20x20.png">';
    output+= '</a> ';

    output+= '<a href="javascript:moveDown('+nr+')">';
    output+= '<img class="arrow down" src="img/arrow_20x20.png">';
    output+= '</a> ';
    
    if(type==0){
        output+= '<a href="javascript:moveLeft('+nr+')">';
    output+= '<img class="arrow left" src="img/arrow_20x20.png">';
    }
    else{
        output+= '<a href="javascript:moveRight('+nr+')">';
    output+= '<img class="arrow right" src="img/arrow_20x20.png">';
    }
    
    output+= '</a> ';
    
    
    output+= '<a href="javascript:removeLine('+nr+')">';
    output+= '<img class="buttonImage" src="img/delete_20x20.png">';
    output+= '</a> ';
    
    output+= '</div>';
    output+= '<div class="headline">';
    if(type!=0){
        output+= '<input class="inactive" type="text" onblur="writeData('+nr+',this.value,\'text\')" value="'+text+'" id="text_'+nr+'" name="text_'+nr+'">';
        output+= '<div class="content" id="show_text_'+nr+'" onclick="activate(\'text_'+nr+'\')"><b>'+text+'</b></div>';
      
    }
    else{
        output+= '&nbsp;';
    }
    output+= '</div>';
    output+= '<div class="item">';
    if(type==0){
        output+= '<input class="inactive" type="text" onblur="writeData('+nr+',this.value,\'text\')" value="'+text+'" id="text_'+nr+'" name="text_'+nr+'">';
        output+= '<div class="content" id="show_text_'+nr+'" onclick="activate(\'text_'+nr+'\')">';
        if(text==""){
            output+= "&nbsp;"
        }
        else{
            output+= text;
        }
        output+='</div>';
      
    }
    else{
        output+= '&nbsp;';
    }
    output+= '</div>';
    output+= '<div class="date">';
    if(type==0){
        output+= '<input class="inactive" type="text" onblur="writeData('+nr+',this.value,\'date\')" value="'+date+'" id="date_'+nr+'" name="date_'+nr+'">';
        output+= '<div class="content" id="show_date_'+nr+'" onclick="activate(\'date_'+nr+'\')">';
        if(date==""){
            output+= "&nbsp;"
        }
        else{
            output+= date;
        }
        output+='</div>';
      
    }
    else{
        output+= '&nbsp;';
    }
    output+= '</div>';
    output+= '<div class="spacer" id="spacer_'+nr+'" onmouseover="showAddButton('+nr+')" onmouseout="hideAddButton('+nr+')">';
    output+= '<a class="addButton inactive" id="between_'+nr+'" href="javascript:addLine('+nr+')">';
    output+= '<div class="addContainer">';
    output+= '&nbsp;</div></a>';
    output+= '</div>';
    
    return output;
}

function writeData(nr,value,field){
    todos[nr][field]=value;
    printList();
}

function showAddButton(nr){
    document.getElementById('spacer_'+nr).style = "height:20px";
    document.getElementById('between_'+nr).className = "active";
}


function hideAddButton(nr){
    document.getElementById('spacer_'+nr).style = "height:5px";
    document.getElementById('between_'+nr).className = "inactive";
}

function switchRows(nr1,nr2){
    var newTodos = new Array();
    for (var i = 0; i < todos.length; i++) {
        if(i==nr1){
            continue;
        }
        if(i==nr2){
            newTodos[i-1] = todos[i];
            newTodos[i] = todos[i-1];
            continue;
        }
        newTodos[i] = todos[i];
    }
    //todos = new Array();
    todos = newTodos;
    printList();
}

function moveUp(nr){
    switchRows(nr-1,nr);
}

function moveDown(nr){
    switchRows(nr,nr+1);
}

function addLine(nr){
    //alert(nr);
    var newTodos = new Array();
    for (var i = 0; i <= nr; i++) {
        newTodos[i] = todos[i];
    }
    newTodos[nr+1] =  {"text": "", "date": "", "type": 0};
    for (var i = nr+1; i < todos.length; i++) {
        newTodos[i+1] = todos[i];
    }
    //todos = new Array();
    todos = newTodos;
    printList();
}

function removeLine(nr){
    //alert(nr);
    var newTodos = new Array();
    for (var i = 0; i < nr; i++) {
        newTodos[i] = todos[i];
    }
    for (var i = nr+1; i < todos.length; i++) {
        newTodos[i-1] = todos[i];
    }
    //todos = new Array();
    todos = newTodos;
    printList();
}

function moveRight(nr){
    todos[nr]["type"]=0;
    printList();
}

function moveLeft(nr){
    todos[nr]["type"]=1;
    printList();
}

function addFirstLine(){
    todos[0] =  {"text": "start here", "date": "", "type": 0};
    printList();
}