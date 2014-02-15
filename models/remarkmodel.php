<?php

class RemarkModel extends Model
{
	function __construct() {
		parent::__construct();
	}
	
	function __destruct() {
		parent::__destruct();
	}

	// just shows all remarks
	public function getUnfilteredList($userid){
		$query = "SELECT 
					id,
					user_id,
					url,
					domain,
					extension,
					title,
					customtitle,
					bookmarkcount,
					clickcount,
					tagcount,
					mediatype,
					created,
					updated,
					clicked
					FROM bookmark WHERE user_id = ? ORDER BY id DESC";
		$counter = 0;
		$entries = array();
		if ($stmt = mysqli_prepare($this->dbLink, $query)) {
			mysqli_stmt_bind_param($stmt, "s", $userid);
			mysqli_stmt_execute($stmt);
			
		
			$result = $stmt->get_result();
			while ($row = $result->fetch_array(MYSQLI_ASSOC)){
				$entries[$counter]=$row;

				//add some additional data for display
				$entries[$counter]['clickclass'] = $this->calculateClickVisibilityClass($entries[$counter]['clickcount']);
				$entries[$counter]['remarkclass'] = $this->calculateRemarkVisibilityClass($entries[$counter]['bookmarkcount']);
				if($entries[$counter]['customtitle']!=""){
					$entries[$counter]['title']=$entries[$counter]['customtitle'];
				}
	
				$date = date_create();
				date_timestamp_set($date, $entries[$counter]['created']);
				$entries[$counter]['date'] = date_format($date, 'Y-m-d');
				$entries[$counter]['time'] = date_format($date, 'H:i');
	
				if($counter!=0){
					if($entries[$counter]['date']==$entries[$counter-1]['date']){
						$entries[$counter]['showdate'] = false;
					}
					else{
						$entries[$counter]['showdate'] = true;
					}
				}
				else{
					$entries[$counter]['showdate'] = true;
				}

				$counter++;
			}
			
			mysqli_stmt_close($stmt);
		}
		
		return $entries;
	}

	// tracks the click on a remarked Link and returns the url
	public function trackClick($userid,$bookmarkid){
		$timeStamp = time();
		
		// write in clicktime table
		$query = "INSERT INTO clicktime (bookmark_id, user_id, created) VALUES (?,?,?)";
		$stmt = mysqli_prepare($this->dbLink, $query);
		mysqli_stmt_bind_param($stmt, "sss", $bookmarkid, $userid, $timeStamp);
		mysqli_stmt_execute($stmt);
		mysqli_stmt_close($stmt);

		// update bookmark table
		$query = "UPDATE bookmark SET clickcount =  clickcount + 1, clicked = ? WHERE id = ?";
		$stmt = mysqli_prepare($this->dbLink, $query);
		mysqli_stmt_bind_param($stmt, "ss", $timeStamp, $bookmarkid);
		mysqli_stmt_execute($stmt);
		mysqli_stmt_close($stmt);
		
		$query = "SELECT url FROM bookmark WHERE id = ?";

		if ($stmt = mysqli_prepare($this->dbLink, $query)) {
			mysqli_stmt_bind_param($stmt, "s", $bookmarkid);
			mysqli_stmt_execute($stmt);
			mysqli_stmt_bind_result($stmt, $url);
			mysqli_stmt_fetch($stmt);
			mysqli_stmt_close($stmt);
		}
		return $url;
	}

	private function getIdByUrl($userid,$url){	
		$normalizedUrl = $this->normalizeUrl($url);
		$query = "SELECT id FROM bookmark WHERE url = ? AND user_id = ?";

		if ($stmt = mysqli_prepare($this->dbLink, $query)) {
			mysqli_stmt_bind_param($stmt, "ss", $normalizedUrl, $userid);
			mysqli_stmt_execute($stmt);
			mysqli_stmt_bind_result($stmt, $id);
			mysqli_stmt_fetch($stmt);
			mysqli_stmt_close($stmt);
		}

		return $id;
	}

	private function saveBookmark($userid,$url,$title){
		$normalizedUrl = $this->normalizeUrl($url);
		$timeStamp = time();
		$urlArray = parse_url($url);
		$domain = isset($urlArray['host']) ? $urlArray['host'] : "";
		$extension = isset($urlArray['path']) ? pathinfo($urlArray['path'], PATHINFO_EXTENSION) : "";
		$mediatype = $this->getMediaType($extension,$domain);
		
		// write in bookmark table
		$query = "INSERT INTO bookmark (url, title, user_id, domain, extension, mediatype, created, updated) VALUES (?,?,?,?,?,?,?,?)";
		$stmt = mysqli_prepare($this->dbLink, $query);
		mysqli_stmt_bind_param($stmt, "ssssssss", $normalizedUrl, $title, $userid, $domain, $extension, $mediatype, $timeStamp, $timeStamp);
		mysqli_stmt_execute($stmt);
		mysqli_stmt_close($stmt);
		
		
		$query = "SELECT id FROM bookmark WHERE url = ? AND user_id = ?";

		if ($stmt = mysqli_prepare($this->dbLink, $query)) {
			mysqli_stmt_bind_param($stmt, "ss", $normalizedUrl, $userid);
			mysqli_stmt_execute($stmt);
			mysqli_stmt_bind_result($stmt, $id);
			mysqli_stmt_fetch($stmt);
			mysqli_stmt_close($stmt);
		}
		
		// write in bookmarktime table
		$query = "INSERT INTO bookmarktime (bookmark_id, user_id, created) VALUES (?,?,?)";
		$stmt = mysqli_prepare($this->dbLink, $query);
		mysqli_stmt_bind_param($stmt, "sss", $id, $userid, $timeStamp);
		mysqli_stmt_execute($stmt);
		mysqli_stmt_close($stmt);
		
		return $id;
	}

	private function saveReBookmark($userid,$id){
		$timeStamp = time();
		
		// write in bookmarktime table
		$query = "INSERT INTO bookmarktime (bookmark_id, user_id, created) VALUES (?,?,?)";
		$stmt = mysqli_prepare($this->dbLink, $query);
		mysqli_stmt_bind_param($stmt, "sss", $id, $userid, $timeStamp);
		mysqli_stmt_execute($stmt);
		mysqli_stmt_close($stmt);

		// update bookmark table
		$query = "UPDATE bookmark SET bookmarkcount =  bookmarkcount + 1, updated = ? WHERE id = ?";
		$stmt = mysqli_prepare($this->dbLink, $query);
		mysqli_stmt_bind_param($stmt, "ss", $timeStamp, $id);
		mysqli_stmt_execute($stmt);
		mysqli_stmt_close($stmt);
		
		return $id;
	}

	public function bookmark($userid,$url,$title){
		$bookmarkid = $this->getIdByUrl($userid,$url);
		if($bookmarkid){
			if($this->saveReBookmark($userid,$bookmarkid))
				return 1;
		}else{
			if($this->saveBookmark($userid,$url,$title))
				return 0;
		}
		return 2;
	}

	private function normalizeUrl($url){
		//TODO implement cutting last / and/or other stuff
		return $url;
	}

	private function calculateRemarkVisibilityClass($number){
		switch($number) {
			 case 0: return 0;
			 case 1: return 1;
			 case 2: return 2;
			 case 3: return 4;
			 case 4: return 6;
		}
		return 8;
	}

	private function calculateClickVisibilityClass($number){
		switch($number) {
			 case 0: return 0;
			 case 1: return 1;
			 case 2: return 2;
			 case 3: return 3;
		}
		if($number<=6)
			return 4;
		if($number<=10)
			return 5;
		if($number<=15)
			return 6;
		if($number<=20)
			return 7;

		return 8;
	}

	private function getMediaType($extension,$domain){
		//mediatypes
		$probablyHTML = 1;
		$image = 2;
		$video = 3;
		$music = 4;
		switch($extension){
			case ("jpg"): return $image;
			case ("jpeg"): return $image;
			case ("gif"): return $image;
			case ("png"): return $image;
		}
		
		switch($domain){
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
