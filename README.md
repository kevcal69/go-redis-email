# go-redis-email
Simple mail sender built with go, redis with vuejs frontend

# Install
Pull this project in your go workspace
```Bash
go get github.com:kevcal69/go-redis-email.git
```


### Compile templates
```Bash
# make sure you have npm
cd pkg/templates
npm install
npm run
```
### Set your environment variable for mailgun 
```Bash
# Example
export MAILGUN_API_URL=https://api.mailgun.net/v3/sandboxbacSomethingkeys.mailgun.org/messages
export MAILGUN_API_KEY=mailgun-api-key123
```
### Run app 
```Bash
# run app/main.go
go run app/main.go
```

### Note: make sure you have redis server running at 
```Bash
# Redis server 
> 0.0.0.0:6379
```