from hashlib import sha256
from os import urandom
from binascii import unhexlify, hexlify


def sha256_hash(input):
    return sha256(input).hexdigest()


# returns hex encoded binary random data of size n bytes
def random_value(n_bytes=32):
    # urandom takes in bytes so by default we want 256 bits == 32 bytes to be generated
    return hexlify(urandom(n_bytes))


def generate_lamport_keys(n=256):
    print("Generating private and public keys")
    public_key = []
    private_key = []
    for n in range(n):
        first, second = random_value(), random_value()
        private_key.append((first, second))
        public_key.append((sha256_hash(first), sha256_hash(second)))

    return public_key, private_key


def sign_message(message, private_key):
    print("Signing message with a private key")
    signature = []
    sha = sha256_hash(message)
    bin_lmsg = unhexlify(sha)
    z = 0
    for x in range(len(bin_lmsg)):
        l_byte = bin(bin_lmsg[x])[2:]
        while len(l_byte) < 8:  # pad the zero's up to 8
            l_byte = '0' + l_byte
        for _ in range(0, 8):
            signature.append(private_key[z][0 if l_byte[-1:] == '0' else 1])
            l_byte = l_byte[:-1]
            z += 1
    return signature


def verify_message_signature(signature, message, pub):  # verify lamport signature
    print("Veryfing message with a signature and public key")
    bin_lmsg = unhexlify(sha256_hash(message))
    verify = []
    z = 0

    for x in range(len(bin_lmsg)):
        # generate a binary string of 8 bits for each byte of 32/256.
        l_byte = bin(bin_lmsg[x])[2:]

        while len(l_byte) < 8:  # pad the zero's up to 8
            l_byte = '0' + l_byte

        for _ in range(0, 8):
            if l_byte[-1:] == '0':
                verify.append((sha256_hash(signature[z]), pub[z][0]))
            else:
                verify.append((sha256_hash(signature[z]), pub[z][1]))
            l_byte = l_byte[:-1]
            z += 1

    for p in range(len(verify)):
        if verify[p][0] == verify[p][1]:
            pass
        else:
            return False

    return True


def main():
    message = 'Hello!'.encode('utf-8')
    pub, priv = generate_lamport_keys()
    signature = sign_message(message, priv)
    print('Calculated Signature of message')
    res = verify_message_signature(signature, message, pub)
    print('Valid' if res else 'Not valid')


main()
