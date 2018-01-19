<?php

namespace Itsmethemojo\Remark;

use Redis;
use Itsmethemojo\File\ConfigReader;

class KeyValueStore
{

    private $redis = null;
    private $config = null;

    public static $ttl = 86400; //60*60*24

    public function __construct($configKey)
    {
        $this->config = ConfigReader::get($configKey, array('REDIS_HOST', 'REDIS_PREFIX'));
        if (!isset($this->config['REDIS_PORT'])) {
            $this->config['REDIS_PORT'] = 6379;
        }
        return $this;
    }

    public function connect($redisAlreadyConnected = null)
    {
        if (!$redisAlreadyConnected) {
            $this->redis = new Redis();
            $this->redis->connect($this->config['REDIS_HOST'], $this->config['REDIS_PORT']);
        } else {
            $this->redis = $redisAlreadyConnected;
        }
        return $this;
    }

    public function set($key, $value, $ttl = 0)
    {
        if (is_array($value) || is_object($value)) {
            $value = "json>" . json_encode($value);
        }
        if (!is_numeric($ttl) || $ttl <= 0) {
            $ttl = KeyValueStore::$ttl;
        }
        return $this->redis->setex($this->transformKeys($key), $ttl, $value);
    }

    public function get($key)
    {
        $value = $this->redis->get($this->transformKeys($key));
        if (substr($value, 0, 5) === "json>") {
            return json_decode(substr($value, 5), true);
        }
        return $value;
    }

    public function mGet($keys)
    {
        return $this->redis->mGet($this->transformKeys($keys));
    }

    public function incr($key)
    {
        return $this->redis->incr($this->transformKeys($key));
    }

    public function isConnected()
    {
        return $this->redis !== null;
    }

    private function transformKeys($keys)
    {
        if (!is_array($keys)) {
            return $this->config['REDIS_PREFIX'].$keys;
        }
        $transformedKeys = [];
        foreach ($keys as $key) {
            $transformedKeys[] = $this->config['REDIS_PREFIX'].$key;
        }
        return $transformedKeys;
    }
}
