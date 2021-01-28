# Cloud Functions

<b>what are cloud functions ?</b>
* Excerpt from google cloud documentation : 
Google Cloud Functions is a serverless execution environment for building and connecting cloud services. 
With Cloud Functions you write simple, single-purpose functions that are attached to events emitted from your cloud infrastructure and services. Your function is triggered when an event being watched is fired. Your code executes in a fully managed environment. There is no need to provision any infrastructure or worry about managing any servers.

<b>Type of events which fires cloud functions</b>
- Http Trigger event like webhooks - can be called using an HTTP call.
- Cloud Storage event Trigger - they are attached to a GCP cloud storage and are fired when there is an event fired from bucket like adding files , updating files , deleting files.
- Cloud Pub/Sub event Trigger - They are attached to a cloud Pub/Sub event. They act as a subscriber.

<b>Imp Notes</b>
- It's important to close/stop the cloud functions once they are done processing.
- If this is not done, then they will try to run upto the timeout value setup for that function which 60 seconds by default. 
- Since you are charged only for the duration the cloud function runs, not closing a function manually will charge you till the timeout value set for that function.
 
 <b>How to close the cloud functions manually ?</b>