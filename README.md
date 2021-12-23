# giftcode-api

Install Golang:
  apt-get install golang

cd into the giftcode-api directory

Set up Go dependencies: 
  go get
  
  
  
Run the mysql createscript (replace <myuser> with your actual username or root):
  mysql -u <myuser> -p < createscript.sql

Note: Don't Forget to create a user and then grant it the right privileges on the "giftcode_api_db" database
