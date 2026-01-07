<?php

$filename = $argv[1];
if ($filename == '') {
	$filename = './periodic-00.hml';
}

$contents = file($filename);

$rows = count($contents);
$tab = 0;

$ret = '';
for ($i = 0; $i < $rows; $i++) {

    $data = explode('<', trim($contents[$i]));

    if (count($data) == 1) {
        echo $contents[$i];
        continue;
    }

    $flag = false;
    $postFlag = false;

    for ($j = 1; $j < count($data); $j++) {
        $line = $data[$j];

        $prefix = substr($line, 0, 1);
        $postfix = substr($line, -2);

        if ($prefix != '/') {
            if ($flag == true) {
                echo "\n";
                $flag = false;
            }
            //echo str_repeat(' ', $tab * 4);
        } else {
            if ($postFlag == true) {
                echo "\n";
            }
        }
        echo '<' . $line;

        if ($prefix != '/') {
            $flag = true;
            $tab++;
        }


        if ($prefix == '/' || $postfix == '/>') {
            $tab--;
        }

        if ($postfix == '/>') {
            $postFlag = true;
        } else {
            $postFlag = false;
        }
    }
}

?>
