<?PHP
namespace App;

use PHPUnit\Framework\TestCase;
require('../src/CleanUrl.php');

class CleanUrlTest extends TestCase
{

	public function testURL()
    {
        $test = new CleanUrl();
        $return = $test->getUrl('http://www.devmt.com.br');
        $this->assertEquals('http', $return['protocol']);
    }

}
