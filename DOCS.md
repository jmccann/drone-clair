
The Hipchat plugin sends build status messages to users and rooms. The below pipeline configuration demonstrates simple usage:

```yaml
pipeline:
  clair:
    image: Unikorn123/drone-clair
    url: http://clair.company.com
    username: johndoe
    password: mysecret
    scan_image: python:2.7
```

# Secrets

The Hipchat plugin supports reading credentials from the Drone secret store. This is strongly recommended instead of storing credentials in the pipeline configuration in plain text.

```diff
pipeline:
  slack:
    image: jmccann/drone-hipchat
    room: my-room
-   auth_token: my-auth-token
```

The above `auth_token` Yaml attribute can be replaced with the `HIPCHAT_AUTH_TOKEN` secret environment variable.

Use the command line utility to add secrets to the store:

```nohighlight
drone secret add --image=jmccann/drone-hipchat \
    octocat/hello-world HIPCHAT_AUTH_TOKEN abcd1234
```

Don't forget to sign the Yaml after making changes:

```nohighlight
drone sign octocat/hello-world
```

Please see the [Drone documentation](http://readme.drone.io/0.5/secrets-with-plugins/) to learn more about secrets.

# Secret Reference

HIPCHAT_AUTH_TOKEN
: HipChat V2 API token

# Parameter Reference

url
: HipChat server URL, defaults to `https://api.hipchat.com`

auth_token
: HipChat V2 API token; use a room or user token with the `Send Notification` scope

room
: ID or URL encoded name of the room

from: drone
: A label to be shown in addition to sender's name

notify: false
: Whether this message should trigger a user notification. See https://www.hipchat.com/docs/apiv2/method/private_message_user
