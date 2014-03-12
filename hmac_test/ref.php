<?php
include("warden-hmac-signer.php");
$w = new WardenHmacSigner();
$d = array();
$d['date'] = time()+$argv[3];
$r = $w->signUrl($argv[1],$argv[2],$d);
echo $r;
