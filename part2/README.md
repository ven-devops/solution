function.go is a pseudo code how the cloud function would work.
The flow is as follows:
1. List subnetworks
2. Store response from "list" operation to DB.

Further improvement:
After retrieving the json response, insert into the DB only the information which is needed.
Authentication to both GCP and DB should come from Env vars.