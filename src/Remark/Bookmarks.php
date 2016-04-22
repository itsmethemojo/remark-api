<?php

namespace Itsmethemojo\Remark;

use Itsmethemojo\Storage\Database;
use Itsmethemojo\Storage\QueryParameters;
use Exception;

class Bookmarks
{
    const NO_TITLE = "NO_TITLE";

    /** @var Database * */
    private $database;

    /** @var int * */
    private $timestamp;

    public function __construct($databaseConfigKey = "remark-mysql", $storageConfigKey = "remark-redis")
    {
        $this->database = new Database($databaseConfigKey,$storageConfigKey);
    }

    public function getAll($userId)
    {
        $params = new QueryParameters();
        $params->add($userId);

        $query = "SELECT
                    b.id as id,
                    b.url as url,
                    b.domain as domain,
                    b.title as title,
                    b.customtitle as customtitle,
                    b.bookmarkcount as remarks,
                    b.clickcount as clicks,
                    bt.created as created,
                    b.updated as updated,
                    b.clicked as clicked
                    FROM bookmark b
                    JOIN bookmarktime bt
                    ON b.id = bt.bookmark_id WHERE b.user_id = ? ORDER BY bt.id DESC";

        return $this->database->read(
            array('allData-' . $userId),
            $query,
            $params,
            false,
            60 * 60 * 24 * 30
        );
    }

    public function click($userId, $bookmarkId)
    {
        if (!$this->isUsersBookmark($userId, $bookmarkId)) {
            throw new Exception("action not allowed");
        }

        $query1  = "INSERT INTO clicktime (bookmark_id, user_id, created) VALUES (?,?,?)";
        $params1 = new QueryParameters();
        $params1->add($bookmarkId)->add($userId)->add($this->getTimestamp());

        $this->database->modify(
            array(),
            $query1,
            $params1
        );

        $query2  = "UPDATE bookmark SET clickcount = clickcount + 1, clicked = ? WHERE id = ?";
        $params2 = new QueryParameters();
        $params2->add($this->getTimestamp())->add($bookmarkId);

        //final operation -> use invalidate tags
        $this->database->modify(
            array('allBookmarkIdsAndUrls', 'allData-' . $userId),
            $query2,
            $params2
        );

        $params3 = new QueryParameters();
        $params3->add($bookmarkId);

        $query3 = "SELECT clickcount FROM bookmark WHERE id = ?";
        $result = $this->database->read(
            array(),
            $query3,
            $params3
        );

        if (count($result) !== 1) {
            throw new Exception("ohohh");
        }

        return array("clicks" => $result[0]['clickcount']);
    }

    public function remark($userId, $url, $title = null)
    {
        $bookmarkId = $this->getBookmarkId($userId, $url);
        if ($bookmarkId === null) {
            $this->insertBookmark($userId, $url, $title);
            $bookmarkId = $this->getBookmarkId($userId, $url);
            if ($bookmarkId === null) {
                throw new Exception("ohohh");
            }
        }

        $params1 = new QueryParameters();
        $params1
            ->add($this->getTimestamp())
            ->add($bookmarkId);

        $query1 = "UPDATE bookmark SET bookmarkcount =  bookmarkcount + 1, updated = ? WHERE id = ?";

        $this->database->modify(
            array('allBookmarkIdsAndUrls', 'allData-' . $userId),
            $query1,
            $params1
        );

        $params2 = new QueryParameters();
        $params2->add($bookmarkId);

        $query2 = "SELECT bookmarkcount FROM bookmark WHERE id = ?";
        $result = $this->database->read(
            array(),
            $query2,
            $params2
        );

        if (count($result) !== 1) {
            throw new Exception("ohohh");
        }

        return array("bookmarkcount" => $result[0]['bookmarkcount']);
    }

    public function delete($userId, $url)
    {
        $bookmarkId = $this->getBookmarkId($userId, $url);
        if ($bookmarkId === null) {
            throw new Exception("can not delete what does not exist");
        }

        $params1 = new QueryParameters();
        $params1->add($bookmarkId);
        $query1  = "DELETE FROM bookmark WHERE id = ?";
        $this->database->modify(
            array('allBookmarkIdsAndUrls', 'allData-' . $userId),
            $query1,
            $params1
        );
    }

    public function updateTitle($userId, $url, $title)
    {
        $bookmarkId = $this->getBookmarkId($userId, $url);
        if ($bookmarkId === null) {
            throw new Exception("can not modify what does not exist");
        }

        $params1 = new QueryParameters();
        $params1
            ->add($title)
            ->add($bookmarkId);
        $query1  = "UPDATE bookmark SET title = ? WHERE id = ?";
        $this->database->modify(
            array('allData-' . $userId),
            $query1,
            $params1
        );
    }

    private function insertBookmark($userId, $url, $title = null)
    {
        $urlArray = parse_url($url);
        if (!$urlArray || !$urlArray['host']
        ) {
            throw new Exception("no valid url");
        }

        if ($title === null) {
            $title = $this->retrieveUrlTitle($url);
        }

        $query1 = "INSERT INTO bookmark (url, title, user_id, domain, created, updated) VALUES (?,?,?,?,?,?)";

        $params1 = new QueryParameters();
        $params1
            ->add($url)
            ->add($title)
            ->add($userId)
            ->add($urlArray['host'])
            ->add($this->getTimestamp())
            ->add($this->getTimestamp());

        $this->database->modify(
            array('allBookmarkIdsAndUrls'),
            $query1,
            $params1
        );
    }

    private function retrieveUrlTitle($url)
    {
        try {
            $client = new \GuzzleHttp\Client();
            $res    = $client->request('GET', $url);

            if ($res->getStatusCode() === 200) {
                $matches = array();
                preg_match(
                    "/<title>(.+)<\/title>/siU",
                    $res->getBody(),
                    $matches
                );
                if (is_array($matches) && isset($matches[1])) {
                    return $matches[1];
                }
            }
        } catch (Exception $e) {
        }

        return self::NO_TITLE;
    }

    private function getBookmarkId($userId, $url)
    {
        $query = "SELECT user_id, id as bookmark_id, url FROM bookmark";

        $availableBookmarks = $this->database->read(
            array('allBookmarkIdsAndUrls'),
            $query
        );

        //this looks dumb
        foreach ($availableBookmarks as $availableBookmark) {
            if ((string) $availableBookmark['user_id'] === (string) $userId && ( $availableBookmark['url']
                === $url || $availableBookmark['url'] === $url . "/" || $availableBookmark['url'] . "/"
                === $url )
            ) {
                return $availableBookmark['bookmark_id'];
            }
        }
        return null;
    }

    private function isUsersBookmark($userId, $bookmarkId)
    {
        $query = "SELECT user_id, id as bookmark_id, url FROM bookmark";

        $availableBookmarks = $this->database->read(
            array('allBookmarkIdsAndUrls'),
            $query
        );

        //this looks dumb
        foreach ($availableBookmarks as $availableBookmark) {
            if ($availableBookmark['bookmark_id'] == $bookmarkId && $availableBookmark['user_id']
                == $userId
            ) {
                return true;
            }
        }
        return false;
    }

    private function getTagsToInvalidate($userId)
    {
        return array(
            'allBookmarkIdsAndUrls',
            'allData-' . $userId,
        );
    }

    private function getTimestamp()
    {
        if ($this->timestamp === null) {
            $this->timestamp = time();
        }
        return $this->timestamp;
    }
}
