# Golang RPG game server

## Background
I created this project to implement separation of attention in a simple CRUD application.
## Overview
RPG game servers are simple server applications that are used to manage administrative needs in an RPG game.
To make the code in this application simpler and more maintainable, I divided the concern of this application into several layers of logic:
   - Entity, a representation of the existing business models in this application
   - Repository, a collection of entities, serves as a gateway to access and manipulate data.
   - Service, a place where the business logic of the entity is defined.
   - Delivery, the gateway for users to execute the related flow business.
## Golang library stack
 #### **ORM**:
   - [gorm](gorm.io/gorm)
 #### **Data transfer object mapping**:
   - [devfeel automapper](github.com/devfeel/mapper)  
 #### **Microservice kit and delivery transport**:
   - [go-chi](github.com/go-chi/chi)
   - [go-kit](github.com/go-kit/kit)
 #### **utility**:
   - [libra](github.com/haritsfahreza/libra)
 #### **Testing**:
   - [apitest](github.com/steinfletcher/apitest)
   - [testify](github.com/stretchr/testify)  
