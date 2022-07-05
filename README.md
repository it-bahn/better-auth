<h1 align="center">🎉 better-auth🎉</h1>
This project is a SaaS running in a serverless envoirment in google cloud platform. It is a authentification service written in go lang and built with docker 
<br/>

<div align="center">       
<img src="https://img.shields.io/github/workflow/status/techonomylabs/better-auth/Docker?label=GCP%20CLOUD%20RUN&style=for-the-badge"/>

[![Docker](https://github.com/techonomylabs/better-auth/actions/workflows/deploy-to-cloud-run.yml/badge.svg)](https://github.com/techonomylabs/better-auth/actions/workflows/deploy-to-cloud-run.yml)

<img src="https://img.shields.io/github/license/techonomylabs/better-auth" />
<a href="https://github.com/techonomylabs/better-auth/issues">
<img src="https://img.shields.io/github/issues/techonomylabs/better-auth" />
</a>
<img src="https://img.shields.io/github/languages/count/techonomylabs/better-auth?style=flat-square"/>

</div>

## What's Inside

- Pure Golang
- RestApi only using standard go packages
- Mongo Golang Driver 
- [crypto/bcrypt](https://golang.org/x/crypto/bcrypt) for password hashing
- Boilerplate golang mongo CRUD 
- Dockerfile
- CI/CD to Google Cloud Run

## How Robust is This Service
- Not using a Relational DB, but a MongoDB
- Not using a JWT,but session management system
- Can not register with the same email twice
- All fields are validated using Regex expressions
- All HTTP request errors are handled using a custom error handler
- All HTTP request are logged using a custom logger
- All User Endpoints are inaccessible if the user logs out

## Live demo

[Live Api Endpoints](https://techonomy-labs-o2k3wv2fsq-uc.a.run.app/api/v1/)

[Documentation](https://documenter.getpostman.com/view/21725756/UzJHRdXy)

[Join The Support Team](https://app.getpostman.com/join-team?invite_code=40a4a16810b9f88648390722e98b8e79)

