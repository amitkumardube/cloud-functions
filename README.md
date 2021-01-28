# Cloud Functions - Getting Started
The understanding of cloud functions is incomplete without understanding of Events and Triggers described later in the course. 

<b>what are cloud functions ?</b>
* Excerpt from google cloud documentation : 
Google Cloud Functions is a serverless execution environment for building and connecting cloud services. 
With Cloud Functions you write simple, single-purpose functions that are attached to events emitted from your cloud infrastructure and services. Your function is triggered when an event being watched is fired. Your code executes in a fully managed environment. There is no need to provision any infrastructure or worry about managing any servers.

<b>Cloud Events</b>
Cloud events are things that happen in your cloud environment. These might be things like changes to data in a database, files added to a storage system, or a new virtual machine instance being created. Events occur whether or not you choose to respond to them

<b>Triggers</b>
You create a response to an event with a trigger. A trigger is a declaration that you are interested in a certain event or set of events. Binding a function to a trigger allows you to capture and act on events.
 
<b>Type of events which fires cloud functions</b>
- Http Trigger event like webhooks - can be called using an HTTP call.
- Cloud Storage event Trigger - they are attached to a GCP cloud storage and are fired when there is an event fired from bucket like adding files , updating files , deleting files.
- Cloud Pub/Sub event Trigger - They are attached to a cloud Pub/Sub event. They act as a subscriber.

<b>Imp Notes</b>
- It's important to close/stop the cloud functions once they are done processing.
- If this is not done, then they will try to run upto the timeout value setup for that function which 60 seconds by default. 
- Since you are charged only for the duration the cloud function runs, not closing a function manually will charge you till the timeout value set for that function.
 
 <b>How to close the cloud functions manually ?</b>
 
 <b>Where to go to learn more? </b>
 <p>Documentation : https://cloud.google.com/functions/docs/concepts/overview#:~:text=Google%20Cloud%20Functions%20is%20a,your%20cloud%20infrastructure%20and%20services.&text=Your%20code%20executes%20in%20a%20fully%20managed%20environment. </p>
 <p>PluralSight course : Google Cloud functions - Getting Started by James Wilson </p>