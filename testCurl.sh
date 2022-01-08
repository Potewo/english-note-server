set -uex
curl -H 'Content-Type:application/json' -d '[{"english": "a", "japanese": "b", "description": "c", "examples": "abc", "similar": "similar", "tags": ["tag1", "tag2"]}]' -i localhost:1323/add
curl -H 'Content-Type:application/json' -i localhost:1323/get
# curl -H 'Content-Type:application/json' -d '[{"enlish": "a", "japanese": "b", "description": "c", "examples": "abc", "similar": "similar", "tags": ["tag1", "tag2"]}]' -i localhost:1323/add

