# SSH server allow password to login

```
docker-compose build
docker-compose up
```

## What's this?

There is a SSH server.
It has a weak rsa key pair.

Let's try to login this server.

```sh
ssh-allow-pw $ ssh -i./private_key/sample_key -p2222 localhost
```