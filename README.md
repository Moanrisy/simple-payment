# Simple Payment app in Go with gin web framework

## Step to run
### Clone this repository
``` sh
git clone git@github.com:Moanrisy/simple-payment.git
```

### Setup postgres database
1. Create database with name simple_payment in postgres
2. Run query in ddl.sql from line 1 to line 41

### Setup env variable 
1. Edit Makefile file to update your env variable (db user, pass, etc)

### Run project
1. Run this command in terminal to run

``` sh
make run
```

2. Open http://localhost:8080/swagger/index.html#/

OR if you prefer postman

2. Import simple-payment-go.postman_collection.json in postman
