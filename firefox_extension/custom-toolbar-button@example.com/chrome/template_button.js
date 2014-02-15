CustomButton = { 

1: function () { 
var url = "http://localhost/reMARK/index.php";
var username = "YOURNAME";
var password = "YOURPASS";

xmlhttp=new XMLHttpRequest();
xmlhttp.open("POST",url,false);
xmlhttp.setRequestHeader("Content-type","application/x-www-form-urlencoded");
xmlhttp.send("remark=1&login=1&user_name="+username+"&user_password="+password+"&url="+encodeURIComponent(window.content.location.href)+"&title="+window.content.document.title);
//alert("username="+username+"&password="+password+"&url="+encodeURIComponent(window.content.location.href)+"&title="+encodeURIComponent(window.content.document.title));
var res = xmlhttp.responseText;
// just responses if something wasn't ok
if(res!="0" && res!="1") 
	alert(res);
  }, 

}
