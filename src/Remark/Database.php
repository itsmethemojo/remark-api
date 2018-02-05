<?php

namespace Itsmethemojo\Remark;

use PDO;
use Exception;

class Database
{

    /** @var PDO **/
    private $pdo = null;

    public function getPdo()
    {
        if ($this->pdo !== null) {
            return $this->pdo;
        }
        $this->pdo = new PDO(
            'mysql:host=' . EnvironmentParameter::get('MYSQL_HOST') .
            ';port=' . EnvironmentParameter::get('MYSQL_PORT', 3306) .
            ';dbname=' . EnvironmentParameter::get('MYSQL_DATABASE') .
            ';charset=utf8',
            EnvironmentParameter::get('MYSQL_USERNAME'),
            EnvironmentParameter::get('MYSQL_PASSWORD')
        );
        return $this->pdo;
    }

    public function query($query, QueryParameters $parameters = null, $exceptionOnEmptyUpdate = false)
    {
        $pdo = $this->getPdo();

        $stmt = $pdo->prepare($query);

        $stmt->execute(
            $parameters === null ? array() : $parameters->toArray()
        );

        if (is_array($stmt->errorInfo()) && $stmt->errorInfo()[1] != null) {
            // this mostly happens if a query is wrong or there is an index duplicate
            throw new Exception($stmt->errorInfo()[2]);
        }

        if ($exceptionOnEmptyUpdate && $stmt->rowCount() === 0) {
            throw new EmptyUpdateException('no line to update');
        }

        return $stmt->fetchAll(PDO::FETCH_ASSOC);
    }
}
