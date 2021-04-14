import jwt from 'k6/x/jwt';

const key = open("private.key", "r");

const signer = new jwt.Signer(key);

export default function () {
  const payload = {
    name: 'Eugene'
  };
  console.log(`mykey => ${signer.sign(payload, 'kid1')}`);
}
