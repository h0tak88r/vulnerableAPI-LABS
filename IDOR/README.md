# IDOR Lab

This lab is designed to demonstrate and help users understand Insecure Direct Object Reference (IDOR) vulnerabilities in web servers/APIs. The lab is implemented in Go programming language to create a simple vulnerable API server.


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

# Testing 

![Screenshot from 2024-03-22 17-38-20](https://github.com/h0tak88r/vulnerableAPI-LABS/assets/108616378/d9edb020-c4f7-4108-a56a-ec5b351e4adb)

![Screenshot from 2024-03-22 17-38-36](https://github.com/h0tak88r/vulnerableAPI-LABS/assets/108616378/989ee42f-0e40-4f5f-b67a-b233348833fd)

![Screenshot from 2024-03-22 17-38-49](https://github.com/h0tak88r/vulnerableAPI-LABS/assets/108616378/8b5c95af-6cdd-4aec-aa88-5201555a3a12)
