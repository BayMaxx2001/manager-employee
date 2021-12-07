## Problem
1. Manager Employees - Teams
2. Connect to Database and save information
4. Write API for employees and teams 
## Getting Started
 1. Install MongoDB 
 2. Create file `.env` in folder `employee/configs` for employee and `team/configs` for team
 3. Configure features in file `.env` the following format in file `env.example`
 4. We must install package: `tkanos/gonfig`,  `github.com/go-chi/chi`, `go-sql-driver/mysql` by the following:
``` console
go get github.com/go-chi/chi
go get github.com/go-sql-driver/mysql
go get github.com/tkanos/gonfig
```
 5. Before running the program, you must set up the MongoDB: 
 	5.1 Schema: employee
	5.2 Collection: employee
	5.3 Schema: team
	5.4 Collection: team
 
 6. Use command:  `go run employee/cmd/httpapi/main.go` run employee serve 
 7. Use command:  `go run team/cmd/httpapi/main.go` run team serve 
 8. API :
	* Get list all employees: 
	```url
        localhost:8282/employees (GET)
	```
	* Get list all teams:
	```url
        localhost:8181/teams (GET)
	```
	* Search employee by id:
	```url
	    localhost:8282/employee/{id} (GET)
	```
    * Search team by id:
	```url
	    localhost:8181/team/{id} (GET)
	```
    *  Create employee: 
    ``` url
        localhost:8282/employee (POST)
    ```
    Example Body: 
	```json
	{
	    "name" : "hien",
	    "gender" :2, 
	    "dob" : "2001-11-11"
	}
	```
    *  Create team: 
    ``` url
        localhost:8181/team (POST)
    ```
    Example Body:
    ```json
	{
	     "name": "pentest"
	}
	```
    * Update employee:
    ``` url
        localhost:8282/employee/{id} (PUT)
    ```
    * Update team:
    ``` url
        localhost:8181/team/{id} (PUT)
    ```
    * Delete employee: 
    ``` url 
        localhost:8282/employee/{id} (DELETE)
    ```
    * Delete team: 
    ``` url 
        localhost:8181/team/{id} (DELETE)
    ```

    * Delete/Add employee join to team:
    ```url 
        localhost:8282/api/v1/event/employee-team (POST)
    ``` 
	* Example Body: 
	``` json 
	{
	    "eid":"5b23489a-d6c7-42f6-9a41-65edd53aaf6f",
	    "tid":"6f31a96c-1981-460f-a387-413c655b9edb"
	}
	```

