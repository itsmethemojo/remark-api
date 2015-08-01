<?php

class RemarkController extends BaseController{    
    
    public function initialize() {
        $this->javascript = array();
        $this->javascript[] = "script/remark.js";
        $this->javascript[] = "script/jquery-2.1.3.min.js";
        $this->javascript[] = "script/jquery.query-object.js";
        
        $this->css = array();
        $this->css[] = "style/remark.css";
        
        $this->favicon = "img/favicon.ico";
        
    }
    
    public function actionComplete(){
        //TODO retrieve userid from Parameter
        $userId = 1;
        $viewParameters['bookmarks'] = $this->model->retrieveCompleteList($userId);
        
        $this->view($viewParameters);
    }    
    
    public function actionOpen(){
        //TODO retrieve userid from Parameter
        $userId = 1;
        $bookmarkId = $this->readParameter('id');
        $viewParameters['target'] = $this->model->trackClick($userId,$bookmarkId);
        $this->disableContainer();
        $this->view($viewParameters);
    }
    
    public function actionEdit(){
        //TODO retrieve userid from Parameter
        $userId = 1;
        $bookmarkId = $this->readParameter('id');
        $bookmarkCustomTitle = $this->readParameter('customtitle');
        $this->model->setTitle($userId,$bookmarkId,$bookmarkCustomTitle);
        $this->disableContainer();
        $this->view();
    }
    
    public function actionRemark(){
        //TODO retrieve userid from Parameter
        $userId = 1;
        $url = $this->readParameter('url');
        $title = $this->readParameter('title');
        $returnStatus = $this->model->saveBookmark($userId,$url,$title);
        $this->disableContainer();
        $this->view(array("status" => $returnStatus));
        
    }
}
?>
