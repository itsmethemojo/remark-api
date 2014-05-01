<?php

class FrontendCalculations {

    public function calculateRemarkVisibilityClass($number) {
        switch ($number) {
            case 0: return 0;
            case 1: return 1;
            case 2: return 2;
            case 3: return 4;
            case 4: return 6;
        }
        return 8;
    }

    public function calculateClickVisibilityClass($number) {
        switch ($number) {
            case 0: return 0;
            case 1: return 1;
            case 2: return 2;
            case 3: return 3;
        }
        if ($number <= 6)
            return 4;
        if ($number <= 10)
            return 5;
        if ($number <= 15)
            return 6;
        if ($number <= 20)
            return 7;

        return 8;
    }

    public function getMediaType($extension, $domain) {
        //mediatypes
        $probablyHTML = 1;
        $image = 2;
        $video = 3;
        $music = 4;
        switch ($extension) {
            case ("jpg"): return $image;
            case ("jpeg"): return $image;
            case ("gif"): return $image;
            case ("png"): return $image;
        }

        switch ($domain) {
            case ("www.youtube.com"): return $video;
            case ("www.youtube.de"): return $video;
            case ("www.vevo.com"): return $video;
            case ("vimeo.com"): return $video;
            case ("www.myvideo.de"): return $video;

            case ("soundcloud.com"): return $music;
        }

        return $probablyHTML;
    }

}

?>