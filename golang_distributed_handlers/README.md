
A sketch of disributed handlers:

* The handlers are for a basic websever
* Each endpoint will confirm to a JSON REST api interface
* Each handler will have conform to the same interface

(The distributed bit:)
* Each handler can be _deployed_ in an arbitrary fashion (at least that's the ultimate goal):
  * Any number of handlers may be placed in a single executable
  * That exectuable will be run in it's own k8s instance (probably use tilt to simulate this)
  * Each executable will need to setup it's own websever (with routing) to handle the requests for the endpoint that it serves
    * The setup of each exectuable must be seperate from the handlers themselves
    * **Ultimately** all of this would be automated somehow - probably with a go generator run from a configuration

* All the endpoints should be available as if they were hosted by a single webserver

