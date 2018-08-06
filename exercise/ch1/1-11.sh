!/bin/sh

cut -d, -f1 misc/url_list.csv | xargs -t -I{} ./fetchall http://{}
