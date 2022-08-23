# autoscaling-functions

## bcrypt

Designed to stress the CPU.

```bash
export PLAINTEXT=test
export HASHED=$(curl -s http://127.0.0.1:8080/function/bcrypt \
--data $PLAINTEXT)
echo Hash: $HASHED

echo Hash validates against plaintext? $(curl -s http://127.0.0.1:8080/function/bcrypt/decode --data "$PLAINTEXT $HASHED")
```

## sleep

Configure via `sleep_duration` environment variable.

Ideal for testing `capacity` based scaling.

## cows

Fast Node.js function ideal for testing RPS-based scaling.

