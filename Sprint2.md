**--API Documentation--**

Server: localhost:3000

**Available end-points**

  - /headsets
  - /headsets/{name}
  - /headsets/{name}/disable
  - /headsets/{name}/enable


**Requests**

- GET /headsets

    Lists all of the headsets

    Response:
            
      [
        {
          "ID": 1,
          "CreatedAt": "2023-02-23T02:45:19-05:00",
          "UpdatedAt": "2023-02-23T02:49:20-05:00",
          "DeletedAt": null,
          "name": "OculusRift",
          "price": 120,
          "wireless": true,
          "rating": 5,
          "description": "Description here",
          "status": false
        },
        {
          "ID": 2,
          "CreatedAt": "2023-02-23T02:52:19-05:00",
          "UpdatedAt": "2023-02-23T02:46:20-05:00",
          "DeletedAt": null,
          "name": "Steam VR",
          "price": 200,
          "wireless": false,
          "rating": 4,
          "description": "Description here",
          "status": false
        }
      ]

- GET /headsets/{name}

    Return a specific headset
    
    Response (/headsets/OculusRift): 
    
      [
        {
          "ID": 1,
          "CreatedAt": "2023-02-23T02:45:19-05:00",
          "UpdatedAt": "2023-02-23T02:49:20-05:00",
          "DeletedAt": null,
          "name": "OculusRift",
          "price": 120,
          "wireless": true,
          "rating": 5,
          "description": "Description here",
          "status": false
        }
      ]

- POST /headsets

    Creates a new headset in the database
    
    Parameters:
      
      name(string): Name of the headset
      price(int): How much the headset costs
      wireless(boolean): If the headset is wireless or not
      rating(int): 1 to 5 star rating of headset
      description(string): Description of the headset
      status(boolean): Is the headset available
    
    Response:
    
      {
        "message": "Successfully added headset!"
      }

- PUT /headsets/{name}

    Update an existing headset
    
    Parameters:
      
      name(string): Name of the headset
      price(int): How much the headset costs
      wireless(boolean): If the headset is wireless or not
      rating(int): 1 to 5 star rating of headset
      description(string): Description of the headset

    Response:

      {
        "message": "Successfully updated headset!"
      }

- PUT /headsets/{name}/enable
    
    Changes headsets status to true
    
    Response:

      {
        "message": "Successfully enabled headset!"
      }

- PUT /headsets/{name}/disable
    
    Changes headsets status to false
    
    Response:

      {
        "message": "Successfully disabled headset!"
      }

- DELETE /headsets/{name}
  
    Deletes an existing headset
    
    Response:

      {
        "message": "Successfully deleted headset!"
      }


      
    
    

    
      
    
      
      
    
      
      
        
      
      
     









