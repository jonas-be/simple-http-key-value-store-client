# Simple Http key value store client

This application is a cli client for the [simple-http-key-value-store](https://github.com/jonas-be/simple-http-key-value-store).

> **Note**
> This project was mainly for learning Go. Especially to write a simple cli app and learn the `jarcoal/httpmock` library. 


## How to use

```./simple-http-key-value-store -m=METHOD -key=KEY -value=VALUE```

### -m
Specify the http method to use
 - `get`
 - `put`
 - `del` or `delete`

### -key
Specify the key 

### -value
Specify the value, if you are using the put method