#Protobuf vs JSON
This project benchmark Protobuffer vs JSON technologies.

---
To proceed I created two servers: One proto-server and one json-server reading messages received in the same way. Then there is one client sending messages in a predefine format either json either protobuf. 

##scenario
A lapstime of 60 seconds is given to the client to send messages to the server. Afterward we display the number of message sent and we clearly see a difference between both protocol:
![alt tag](https://github.com/jhayotte/protobufvsjson/blob/master/perfcompare-nbmessage.JPG)

The difference is due to several reasons:
- protobuf gives a better binary serialisation and we consume less bandwith.
- our server reads messages faster due to the schema and contract between the client and server. As opposed to JSON it avoids to do reflection to determine the type of message received. 

##Profiling CPU

Run our profiling tool when the client is talking to the server with this command: `go tool pprof http://localhost:6060/debug/pprof/profile`

Install Graphviz to visualize pprof result http://www.graphviz.org/Download_windows.php

the top command shows

