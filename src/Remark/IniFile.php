<?php

namespace Itsmethemojo\Remark;

use Exception;

class IniFile
{

    public static function readConfig($key, $mustHaveKeys = array())
    {

        $filePath = IniFile::getConfigPath().'/'.$key.'.ini';

        if (!file_exists($filePath)) {
            throw new Exception('the following ini file is missing ' . $filePath);
        }

        $data = parse_ini_file($filePath);
        foreach ($mustHaveKeys as $mustHaveKey) {
            if (!array_key_exists($mustHaveKey, $data)) {
                throw new Exception("the file \"" . $key . ".ini\" misses the required key \"" . $mustHaveKey . "\"");
            }
        }

        return $data;
    }

    public static function getConfigPath()
    {
        $currentFileDir = __DIR__ . '/';
        if (strpos($currentFileDir, '/vendor/')) {
            return preg_split('/\/vendor\//', $currentFileDir, 2)[0] . '/config';
        }
        if (strpos($currentFileDir, '/src/')) {
            return preg_split('/\/src\//', $currentFileDir, 2)[0] . '/config';
        }
        throw new Exception('unusual file strucure, expecting this classfile somewhere in vendor or src folder');
    }
}
