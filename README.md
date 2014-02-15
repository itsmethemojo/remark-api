reMARK
======

simple bookmarking Tool consisting of a webpage (php,mysql) and a firefox extension

Bookmarking is easy. The maintenance of the own bookmark library is the pain in the ass. So why not just let the bookmarks sort themselves. This bookmark tool counts the times a site was bookmarked by you. So if you don't know if this site is already in your library, just bookmark it (again maybe). Same Situation with the clicks on the bookmarked sites. And over the time your important bookmarks will stand out clearly without you wasting your time creating any lists.

i used the php-login-advanced for authentification https://github.com/panique/php-login-advanced
therefore see the installation instructions

The firefox extension is a really ugly adaption of a tutorial I found. Maybe there is someone out there who understands and likes the fundamental priciples of a firefox extensions enough to help me out with that. ;)

v0.1
====

I already wrote a version with more features, but I decided to refactor it. So here is the first draft.

 - bookmark/remark sites via the firefox extension
 - list all bookmarks
 - count remarks/clicks
 - installation instructions

coming Up
=========
 
 - system requirements catalog
 - database access refactoring
 - editing bookmarks
 - tagging
 - filter/sort list
 - individual statistics

Installation
============

 - use the database.sql to create (database and) tables
 - modify the config/template_db.php and save it as db.php
 - modify the firefox_extension/custom-toolbar-button@example.com/chrome/template_button.js and save it as button.js
 - pack the content of the firefox_extension/custom-toolbar-button@example.com to a zip, change the extension zip to xpi and import the extension in your firefox
 - checkout the https://github.com/panique/php-login-advanced in the subfolder php-login-advanced
 - follow the instructions here but skip the sql scripts https://github.com/panique/php-login-advanced#installation-quick-setup
 - you need to touch two files in the php-login-advanced folder to get this working


 **php-login-advanced/index.php**

> if ($login->isUserLoggedIn() == true) {

> // the user is logged in. you can do whatever you want here.

> // for demonstration purposes, we simply show the "you are logged in" view.

> **header('Location: http://localhost/reMARK/index.php');**

> } else {


 **php-login-advanced/classes/Login.php**

> private function loginWithSessionData(){

> $this->user_name = $_SESSION['user_name'];

> $this->user_email = $_SESSION['user_email'];

> **$this->user_id = $_SESSION['user_id'];**

additionally add this function

> public function getUserid(){

> return $this->user_id;

> }
