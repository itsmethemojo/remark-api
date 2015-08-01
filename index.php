<?php

if (file_exists('vendor/loginService.php')) {
    include 'vendor/loginService.php';
}

include 'vendor/mvc-core/autoloader.php';
include 'autoloader.php';


$remark = $remarkController;

if(!isset($_POST["action"]) && !isset($_GET["action"])){
    $remark->actionComplete();
    exit();
}

if(isset($_GET["action"])){
    switch ($_GET["action"]){
        case "open" :
            $remark->actionOpen();
            exit();
    }
}

if(isset($_POST["action"])){
    switch ($_POST["action"]){
        case "edit" :
            $remark->actionEdit();
            exit();
        case "remark" :
            $remark->actionRemark();
            exit();
    }
    
}




echo "action not implemented";