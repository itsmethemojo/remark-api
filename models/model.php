<?php

class Model 
{
	/**
		initializes variables with null
		IMPORTANT NOTE: if inherit from this class you have to reimpliment __construct with parent::__construct()
	*/
	function __construct() {
		include HOME . DS . 'config' . DS . 'db.php';
		echo '';
		$this->dbLink = mysqli_connect($database_host, 
													$database_username, 
													$database_password, 
													$database_databasename, 
													$database_port);
		if(!$this->dbLink){
			throw new Exception('Couldn\'t connect to db. Check db config!');
		}
	}	
	
	/**
		automaticly closes Connection of the mysql Object
		IMPORTANT NOTE: if inherit from this class you have to reimpliment __destruct with parent::__destruct()
	*/
	function __destruct() {
		//echo "Destroying everything";
		// close mysql connection
		if($this->dbLink){
			mysqli_close($this->dbLink);
		}
	}
}
