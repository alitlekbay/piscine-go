#!/bin/bash
curl https://raw.githubusercontent.com/kigiri/superhero-api/master/api/all.json | jq -jr ".[] | select(.id==$HERO_ID) | .connections | .relatives"
