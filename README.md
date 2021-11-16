## Problem
1. Manager Employees - Teams
2. Connect to Database and save information
4. Write API for employees and teams 
## Getting Started
 1. Install MySQL 
 2. Create file `dev_config.json` in folder `pkg/config`  
 3. Configure features in file `config/dev_config.json` the following format: 
	 ``` json
	 { 
		 "DB_USERNAME": <YOUR_USERNAME>,
		 "DB_PASSWORD": <YOUR_PASSWORD>,
		 "DB_PORT": "3306",
		 "DB_HOST": "127.0.0.1",
		 "DB_NAME": "manager_employee"
	 }
	 ```
 4. We must install package: `tkanos/gonfig`,  `github.com/go-chi/chi`, `go-sql-driver/mysql` by the following:
``` console
go get github.com/go-chi/chi
go get github.com/go-sql-driver/mysql
go get github.com/tkanos/gonfig
```
 5. Before running the program, you must set up the database by running the command: 
 
 	`go run build/database/build_db.go`
 
 6. Use command:  `go run cmd/web/*.go` run program 
 7. API :
	* Get list all employees: 
	```url
        localhost:8080/employees (GET)
	```
	* Get list all teams:
	```url
        localhost:8080/teams (GET)
	```
	* Search employee by id:
	```url
	    localhost:8080/employee/{id} (GET)
	```
    * Search team by id:
	```url
	    localhost:8080/team/{id} (GET)
	```
    *  Create employee: 
    ``` url
        localhost:8080/employee (POST)
    ```
    *  Create team: 
    ``` url
        localhost:8080/team (POST)
    ```
    * Update employee:
    ``` url
        localhost:8080/employee/{id} (PUT)
    ```
    * Update team:
    ``` url
        localhost:8080/team/{id} (PUT)
    ```
    * Delete employee: 
    ``` url 
        localhost:8080/employee/{id} (DELETE)
    ```
    * Delete team: 
    ``` url 
        localhost:8080/team/{id} (DELETE)
    ```

    * Delete/Add employee join to team:
    ```url 
        localhost:8080/employee/{idEmp}/team/{idTeam} (POST/DELETE)
    ``` 


