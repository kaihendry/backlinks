redo-ifchange $2.mdwn
# Scan the changed file for links and add link to the end of the file it points to
# https://salsa.debian.org/debian/ikiwiki/-/blob/debian/master/IkiWiki/Plugin/link.pm#L136
grep -HoP "\[(.*?)\]\((.*?)\)" $2.mdwn | while IFS=: read -r a b
do
	echo Backlinks $a $b 1>&2
	page=${b##*\(} page=${page%)*}
	echo Add to ${page%.*}.bl, the back link to $a 1>&2
	mkdir -p bl/${page%.*}.bl || true
	touch bl/${page%.*}.bl/${a%.*} 1>&2
done
