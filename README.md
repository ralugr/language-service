# Language service
Provides a list of banned words.

## Client guide
#### There are 2 ways to get access to the list:
1. Use the **GET localhost:8080/list** to retrieve the list on need.
2. Use **POST localhost:8080/subscribe** to subscribe with a token (a unique key) and a URL. Once the list is updated, you will receive it through the specified URL.

#### Models:
- To subscribe, please make sure the body has the following keys in a JSON format:   
`token(string)` `url(string)`  

Example:  
> POST localhost:8080/subscribe  
>> Body  
 {  
    "token": "123456%^sdf5ns*(wrh2we4fb28341837",  
    "url": "http://localhost:8081/notify"  
  }  

- You will be notified at the given URL in the following JSON format:   
`token(string)` `words(string)`


## Design decisions
The language service follows the implementation of the observer pattern adapted to microservices. To achieve this behavior, the service stores a list of subscribers, each having a token and a URL. The subscriber will be notified using POST on the given URL as soon as the list updates. The token can be used by the client to validate if the POST was indeed made by the language service.

The service also provides a route for updating the list **POST localhost:8080/update_list**. This is used to simulate an update to the banned words list. Used for testing purposes only. 

Example:  
> POST localhost:8080/update_list  
>> Body  
 ["cat", "dog","bird", "zebra"]
 
 
## Limitations
- The client must have a web api and implement POST on the given URL.



