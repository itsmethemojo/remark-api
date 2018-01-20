<?php

namespace Itsmethemojo\Remark;

use Exception;

class QueryParameters
{

    private $parameters = [];

    public function add($parameter)
    {
        $this->parameters[] = $parameter;
        return $this;
    }

    public function toArray()
    {
        return $this->parameters;
    }
}
