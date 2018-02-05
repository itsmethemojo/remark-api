<?php

namespace Itsmethemojo\Remark;

use Exception;

class Project
{

    public static function readEnvironmentConfig()
    {

        $filePath = Project::getPath().'/config/.env';

        if (!file_exists($filePath)) {
            return;
        }

        foreach (parse_ini_file($filePath) as $key => $value) {
            putenv($key . '="' . $value .'"');
        }
    }

    public static function getPath()
    {
        $currentFileDir = __DIR__ . '/';
        if (strpos($currentFileDir, '/vendor/')) {
            return preg_split('/\/vendor\//', $currentFileDir, 2)[0];
        }
        if (strpos($currentFileDir, '/src/')) {
            return preg_split('/\/src\//', $currentFileDir, 2)[0];
        }
        throw new Exception('unusual file strucure, expecting this classfile somewhere in vendor or src folder');
    }
}
