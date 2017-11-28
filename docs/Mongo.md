# Query

| **Query** | **Valid** |
|-----------|-----------|
|```db.resources.find({ data : { $elemMatch : { nickName : "Babbi" }}})``` | ok |
|```db.resources.find({ data : { $elemMatch : { $and:[ { nickName : "Babbi" }, { _urn : "urn:ietf:params:scim:schemas:extension:enterprise:2.0:User"  } ]}    }} )``` | ok |
|```db.resources.find({ data : { $elemMatch : { $and:[ { "name.familyName" : "Jensen" }, { _urn : "urn:ietf:params:scim:schemas:extension:enterprise:2.0:User"  } ]}    }} )``` | ok |
|```db.resources.find({ data : { $elemMatch : { "name.familyName" : "Jensen" }    }} )``` | ok |