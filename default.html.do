redo-ifchange $2.mdwn $2.bl
echo $2.mdwn $2.bl 1>&2
cmark < $2.mdwn

echo "<h1>Backlinks</h1>"

while read bl
do
	page=${bl%.*}
	echo "[${page}](${page}.html)" | cmark
done < $2.bl
