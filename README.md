# Freshpaint FHIR server

We need a FHIR server/API to code against for the Data Activation project.

We chose Blaze for its ease of use, and for great documentation on its API.

See the URLs below for a test suite that runs against Blaze to make sure it
conforms to FHIR API standards. The examples show known working HTTP requests
that we can use as examples/templates.
https://touchstone.aegis.net/touchstone/conformance/detail?suite=FHIR4-0-1-Basic-Server&sVersion=6&testSystem=5e53ff030a120e4b3ee1dae1&supportedOnly=true&cb=%2fFHIR4-0-1-Basic&format=ALL&published=true
https://touchstone.aegis.net/touchstone/scriptexecution?exec=202103070728502289371123&qn=/FHIR4-0-1-Basic/A-C/Appointment/Client%20Assigned%20Id/Appointment-client-id-json#WAITING_OPER

## Blaze: Clojure FHIR server

https://github.com/samply/blaze/tree/master

We'll run Blaze via Docker, so we need to do some basic Docker setup

```sh
docker volume create blaze-data
docker run -p 8080:8080 -v blaze-data:/app/data samply/blaze:0.24
```

Next, install the `blazectl` helper utility - this is useful for controlling the
`blaze` instance.
https://github.com/samply/blaze/blob/master/docs/importing-data.md

```sh
cd ~/Downloads
curl -LO https://github.com/samply/blazectl/releases/download/v0.13.0/blazectl-0.13.0-darwin-amd64.tar.gz
tar xzf blazectl-0.13.0-darwin-amd64.tar.gz
sudo mv ./blazectl /usr/local/bin/blazectl
```

If you started with an empty install, check to make sure there are no resources
yet (your drive should be blank).

```sh
blazectl --server http://localhost:8080/fhir count-resources
```

Next, download a patient data bundle generator and generate a patient bundle
https://github.com/samply/bbmri-fhir-gen

```sh
cd ~/Downloads
curl -LO https://github.com/samply/bbmri-fhir-gen/releases/download/v0.4.0/bbmri-fhir-gen-0.4.0-darwin-arm64.tar.gz
tar xzf bbmri-fhir-gen-0.4.0-darwin-arm64.tar.gz
sudo mv ./bbmri-fhir-gen /usr/local/bin/bbmri-fhir-gen
```

Generate the test data and upload it to the blaze instance

```sh
mkdir fhir-test-data
bbmri-fhir-gen fhir-test-data
blazectl --server http://localhost:8080/fhir upload fhir-test-data
```

Test cases:
https://touchstone.aegis.net/touchstone/scriptexecution?exec=202103070728502289371123&qn=/FHIR4-0-1-Basic/A-C/Appointment/Client%20Assigned%20Id/Appointment-client-id-json#WAITING_OPER

### Other patient data generators

Synthea does not create "Appointment" resources - you can use "encounter" instead
https://github.com/synthetichealth/synthea/issues/1007

# Alternative FHIR servers

There are other open-source FHIR servers we could have used.

HAPI-FHIR is a popular option, but we decided against using it because it was
enterprise Java and felt difficult & unpleasant to work with

https://github.com/fhir-fuel/awesome-FHIR
https://www.redoxengine.com/blog/popular-open-source-fhir-libraries/
https://hapifhir.io/hapi-fhir/docs/getting_started/downloading_and_importing.html
https://github.com/FirelyTeam/fhirstarters/tree/master/java/hapi-fhirstarters-simple-server/
https://hapifhir.io/hapi-fhir/docs/server_plain/server_types.html#plain-server-facade
https://github.com/hapifhir/hapi-fhir-jpaserver-starter

# Uploading data generated from Synthea to Blaze
Follow the documentation from Synthea on how to locally generate patient test data: https://github.com/synthetichealth/synthea/wiki/Basic-Setup-and-Running

Generate data from Synthea as a transaction type: java -jar synthea-with-dependencies.jar --exporter.fhir.transaction_bundle true
Upload a directory: blazectl --server http://localhost:8080/fhir upload /PWD/output/fhir/
