// t1meMa$heen
// 3d$h33ran

cname = "resources"

print(" ")

res0 = db.createCollection(cname);

if (res0.ok) {
    print("Collection " + cname + " ok.");

    res = db[cname].insertMany([
        {
            "id": "2819c223-7f76-453a-919d-ab1234567891",
            "externalId": "",
            "meta": {
                "created": ISODate("2010-01-23T04:56:22.000Z"),
                "lastModified": ISODate("2011-05-13T04:42:34.000Z"),
                "version": "W/\"a330bc54f0671c9\"",
                "location": "/v2/Users/2819c223-7f76-453a-919d-ab1234567891",
                "resourceType": "User"
            },
            "urn:ietf:params:scim:schemas:extension:life:2°0:User" : {
                "taxCode": "1234567890",
                "gender": "female"
            },
            "urn:ietf:params:scim:schemas:extension:enterprise:2°0:User": {
                "division": "Theme Park",
                "department": "Tour Operations",
                "manager": {
                    "§ref": "../Users/26118915-6090-4610-87e4-49d8ca9f808d",
                    "displayName": "John Smith",
                    "value": "26118915-6090-4610-87e4-49d8ca9f808d"
                },
                "employeeNumber": "701984",
                "costCenter": "4130",
                "organization": "Universal Studios"
            },
            "urn:ietf:params:scim:schemas:core:2°0:User": {
                "userName": "tfork@example.com",
                "name": {
                    "givenName": "Tiffany",
                    "middleName": "Geraldine",
                    "honorificPrefix": "Ms.",
                    "honorificSuffix": "II",
                    "formatted": "Ms. Tiffany G Fork, II",
                    "familyName": "Fork"
                },
                "title": "Electronic home entertainment equipment installer",
                "locale": "en-US",
                "emails": [
                    {
                        "value": "tfork@example.com",
                        "type": "work",
                        "primary": true
                    },
                    {
                        "value": "tiffy@fork.org",
                        "type": "home"
                    }
                ],
                "addresses": [
                    {
                        "region": "NC",
                        "postalCode": "27401",
                        "country": "USA",
                        "type": "work",
                        "primary": true,
                        "formatted": "637 Keyser Ridge Road\nGreensboro, NC 27401 USA",
                        "streetAddress": "637 Keyser Ridge Road",
                        "locality": "Greensboro"
                    },
                    {
                        "type": "home",
                        "formatted": "456 Greensboro Blvd\nGreensboro, NC 27401 USA",
                        "streetAddress": "456 Greensboro Blvd",
                        "locality": "Greensboro",
                        "region": "NC",
                        "postalCode": "27401",
                        "country": "USA"
                    }
                ],
                "x509Certificates": [
                    {
                        "value": BinData(0, "MIIDQzCCAqygAwIBAgICEAAwDQYJKoZIhvcNAQEFBQAwTjELMAkGA1UEBhMCVVMxEzARBgNVBAgMCkNhbGlmb3JuaWExFDASBgNVBAoMC2V4YW1wbGUuY29tMRQwEgYDVQQDDAtleGFtcGxlLmNvbTAeFw0xMTEwMjIwNjI0MzFaFw0xMjEwMDQwNjI0MzFaMH8xCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRQwEgYDVQQKDAtleGFtcGxlLmNvbTEhMB8GA1UEAwwYTXMuIEJhcmJhcmEgSiBKZW5zZW4gSUlJMSIwIAYJKoZIhvcNAQkBFhNiamVuc2VuQGV4YW1wbGUuY29tMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA7Kr+Dcds/JQ5GwejJFcBIP682X3xpjis56AK02bc1FLgzdLI8auoR+cC9/Vrh5t66HkQIOdA4unHh0AaZ4xL5PhVbXIPMB5vAPKpzz5iPSi8xO8SL7I7SDhcBVJhqVqr3HgllEG6UClDdHO7nkLuwXq8HcISKkbT5WFTVfFZzidPl8HZ7DhXkZIRtJwBweq4bvm3hM1Os7UQH05ZS6cVDgweKNwdLLrT51ikSQG3DYrl+ft781UQRIqxgwqCfXEuDiinPh0kkvIi5jivVu1Z9QiwlYEdRbLJ4zJQBmDrSGTMYn4lRc2HgHO4DqB/bnMVorHB0CC6AV1QoFK4GPe1LwIDAQABo3sweTAJBgNVHRMEAjAAMCwGCWCGSAGG+EIBDQQfFh1PcGVuU1NMIEdlbmVyYXRlZCBDZXJ0aWZpY2F0ZTAdBgNVHQ4EFgQU8pD0U0vsZIsaA16lL8En8bx0F/gwHwYDVR0jBBgwFoAUdGeKitcaF7gnzsNwDx708kqaVt0wDQYJKoZIhvcNAQEFBQADgYEAA81SsFnOdYJtNg5Tcq+/ByEDrBgnusx0jloUhByPMEVkoMZ3J7j1ZgI8rAbOkNngX8+pKfTiDz1RC4+dx8oU6Za+4NJXUjlL5CvV6BEYb1+QAEJwitTVvxB/A67g42/vzgAtoRUeDov1+GFiBZ+GNF/cAYKcMtGcrs2i97ZkJMo=")
                    }
                ],
                "profileUrl": "https://login.example.com/tfork",
                "preferredLanguage": "en-US",
                "active": true,
                "photos": [
                    {
                        "type": "photo",
                        "value": "https://photos.example.com/profilephoto/72930000000Ccne/G"
                    },
                    {
                        "value": "https://photos.example.com/profilephoto/72930000000Ccne/H",
                        "type": "thumbnail"
                    }
                ],
                "displayName": "Tiffany Fork",
                "userType": "Employee",
                "password": "$2a$10$5BxNICIsQHcJpCpK8OR9OeMSjVBvU4MCdT46CvBDzXp8VA4uuX6wO",
                "phoneNumbers": [
                    {
                        "value": "336-485-7643",
                        "type": "work"
                    },
                    {
                        "value": "336-485-3456",
                        "type": "mobile"
                    }
                ],
                "ims": [
                    {
                        "value": "tiffyaim",
                        "type": "aim"
                    }
                ],
                "nickName": "Tiffy",
                "timezone": "America/Los_Angeles",
                "groups": [
                ]
            },
            "schemas": [
                "urn:ietf:params:scim:schemas:core:2.0:User",
                "urn:ietf:params:scim:schemas:extension:life:2.0:User"
            ]
        },
        {
            "id": "2819c223-7f76-453a-919d-ab1234567892",
            "externalId": "",
            "meta": {
                "created": ISODate("2010-01-23T04:56:22.000Z"),
                "lastModified": ISODate("2011-05-13T04:42:34.000Z"),
                "version": "W/\"a330bc54f0671c9\"",
                "location": "/v2/Users/2819c223-7f76-453a-919d-ab1234567892",
                "resourceType": "User"
            },
            "urn:ietf:params:scim:schemas:extension:life:2°0:User" : {
                "taxCode": "1234567790",
                "gender": "female"
            },
            "urn:ietf:params:scim:schemas:extension:enterprise:2°0:User": {
                "division": "Theme Park",
                "department": "Roller Coaster Operator",
                "manager": {
                    "§ref": "../Users/26118915-6090-4610-87e4-49d8ca9f707c",
                    "displayName": "Johnny Smith II",
                    "value": "26118915-6090-4610-87e4-49d8ca9f707c"
                },
                "employeeNumber": "489107",
                "costCenter": "8260",
                "organization": "Universal Studios"
            },
            "urn:ietf:params:scim:schemas:core:2°0:User": {
                "userName": "ajames@example.com",
                "name": {
                    "givenName": "Alexandra",
                    "middleName": "Brandi",
                    "honorificPrefix": "Ms.",
                    "honorificSuffix": "III",
                    "formatted": "Ms. Alexandra B James, III",
                    "familyName": "James"
                },
                "title": "Electronic home entertainment equipment installer",
                "locale": "en-US",
                "emails": [
                    {
                        "value": "ajames@example.com",
                        "type": "work",
                        "primary": true
                    },
                    {
                        "value": "alex@james.org",
                        "type": "home"
                    }
                ],
                "addresses": [
                    {
                        "region": "NC",
                        "postalCode": "27401",
                        "country": "USA",
                        "type": "work",
                        "primary": true,
                        "formatted": "647 Keyser Ridge Road\nGreensboro, NC 27401 USA",
                        "streetAddress": "647 Keyser Ridge Road",
                        "locality": "Greensboro"
                    },
                    {
                        "type": "home",
                        "formatted": "654 Greensboro Blvd\nGreensboro, NC 27401 USA",
                        "streetAddress": "654 Greensboro Blvd",
                        "locality": "Greensboro",
                        "region": "NC",
                        "postalCode": "27401",
                        "country": "USA"
                    }
                ],
                "x509Certificates": [
                    {
                        "value": BinData(0, "MIIDQzCCAqygAwIBAgICEAAwDQYJKoZIhvcNAQEFBQAwTjELMAkGA1UEBhMCVVMxEzARBgNVBAgMCkNhbGlmb3JuaWExFDASBgNVBAoMC2V4YW1wbGUuY29tMRQwEgYDVQQDDAtleGFtcGxlLmNvbTAeFw0xMTEwMjIwNjI0MzFaFw0xMjEwMDQwNjI0MzFaMH8xCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRQwEgYDVQQKDAtleGFtcGxlLmNvbTEhMB8GA1UEAwwYTXMuIEJhcmJhcmEgSiBKZW5zZW4gSUlJMSIwIAYJKoZIhvcNAQkBFhNiamVuc2VuQGV4YW1wbGUuY29tMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA7Kr+Dcds/JQ5GwejJFcBIP682X3xpjis56AK02bc1FLgzdLI8auoR+cC9/Vrh5t66HkQIOdA4unHh0AaZ4xL5PhVbXIPMB5vAPKpzz5iPSi8xO8SL7I7SDhcBVJhqVqr3HgllEG6UClDdHO7nkLuwXq8HcISKkbT5WFTVfFZzidPl8HZ7DhXkZIRtJwBweq4bvm3hM1Os7UQH05ZS6cVDgweKNwdLLrT51ikSQG3DYrl+ft781UQRIqxgwqCfXEuDiinPh0kkvIi5jivVu1Z9QiwlYEdRbLJ4zJQBmDrSGTMYn4lRc2HgHO4DqB/bnMVorHB0CC6AV1QoFK4GPe1LwIDAQABo3sweTAJBgNVHRMEAjAAMCwGCWCGSAGG+EIBDQQfFh1PcGVuU1NMIEdlbmVyYXRlZCBDZXJ0aWZpY2F0ZTAdBgNVHQ4EFgQU8pD0U0vsZIsaA16lL8En8bx0F/gwHwYDVR0jBBgwFoAUdGeKitcaF7gnzsNwDx708kqaVt0wDQYJKoZIhvcNAQEFBQADgYEAA81SsFnOdYJtNg5Tcq+/ByEDrBgnusx0jloUhByPMEVkoMZ3J7j1ZgI8rAbOkNngX8+pKfTiDz1RC4+dx8oU6Za+4NJXUjlL5CvV6BEYb1+QAEJwitTVvxB/A67g42/vzgAtoRUeDov1+GFiBZ+GNF/cAYKcMtGcrs2i97ZkJMo=")
                    }
                ],
                "profileUrl": "https://login.example.com/ajames",
                "preferredLanguage": "en-US",
                "active": true,
                "photos": [
                    {
                        "type": "photo",
                        "value": "https://photos.example.com/profilephoto/72930000000Ccne/G"
                    },
                    {
                        "value": "https://photos.example.com/profilephoto/72930000000Ccne/H",
                        "type": "thumbnail"
                    }
                ],
                "displayName": "Alexandra James",
                "userType": "Employee",
                "password": "$2a$10$TPfDp8mAmfrAkIWTDrEuK.lSs75wsh1pdGrvPaRwxgvzpNbiPNBPm",
                "phoneNumbers": [
                    {
                        "value": "336-485-6543",
                        "type": "work"
                    },
                    {
                        "value": "336-485-9876",
                        "type": "mobile"
                    }
                ],
                "ims": [
                    {
                        "value": "alexaim",
                        "type": "aim"
                    }
                ],
                "nickName": "Alex",
                "timezone": "America/Los_Angeles",
                "groups": [
                    {
                        "display": "Employees",
                        "value": "e9e30dba-f08f-4109-8486-d5c6a331660a",
                        "§ref": "/v2/Groups/e9e30dba-f08f-4109-8486-d5c6a331660a"
                    }
                ]
            },
            "schemas": [
                "urn:ietf:params:scim:schemas:core:2.0:User",
                "urn:ietf:params:scim:schemas:extension:life:2.0:User"
            ]
        },
        {
            "schemas": [
                "urn:ietf:params:scim:schemas:core:2.0:Group"
            ],
            "id": "e9e30dba-f08f-4109-8486-d5c6a331660a",
            "urn:ietf:params:scim:schemas:core:2°0:Group": {
                "displayName": "Salinesss",
                "members": [
                    {
                        "value": "2819c223-7f76-453a-919d-ab1234567892",
                        "§ref": "/v2/Users/2819c223-7f76-453a-919d-ab1234567892"
                    },
                ],
            },
            "meta": {
                "resourceType": "Group",
                "created": ISODate("2010-01-23T04:56:22.000Z"),
                "lastModified": ISODate("2011-05-13T04:42:34.000Z"),
                "version": "W/\"a330bc54f0671c9\"",
                "location": "/v2/Groups/e9e30dba-f08f-4109-8486-d5c6a331660a"
            }
        }
    ])

    if (res.acknowledged) {
        print("Insertion of resources ok.")
    } else {
        print("Error inserting resource. Exiting.")
        quit()
    }
}

print(" ")