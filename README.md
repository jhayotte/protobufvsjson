#Protobuf vs JSON
This project benchmark Protobuffer vs JSON technologies.

---
To proceed I created two servers: one proto-server and one json-server reading messages received. Then there is one client sending messages in a predefine format either json either protobuf. 

#Scenario and Profiling
A lapstime of 60 seconds is given to the client to send messages to the server. Afterward we display the number of message sent and we clearly see a difference between both protocol:

![alt tag](https://github.com/jhayotte/protobufvsjson/blob/master/perfcompare-nbmessage.JPG)

The difference is due to several reasons:
- protobuf gives a better binary serialisation and therefore we consume less bandwith.
- with protobuf our server reads messages faster due to the schema and contract between the client and server. As opposed to JSON it avoids to do marshall operation to determine the type of message received. 

##Profiling CPU

Run pprof when the client is talking to the server with this command: `go tool pprof http://localhost:6060/debug/pprof/profile`

###Top command:
####Top json
![alt tag](https://github.com/jhayotte/protobufvsjson/blob/master/pprof-json.JPG)

####Top proto
![alt tag](https://github.com/jhayotte/protobufvsjson/blob/master/pprof-proto.JPG)

We clearly see that in JSON we spend a lot of CPU cycle in JSON marshal and we also see that we take less CPU with proto than with json.

### command web
Install Graphviz to visualize pprof result http://www.graphviz.org/Download_windows.php

You can then visualize a graph with all part 
 
