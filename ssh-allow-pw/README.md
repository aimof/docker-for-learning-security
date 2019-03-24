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
ssh-allow-pw $ ssh -i./private_key/sample_key -p2222 root@localhost
```

## If you don't have the key, how do you login this server?

There is a way to login.

## WARNING

DON'T use this key pair.
DON'T publish it.