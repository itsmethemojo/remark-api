<?php

namespace Itsmethemojo\Remark;

use Exception;

class EnvironmentParameter
{
    public static function get($key, $ifNotSet = null)
    {
        $value = getenv($key);
        if ($value !== false) {
            return $value;
        }
        if ($ifNotSet === null) {
            throw new Exception("missing parameter: " . $key);
        }
        return $ifNotSet;
    }
}
