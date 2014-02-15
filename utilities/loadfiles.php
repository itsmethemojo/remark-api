<?php	
	foreach (glob("models/*.php") as $filename) {
		require_once($filename);
		//echo $filename.'<br>';
	}

	foreach (glob("controllers/*.php") as $filename) {
      require_once($filename);
		//echo $filename.'<br>';
	}

	$pi = pathinfo(__FILE__);
	$actFile = $pi['filename'];

	foreach (glob("utilities/*.php") as $filename) {
		if($filename!='utilities/'.$actFile.'.php'){
	      require_once($filename);
			//echo $filename.'<br>';
		}
	}

?>
