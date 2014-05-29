<?php

class TodoModel extends Model
{
	public function __construct($dbLink) {
            error_log("Todo Model construct");
            parent::__construct($dbLink);
	}
	
        
        function getData($userid){
            $query = "SELECT data FROM todo WHERE user_id = ? order by id DESC LIMIT 0,1";
            $json = null;
            if ($stmt = mysqli_prepare($this->dbLink, $query)) {
                mysqli_stmt_bind_param($stmt, "s", $userid);
                mysqli_stmt_execute($stmt);
                mysqli_stmt_bind_result($stmt, $json);
                mysqli_stmt_fetch($stmt);
                mysqli_stmt_close($stmt);
            }
            
            return $json;
        }
        
        function saveData($userid,$json){
            $timeStamp = time();

            //TODO add created to table
            $query = "INSERT INTO todo (user_id, data) VALUES (?,?)";
            //$query = "INSERT INTO todo (user_id, created, data) VALUES (?,?,?)";
            $stmt = mysqli_prepare($this->dbLink, $query);
            //mysqli_stmt_bind_param($stmt, "sss", $userid, $timeStamp, $json);
            mysqli_stmt_bind_param($stmt, "ss", $userid, $json);
            mysqli_stmt_execute($stmt);
            mysqli_stmt_close($stmt);
            
            return $this->getData($userid);
        }
}

