<?php
class CleanURL
{
	public function getUrl($url)
	{
		$protocol = explode('://', $url);
		$host = explode('.', $protocol[1]);
		$return = array(
			'protocol' => $protocol[0],
			'host' => $host[0],
		);
		if ($protocol[0] == 'http') {
			$return['domain'] = $host[1] . '.' .  $host[2];
		} else {
			$domain = explode('@', $protocol[1]);
			$return['domain'] = $protocol[1];
			$user_pass = explode('%', $protocol[0]);
			$return['user'] = $user_pass[0];
			$return['pass'] = $user_pass[1];
		}
		return $return;
	}
}