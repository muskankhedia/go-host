# go-host
Hosting a golang webapp using Heroku

## To host a Golang web server using heroku:

1. Build the web application using command `go build -o bin/go-host -v .`
2. Create a Procfile with the details
3. Create the other necessary files such as `heroku.yml`, `go.mod`, `app.json`.
4. Run `heroku local` to test in localhost:5000
5. Run `heroku create`
6. Run `git push heroku master`
7. Run `heroku open`

## Runnig the go-host app
```sh
  $ git clone https://github.com/muskankhedia/go-host.git
  $ cd go-host
  $ go build -o bin/go-host -v .
  $ heroku local
  $ heroku create
  $ git push heroku master
  $ heroku open
 ```

Reference: [https://devcenter.heroku.com/articles/getting-started-with-go#set-up](https://devcenter.heroku.com/articles/getting-started-with-go#set-up)
