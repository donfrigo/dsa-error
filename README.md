# dsa-error
A small repo that illustrates a problem with DSA key parsing.

I have created the DSA keys using the following commands 

```openssl dsaparam -out params.pem 3072
openssl gendsa -out key.pem params.pem
openssl req -new -key key.pem -out req.pem
openssl x509 -req -in req.pem -signkey key.pem -out certificate.cer
````

Upon running the code I get the following error:

 ```panic: tls: failed to parse private key```
 
 This repo was created to illustrate the issue I have on [Stackoverflow](https://stackoverflow.com/questions/61308905/parsing-dsa-key-with-golang-into-tls-config-object?noredirect=1#comment108459960_61308905)
