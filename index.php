<?php

// check if logged in

require_once('php-login-advanced/config/config.php');
require_once('php-login-advanced/translations/en.php');
require_once('php-login-advanced/libraries/PHPMailer.php');
require_once('php-login-advanced/classes/Login.php');
$login = new Login();

if ($login->isUserLoggedIn() != true) {
   header( 'Location: http://localhost/reMARK/php-login-advanced/index.php' );
	exit;
}

// logged in, proceed

define ('DS', DIRECTORY_SEPARATOR);
define ('HOME', dirname(__FILE__));

ini_set ('display_errors', 1);

require_once('utilities/loadfiles.php');

if(isset($_GET['openid'])){
	$remark = new RemarkController("remark", "openlink");
	$remark->trackClick($login->getUserid(),$_GET['openid']);
}
elseif(isset($_POST['remark'])){
	$remark = new RemarkController("remark", "saverequest");
	$remark->bookmark($login->getUserid(),$_POST['url'],$_POST['title']);
}
else{
	$remark = new RemarkController("remark", "unfilteredlist");
	$remark->unfilteredList($login->getUserid());
}
