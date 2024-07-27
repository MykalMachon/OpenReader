# Notes 

This file is meant to be notes for what I need to do this in repo and general notes on what I need to build 

## Core services

- [ ] Ingest
  - [ ] system that on startup ingests books 
  - [ ] read in books folder if it exists, ask for one if there isn't one, or create one 
  - [ ] create books items with metadata for each book in a `books` table
- [ ] Reader
  - [ ] system that is responsible for acting on books 
  - [ ] open/close books 
  - [ ] read metadata about books 
  - [ ] extract text and images about books 
  - [ ] allow you to move forward and backward in a book
  - [ ] keep track of where the user is in a book
  - [ ] keep track of the status of a book (new, reading, read)

## Command Line Interface

This system interacts with the [core services](#core-services) to allow the user to read books from the command line. 

## REST API

This would be a simple web service that interacts with [core serivces](#core-services) and allow users to request book information and build their own abstraction on top of the core services.

## E-Reader 

This is the ultimate *goal* of the project, and would be an actual embedded go program that would be installed on hardware.