<?php
header("Cache-Control: public, max-age=0, no-cache");
header('Cache-Control: no-store, no-cache, must-revalidate');
header('Cache-Control: post-check=0, pre-check=0', false);
header('Pragma: no-cache');


require __DIR__ . '/../vendor/autoload.php';

use Itsmethemojo\Authentification\TwitterExtended;
use Itsmethemojo\Authentification\Redirect;
use Itsmethemojo\Authentification\ParameterException;
use Itsmethemojo\File\Config;

$iniFilename = 'remark';

try {
    $debug = boolval(Config::get($iniFilename, array('DEBUG_MODE'))['DEBUG_MODE']);
} catch (Exception $e) {
    $debug = false;
}

$config = [
    'settings' => [
        'displayErrorDetails' => $debug,
        'iniFileName' => $iniFilename,
        'loginIniFileName' => 'login'
    ],
];

$app = new \Slim\App($config);

//setup caching
$app->add(new \Slim\HttpCache\Cache());
$container = $app->getContainer();
$container['cache'] = function () {
    return new \Slim\HttpCache\CacheProvider();
};

$app->get(
    '/',
    function ($request, $response, $args) {
        $twitter = new TwitterExtended($this->get('settings')['loginIniFileName']);
        if (!$twitter->isLoggedIn()) {
            $output = $response->withStatus(401)->withJson(array("status" => "not authorized"));
            return $this->cache->allowCache($output, 'public', 0);
        }

        $userId    = 1;
        $bookmarks = new Itsmethemojo\Remark\Bookmarks(
            $this->get('settings')['iniFileName'],
            $this->get('settings')['iniFileName']
        );
        $data      = $bookmarks->getAll($userId);

        return $response->withJson($data);
    }
);

//TODO change this to post
$app->get(
    '/click/{id}/',
    function ($request, $response, $args) {
        $twitter = new TwitterExtended($this->get('settings')['loginIniFileName']);
        if (!$twitter->isLoggedIn()) {
            $output = $response->withStatus(401)->withJson(array("status" => "not authorized"));
            return $this->cache->allowCache($output, 'public', 0);
        }

        $userId    = 1;
        $bookmarks = new Itsmethemojo\Remark\Bookmarks(
            $this->get('settings')['iniFileName'],
            $this->get('settings')['iniFileName']
        );
        $data      = $bookmarks->click($userId, $args['id']);
        return $response->withJson($data);
    }
);

//TODO change this to post
$app->get(
    '/remark/',
    function ($request, $response, $args) {
        $twitter = new TwitterExtended($this->get('settings')['loginIniFileName']);
        if (!$twitter->isLoggedIn()) {
            $output = $response->withStatus(401)->withJson(array("status" => "not authorized"));
            return $this->cache->allowCache($output, 'public', 0);
        }

        $userId = 1;
        if (!$request->getParam('url')) {
            return $response->withStatus(400)->withJson(
                array(
                    'error' => 'missing parameter: url'
                )
            );
        }
        $bookmarks = new Itsmethemojo\Remark\Bookmarks(
            $this->get('settings')['iniFileName'],
            $this->get('settings')['iniFileName']
        );
        $data = $bookmarks->remark(
            $userId,
            $request->getParam('url'),
            $request->getParam('title')
        );

        // to use it in browser extension??
        //header("Access-Control-Allow-Origin: *");
        return $response->withJson($data);
    }
);

$app->run();
