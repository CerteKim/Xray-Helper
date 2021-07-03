package main

var upstatConfig = `# {{.Name}} {{.Description}}
description     "{{.Description}}"
author          "Pichu Chen <pichu@tih.tw>"
start on runlevel [2345]
stop on runlevel [016]
respawn
#kill timeout 5
exec {{.Path}} {{.Args}} >> /var/log/{{.Name}}.log 2>> /var/log/{{.Name}}.err
`
