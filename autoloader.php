<?php

include_once realpath(dirname(__FILE__)).DIRECTORY_SEPARATOR.'Configuration'.DIRECTORY_SEPARATOR.'settings.php';
include_once realpath(dirname(__FILE__)).DIRECTORY_SEPARATOR.'Classes'.DIRECTORY_SEPARATOR.'Controller'.DIRECTORY_SEPARATOR.'RemarkController.php';
$remarkController = new RemarkController($dbConfig,realpath(dirname(__FILE__)).DIRECTORY_SEPARATOR);
