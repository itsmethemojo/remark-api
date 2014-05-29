<?php

class Model {
    
    public function __construct($dbLink=NULL) {
        if($dbLink==NULL){
            error_log("Model dblink created");
            $this->connect();
        }
        else{
            error_log("Model dblink copied");
            $this->dbLink = $dbLink;
        }
    }

    public function connect(){
        include HOME . DS . 'config' . DS . 'db.php';
        $this->dbLink = mysqli_connect($database_host, $database_username, $database_password, $database_databasename, $database_port);
        if (!$this->dbLink) {
            throw new Exception('Couldn\'t connect to db. Check db config!');
        }
    }

    public function disconnect() {
        //echo "Destroying everything";
        // close mysql connection
        if ($this->dbLink) {
            mysqli_close($this->dbLink);
        }
    }
    
    function getUserMapping(){
        //
    }

}
