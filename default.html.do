redo-ifchange $2.mdwn $2.bl
cmark < $2.mdwn

echo "<h1>Backlinks</h1>"

grep $2 *.bl | while IFS=: read -r a b
do
	echo $a backlinks to $b >&2
	echo "[${a%.*}](${a%.*}.html)" | cmark
done
