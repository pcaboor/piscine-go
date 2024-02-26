HERO_ID=1 

curl -s "https://zone01normandie.org/assets/superhero/all.json" | jq -r --arg id "$HERO_ID" '.[] | select(.id==$id) | .connections.relatives | gsub("\""; "")'