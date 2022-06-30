# mytheresa coding assignment

This project is an implementation to a basic REST api to get products. This API should also support filtering by category as well as price as stated in the problem statement. 

The overall project is built using [go-gin framework](https://github.com/gin-gonic/gin). Its build using gin as it's much faster than traditional go routers
by 40 times. Also, having a framework provides better maintainability to the project as we can focus more on solving the 
problem than maintaining or creating boiler-plates. It's not much different than traditional go so the reader of the code shouldn't
have difficulty reading the code. 

The overall project is divided into **3 packages**. The main reason for doing so is to provide better maintainability of code
and make it easier for any developer to add changes with ease. The overall code strongly follow Open-Close principle for software programming. Which simply means Software entities (classes, modules, functions, etc.) should be open for extension, but closed for modification.
We could add any new filtering condition as well as new discount logic without affecting the existing one's.
[Effective GO](https://go.dev/doc/effective_go) coding construct is also followed throughout the code.  
**Note:** You would find extensive comments throughout the code which would also explain my reason fo choosing this design.

1. **types** : For defining all the types required in the project. 
2. **filter** : For maintaining any filtering condition. 
   1. Example: We would like to filter by category as one of the requirements, Thus I added a separate filter named by_category.go which will provide the simple implementation of one filter individually.
3. **discount**: All the discount related logic is kept in here.
**Assumptions:**
1. Strictly following the assumptions defined in problem statement.
2. I am maintaining a simple json file as the product data store which is kept in memory at run time. Also introducing a DB would introduce
create extra dependencies like installation.
3. Some test cases would need to be changed if we make changes to the data store.
4. I am assuming GO lang is installed correctly on the machine which will try to execute the binaries. 
5. If you are cloning this repo, then you might need to change the permissions of the file. 
   1. for any linux :
      1. `chmod +x mytheresa-darwin-amd64`
   2. For windows, try executing the files. I didn't test on Windows as I don't have a windows operating system.

#**Commands**

**How to run from console on a linux based machine only:**
1. Clone this directory
2. Move into the directory using `cd mytheresa`
3. `chmod +x mytheresa-darwin-amd64` // Permissions are internal to an operating system and I can't control the run time unless permissions are given for a file to be executable.
4. `./mytheresa-darwin-amd64` 


#**Test cases**
1. All the packages contain separate test cases
2. To run all the test cases you could use below command:
   1. `go test ./...`
   2. Which will result in below kind of output
   3. ``` 
      ok      github.com/anujpradhaan/mytheresa       1.246s
      ok      github.com/anujpradhaan/mytheresa/discount      1.318s
      ok      github.com/anujpradhaan/mytheresa/filter        0.794s
      ?       github.com/anujpradhaan/mytheresa/types [no test files]
   ```
3. You could ignore the `github.com/anujpradhaan/mytheresa/types` directory as it doesn't have any logic. Thus, no test cases. 
   1. the statement "no test file" only mean the package doesn't have test file. **It's not a failure of test cases.**
4. Or if you prefer checking all the test cases being executed, then you could try:
   1. `go test -v ./...` it will run all the test cases in verbose mode.
