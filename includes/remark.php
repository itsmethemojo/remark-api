<?php

$remark = new RemarkController("remark");

if(isset($_GET['openid'])){
    $remark->trackClick($userid,$_GET['openid']);
}
elseif(isset($_POST['remark'])){
    $remark->bookmark($userid,$_POST['url'],$_POST['title']);
}
else{
    $remark->unfilteredList($userid);
}


?>

