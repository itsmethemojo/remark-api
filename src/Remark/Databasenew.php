<?php

namespace Itsmethemojo\Remark;

use Itsmethemojo\File\ConfigReader;
use PDO;
use Exception;

class Databasenew
{

    /** @var PDO **/
    private $pdo = null;

    /** @var String **/
    private $iniFile = null;

    public function __construct($iniFile = 'mysql')
    {
        $this->iniFile = $iniFile;
    }

    public function getPdo()
    {
        if ($this->pdo !== null) {
            return $this->pdo;
        }
        $config = ConfigReader::get(
            $this->iniFile,
            array('MYSQL_USERNAME', 'MYSQL_PASSWORD', 'MYSQL_HOST', 'MYSQL_DATABASE')
        );
        $this->pdo = new PDO(
            'mysql:host=' . $config['MYSQL_HOST'] .
            ';port=' . $config['MYSQL_PORT'] .
            ';dbname=' . $config['MYSQL_DATABASE'] .
            ';charset=utf8',
            $config['MYSQL_USERNAME'],
            $config['MYSQL_PASSWORD']
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

        //TODO check this array more proper
        //TODO use custom exception
        if (is_array($stmt->errorInfo()) && $stmt->errorInfo()[1] != null) {
            throw new Exception($stmt->errorInfo()[2]);
        }

        if ($exceptionOnEmptyUpdate && $stmt->rowCount() === 0) {
            throw new EmptyUpdateException('no line to update');
        }

        return $stmt->fetchAll(PDO::FETCH_ASSOC);
    }
}