//TODO replace with jquery lib?
function gup( name )
{
  name = name.replace(/[\[]/,"\\\[").replace(/[\]]/,"\\\]");
  var regexS = "[\\?&]"+name+"=([^&#]*)";
  var regex = new RegExp( regexS );
  var results = regex.exec( window.location.href );
  if( results === null )
    return null;
  else
    return results[1];
}

function initializeFiler(){
    $(".filter .box").val(gup("s"));
}

function changeFiler(searchString){
    if(searchString===""){
        history.replaceState(null, null, "?");
    }
    else{
        history.replaceState(null, null, "?s="+searchString);
    }
}

function filterList(searchString){
    changeFiler(searchString);
    printBookmarkList();
}

function retrieveFilter(){
    var filter = new Array();
    filter['string'] = gup("s");
    return filter;
}

function isNotInFilter(bookmark,filter){
    if(!filter["string"]){
        return false;
    }
    
    if(filter["string"]===""){
        return false;    
    }
    
    if(bookmark['title'].toLowerCase().search(filter["string"].toLowerCase()) > -1){
        return false;
    }
        
    if(bookmark['url'].toLowerCase().search(filter["string"].toLowerCase()) > -1){
        return false;
    }
    
    return true;
}



function printBookmarkList(){
    var filter = retrieveFilter();
    lastId = -1;
    lastDate = "";
    markup = '<table class="bookmarks">';

    for (var i = 0; i < bookmarks.length; i++) {
        
        if(isNotInFilter(bookmarks[i],filter)){
            continue;
        }
        
        markup += printBookmarkListLine(i,lastId,lastDate,filter);
        lastId = bookmarks[i]['id'];
        
        lastDate = extractDateString(bookmarks[i]['created']);
    }

    markup += '</table>';

    //TODO switch to jquery notation
    document.getElementById("bookmarksContainer").innerHTML = markup;

}

function printBookmarkListLine(bookmarkPosition,lastId,lastDate,filter){
    
    bookmark = bookmarks[bookmarkPosition];
    
    if(bookmark['id'] === lastId){
        return "";
    }    
    var dateString = extractDateString(bookmark['created']);
    if(dateString === lastDate){
        extraLine = '';
    }else{
        extraLine = '<tr><td class="date date' + dateString + '">' + dateString + '</td></tr>';
    }
    
    if(bookmark['customtitle']===""){
        displayedTitle = bookmark['title'];
        
    }
    else{
        displayedTitle = bookmark['customtitle'];
    }
    
    remarkTitle = bookmark['bookmarkcount']=="0" ? '' : ' title="'+bookmark['bookmarkcount']+'"';
    clickTitle = bookmark['clickcount']=="0" ? '' : ' title="'+bookmark['clickcount']+'"';
    return extraLine + '<tr class="line position'+bookmarkPosition+'">' +
            '<td class="date"></td>' +
            '<td class="time">' + extractTimeString(bookmark['created']) + '</td>' +
            '<td><img class="icon-padding visibility-level-' + getRemarkVisibility(bookmark['bookmarkcount']) + '" src="img/remark-' + getIconVersion(bookmark['bookmarkcount']) + '.png"'+remarkTitle+'></td>' +
            '<td><img class="icon-padding visibility-level-' + getClickCountVisibility(bookmark['clickcount']) + '" src="img/click-' + getIconVersion(bookmark['clickcount']) + '.png"'+clickTitle+'></td>' +
            '<td><a href="javascript:openDetails(' + bookmarkPosition + ')"><img class="icon-padding" src="img/edit.png"></a></td>' +
            '<td><a href="index.php?action=open&id=' + bookmark['id'] +'">' + 
            displayedTitle + 
            '</a>' + 
            '<br><span class="domain">' + bookmark['domain'] + '</span>'
            '</td>' +
            '</tr>';
}

function getRemarkVisibility(number) {
    switch (parseInt(number)) {
        case 0: return 0;
        case 1: return 1;
        case 2: return 2;
        case 3: return 4;
        case 4: return 6;
    }
    return 8;
}

function getClickCountVisibility(number) {
    switch (parseInt(number)) {
        case 0: return 0;
        case 1: return 1;
        case 2: return 2;
        case 3: return 3;
    }
    if (number <= 6)
        return 4;
    if (number <= 10)
        return 5;
    if (number <= 15)
        return 6;
    if (number <= 20)
        return 7;

    return 8;
}

function getIconVersion(number){
    if(parseInt(number) === 0){
       return 0;
    }
    return 1;
}

function extractDateString(timestamp){
    date = new Date(timestamp * 1000);
    return date.getFullYear() + "-" + printTwoDigits(date.getMonth() + 1) + "-" + printTwoDigits(date.getDate());
}

function extractTimeString(timestamp){
    date = new Date(timestamp * 1000);
    return printTwoDigits(date.getHours()) + ":" + printTwoDigits(date.getMinutes());
    
}

//TODO make this better
function printTwoDigits(number){
    if(number<10){
        return "0"+number;
    }
    return number;
}

function openDetails(bookmarkPosition){
    
    $( ".line.details" ).remove();
    
    detailsHtml = '<tr class="line details"><td colspan="5"></td><td>';
    detailsHtml += '<input class="id" type="hidden" name="customtitle" value="'+bookmarks[bookmarkPosition]['id']+'"/>';
    detailsHtml += '<input class="customtitle" type="text" name="customtitle" value="'+bookmarks[bookmarkPosition]['customtitle']+'"/><br>';
    detailsHtml += '<span class="orgtitle">'+bookmarks[bookmarkPosition]['title']+'</span><br>';
    detailsHtml += '<input type="button" onclick="discardDetails();" value="x">';
    detailsHtml += '<input type="button" onclick="saveDetails();" value="save">';
    detailsHtml += '</td></tr>';
    
    $(".position"+bookmarkPosition).after(detailsHtml);
    
}

function discardDetails(){
    $( ".line.details" ).remove();
}

function saveDetails(){    
    var customTitle = $(".line.details .customtitle").val();
    var id = $(".line.details .id").val();
    $.post( "index.php", {
        id: id,
        customtitle: customTitle, 
        action: "edit" 
    }).done(function( responseData ) {
        location.href = location.href;
        //alert( "Data Loaded: " + responseData );
    });
}

/*
 * TODO implement set
function 
for (let item of mySet.values()) console.log(item);
*/