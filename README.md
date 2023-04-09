
# Overview

## What is it?
Custom rolled  replacement for Heimdall and Homer.
- https://github.com/linuxserver/Heimdall
- https://github.com/bastienwirtz/Homer

## Why
- It is actually something that I will use
- Good oportunity to get more experice developing a REST api in Go
- Get experience developing a website that isn't just a SPA

# Requirements

Website that allows user to create and organize Application Launchers. An Application Launcher is a clickable tile
that, when clicked, opens a new browser page based on how the Application Launcher is configured. Application Launchers
can be organized by adding them to Groups. Additionally, Application Laucnhers should be mutable and removeable. 

Groups must:
- have a mutable list of Application Launchers
- have a title describing what the group contains
- have an icon to make the group easier to recognize

Application Launchers must:
- have a title describing what the application launcher launches
- have an icon to make the application launcher easier to recognize
- belong to a single group

# Design
## Icons
- must be the name of the google material icon
## Primary Resources
```
Group {
  ID
  Title
  Icon
  ApplicationLauncher[]
  CreatedAt
  UpdatedAt
  DeletedAt
}
```

```
ApplicationLauncher {
  ID
  Title
  Icon
  ApplicationUrl
  GroupID
  CreatedAt
  UpdatedAt
  DeletedAt
}
```
## API
### POST - /v1/applicationLaunchers
Request:
```
{
  "title": string,
  "icon": svg,
  "application_url": string,
  "group_id": number
}
```

Response:
```
{
  "id": string,
  "title": string,
  "icon": string,
  "application_url": string,
  "group_id": number,
  "created_at": date,
  "updated_at": date
}
```
Errors:
- 400
- 422
- 500

### PATCH - /v1/applicationLaunchers/:id

### POST - /v1/groups

### GET - /v1/groups

### PATCH - /v1/groups/:id

