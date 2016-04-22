<?php

header("Cache-Control: public, max-age=0, no-cache");
header('Cache-Control: no-store, no-cache, must-revalidate');
header('Cache-Control: post-check=0, pre-check=0', false);
header('Pragma: no-cache');

error_reporting(E_ALL | E_STRICT);
ini_set('display_errors', 1);

require __DIR__ . '/../vendor/autoload.php';

$twitter = new Itsmethemojo\Authentification\TwitterExtended();
$loginUser = $twitter->getLoginUser();

if ($twitter->isLoggedIn()) {
    $url = "index.html";
    header("Location: " . $url);
    echo "redirecting<script>location.href=location.href</script>redirecting to " . $url;
    exit;
}
?>
ohoh