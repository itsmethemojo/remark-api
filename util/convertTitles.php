<?php

//TODO read credentials from config
$dbLink = new mysqli("localhost:3306", "root", "root", "remark");
if ($dbLink->connect_errno) {
    throw new Exception("database mssing?", NULL, NULL);
}
//var_dump($dbLink);

$query = "SELECT title,id FROM bookmark";

$stmt = $dbLink->prepare($query);
$stmt->execute();
$stmtResult = $stmt->get_result();
        

$resultArr = array();
        
while ($row = $stmtResult->fetch_array(MYSQLI_ASSOC)){
    $resultArr[] = $row;
}
$stmt->close();
    
foreach ($resultArr as $row){
    $orgTitle = $row['title'];
    $newTitle = $row['title'];
    $newTitle = preg_replace('/ä/', '&auml;', $newTitle);
    $newTitle = preg_replace('/ö/', '&ouml;', $newTitle);
    $newTitle = preg_replace('/ü/', '&uuml;', $newTitle);
    $newTitle = preg_replace('/Ä/', '&Auml;', $newTitle);
    $newTitle = preg_replace('/Ö/', '&Ouml;', $newTitle);
    $newTitle = preg_replace('/Ü/', '&Uuml;', $newTitle);
    $newTitle = preg_replace('/ß/', '&szlig;', $newTitle);
    $newTitle = preg_replace('/[^A-Za-z0-9\- \|\.&;]/', '', $newTitle);
    
    if($orgTitle==$newTitle){
        continue;
    }
    
    echo $newTitle."\n";
    
    $stmt = $dbLink->prepare("update bookmark set title = ? where id = ?");
    $stmt->bind_param("ss", $newTitle,$row['id']);
    $stmt->execute();
    $stmtResult = $stmt->get_result();
    $stmt->close();
}
        


        
?>