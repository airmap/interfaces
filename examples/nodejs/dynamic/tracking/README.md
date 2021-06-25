This is an example of how to use dynamically loaded gRPC and dynamically-typed protobufs.

The example clients can be installed by running:
```sh
cd examples/nodejs/dynamic/tracking
npm install
```

The traffic processor client can be run by:
```
EXPORT CLIENT_ID={YOUR_CLIENT_ID}
EXPORT CLIENT_SECRET={YOUR_CLIENT_SECRET}
node processor.js
```

The traffic provider client can be run by:
```
EXPORT CLIENT_ID={YOUR_CLIENT_ID}
EXPORT CLIENT_SECRET={YOUR_CLIENT_SECRET}
EXPORT PROVIDER_ID={YOUR_AIRMAP_ISSUED_PROVIDER_ID}
node provider.js
```

The processor monitor can be run by:
```
EXPORT CLIENT_ID={YOUR_CLIENT_ID}
EXPORT CLIENT_SECRET={YOUR_CLIENT_SECRET}
node monitor.js
```

This monitor is designed to be run continously and pulls the data from the tracking API endpoint. Some error conditions are handled by a restart process. This will write a log file, **log.log** which is appended to every 1 minute. There is also a **metric** file which contains the latest value, updated every 1 minute.


