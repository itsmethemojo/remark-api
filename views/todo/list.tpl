<!DOCTYPE html>
<html>
    <head>
        <meta charset="utf-8" />
        <title><?php echo $title; ?></title>
        <link href="img/favicon.ico" rel="shortcut icon" type="image/x-icon">
        <link rel="stylesheet" type="text/css" href="src/style/todo.css">
        <script type="text/javascript">
            if('<?php echo $json; ?>'==''){
                var archiv = new Array();
                var todos = new Array();
            }
            else{
                var archiv = JSON.parse('<?php echo $json; ?>');
                var todos = JSON.parse('<?php echo $json; ?>');
            }
            </script>
        <script type="text/javascript" src="src/script/todo.js" ></script>
    </head>
    <body onload="printList();">
        <div id="container">

        </div>
        <form action="?app=todo" method="POST">
            <input type="hidden" name="data" id="data"/>
            <input type="submit" name="save" id="save" value="save"/>
        </form>
    </body>
</html>