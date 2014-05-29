reMARK v0.2
===========

simple bookmarking tool consisting of a webpage (php,mysql) and a firefox extension

Bookmarking is easy. The maintenance of the own bookmark library is the pain in the ass. So why not just let the bookmarks sort themselves. This bookmark tool counts the times a site was bookmarked by you. So if you don't know if this site is already in your library, just bookmark it (again maybe). Same Situation with the clicks on the bookmarked sites. And over the time your important bookmarks will stand out clearly without you wasting your time creating any lists.

Used as a everyday tool, I thought it might be interesting to add some self managing features, i have never found the correct software/apps for. I am currently thinking about a todo list, contacts/birtdays and reminders. Maybe a notification bar System and/or email notification. Let's see.

I used the php-login for authentification https://github.com/panique/php-login

The firefox extension is a really ugly adaption of a tutorial I found. Maybe there is someone out there who understands and likes the fundamental priciples of a firefox extensions enough to help me out with that. ;)

v0.1
====

I already wrote a version with more features, but I decided to refactor it. So here is the first draft.

 - bookmark/remark sites via the firefox extension
 - list all bookmarks
 - count remarks/clicks
 - installation instructions

v0.2
====

 - changed login application to the newest php-login https://github.com/panique/php-login
 - added a todo list feature
 - refactored installation process and added new install instructions
 

coming up
=========
 
 - database access refactoring
 - editing bookmarks
 - tagging
 - filter/sort list
 - individual statistics
 - birtdays timeline / contact list
 - simple reminders/events
 - notification bar

installation
============

 - delete the "template_" from the files in the config folder and modify the content
 - run all the sql scripts in the _installation directory
 - modify the firefox_extension/custom-toolbar-button@example.com/chrome/template_button.js and save it as button.js
 - pack the content of the firefox_extension/custom-toolbar-button@example.com to a zip, change the extension zip to xpi and import the extension in your firefox
 - delete the php-login folder
 - make sure you have composer installed (howto http://www.dev-metal.com/install-update-composer-windows-7-ubuntu-debian-centos/)
 - run sudo composer create-project panique/php-login [YOUR_PATH]/reMARK/php-login dev-master (howto http://www.dev-metal.com/install-php-login-nets-4-full-mvc-framework-login-script-ubuntu/)
 - checkout the php-login folder again and overwrite the files
