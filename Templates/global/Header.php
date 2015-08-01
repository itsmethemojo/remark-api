<!DOCTYPE html>
<html>
    <head>
        <title><?php echo $this->meta['title'];?></title>
        <meta charset="UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1.0">
        
        <?php if(isset($this->javascript)) foreach($this->javascript as $path){?>
        <script src="<?php echo $path;?>" type="text/javascript"></script>
        <?php }?>
        <?php if(isset($this->css)) foreach($this->css as $path){?>
        <script src="<?php echo $path;?>"></script>
        <link rel="stylesheet" type="text/css" href="<?php echo $path;?>">
        <?php }?>
    </head>
    <body>