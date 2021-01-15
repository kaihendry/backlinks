redo-ifchange $2.mdwn
# Scan the changed file for links and add link to the end of the file it points to
# https://salsa.debian.org/debian/ikiwiki/-/blob/debian/master/IkiWiki/Plugin/link.pm#L136
grep -HoP "\[(.*?)\]\((.*?)\)" $2.mdwn | while IFS=: read -r a b
do
	page=${b##*\(} page=${page%)*}
	echo Backlinks $a $page 1>&2
	mkdir -v ${page%.*}.bldir || true
	touch ${page%.*}.bldir/${a%.*}
	echo Added to ${page%.*}.bldir, the back link to $a
done
