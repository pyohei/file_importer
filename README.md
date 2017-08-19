# Sleep Cycle data converter

Convert [Sleep Cycle Exported data](https://s.sleepcycle.com/) into MySQL data.
This tool is for myself, and if you want to use it, be careful when using. 

## Usage

1. Install golang and MySQL
1. Install this repository
1. Create database with `sql/createtable.sql`
1. Edit `main.go`. Change database connection information.
1. Execute script with `go run main.go filename`.  
   You set target file as first argument.

## Licence

* MIT
