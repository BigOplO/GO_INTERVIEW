# Introduction
1.Using contentfulAPI to get the data and store in postgreSQL by using CLI(Cobra)  
2.Create the graphQL server to operate the data from postgreSQL  

# Setting
- Enter the psql and create the postgreSQL database
```sh
  psql postgres
  CREATE DATABASE breads;
```
- DATABASE setting
```sh
host:localhost  
port:5432(default)  
#user: You have to change the username in line15 in db.go to your username.  
SSL:disable
```
# Operating
- Build the project
```sh
go build -o bread-info
```

- Use contentful to get data and store into psql
```sh
./bread-info contentful
```

- Run the graphQL server and operate the graphQL
```sh
./bread-info server
```
Open localhost:8080 in your computer and write the description in left side. 
1.GetAllBreads
```sh
query GetAllBread{
  bread{
      id
      name
      createdAt
}
}
```
2.GetBread
```sh
query GetBread{
  bread(id: "bread_id"){
      id
      name
      createdAt
}
}
```

