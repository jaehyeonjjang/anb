<?php

$search_array = array(
    'user' => array('!loginid', '!email', array('level', '@name', '@loginid', '@email'), array('level'), array('company', 'status')),
    'findpasswd' => array('!link'),
    'sendsms' => array('status'),
    'sendmail' => array('status'),
    'image' => array('apt', array('apt', '<level'), array('apt', 'parent')),
    'work' => array('image'),
    'apt' => array('@name', array('company', 'status')),
    'aptuser' => array('apt', array('apt', 'level'), array('!apt', 'user'), 'user'),
    'aptuserlist' => array('user', 'apt', array('apt', 'aptlevel'), array('!apt', 'user'), array('user', 'aptstatus')),
    'data' => array(array('apt', 'image'), 'apt',  array('image', 'name', 'imagename')),
    'status' => array('type', 'company', array('!company', 'type', 'name'), array('company', 'type')),
    'statuscategory' => array('type', 'company'),
	'imagefloor' => array(array('!image', 'name', 'imagename'), '!target'),
	'report' => array('status'),
    'picture' => array('apt')
    );

$alias_array = array(
    );

$increase_array = array(
    );

$delete_array = array(
    'aptuser' => array('apt'),
    'image' => array('apt'),
    'data' => array(array('image', 'user'), 'image'),
    'sync' => array('image')
    );

$update_array = array(
    'apt' => array(array('where' => 'id', 'field' => array('report'))),
    'report' => array(array('where' => 'id', 'field' => array('status')))
    );

$sum_array = array(
    );

$sum2_array = array(
    );

$cache_array = array();
$cache_search_array = array(
    );
?>
