<?php

namespace Itsmethemojo\Remark;

use Exception;

class QueryParameters
{

    private $parameters = [];

    public function add($parameter)
    {
        if (!is_int($parameter)
            && !is_float($parameter)
            && !is_string($parameter)
        ) {
            //TODO use custom exception
            throw new Exception(
                'query parameter must be either int, float or string'
            );
        }
        $this->parameters[] = $parameter;
        return $this;
    }

    public function toArray()
    {
        return $this->parameters;
    }
}
