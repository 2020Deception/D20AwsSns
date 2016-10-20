# D20AwsSns
Amazon Web Service Simple Notification Service Script

# Installation

First type 
```
go get github.com/aws/aws-sdk-go
```

# Instructions

from the terminal/command line :  
```
go run sns.go
```  
 -s the flag to set up a topic. Pass the topic name and aws region, in that order.  
 -r the flag to register a device. Pass the device token, platformARN and aws region, in that order.  
 -p the flag to push a notification. Pass the deviceARN (target), text and aws region, in that order.  
