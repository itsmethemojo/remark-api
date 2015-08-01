var buttons = require('sdk/ui/button/action');
var tabs = require("sdk/tabs");
var Request = require("sdk/request").Request;

var button = buttons.ActionButton({
  id: "remark-button",
  label: "reMARK this Page",
  icon: {
    "16": "./remark-icon-16.png",
    "32": "./remark-icon-32.png",
    "64": "./remark-icon-64.png"
  },
  onClick: handleClick
});

function trim(s){ 
  return ( s || '' ).replace( /^[\s\n]+|[\s\n]+$/g, '' ); 
}

function handleClick(state) {

    var remarkUrl = "https://itsmethemojo.com/reMARK/index.php";
    var currentUrl = tabs.activeTab.url;
    var currentTitle = tabs.activeTab.title;    
    
    remarkRequest = Request({
        url: remarkUrl,
        content: {
            action : "remark",
            url : currentUrl,
            title : currentTitle
            
        },
        onComplete: function (response) {
            tabs.open(remarkUrl);
            //TODO add information if bookmark store was succesfull
            //if (!response || response.status===0 || trim(response.text)!=="1" || trim(response.text)!=="0") {
            
                
            //} else {
                /*
                console.log("HTTP HEADERS:");
                for (var headerName in response.headers) {
                    console.log(headerName + " : " + response.headers[headerName]);
                }
                console.log("------------");
                console.log("HTTP STATUS: "+response.status);
                console.log("------------");
                console.log("HTTP RESPONSE:");
                console.log(response.text);
                */
            //}
        }
    });
    
    remarkRequest.post();

}