# Readme
## This is a very simple code to fetch the gitpr based on the status

## how to setup
- git clone <url>
- move the Dockerfile out of gitpr directory, after this your directory structure will look like as below
-

    ├── Dockerfile
    └── gitpr
        ├── Dockerfile
        ├── cmd
        │   └── main.go
        ├── gitpr
        ├── gitpr.exe
        └── pkg
- Now build the image
    
    ``` docker build -t go_lang . ```  
- Run the container

    ``` docker run -t --name gitpr -d go_lang  ```
- Go inside the container

    ``` docker exec -it gitpr  bash  ```
- now you are inside the container, go to codebase to run the app
    ``` 
    $ /usr/local/go/src# cd cmd/gitpr/
    $ /usr/local/go/src# ls
    main.go
    $ go run main.go
    Title: Bump rack from 2.0.3 to 2.2.3.1
    Url: https://api.github.com/repos/basecamp/local_time/pulls/121
    CreatedAt: 2022-05-27T17:41:32Z
    ----------------------
    Title: Bump nokogiri from 1.8.0 to 1.13.6
    Url: https://api.github.com/repos/basecamp/local_time/pulls/119
    CreatedAt: 2022-05-18T20:53:29Z
    ----------------------



