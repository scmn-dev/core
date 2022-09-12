# $PASSWORD_MANAGER_NAME Core

> **$PASSWORD_MANAGER_NAME Core** is the Backend Infrastructure for $PASSWORD_MANAGER_NAME products.

**`$PASSWORD_MANAGER_NAME core` prefers to run in unix 🐧 machine**

## $PASSWORD_MANAGER_NAME Products

**1. $PASSWORD_MANAGER_NAME CLI.**

**2. $PASSWORD_MANAGER_NAME Hub.**

**3. $PASSWORD_MANAGER_NAME Web Extension.**

## Security

**1. $PASSWORD*MANAGER_NAME uses The Advanced Encryption Standard (AES) encryption algorithm with Galois/Counter Mode (GCM) symmetric-key cryptographic mode. Passwords encrypted with AES can only be decrypted with the passphrase defined in the \_config.yml* file.**

**2. Endpoints are protected with security middlewares against attacks like XSS.**

**3. Against SQL injection, $PASSWORD_MANAGER_NAME uses Gorm package to handle database queries which clears all queries.**

**4. There is rate limiter for signin attempts against brute force attacks.**

## Secman

> 👊 Human-friendly and amazing open-source secrets manager.
<p align="center">
  <img src="https://raw.githubusercontent.com/scmn-dev/.github/main/assets/secman.svg" />
</p>

<p align="center">
  <a href="https://secman.dev"><strong>secman.dev</strong></a>
</p>

> `secman` is a password manager can store, retrieves, generates, and synchronizes passwords, The most important difference is secman is not GPG cored. Instead, it uses a master password to securely store your passwords. and you can easily manage your passwords from everywhere with **Secman Cloud** 😉.
