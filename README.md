# SHIFT Project - Team 2

<p align="center">
  <img src="https://github.com/WomenPlusPlus/deploy-impact-23-shift-2/blob/main/docs/logo.png"/ alt="Shift2">
</p>

## :bookmark_tabs: Table of Contents
<!-- TOC -->
* [Our Solution](#bulb-our-solution)
* [Site Map](#dart-site-map)
* [Designs](#art-designs)
* [Tech Stack](#toolbox-tech-stack)
* [Architecture](#sparkles-architecture)
* [Database](#game_die-database)
* [Third-party libraries](#books-third-party-libraries)
* [License file](#memo-license-file)
* [Getting started](#computer-getting-started)
* [Team members](#woman_technologist-team-members-man_technologist)
* [Further development](#pushpin-further-development)
<!-- TOC -->

## :bulb: Our Solution
* TODO
* Describe it
* ...
* Matching algorithm:
  * A matching algorithm for candidates to recommend jobs and for companies to recommend candidates for jobs. It was not implemented due to time and capacity constraints, however, the idea can be found in [matching.md file](https://github.com/WomenPlusPlus/deploy-impact-23-shift-2/blob/main/docs/matching.md). 

## :dart: Site Map
* [The site map design](https://github.com/WomenPlusPlus/deploy-impact-23-shift-2/blob/main/docs/site_map.png)

## :art: Designs
* The designs were created using [Figma](https://www.figma.com/file/3BlcYSbbfCmx8oc5XdCKUN/Shift?type=design&node-id=1-1428&mode=design&t=UaqFM0xV7kPHePNj-0).
 
## :toolbox: Tech Stack
| Backend     | Frontend | Database | Cloud Storage | Deployments |
|-------------|----------|----------|---------------|-------------|
| Golang      |Angular   |PostgreSQL| GCP           | GCP         |
| Gorilla mux |DaisyUI   |Auth0     |               | Vercel      |
| SQLX        |NgRX      |          |               |             |

## :sparkles: Architecture
* [The architecture diagram](https://github.com/WomenPlusPlus/deploy-impact-23-shift-2/blob/main/docs/architecture.jpg)

## :game_die: Database
* [The database schema](https://miro.com/app/board/uXjVNfUchWk=/?share_link_id=4352832909)

## :books: Third-party libraries
For the Front-End: 
* [Tailwind CSS framework](https://tailwindcss.com/) 
* [DaisyUI - component library for Tailwind CSS](https://daisyui.com/)
* [RxJS - Reactive Extensions Library for JavaScript](  https://rxjs.dev/)
* [NGRX - Reactive State for Angular](https://ngrx.io/) 

## :memo: License file
* [GNU General Public License v3.0](https://github.com/WomenPlusPlus/deploy-impact-23-shift-2/blob/main/LICENSE)  

## :computer: Getting started
This project has a deployed version, which you can view [here](https://shift2-deployimpact.vercel.app/).

You can also run it locally, using the following instructions:

### Back-End Prerequisites:
* [Go](https://go.dev/doc/install)
* [Migrate for go](https://github.com/golang-migrate/migrate)
* [Docker](https://docs.docker.com/get-docker/)

### Front-End Prerequisites:
* [Node.js](https://nodejs.org/en)
* [npm package manager](https://docs.npmjs.com/downloading-and-installing-node-js-and-npm) (which is installed with Node.js by default)
* [Angular CLI](https://angular.io/guide/setup-local#install-the-angular-cli)

### Instructions

* Clone the repository on your computer. You can find instructions for that [here](https://docs.github.com/en/get-started/getting-started-with-git/about-remote-repositories#cloning-with-ssh-urls)
* Go the src folder of the project using `cd src`
* Install the necessary packages using `npm i`
* For the Back-End:
  * Go to the src/backend folder using `cd backend`
  * `SET POSTGRESQL_URL='postgres://postgres:shift2023@0.0.0.0:5432/postgres?sslmode=disable'`
  * `make DATA_PATH=/any/path/to/data docker-run-db` where you can specify any path on your computer 
  * run `migrate -database ${POSTGRESQL_URL} -path internal/db/migrations up` only the 1st time, to get all the changes from the db
  * `make run`
  * Now the backend is running on [http://localhost:8080/](http://localhost:8080/)
* For the Front-End:
  * In another terminal, on the src folder run `npm i --prefix frontend`
  * Go to the src/frontend folder using `cd frontend`
  * run `npm start`
* Now you can visit [http://localhost:4200/](http://localhost:4200/) on your browser to view the SHIFT website.

### Note: 

The Front-End services are performing HTTP Requests on our [deployed back-end](https://shift2-deployimpact.xyz). If you want to perform HTTP Requests to the local back-end, make sure to change the **API_BASE_URL** in the file *src/frontend/src/environments/environment.ts* to 'http://localhost:8080'

In order to view the content of our website, you are required to login. 

We are using [Auth0](https://auth0.com/) for User Authentication.

[Here](https://github.com/WomenPlusPlus/deploy-impact-23-shift-2/blob/main/docs/credentials.md) are some demo users you can use.

## :woman_technologist: Team members :man_technologist:
| Name                                                                         | Role            | GitHub                                            |
|------------------------------------------------------------------------------|-----------------|---------------------------------------------------|
| [Adamantia Milia](https://www.linkedin.com/in/adamantia-milia/)              | Frontend        | [mandyjker](https://github.com/mandyjker)         |
| [Adrianna Zielińska](https://www.linkedin.com/in/adriannazielinska/)         | Backend, Data Science    | [adriannaziel](https://github.com/adriannaziel)   |
| [Bianca Alves](https://www.linkedin.com/in/biancaalves/)                     | Project Manager | [biancamnalves](https://github.com/biancamnalves) |
| [Hannah Rüfenacht](https://www.linkedin.com/in/hannahrufenacht/)             | UX/UI           | [hrrenee15](https://github.com/hrrenee15)         |
| [Istvan Zsigmond](https://www.linkedin.com/in/istvan-zsigmond/)              | Backend         | [istvzsig](https://github.com/istvzsig)           |
| [João Rodrigues](https://www.linkedin.com/in/jo%C3%A3o-rodrigues-84268613b/) | Fullstack       | [jotar910](https://github.com/jotar910)           |

## :pushpin: Further development
* For any further development, please contact [João Rodrigues](https://www.linkedin.com/in/jo%C3%A3o-rodrigues-84268613b/).
