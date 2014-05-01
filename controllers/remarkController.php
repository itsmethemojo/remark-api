<?php

class RemarkController extends Controller
{
	public function __construct($model)
	{

		parent::__construct($model);
		$this->_setModel($model);
	}
	
	public function unfilteredList($userid)
	{
		try {
			
			$items = $this->_model->getUnfilteredList($userid);
                        $this->_setView("unfilteredlist");
			$this->_view->set('items', $items);
			$this->_view->set('title', 'reMARK');
			return $this->_view->output();
			
		} catch (Exception $e) {
			echo '<h1>Application error:</h1>' . $e->getMessage();
		}
	}

	public function trackClick($userid,$bookmarkid)
	{
		try {
			
			$url = $this->_model->trackClick($userid,$bookmarkid);
                        $this->_setView("openlink");
			$this->_view->set('url', $url);
			return $this->_view->output();
			
		} catch (Exception $e) {
			echo '<h1>Application error:</h1>' . $e->getMessage();
		}
	}

	public function bookmark($userid,$url,$title)
	{
		try {
			
			$id = $this->_model->bookmark($userid,$url,$title);
                        $this->_setView("saverequest");
			$this->_view->set('id', $id);
			return $this->_view->output();
			
		} catch (Exception $e) {
			echo '<h1>Application error:</h1>' . $e->getMessage();
		}
	}
	
	
	
}