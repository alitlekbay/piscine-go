#!/bin/bash
curl https://raw.githubusercontent.com/kigiri/superhero-api/master/api/all.json | jq -r ".[] | select(.id==$HERO_ID) | .connections | .relatives"
