jq -r'.[] | select(.id == 170) | .name super_hero.json' 
jq -r'.[] | select(.id == 170) | .powerstats super_hero.json' 
jq -r'.[] | select(.id == 170) | .gender super_hero.json' 