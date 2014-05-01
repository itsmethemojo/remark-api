<?php

class TodoController extends Controller
{
	public function __construct($model)
	{

		parent::__construct($model);
		$this->_setModel($model);
	}
	
	public function showList($userid)
	{
		try {
			
			$json = $this->_model->getData($userid);
                        $this->_setView("list");
			$this->_view->set('json', $json);
			$this->_view->set('title', 'TODO');
			return $this->_view->output();
			
		} catch (Exception $e) {
			echo '<h1>Application error:</h1>' . $e->getMessage();
		}
	}
        
        public function save($userid,$json)
	{
		try {
			$json = $this->_model->saveData($userid,$json);
                        $this->_setView("redirect");
			$this->_view->set('json', $json);
			$this->_view->set('title', 'TODO');
			return $this->_view->output();
			
		} catch (Exception $e) {
			echo '<h1>Application error:</h1>' . $e->getMessage();
		}
	}
}
?>