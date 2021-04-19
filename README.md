# JWT xk6 extension

kx6 extension which provides posibility working with JWT.

> **WARNING** This is only PoC.

## Prerequisites & Build own k6

Very recommend to read https://k6.io/blog/extending-k6-with-xk6/#building-the-extension-with-xk6

## Usage

```
import jwt from 'k6/x/jwt';

// a pthe to the private key
const privateKeyPath = '/keys/private.key';

const signer = new jwt.Signer(privateKeyPath);

const payload = {
  foo: 'bar'
};

// key ID
const keyId = 'kid1';

const token = toolSigner.sign(payload, keyId);

```