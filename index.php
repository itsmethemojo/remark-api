<?php

// check if logged in

require_once('php-login/application/libs/Session.php');

Session::init();

if (!Session::get('user_logged_in')) {
   header( 'Location: http://localhost/reMARK/php-login/index.php' );
	exit;
}
// logged in, proceed
$userid = Session::get('user_id');

define ('DS', DIRECTORY_SEPARATOR);
define ('HOME', dirname(__FILE__));

ini_set ('display_errors', 1);

require_once('utilities/loadfiles.php');

if(!isset($_GET['app'])){
    $app = "remark"; 
}
else{
    $app = $_GET['app'];
}

if (file_exists('includes/'.$app.'.php')) {
    require 'includes/'.$app.'.php';
}
else{
    echo 'no such application';
}