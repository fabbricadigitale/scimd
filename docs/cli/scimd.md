## scimd

SCIMD is ...

### Synopsis

...
	
	SCIM 2 RFC - published under the IETF - defines an open API for managing identities.
Complete documentation is available at ...

```
scimd [flags]
```

### Options

```
  -c, --config string                    the path of directory containing the configuration resources
      --debug                            wheter to enable or not the debug mode
      --enable-self                      whether to enable or not the creation of an endpoint for the authenticated entity
  -h, --help                             help for scimd
      --page-size int                    the page size the server has to use (default 10)
  -p, --port int                         port to run the server on (default 8787)
  -s, --service-provider-config string   the path of the service provider config to use
      --storage-coll string              the storage's collection name (default "resources")
      --storage-host string              the storage's address (default "0.0.0.0")
      --storage-name string              the storage's database name (default "scimd")
      --storage-port int                 the storage's port (default 27017)
      --storage-type string              type of storage to use (default "mongo")
```

### SEE ALSO

* [scimd get-config](scimd_get-config.md)	 - Get the default configuration
* [scimd get-service-provider-config](scimd_get-service-provider-config.md)	 - Get the default service provider configuration
* [scimd print-config](scimd_print-config.md)	 - Print the current configuration
* [scimd version](scimd_version.md)	 - Print the version number of scimd

