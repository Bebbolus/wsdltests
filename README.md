# wsdltests
a set of tests run to make golang use wsdl and then consume the service

# background
found several projects that try to get input as wsdl and generate golang code to consume the service
https://github.com/droyo/go-xml
https://github.com/fiorix/wsdl2go
https://github.com/hooklift/gowsdl

## refers
This article describe how to do without any libs
https://medium.com/eaciit-engineering/soap-wsdl-request-in-go-language-3861cfb5949e

## tools
This is a wsdl from free online soap service
http://www.dneonline.com/calculator.asmx?wsdl

# conclusion
Found that this is the most followed project
https://github.com/hooklift/gowsdl

basically it works faster then all the others:
you get the wsdl, run the command gowsdl yourwsdlfile, it generates a file you have to include into your package, then in your main you can call the service function passing the values.
Inspecting the generated file you can see all the property and method you can apply to the request.

You can find the examples in this repo
