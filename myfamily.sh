HERO_ID=$HERO_ID

curl -s "https://zone01normandie.org/assets/superhero/all.json" | jq  -r '.[] | select(.id==1) | .connections.relatives | gsub("\""; "")'



