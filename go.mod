module fnm-api

go 1.16

require (
	github.com/TraiOi/dbs v0.0.0
	github.com/TraiOi/switch_api v0.0.0
	github.com/TraiOi/util v0.0.0
	github.com/TraiOi/checkip v0.0.0
	github.com/TraiOi/fnm v0.0.0
	github.com/yl2chen/cidranger v1.0.2 // indirect
	go.mongodb.org/mongo-driver v1.5.0 // indirect
)

replace github.com/TraiOi/dbs => ./dbs
replace github.com/TraiOi/switch_api => ./switch_api

replace github.com/TraiOi/util => ./util

replace github.com/TraiOi/checkip => ./checkip
replace github.com/TraiOi/fnm => ./fnm
