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
        $params
            ->add($this->getTimestamp())
            ->add($bookmarkId)
            ->add($userId);
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
        $params = new QueryParameters();
        $params
            ->add($this->getTimestamp())
            ->add($url)
            ->add($userId);
        $updateRemarkCountQuery = "UPDATE bookmark SET bookmarkcount =  bookmarkcount + 1, updated = ? "
                                  . "WHERE url = ? AND user_id = ?";

        try {
            $this->getDatabaseNew()->query(
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
            $params
                ->clear()
                ->add($url)
                ->add($title)
                ->add($userId)
                ->add($urlArray['host'])
                ->add($this->getTimestamp())
                ->add($this->getTimestamp())
                ->add(1);

            // throws exception when no insert is made ???
            $this->getDatabaseNew()->query(
                $insertBookmarkQuery,
                $params,
                true
            );
        }

        $params
            ->clear()
            ->add($url)
            ->add($userId);

        $readClickCountQuery = "SELECT bookmarkcount, id FROM bookmark WHERE url = ? AND user_id = ?";

        $result = $this->getDatabaseNew()->query(
            $readClickCountQuery,
            $params
        );

        if (count($result) !== 1) {
            throw new Exception("ohohh");
        }

        $params
            ->clear()
            ->add($this->getTimestamp())
            ->add($result[0]['id'])
            ->add($userId);

        $insertRemarkTimeQuery = "INSERT INTO bookmarktime (created, bookmark_id, user_id) VALUES (?,?,?)";

        $this->getDatabaseNew()->query(
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
