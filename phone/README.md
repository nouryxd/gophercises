# Phone

## Postgresql

[Arch Wiki](https://wiki.archlinux.org/title/PostgreSQL)

Install postgresql
    $ yay postgresql

Create DB as postgres user
    $ sudo -iu postgres
    [postgres]$ createdb coolDatabaseName

Connect to DB
    $ psql -d coolDatabaseName
