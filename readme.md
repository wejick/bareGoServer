# Bare Go Server
Bootstrapping Go application server is not hard but it's always easier if we can just use existing boiler plate to structure our app.
I keep on using standard library except for the http server, since httprouter is really fast and provide easier interface than default Go http muxer.
You can find simple app structure and hello world implementation as example.

# Getting and building
 ### Getting  
  ```
   go get github.com/wejick/bareGoServer
  ```
 ### Building
  ```
   go build 
  ```
 ### Running
  ```
   ./bareGoServer
  ```