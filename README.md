# Make This Application work

1) clone
2) create database and update db/sqlboiler.toml
3) install sqlboiler
4) run sqlboiler.toml
    go install github.com/volatiletech/sqlboiler/v4@latest
    go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-psql@latest
    make sure go bin folder is added to path
    "https://github.com/volatiletech/sqlboiler"
5) install golang-migrate 
    https://github.com/golang-migrate/migrate/tree/master/cmd/migrate
6) run '''make migrate_up''', This will create a table schema for you
7) genrate models by '''make models_update'''
8) start the application by '''make'''