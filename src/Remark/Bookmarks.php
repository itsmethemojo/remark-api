<?php

namespace Itsmethemojo\Remark;

use Exception;

class Bookmarks
{
    const NO_TITLE = "NO_TITLE";

    /** @var Database * */
    private $database;

    /** @var int * */
    private $timestamp;

    /** @var String **/
    private $iniFile = null;

    public function __construct($iniFile)
    {
        $this->iniFile = $iniFile;
    }

    public function getAll($userId)
    {
        $params = (new QueryParameters())
                  ->add($userId);

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

        return $this->getDatabase()->query(
            $query,
            $params
        );
    }

    private function getDatabase()
    {
        if ($this->database !== null) {
            return $this->database;
        }

        $this->database = new Database($this->iniFile);

        return $this->database;
    }

    public function click($userId, $bookmarkId)
    {
        $params = (new QueryParameters())
                  ->add($this->getTimestamp())
                  ->add($bookmarkId)
                  ->add($userId);
        $updateClickCountQuery = "UPDATE bookmark SET clickcount = clickcount + 1, clicked = ? "
                                  . "WHERE id = ? AND user_id = ?";
        // first i update the overall count.
        // if there is no updated line this bookmark does not belong to this user or does not exist
        try {
            $this->getDatabase()->query(
                $updateClickCountQuery,
                $params,
                true
            );
        } catch (EmptyUpdateException $e) {
            throw new Exception('not your bookmark');
        }


        $insertClickTimeQuery = "INSERT INTO clicktime (created, bookmark_id, user_id) VALUES (?,?,?)";
        $this->getDatabase()->query(
            $insertClickTimeQuery,
            $params
        );

        $params = (new QueryParameters())
                  ->add($bookmarkId);

        $readClickCountQuery = "SELECT clickcount FROM bookmark WHERE id = ?";
        $result = $this->getDatabase()->query(
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
        $params = (new QueryParameters())
                  ->add($this->getTimestamp())
                  ->add($url)
                  ->add($userId);
        $updateRemarkCountQuery = "UPDATE bookmark SET bookmarkcount =  bookmarkcount + 1, updated = ? "
                                  . "WHERE url = ? AND user_id = ?";

        try {
            $this->getDatabase()->query(
                $updateRemarkCountQuery,
                $params,
                true
            );
        } catch (EmptyUpdateException $e) {
            // so this bookmark is new
            $urlArray = parse_url($url);

            if (!$urlArray || !$urlArray['host']
            ) {
                throw new Exception("no valid url");
            }

            if ($title === null) {
                $title = $this->retrieveUrlTitle($url);
            }

            $insertBookmarkQuery = "INSERT INTO bookmark "
                                    . "(url, title, user_id, domain, created, updated, bookmarkcount) "
                                    . "VALUES (?,?,?,?,?,?,?)";
            $params = (new QueryParameters())
                      ->add($url)
                      ->add($title)
                      ->add($userId)
                      ->add($urlArray['host'])
                      ->add($this->getTimestamp())
                      ->add($this->getTimestamp())
                      ->add(1);

            // throws exception when no insert is made ???
            $this->getDatabase()->query(
                $insertBookmarkQuery,
                $params,
                true
            );
        }

        $params = (new QueryParameters())
                  ->add($url)
                  ->add($userId);

        $readClickCountQuery = "SELECT bookmarkcount, id FROM bookmark WHERE url = ? AND user_id = ?";

        $result = $this->getDatabase()->query(
            $readClickCountQuery,
            $params
        );

        if (count($result) !== 1) {
            throw new Exception("ohohh");
        }

        $params = (new QueryParameters())
                  ->add($this->getTimestamp())
                  ->add($result[0]['id'])
                  ->add($userId);

        $insertRemarkTimeQuery = "INSERT INTO bookmarktime (created, bookmark_id, user_id) VALUES (?,?,?)";

        $this->getDatabase()->query(
            $insertRemarkTimeQuery,
            $params,
            true
        );

        return array("bookmarkcount" => $result[0]['bookmarkcount']);
    }

    private function retrieveUrlTitle($url)
    {
        //TODO set proper user-agent
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

    private function getTimestamp()
    {
        if ($this->timestamp === null) {
            $this->timestamp = time();
        }
        return $this->timestamp;
    }
}
