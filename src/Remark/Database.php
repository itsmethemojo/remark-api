<?php

namespace Itsmethemojo\Remark;

use Itsmethemojo\File\ConfigReader;
use Itsmethemojo\Remark\QueryParameters;
use Itsmethemojo\Remark\KeyValueStore;
use PDO;
use Exception;

class Database
{

    /** @var PDO **/
    private $database = null;

    /** @var KeyValueStore **/
    private $keyValueStore = null;

    /** @var mixed**/
    private $configuration = array();

    public function __construct($databaseConfigKey = "mysql", $storageConfigKey = "redis")
    {
        $this->configuration['database'] = ConfigReader::get(
            $databaseConfigKey,
            array('MYSQL_USERNAME', 'MYSQL_PASSWORD', 'MYSQL_HOST', 'MYSQL_DATABASE')
        );



        if (!array_key_exists('MYSQL_PREFIX', $this->configuration['database'])) {
            $this->configuration['database']['MYSQL_PREFIX'] = '';
        }
        if (!array_key_exists('MYSQL_PORT', $this->configuration['database'])) {
            $this->configuration['database']['MYSQL_PORT'] = 3306;
        }

        $this->keyValueStore = new KeyValueStore($storageConfigKey);
    }

    public function connect($pdo = null, $redis = null)
    {
        $this->mysqlLazyConnect($pdo);
        $this->redisLazyConnect($redis);
    }

    public function read($tags, $query, QueryParameters $parameters = null, $notSaveIfEmptyResult = false, $ttl = 0)
    {
        if (count($tags) === 0) {
            return $this->mysqlFetch($query, $parameters);
        }
        //check md5 performance
        $key = $this->getTagsPrefix($tags);
        $toHash = $query;
        if ($parameters !== null) {
            $toHash .= implode('-', $parameters->toArray());
        }

        $key .= md5($toHash);
        $this->redisLazyConnect();
        $cached = $this->keyValueStore->get($key);

        if (!$cached) {
            $cached = $this->mysqlFetch($query, $parameters);
            if (!$notSaveIfEmptyResult) {
                $this->keyValueStore->set($key, $cached, $ttl);
            }
        }
        return $cached;
    }

    public function modify($invalidateTags, $query, QueryParameters $parameters = null)
    {
        $this->incrementTags($invalidateTags);
        $this->mysqlFetch($query, $parameters);
    }

    public function killCache($invalidateTags)
    {
        $this->incrementTags($invalidateTags);
    }

    public function putInStore($key, $value, $ttl)
    {
        $this->redisLazyConnect();
        return $this->keyValueStore->set($key, $value, $ttl);
    }

    public function getFromStore($key)
    {
        $this->redisLazyConnect();
        return $this->keyValueStore->get($key);
    }
    //========================================
    //mysql functions

    private function mysqlFetch($query, QueryParameters $parameters = null)
    {
        $this->mysqlLazyConnect();

        $stmt = $this->database->prepare(
            str_replace(
                '#',
                $this->configuration['database']['MYSQL_PREFIX'],
                $query
            )
        );

        $stmt->execute(
            $parameters === null ? array() : $parameters->toArray()
        );

        //TODO check this array more proper
        if (is_array($stmt->errorInfo()) && $stmt->errorInfo()[1] != null) {
            throw new Exception($stmt->errorInfo()[2]);
        }

        return $stmt->fetchAll(PDO::FETCH_ASSOC);
    }

    private function mysqlLazyConnect($pdo = null)
    {
        if ($this->database === null) {
            if ($pdo) {
                $this->database = $pdo;
                return;
            }
            $this->database = new PDO(
                'mysql:host=' . $this->configuration['database']['MYSQL_HOST'] .
                ';port=' . $this->configuration['database']['MYSQL_PORT'] .
                ';dbname=' . $this->configuration['database']['MYSQL_DATABASE'] .
                ';charset=utf8',
                $this->configuration['database']['MYSQL_USERNAME'],
                $this->configuration['database']['MYSQL_PASSWORD']
            );
        }
    }

    //========================================
    //redis functions

    private function getTagsPrefix($tags)
    {
        $this->redisLazyConnect();
        $tagsPrefix = '';
        $tagCounts = $this->keyValueStore->mGet($this->transformTags($tags));
        for ($index = 0; $index < count($tags); $index++) {
            if ($tagCounts[$index] === false) {
                $this->keyValueStore->set('tag-' . $tags[$index], 1);
                $tagCounts[$index] = 1;
            }
            $tagsPrefix .= $tags[$index] . $tagCounts[$index] . '-';
        }
        return $tagsPrefix;
    }

    private function transformTags($tags)
    {
        $transformedTags = array();
        foreach ($tags as $tag) {
            //TODO validate tags (length, characters)
            $transformedTags[] = 'tag-' . $tag;
        }
        return $transformedTags;
    }

    private function incrementTags($tags)
    {
        $this->redisLazyConnect();
        foreach ($this->transformTags($tags) as $tag) {
            $this->keyValueStore->incr($tag);
        }
    }

    private function redisLazyConnect($redis = null)
    {
        if (!$this->keyValueStore->isConnected()) {
            $this->keyValueStore->connect($redis);
        }
    }
}
