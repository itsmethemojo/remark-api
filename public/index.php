<?php
header("Cache-Control: public, max-age=0, no-cache");
header('Cache-Control: no-store, no-cache, must-revalidate');
header('Cache-Control: post-check=0, pre-check=0', false);
header('Pragma: no-cache');

error_reporting(E_ALL | E_STRICT);
ini_set('display_errors', 1);

require __DIR__ . '/../vendor/autoload.php';

$app = new \Slim\App();

$app->get(
    '/',
    function ($request, $response, $args) {
        $twitter = new Itsmethemojo\Authentification\TwitterExtended();
        if (!$twitter->isLoggedIn()) {
            return $response->withStatus(401)->withJson(array("message" => "not authorized"));
        }
        try {
            //TODO remove this workaround
            if ($twitter->getLoginUser()["id"] !== "itsmethemojo") {
                return $response->withStatus(401)->withJson(array("message" => "not authorized"));
            }
            $userId    = 1;
            $bookmarks = new Itsmethemojo\Remark\Bookmarks();
            $data      = $bookmarks->getAll($userId);

            return $response->withJson($data);
        } catch (Exception $ex) {
            return $response->withStatus(400)->withJson(
                array(
                    'error' => $ex->getMessage()
                )
            );
        }
    }
);

//TODO change this to post
$app->get(
    '/click/{id}/',
    function ($request, $response, $args) {
        $twitter = new Itsmethemojo\Authentification\TwitterExtended();
        if (!$twitter->isLoggedIn()) {
            return $response->withStatus(401)->withJson(array("message" => "not authorized"));
        }

        try {
            //TODO remove this workaround
            if ($twitter->getLoginUser()["id"] !== "itsmethemojo") {
                return $response->withStatus(401)->withJson(array("message" => "not authorized"));
            }
            $userId    = 1;
            $bookmarks = new Itsmethemojo\Remark\Bookmarks();
            $data      = $bookmarks->click($userId, $args['id']);
            return $response->withJson($data);
        } catch (Exception $ex) {
            return $response->withStatus(400)->withJson(
                array(
                    'error' => $ex->getMessage()
                )
            );
        }
    }
);

//TODO change this to post
$app->get(
    '/remark/',
    function ($request, $response, $args) {
        $twitter = new Itsmethemojo\Authentification\TwitterExtended();
        if (!$twitter->isLoggedIn()) {
            return $response->withStatus(401)->withJson(array("message" => "not authorized"));
        }
        try {
            //TODO remove this workaround
            if ($twitter->getLoginUser()["id"] !== "itsmethemojo") {
                return $response->withStatus(401)->withJson(array("message" => "not authorized"));
            }
            $userId    = 1;
            if (!$request->getParam('url')) {
                return $response->withStatus(400)->withJson(
                    array(
                        'error' => 'missing parameter: url'
                    )
                );
            }
            $bookmarks = new Itsmethemojo\Remark\Bookmarks();
            $data = $bookmarks->remark(
                $userId,
                $request->getParam('url'),
                $request->getParam('title')
            );
            //header("Access-Control-Allow-Origin: *");
            return $response->withJson($data);
        } catch (Exception $ex) {
            return $response->withStatus(400)->withJson(
                array(
                    'error' => $ex->getMessage()
                )
            );
        }
    }
);

$app->run();
