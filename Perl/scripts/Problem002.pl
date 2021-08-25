sub f_nthTerm {
    my $n = @_[0];
    my $a = 0; my $b = 1;
    for( $i = 0; $i < $n; $i++) {
        $b += $a;
        ($a, $b) = ($b, $a);
    }
    return $a;
}


$start  = time();
($i, $ans) = (1, 0);
$k = f_nthTerm($i);

while ($k < 4000000){
    $k = f_nthTerm($i);
    if ($k % 2 == 0){ $ans += $k; }
    $i++;
}

$exc_time = time() - $start;
print "\nAnswer: $ans\n";
print "Time Taken: $exc_time seconds\n\n";
