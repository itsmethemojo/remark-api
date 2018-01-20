<?php

namespace Itsmethemojo\Remark;

//use Itsmethemojo\Storage\Database;
//use Itsmethemojo\Storage\QueryParameters;
use Exception;

class Bookmarks
{
    const NO_TITLE = "NO_TITLE";

    /** @var Database * */
    private $database;

    /** @var Databasenew * */
    private $databasenew;

    /** @var int * */
    private $timestamp;

    /** @var String **/
    private $iniFile = null;

    public function __construct($databaseConfigKey = "remark-mysql", $storageConfigKey = "remark-redis")
    {
        //TODO change this final
        //$this->iniFile = $iniFile;
        $this->iniFile = $databaseConfigKey;
        $this->database = new Database($databaseConfigKey, $storageConfigKey);
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

        return $this->getDatabaseNew()->query(
            $query,
            $params
        );
    }

    private function getDatabaseNew()
    {
        if ($this->databasenew !== null) {
            return $this->databasenew;
        }

        $this->databasenew = new Databasenew($this->iniFile);

        return $this->databasenew;
    }

    public function click($userId, $bookmarkId)
    {
        $params = new QueryParameters();
        $params->add($this->getTimestamp())->add($bookmarkId)->add($userId);
        $updateClickCountQuery = "UPDATE bookmark SET clickcount = clickcount + 1, clicked = ? "
                                  . "WHERE id = ? AND user_id = ?";
        // first i update the overall count.
        // if there is no updated line this bookmark does not belong to this user or does not exist
        try {
            $this->getDatabaseNew()->query(
                $updateClickCountQuery,
                $params,
                true
            );
        } catch (EmptyUpdateException $e) {
            throw new Exception('not your bookmark');
        }


        $insertClickTimeQuery = "INSERT INTO clicktime (created, bookmark_id, user_id) VALUES (?,?,?)";
        $this->getDatabaseNew()->query(
            $insertClickTimeQuery,
            $params
        );

        $params->clear()->add($bookmarkId);

        $readClickCountQuery = "SELECT clickcount FROM bookmark WHERE id = ?";
        $result = $this->database->read(
            array(),
            $readClickCountQuery,
            $params
        );

        if (count($result) !== 1) {
            throw new Exception("ohohh");
        }

        return array("clicks" => $result[0]['clickcount']);
    }


    public function remark($userId, $url, $title = null)
    {
        // refactor idea
        // first update again and catch exception
        // if exception insert
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
            ->add($bookmarkId)
            ->add($userId)
            ->add($this->getTimestamp());

        $query1 = "INSERT INTO bookmarktime (bookmark_id, user_id, created) VALUES (?,?,?)";

        $this->database->modify(
            array('allBookmarkIdsAndUrls', 'allData-' . $userId),
            $query1,
            $params1
        );

        $params2 = new QueryParameters();
        $params2
            ->add($this->getTimestamp())
            ->add($bookmarkId);

        $query2 = "UPDATE bookmark SET bookmarkcount =  bookmarkcount + 1, updated = ? WHERE id = ?";

        $this->database->modify(
            array('allBookmarkIdsAndUrls', 'allData-' . $userId),
            $query2,
            $params2
        );

        $params3 = new QueryParameters();
        $params3->add($bookmarkId);

        $query3 = "SELECT bookmarkcount FROM bookmark WHERE id = ?";
        $result = $this->database->read(
            array(),
            $query3,
            $params3
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
            array('allBookmarkIdsAndUrls', 'allData-' . $userId),
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
