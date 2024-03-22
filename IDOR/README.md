# API IDOR-LAB

## Endpoints

- `/api/user`: Endpoint to interact with user data.

## Vulnerabilities

### Insecure Direct Object Reference (IDOR)

- **Description**: This lab contains an IDOR vulnerability that allows users to access or manipulate resources (user data) they are not authorized to access.
- **Exploitation**: Users can exploit this vulnerability by tampering with the `id` parameter in the URL when making GET requests to `/api/user`. They can access other users' data by changing the value of `id`.
- **Example**: To retrieve user data with ID 0, make a GET request to `/api/user?id=0`. An attacker can change the ID to access other users' data.

## Learning Objectives

- Understand the concept of Insecure Direct Object Reference (IDOR) vulnerabilities.
- Learn how to identify and exploit IDOR vulnerabilities in web servers/APIs.
- Learn best practices for secure API design and authorization.

## Disclaimer

This lab is for educational purposes only. Do not use it to perform any unauthorized or malicious activities. The creators of this lab are not responsible for any misuse or damage caused.
