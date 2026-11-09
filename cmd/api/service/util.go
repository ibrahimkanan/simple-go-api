package service

import(

)

type Data struct{
	UserId int
	Id int 
	Title string
	Body string
}

type Payload struct{
	Data []Data
}