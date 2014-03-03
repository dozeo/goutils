<?php
include("warden-hmac-signer.php");
$w = new WardenHmacSigner();
$r = $w->signUrl($argv[1],$argv[2]);
echo $r;
