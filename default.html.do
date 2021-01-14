redo-ifchange $2.mdwn $2.bl
cmark < $2.mdwn

echo "<h1>Backlinks</h1>"

ls $2.bl | while read bl
do
	page=${bl%.*}
	echo "[${page}](${page}.html)" | cmark
done
