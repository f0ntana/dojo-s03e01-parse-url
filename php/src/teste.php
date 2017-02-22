<?php
require_once 'CleanURL.php';


class Test extends PHPUnit_Framework_TestCase

{
    public function testURL()
    {
        $test = new CleanURL();
        $return = $test->getUrl('http://www.devmt.com.br');
        $this->assertEquals('http', $return['protocol']);
    }
}