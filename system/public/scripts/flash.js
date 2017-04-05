function flash($class, $string) {
	$class.html($string);
	$class.show().delay(5000).fadeOut();
}

