set -uex
curl -H 'Content-Type:application/json' -d '[{"english": "a", "japanese": "b", "description": "c", "examples": "abc", "similar": "similar", "tags": ["tag1", "tag2"]}]' localhost:1323/add
