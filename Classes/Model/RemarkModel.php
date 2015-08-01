
<?php


class RemarkModel extends BaseModel{
    
    public function retrieveCompleteList($userId){
        $query = "SELECT 
                    b.id as id,
                    b.url as url,
                    b.domain as domain,
                    b.extension as extension,
                    b.title as title,
                    b.customtitle as customtitle,
                    b.bookmarkcount as bookmarkcount,
                    b.clickcount as clickcount,
                    b.tagcount as tagcount,
                    b.mediatype as mediatype,
                    bt.created as created,
                    b.updated as updated,
                    b.clicked as clicked
                    FROM bookmark b 
                    JOIN bookmarktime bt
                    ON b.id = bt.bookmark_id WHERE b.user_id = ? ORDER BY bt.id DESC";
        return $this->queryDatabase($query, array($userId));
    }        

    // tracks the click on a remarked Link and returns the url
    public function trackClick($userId,$bookmarkId){
        $timeStamp = time();

        $query1 = "INSERT INTO clicktime (bookmark_id, user_id, created) VALUES (?,?,?)";
        $this->queryDatabase($query1, array($bookmarkId, $userId, $timeStamp));

        $query2 = "UPDATE bookmark SET clickcount = clickcount + 1, clicked = ? WHERE id = ?";
        $this->queryDatabase($query2, array($timeStamp, $bookmarkId));

        $query3 = "SELECT url FROM bookmark WHERE id = ?";
        $result3 = $this->queryDatabase($query3, array($bookmarkId));

        return $result3[0]['url'];
    }
        
    public function setTitle($userId,$bookmarkId,$title){
        $query = "UPDATE bookmark
                    SET customtitle = ?
                    WHERE user_id = ?
                    AND id = ?";
        return $this->queryDatabase($query, array($title,$userId,$bookmarkId));
    }
    
    private function normalizeUrl($url){
        //TODO implement cutting last / and/or other stuff
        return $url;
    }
    
    private function getIdByUrl($userId,$url){	
        $normalizedUrl = $this->normalizeUrl($url);
        $query = "SELECT id FROM bookmark WHERE url = ? AND user_id = ?";
        $result = $this->queryDatabase($query, array($normalizedUrl, $userId));
        if($result && isset($result[0]) && isset($result[0]['id'])){
            return $result[0]['id'];
        }
        return false;
    }
    
    public function saveBookmark($userId,$url,$title){
        $timeStamp = time();
        $bookmarkId = $this->getIdByUrl($userId,$url);
        if(!$bookmarkId){
            $normalizedUrl = $this->normalizeUrl($url);
            
            $urlArray = parse_url($url);
            $domain = isset($urlArray['host']) ? $urlArray['host'] : "";
            $extension = isset($urlArray['path']) ? pathinfo($urlArray['path'], PATHINFO_EXTENSION) : "";
            $mediatype = $this->getMediaType($extension,$domain);

            // write in bookmark table
            $query1 = "INSERT INTO bookmark (url, title, user_id, domain, extension, mediatype, created, updated) VALUES (?,?,?,?,?,?,?,?)";
            $this->queryDatabase($query1, array($normalizedUrl, $title, $userId, $domain, $extension, $mediatype, $timeStamp, $timeStamp));

            $bookmarkId = $this->getIdByUrl($userId,$url);
        }
        else{
            // update bookmark table
            $query1 = "UPDATE bookmark SET bookmarkcount =  bookmarkcount + 1, updated = ? WHERE id = ?";
            $this->queryDatabase($query1, array($timeStamp, $bookmarkId));
        }
		
        // write in bookmarktime table
        $query2 = "INSERT INTO bookmarktime (bookmark_id, user_id, created) VALUES (?,?,?)";
        $this->queryDatabase($query2, array($bookmarkId, $userId, $timeStamp));

        return 1; //TODO maybe return something more usefull
 
    }
    
    private function getMediaType($extension, $domain) {
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
    
    /**
     * removes special chars that can break the json
     */
    private function normalizeString($inputString){
        
        return preg_replace($pattern, $replacement, $subject);
    }

}
