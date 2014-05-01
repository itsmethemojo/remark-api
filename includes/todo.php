<?php

$todo = new TodoController("todo");

if(isset($_POST['save'])){
    $todo->save($userid, $_POST['data']);
}
else{
    $todo->showList($userid);
}
?>